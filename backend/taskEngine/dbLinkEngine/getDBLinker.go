/**
 * @date: 2022/3/9
 * @desc: ...
 */

package dbLinkEngine

import (
	"errors"
)

//
// GetDBLinker
// @Description: 返回数据库连接
// @param dbType:
// @param dbOptions:
// @return dbLinker:
// @return err:
//
func GetDBLinker(dbOptions DataBaseOption) (dbLinker DBLinker, err error) {
	switch dbOptions.DBType {
	case "vertica":
		dbLinker, err = NewVerticaLink(dbOptions)
	case "oracle":
		dbLinker, err = NewOracleLink(dbOptions)
	case "mysql":
		dbLinker, err = NewMysqlLink(dbOptions)
	case "postgres":
		dbLinker, err = NewPostgresLink(dbOptions)
	default:
		dbLinker, err = nil, errors.New("不支持的数据库类型")
	}
	if err != nil {
		return nil, err
	}
	return dbLinker, nil
}
