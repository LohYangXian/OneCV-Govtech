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
	DBName     string
	DBUser     string
	DBPassword string
	Port       string
}

type TestDatabaseConfig struct {
	TestDBName     string
	TestDBUser     string
	TestDBPassword string
	TestPort       string
}
