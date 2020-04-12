package FILEstruct

import "fmt"
import "../Change"
//import "../FILEstruct"
func (Shdr Elf32_Shdr) Findshstrtab(Ehdr Elf32_Ehdr,data []byte)[]byte{
	//var num int64 = int64(Ehdr.e_shnum)//获取到名称表总数
	//先找到.shstrtab段的数据拿到节区名称表
	var strtab = int64(Ehdr.e_shstrndx)
	shsize := int64(Ehdr.e_shentsize)
	var t_Index int64= int64(Ehdr.e_shoff)
	TabStIndex := t_Index + strtab*shsize
	t_data := data[int(TabStIndex):int(TabStIndex+shsize)]
	//fmt.Println(t_data)
	//获取节区大小
	t_size := t_data[20:24]
	t_size = Change.LSBtoMSB(t_size)
	size := Change.BytesToInt32(t_size)
	//fmt.Println(int64(size))
	//获取数据偏移地址
	t_offset := t_data[16:20]
	t_offset = Change.LSBtoMSB(t_offset)
	offset := Change.BytesToInt32(t_offset)
	//fmt.Println(Change.DecHex(int64(offset)))
	t_data = data[int(offset):int(offset+size)]
	return t_data
}

func (Shdr Elf32_Shdr) FindName(index int32,data []byte)string{
	var t_str string
	for{
		t_str += string(data[index])
		index++
		if(data[index] == 0){
			break
		}
	}
	return t_str
}

func (Shdr Elf32_Shdr) PHeader(Ehdr Elf32_Ehdr){
	fmt.Println("There are",Ehdr.e_shnum,
		"section headers, starting at offset",
		"0x"+
		Change.DecHex(int64(Ehdr.e_shoff))+":",
		)
	fmt.Println()
	fmt.Println("Section Headers:")
	fmt.Println("[Nr] Name              Type            Addr     Off    Size   ES Flg Lk Inf Al")
}

