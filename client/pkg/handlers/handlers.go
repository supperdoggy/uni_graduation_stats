package handlers

import (
	"fmt"

	"github.com/supperdoggy/diploma_university_statistics_tool/client/pkg/clients/states"
	"github.com/supperdoggy/diploma_university_statistics_tool/client/pkg/localization"
	"github.com/supperdoggy/diploma_university_statistics_tool/client/pkg/service"
	"go.uber.org/zap"
	"gopkg.in/tucnak/telebot.v2"
)

type IHandler interface {
	Start(m *telebot.Message)
	ProcessMsg(m *telebot.Message)
	Schools(m *telebot.Message)
	TopCompaniesBySchool(m *telebot.Message)
	TopHiredDegrees(m *telebot.Message)
	TopSchoolsByCompany(m *telebot.Message)
	SchoolDegrees(m *telebot.Message)
}

type handler struct {
	log          *zap.Logger
	srv          service.IService
	stateManager *states.StateManager
	loc          *localization.Localization
	bot          *telebot.Bot
}

func NewHandler(log *zap.Logger, srv service.IService, bot *telebot.Bot, sm *states.StateManager) IHandler {
	return &handler{
		log:          log,
		srv:          srv,
		bot:          bot,
		stateManager: sm,
		loc:          localization.NewLocalization(),
	}
}

func (h *handler) Start(m *telebot.Message) {

	h.log.Info("Start command received", zap.Any("m", m))

	_, err := h.bot.Reply(m, "Hello, I'm a bot that can help you with university statistics")
	if err != nil {
		h.log.Error("Error while replying", zap.Error(err))
		return
	}
}

func (h *handler) ProcessMsg(m *telebot.Message) {
	h.log.Info("Message received", zap.Any("m", m))

	state, _ := h.stateManager.GetState(m.Sender.ID)

	switch state {
	case states.NONE:
		return
	case states.TopCompaniesBySchool_INPUT:
		h.TopCompaniesBySchool(m)
	case states.TopHiredDegrees_SCHOOL_INPUT:
		h.TopHiredDegrees(m)
	case states.TopHiredDegrees_COMPANY_INPUT:
		h.TopHiredDegrees(m)
	case states.TopSchoolsByCompany_INPUT:
		h.TopSchoolsByCompany(m)
	case states.SchoolDegrees_INPUT:
		h.SchoolDegrees(m)
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
		if i == 40 {
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

func (h *handler) TopCompaniesBySchool(m *telebot.Message) {
	st, _ := h.stateManager.GetState(m.Sender.ID)
	if st != states.TopCompaniesBySchool_INPUT {
		h.stateManager.SetState(m.Sender.ID, states.TopCompaniesBySchool_INPUT, nil)
		_, _ = h.bot.Send(m.Sender, h.loc.Get("TopCompaniesBySchool_INPUT"))
		return
	}
	defer h.stateManager.SetState(m.Sender.ID, states.NONE, nil)

	companyName := m.Text

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
		text += fmt.Sprintf("%v: %s - %v працівників\n", i+1, company.Name, company.Count)
	}

	_, err = h.bot.Reply(m, text)
	if err != nil {
		h.log.Error("Error while replying", zap.Error(err))
		return
	}
}

func (h *handler) TopHiredDegrees(m *telebot.Message) {
	st, values := h.stateManager.GetState(m.Sender.ID)
	if st != states.TopHiredDegrees_SCHOOL_INPUT && st != states.TopHiredDegrees_COMPANY_INPUT {
		h.stateManager.SetState(m.Sender.ID, states.TopHiredDegrees_SCHOOL_INPUT, nil)
		_, _ = h.bot.Send(m.Sender, h.loc.Get("TopHiredDegrees_SCHOOL_INPUT"))
		return
	} else if st == states.TopHiredDegrees_SCHOOL_INPUT {
		h.stateManager.SetState(m.Sender.ID, states.TopHiredDegrees_COMPANY_INPUT, []string{m.Text})
		_, _ = h.bot.Send(m.Sender, h.loc.Get("TopHiredDegrees_COMPANY_INPUT"))
		return
	}

	defer h.stateManager.SetState(m.Sender.ID, states.NONE, nil)

	school := values[0]
	company := m.Text

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
		text += fmt.Sprintf("%v: %s - %v працівників\n", i+1, degree.Name, degree.Count)
	}

	_, err = h.bot.Reply(m, text)
	if err != nil {
		h.log.Error("Error while replying", zap.Error(err))
		return
	}
}

func (h *handler) TopSchoolsByCompany(m *telebot.Message) {
	st, _ := h.stateManager.GetState(m.Sender.ID)
	if st != states.TopSchoolsByCompany_INPUT {
		h.stateManager.SetState(m.Sender.ID, states.TopSchoolsByCompany_INPUT, nil)
		_, _ = h.bot.Send(m.Sender, h.loc.Get("TopSchoolsByCompany_INPUT"))
		return
	}

	defer h.stateManager.SetState(m.Sender.ID, states.NONE, nil)

	company := m.Text

	resp, err := h.srv.TopSchoolsByCompany(company)
	if err != nil {
		h.log.Error("Error while getting schools", zap.Error(err))
		return
	}

	if resp.Error != "" {
		h.log.Error("Error while getting schools", zap.Error(err))
		h.bot.Reply(m, resp.Error)
		return
	}

	if len(resp.Schools) == 0 {
		h.log.Error("No schools found")
		h.bot.Reply(m, "No schools found")
		return
	}

	var text string
	for i, school := range resp.Schools {
		if i > 30 {
			break
		}
		text += fmt.Sprintf("%v: %s - %v навчальні програми\n", i+1, school.Name, school.Count)
	}

	_, err = h.bot.Reply(m, text)
	if err != nil {
		h.log.Error("Error while replying", zap.Error(err))
		return
	}
}

func (h *handler) SchoolDegrees(m *telebot.Message) {
	st, _ := h.stateManager.GetState(m.Sender.ID)
	if st != states.SchoolDegrees_INPUT {
		h.stateManager.SetState(m.Sender.ID, states.SchoolDegrees_INPUT, nil)
		_, _ = h.bot.Send(m.Sender, h.loc.Get("SchoolDegrees_INPUT"))
		return
	}

	defer h.stateManager.SetState(m.Sender.ID, states.NONE, nil)

	school := m.Text

	resp, err := h.srv.SchoolDegrees(school)
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
		text += fmt.Sprintf("%v: %s - %v students\n", i+1, degree.Degree, degree.Count)
	}

	_, err = h.bot.Reply(m, text)
	if err != nil {
		h.log.Error("Error while replying", zap.Error(err))
		return
	}

}
