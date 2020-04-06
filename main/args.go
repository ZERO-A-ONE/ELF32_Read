package main

import (
	"flag"
)

var File_name = flag.String("f","","文件名称")
var All = flag.Bool("a",false,"显示全部信息,等价于 -h -l -S -s -r -d -V -A -I")
var File_header = flag.Bool("h",false,"显示elf文件开始的文件头信息")
var Program_headers = flag.Bool("l",false,"显示程序头（段头）信息(如果有的话)")
var Section_headers = flag.Bool("S",false,"显示节头信息(如果有的话)")
var Section_groups = flag.Bool("g",false,"显示节组信息(如果有的话)")
var Section_details = flag.Bool("t",false,"显示节的详细信息(-S的)")
var Symbols = flag.Bool("s",false,"显示符号表段中的项（如果有的话）")
var Headers = flag.Bool("e",false,"显示全部头信息，等价于: -h -l -S")
var Notes = flag.Bool("n",false,"显示note段（内核注释）的信息")
var Relocs = flag.Bool("r",false,"显示可重定位段的信息")
var Dynamic = flag.Bool("d",false,"显示动态段的信息")
var Arch_specific = flag.Bool("A",false,"显示CPU构架信息")

func ReadArgs()  {
	flag.Parse()
	/*
	if *File_name == ""{
		fmt.Println("请指定文件名 -n name")
	}*/
}