func (Shdr Elf32_Shdr) Mainread(Ehdr Elf32_Ehdr,data []byte){
	var Shnum int64 = int64(Ehdr.e_shnum) //获取Section Header Table数量
	var StartIndex int64= int64(Ehdr.e_shoff)
	for i:=0;i<int(Shnum);i++ {
		Shdr.SonRead(i,StartIndex,data,int64(Ehdr.e_shentsize),Ehdr)
		fmt.Println()
		StartIndex = StartIndex + int64(Ehdr.e_shentsize)
	}
	fmt.Println("Key to Flags:")
	fmt.Println("  W (write), A (alloc), X (execute), M (merge), S (strings)")
	fmt.Println("  I (info), L (link order), G (group), T (TLS), E (exclude), x (unknown)")
	fmt.Println("  O (extra OS processing required) o (OS specific), p (processor specific)")
}
func (Shdr Elf32_Shdr) SonRead(i int,StartIndex int64,data []byte,size int64,Ehdr Elf32_Ehdr){
	var t_data []byte = data[StartIndex:StartIndex+size]
	var t_shdr Elf32_Shdr
	ShType := map[int]string{
		0x0:"NULL",/* Inactive section header */
		0x1:"PROGBITS",/* Information defined by the program */
		0x2:"SYMTAB",/* Symbol table - not DLL */
		0x3:"STRTAB",/* String table */
		0x4:"RELA",/* Explicit addend relocations, Elf64_Rela */
		0x5:"HASH",/* Symbol hash table */
		0x6:"DYNAMIC",/* Information for dynamic linking */
		0x7:"NOTE",/* A Note section */
		0x8:"NOBITS",/* Like SHT_PROGBITS with no data */
		0x9:"REL",/* Implicit addend relocations, Elf64_Rel */
		0xA:"SHLIB",/* Currently unspecified semantics */
		0xD:"DYNSYM",/* Symbol table for a DLL */
		0xE:"INIT_ARRAY",/* Array of constructors */
		0xF:"FINI_ARRAY",/* Array of deconstructors */
		0x10:"PREINIT_ARRAY",/* Array of pre-constructors */
		0x11:"GROUP",/* Section group */
		0x12:"SYMTAB_SHNDX",/* Extended section indeces */
		0x13:"NUM",/* Number of defined types */
		0x60000000:"LOOS",/* Lowest OS-specific section type */
		0x6ffffff5:"GNU_ATTRIBUTES",/* Object attribuytes */
		0x6ffffff6:"GNU_HASH",/* GNU-style hash table */
		0x6ffffff7:"GNU_LIBLIST",/* Prelink library list */
		0x6ffffff8:"CHECKSUM",/* Checksum for DSO content */
		0x6ffffffa:"LOSUNW",/* Sun-specific low bound */
		0x6ffffffb:"SUNW_COMDAT",
		0x6ffffffc:"SUNW_syminfo",
		0x6ffffffd:"GNU_verdef",/* Version definition section */
		0x6ffffffe:"GNU_verdneed",/* Version needs section */
		0x6fffffff:"GNY_versym",/* Version symbol table */
		0x70000000:"LOPROC",/* Start of processor-specific section type */
		0x7fffffff:"HIPROC",/* End of processor-specific section type */
		0x80000000:"LOUSER",/* Start of application-specific */
		0x8fffffff:"HIUSER",/* Ennd of application-specific */
	}
	ShFlag := map[int]string{
		0x0:"",
		0x1:"W",
		0x2:"A",
		0x3:"WA",
		0x4:"",
		0x5:"",
		0x6:"AX",
		0x30:"MS",
		0x42:"AI",
	}

	//fmt.Println(t_data)
	//sh_name Elf32_Word
	tmp := t_data[:4]
	tmp = Change.LSBtoMSB(tmp)
	t_shdr.sh_name = Elf32_Word(Change.BytesToInt32(tmp))
	//sh_type Elf32_Word
	tmp = t_data[4:8]
	tmp = Change.LSBtoMSB(tmp)
	t_shdr.sh_type = Elf32_Word(Change.BytesToInt32(tmp))
	//sh_flags Elf32_Word
	tmp = t_data[8:12]
	tmp = Change.LSBtoMSB(tmp)
	t_shdr.sh_flags = Elf32_Word(Change.BytesToInt32(tmp))
	//sh_addr Elf32_Addr
	tmp = t_data[12:16]
	tmp = Change.LSBtoMSB(tmp)
	t_shdr.sh_addr = Elf32_Addr(Change.BytesToInt32(tmp))
	//sh_offset Elf32_Off
	tmp = t_data[16:20]
	tmp = Change.LSBtoMSB(tmp)
	t_shdr.sh_offset = Elf32_Off(Change.BytesToInt32(tmp))
	//sh_size Elf32_Word
	tmp = t_data[20:24]
	tmp = Change.LSBtoMSB(tmp)
	t_shdr.sh_size = Elf32_Word(Change.BytesToInt32(tmp))
	//sh_link Elf32_Word
	tmp = t_data[24:28]
	tmp = Change.LSBtoMSB(tmp)
	t_shdr.sh_link = Elf32_Word(Change.BytesToInt32(tmp))
	//sh_info Elf32_Word
	tmp = t_data[28:32]
	tmp = Change.LSBtoMSB(tmp)
	t_shdr.sh_info = Elf32_Word(Change.BytesToInt32(tmp))
	//sh_addralign Elf32_Word
	tmp = t_data[32:36]
	tmp = Change.LSBtoMSB(tmp)
	t_shdr.sh_addralign = Elf32_Word(Change.BytesToInt32(tmp))
	//sh_entsize Elf32_Word
	tmp = t_data[36:40]
	tmp = Change.LSBtoMSB(tmp)
	t_shdr.sh_entsize = Elf32_Word(Change.BytesToInt32(tmp))
	//fmt.Println(t_shdr)
	//[Nr]
	fmt.Print("[")
	fmt.Printf("%2v",i)
	fmt.Print("] ")
	//Name
	if t_shdr.sh_name == 0{
		t_Str := "0"
		fmt.Printf("%-18s",t_Str)
	}else {
		ShName := Shdr.Findshstrtab(Ehdr,data)
		t_Str := Shdr.FindName(int32(t_shdr.sh_name),ShName)
		fmt.Printf("%-18s",t_Str)
	}
	//Type
	tmpstr := ShType[int(int32(t_shdr.sh_type))]
	if tmpstr ==""{
		fmt.Printf("%-16s","Unknown")
	}else{
		fmt.Printf("%-16s",tmpstr)
	}
	//Addr
	tmpstr = Change.DecHex(int64(t_shdr.sh_addr))
	fmt.Printf("%-9s",tmpstr)
	//Off
	tmpstr = Change.DecHex(int64(t_shdr.sh_offset))
	fmt.Printf("%-7s",tmpstr)
	//Size
	tmpstr = Change.DecHex(int64(t_shdr.sh_size))
	fmt.Printf("%-7s",tmpstr)
	//ES
	tmpstr = Change.DecHex(int64(t_shdr.sh_entsize))
	fmt.Printf("%-3s",tmpstr)
	//Flg
	tmpstr = ShFlag[int(int32(t_shdr.sh_flags))]
	fmt.Printf("%3s",tmpstr)
	//Lk
	tmpstr = Change.DecHex(int64(t_shdr.sh_link))
	fmt.Printf("%3s",tmpstr)
	//Inf
	tmpstr = Change.DecHex(int64(t_shdr.sh_info))
	fmt.Print(" ")
	fmt.Printf("%3s",tmpstr)
	//Al
	tmpstr = Change.DecHex(int64(t_shdr.sh_addralign))
	fmt.Print(" ")
	fmt.Printf("%2s",tmpstr)
}