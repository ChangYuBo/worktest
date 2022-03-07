package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("./test.txt") //创建文件
	if err != nil {
		fmt.Println("error")
		return
	}
	defer file.Close()

	for i := 0; i < 5; i++ {
		file.WriteString("hello\n")
		file.Write([]byte("world\n"))
	}

	// file, _ := os.Open("./test.txt")

	// defer file.Close()

	// // 字节切片缓存 存放每次读取的字节
	// buf := make([]byte, 1024)

	// // 该字节切片用于存放文件所有字节
	// var bytes []byte

	// for {
	// 	// 返回本次读取的字节数
	// 	count, err := file.Read(buf)

	// 	// 检测是否到了文件末尾
	// 	if err == io.EOF {
	// 		break
	// 	}

	// 	// 取出本次读取的数据
	// 	currBytes := buf[:count]

	// 	// 将读取到的数据 追加到字节切片中
	// 	bytes = append(bytes, currBytes...)
	// 	for _, i := range bytes {
	// 		fmt.Println(string(i))
	// 	}
	// }

	// 将字节切片转为字符串 最后打印出来文件内容
	//fmt.Println(string(bytes))

	content, err := os.OpenFile("./test.txt", os.O_RDWR|os.O_APPEND, 0666) //打开文件
	if err != nil {
		fmt.Println(err)
		return
	}
	defer content.Close()
	temp := make([]byte, 1024)

	content.Read(temp)

	fmt.Println(string(temp))

}
