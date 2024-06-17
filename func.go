package common_tools

import (
	"log"
	"net"
	"os"
	"strings"
	"time"
)

// 获取当前时间
func GetNowTimeStr() string {
	return time.Unix(time.Now().Unix(), 0).Format("2006-01-02 15:04:05")
}

// 获取本机ip
func GetLocalIp() string {
	conn, err := net.Dial("udp", "8.8.8.8:53")
	if err != nil {
		log.Printf("get local addr err:%v", err)
		return ""
	}
	localIp := strings.Split(conn.LocalAddr().String(), ":")[0]
	return localIp
}

// 获取主机名
func GetHostNmae() string {
	name, _ := os.Hostname()
	return name
}
