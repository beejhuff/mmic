package mmic

import (
	log "github.com/Sirupsen/logrus"

	"fmt"
	"github.com/eternnoir/mmic/mmic/config"
	"github.com/eternnoir/mmic/mmic/handlers"
	"github.com/eternnoir/mmic/mmic/handlers/trello"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/labstack/echo/middleware"
)

type Mmic struct {
	Port       string
	HttpServer *echo.Echo
	RouteMap   map[string]handlers.Handler
}

func NewMmic(config *config.MmicConfig) (*Mmic, error) {
	log.Infof("Create Mnic by mmic config: %#v", config)
	mmic := &Mmic{Port: config.Port, HttpServer: echo.New()}
	mmic.RouteMap = make(map[string]handlers.Handler)
	err := mmic.bindHandlers(config)
	if err != nil {
		return nil, err
	}

	return mmic, nil
}

func (mmic *Mmic) bindHandlers(config *config.MmicConfig) error {
	err := mmic.bindTrelloHandlers(config.TrelloConfigs)
	if err != nil {
		log.Errorf("Bind handler fail. %s", err)
		return err
	}

	return nil
}

func (mmic *Mmic) bindTrelloHandlers(configs []config.TrelloConfig) error {
	for _, config := range configs {
		if _, inmap := mmic.RouteMap[config.RoutePath]; inmap {
			return fmt.Errorf("Rout path %s already exist.", config.RoutePath)
		}
		th := trello.NewTrelloHandler(config)
		log.Debugf("Bind %#v", config)
		mmic.HttpServer.Post(config.RoutePath, th)
		mmic.RouteMap[config.RoutePath] = th
	}
	return nil
}

func (mmic *Mmic) Start() {
	mmic.HttpServer.Use(middleware.Logger())
	mmic.HttpServer.Use(middleware.Recover())
	mmic.HttpServer.Run(standard.New(":" + mmic.Port))
}
