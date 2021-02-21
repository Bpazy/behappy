package really

import (
	"encoding/json"
	"fmt"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"reflect"
	"time"
)

// configuration 总配置
type configuration struct {
	DotaMax *dotaMaxConfig `mapstructure:"dotamax"`
}

// dotaMaxConfig Dota Max 相关配置
type dotaMaxConfig struct {
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Cookies  string `mapstructure:"cookies"`
}

// check 校验配置文件必填项
func (c *configuration) check() {

}
func ToTimeHookFunc() mapstructure.DecodeHookFunc {
	return func(
		f reflect.Type,
		t reflect.Type,
		data interface{}) (interface{}, error) {
		if t != reflect.TypeOf(time.Time{}) {
			return data, nil
		}

		switch f.Kind() {
		case reflect.String:
			return time.Parse(time.RFC3339, data.(string))
		case reflect.Float64:
			return time.Unix(0, int64(data.(float64))*int64(time.Millisecond)), nil
		case reflect.Int64:
			return time.Unix(0, data.(int64)*int64(time.Millisecond)), nil
		default:
			return data, nil
		}
		// Convert it by parsing
	}
}

func Decode(input map[string]interface{}, result interface{}) error {
	decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		Metadata: nil,
		DecodeHook: mapstructure.ComposeDecodeHookFunc(
			ToTimeHookFunc()),
		Result: result,
	})
	if err != nil {
		return err
	}

	if err := decoder.Decode(input); err != nil {
		return err
	}
	return err
}
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

func (c *configuration) SetCookies(cookies []*http.Cookie) {
	jsonCookies, err := json.Marshal(cookies)
	if err != nil {
		log.Fatalf("序列化 JSON 错误: %+v", err)
	}
	c.DotaMax.Cookies = string(jsonCookies)
	viper.Set("dotamax", c.DotaMax)

	c.SaveConfig()
}

func (c *configuration) SaveConfig() {
	err := viper.WriteConfig()
	if err != nil {
		log.Fatalf("保存配置文件失败: %+v", err)
	}
}

func createDefaultConfigFile() {
	viper.Set("dotamax", dotaMaxConfig{})

	err := viper.SafeWriteConfig()
	if err != nil {
		log.Fatalf("初始化配置文件失败: %+v", err)
	}
	userHomeDir, _ := os.UserHomeDir()
	fmt.Println("请填写配置文件: " + filepath.Join(userHomeDir, ".really.yaml"))
	os.Exit(0)
}
