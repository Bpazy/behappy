package really

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
)

// configuration 总配置
type configuration struct {
	Mirai      *MiraiConfig      `mapstructure:"mirai"`
	SteamAPI   *SteamAPIConfig   `mapstructure:"steamapi"`
	DataSource *DataSourceConfig `mapstructure:"datasource"`
}

type MiraiConfig struct {
	BotQQ string `mapstructure:"botqq"`
}

type SteamAPIConfig struct {
	Key string `mapstructure:"key"`
}

type DataSourceConfig struct {
	Url string `mapstructure:"url"`
}

// check 校验配置文件必填项
func (c *configuration) check() {

}

var config = InitConfig()

func InitConfig() *configuration {
	viper.SetConfigName(".really")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("$HOME")
	err := viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			createDefaultConfigFile()
		} else {
			logrus.Fatalf("保存配置文件失败: %+v", err)
		}
	}

	var c configuration
	err = viper.Unmarshal(&c)
	if err != nil {
		logrus.Fatalf("读取配置文件失败: %+v", err)
	}

	c.check()

	return &c
}

func (c *configuration) SaveConfig() {
	err := viper.WriteConfig()
	if err != nil {
		logrus.Fatalf("保存配置文件失败: %+v", err)
	}
}

func createDefaultConfigFile() {
	viper.Set("mirai", MiraiConfig{})
	viper.Set("steamapi", SteamAPIConfig{})
	viper.Set("datasource", DataSourceConfig{})

	err := viper.SafeWriteConfig()
	if err != nil {
		logrus.Fatalf("初始化配置文件失败: %+v", err)
	}
	userHomeDir, _ := os.UserHomeDir()
	fmt.Println("请填写配置文件: " + filepath.Join(userHomeDir, ".really.yaml"))
	os.Exit(0)
}
