package config

type XcbbRobotConfig struct {
	Loglevel int `yaml:"Loglevel"`
	Logfile string `yaml:"Logfile"`
	Robotlist string `yaml:"Robotlist"`
	Server string `yaml:"Server"`
}

type Testccconf struct {
	Host string `yaml:"host"`
	User string `yaml:"user"`
	Pwd string `yaml:"pwd"`
	Dbname string `yaml:"dbname"`
}


