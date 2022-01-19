package cmd

import (
	"log"

	"github.com/dgraph-io/badger/v3"
)

// CommonOptionsCommander extends flags.Commander with SetCommon
// All commands should implement this interfaces
type CommonOptionsCommander interface {
	SetCommon(commonOpts CommonOpts)
	Execute(args []string) error
}

// CommonOpts sets externally from main, shared across all commands
type CommonOpts struct {
	Revision string
	Logger   *log.Logger
	DB       *badger.DB
}

// SetCommon satisfies CommonOptionsCommander interface and sets common option fields
// The method called by main for each command
func (c *CommonOpts) SetCommon(commonOpts CommonOpts) {
	c.Revision = commonOpts.Revision
	c.Logger = commonOpts.Logger
	c.DB = commonOpts.DB
}
