package en

const (
	GreetingEN = `
	Hello, User!
	Please, authorize first with password, then you can write your notes 🗒🗒🗒`
	HelpEN = `
	/start - Authentication, first thing that you need to do
	/notes - All of your notes will be printed
	/clear - Deletes 1 note by number. After command, you need to verify your choice and pick number of note that will be deleted,
		for example: yes, 1
	/clearall - Deletes all of your notes
	/whoami - Prints your login name, tg handle and tg link
		we hope you have not forgotten who you are 🙃
	/help - Prints this information
	
	Bot is a very very MVP
	For any questions, please contact @fselischev`
	AuthorizationFailedEN  = `You need to authorize first using /start command`
	ClearVerificationEN    = `Are you sure about this? Picked note will be deleted`
	ClearallVerificationEN = `Are you sure about this? All your notes will be deleted`
	CommandNotSupportedEN  = `
This command in not supported :(
/help shows all available commands`
	ClearYesEN          = "Note is successfully deleted"
	ClearNoEN           = "There is no note with this number, try again with correct one"
	ClearIncorrectEN    = "Incorrect answer.. Write yes or no and the number, for example: yes, 1"
	ClearallYesEN       = "Your notes are successfully deleted!"
	ClearallNoEN        = "Well, your notes stayed in place"
	ClearallIncorrectEN = "Incorrect answer.. Write yes or no"
	EmptyNotesEN        = "You don't have notes yet :("
)
