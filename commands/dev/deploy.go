package dev

import (
	"github.com/spf13/cobra"

	"github.com/nanobox-io/nanobox/processor"
	"github.com/nanobox-io/nanobox/util/print"
	"github.com/nanobox-io/nanobox/validate"
)

// DeployCmd ...
var DeployCmd = &cobra.Command{
	Use:   "deploy",
	Short: "Deploys a build package into your dev platform and starts all data services.",
	Long: ``,
	PreRun: validate.Requires("provider", "provider_up", "built", "dev_isup"),
	Run:    deployFn,
}

// deployFn ...
func deployFn(ccmd *cobra.Command, args []string) {
	print.OutputCommandErr(processor.Run("dev_deploy", processor.DefaultControl))
}