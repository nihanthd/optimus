package service

import (
	"go.uber.org/zap"
)

type Service struct {
	Logger *zap.Logger `inject:""`
}

func New() (*Service, error) {

	return nil, nil
}
