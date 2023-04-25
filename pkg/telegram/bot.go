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
		msg           tgbotapi.MessageConfig
		password      string
		firstname     string
		isStart       = false
		isClear       = false
		isClearall    = false
		numberedNotes = linkedhashmap.New()
		lang          string
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

		firstname = upd.Message.From.FirstName

		resp := response.NewResponse(lang)

		if upd.Message.IsCommand() {
			msg = tgbotapi.NewMessage(upd.Message.Chat.ID, "")
			switch upd.Message.Command() {
			case "start":
				isStart = true
				msg.Text = resp.Greeting()
				if _, err := b.bot.Send(msg); err != nil {
					log.Fatal(err)
				}
				continue
			case "help":
				msg.Text = resp.Help()
			case "notes":
				if password == "" {
					msg.Text = resp.AuthorizationFailed()
					msg.ReplyToMessageID = upd.Message.MessageID
					b.sendMessage(msg)
					continue
				}
				numberedNotes.Clear()
				smlp, err := db.Query(fmt.Sprintf("SELECT note FROM users WHERE password = '%s'", password))
				if err != nil {
					log.Panic(err)
				}
				i := 1
				for smlp.Next() {
					var note string
					if err := smlp.Scan(&note); err != nil {
						log.Panic(err)
					}
					numberedNotes.Put(i, fmt.Sprintf("%d. %s", i, note))
					i++
				}
				if numberedNotes.Size() == 0 {
					msg.Text = resp.EmptyNotes()
				} else {
					msg.Text = resp.GiveNotes(numberedNotes.Values())
				}
			case "clear":
				if numberedNotes.Size() == 0 {
					msg.Text = resp.EmptyNotes()
				} else {
					if password == "" {
						msg.Text = resp.AuthorizationFailed()
						msg.ReplyToMessageID = upd.Message.MessageID
						b.sendMessage(msg)
						continue
					}
					isClear = true
					msg.Text = resp.ClearVerification()
				}
			case "clearall":
				if numberedNotes.Size() == 0 {
					msg.Text = resp.EmptyNotes()
				} else {
					if password == "" {
						msg.Text = resp.AuthorizationFailed()
						msg.ReplyToMessageID = upd.Message.MessageID
						b.sendMessage(msg)
						continue
					}
					isClearall = true
					msg.Text = resp.ClearallVerification()
				}
			case "whoami":
				if password == "" {
					msg.Text = resp.AuthorizationFailed()
					msg.ReplyToMessageID = upd.Message.MessageID
					b.sendMessage(msg)
					continue
				}
				msg.Text = resp.WhoAmI(password, upd)
				msg.ReplyToMessageID = upd.Message.MessageID
			default:
				msg.Text = resp.CommandNotSupported()
				msg.ReplyToMessageID = upd.Message.MessageID
			}
		} else {
			switch {
			case isStart:
				password = upd.Message.Text
				log.Println(resp.WhoAmI(password, upd))
				isStart = false
				msg.Text = resp.AuthorizationSuccess(firstname)
			case isClear:
				splitMsg := strings.Split(upd.Message.Text, ",")
				ans := strings.ToLower(strings.TrimSpace(splitMsg[0]))
				switch ans {
				case "yes", "да":
					if len(splitMsg) == 1 {
						msg.Text = resp.ClearIncorrect()
						b.sendMessage(msg)
						continue
					}
					key := strings.TrimSpace(splitMsg[1])
					isClear = false
					v, _ := strconv.Atoi(key)
					if v > numberedNotes.Size() {
						msg.Text = resp.ClearNo()
						b.sendMessage(msg)
						continue
					}
					gt, _ := numberedNotes.Get(v)
					_, err := db.Exec(fmt.Sprintf("DELETE FROM `users` WHERE `password` = '%s' AND `note` = '%s'", password, strings.TrimPrefix(gt.(string), fmt.Sprintf("%d. ", v))))
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
					_, err := db.Exec(fmt.Sprintf("DELETE FROM `users` WHERE `password` = '%s'", password))
					numberedNotes.Clear()
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
					_, err := db.Exec(fmt.Sprintf("INSERT INTO users (password, note) VALUES('%s', '%s')", password, strings.TrimSpace(note)))
					if err != nil {
						log.Panic(err)
					}
				}
				msg.Text = resp.DataSavedSuccess(firstname)
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
