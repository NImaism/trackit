package notifier

import (
	"fmt"
	"log"
	"time"

	"github.com/nimaism/trackit/internal/config"
	"github.com/nimaism/trackit/internal/notifier/platforms"
)

type Notifier interface {
	Notify(message string) error
}

func NewNotifiers(cfg *config.Config) ([]Notifier, error) {
	var notifiers []Notifier

	if cfg.Notifier.Discord.Enabled {
		if cfg.Notifier.Discord.WebhookURL == "" {
			return nil, fmt.Errorf("Discord is enabled but WebhookURL is missing")
		}
		discordNotifier := platforms.NewDiscordNotifier(cfg.Notifier.Discord.WebhookURL)
		notifiers = append(notifiers, discordNotifier)
	}

	if cfg.Notifier.Telegram.Enabled {
		if cfg.Notifier.Telegram.BotToken == "" || cfg.Notifier.Telegram.ChatID == "" {
			return nil, fmt.Errorf("Telegram is enabled but BotToken or ChatID is missing")
		}
		telegramNotifier := platforms.NewTelegramNotifier(cfg.Notifier.Telegram.BotToken, cfg.Notifier.Telegram.ChatID)
		notifiers = append(notifiers, telegramNotifier)
	}

	if len(notifiers) == 0 {
		log.Println("No notifiers are enabled in the configuration. Notifications will not be sent.")
	}

	return notifiers, nil
}

func GenerateShortAlertMessage(url string) string {
	timestamp := time.Now().UTC().Format("2006-01-02 15:04:05 UTC")
	return fmt.Sprintf(
		"🚨 Alert: Change detected at %s\n📅 %s",
		url, timestamp,
	)
}

func Alert(notifiers []Notifier, url string) {
	message := GenerateShortAlertMessage(url)
	for _, notifier := range notifiers {
		err := notifier.Notify(message)
		if err != nil {
			log.Printf("Failed to send notification: %v", err)
		}
	}
}
