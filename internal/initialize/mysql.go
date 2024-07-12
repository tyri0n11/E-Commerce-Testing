package initialize

import (
	"fmt"
	"time"

	"github.com/tyri0n11/Muffin/global"
	"github.com/tyri0n11/Muffin/internal/po"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func checkError(err error, s string) {
	if err != nil {
		global.Logger.Error(s, zap.Error(err))
		panic(err)
	}
}

func InitMySQL() {
	m := global.Config.Mysql

	dsn := "%s:%s@tcp(%s:%v)/%s?charset=utf8mb4&parseTime=True&loc=local"
	var s = fmt.Sprintf(dsn, m.User, m.Password, m.Host, m.Port, m.Database)

	db, err := gorm.Open(mysql.Open(s), &gorm.Config{
		SkipDefaultTransaction: false,
	})
	checkError(err, "InitMysql initialization error")
	global.Logger.Info("Initializing MySQL successfully")
	global.Mdb = db
	setPool()
	migrateTables()
}

func setPool() {
	m := global.Config.Mysql
	sqlDb, err := global.Mdb.DB()
	checkError(err, "Set Pool")
	sqlDb.SetConnMaxIdleTime(time.Duration(m.MaxIdleConns))
	sqlDb.SetMaxOpenConns(m.MaxOpenConns)
	sqlDb.SetConnMaxLifetime(time.Duration(m.ConnMaxLifetime))

}
func migrateTables() {
	err := global.Mdb.AutoMigrate(
		&po.User{},
		&po.Role{},
	)
	if err != nil {
		fmt.Print("migrating table error")
	}
}
