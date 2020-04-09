package FILEstruct

import "fmt"
import "../Change"
func (Ehdr Elf32_Ehdr) PHeader(){
	fmt.Println("Magic: ",Ehdr.Ehdr_Magic)
	//Class
	fmt.Print("Class: ")
	if Ehdr.Ehdr_Class == 1{
		fmt.Println("ELF32")
	}else if Ehdr.Ehdr_Class == 2{
		fmt.Println("ELF64")
	}else{
		fmt.Println("Invalid class")
	}
	//Data
	fmt.Print("Data: ")
	if Ehdr.Ehdr_Data == 1{
		fmt.Println("2's complement, little endian")
	}else if Ehdr.Ehdr_Data == 2{
		fmt.Println("2's complement, big endian")
	}else {
		fmt.Println("Invaild data encoding")
	}
	//Version
	fmt.Print("Version: ")
	if Ehdr.Ehdr_Version == 1{
		fmt.Println("1 (current)")
	} else{
		fmt.Println("Invaild version")
	}
	//OS/ABI
	fmt.Print("OS/ABI: ")
	if Ehdr.Ehdr_OS == 0{
		fmt.Println("UNIX - System V")
	}else if Ehdr.Ehdr_OS == 1{
		fmt.Println("HPUX")
	}else if Ehdr.Ehdr_OS == 255{
		fmt.Println("Standalone")
	}else{
		fmt.Println("Unknown")
	}
	//ABI Version
	fmt.Println("ABI Version: 0x"+Change.DecHex(int64(Ehdr.Ehdr_ABIV)))
	//Type
	fmt.Print("Type: ")
	switch Ehdr.e_type {
	case 0:fmt.Println("No file type")
	case 1:fmt.Println("Relocatable")
	case 2:fmt.Println("Executable")
	case 3:fmt.Println("Shared object")
	case 4:fmt.Println("Core")
	case 65024:fmt.Println("OS specific")
	case 65279:fmt.Println("OS specific")
	case 65280:fmt.Println("Processor specific")
	case 65535:fmt.Println("Processor specific")
	default:
		fmt.Println("Unknown" )
	}
	//Machine
	fmt.Print("Machine: ")
	switch Ehdr.e_machine {
	case 0:fmt.Println("No machine")
	case 1:fmt.Println("AT&T WE32100")
	case 2:fmt.Println("SPARC")
	case 3:fmt.Println("Intel 80386")
	case 4:fmt.Println("Motorola 68000")
	case 5:fmt.Println("Motorola 88000")
	case 6:fmt.Println("Reserved for future use")
	case 7:fmt.Println("Intel 80860")
	case 8:fmt.Println("MIPS I Architecture")
	case 9:fmt.Println("Reserved for future use")
	case 10:fmt.Println("MIPS RS3000 Little-endian")
	case 11,12,13,14:fmt.Println("Reserved for future use")
	case 15:fmt.Println("Hewlett-Packard PA-RISC")
	case 16:fmt.Println("Reserved for future use")
	case 17:fmt.Println("Fujitsu VPP500")
	case 18:fmt.Println("Enhanced instruction set SPARC")
	case 19:fmt.Println("Intel 80960")
	case 20:fmt.Println("Power PC")
	case 21,22,23,24,25,26,27,28,29,30,31,32,33,34,35:fmt.Println("Reserved for future use")
	case 36:fmt.Println("NEC V800")
	case 37:fmt.Println("Fujitsu FR20")
	case 38:fmt.Println("TRW RH-32")
	case 39:fmt.Println("Motorola RCE")
	case 40:fmt.Println("Advanced RISC Machines ARM")
	case 41:fmt.Println("Digital Alpha")
	case 42:fmt.Println("Hitachi SH")
	case 43:fmt.Println("SPARC Version 9")
	case 44:fmt.Println("Siemens Tricore embedded processor")
	case 45:fmt.Println("Argonaut RISC Core, Argonaut Technologies Inc.")
	case 46:fmt.Println("Hitachi H8/300")
	case 47:fmt.Println("Hitachi H8/300H")
	case 48:fmt.Println("Hitachi H8S")
	case 49:fmt.Println("Hitachi H8/500")
	case 50:fmt.Println("Intel MercedTM Processor")
	case 51:fmt.Println("Stanford MIPS-X")
	case 52:fmt.Println("Motorola Coldfire")
	case 53:fmt.Println("Motorola M68HC12")
	default:
		fmt.Println("Unknown" )
	}
	//Version
	fmt.Println("Version: 0x"+Change.DecHex(int64(Ehdr.e_version)))
	//Entry point address
	fmt.Println("Entry point address: 0x"+Change.DecHex(int64(Ehdr.e_entry)))
	//Start of program headers
	fmt.Println("Start of program headers: 0x"+Change.DecHex(int64(Ehdr.e_phoff)),"(bytes into file)")
	//Start of section headers
	fmt.Println("Start of section headers: 0x"+Change.DecHex(int64(Ehdr.e_shoff)),"(bytes into file)")
	//Flags
	fmt.Println("FLags: 0x"+Change.DecHex(int64(Ehdr.e_flags)))
	//Size of this header
	fmt.Println("Size of this header: 0x"+Change.DecHex(int64(Ehdr.e_ehsize)),"(bytes)")
	//Size of program headers
	fmt.Println("Size of program headers: 0x"+Change.DecHex(int64(Ehdr.e_phentsize)),"(bytes)")
	//Number of program headers
	fmt.Println("Number of program headers: 0x"+Change.DecHex(int64(Ehdr.e_phnum)))
	//Size of section headers
	fmt.Println("Size of section headers: 0x"+Change.DecHex(int64(Ehdr.e_shentsize)),"(bytes)")
	//Number of section headers
	fmt.Println("Number of section headers: 0x"+Change.DecHex(int64(Ehdr.e_shnum)))
	//Section header string table index
	fmt.Println("Section header string table index: 0x"+Change.DecHex(int64(Ehdr.e_shstrndx)))
}
func (phdr Elf32_phdr) PHeader(Ehdr Elf32_Ehdr){
	var T_type string
	switch Ehdr.e_type {
	case 0:T_type = "No file type"
	case 1:T_type = "Relocatable"
	case 2:T_type = "Executable"
	case 3:T_type = "Shared object"
	case 4:T_type = "Core"
	case 65024:T_type = "OS specific"
	case 65279:T_type = "OS specific"
	case 65280:T_type = "Processor specific"
	case 65535:T_type = "Processor specific"
	default:
		T_type = "Unknown"
	}
	fmt.Println("Elf file type is "+T_type)
	fmt.Println("Entry point 0x"+Change.DecHex(int64(Ehdr.e_entry)))
	fmt.Print("There are ",Ehdr.e_phnum," program headers, ")
	fmt.Println("starting at offset",Ehdr.e_phoff)
	fmt.Println("Program Headers:")
	fmt.Println("Type           Offset   VirtAddr   PhysAddr   FileSiz MemSiz  Flg Align")

}
