// Copyright (c) 2015 Pagoda Box Inc
//
// This Source Code Form is subject to the terms of the Mozilla Public License, v.
// 2.0. If a copy of the MPL was not distributed with this file, You can obtain one
// at http://mozilla.org/MPL/2.0/.
//

package commands

//
import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"

	"github.com/pagodabox/nanobox-cli/config"
	"github.com/pagodabox/nanobox-cli/util"
	"github.com/pagodabox/nanobox-golang-stylish"
)

//
var destroyCmd = &cobra.Command{
	Use:   "destroy",
	Short: "Destroys the nanobox VM",
	Long: `
Description:
  Destroys the nanobox VM by issuing a "vagrant destroy"`,

	Run: nanoDestroy,
}

// nanoDestroy
func nanoDestroy(ccmd *cobra.Command, args []string) {

	// if the command is being run with the "remove" flag, it means an entry needs
	// to be removed from the hosts file and execution yielded back to the parent
	if len(args) > 0 && args[0] == "remove-entry" {
		util.RemoveDevDomain()
		os.Exit(0)
	}

	//
	// destroy the vm; this needs to happen before cleaning up the app to ensure
	// there is a Vagrantfile to run the command with (otherwise it will just get
	// re-created)
	if err := util.RunVagrantCommand(exec.Command("vagrant", "destroy", "--force")); err != nil {

		// dont care if the project no longer exists... thats what we're doing anyway
		if err != err.(*os.PathError) {
			config.Fatal("[commands/destroy] util.RunVagrantCommand() failed", err.Error())
		}
	}

	// remove app; this needs to happen after the VM is destroyed so that the app
	// isn't just created again upon running the vagrant command
	fmt.Printf(stylish.Bullet("Deleting nanobox files (%s)", config.AppDir))
	if err := os.RemoveAll(config.AppDir); err != nil {
		config.Fatal("[commands/destroy] os.RemoveAll() failed", err.Error())
	}

	// attempt to remove the entry regardless of whether its there or not
	util.SudoExec("destroy remove-entry", fmt.Sprintf("Removing %s domain from /etc/hosts", config.Nanofile.Domain))
}
