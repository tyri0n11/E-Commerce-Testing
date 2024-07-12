package initialize

import (
	"fmt"

	"github.com/tyri0n11/Muffin/global"
)

func Run() {

	LoadConfig()
	m := global.Config.Mysql
	fmt.Println("Loading configuration mysql...", m.User, m.Host, m.Port, m.Database)
	InitLogger()
	InitMySQL()
	InitRedis()
	r := InitRouter()

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}
