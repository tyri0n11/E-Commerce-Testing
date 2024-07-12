package setting

type Config struct {
	Mysql  MySQLSetting  `mapstructure:"mysql"`
	Logger LoggerSetting `mapstructure:"logger"`
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