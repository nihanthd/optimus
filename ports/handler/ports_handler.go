package handler

import (
	"bitbucket.org/nihanthd/optimus/pi"
	"github.com/labstack/echo"
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

func (p *PortsHandler) GetPortStatus(c echo.Context) error {
	return p.PI.Status(c)
}
