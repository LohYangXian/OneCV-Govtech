package config

type Configurations struct {
	Server      ServerConfig
	Database    DatabaseConfig
	Environment string
}

type ServerConfig struct {
	Port string
}

type DatabaseConfig struct {
	DBName     string
	DBUser     string
	DBPassword string
	Port       string
}
