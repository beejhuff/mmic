package trello

import (
	log "github.com/Sirupsen/logrus"
	"github.com/eternnoir/mmic/mmic/config"
	"github.com/eternnoir/mmic/mmic/payloads/trello"
	"github.com/labstack/echo"
	"net/http"
)

type TrelloHandler struct {
	Config    *config.TrelloConfig
	RoutePath string
}

func NewTrelloHandler(config config.TrelloConfig) *TrelloHandler {
	th := &TrelloHandler{Config: &config}
	return th
}

func (handler *TrelloHandler) Handle(c echo.Context) error {
	trelloUpdate := &trello.UpdateAction{}
	if err := c.Bind(trelloUpdate); err != nil {
		return err
	}
	log.Infof("TrelloHandler get new update %#v", trelloUpdate)
	return c.String(http.StatusOK, "")
}
