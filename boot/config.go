package boot

import (
	"github.com/spf13/viper"
	"mermaid/global"
)

func bootConfig() error {
	vip := viper.New()
	vip.SetConfigName("config")
	vip.SetConfigType("yaml")
	vip.AddConfigPath("./config/")
	err := vip.ReadInConfig()
	if err != nil {
		return err
	}
	err = vip.Unmarshal(global.Conf)
	if err != nil {
		return err
	}
	return nil
}
