package http

import "github.com/sirupsen/logrus"

type UserController struct {
	Log *logrus.Logger
}

func NewUserController(log *logrus.Logger) *UserController {
	return &UserController{
		Log: log,
	}
}
