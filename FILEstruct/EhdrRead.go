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
	switch t_phdr.p_type {
	case 0:fmt.Printf("%-13s","NULL")
	case 1:fmt.Printf("%-13s","LOAD")
	case 2:fmt.Printf("%-13s","DYNAMIC")
	case 3:fmt.Printf("%-13s","INTERP")
	case 4:fmt.Printf("%-13s","NOTE")
	case 5:fmt.Printf("%-13s","SHLIB")
	case 6:fmt.Printf("%-13s","PHDR")
	case 7:fmt.Printf("%-13s","TLS")
	case 8:fmt.Printf("%-13s","NUM")
	case 0x60000000:fmt.Printf("%-13s","LOOS")
	case 0x6474e550:fmt.Printf("%-13s","GNU_EH_FRAME")
	case 0x6474e551:fmt.Printf("%-13s","GNU_STACK")
	case 0x6474e552:fmt.Printf("%-13s","GNU_RELRO")
	case 0x6ffffffa:fmt.Printf("%-13s","LOSUNW")
	case 0x6ffffffb:fmt.Printf("%-13s","SUNWSTACK")
	case 0x6fffffff:fmt.Printf("%-13s","HISUNW")
	case 0x70000000:fmt.Printf("%-13s","LOPROC")
	case 0x7fffffff:fmt.Printf("%-13s","HIPROC")
	//ARM Sections
	case 0x70000001:fmt.Printf("%-13s","HIPROC")
	case 0x70000002:fmt.Printf("%-13s","PREEMPTMAP")
	case 0x70000003:fmt.Printf("%-13s","ATTRIBUTES")
	case 0x70000004:fmt.Printf("%-13s","DEBUGOVERLAY")
	case 0x70000005:fmt.Printf("%-13s","OVERLAYSECTION")
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
	switch t_phdr.p_flage {
	case 0:fmt.Printf("%-4s","N")
	case 1:fmt.Printf("%-4s","__E")
	case 2:fmt.Printf("%-4s","_W_")
	case 3:fmt.Printf("%-4s","_WE")
	case 4:fmt.Printf("%-4s","R__")
	case 5:fmt.Printf("%-4s","R_E")
	case 6:fmt.Printf("%-4s","RW_")
	case 7:fmt.Printf("%-4s","RWE")
	}
	//p_align
	tstr = "0x"+Change.DecHex(int64(t_phdr.p_align))
	fmt.Printf("%-8s\n",tstr)
	//fmt.Println(t_phdr)
}