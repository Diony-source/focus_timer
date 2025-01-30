package models

import (
	"github.com/sirupsen/logrus"
)

// Timer struct (Temiz kod için models klasörüne taşındı)
type Timer struct {
	TotalSeconds int
	Pause        bool
	Stop         bool
	Log          *logrus.Logger
}
