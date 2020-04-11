package FILEstruct
import (
	"../Change"
	"fmt"
)
func (phdr Elf32_phdr) MainRead(Ehdr Elf32_Ehdr,data []byte){
	var Phnum int64 = int64(Ehdr.e_phnum) //获取Program Header数量
	var StartIndex int64= int64(Ehdr.e_phoff)
	for i:=0;i<int(Phnum);i++ {
		phdr.SonRead(StartIndex,data,int64(Ehdr.e_phentsize))
		StartIndex = StartIndex + int64(Ehdr.e_phentsize)
	}
	fmt.Println("")
}
func (phdr Elf32_phdr) SonRead(StartIndex int64,data []byte,size int64){
	PType := map[int]string{
		0x0:"NULL",
		0x1:"LOAD",
		0x2:"DYNAMIC",
		0x3:"INERP",
		0x4:"NOTE",
		0x5:"SHLIB",
		0x6:"PHDR",
		0x7:"TLS",
		0x8:"NUM",
		0x60000000:"LOOS",
		0x6474e550:"GNU_EH_FRAME",
		0x6474e551:"GNU_STACK",
		0x6474e552:"GNU_RELRO",
		0x6ffffffa:"LOSUNW",
		0x6ffffffc:"SUNWBSS",
		0x6ffffffb:"SUNWSTACK",
		0x6fffffff:"HISUNW ",
		0x6ffffffe:"HIOS",
		0x70000000:"LOPROC",
		0x7fffffff:"HIPROC",
		// ARM Sections
		0x70000001:"ARM_EXIDX",
		0x70000002:"ARM_PREEMPTMAP",
		0x70000003:"ARM_ATTRIBUTES",
		0x70000004:"ARM_DEBUGOVERLAY",
		0x70000005:"ARM_OVERLAYSECTION",
	}
	PFlag := map[int]string{
		0:"N",
		1:"__E",
		2:"_W_",
		3:"_WE",
		4:"R__",
		5:"R_E",
		6:"RW_",
		7:"RWE",
	}

	var t_phdr Elf32_phdr
	var t_data []byte = data[StartIndex:StartIndex+size]
	//fmt.Println(t_data)
	//p_type Elf32_Word
	tmp := t_data[:4]
	tmp = Change.LSBtoMSB(tmp)
	t_phdr.p_type = Elf32_Word(Change.BytesToInt32(tmp))
	//p_offset Elf32_Off
	tmp = t_data[4:8]
	tmp = Change.LSBtoMSB(tmp)
	t_phdr.p_offset = Elf32_Off(Change.BytesToInt32(tmp))
	//p_vaddr Elf32_Addr
	tmp = t_data[8:12]
	tmp = Change.LSBtoMSB(tmp)
	t_phdr.p_vaddr = Elf32_Addr(Change.BytesToInt32(tmp))
	//p_paddr Elf32_Addr
	tmp = t_data[12:16]
	tmp = Change.LSBtoMSB(tmp)
	t_phdr.p_paddr = Elf32_Addr(Change.BytesToInt32(tmp))
	//p_filesz Elf32_Word
	tmp = t_data[16:20]
	tmp = Change.LSBtoMSB(tmp)
	t_phdr.p_filesz = Elf32_Word(Change.BytesToInt32(tmp))
	//p_memsz Elf32_Word
	tmp = t_data[20:24]
	tmp = Change.LSBtoMSB(tmp)
	t_phdr.p_memsz = Elf32_Word(Change.BytesToInt32(tmp))
	//p_flage Elf32_Word
	tmp = t_data[24:28]
	tmp = Change.LSBtoMSB(tmp)
	t_phdr.p_flage = Elf32_Word(Change.BytesToInt32(tmp))
	//p_align Elf32_Word
	tmp = t_data[28:32]
	tmp = Change.LSBtoMSB(tmp)
	t_phdr.p_align = Elf32_Word(Change.BytesToInt32(tmp))
	//开始打印
	//p_type
	tmpstr := PType[int(int32(t_phdr.p_type))]
	if tmpstr ==""{
		fmt.Printf("%-13s","Unknown")
	}else{
		fmt.Printf("%-13s",tmpstr)
	}
	//p_offset
	fmt.Print("  ")
	tstr := "0x"+Change.DecHex(int64(t_phdr.p_offset))
	fmt.Printf("%-9s",tstr)
	//p_vaddr
	tstr = "0x"+Change.DecHex(int64(t_phdr.p_vaddr))
	fmt.Printf("%-11s",tstr)
	//p_paddr
	tstr = "0x"+Change.DecHex(int64(t_phdr.p_paddr))
	fmt.Printf("%-11s",tstr)
	//p_filesz
	tstr = "0x"+Change.DecHex(int64(t_phdr.p_filesz))
	fmt.Printf("%-8s",tstr)
	//p_memsz
	tstr = "0x"+Change.DecHex(int64(t_phdr.p_memsz))
	fmt.Printf("%-8s",tstr)
	//p_flage
	//R:Read W:Write E:Exec N:None
	tmpstr = PFlag[int(int32(t_phdr.p_flage))]
	if tmpstr ==""{
		fmt.Printf("%-13s","Unknown")
	}else{
		fmt.Printf("%-4s",tmpstr)
	}
	//p_align
	tstr = "0x"+Change.DecHex(int64(t_phdr.p_align))
	fmt.Printf("%-8s\n",tstr)
	//fmt.Println(t_phdr)
}

func (phdr Elf32_phdr) PHeader(Ehdr Elf32_Ehdr){
	EhdrType := map[int]string{
		0x0 : "NONE",//No file type
		0x1 : "REL",//Relocatable
		0x2 : "EXEC",//Executable
		0x3 : "DYN",//Shared object
		0x4 : "CORE",//Core
		0xfe00 : "LOOS",//OS specific
		0xfeff : "HIOS",//OS specific
		0xff00 : "LOPROC",//Processor specific
		0xffff : "HIPROC",//Processor specific
	}
	var T_type string
	tmpstr := EhdrType[int(int32(Ehdr.e_type))]
	if tmpstr == ""{
		T_type = "Unknown"
	}else {
		T_type =tmpstr
	}
	fmt.Println("Elf file type is "+T_type)
	fmt.Println("Entry point 0x"+Change.DecHex(int64(Ehdr.e_entry)))
	fmt.Print("There are ",Ehdr.e_phnum," program headers, ")
	fmt.Println("starting at offset",Ehdr.e_phoff)
	fmt.Println("Program Headers:")
	fmt.Println("Type           Offset   VirtAddr   PhysAddr   FileSiz MemSiz  Flg Align")

}