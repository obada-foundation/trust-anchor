package cmd

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/obada-foundation/trust-anchor/rest/api"
)

type ServerCommand struct {
	Port    int      `long:"port" env:"TA_PORT" default:"80" description:"port"`
	Address string   `long:"address" env:"TA_ADDRESS" default:"" description:"listening address"`
	SSL     SSLGroup `group:"ssl" namespace:"ssl" env-namespace:"SSL"`

	CommonOpts
}

// SSLGroup defines options group for server ssl params
type SSLGroup struct {
	Type string `long:"type" env:"TYPE" description:"ssl support" choice:"acme" choice:"none" choice:"static" default:"none"` // nolint
	Cert string `long:"cert" env:"CERT" description:"path to cert.pem file"`
	Key  string `long:"key" env:"KEY" description:"path to key.pem file"`
}

// serverApp holds all active objects
type serverApp struct {
	*ServerCommand
	restSrv    *api.Rest
	terminated chan struct{}
}

// Execute is the entry point for "server" command, called by flag parser
func (s *ServerCommand) Execute(_ []string) error {
	s.Logger.Printf("[INFO] starting server on port %s:%d", s.Address, s.Port)

	ctx, cancel := context.WithCancel(context.Background())

	go func() { // catch signal and invoke graceful termination
		stop := make(chan os.Signal, 1)
		signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
		<-stop
		log.Printf("[WARN] interrupt signal")
		cancel()
	}()

	app := s.newServerApp()

	if err := app.run(ctx); err != nil {
		log.Printf("[ERROR] server terminated with error %+v", err)
		return err
	}

	return nil
}

// newServerApp prepares application and return it with all active parts
// doesn't start anything
func (s *ServerCommand) newServerApp() *serverApp {
	sslConfig := s.makeSSLConfig()

	srv := &api.Rest{
		Logger:    s.Logger,
		SSLConfig: sslConfig,
	}

	return &serverApp{
		ServerCommand: s,
		restSrv:       srv,
		terminated:    make(chan struct{}),
	}
}

func (s *ServerCommand) makeSSLConfig() (config api.SSLConfig) {
	switch s.SSL.Type {
	case "none":
		config.Type = api.None

	}

	config.Cert = s.SSL.Cert
	config.Key = s.SSL.Key

	return config
}

// Run all application objects
func (a *serverApp) run(ctx context.Context) error {
	go func() {
		// shutdown on context cancellation
		<-ctx.Done()
		log.Print("[INFO] shutdown initiated")
		a.restSrv.Shutdown()
	}()

	a.restSrv.Run(a.Address, a.Port)

	close(a.terminated)
	return nil
}

// Wait for application completion (termination)
func (a *serverApp) Wait() {
	<-a.terminated
}
