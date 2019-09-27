package models

type AppConfigProps struct {
	BundleConfig BundleProperty `yaml:"bundleConfig"`
}

//bundle config
type BundleProperty struct {
	BasePath string `yaml:"basePath"`
}
