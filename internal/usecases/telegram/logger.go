package telegram

import (
	"context"
	"fmt"
	"log/slog"
)

const errAttrKey = "error"

type TelegramHandler struct {
	tg Sender
}

func NewLogger() (*slog.Logger, error) {
	tgSender, err := NewSender()
	if err != nil {
		return nil, err
	}

	return slog.New(
		&TelegramHandler{
			tg: tgSender,
		},
	), nil
}

func (h *TelegramHandler) Handle(ctx context.Context, record slog.Record) error {
	label := ""

	switch record.Level {
	case slog.LevelDebug:
		label = "⚙️"
	case slog.LevelInfo:
		label = "✅"
	case slog.LevelError:
		label = "❌"
	case slog.LevelWarn:
		label = "⚠️"
	}

	var errInfo string

	record.Attrs(func(a slog.Attr) bool {
		if a.Key == errAttrKey {
			errInfo = a.Value.String()
			return true
		}

		return false
	})

	msg := fmt.Sprintf("Status: %s\nTime: %s\nMessage: %s\n", label, record.Time, record.Message)
	if errInfo != "" {
		msg += fmt.Sprintf("Error: %s\n", errInfo)
	}

	h.tg.Send(msg)

	return nil
}

func (h *TelegramHandler) Enabled(ctx context.Context, level slog.Level) bool {
	return true
}

func (h *TelegramHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return nil
}

func (h *TelegramHandler) WithGroup(group string) slog.Handler {
	return nil
}

func LogError(log *slog.Logger, msg string, err error) {
	if log == nil {
		return
	}

	log.Error(msg, slog.Any(errAttrKey, err))
}
