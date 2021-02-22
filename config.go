package really

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"os"
	"path/filepath"
)

// configuration 总配置
type configuration struct {
	Mirai *miraiConfig `mapstructure:"mirai"`
}

// miraiConfig Dota Max 相关配置
type miraiConfig struct {
	BotQQ string `mapstructure:"botqq"`
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
			log.Fatalf("保存配置文件失败: %+v", err)
		}
	}

	var c configuration
	err = viper.Unmarshal(&c)
	if err != nil {
		log.Fatalf("读取配置文件失败: %+v", err)
	}

	c.check()

	return &c
}

func (c *configuration) SaveConfig() {
	err := viper.WriteConfig()
	if err != nil {
		log.Fatalf("保存配置文件失败: %+v", err)
	}
}

func createDefaultConfigFile() {
	viper.Set("botqq", miraiConfig{})

	err := viper.SafeWriteConfig()
	if err != nil {
		log.Fatalf("初始化配置文件失败: %+v", err)
	}
	userHomeDir, _ := os.UserHomeDir()
	fmt.Println("请填写配置文件: " + filepath.Join(userHomeDir, ".really.yaml"))
	os.Exit(0)
}
