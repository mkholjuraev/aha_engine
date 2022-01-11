package utils

import "github.com/spf13/viper"

type Config struct {
	DBScheme              string `mapstructure:"DB_SCHEME"`
	DBUsername            string `mapstructure:"DB_USERNAME"`
	DBPassword            string `mapstructure:"DB_PASSWORD"`
	DBDatabase            string `mapstructure:"DB_DATABASE"`
	DBHost                string `mapstructure:"DB_HOST"`
	DBPort                int    `mapstructure:"DB_PORT"`
	DBSync                bool   `mapstructure:"DB_SYNC"`
	DB_IS_LOGGING_ENALBED bool   `mapstructure:"DB_IS_LOGGING_ENALBED"`
}

func LoadDbConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("db_config")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
