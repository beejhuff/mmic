package trello

import (
	"bytes"
	log "github.com/Sirupsen/logrus"
	"github.com/eternnoir/mmic/mmic/config"

	"github.com/eternnoir/mmic/mmic/senders"

	"fmt"
	"github.com/eternnoir/mmic/mmic/payloads/mattermost"
	"github.com/eternnoir/mmic/mmic/payloads/trello"
	"github.com/labstack/echo"
	"net/http"
	"strings"
	"text/template"
)

type TrelloHandler struct {
	Config    *config.TrelloConfig
	RoutePath string
	Sender    senders.Sender
}

func NewTrelloHandler(config config.TrelloConfig) *TrelloHandler {
	th := &TrelloHandler{Config: &config}
	th.Sender = senders.NewMatterMostSender(config.TargetMM)
	return th
}

func (th *TrelloHandler) Handle(c echo.Context) error {
	trelloUpdate := &trello.UpdateAction{}
	if err := c.Bind(trelloUpdate); err != nil {
		return err
	}
	log.Infof("TrelloHandler get new update %#v", trelloUpdate)
	if trelloUpdate.GetType() == "" {
		log.Infof("Not Support type. %#v", trelloUpdate)
		return c.String(http.StatusOK, "")
	}
	payload, err := th.convertToMMPayload(trelloUpdate)
	if err != nil {
		log.Errorf("Convert payload fail.%s", err)
		return err
	}
	serr := th.Sender.Send(payload)
	if serr != nil {
		log.Errorf("Send Paylod fail.%s", err)
		return err
	}
	return c.String(http.StatusOK, "")
}

func (th *TrelloHandler) convertToMMPayload(trelloupdate *trello.UpdateAction) (*mattermost.MatterMostPayload, error) {
	tmpl, err := template.New("mmsendertemplate").Parse(th.Config.TextTemplate)
	if err != nil {
		return nil, fmt.Errorf("Parse template error. %s", err)
	}
	var doc bytes.Buffer
	err = tmpl.Execute(&doc, *trelloupdate)
	if err != nil {
		return nil, fmt.Errorf("Template Exec error. %s", err)
	}
	payload := &mattermost.MatterMostPayload{}
	mmconfig := th.Config.TargetMM

	if mmconfig.Channel != "" {
		payload.Channel = &mmconfig.Channel
	}
	if mmconfig.Username != "" {
		payload.Username = &mmconfig.Username
	}
	textStr := doc.String()
	if strings.TrimSpace(textStr) == "" {
		return nil, fmt.Errorf("Nothing to send. %#v", trelloupdate)
	}
	payload.Text = &textStr

	return payload, nil
}
