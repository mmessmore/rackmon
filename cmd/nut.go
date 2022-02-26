/*
Copyright © 2022 Michael Messmore <mike@messmore.org>

*/
package cmd

import (
	"github.com/mmessmore/rackmon/nut"
	"github.com/spf13/cobra"
)

// nutCmd represents the nut command
var nutCmd = &cobra.Command{
	Use:   "nut",
	Short: "NUT UPS client",
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
		defaultPrefix := makePrefix(defaultPrefixFormat, "nut", false)
		prefix, _ := cmd.Flags().GetString("prefix")

		// if they didn't change the default, set it back to do the
		// hostname formatting right
		// there's probably a better way to do this
		if prefix == defaultPrefix {
			prefix = defaultPrefixFormat
		}
		short, _ := cmd.Flags().GetBool("short")
		name, _ := cmd.Flags().GetString("upsname")

		prefix = makePrefix(prefix, "nut", short)
		nut.Run(prefix, name)
	},
}

func init() {
	rootCmd.AddCommand(nutCmd)
	defaultPrefix := makePrefix(defaultPrefixFormat, "nut", false)
	nutCmd.Flags().StringP("prefix", "p", defaultPrefix, "Metric prefix")
	nutCmd.Flags().BoolP("short", "s", false, "Use short hostname in prefix")
	nutCmd.Flags().StringP("upsname", "n", "main", "UPS name from NUT")
}
