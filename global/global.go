package global

import (
	"github.com/tyri0n11/Muffin/pkg/logger"
	"github.com/tyri0n11/Muffin/pkg/setting"
	"gorm.io/gorm"
)

var (
	Config setting.Config
	Logger *logger.LoggerZap
	Mdb    *gorm.DB
)
