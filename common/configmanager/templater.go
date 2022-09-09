package configmanager

type BaseConfig struct {
	Data struct {
		Mysql MysqlConfig `json:"mysql"`
		Mongo MongoConfig `json:"mongo"`
		Redis RedisConfig `json:"redis"`
	} `json:"data"`
	Server struct {
		Http HttpConfig `json:"http"`
		Grpc GrpcConfig `json:"grpc"`
	} `json:"server"`
}

type (
	HttpConfig struct {
		Addr    string   `json:"addr"`
		Timeout Duration `json:"timeout,default=2s"`
	}
	GrpcConfig struct {
		Addr    string   `json:"addr"`
		Timeout Duration `json:"timeout,default=2s"`
	}
)

type (
	MysqlConfig struct {
		Host string `json:"host"`
	}
	MongoConfig struct {
		Host           string   `json:"host"`
		DataBase       string   `json:"data_base"`
		ConnectTimeout Duration `json:"connect_timeout,default=3s"`
	}
	RedisConfig struct {
		Host string `json:"host"`
	}
)
