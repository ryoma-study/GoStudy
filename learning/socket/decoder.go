package main

import (
	"encoding/binary"
	"fmt"
)

func main() {
	sourceData := "Hello, World!"
	encodeData := encoder(sourceData)

	decodeData := decoder(encodeData)
	fmt.Printf("\nsourceData: %v\nencodeData： %v\ndecodeData: %v\n", sourceData, encodeData, decodeData)
}

const (
	// size
	// 包长度
	_packSize = 4
	// 头长度
	_headerSize = 2
	// 版本号
	_verSize = 2
	// 操作码？
	_opSize = 4
	// 序号
	_seqSize       = 4
	_rawHeaderSize = _packSize + _headerSize + _verSize + _opSize + _seqSize
	// offset
	_packOffset   = 0
	_headerOffset = _packOffset + _packSize
	_verOffset    = _headerOffset + _headerSize
	_opOffset     = _verOffset + _verSize
	_seqOffset    = _opOffset + _opSize
)

func decoder(data []byte) string {
	packetLen := binary.BigEndian.Uint32(data[_packOffset:_headerOffset])
	fmt.Printf("packetLen: %v\n", packetLen)

	headerLen := binary.BigEndian.Uint16(data[_headerOffset:_verOffset])
	fmt.Printf("headerLen: %v\n", headerLen)

	version := binary.BigEndian.Uint16(data[_verOffset:_opOffset])
	fmt.Printf("version: %v\n", version)

	operation := binary.BigEndian.Uint32(data[_opOffset:_seqOffset])
	fmt.Printf("operation: %v\n", operation)

	sequence := binary.BigEndian.Uint32(data[_seqOffset:_rawHeaderSize])
	fmt.Printf("sequence: %v\n", sequence)

	body := string(data[_rawHeaderSize:])
	fmt.Printf("body: %v\n", body)

	return body
}

func encoder(body string) []byte {
	packetLen := len(body) + _rawHeaderSize
	ret := make([]byte, packetLen)

	binary.BigEndian.PutUint32(ret[_packOffset:], uint32(packetLen))
	binary.BigEndian.PutUint16(ret[_headerOffset:], uint16(_rawHeaderSize))

	version := 28
	binary.BigEndian.PutUint16(ret[_verOffset:], uint16(version))
	operation := 36
	binary.BigEndian.PutUint32(ret[_opOffset:], uint32(operation))
	sequence := 48
	binary.BigEndian.PutUint32(ret[_seqOffset:], uint32(sequence))

	byteBody := []byte(body)
	copy(ret[_rawHeaderSize:], byteBody)

	return ret
}
