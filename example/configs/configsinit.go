// configs project configs.go
package configs

import (
	"io/ioutil"
	"path/filepath"

	log "github.com/cihub/seelog"
	"github.com/tigershi/singletongo/example/models"
	"gopkg.in/yaml.v2"
)

var AppConfig *models.AppConfigProps
var Logger log.LoggerInterface

func init() {
	initLog()
	loadAppConfig()
}
func loadAppConfig() {
	configPath := "configs" + string(filepath.Separator) + "app.yml"
	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		panic("read app config path error\n" + err.Error())
	}
	Logger.Info("begin read configuration file")
	AppConfig = new(models.AppConfigProps)
	//把yaml形式的字符串解析成struct类型
	yaml.Unmarshal(data, AppConfig)

}

func initLog() {
	var err error
	logPath := "configs" + string(filepath.Separator) + "seelog.xml"
	Logger, err = log.LoggerFromConfigAsFile(logPath)
	if err != nil {
		log.Errorf("parse config.xml error: %v", err)
		return
	}
}
