package setting

type Config struct {
	Mysql  MySQLSetting  `mapstructure:"mysql"`
	Logger LoggerSetting `mapstructure:"logger"`
	Redis  RedisSetting  `mapstructure:"redis"`
	Server ServerSetting `mapstructure:"server"`
}
type ServerSetting struct {
	Port int    `mapstructure:"port"`
	Mode string `mapstructure:"mode"`
}

type MySQLSetting struct {
	Host            string `mapstructure:"host"`
	Port            string `mapstructure:"port"`
	User            string `mapstructure:"username"`
	Password        string `mapstructure:"password"`
	Database        string `mapstructure:"dbname"`
	MaxIdleConns    int    `mapstructure:"maxIdleConnections"`
	MaxOpenConns    int    `mapstructure:"maxOpenConnections"`
	ConnMaxLifetime int    `mapstructure:"connMaxLifetime"`
}

type LoggerSetting struct {
	Level      string `mapstructure:"log_level"`
	Filename   string `mapstructure:"file_log_name"`
	MaxSize    int    `mapstructure:"maxSize"`
	MaxBackups int    `mapstructure:"maxBackups"`
	MaxAge     int    `mapstructure:"maxAge"`
	Compress   bool   `mapstructure:"compress"`
}

type RedisSetting struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Password string `mapstructure:"password"`
	Database int    `mapstructure:"database"`
}
