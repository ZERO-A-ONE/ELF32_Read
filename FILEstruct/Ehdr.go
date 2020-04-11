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

func (Ehdr Elf32_Ehdr) PHeader(){

	EhdrClass := map[int]string{
		1: "ELF32",
		2: "ELF64",
	}
	EhdrData := map[int]string{
		1:"2's complement, little endian",
		2:"2's complement, big endian",
	}
	EhdrVersion := map[int]string{
		1:"1 (current)",
	}
	EhdrOSABI := map[int]string{
		0x0 : "UNIX - System V",
		0x1 : "Hewlett-Packard HP-UX",
		0x2 : "NetBSD",
		0x3 : "Linux",
		0x6 : "Sun Solaris",
		0x7 : "AIX",
		0x8 : "IRIX",
		0x9 : "FreeBSD",
		0xA : "Compaq TRU64 UNIX",
		0xB : "Novell Modesto",
		0xC : "Open BSD",
		0xD : "Open VMS",
		0xE : "Hewlett-Packard Non-Stop Kernel",
		0xF : "Amiga Research OS",
		0x40: "ARM EABI",
		0x61: "ARM",
		0xFF: "Standalone (embedded applications)",
	}
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
	EhdrMachine := map[int]string{
		0 : "No machine",
		1 : "AT&T WE 32100",
		2 : "SPARC",
		3 : "Intel 80386",
		4 : "Motorola 68000",
		5 : "Motorola 88000",
		6 : "Reserved for future use (was EM_486)",
		7 : "Intel 80860",
		8 : "MIPS I Architecture",
		9 : "IBM System/370 Processor",
		10 : "MIPS RS3000 Little-endian",
		11 : "Reserved for future use",
		12 : "Reserved for future use",
		13 : "Reserved for future use",
		14 : "Reserved for future use",
		15 : "Hewlett-Packard PA-RISC",
		16 : "Reserved for future use",
		17 : "Fujitsu VPP500",
		18 : "Enhanced instruction set SPARC",
		19 : "Intel 80960",
		20 : "PowerPC",
		21 : "64-bit PowerPC",
		22 : "IBM System/390 Processor",
		23 : "Reserved for future use",
		24 : "Reserved for future use",
		25 : "Reserved for future use",
		26 : "Reserved for future use",
		27 : "Reserved for future use",
		28 : "Reserved for future use",
		29 : "Reserved for future use",
		30 : "Reserved for future use",
		31 : "Reserved for future use",
		32 : "Reserved for future use",
		33 : "Reserved for future use",
		34 : "Reserved for future use",
		35 : "Reserved for future use",
		36 : "NEC V800",
		37 : "Fujitsu FR20",
		38 : "TRW RH-32",
		39 : "Motorola RCE",
		40 : "Advanced RISC Machines ARM",
		41 : "Digital Alpha",
		42 : "Hitachi SH",
		43 : "SPARC Version 9",
		44 : "Siemens TriCore embedded processor",
		45 : "Argonaut RISC Core, Argonaut Technologies Inc.",
		46 : "Hitachi H8/300",
		47 : "Hitachi H8/300H",
		48 : "Hitachi H8S",
		49 : "Hitachi H8/500",
		50 : "Intel IA-64 processor architecture",
		51 : "Stanford MIPS-X",
		52 : "Motorola ColdFire",
		53 : "Motorola M68HC12",
		54 : "Fujitsu MMA Multimedia Accelerator",
		55 : "Siemens PCP",
		56 : "Sony nCPU embedded RISC processor",
		57 : "Denso NDR1 microprocessor",
		58 : "Motorola Star*Core processor",
		59 : "Toyota ME16 processor",
		60 : "STMicroelectronics ST100 processor",
		61 : "Advanced Logic Corp. TinyJ embedded processor family",
		62 : "AMD x86-64 architecture",
		63 : "Sony DSP Processor",
		64 : "Digital Equipment Corp. PDP-10",
		65 : "Digital Equipment Corp. PDP-11",
		66 : "Siemens FX66 microcontroller",
		67 : "STMicroelectronics ST9+ 8/16 bit microcontroller",
		68 : "STMicroelectronics ST7 8-bit microcontroller",
		69 : "Motorola MC68HC16 Microcontroller",
		70 : "Motorola MC68HC11 Microcontroller",
		71 : "Motorola MC68HC08 Microcontroller",
		72 : "Motorola MC68HC05 Microcontroller",
		73 : "Silicon Graphics SVx",
		75 : "Digital VAX",
		76 : "Axis Communications 32-bit embedded processor",
		77 : "Infineon Technologies 32-bit embedded processor",
		78 : "Element 14 64-bit DSP Processor",
		79 : "LSI Logic 16-bit DSP Processor",
		80 : "Donald Knuth's educational 64-bit processor",
		81 : "Harvard University machine-independent object files",
		82 : "SiTera Prism",
		83 : "Atmel AVR 8-bit microcontroller",
		84 : "Fujitsu FR30",
		85 : "Mitsubishi D10V",
		86 : "Mitsubishi D30V",
		87 : "NEC v850",
		88 : "Mitsubishi M32R",
		89 : "Matsushita MN10300",
		90 : "Matsushita MN10200",
		91 : "picoJava",
		92 : "OpenRISC 32-bit embedded processor",
		93 : "ARC Cores Tangent-A5",
		94 : "Tensilica Xtensa Architecture",
		95 : "Alphamosaic VideoCore processor",
		96 : "Thompson Multimedia General Purpose Processor",
		97 : "National Semiconductor 32000 series",
		98 : "Tenor Network TPC processor",
		99 : "Trebia SNP 1000 processor",
		100 : "STMicroelectronics (www.st.com) ST200 microcontroller",
		101 : "Ubicom IP2xxx microcontroller family",
		102 : "MAX Processor",
		103 : "National Semiconductor CompactRISC microprocessor",
		104 : "Fujitsu F2MC16",
		105 : "Texas Instruments embedded microcontroller msp430",
		106 : "Analog Devices Blackfin (DSP) processor",
		107 : "S1C33 Family of Seiko Epson processors",
		108 : "Sharp embedded microprocessor",
		109 : "Arca RISC Microprocessor",
		110 : "Microprocessor series from PKU-Unity Ltd. and MPRC of Peking University",
	}


	fmt.Println("ELF Header:")
	fmt.Printf("%-9s","Magic: ")
	fmt.Println(Ehdr.Ehdr_Magic)
	//Class
	fmt.Printf("%-35s","Class: ")
	tmpstr := EhdrClass[Ehdr.Ehdr_Class]
	if tmpstr == ""{
		fmt.Println("Invalid class")
	}else{
		fmt.Println(tmpstr)
	}
	//Data
	fmt.Printf("%-35s","Data: ")
	tmpstr = EhdrData[Ehdr.Ehdr_Data]
	if tmpstr == ""{
		fmt.Println("Invaild data encoding")
	}else {
		fmt.Println(tmpstr)
	}
	//Version
	fmt.Printf("%-35s","Version: ")
	tmpstr = EhdrVersion[Ehdr.Ehdr_Version]
	if tmpstr == ""{
		fmt.Println("Invaild version")
	} else{
		fmt.Println(tmpstr)
	}
	//OS/ABI
	fmt.Printf("%-35s","OS/ABI: ")
	tmpstr = EhdrOSABI[Ehdr.Ehdr_OS]
	if tmpstr == ""{
		fmt.Println("Unknown")
	} else{
		fmt.Println(tmpstr)
	}
	//ABI Version
	fmt.Printf("%-35s","ABI Version:")
	fmt.Println(int64(Ehdr.Ehdr_ABIV))
	//Type
	fmt.Printf("%-35s","Type: ")
	tmpstr = EhdrType[int(int32(Ehdr.e_type))]
	if tmpstr == ""{
		fmt.Println("Unknown" )
	}else {
		fmt.Println(tmpstr )
	}
	//Machine
	fmt.Printf("%-35s","Machine: ")
	tmpstr = EhdrMachine[int(int32(Ehdr.e_machine))]
	if tmpstr == ""{
		fmt.Println("Unknown" )
	}else {
		fmt.Println(tmpstr)
	}
	//Version
	fmt.Printf("%-35s","Version:")
	fmt.Println("0x"+Change.DecHex(int64(Ehdr.e_version)))
	//Entry point address
	fmt.Printf("%-35s","Entry point address:")
	fmt.Println("0x"+Change.DecHex(int64(Ehdr.e_entry)))
	//Start of program headers
	fmt.Printf("%-35s","Start of program headers:")
	fmt.Println(int64(Ehdr.e_phoff),"(bytes into file)")
	//Start of section headers
	fmt.Printf("%-35s","Start of section headers:")
	fmt.Println(int64(Ehdr.e_shoff),"(bytes into file)")
	//Flags
	fmt.Printf("%-35s","FLags:")
	fmt.Println("0x"+Change.DecHex(int64(Ehdr.e_flags)))
	//Size of this header
	fmt.Printf("%-35s","Size of this header:")
	fmt.Println(int64(Ehdr.e_ehsize),"(bytes)")
	//Size of program headers
	fmt.Printf("%-35s","Size of program headers:")
	fmt.Println(int64(Ehdr.e_phentsize),"(bytes)")
	//Number of program headers
	fmt.Printf("%-35s","Number of program headers:")
	fmt.Println(int64(Ehdr.e_phnum))
	//Size of section headers
	fmt.Printf("%-35s","Size of section headers:")
	fmt.Println(int64(Ehdr.e_shentsize),"(bytes)")
	//Number of section headers
	fmt.Printf("%-35s","Number of section headers:")
	fmt.Println(int64(Ehdr.e_shnum))
	//Section header string table index
	fmt.Printf("%-35s","Section header string table index:")
	fmt.Println(int64(Ehdr.e_shstrndx))
}