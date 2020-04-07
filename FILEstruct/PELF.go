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
	case 2:fmt.Println("Sparc")
	case 3:fmt.Println("Intel 80386")
	case 4:fmt.Println("Moto 68K")
	case 5:fmt.Println("Moto 88K")
	case 7:fmt.Println("Intel 80860")
	case 8:fmt.Println("")
	case 10:fmt.Println("")
	case 40:fmt.Println("ARM")
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
