// bundle project bundle.go
package bundle

import (
	"os"
	"path/filepath"

	"github.com/tigershi/singletongo/service-core/configs"
)

func checkFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

func getTargetFile(productName string, version string, component string, language string) string {
	sep := string(filepath.Separator)
	targPath := configs.AppConfig.BundleConfig.BasePath + sep + "l10n" + sep + bundles + sep + productName + sep + version + sep + component + sep + "messages_" + language + ".json"
	return targPath
}
