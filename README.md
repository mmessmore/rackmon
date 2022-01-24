# rackmon

```
./rackmon

This monitors the temperature of a Dallas 1-Wire temperature
sensor and is intended to work with the telegraf execd plugin.

This runs as a daemon outputting Graphite formatted metrics
to STDOUT when it receives a SIGUSR1.  It will exit on
SIGTERM or SIGINT (^C).

USAGE

  ./rackmon [metric_prefix...]
  ./rackmon -h

ARGUMENTS
  metric_prefix	all arguments are joined by a '.' as the metric name

OPTIONS
  -h	This handy usage message

```
