package ru

const (
	GreetingRU = `
Привет, пользователь!
Пожалуйста, авторизуйтесь с помощью пароля, потом можете писать свои заметки 🗒🗒🗒`
	HelpRU = `
/start - Аутентификация, первое, что вам нужно сделать
/notes - Все заметки будут показаны в виде списка
/clear - Удалить 1 заметку по номеру. После команды необходимо дать подтверждение и выбрать номер заметки, которая будет удалена,
	например: да, 1
/clearall - Удалить все заметки
/whoami - Выводит имя пользователя, tg handle и tg link
	мы надеемся, что вы не забыли, кто вы 🙃
/help - Выводит эту информацию

Бот является очень очень MVP
По всем вопросам и найденным багам обращайтесь к @fselischev`
	AuthorizationFailedRU  = "Тебе нужно авторизоваться через команду /start"
	ClearVerificationRU    = `Уверены насчёт этого? Выбранная заметка будет удалена`
	ClearallVerificationRU = `Точно хотите удалить все заметки?`
	CommandNotSupportedRU  = `
Это команда не поддерживается :( 
/help показывает все доступные команды`
	ClearYesRU          = "Заметка успешно удалена"
	ClearNoRU           = "Заметки с таким номером нет, попробуйте ещё раз"
	ClearIncorrectRU    = "Неправильный ответ.. Напишите да или нет и номер заметки, например: да, 1"
	ClearallYesRU       = "Ваши заметки успешно удалены!"
	ClearallNoRU        = "Окей, ваши заметки остались на своём месте"
	ClearallIncorrectRU = "Неправильный ответ.. Напишите да или нет"
	EmptyNotesRU        = "Пока что у вас нет заметок :("
)