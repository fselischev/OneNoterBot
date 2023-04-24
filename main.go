package main

import (
	"OneNoterBot/pkg/telegram"
	_ "OneNoterBot/response"
	"database/sql"
	"flag"
	_ "github.com/go-sql-driver/mysql"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

func main() {
	db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/golang")
	if err != nil {
		log.Fatal("Can't open database")
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatal("Can't close database", err)
		}
	}(db)

	bot, err := tgbotapi.NewBotAPI(mustToken())
	if err != nil {
		log.Fatal(err)
	}

	telegramBot := telegram.NewBot(bot)
	telegramBot.Start(db)
}

func mustToken() string {
	token := flag.String(
		"tg-bot-token",
		"",
		"token for accessing telegram bot",
	)
	flag.Parse()
	if *token == "" {
		log.Fatal("token is not specified")
	}
	return *token
}
