package main

import (
	"errors"
	"fmt"
	"os/exec"
	"runtime"
    "strings"
    "time"
)

//  收集linux信息
func collect_linux() (string, error) {
	cmd := exec.Command("dmidecode", "-t", "system")
	// cmd := exec.Command("pwd")

	out, _ := cmd.Output()

	for _, line := range strings.Split(string(out), "\n") {
		if strings.Contains(line, "Serial") {
			//  取出首位空格
			s := strings.Split(line, ":")
			sn := s[len(s)-1]
			sn = strings.TrimSpace(sn)
			return sn, nil
		}
	}
	return "", errors.New("未匹配到序列号")
}


func main() {
	os := runtime.GOOS
	if os == "windows" {
        fmt.Printf("windows")
        time.Sleep(time.Duration(2) * time.Second)
	} else if os == "darwin" {
		fmt.Printf("苹果")
	} else if os == "linux" { //  处理linux系统相关
		sn, err := collect_linux()
		if err == nil {
			fmt.Println(sn)
		} else {
			fmt.Println(err)
		}
	} else {
		fmt.Printf("找不到操作系统%s, 请联系管理员!\n", os)
	}
}
