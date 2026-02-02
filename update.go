package main

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		writeLog("参数为空", "ERROR")
		return
	}
	exe := args[0]
	updateTmp := exe + ".tmp"
	time.Sleep(2 * time.Second)
	if !fileExists(exe) {
		writeLog("不存在: "+exe, "ERROR")
		return
	}
	if !fileExists(updateTmp) {
		writeLog("不存在: "+updateTmp, "ERROR")
		return
	}
	err := os.Remove(exe)
	if err != nil {
		writeLog("删除文件错误: "+err.Error(), "ERROR")
		return
	}
	err = os.Rename(updateTmp, exe)
	if err != nil {
		writeLog("重命名错误: "+err.Error(), "ERROR")
		return
	}

	err = exec.Command(exe, args[1:]...).Run()
	if err != nil {
		writeLog("重启程序错误: "+err.Error(), "ERROR")
		return
	}
	writeLog("更新成功", "INFO")

	// 等待3秒, 确保程序已退出
	time.Sleep(2 * time.Second)
}

func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return true
}

func writeLog(s string, level string) {
	exePath, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}

	dir := filepath.Dir(exePath)

	logfile := filepath.Join(dir, "ani-rss-update.log")

	timeStr := time.Now().Format("2006-01-02 15:04:05")

	s = timeStr + "\t" + level + "\t" + s + "\n"

	file, err := os.OpenFile(logfile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)

	_, err = file.WriteString(s)
	if err != nil {
		log.Fatal(err)
	}
}
