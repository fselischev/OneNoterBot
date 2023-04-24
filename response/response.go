package response

import (
	en "OneNoterBot/response/en"
	"OneNoterBot/response/ru"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strings"
)

type Response struct {
	lang string
}

func NewResponse(lang string) *Response {
	return &Response{lang: lang}
}

func (r *Response) GiveNotes(notes []interface{}) string {
	vs := make([]string, len(notes))
	for i, v := range notes {
		vs[i] = v.(string)
	}
	var pref string
	switch r.lang {
	case "en":
		pref = "Here's your notes"
	case "ru":
		pref = "Твои заметки"
	}
	return fmt.Sprintf("%s:\n"+strings.Join(vs, "\n"), pref)
}

func (r *Response) AuthorizationSuccess(username string) string {
	switch r.lang {
	case "en":
		return fmt.Sprintf("Authorized success, %s\nNow, enter your notes", username)
	case "ru":
		return fmt.Sprintf("Авторизация прошла успешно, %s\nТеперь вы можете делать заметки", username)
	default:
		return "not supported"
	}
}

func (r *Response) DataSavedSuccess(username string) string {
	var pref string
	switch r.lang {
	case "en":
		pref = "Got your data"
	case "ru":
		pref = "Записал"
	default:
		return "not supported"
	}
	if username != "" {
		return fmt.Sprintf("%s, %s", username, pref)
	}
	return r.AuthorizationFailed()
}

func (r *Response) WhoAmI(username string, upd tgbotapi.Update) string {
	switch r.lang {
	case "en":
		return fmt.Sprintf("You are loged as %s\ntg handle @%s\ntg link https://t.me/%s", username, upd.Message.From.UserName, upd.Message.From.UserName)
	case "ru":
		return fmt.Sprintf("Вы авторизованы как %s\ntg handle @%s\ntg link https://t.me/%s", username, upd.Message.From.UserName, upd.Message.From.UserName)
	default:
		return "not supported"
	}
}

func (r *Response) Greeting() string {
	switch r.lang {
	case "en":
		return en.GreetingEN
	case "ru":
		return ru.GreetingRU
	default:
		return "not supported"
	}
}

func (r *Response) Help() string {
	switch r.lang {
	case "en":
		return en.HelpEN
	case "ru":
		return ru.HelpRU
	default:
		return "not supported"
	}
}

func (r *Response) AuthorizationFailed() string {
	switch r.lang {
	case "en":
		return en.AuthorizationFailedEN
	case "ru":
		return ru.AuthorizationFailedRU
	default:
		return "not supported"
	}
}

func (r *Response) EmptyNotes() string {
	switch r.lang {
	case "en":
		return en.EmptyNotesEN
	case "ru":
		return ru.EmptyNotesRU
	default:
		return "not supported"
	}
}

func (r *Response) ClearVerification() string {
	switch r.lang {
	case "en":
		return en.ClearVerificationEN
	case "ru":
		return ru.ClearVerificationRU
	default:
		return "not supported"
	}
}

func (r *Response) ClearallVerification() string {
	switch r.lang {
	case "en":
		return en.ClearallVerificationEN
	case "ru":
		return ru.ClearallVerificationRU
	default:
		return "not supported"
	}
}

func (r *Response) CommandNotSupported() string {
	switch r.lang {
	case "en":
		return en.CommandNotSupportedEN
	case "ru":
		return ru.CommandNotSupportedRU
	default:
		return "not supported"
	}
}

func (r *Response) ClearYes() string {
	switch r.lang {
	case "en":
		return en.ClearYesEN
	case "ru":
		return ru.ClearYesRU
	default:
		return "not supported"
	}
}

func (r *Response) ClearallNo() string {
	switch r.lang {
	case "en":
		return en.ClearallNoEN
	case "ru":
		return ru.ClearallNoRU
	default:
		return "not supported"
	}
}
func (r *Response) ClearallYes() string {
	switch r.lang {
	case "en":
		return en.ClearallYesEN
	case "ru":
		return ru.ClearallYesRU
	default:
		return "not supported"
	}
}

func (r *Response) ClearallIncorrect() string {
	switch r.lang {
	case "en":
		return en.ClearallIncorrectEN
	case "ru":
		return ru.ClearallIncorrectRU
	default:
		return "not supported"
	}
}
