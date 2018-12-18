package handler

import (
	"github.com/labstack/echo"
	"github.com/nihanthd/optimus/pi"
)

type PinsHandler struct {
	PI *pi.PI
}

func NewPinsHandler(pi *pi.PI) *PinsHandler {
	p := &PinsHandler{
		PI: pi,
	}
	return p
}

func (p *PinsHandler) TogglePin(c echo.Context) error {
	return p.PI.Turn(c)
}

func (p *PinsHandler) GetPinStatus(c echo.Context) error {
	return p.PI.Status(c)
}
