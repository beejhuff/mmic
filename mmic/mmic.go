package mmic

import (
	log "github.com/Sirupsen/logrus"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/labstack/echo/middleware"
)

type Mmic struct {
	Port string
}

func NewMmic() *Mmic {
	return nil
}

func (mmic *Mmic) Start() error {
	return nil
}
