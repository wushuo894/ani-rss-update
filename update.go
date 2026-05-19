package main

import (
	"flag"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

func main() {
	var source string
	var target string
	var args string

	flag.StringVar(&source, "source", "", "源文件路径")
	flag.StringVar(&target, "target", "", "目标文件路径")
	flag.StringVar(&args, "args", "", "启动参数")
	flag.Parse()

	writeLog("==========", "INFO")
	writeLog("source: "+source, "INFO")
	writeLog("target: "+target, "INFO")
	writeLog("args: "+args, "INFO")

	if source == "" || target == "" {
		writeLog("source 和 target 参数不能为空", "ERROR")
		return
	}

	time.Sleep(2 * time.Second)

	if !fileExists(source) {
		writeLog("不存在: "+source, "ERROR")
		return
	}
	if !fileExists(target) {
		writeLog("不存在: "+target, "ERROR")
		return
	}

	// 删除旧版
	err := os.Remove(target)
	if err != nil {
		writeLog("删除文件错误: "+err.Error(), "ERROR")
		return
	}

	// 将新版重命名
	err = os.Rename(source, target)
	if err != nil {
		writeLog("重命名错误: "+err.Error(), "ERROR")
		return
	}

	var cmdArgs []string
	if args != "" {
		cmdArgs = strings.Split(args, " ")
	}

	err = exec.Command(target, cmdArgs...).Run()
	if err != nil {
		writeLog("重启程序错误: "+err.Error(), "ERROR")
		return
	}
	writeLog("更新成功", "INFO")

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
