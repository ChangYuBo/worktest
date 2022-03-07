package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	// 打开源文件
	srcFile, err := os.Open("./test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	// 创建新文件并打开
	dstFile, err2 := os.OpenFile("./abc2.txt", os.O_CREATE|os.O_APPEND, 0666) //创建  追加
	if err2 != nil {
		fmt.Println(err2)
		return
	}
	// 缓冲读取
	buf := make([]byte, 1024)
	var content []byte

	for {
		n, err := srcFile.Read(buf[:])
		if err == io.EOF {
			// 读取结束
			break
		}
		if err != nil {
			fmt.Println("read file err ", err)
			return
		}
		content = append(content, buf[:n]...)
		v := string(content)
		for _, x := range v {
			switch {
			case x >= 65 && x <= 90:
				str := strings.ToLower(string(x))
				dstFile.WriteString(str)
			case x >= 97 && x <= 121:
				str := strings.ToUpper(string(x))
				dstFile.WriteString(str)
			default:
				dstFile.WriteString(string(x))
			}

		}
		defer srcFile.Close()
		defer dstFile.Close()
	}
	fmt.Println("转换完成")
}
