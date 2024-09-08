package config

type Config struct {
	Server Server `yaml:"Server"`
	Jwt    Jwt    `yaml:"Jwt"`
	Log    Log    `yaml:"Log"`
	Mysql  Mysql  `yaml:"Mysql"`
	Redis  Redis  `yaml:"Redis"`
}

type Server struct {
	Name string `yaml:"Name"`
	Host string `yaml:"Host"`
	Port int    `yaml:"Port"`
}

type Jwt struct {
	Secret string `yaml:"Secret"`
	Expire int    `yaml:"Expire"`
}

type Log struct {
	Dir   string `yaml:"Dir"`
	Level string `yaml:"Level"`
}

type Mysql struct {
	Host string `yaml:"Host"`
	Port int    `yaml:"Port"`
	User string `yaml:"User"`
	Pwd  string `yaml:"Pwd"`
	Db   string `yaml:"Db"`
}

type Redis struct {
	Host string `yaml:"Host"`
	Port int    `yaml:"Port"`
	User string `yaml:"User"`
	Pwd  string `yaml:"Pwd"`
	Db   int    `yaml:"Db"`
}
