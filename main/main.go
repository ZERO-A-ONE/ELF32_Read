package main

import (
	"../FILEstruct"
	"fmt"
)
func main() {
	ReadArgs()
	var data []byte
	data = Openfile("D:/GOL/ELFRead/hello")
	//fmt.Println(data)
	//fmt.Println(len(data))
	var Ehdr FILEstruct.Elf32_Ehdr
	//fmt.Println(Ehdr)
	//Ehdr = Ehdr.ReadMagic(data)
	//fmt.Println(Ehdr)
	Ehdr = Ehdr.ReadHeader(data)
	//fmt.Println(Ehdr)
	Ehdr.PHeader()
	fmt.Println("")
	var phdr FILEstruct.Elf32_phdr
	phdr.PHeader(Ehdr)
	phdr.MainRead(Ehdr,data)
}

