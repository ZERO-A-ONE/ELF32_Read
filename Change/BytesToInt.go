package Change
import (
	"bytes"
	"encoding/binary"
	"fmt"
)
//isSymbol表示有无符号
func BytesToInt(b []byte, isSymbol bool)  (int, error){
	if isSymbol {
		return bytesToIntS(b)
	}
	return bytesToIntU(b)
}
//字节数(大端)组转成int(无符号的)
func bytesToIntU(b []byte) (int, error) {
	if len(b) == 3 {
		b = append([]byte{0},b...)
	}
	bytesBuffer := bytes.NewBuffer(b)
	switch len(b) {
	case 1:
		var tmp uint8
		err := binary.Read(bytesBuffer, binary.BigEndian, &tmp)
		return int(tmp), err
	case 2:
		var tmp uint16
		err := binary.Read(bytesBuffer, binary.BigEndian, &tmp)
		return int(tmp), err
	case 4:
		var tmp uint32
		err := binary.Read(bytesBuffer, binary.BigEndian, &tmp)
		return int(tmp), err
	default:
		return 0,fmt.Errorf("%s", "BytesToInt bytes lenth is invaild!")
	}
}
//字节数(大端)组转成int(有符号)
func bytesToIntS(b []byte) (int, error) {
	if len(b) == 3 {
		b = append([]byte{0},b...)
	}
	bytesBuffer := bytes.NewBuffer(b)
	switch len(b) {
	case 1:
		var tmp int8
		err := binary.Read(bytesBuffer, binary.BigEndian, &tmp)
		return int(tmp), err
	case 2:
		var tmp int16
		err := binary.Read(bytesBuffer, binary.BigEndian, &tmp)
		return int(tmp), err
	case 4:
		var tmp int32
		err := binary.Read(bytesBuffer, binary.BigEndian, &tmp)
		return int(tmp), err
	default:
		return 0,fmt.Errorf("%s", "BytesToInt bytes lenth is invaild!")
	}
}
//整形转换成字节
func IntToByte(n int) []byte {
	x := int32(n)
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, x)
	return bytesBuffer.Bytes()
}
//字节转换成整形
func ByteToInt(b []byte) int {
	bytesBuffer := bytes.NewBuffer(b)

	var x int32
	binary.Read(bytesBuffer, binary.BigEndian, &x)

	return int(x)
}
func Int16ToBytes(i int16) []byte {
	buf := make([]byte, 8)
	binary.BigEndian.PutUint16(buf, uint16(i))
	return buf
}

func BytesToInt16(buf []byte) int16 {
	return int16(binary.BigEndian.Uint16(buf))
}
func Int32ToBytes(i int32) []byte {
	buf := make([]byte, 16)
	binary.BigEndian.PutUint32(buf, uint32(i))
	return buf
}

func BytesToInt32(buf []byte) int32 {
	return int32(binary.BigEndian.Uint32(buf))
}
//小端序转大端序
func LSBtoMSB(lsb []byte) []byte{
	buf := make([]byte, len(lsb))
	index := 0
	for i:= len(lsb)-1;i>=0;i--{
		buf[index] = lsb[i]
		index++
	}
	return buf
}