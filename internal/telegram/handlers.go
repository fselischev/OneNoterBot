package telegram

import (
	"OneNoterBot/internal/response"

	"github.com/emirpasic/gods/maps/linkedhashmap"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func ClearHandler(password string, msg *tgbotapi.MessageConfig, resp *response.Response, upd tgbotapi.Update, numberedNotes *linkedhashmap.Map) (fl bool) {
	if password == "" {
		msg.Text = resp.AuthorizationFailed()
		msg.ReplyToMessageID = upd.Message.MessageID
	} else {
		if numberedNotes.Size() == 0 {
			msg.Text = resp.EmptyNotes()
		} else {
			fl = true
		}
	}
	return
}

func WhoamiHandler(password string, msg *tgbotapi.MessageConfig, resp *response.Response, upd tgbotapi.Update) {
	if password == "" {
		msg.Text = resp.AuthorizationFailed()
		msg.ReplyToMessageID = upd.Message.MessageID
	} else {
		msg.Text = resp.WhoAmI(password, upd)
		msg.ReplyToMessageID = upd.Message.MessageID
	}
}
