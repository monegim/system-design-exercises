package main

import (
	"flag"
	"fmt"
)

func main() {
	debug := flag.Bool("debug", false, "enable debug logging")
	quiet := flag.Bool("quiet", false, "do not log informational messages (takes precedence over debug)")
	json := flag.Bool("json", false, "enable JSON logging")
	test := flag.Bool("test", false, "test crontab (does not run jobs)")
	prometheusListen := flag.String(
		"prometheus-listen-address",
		"",
		fmt.Sprintf("give a valid ip[:port] address to expose Prometheus metrics at /metrics (port defaults to %s), "+
		"use 0.0.0.0 for all network interfaces.", prometheus_metrics.DefaultPort),
	)
	splitLogs := flag.Bool("split-logs", false, "split log output into stdout/stderr")
	passthroughLogs := flag.Bool("passthrough-logs", false, "passthrough logs from commands, do not wrap them in Supercronic logging")
	sentry := flag.String("sentry-dsn", "", "enable Sentry error logging, using provided DSN")
	sentryAlias := flag.String("sentryDsn", "", "alias for sentry-dsn")
	overlapping := flag.Bool("overlapping", false, "enable tasks overlapping")
	flag.Parse()
}
