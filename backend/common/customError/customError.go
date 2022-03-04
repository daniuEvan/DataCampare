/**
 * @date: 2022/3/3
 * @desc: ...
 */

package customError

import "errors"

var (
	NotImplementedError      = errors.New("该功能暂未实现")
	InternalServerError      = errors.New("服务异常")
	BadRequestError          = errors.New("请求失败")
	ResourceNotFountError    = errors.New("找不到资源")
	UnprocessableEntityError = errors.New("对象不存在")
	DatabaseConnectError     = errors.New("数据库连接异常")
)
