package main

import (
	"fmt"
	"os"
	"time"

	"viper/configuration"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// 환경변수는 PROFILE을 확인하기 위해 하나만 설정합니다.
func initProfile() string {
	var profile string
	profile = os.Getenv("GO_PROFILE")

	if len(profile) <= 0 {
		profile = "local"
	}
	fmt.Println("GOLANG_PROFILE: " + profile)
	return profile
}

// PROFILE을 기반으로 config파일을 읽고 전역변수에 언마샬링 해줍니다.
func setRuntimeConfig(profile string) {

	// 환경변수에서 읽어온 profile이름의 yaml파일을 configPath로 설정합니다.
	// viper.SetConfigName(profile)
	viper.SetConfigName("local")
	viper.SetConfigType("yaml")

	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	// viper는 읽어온 설정파일의 정보를 가지고있으니, 전역변수에 언마샬링해
	// 애플리케이션의 원하는곳에서 사용하도록 합니다.
	err = viper.Unmarshal(&configuration.RuntimeConf)
	if err != nil {
		panic(err)
	}

	viper.WatchConfig()

	// 일반적인 프로필 설정은 이 위까지 적용하시면 됩니다.
	// viper는 설정파일이 변경된 이벤트를 핸들링할수 있습니다.
	// 설정파일이 변경되면 다시 언마샬링해줍니다.

	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
		var err error
		err = viper.ReadInConfig()
		if err != nil {
			fmt.Println(err)
			return
		}
		err = viper.Unmarshal(&configuration.RuntimeConf)
		if err != nil {
			fmt.Println(err)
			return
		}
	})

	viper.Get("server")
}

func main() {

	// profile := initProfile()

	// setRuntimeConfig(profile)
	setRuntimeConfig("meel")

	// viper.AddConfigPath(".")

	// viper.SetDefault("ContentDir", "content")

	// err := viper.ReadInConfig() // Find and read the config file
	// if err != nil {             // Handle errors reading the config file
	// 	panic(fmt.Errorf("fatal error config file: %w", err))
	// }

	fmt.Println("db type: ", configuration.RuntimeConf.Datasource.DbType)
	for {
		<-time.After(time.Second * 3)
		fmt.Println("db type: ", configuration.RuntimeConf.Datasource.DbType)
	}

}
