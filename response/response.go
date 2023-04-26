package response

import (
	"OneNoterBot/response/en"
	"OneNoterBot/response/ru"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/samber/lo"
	"strings"
)

const notSupported = "not supported"

type Response struct {
	lang string
}

func NewResponse(lang string) *Response {
	return &Response{lang: lang}
}

func (r *Response) GiveNotes(notes []interface{}) string {
	var pref string
	switch r.lang {
	case "en":
		pref = "Here's your notes"
	case "ru":
		pref = "Ваши заметки"
	}
	return fmt.Sprintf("%s:\n"+strings.Join(lo.Map(notes, func(v interface{}, i int) string { return v.(string) }), "\n"), pref)
}

func (r *Response) AuthorizationSuccess(firstname string) string {
	switch r.lang {
	case "en":
		return fmt.Sprintf("Authorized success, %s\nNow, enter your notes", firstname)
	case "ru":
		return fmt.Sprintf("Авторизация прошла успешно, %s\nТеперь вы можете делать заметки", firstname)
	default:
		return notSupported
	}
}

func (r *Response) DataSavedSuccess(firstname string) string {
	var pref string
	switch r.lang {
	case "en":
		pref = "Got your data"
	case "ru":
		pref = "Записал"
	default:
		return notSupported
	}
	if firstname != "" {
		return fmt.Sprintf("%s, %s", pref, firstname)
	}
	return r.AuthorizationFailed()
}

func (r *Response) WhoAmI(username string, upd tgbotapi.Update) string {
	switch r.lang {
	case "en":
		return fmt.Sprintf("You are loged with password %s\ntg handle @%s", username, upd.Message.From.UserName)
	case "ru":
		return fmt.Sprintf("Вы авторизованы с паролем %s\ntg handle @%s", username, upd.Message.From.UserName)
	default:
		return notSupported
	}
}

func (r *Response) Greeting() string {
	switch r.lang {
	case "en":
		return en.GreetingEN
	case "ru":
		return ru.GreetingRU
	default:
		return notSupported
	}
}

func (r *Response) Help() string {
	switch r.lang {
	case "en":
		return en.HelpEN
	case "ru":
		return ru.HelpRU
	default:
		return notSupported
	}
}

func (r *Response) AuthorizationFailed() string {
	switch r.lang {
	case "en":
		return en.AuthorizationFailedEN
	case "ru":
		return ru.AuthorizationFailedRU
	default:
		return notSupported
	}
}

func (r *Response) EmptyNotes() string {
	switch r.lang {
	case "en":
		return en.EmptyNotesEN
	case "ru":
		return ru.EmptyNotesRU
	default:
		return notSupported
	}
}

func (r *Response) ClearVerification() string {
	switch r.lang {
	case "en":
		return en.ClearVerificationEN
	case "ru":
		return ru.ClearVerificationRU
	default:
		return notSupported
	}
}

func (r *Response) ClearallVerification() string {
	switch r.lang {
	case "en":
		return en.ClearallVerificationEN
	case "ru":
		return ru.ClearallVerificationRU
	default:
		return notSupported
	}
}

func (r *Response) CommandNotSupported() string {
	switch r.lang {
	case "en":
		return en.CommandNotSupportedEN
	case "ru":
		return ru.CommandNotSupportedRU
	default:
		return notSupported
	}
}

func (r *Response) ClearYes() string {
	switch r.lang {
	case "en":
		return en.ClearYesEN
	case "ru":
		return ru.ClearYesRU
	default:
		return notSupported
	}
}

func (r *Response) ClearNo() string {
	switch r.lang {
	case "en":
		return en.ClearNoEN
	case "ru":
		return ru.ClearNoRU
	default:
		return notSupported
	}
}

func (r *Response) ClearallNo() string {
	switch r.lang {
	case "en":
		return en.ClearallNoEN
	case "ru":
		return ru.ClearallNoRU
	default:
		return notSupported
	}
}
func (r *Response) ClearallYes() string {
	switch r.lang {
	case "en":
		return en.ClearallYesEN
	case "ru":
		return ru.ClearallYesRU
	default:
		return notSupported
	}
}

func (r *Response) ClearallIncorrect() string {
	switch r.lang {
	case "en":
		return en.ClearallIncorrectEN
	case "ru":
		return ru.ClearallIncorrectRU
	default:
		return notSupported
	}
}

func (r *Response) ClearIncorrect() string {
	switch r.lang {
	case "en":
		return en.ClearIncorrectEN
	case "ru":
		return ru.ClearIncorrectRU
	default:
		return notSupported
	}
}

func (r *Response) IsPositive(answer string) bool {
	return answer == "yes" || answer == "да"
}

func (r *Response) IsNegative(answer string) bool {
	return answer == "no" || answer == "нет"
}
