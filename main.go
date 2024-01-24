package palworld_monitor

import (
	"log/slog"
	"os"
	palrcon "palworld_monitor/pkg/palcron"
	"time"
)

var (
	rconEndpoint = "192.168.1.6:25575"
	rconPassword = "197210"

	intervalRaw = "5s"
	interval    time.Duration

	timeoutRaw = "1s"
	timeout    time.Duration
)

func init() {
	var err error

	if timeoutRaw == "" {
		timeoutRaw = "1s"
	}

	timeout, err = time.ParseDuration(timeoutRaw)
	if err != nil {
		slog.Error("failed to parse timeout", "error", err)
		os.Exit(1)
	}

	if intervalRaw == "" {
		intervalRaw = "5s"
	}

	interval, err = time.ParseDuration(intervalRaw)

	if err != nil {
		slog.Error("failed to parse interval", "error", err)
		os.Exit(1)
	}
}

func main() {
	palRCON := palrcon.NewPalRCON(rconEndpoint, rconPassword)
	palRCON.SetTimeout(timeout)

}
