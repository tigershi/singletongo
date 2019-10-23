// bundle project bundle.go
package bundle

import (
	"os"
	"path/filepath"

	"github.com/tigershi/singletongo/example/configs"
)

func checkFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

func getTargetFile(productName string, version string, component string, language string) string {
	path, _ := filepath.Abs(configs.AppConfig.BundleConfig.BasePath)
	targPath := filepath.Join(path, "l10n", "bundles", productName, version, component, "messages_"+language+".json")
	return targPath
}
