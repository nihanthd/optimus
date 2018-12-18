package pi

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/labstack/echo"

	"github.com/stianeikeland/go-rpio"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type Config struct {
	Enabled bool `yaml:"enabled"`
}

type PI struct {
	config *Config
}

type Pin struct {
	PinNumber int `json:"pin",omitempty`
	Status    int `json:"status"`
}

func NewPI(lc fx.Lifecycle, log *zap.Logger, config *Config) (*PI, error) {
	log.Info("Initializing PI")

	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			if config.Enabled {
				err := rpio.Open()
				return err
			}
			return nil
		},
		OnStop: func(context.Context) error {
			if config.Enabled {
				err := rpio.Close()
				return err
			}
			return nil
		},
	})
	return &PI{config: config}, nil
}

func (pi *PI) Turn(c echo.Context) error {
	pinNumber, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusBadRequest, "Wrong Port Number")
	}
	p := Pin{}
	var data []byte
	_, err = c.Request().Body.Read(data)
	if err != nil {
		return c.String(http.StatusBadRequest, "Unable to read data")
	}
	err = json.Unmarshal(data, p)
	if err != nil {
		return c.String(http.StatusBadRequest, "Unable to read json data")
	}
	if pi.config.Enabled {
		pin := rpio.Pin(pinNumber)
		if p.Status == 0 {
			pin.Low()
		} else {
			pin.High()
		}
		response, _ := json.Marshal(p)
		return c.String(http.StatusOK, string(response))
	}
	return c.String(http.StatusOK, "Not running on raspberry pi")
}

func (pi *PI) Status(c echo.Context) error {
	pinNumber, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusBadRequest, "Wrong Port Number")
	}
	if pi.config.Enabled {
		pin := rpio.Pin(pinNumber)
		status := rpio.ReadPin(pin)
		response := Pin{
			PinNumber: pinNumber,
			Status:    int(status),
		}
		responseBody, _ := json.Marshal(response)
		return c.String(http.StatusOK, string(responseBody))
	}
	return c.String(http.StatusOK, "Not running on raspberry pi")
}
