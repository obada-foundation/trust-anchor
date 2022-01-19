package api

import (
	"crypto/tls"
	"net/http"
	"net/url"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"golang.org/x/crypto/acme/autocert"
)

// sslMode defines ssl mode for rest server
type sslMode int8

const (
	// None defines to run http server only
	None sslMode = iota

	// Static defines to run both https and http server. Redirect http to https
	Static

	// Auto defines to run both https and http server. Redirect http to https. Https server with acme support
	Auto
)

// SSLConfig holds all ssl params for rest server
type SSLConfig struct {
	Type         sslMode
	Port         int
	Cert         string
	Key          string
	ACMELocation string
	ACMEEmail    string
}

// httpToHTTPSRouter creates new router which does redirect from http to https server  with default middlewares. Used in 'static' ssl mode.
func (s *Rest) httpToHTTPSRouter() chi.Router {
	s.Logger.Printf("[DEBUG] create https-to-http redirect routes")
	router := chi.NewRouter()

	router.Handle("/*", s.redirectHandler())
	return router
}

// nolint
func (s *Rest) redirectHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		newURL := s.TrustAnchorURL + r.URL.Path
		if r.URL.RawQuery != "" {
			newURL += "?" + r.URL.RawQuery
		}
		http.Redirect(w, r, newURL, http.StatusTemporaryRedirect)
	})
}

func (s *Rest) makeAutocertManager() *autocert.Manager {
	return &autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		Cache:      autocert.DirCache(s.SSLConfig.ACMELocation),
		HostPolicy: autocert.HostWhitelist(s.getRemarkHost()),
		Email:      s.SSLConfig.ACMEEmail,
	}
}

// makeHTTPSAutoCertServer makes https server with autocert mode (LE support)
func (s *Rest) makeHTTPSAutocertServer(address string, port int, router http.Handler, m *autocert.Manager) *http.Server {
	server := s.makeHTTPServer(address, port, router)
	cfg := s.makeTLSConfig()
	cfg.GetCertificate = m.GetCertificate
	server.TLSConfig = cfg
	return server
}

// makeHTTPSServer makes https server for static mode
func (s *Rest) makeHTTPSServer(address string, port int, router http.Handler) *http.Server {
	server := s.makeHTTPServer(address, port, router)
	server.TLSConfig = s.makeTLSConfig()
	return server
}

// httpChallengeRouter creates new router which performs ACME "http-01" challenge response
// with default middlewares. This part is necessary to obtain certificate from LE.
// If it receives not a acme challenge it performs redirect to https server.
// Used in 'auto' ssl mode.
func (s *Rest) httpChallengeRouter(m *autocert.Manager) chi.Router {
	s.Logger.Printf("[DEBUG] create http-challenge routes")
	router := chi.NewRouter()
	router.Use(middleware.Throttle(1000), middleware.RealIP, Recoverer(s.Logger))
	router.Use(middleware.Throttle(1000), middleware.Timeout(60*time.Second))

	router.Handle("/*", m.HTTPHandler(s.redirectHandler()))
	return router
}

// getHost returns hostname for remark server.
// For example for trustAnchorURL https://trust-anchor.obada.com:443 it should return trust-anchor.obada.com
func (s *Rest) getRemarkHost() string {
	u, err := url.Parse(s.TrustAnchorURL)
	if err != nil {
		return ""
	}
	return u.Hostname()
}

func (s *Rest) makeTLSConfig() *tls.Config {
	return &tls.Config{
		PreferServerCipherSuites: true,
		CipherSuites: []uint16{
			tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
			// tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305,
			// tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305,
			tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
			tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
			// tls.TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA,
		},
		MinVersion: tls.VersionTLS12,
		CurvePreferences: []tls.CurveID{
			tls.CurveP256,
			tls.X25519,
			tls.CurveP384,
		},
	}
}
