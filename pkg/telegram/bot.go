package telegram

import (
	"OneNoterBot/response"
	"database/sql"
	"fmt"
	"github.com/emirpasic/gods/maps/linkedhashmap"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"strconv"
	"strings"
)

type Bot struct {
	bot *tgbotapi.BotAPI
}

func NewBot(bot *tgbotapi.BotAPI) *Bot {
	return &Bot{bot: bot}
}

//nolint:funlen
//nolint:gocognit
func (b *Bot) Start(db *sql.DB) {
	updConfig := tgbotapi.NewUpdate(0)

	updConfig.Timeout = 30

	upds := b.bot.GetUpdatesChan(updConfig)

	var (
		msg        tgbotapi.MessageConfig
		username   string
		isStart    = false
		isClear    = false
		isClearall = false
		notes      = linkedhashmap.New()
		lang       string
	)

	for upd := range upds {
		if upd.Message == nil {
			continue
		}

		lang = upd.Message.From.LanguageCode
		switch lang {
		case "en", "ru":
		default:
			lang = "en"
		}

		resp := response.NewResponse(lang)

		if upd.Message.IsCommand() {
			msg = tgbotapi.NewMessage(upd.Message.Chat.ID, "")
			switch upd.Message.Command() {
			case "start":
				isStart = true
				msg.Text = resp.Greeting()
				if _, err := b.bot.Send(msg); err != nil {
					log.Panic(err)
				}
				continue
			case "help":
				msg.Text = resp.Help()
			case "notes":
				if username == "" {
					msg.Text = resp.AuthorizationFailed()
					msg.ReplyToMessageID = upd.Message.MessageID
					b.sendMessage(msg)
					continue
				}
				notes.Clear()
				smlp, err := db.Query(fmt.Sprintf("SELECT note FROM users WHERE name = '%s'", username))
				if err != nil {
					log.Panic(err)
				}
				defer func(smlp *sql.Rows) {
					err := smlp.Close()
					if err != nil {
						log.Panic(err)
					}
				}(smlp)
				i := 1
				for smlp.Next() {
					var note string
					if err := smlp.Scan(&note); err != nil {
						log.Panic(err)
					}
					notes.Put(i, fmt.Sprintf("%d. %s", i, note))
					i++
				}
				if notes.Size() == 0 {
					msg.Text = resp.EmptyNotes()
				} else {
					msg.Text = resp.GiveNotes(notes.Values())
				}
			case "clear":
				if notes.Size() == 0 {
					msg.Text = resp.EmptyNotes()
				} else {
					if username == "" {
						msg.Text = resp.AuthorizationFailed()
						msg.ReplyToMessageID = upd.Message.MessageID
						b.sendMessage(msg)
						continue
					}
					isClear = true
					msg.Text = resp.ClearVerification()
				}
			case "clearall":
				if notes.Size() == 0 {
					msg.Text = resp.EmptyNotes()
				} else {
					if username == "" {
						msg.Text = resp.AuthorizationFailed()
						msg.ReplyToMessageID = upd.Message.MessageID
						b.sendMessage(msg)
						continue
					}
					isClearall = true
					msg.Text = resp.ClearallVerification()
				}
			case "whoami":
				if username == "" {
					msg.Text = resp.AuthorizationFailed()
					msg.ReplyToMessageID = upd.Message.MessageID
					b.sendMessage(msg)
					continue
				}
				msg.Text = resp.WhoAmI(username, upd)
				msg.ReplyToMessageID = upd.Message.MessageID
			default:
				msg.Text = resp.CommandNotSupported()
				msg.ReplyToMessageID = upd.Message.MessageID
			}
		} else {
			switch {
			case isStart:
				username = upd.Message.Text
				log.Println(resp.WhoAmI(username, upd))
				isStart = false
				msg.Text = resp.AuthorizationSuccess(username)
			case isClear:
				splited := strings.Split(upd.Message.Text, ",")
				ans := strings.ToLower(strings.TrimSpace(splited[0]))
				switch ans {
				case "yes", "да":
					key := strings.TrimSpace(splited[1])
					isClear = false
					v, _ := strconv.Atoi(key)
					gt, _ := notes.Get(v)
					_, err := db.Exec(fmt.Sprintf("DELETE FROM `users` WHERE `name` = '%s' AND `note` = '%s'", username, strings.TrimPrefix(gt.(string), fmt.Sprintf("%d. ", v))))
					if err != nil {
						log.Panic(err)
					}
					msg.Text = resp.ClearYes()
				case "no", "нет":
					isClear = false
					msg.Text = resp.ClearallNo()
				default:
					msg.Text = resp.ClearallIncorrect()
				}
			case isClearall:
				verificationMsg := strings.TrimSpace(strings.ToLower(upd.Message.Text))
				switch {
				case resp.IsPositive(verificationMsg):
					isClearall = false
					_, err := db.Exec(fmt.Sprintf("DELETE FROM `users` WHERE `name` = '%s'", username))
					notes.Clear()
					if err != nil {
						log.Panic(err)
					}
					msg.Text = resp.ClearallYes()
				case resp.IsNegative(verificationMsg):
					isClearall = false
					msg.Text = resp.ClearallNo()
				default:
					msg.Text = resp.ClearallIncorrect()
				}
			default:
				msgNotes := strings.Split(upd.Message.Text, ",")
				for _, note := range msgNotes {
					_, err := db.Exec(fmt.Sprintf("INSERT INTO users (name, note) VALUES('%s', '%s')", username, strings.TrimSpace(note)))
					if err != nil {
						log.Panic(err)
					}
				}
				msg.Text = resp.DataSavedSuccess(username)
				msg.ReplyToMessageID = upd.Message.MessageID
			}
		}
		b.sendMessage(msg)
	}
}

func (b *Bot) sendMessage(msg tgbotapi.MessageConfig) {
	if _, err := b.bot.Send(msg); err != nil {
		log.Panic(err)
	}
}
