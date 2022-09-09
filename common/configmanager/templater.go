package configmanager

type BaseConfig struct {
	Data struct {
		Mysql MysqlConfig
		Mongo MongoConfig
		Redis RedisConfig
	}

	Server struct {
	}
}



type MysqlConfig struct {
	Host string
}

type MongoConfig struct {
	Host           string   `json:"host"`
	DataBase       string   `json:"data_base"`
	ConnectTimeout Duration `json:"connect_timeout,default=3s"`
}

type RedisConfig struct {
	Host string
}
