package frame

import (
	"encoding/binary"
	"errors"
	"io"
)

type FramePayload []byte

type StreamFrameCodec interface {
	Encode(io.Writer, FramePayload) error   // data -> frame，并写入io.Writer
	Decode(io.Reader) (FramePayload, error) // 从io.Reader中提取frame payload，并返回给上层
}

//errors.New 使用字符串创建一个错误,请类比fmt包的Errorf方法，差不多可以认为是New(fmt.Sprintf(...))

var ErrShortWrite = errors.New("short write")
var ErrShortRead = errors.New("short read")

type myFrameCodec struct{}

func NewMyFrameCodec() StreamFrameCodec {
	return &myFrameCodec{}
}

func (p *myFrameCodec) Encode(w io.Writer, framePayload FramePayload) error {
	var f = framePayload
	var totalLen = int32(len(framePayload)) + 4

	/*Write 将data的binary编码格式写入w，data必须是定长值、定长值的切片、定长值的指针。
	order指定写入数据的字节序，写入结构体时，名字中有'_'的字段会置为0。*/
	err := binary.Write(w, binary.BigEndian, &totalLen)
	if err != nil {
		return err
	}

	n, err := w.Write([]byte(f)) // write the frame payload to outbound stream
	if err != nil {
		return err
	}

	if n != len(framePayload) {
		return ErrShortWrite
	}
	return nil
}

func (p *myFrameCodec) Decode(r io.Reader) (FramePayload, error) {
	var totalLen int32

	/*从r中读取binary编码的数据并赋给data，data必须是一个指向定长值的指针或者定长值的切片。
	从r读取的字节使用order指定的字节序解码并写入data的字段里当写入结构体是，
	名字中有'_'的字段会被跳过，这些字段可用于填充（内存空间）。*/
	err := binary.Read(r, binary.BigEndian, &totalLen)
	if err != nil {
		return nil, err
	}

	buf := make([]byte, totalLen-4)
	/*ReadFull从r精确地读取len(buf)字节数据填充进buf。函数返回写入的字节数和错误（如果没有读取足够的字节）。
	只有没有读取到字节时才可能返回EOF；如果读取了有但不够的字节时遇到了EOF，函数会返回ErrUnexpectedEOF。
	只有返回值err为nil时，返回值n才会等于len(buf)。*/
	n, err := io.ReadFull(r, buf)
	if err != nil {
		return nil, err
	}
	if n != int(totalLen-4) {
		return nil, ErrShortRead
	}
	return FramePayload(buf), nil
}
