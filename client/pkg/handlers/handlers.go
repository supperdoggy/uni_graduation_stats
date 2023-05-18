package handlers

import (
	"fmt"
	"strings"

	"github.com/supperdoggy/diploma_university_statistics_tool/client/pkg/service"
	"go.uber.org/zap"
	"gopkg.in/tucnak/telebot.v2"
)

type IHandler interface {
	Schools(m *telebot.Message)
	TopCompanies(m *telebot.Message)
}

type handler struct {
	log *zap.Logger
	srv service.IService
	bot *telebot.Bot
}

func NewHandler(log *zap.Logger, srv service.IService, bot *telebot.Bot) *handler {
	return &handler{
		log: log,
		srv: srv,
		bot: bot,
	}
}

func (h *handler) Schools(m *telebot.Message) {
	schools, err := h.srv.Schools()
	if err != nil {
		h.log.Error("Error while getting schools", zap.Error(err))
		return
	}

	var text string

	for i, school := range schools.Schools {
		if i == 10 {
			text += school.Name
			break
		}
		text += school.Name + "\n"
	}

	_, err = h.bot.Reply(m, text)
	if err != nil {
		h.log.Error("Error while replying", zap.Error(err))
		return
	}
}

func (h *handler) TopCompanies(m *telebot.Message) {

	msg := strings.Split(m.Text, "/top_companies ")
	if len(msg) < 2 {
		h.log.Error("Error while getting school name")
		return
	}

	companyName := msg[1]

	resp, err := h.srv.TopCompanies(companyName)
	if err != nil {
		h.log.Error("Error while getting companies", zap.Error(err))
		return
	}

	if resp.Error != "" {
		h.log.Error("Error while getting companies", zap.Error(err))
		h.bot.Reply(m, resp.Error)
		return
	}

	if len(resp.Companies) == 0 {
		h.log.Error("No companies found")
		h.bot.Reply(m, "No companies found")
		return
	}

	var text string
	for i, company := range resp.Companies {
		if i > 30 {
			break
		}
		text += fmt.Sprintf("%v: %s - %v employees\n", i+1, company.Name, company.Count)
	}

	_, err = h.bot.Reply(m, text)
	if err != nil {
		h.log.Error("Error while replying", zap.Error(err))
		return
	}

}

func (h *handler) TopHiredDegrees(m *telebot.Message) {
	msg := strings.Split(m.Text, "/top_hired_degrees ")
	if len(msg) < 2 {
		h.log.Error("Error while getting school name")
		return
	}

	school := strings.Split(msg[1], "&")[0]
	company := strings.Split(msg[1], "&")[1]

	resp, err := h.srv.TopHiredDegrees(school, company)
	if err != nil {
		h.log.Error("Error while getting degrees", zap.Error(err))
		return
	}

	if resp.Error != "" {
		h.log.Error("Error while getting degrees", zap.Error(err))
		h.bot.Reply(m, resp.Error)
		return
	}

	if len(resp.Degrees) == 0 {
		h.log.Error("No degrees found")
		h.bot.Reply(m, "No degrees found")
		return
	}

	var text string
	for i, degree := range resp.Degrees {
		if i > 30 {
			break
		}
		text += fmt.Sprintf("%v: %s - %v employees\n", i+1, degree.Name, degree.Count)
	}

	_, err = h.bot.Reply(m, text)
	if err != nil {
		h.log.Error("Error while replying", zap.Error(err))
		return
	}

}
