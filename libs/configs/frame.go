package configs

type Config struct {
	Mode     string   `yaml:"mode"`
	Server   Server   `yaml:"server"`
	Dingtalk Dingtalk `yaml:"dingtalk"`
	Log      Log      `yaml:"log"`
	Database Database `yaml:"databases"`
	Host     Host     `yaml:"host"`
}
type Server struct {
	AppName string `yaml:"AppName"`
	Address string `yaml:"address"`
	Port    string `yaml:"port"`
}
type Dingtalk struct {
	Host    string `yaml:"host"`
	Webhook string `yaml:"webhook"`
	Secret  string `yaml:"secret"`
}
type Log struct {
	AccessKeyID     string `yaml:"AccessKeyID"`
	AccessKeySecret string `yaml:"AccessKeySecret"`
	LogStore        string `yaml:"LogStore"`
	Endpoint        string `yaml:"Endpoint"`
	Project         string `yaml:"Project"`
	LookAddr        string `yaml:"LookAddr"`
}
type Mysql struct {
	Username     string `yaml:"username"`
	Password     string `yaml:"password"`
	Address      string `yaml:"address"`
	Dbname       string `yaml:"dbname"`
	Asname       string `yaml:"asname"`
	MaxOpenConns int    `yaml:"maxOpenConns"`
}
type Redis struct {
	Address      string `yaml:"Address"`
	Port         int    `yaml:"Port"`
	Db           int    `yaml:"Db"`
	Password     string `yaml:"Password"`
	PoolSize     int    `yaml:"PoolSize"`
	MinIdleConns int    `yaml:"MinIdleConns"`
	DialTimeout  int    `yaml:"DialTimeout"`
	Asname       string `yaml:"asname"`
}
type Database struct {
	Mysql []Mysql `yaml:"mysql"`
	Redis []Redis `yaml:"redis"`
}
type Host struct {
	OssUpload string `yaml:"ossUpload"`
}
