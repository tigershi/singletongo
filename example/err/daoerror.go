//define the singleton dao
package err

import "github.com/tigershi/singletongo/example/models"

//bundle module error
var (
	NO_JSON_FILE_ERR = models.NewSingtonError(100000, "not found the json file")
)
