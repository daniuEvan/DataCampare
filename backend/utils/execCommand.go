/**
 * @date: 2022/3/3
 * @desc: ...
 */

package utils

import (
	"io/ioutil"
	"os/exec"
	"runtime"
)

//
// ExecCommand
// @Description: 执行系统命令
// @param arg:
// @return output:
// @return err:
//
func ExecCommand(args ...string) (output string, err error) {
	name := "/bin/bash"
	c := "-c"
	// 根据系统设定不同的命令name
	if runtime.GOOS == "windows" {
		name = "cmd"
		c = "/C"
	}
	argList := append([]string{c}, args...)
	cmd := exec.Command(name, argList...)

	//创建获取命令输出管道
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return
	}
	//执行命令
	if err = cmd.Start(); err != nil {
		return
	}
	//读取所有输出
	bytes, err := ioutil.ReadAll(stdout)
	if err != nil {
		return
	}
	if err = cmd.Wait(); err != nil {
		return
	}
	output = string(bytes)
	return
}
