package main

import (
	"fmt"
	"log"
	"os"

	badger "github.com/dgraph-io/badger/v3"
	flags "github.com/jessevdk/go-flags"
	"github.com/obada-foundation/trust-anchor/cmd"
)

var revision = "unknown"

type opts struct {
	DBPath    string            `long:"db-path" description:"Path to database files" default:"/var/obada/trust-anchor/data"`
	ServerCmd cmd.ServerCommand `command:"server"`
}

func main() {
	fmt.Printf("TrustAnchor %s\n(c)OBADA 2022\n\n", revision)

	logger := log.New(os.Stdout, " :: TRUST-ANCHOR  :: ", log.LstdFlags|log.Lmicroseconds|log.Lshortfile)

	var o opts

	if err := run(logger, &o); err != nil {
		if flagsErr, ok := err.(*flags.Error); ok && flagsErr.Type == flags.ErrHelp {
			os.Exit(0)
		} else {
			logger.Println(err)
			os.Exit(1)
		}
	}

}

func run(logger *log.Logger, o *opts) error {
	p := flags.NewParser(o, flags.Default)
	p.CommandHandler = func(command flags.Commander, args []string) error {
		db, err := badger.Open(badger.DefaultOptions(o.DBPath))
		if err != nil {
			return err
		}
		defer db.Close()

		c := command.(cmd.CommonOptionsCommander)
		c.SetCommon(cmd.CommonOpts{
			Revision: revision,
			Logger:   logger,
			DB:       db,
		})

		err = c.Execute(args)
		if err != nil {
			logger.Printf("[ERROR] failed with %+v", err)
		}
		return err
	}

	if _, err := p.Parse(); err != nil {
		return err
	}

	return nil
}
