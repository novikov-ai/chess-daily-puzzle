package telegram

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Sender interface {
	Send(string) error
}

type sender struct {
	bot *tgbotapi.BotAPI
}

func NewSender() (*sender, error) {
	token := os.Getenv("TELEGRAM_BOT_TOKEN")
	if token == "" {
		return nil, errors.New("telegram bot token is empty")
	}

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}

	return &sender{
		bot: bot,
	}, nil
}

func (s *sender) Send(msg string) error {
	chatID := os.Getenv("TELEGRAM_CHAT_ID")
	if chatID == "" {
		return errors.New("telegram chat ID is empty")
	}

	id, err := strconv.Atoi(chatID)
	if err != nil {
		return err
	}

	_, err = s.bot.Send(composeMessage(int64(id), msg))
	if err != nil {
		return err
	}

	return nil
}

func composeMessage(chatID int64, msg string) tgbotapi.MessageConfig {
	config := tgbotapi.NewMessage(chatID, fmt.Sprintf("`%s`", msg))
	config.ParseMode = "MarkdownV2"

	return config
}
