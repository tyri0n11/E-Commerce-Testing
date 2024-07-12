package initialize

import (
	"github.com/tyri0n11/Muffin/global"
	"github.com/tyri0n11/Muffin/pkg/logger"
)

func InitLogger() {
	global.Logger = logger.NewLogger(global.Config.Logger)
}
