package main

import "fmt"
import "../ReadFile"

func main() {
	ReadArgs()
	var data []byte
	data = ReadFile.Openfile("D:/GOL/ELFRead/hello")
	fmt.Println(data)
}

