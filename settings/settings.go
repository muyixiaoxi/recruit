package settings

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var Conf = new(AppConfig)

type AppConfig struct {
	*ProjectConfig `mapstructure:"project"`
	*LogConfig     `mapstructure:"log"`
	*MySQLConfig   `mapstructure:"mysql"`
	*RedisConfig   `mapstructure:"redis"`
	*UploadImg     `mapstructure:"uploadImg"`
}
type ProjectConfig struct {
	Name string `mapstructure:"name"`
	Mode string `mapstructure:"mode"`
	Port int    `mapstructure:"port"`
}

type MySQLConfig struct {
	Host         string `mapstructure:"host"`
	User         string `mapstructure:"user"`
	Password     string `mapstructure:"password"`
	DB           string `mapstructure:"dbname"`
	Port         int    `mapstructure:"port"`
	MaxOpenConns int    `mapstructure:"max_open_conns"`
	MaxIdleConns int    `mapstructure:"max_idle_conns"`
}

type RedisConfig struct {
	Host         string `mapstructure:"host"`
	Password     string `mapstructure:"password"`
	Port         int    `mapstructure:"port"`
	DB           int    `mapstructure:"db"`
	PoolSize     int    `mapstructure:"pool_size"`
	MinIdleConns int    `mapstructure:"min_idle_conns"`
}

type LogConfig struct {
	Level      string `mapstructure:"level"`
	Filename   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxAge     int    `mapstructure:"max_age"`
	MaxBackups int    `mapstructure:"max_backups"`
}
type UploadImg struct {
	ImageSavePath   string   `mapstructure:"imageSavePath"`
	ImagePrefixUrl  string   `mapstructure:"imagePrefixUrl"`
	ImageMaxSize    int      `mapstructure:"imageMaxSize"`
	ImageAllowExits []string `mapstructure:"imageAllowExits"`
	TimeFormat      string   `mapstructure:"timeFormat"`
	RuntimeRootPath string   `mapstructure:"runtimeRootPath"`
}

func Init() error {
	viper.SetConfigFile("./config.yaml")

	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("夭寿啦~配置文件被人修改啦...")
		viper.Unmarshal(&Conf)
	})

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Errorf("ReadInConfig failed, err: %v", err)
		return err
	}
	if err := viper.Unmarshal(&Conf); err != nil {
		fmt.Errorf("unmarshal to Conf failed, err:%v", err)
		return err
	}
	return err
}
