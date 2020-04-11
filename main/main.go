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
	var Phdr FILEstruct.Elf32_phdr
	Phdr.PHeader(Ehdr)
	Phdr.MainRead(Ehdr,data)
	var Shdr FILEstruct.Elf32_Shdr
	//ShName := Shdr.Findshstrtab(Ehdr,data)
	//t_Str := Shdr.FindName(0x1,ShName)
	//fmt.Println(t_Str)
	Shdr.PHeader(Ehdr)
	Shdr.Mainread(Ehdr,data)
	//Shdr.Findshstrtab(Ehdr,data)
}

