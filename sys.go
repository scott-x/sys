/*
* @Author: scottxiong
* @Date:   2020-10-23 11:30:17
* @Last Modified by:   scottxiong
* @Last Modified time: 2020-10-23 14:27:14
 */
package sys

import (
	"os/exec"
	"runtime"
	"strings"
)

func runInWindows(cmd string) (string, error) {
	result, err := exec.Command("cmd", "/c", cmd).Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(result)), err
}

func RunCommand(cmd string) (string, error) {
	if runtime.GOOS == "windows" {
		return runInWindows(cmd)
	} else {
		return runInLinux(cmd)
	}
}

func runInLinux(cmd string) (string, error) {
	result, err := exec.Command("/bin/sh", "-c", cmd).Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(result)), err
}

//check if the process is running by given name
func IsProcessRunning(serverName string) (bool, error) {
	// a := `ps ux | awk '/` + serverName + `/ && !/awk/ {print $2}'` => works on mac, not linux
	a := `ps -ef | awk '/` + serverName + `/ && !/awk/ {print $2}'`
	pid, err := RunCommand(a)
	if err != nil {
		return false, err
	}
	return pid != "", nil
}

//get the process id by given name
func GetPid(serverName string) (string, error) {
	a := `ps ux | awk '/` + serverName + `/ && !/awk/ {print $2}'`
	pid, err := RunCommand(a)
	return pid, err
}
