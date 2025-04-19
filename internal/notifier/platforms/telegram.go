package platforms

import (
	"fmt"
	"net/http"
	"net/url"
)

type TelegramNotifier struct {
	botToken string
	chatID   string
}

func NewTelegramNotifier(botToken, chatID string) *TelegramNotifier {
	return &TelegramNotifier{
		botToken: botToken,
		chatID:   chatID,
	}
}

func (n *TelegramNotifier) Notify(message string) error {
	apiURL := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", n.botToken)
	data := url.Values{}
	data.Set("chat_id", n.chatID)
	data.Set("text", message)

	resp, err := http.PostForm(apiURL, data)
	if err != nil {
		return fmt.Errorf("failed to send Telegram notification: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected response from Telegram API: %d", resp.StatusCode)
	}

	return nil
}
