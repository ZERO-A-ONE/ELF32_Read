package FILEstruct

import (
	"../Change"
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