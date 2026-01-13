package config

type DB struct {
	Host     string `env:"POSTGRESQL_HOST,required"`
	Port     int    `env:"POSTGRESQL_PORT,required"`
	Username string `env:"POSTGRESQL_USERNAME,required"`
	Password string `env:"POSTGRESQL_PASSWORD,required"`
	Database string `env:"POSTGRESQL_DATABASE,required"`
	Schema   string `env:"POSTGRESQL_SCHEMA,default=public"`
	Timezone string `env:"POSTGRESQL_TIMEZONE,default=UTC"`
	LogSQL   bool   `env:"POSTGRESQL_LOG_SQL,default=false"`
}
