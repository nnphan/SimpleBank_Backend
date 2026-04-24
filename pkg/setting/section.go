package setting

type Config struct {
	Server     ServerSettings     `mapstructure:"server"`
	Postgresql PostgreSQLSettings `mapstructure:"postgresql"`
	Logger     LoggerSettings     `mapstructure:"logger"`
	Redis      RedisSettings      `mapstructure:"redis"`
}

type ServerSettings struct {
	Port int    `mapstructure:"port"`
	Mode string `mapstructure:"mode"`
}

type RedisSettings struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
}

type PostgreSQLSettings struct {
	Driver          string `mapstructure:"driver"`
	Host            string `mapstructure:"host"`
	Port            int    `mapstructure:"port"`
	DBName          string `mapstructure:"dbname"`
	User            string `mapstructure:"user"`
	Password        string `mapstructure:"password"`
	SSLMode         string `mapstructure:"sslmode"`
	MaxIdleConns    int    `mapstructure:"max_idle_conns"`
	MaxOpenConns    int    `mapstructure:"max_open_conns"`
	ConnMaxLifetime int    `mapstructure:"conn_max_lifetime"`
}

// Setting for logger load from yaml config file
type LoggerSettings struct {
	LogLevel    string `mapstructure:"log_level"`
	FileLogName string `mapstructure:"file_log_name"`
	MaxBackups  int    `mapstructure:"max_backups"`
	MaxSize     int    `mapstructure:"max_size"`
	MaxAge      int    `mapstructure:"max_age"`
	Compress    bool   `mapstructure:"compress"`
}