/*
* @Author: scottxiong
* @Date:   2020-10-23 11:30:17
* @Last Modified by:   scottxiong
* @Last Modified time: 2020-10-23 14:27:14
 */
package sys

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
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
	if runtime.GOOS == "windows" {
		flag, _, _ := isProcessExistOnWindows(serverName)
		return flag, nil
	}
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
	if runtime.GOOS == "windows" {
		_, _, pid := isProcessExistOnWindows(serverName)
		return strconv.Itoa(pid), nil
	}
	a := `ps -ef | awk '/` + serverName + `/ && !/awk/ {print $2}'`
	pid, err := RunCommand(a)
	return pid, err
}

func isProcessExistOnWindows(appName string) (bool, string, int) {
	appary := make(map[string]int)
	cmd := exec.Command("cmd", "/C", "tasklist")
	output, _ := cmd.Output()
	//fmt.Printf("fields: %v\n", output)
	n := strings.Index(string(output), "System")
	if n == -1 {
		fmt.Println("no find")
		os.Exit(1)
	}
	data := string(output)[n:]
	fields := strings.Fields(data)
	for k, v := range fields {
		if v == appName {
			appary[appName], _ = strconv.Atoi(fields[k+1])

			return true, appName, appary[appName]
		}
	}

	return false, appName, -1
}
