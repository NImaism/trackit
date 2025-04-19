package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/nimaism/trackit/internal/change"
	"github.com/nimaism/trackit/internal/config"
	"github.com/nimaism/trackit/internal/notifier"
	"github.com/nimaism/trackit/internal/store"
	"github.com/nimaism/trackit/pkg/version"

	"github.com/projectdiscovery/goflags"
)

func parseFlags() string {
	var configPath string
	var disableUpdateCheck bool

	flagSet := goflags.NewFlagSet()
	flagSet.SetDescription("TrackIt is your homie for tracking website changes..")
	flagSet.StringVar(&configPath, "config", "config.yaml", "Path to the configuration file")
	flagSet.BoolVarP(&disableUpdateCheck, "disable-update-check", "duc", false, "Disable automatic update check")
	flagSet.Parse()

	if !disableUpdateCheck {
		version.CheckLatestVersion()
	}

	return configPath
}

func main() {
	version.ShowVersion()

	configPath := parseFlags()

	cfg, err := config.Load(configPath)
	if err != nil {
		log.Fatalf("Failed to load configuration from %s: %v", configPath, err)
	}

	storage := store.NewStore(cfg.StorageFile)

	notifiers, err := notifier.NewNotifiers(cfg)
	if err != nil {
		log.Fatalf("Failed to initialize notifiers: %v", err)
	}

	detector := change.New(notifiers, *storage, cfg)

	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, os.Interrupt, syscall.SIGTERM)

	go func() {
		detector.Start()
	}()

	<-stopChan
	log.Println("Shutting down gracefully...")
}
