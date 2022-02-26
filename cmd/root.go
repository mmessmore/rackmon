/*
Copyright © 2022 Michael Messmore <mike@messmore.org>

*/
package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
)

const defaultPrefixFormat = "HOSTNAME.sensors"

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "rackmon",
	Short: "Telegraf integrations for measuring stuff for my SOHO rack.",
	Long: `A set of sensors to integrate with Telegraf's execd plugin.

All emit metrics in Graphite format on SIGHUP or SIGUSR1.

Currently this supports:
- d1w: Dallas 1-Wire temperature sensor via Linux kernel support
- nut: UPS metrics via NUT
`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.rackmon.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// makePrefix replaces HOSTNAME with the short or long name of the local host
// based on whether or not short is true
func makePrefix(prefix string, monitor string, short bool) string {
	// This does not use Sprintf to avoid evil
	hostname, _ := os.Hostname()
	if short {
		hostname = strings.Split(hostname, ".")[0]
	}
	prefix = strings.ReplaceAll(prefix, "HOSTNAME", hostname)
	prefix = prefix + "." + monitor
	return prefix
}
