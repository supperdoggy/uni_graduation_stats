package main

import (
	"github.com/supperdoggy/diploma_university_statistics_tool/client/pkg/cfg"
	"github.com/supperdoggy/diploma_university_statistics_tool/client/pkg/clients/states"
	"github.com/supperdoggy/diploma_university_statistics_tool/client/pkg/clients/unistats"
	"github.com/supperdoggy/diploma_university_statistics_tool/client/pkg/handlers"
	"github.com/supperdoggy/diploma_university_statistics_tool/client/pkg/service"
	"go.uber.org/zap"
	"gopkg.in/tucnak/telebot.v2"
)

func main() {
	config := cfg.NewConfig()
	log, _ := zap.NewDevelopment()

	stateManager := states.NewStateManager()
	api := unistats.NewUniStats(log, config.UniStatsURL)
	srv := service.NewService(log, api)
	bot, err := telebot.NewBot(telebot.Settings{
		Token: config.Token,
		Poller: &telebot.LongPoller{
			Timeout: 10,
		},
	})
	h := handlers.NewHandler(log, srv, bot, stateManager)

	bot.Handle(telebot.OnText, h.ProcessMsg)
	bot.Handle("/start", h.Start)

	bot.Handle("/schools", h.Schools)
	bot.Handle("/top_companies_by_school", h.TopCompaniesBySchool)
	bot.Handle("/top_hired_degrees", h.TopHiredDegrees)
	bot.Handle("/top_schools_by_company", h.TopSchoolsByCompany)
	bot.Handle("/school_degrees", h.SchoolDegrees)

	if err != nil {
		log.Fatal("Error while creating bot", zap.Error(err))
	}

	log.Info("Bot started")

	bot.Start()

}
