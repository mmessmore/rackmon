/*
Copyright © 2022 Michael Messmore <mike@messmore.org>

*/
package cmd

import (
	"os"
	"strings"

	"github.com/mmessmore/rackmon/d1w"
	"github.com/spf13/cobra"
)

const defaultPrefixFormat = "HOSTNAME.sensors.d1w"

// d1wCmd represents the d1w command
var d1wCmd = &cobra.Command{
	Use:   "d1w",
	Short: "Dallas 1-Wire temp sensor",
	Long: `Runs via the execd plugin to Telegraf for a Dallas 1-Wire
temperature sensor.  This will emit graphite formatted metric on SIGUSR1
or SIGHUP.

Metric is HOSTNAME.sensors.d1w.DEVICE by default.  DEVICE will look like
28-000005895125.

By default HOSTNAME is the FQDN returned by 'uname -n'.  The --short flag
reduces this to the name up until the first '.'.  For example, a host named
'foo.example.com' would be reduced to 'foo'.

If --prefix is defined, "HOSTNAME" is replaced by the hostname or short
hostname if --short is specified.


Example Telegraf config:
[[inputs.execd]]
  command = ["rackmon", "d1w"]
  signal = "SIGUSR1"
  restart_delay = "10s"
  data_format = "graphite"

KNOWN BUGS
This only will output the first detected w1 device and will die tragically if
it's not a temperature sensor.

It should support multiples, allow mappings of device names to friendly names,
and skip non-temperature sensors.
`,
	Run: func(cmd *cobra.Command, args []string) {
		defaultPrefix := makePrefix(defaultPrefixFormat, false)
		prefix, _ := cmd.Flags().GetString("prefix")

		// if they didn't change the default, set it back to do the
		// hostname formatting right
		// there's probably a better way to do this
		if prefix == defaultPrefix {
			prefix = defaultPrefixFormat
		}
		short, _ := cmd.Flags().GetBool("short")

		prefix = makePrefix(prefix, short)
		d1w.Run(prefix)
	},
}

func init() {
	rootCmd.AddCommand(d1wCmd)
	defaultPrefix := makePrefix(defaultPrefixFormat, false)
	d1wCmd.Flags().StringP("prefix", "p", defaultPrefix, "Metric prefix")
	d1wCmd.Flags().BoolP("short", "s", false, "Use short hostname in prefix")
}

// makePrefix replaces HOSTNAME with the short or long name of the local host
// based on whether or not short is true
func makePrefix(prefix string, short bool) string {
	// This does not use Sprintf to avoid evil
	hostname, _ := os.Hostname()
	if short {
		hostname = strings.Split(hostname, ".")[0]
	}
	prefix = strings.ReplaceAll(prefix, "HOSTNAME", hostname)
	return prefix
}
