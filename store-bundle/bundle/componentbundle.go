// bundle project bundle.go
package bundle

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/tigershi/singletongo/service-core/configs"
	"github.com/tigershi/singletongo/service-facade/models"
)

type bundleComponentDaoImpl struct {
}

var instance *bundleComponentDaoImpl
var once sync.Once

func GetbundleComponentDaoInstance() *bundleComponentDaoImpl {
	once.Do(func() {
		instance = &bundleComponentDaoImpl{}
	})
	return instance
}
func (b *bundleComponentDaoImpl) GetComponentTranslation(productName string, version string, component string, language string) (*models.ComponentStore, error) {

	targPath := getTargetFile(productName, version, component, language)
	filePtr, err := os.Open(targPath)
	if err != nil {
		error = err
		return
	}
	defer filePtr.Close()

	var result models.ComponentStore

	decodermd := json.NewDecoder(filePtr)
	err = decodermd.Decode(&result)
	if err != nil {

		return nil, err
	}

	return &result, nil
}
func (b *bundleComponentDaoImpl) UpdateComponentTranslation(compContext *models.ComponentStore) (*models.ComponentStore, error) {
	targPath := getTargetFile(compContext.Product, compContext.Version, compContext.Component, compContext.Language)
	filePtr, err := os.Open(targPath)
	if err != nil && os.IfNotExist(err) {
		filePtr, err1 = os.Create(targPath)
		if err1 != nil {
			return nil, err1
		}
		compJson, err2 := json.MarshalIndent(compContext, "", "   ")
		if err2 != nil {
			return nil, err2
		}
		filePtr.WriteString(compJson)
		return compContext, nil
	}
	defer filePtr.Close()
	var queryComp models.ComponentStore

	decodermd := json.NewDecoder(filePtr)
	err = decodermd.Decode(&queryComp)
	if err != nil {
		return nil, err
	}
	for k, v := range compContext.Messages {
		queryComp.Messages[k] = v
	}
	mixJson, mixerr := json.MarshalIndent(queryComp, "", "   ")
	if mixerr != nil {
		return nil, mixerr
	}
	filePtr.WriteString(mixJson)
	return &queryComp, nil
}

func (b *bundleComponentDaoImpl) DelComponentTranslation(productName string, version string, component string, language string) error {

	targPath := getTargetFile(productName, version, component, language)
	return os.Remove(targPath)

}
