/*
Copyright © 2022 Yoshino-s

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/
package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/yoshino-s/ElectronInjector/injector"
)

// rootCmd represents the base command when called without any subcommands
var (
	inject  string
	rootCmd = &cobra.Command{
		Use:   "ElectronInjector",
		Short: "Another injector on encrypted electron app",
		// Uncomment the following line if your bare application
		// has an action associated with it:
		Run: func(cmd *cobra.Command, args []string) {
			var prog string
			if len(args) == 0 {
				prog = "./main.node"
			} else {
				prog = strings.Join(args, " ")
			}
			if err := injector.Inject(prog, inject); err != nil {
				fmt.Println(err)
				fmt.Print("Press enter to exit...")
				input := bufio.NewScanner(os.Stdin)
				input.Scan()
			}
		},
	}
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.MousetrapHelpText = ""
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVarP(&inject, "inject", "i", "crack", "Inject file (Internal Payload: hello|dump|crack)")
}
