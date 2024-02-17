package config

type Configurations struct {
	Server       ServerConfig
	Database     DatabaseConfig
	TestDatabase TestDatabaseConfig
	Environment  string
}

type ServerConfig struct {
	Port string
}

type DatabaseConfig struct {
	DBHost     string
	DBName     string
	DBUser     string
	DBPassword string
	Port       string
}

type TestDatabaseConfig struct {
	TestDBHost     string
	TestDBName     string
	TestDBUser     string
	TestDBPassword string
	TestPort       string
}
