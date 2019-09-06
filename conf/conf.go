package conf

var (
	Conf config // holds the global app config.
)

type config struct {
	// 应用配置
	App app

	//Jwt
	Jwt jwt `toml:"jwt"`

	// MySQL
	DB database `toml:"db"`

	// 静态资源
	Static static

	// Redis
	Redis redis
}

type app struct {
	Name string `toml:"name"`
}

type jwt struct {
	Secret string `toml:"secret"`
}

type static struct {
	Type string `toml:"type"`
}

type database struct {
	DataBase string `toml:"database"`
	User     string `toml:"user"`
	Pwd      string `toml:"password"`
	Host     string `toml:"host"`
	Port     string `toml:"port"`
}

type redis struct {
	Host string `toml:"host"`
	Port string `toml:"port"`
	Auth string `toml:"auth"`
}
