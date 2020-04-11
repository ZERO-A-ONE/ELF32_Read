package FILEstruct
const EI_NIDENT = 16
type Elf32_Addr = uint32
type Elf32_Half = uint16
type Elf32_Off = uint32
type Elf32_SWord = int32
type Elf32_Word = uint32
type u_char = byte
//ELF_Header
type Elf32_Ehdr struct {
	//e_ident [EI_NIDENT]u_char
	Ehdr_Magic string
	Ehdr_Class int
	Ehdr_Data int
	Ehdr_Version int
	Ehdr_OS int
	Ehdr_ABIV int
	/*
	e_ident 这是一个数组,其每个字节又都有所代表的含义
	EI_MAG0 - EI_MAG3 文件标识就是平时所说的ELF头
	EI_CLASS 文件类,其实代表的是32位/64位程序
	EI_DATA 数据编码,一般都是01[td]
	EI_VERSION 文件版本,固定值01 EV_CURRENT
	EI_PAD 呃…就是一堆全是00的用来补全大小的数组
	EI_NIDENT 说是e_ident数组的大小,但我看了好几个so都是00
	*/
	e_type Elf32_Half
	//标识文件类型
	e_machine Elf32_Half
	//声明ABI
	e_version Elf32_Word
	//跟ident[]里的EI_VERSION一样,为01
	e_entry Elf32_Addr
	//可执行程序入口点地址
	e_phoff Elf32_Off
	//Program Header Offset,程序头部表索引地址,没有则为0
	e_shoff Elf32_Off
	//Section Header Offset,节区表索引地址,没有则为0
	e_flags Elf32_Word
	//保存与文件相关的，特定于处理器的标志。
	e_ehsize Elf32_Half
	//ELF_Header Size,ELF头部的大小
	e_phentsize Elf32_Half
	//程序头部表的单个表项的大小
	e_phnum Elf32_Half
	//程序头部表的表项数
	e_shentsize Elf32_Half
	//节区表的单个表项的大小
	e_shnum Elf32_Half
	//节区表的表项数
	e_shstrndx Elf32_Half
	//String Table Index,在节区表中有一个存储各节区名称的节区,这里表示名称表在第几个节区。
}
//Program Header
type Elf32_phdr struct {
	p_type Elf32_Word
	//此数组元素描述的段的类型，或者如何解释此数组元素的信息
	p_offset Elf32_Off
	//此成员给出从文件头到该段第一个字节的偏移
	p_vaddr Elf32_Addr
	//此成员给出段的第一个字节将被放到内存中的虚拟地址
	p_paddr Elf32_Addr
	//此成员仅用于与物理地址相关的系统中。System V忽略所有应用程序的物理地址信息
	p_filesz Elf32_Word
	//此成员给出段在文件映像中所占的字节数。可以为0
	p_memsz Elf32_Word
	//此成员给出段在内存映像中占用的字节数。可以为0
	p_flage Elf32_Word
	//此成员给出与段相关的标志(read、write、exec)
	p_align Elf32_Word
	//此成员给出段在文件中和内存中如何对齐
	//字节对其,p_vaddr 和 p_offset 对 p_align 取模后应该等于0
}
//Section Header Table
type Elf32_Shdr struct {
	sh_name Elf32_Word
	//节区名称,此处是一个在名称节区的索引
	sh_type Elf32_Word
	//节区类型
	sh_flags Elf32_Word
	//同Program Header的p_flags
	sh_addr Elf32_Addr
	//节区索引地址
	sh_offset Elf32_Off
	//节区相对于文件的偏移地址
	sh_size Elf32_Word
	//节区的大小
	sh_link Elf32_Word
	//此成员给出节区头部表索引链接
	sh_info Elf32_Word
	//此成员给出附加信息
	sh_addralign Elf32_Word
	//某些节区带有地址对齐约束
	sh_entsize Elf32_Word
	//某些节区中包含固定大小的项目,如符号表
}