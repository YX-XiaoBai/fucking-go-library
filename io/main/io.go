package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	// 操作终端
	//var buf [16]byte
	//os.Stdin.Read(buf[:])
	//os.Stdin.WriteString(string(buf[:]))

	// Open()
	//file, err := os.Open("./io.go")
	//if err != nil {
	//	fmt.Println("open file failed! Error:", err)
	//	return
	//}
	//fmt.Println("open file success!")
	//file.Close()

	// OpenFile()
	//file1, err := os.OpenFile("io.log", os.O_RDONLY|os.O_CREATE, 0755)
	//if err != nil {
	//	log.Fatal("Error:", err)
	//}
	//fmt.Println("open file1 success")
	//if err := file1.Close(); err != nil {
	//	log.Fatal("Error:", err)
	//}
	//fmt.Println("close file1 success")

	// WriteXx() & Create()
	//file, err := os.Create("./io.txt")
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//defer file.Close()
	//for i := 0; i < 5; i ++ {
	//	file.WriteString("ab\n")
	//	file.Write([]byte("cd\n"))
	//}
	//file.WriteAt([]byte("fe"), 3)

	// Read() & ReadAt()
	//file, err := os.Open("./io.txt")
	//if err != nil {
	//	fmt.Println("open file err:", err)
	//	return
	//}
	//defer file.Close()
	//var buf [128]byte
	//var content []byte
	//for {
	//	n, err := file.Read(buf[:])
	//	//n, err := file.ReadAt(buf[:], int64(0))
	//	//fmt.Println(string(buf[:n]))
	//	if err == io.EOF {
	//		break
	//	}
	//	if err != nil{
	//		fmt.Println("read file err", err)
	//		return
	//	}
	//	content = append(content, buf[:n]...)
	//}
	//
	//fmt.Println(string(content))

	// 拷贝文件
	srcFile, err := os.Open("./io.txt")
	if err != nil {
		fmt.Println("E1:", err)
		return
	}
	newFile, err2 := os.Create("./io2.txt")
	if err2 != nil {
		fmt.Println("E2", err2)
		return
	}
	buf := make([]byte, 1024)
	for {
		n, err := srcFile.Read(buf)
		if err == io.EOF {
			fmt.Println("Read Done")
			break
		}
		if err != nil {
			fmt.Println(err)
			break
		}
		newFile.Write(buf[:n])
	}
	srcFile.Close()
	newFile.Close()
}

