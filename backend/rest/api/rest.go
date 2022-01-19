package api

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth/v5"
	"github.com/obada-foundation/trust-anchor/ta"
	"github.com/obada-foundation/trust-anchor/ta/compliance"
)

const hardBodyLimit = 1024 * 64 // limit size of body

// Rest is a rest access server
type Rest struct {
	Version string

	AuthTokenSvc  *jwtauth.JWTAuth
	TaTokenSvc    *ta.TokenService
	ComplianceSvc *compliance.Service

	Logger         *log.Logger
	TrustAnchorURL string
	SSLConfig      SSLConfig
	httpServer     *http.Server
	httpsServer    *http.Server
	lock           sync.Mutex

	publicRest  public
	privateRest private
}

// Run the lister and request's router, activate rest server
func (s *Rest) Run(address string, port int) {
	if address == "*" {
		address = ""
	}

	switch s.SSLConfig.Type {
	case None:
		s.Logger.Printf("[INFO] activate http rest server on %s:%d", address, port)

		s.lock.Lock()
		s.httpServer = s.makeHTTPServer(address, port, s.routes())
		s.httpServer.ErrorLog = s.Logger
		s.lock.Unlock()

		err := s.httpServer.ListenAndServe()
		s.Logger.Printf("[WARN] http server terminated, %s", err)
	case Static:
		s.Logger.Printf("[INFO] activate https server mode on %s:%d", address, port)

		s.lock.Lock()
		s.httpsServer = s.makeHTTPSServer(address, port, s.routes())
		s.httpsServer.ErrorLog = s.Logger
		s.lock.Unlock()

		err := s.httpsServer.ListenAndServeTLS(s.SSLConfig.Cert, s.SSLConfig.Key)
		s.Logger.Printf("[WARN] https server terminated, %s", err)
	case Auto:
		s.Logger.Printf("[INFO] activate https server in 'auto' mode on %s:%d", address, s.SSLConfig.Port)

		m := s.makeAutocertManager()
		s.lock.Lock()
		s.httpsServer = s.makeHTTPSAutocertServer(address, s.SSLConfig.Port, s.routes(), m)
		s.httpsServer.ErrorLog = s.Logger

		s.httpServer = s.makeHTTPServer(address, port, s.httpChallengeRouter(m))
		s.httpsServer.ErrorLog = s.Logger

		s.lock.Unlock()

		go func() {
			s.Logger.Printf("[INFO] activate http challenge server on port %d", port)

			err := s.httpServer.ListenAndServe()
			s.Logger.Printf("[WARN] http challenge server terminated, %s", err)
		}()

		err := s.httpsServer.ListenAndServeTLS("", "")
		s.Logger.Printf("[WARN] https server terminated, %s", err)
	}
}

func (s *Rest) routes() chi.Router {
	router := chi.NewRouter()

	mdLogger := middleware.RequestLogger(&middleware.DefaultLogFormatter{Logger: s.Logger, NoColor: false})

	router.Use(mdLogger)
	router.Use(middleware.Throttle(1000), middleware.RealIP, Recoverer(s.Logger))
	router.Use(AppInfo("Trust Anchor", "OBADA Foundation", s.Version), Ping)

	s.publicRest, s.privateRest = s.handlerGroups()

	// api routes
	router.Route("/api/v1", func(rapi chi.Router) {
		rapi.Group(func(rpriv chi.Router) {
			rpriv.Use(jwtauth.Verifier(s.AuthTokenSvc))
			rpriv.Use(jwtauth.Authenticator)

			rpriv.Post("/nft/register", s.privateRest.registerNFT)
		})

		rapi.Group(func(rpub chi.Router) {
			rpub.Post("/user/register", s.publicRest.registerUser)
			rpub.Get("/token", s.publicRest.token)
			rpub.Get("/verify", s.publicRest.verifyToken)
		})
	})

	return router
}

func (s *Rest) handlerGroups() (public, private) {
	pubGrp := public{
		logger:        s.Logger,
		complianceSvc: s.ComplianceSvc,
		authTokenSvc:  s.AuthTokenSvc,
	}

	privGrp := private{
		logger:     s.Logger,
		taTokenSvc: s.TaTokenSvc,
	}

	return pubGrp, privGrp
}

func (s *Rest) makeHTTPServer(address string, port int, router http.Handler) *http.Server {
	return &http.Server{
		Addr:              fmt.Sprintf("%s:%d", address, port),
		Handler:           router,
		ReadHeaderTimeout: 5 * time.Second,
		// WriteTimeout:      120 * time.Second, // TODO: such a long timeout needed for blocking export (backup) request
		IdleTimeout: 30 * time.Second,
	}
}

// Shutdown rest http server
func (s *Rest) Shutdown() {
	s.Logger.Print("[WARN] shutdown rest server")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	s.lock.Lock()

	if s.httpServer != nil {
		if err := s.httpServer.Shutdown(ctx); err != nil {
			s.Logger.Printf("[DEBUG] http shutdown error, %s", err)
		}
		s.Logger.Print("[DEBUG] shutdown http server completed")
	}

	if s.httpsServer != nil {
		s.Logger.Print("[WARN] shutdown https server")
		if err := s.httpsServer.Shutdown(ctx); err != nil {
			s.Logger.Printf("[DEBUG] https shutdown error, %s", err)
		}
		s.Logger.Print("[DEBUG] shutdown https server completed")
	}
	s.lock.Unlock()
}
