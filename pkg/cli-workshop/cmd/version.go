package cmd

import (
	"fmt"
	"runtime"

	"github.com/spf13/cobra"
	"github.com/tehcyx/go-cli-workshop/pkg/cli-workshop/core"
)

//Version contains the binary version injected by the build system
var Version string

//GitCommit contains the git commit sha that the binary was built with, injected by the build system
var GitCommit string

//VersionOptions defines available options for the command
type VersionOptions struct {
	*core.Options
}

//NewVersionOptions creates options with default values
func NewVersionOptions(o *core.Options) *VersionOptions {
	return &VersionOptions{Options: o}
}

//NewVersionCmd creates a new version command
func NewVersionCmd(o *VersionOptions) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "version",
		Short: "Version of the workshop CLI",
		Long:  `Prints the version of workshop CLI`,
		RunE:  func(_ *cobra.Command, _ []string) error { return o.Run() },
	}

	return cmd
}

//Run runs the command
func (o *VersionOptions) Run() error {

	version := Version
	if version == "" {
		version = "N/A"
	}
	commit := GitCommit
	if commit == "" {
		commit = "N/A"
	}
	gover := runtime.Version()
	fmt.Printf("workshop CLI version: %s, commit: %s, go version: %s\n", version, commit, gover)

	return nil
}
