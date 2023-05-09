package main

import (
	_ "OneNoterBot/internal/response"
	"OneNoterBot/internal/telegram"
	"OneNoterBot/pkg/logging"
	"database/sql"
	"flag"

	_ "github.com/go-sql-driver/mysql"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	logger := logging.GetLogger()
	logger.Info("Open db")
	db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/golang")
	if err != nil {
		logger.Fatal("Can't open database")
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			logger.Fatal("Can't close database", err)
		}
	}(db)

	bot, err := tgbotapi.NewBotAPI(mustToken(logger))
	if err != nil {
		logger.Fatal(err)
	}
	telegramBot := telegram.NewBot(bot)
	logger.Info("Bot starter")
	telegramBot.Start(logger, db)
}

func mustToken(logger *logging.Logger) string {
	token := flag.String(
		"tg-bot-token",
		"",
		"token for accessing telegram bot",
	)
	flag.Parse()
	if *token == "" {
		logger.Fatal("token is not specified")
	}
	return *token
}
