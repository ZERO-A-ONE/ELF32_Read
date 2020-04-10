package main

import (
	"fmt"
	"os"
)
func Openfile(Fpath string) []byte{
	//1.打开文件
	fp, err := os.Open(Fpath)
	if err != nil {
		fmt.Println("Open File Error",err)
	}else{
		//fmt.Println("打开文件成功")
	}
	//2.关闭文件
	defer func() {
		if err = fp.Close(); err != nil {
			//fmt.Println("关闭文件失败",err)
		}else{
			//fmt.Println("关闭文件成功")
		}
	}()
	//3.文件信息获取
	finfo, err := os.Stat(Fpath)
	if err != nil {
		fmt.Println(err)
	}else{
		//fmt.Println(finfo)
		fmt.Println("Name: ",finfo.Name())
		fmt.Println("Size: ",finfo.Size())
		fmt.Println("ModTime: ",finfo.ModTime())
		fmt.Println("IsDir: ",finfo.IsDir())
	}
	var b []byte = make([]byte,2*finfo.Size())
	_, err = fp.Read(b)
	if err != nil{
		fmt.Println("Read File Error",err)
	}else{
		//fmt.Println("读取文件成功")
	}
	return b
	//4.读取文件
	/*
	r := bufio.NewReader(fp) //接收一个句柄,返回一个句柄,利用返回的句柄来读取数据
	ByteT := make([]uint8, 1)
	for {
		buff, err := r.ReadBytes(nil)
		//fmt.Println(reflect.TypeOf(buff))
		//fmt.Print(string(buff))
		for _, value := range buff {
			ByteT = append(ByteT, value)
		}
		if err == io.EOF {
			if err != nil {
				fmt.Print("读取失败")
				fmt.Print(err)
				break
			}
		}
	}
	return ByteT*/
}