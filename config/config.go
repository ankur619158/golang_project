package config

type Config struct {
	dbConfig  *databaseConfig
	srvConfig *serverConfig
}

func NewConfig() (*Config, error) {
	srvCfg, srvConfigErr := newServerConfig()
	if srvConfigErr != nil {
		return nil, srvConfigErr
	}
	dbCfg, dbConfigErr := newDatabaseConfig(srvCfg.env)
	if dbConfigErr != nil {
		return nil, dbConfigErr
	}

	return &Config{
		dbConfig: dbCfg,
	}, nil
}

func (c *Config) GetPort() string {
	return c.srvConfig.port
}

func (c *Config) GetEnv() string {
	return c.srvConfig.env
}

func (c *Config) GetDBDriver() string {
	return c.dbConfig.dbDriver
}

func (c *Config) GetDBUser() string {
	return c.dbConfig.dbUser
}

func (c *Config) GetDBHost() string {
	return c.dbConfig.dbHost
}

func (c *Config) GetDBName() string {
	return c.dbConfig.dbName
}

func (c *Config) GetConnStringWithDB() string {
	return c.dbConfig.setupConnectionString()
}
