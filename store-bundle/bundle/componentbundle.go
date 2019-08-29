// bundle project bundle.go
package bundle

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
func (b *bundleComponentDaoImpl) getComponentTranslation(productName string, version string, component string, language string) (models.ComponentStore, error) {

}
func (b *bundleComponentDaoImpl) UpdateComponentTranslation(compContent models.ComponentStore) (models.ComponentStore, error) {

}
func (b *bundleComponentDaoImpl) DelComponentTranslation(productName string, version string, component string, language string) error {

}
