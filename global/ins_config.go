package global

var Conf = &Config{}

type Config struct {
	Db     *db
	Redis  *redis
	Log    *log
	Server *server
}

type db struct {
	Type        string
	Dsn         string
	MaxIdle     int
	MaxOpen     int
	LogMode     bool
	MaxLifeTime int
}

type redis struct {
	Host string
}

type log struct {
	Dir   string
	Level string
}

type server struct {
	Addr string
}
