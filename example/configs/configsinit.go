// configs project configs.go
package configs

import (
	"io/ioutil"
	"log"
	"path/filepath"

	"github.com/tigershi/singletongo/example/models"
	"gopkg.in/yaml.v2"
)

var AppConfig *models.AppConfigProps

func init() {
	loadAppConfig()
}
func loadAppConfig() {
	configPath := "configs" + string(filepath.Separator) + "app.yml"
	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		panic("read app config path error\n" + err.Error())
	}
	log.Println("begin read configuration file")
	AppConfig = new(models.AppConfigProps)
	//把yaml形式的字符串解析成struct类型
	yaml.Unmarshal(data, AppConfig)

}
