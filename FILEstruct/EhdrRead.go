package FILEstruct

import (
	"../Change"
	"fmt"
)
func (Ehdr Elf32_Ehdr) ReadMagic(data []byte)Elf32_Ehdr{
	var MagicNum = data[:16]
	//Ehdr.e_ident = u_char(MagicNum)
	for _,value := range MagicNum{
		Ehdr.Ehdr_Magic += Change.DecHex(int64(value))
		Ehdr.Ehdr_Magic += " "
	}
	//Ehdr_Class
	Ehdr.Ehdr_Class = int(MagicNum[4])
	//Ehdr_Data
	Ehdr.Ehdr_Data = int(MagicNum[5])
	//Ehdr_Version
	Ehdr.Ehdr_Version = int(MagicNum[6])
	//Ehdr_OS
	Ehdr.Ehdr_OS = int(MagicNum[7])
	//Ehdr_ABIV
	Ehdr.Ehdr_ABIV = int(MagicNum[8])
	return Ehdr
}
func (Ehdr Elf32_Ehdr) ReadHeader(data []byte) Elf32_Ehdr{
	Ehdr = Ehdr.ReadMagic(data)
	//e_type
	tmp := data[16:18]
	tmp = Change.LSBtoMSB(tmp)
	Ehdr.e_type = Elf32_Half(Change.BytesToInt16(tmp))
	//e_machine
	tmp = data[18:20]
	tmp = Change.LSBtoMSB(tmp)
	Ehdr.e_machine = Elf32_Half(Change.BytesToInt16(tmp))
	//e_version
	tmp = data[20:24]
	tmp = Change.LSBtoMSB(tmp)
	Ehdr.e_version = Elf32_Word(1)
	//e_entry
	tmp = data[24:28]
	tmp = Change.LSBtoMSB(tmp)
	Ehdr.e_entry = Elf32_Addr(Change.BytesToInt32(tmp))
	//e_phoff
	tmp = data[28:32]
	tmp = Change.LSBtoMSB(tmp)
	Ehdr.e_phoff = Elf32_Off(Change.BytesToInt32(tmp))
	//e_shoff
	tmp = data[32:36]
	tmp = Change.LSBtoMSB(tmp)
	Ehdr.e_shoff = Elf32_Off(Change.BytesToInt32(tmp))
	//e_flags
	tmp = data[36:40]
	tmp = Change.LSBtoMSB(tmp)
	Ehdr.e_flags = Elf32_Word(Change.BytesToInt32(tmp))
	//e_ehsize
	tmp = data[40:42]
	tmp = Change.LSBtoMSB(tmp)
	Ehdr.e_ehsize = Elf32_Half(Change.BytesToInt16(tmp))
	//e_phentsize
	tmp = data[42:44]
	tmp = Change.LSBtoMSB(tmp)
	Ehdr.e_phentsize = Elf32_Half(Change.BytesToInt16(tmp))
	//e_phnum
	tmp = data[44:46]
	tmp = Change.LSBtoMSB(tmp)
	Ehdr.e_phnum = Elf32_Half(Change.BytesToInt16(tmp))
	//e_shentsize
	tmp = data[46:48]
	tmp = Change.LSBtoMSB(tmp)
	Ehdr.e_shentsize = Elf32_Half(Change.BytesToInt16(tmp))
	//e_shnum
	tmp = data[48:50]
	tmp = Change.LSBtoMSB(tmp)
	Ehdr.e_shnum = Elf32_Half(Change.BytesToInt16(tmp))
	//e_shstrndx
	tmp = data[50:52]
	tmp = Change.LSBtoMSB(tmp)
	Ehdr.e_shstrndx = Elf32_Half(Change.BytesToInt16(tmp))
	return Ehdr
}

func (phdr Elf32_phdr) MainRead(Ehdr Elf32_Ehdr,data []byte){
	var Phnum int64 = int64(Ehdr.e_phnum) //获取Program Header数量
	var StartIndex int64= int64(Ehdr.e_phoff)
	for i:=0;i<int(Phnum);i++ {
		phdr.SonRead(StartIndex,data,int64(Ehdr.e_phentsize))
		StartIndex = StartIndex + int64(Ehdr.e_phentsize)
	}
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