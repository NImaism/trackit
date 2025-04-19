package config

type Config struct {
	URLsFile    string `yaml:"urls_file"`
	Interval    uint16 `yaml:"interval"`
	StorageFile string `yaml:"storage_file"`
	Concurrency uint8  `yaml:"concurrency"`
	Notifier    struct {
		Discord struct {
			Enabled    bool   `yaml:"enabled"`
			WebhookURL string `yaml:"webhook_url"`
		} `yaml:"discord"`
		Telegram struct {
			Enabled  bool   `yaml:"enabled"`
			BotToken string `yaml:"bot_token"`
			ChatID   string `yaml:"chat_id"`
		} `yaml:"telegram"`
	} `yaml:"notifier"`
	Network struct {
		TimeoutSec      int  `yaml:"timeout_sec"`
		VerifySSL       bool `yaml:"verify_ssl"`
		DisableRedirect bool `yaml:"disable_redirect"`
	} `yaml:"network"`
}
