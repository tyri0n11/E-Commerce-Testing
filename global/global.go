package global

import (
	"github.com/redis/go-redis/v9"
	"github.com/tyri0n11/Muffin/pkg/logger"
	"github.com/tyri0n11/Muffin/pkg/setting"
	"gorm.io/gorm"
)

var (
	Config setting.Config
	Logger *logger.LoggerZap
	Mdb    *gorm.DB
	Rdb    *redis.Client
)
