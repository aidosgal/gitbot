package app

import (
	"log/slog"

	"github.com/aidosgal/git-bot/internal/config"
)

type TelegramBot struct {
	log *slog.Logger
	cfg *config.Config
}

func NewTelegramBot(log *slog.Logger, cfg *config.Config) *TelegramBot {
	return &TelegramBot{
		log: log,
		cfg: cfg,
	}
}

func (t *TelegramBot) Run() error {
	return nil
}
