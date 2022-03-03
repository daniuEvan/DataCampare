/**
 * @date: 2022/3/3
 * @desc: ...
 */

package customError

import "errors"

var (
	NotImplementedError = errors.New("该功能暂未实现")
)
