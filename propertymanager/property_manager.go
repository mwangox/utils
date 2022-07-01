package propertymanager

import (
	"flag"
	"log"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var ConfigDir string

func Initialize() {

	flag.StringVar(&ConfigDir, "config-dir", "./conf", "Configuration directory")
	flag.Parse()

	viper.SetConfigName("application")
	viper.AddConfigPath(ConfigDir)
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Println("Config file changed:", e.Name)
	})

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Failed to read config file:", err)
	}
}

func GetStringProperty(key string, defaultValue ...string) string {
	keyValue := viper.GetString(key)
	if keyValue == "" && len(defaultValue) != 0 {
		return defaultValue[0]
	}
	return keyValue
}

func GetIntProperty(key string, defaultValue ...int) int {
	keyValue := viper.GetInt(key)
	if keyValue == 0 && len(defaultValue) != 0 {
		return defaultValue[0]
	}
	return keyValue
}

func GetBoolProperty(key string) bool {
	return viper.GetBool(key)
}
