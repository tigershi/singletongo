package bundle

import (
	"github.com/tigershi/singletongo/service-core/level3"
)

func init() {
	level3.InitComponentDao(GetbundleComponentDaoInstance())
}
