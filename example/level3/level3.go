// leve3.go
package level3

import "github.com/tigershi/singletongo/example/dao"

var componentDao dao.ComponentDao

func InitComponentDao(dao dao.ComponentDao) {
	componentDao = dao
}
