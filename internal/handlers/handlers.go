package handlers

import (
	tgbotapi "github.com/go-telegram/bot"
	"skat_bot/internal/fsm"
	"skat_bot/internal/keyboard"
	"skat_bot/internal/service"
	log "skat_bot/pkg/logger"
)

type BotHandler struct {
	service service.Service
	log     *log.Logger
	fsm     fsm.Fsm
}

func New(bot *tgbotapi.Bot, s service.Service, log *log.Logger) {
	f := fsm.New()

	botHandler := &BotHandler{s, log, f}

	bot.RegisterHandler(tgbotapi.HandlerTypeCallbackQueryData,
		keyboard.Back,
		tgbotapi.MatchTypeContains,
		botHandler.Back())
	//get my skats
	bot.RegisterHandler(tgbotapi.HandlerTypeCallbackQueryData,
		keyboard.GetMySkatsCommand,
		tgbotapi.MatchTypeContains,
		botHandler.MySkats())
	//add skat beta
	bot.RegisterHandler(tgbotapi.HandlerTypeCallbackQueryData,
		keyboard.AddSkatCommand,
		tgbotapi.MatchTypeContains,
		botHandler.AddSkatBeta())
	bot.RegisterHandler(tgbotapi.HandlerTypeCallbackQueryData,
		keyboard.AddSkatInstituteCommand,
		tgbotapi.MatchTypeContains,
		botHandler.AddSkatBetaInstitute())
	bot.RegisterHandler(tgbotapi.HandlerTypeCallbackQueryData,
		keyboard.AddSkatSubjectNameCommand,
		tgbotapi.MatchTypeContains,
		botHandler.AddSkatBetaSubjectName())
	bot.RegisterHandler(tgbotapi.HandlerTypeCallbackQueryData,
		keyboard.AddSkatSemesterCommand,
		tgbotapi.MatchTypeContains,
		botHandler.AddSkatBetaSemester())
	bot.RegisterHandler(tgbotapi.HandlerTypeCallbackQueryData,
		keyboard.AddSkatSubjectTypeCommand,
		tgbotapi.MatchTypeContains,
		botHandler.AddSkatBetaSubjectType())
	bot.RegisterHandler(tgbotapi.HandlerTypeCallbackQueryData,
		keyboard.AddSkatVariantTypeCommand,
		tgbotapi.MatchTypeContains,
		botHandler.AddSkatBetaVariantType())

	//get skat beta
	bot.RegisterHandler(tgbotapi.HandlerTypeCallbackQueryData,
		keyboard.GetSkatCommand,
		tgbotapi.MatchTypeContains,
		botHandler.GetSkatBeta())
	bot.RegisterHandler(tgbotapi.HandlerTypeCallbackQueryData,
		keyboard.GetSkatInstituteCommand,
		tgbotapi.MatchTypeContains,
		botHandler.GetSkatBetaInstitute())
	bot.RegisterHandler(tgbotapi.HandlerTypeCallbackQueryData,
		keyboard.GetSkatSubjectNameCommand,
		tgbotapi.MatchTypeContains,
		botHandler.GetSkatBetaSubjectName())
	bot.RegisterHandler(tgbotapi.HandlerTypeCallbackQueryData,
		keyboard.GetSkatSemesterCommand,
		tgbotapi.MatchTypeContains,
		botHandler.GetSkatBetaSemester())
	bot.RegisterHandler(tgbotapi.HandlerTypeCallbackQueryData,
		keyboard.GetSkatSubjectTypeCommand,
		tgbotapi.MatchTypeContains,
		botHandler.GetSkatBetaSubjectType())

	//download skat

	bot.RegisterHandler(tgbotapi.HandlerTypeCallbackQueryData,
		"variant", tgbotapi.MatchTypePrefix, botHandler.DownloadFile())
	//paginators
	bot.RegisterHandler(tgbotapi.HandlerTypeCallbackQueryData,
		keyboard.PageInstitutePaginatorData, tgbotapi.MatchTypeContains, botHandler.InstitutePaginator())
	bot.RegisterHandler(tgbotapi.HandlerTypeCallbackQueryData,
		keyboard.PageSemesterPaginatorData, tgbotapi.MatchTypeContains, botHandler.SemesterPaginator())
	bot.RegisterHandler(tgbotapi.HandlerTypeCallbackQueryData,
		keyboard.PageSubjectNamePaginatorData, tgbotapi.MatchTypeContains, botHandler.SubjectNamePaginator())
	bot.RegisterHandler(tgbotapi.HandlerTypeCallbackQueryData,
		keyboard.PageSubjectTypePaginatorData, tgbotapi.MatchTypeContains, botHandler.SubjecTypePaginator())
	//pass callback
	bot.RegisterHandler(tgbotapi.HandlerTypeCallbackQueryData,
		"pass", tgbotapi.MatchTypePrefix, botHandler.Pass())
	//start
	//bot.RegisterHandler(tgbotapi.HandlerTypeMessageText,
	//	"/start", tgbotapi.MatchTypeExact, Start())
	bot.RegisterHandler(tgbotapi.HandlerTypeMessageText,
		"/menu", tgbotapi.MatchTypeExact, botHandler.Menu())
	////back
	//bot.RegisterHandler(tgbotapi.HandlerTypeMessageText,
	//	keyboard.BackCommand, tgbotapi.MatchTypeExact, back.undo())

}
