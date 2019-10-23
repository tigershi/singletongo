package bundle

import (
	"github.com/tigershi/singletongo/example/level3"
)

func init() {
	level3.InitComponentDao(GetbundleComponentDaoInstance())
}
