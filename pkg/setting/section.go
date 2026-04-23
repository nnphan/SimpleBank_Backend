package setting

type Config struct {
	Postgresql PostgreSQLSettings `mapstructure:"postgresql"`
}

type PostgreSQLSettings struct {
	Driver   string `mapstructure:"driver"`
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	DBName     string `mapstructure:"dbname"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	SSLMode  string `mapstructure:"sslmode"`
	MaxIdleConns int    `mapstructure:"max_idle_conns"`
	MaxOpenConns int    `mapstructure:"max_open_conns"`
	ConnMaxLifetime int   `mapstructure:"conn_max_lifetime"`
}