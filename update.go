package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Println("参数为空")
		return
	}
	exe := args[0]
	updateTmp := exe + ".tmp"
	time.Sleep(1 * time.Second)
	if !fileExists(exe) {
		fmt.Println("不存在: ", exe)
		return
	}
	if !fileExists(updateTmp) {
		fmt.Println("不存在: ", updateTmp)
		return
	}
	err := os.Remove(exe)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = os.Rename(updateTmp, exe)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = exec.Command(exe, args[1:]...).Run()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("更新成功")
	time.Sleep(3 * time.Second)
}

func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return true
}
