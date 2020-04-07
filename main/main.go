package main

import "../FILEstruct"
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
}

