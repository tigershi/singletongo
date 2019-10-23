// bundle project bundle.go
package bundle

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"

	log "github.com/sirupsen/logrus"
	"github.com/tigershi/singletongo/example/models"
)

type bundleComponentDaoImpl struct {
}

var instance *bundleComponentDaoImpl
var bundleComponentDaoImplOnce sync.Once

func GetbundleComponentDaoInstance() *bundleComponentDaoImpl {
	bundleComponentDaoImplOnce.Do(func() {
		instance = &bundleComponentDaoImpl{}
	})
	return instance
}
func (b *bundleComponentDaoImpl) GetComponentTranslation(productName string, version string, component string, language string) (*models.ComponentStore, error) {

	targPath := getTargetFile(productName, version, component, language)
	filePtr, err := os.Open(targPath)

	if err != nil {

		return nil, err
	}
	defer filePtr.Close()

	result := new(models.ComponentStore)

	decodermd := json.NewDecoder(filePtr)
	err = decodermd.Decode(result)
	if err != nil {

		return nil, err
	}

	return result, nil
}
func (b *bundleComponentDaoImpl) UpdateComponentTranslation(compContext *models.ComponentStore) (*models.ComponentStore, error) {

	targPath := getTargetFile(compContext.Product, compContext.Version, compContext.Component, compContext.Language)

	log.WithFields(log.Fields{
		"targetPath": targPath,
	}).Info("--------------component bundle store---------")

	if !checkFileIsExist(targPath) {

		compJson, err2 := json.MarshalIndent(compContext, "", "   ")
		if err2 != nil {
			log.WithFields(log.Fields{
				"targetPath": targPath,
			}).Error("--------------convert json Error---------")

		}

		if err9 := ioutil.WriteFile(targPath, compJson, 0666); err9 != nil {
			fmt.Println(targPath)
			dirPath, _ := filepath.Split(targPath)
			fmt.Println(dirPath)
			errmk := os.MkdirAll(dirPath, 0777)

			if errmk != nil {
				fmt.Println(errmk)
			}
			if err7 := ioutil.WriteFile(targPath, compJson, 0666); err7 != nil {
				log.WithFields(log.Fields{
					"targetPath": string(compJson),
				}).Error("--------------write json json Error---------")
				return nil, err7
			}
		}
		return compContext, nil
	} else {
		filePtr, err3 := os.Open(targPath)
		if err3 != nil {
			log.WithFields(log.Fields{
				"targetPath": targPath,
			}).Error("--------------open component bundle store Error---------")
			return nil, err3
		}
		queryComp := new(models.ComponentStore)
		decodermd := json.NewDecoder(filePtr)
		err4 := decodermd.Decode(queryComp)
		if err4 != nil {
			return nil, err4
		}
		filePtr.Close()
		for k, v := range compContext.Messages {
			queryComp.Messages[k] = v
		}
		mixJson, mixerr := json.MarshalIndent(queryComp, "", "   ")
		if mixerr != nil {
			return nil, mixerr
		}
		if err6 := ioutil.WriteFile(targPath, mixJson, 0666); err6 != nil {
			return nil, err6
		}
		return queryComp, nil
	}

}

func (b *bundleComponentDaoImpl) DelComponentTranslation(productName string, version string, component string, language string) error {

	targPath := getTargetFile(productName, version, component, language)
	return os.Remove(targPath)

}
