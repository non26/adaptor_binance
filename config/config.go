package serviceconfig

import "github.com/spf13/viper"

func ReadConfig() (c *ServiceConfig, err error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")
	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		return nil, err
	}
	viper.Unmarshal(&c)
	return c, nil
}
