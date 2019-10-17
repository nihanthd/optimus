package handler

import (
	"github.com/labstack/echo"
	"github.com/nihanthd/optimus/pi"
)

type PortsHandler struct {
	PI *pi.PI
}

func NewPortsHandler(pi *pi.PI) *PortsHandler {
	p := &PortsHandler{
		PI: pi,
	}
	return p
}

func (p *PortsHandler) TogglePort(c echo.Context) error {
	return p.PI.Turn(c)
}

func (p *PortsHandler) GetPortStatus(c echo.Context) error {
	return p.PI.Status(c)
}
