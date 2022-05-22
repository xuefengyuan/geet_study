package main

import (
    "encoding/binary"
    "fmt"
)

func main() {

    data := encoder("hello goim")

    decoder(data)
}

// decoder
//  4 bytes Package Length，包长度
//  2 bytes Header Length，头长度
//  2 bytes Protocol Version，协议版本
//  4 bytes Operation，操作码
//  4 bytes Sequence 请求序号ID
//  Packlen-Headerlen(包长度-头长度) Body，包内容
func decoder(data []byte) {

    if len(data) <= 16 {
        fmt.Println("data len < 16 len")
        return
    }
    packLen := binary.BigEndian.Uint32(data[:4]) // 获取前面4 byte整个包的长度
    fmt.Printf("packLen : %d\n", packLen)

    headerLen := binary.BigEndian.Uint16(data[4:6]) // 获取 4-6 byte head的长
    fmt.Printf("headLen : %d\n", headerLen)

    version := binary.BigEndian.Uint16(data[6:8]) // 获取6-8 byte 协议版本号
    fmt.Printf("version : %d\n", version)

    operation := binary.BigEndian.Uint32(data[8:12]) // 获取8-12 byte 操作码
    fmt.Printf("operation : %d\n", operation)

    sequence := binary.BigEndian.Uint32(data[12:16]) // 获取12-16 byte 获取序列号，唯一标识
    fmt.Printf("sequence : %d\n", sequence)

    body := string(data[16:])
    fmt.Printf("body : %s\n", body)
}

// encoder
//  4 bytes Package Length，包长度
//  2 bytes Header Length，头长度
//  2 bytes Protocol Version，协议版本
//  4 bytes Operation，操作码
//  4 bytes Sequence 请求序号ID
//  Packlen-Headerlen(包长度-头长度) Body，包内容
func encoder(body string) []byte {
    headerLen := 16
    packeLen := len(body) + headerLen
    ret := make([]byte, packeLen)

    binary.BigEndian.PutUint32(ret[:4], uint32(packeLen))   // 写入整体消息长度
    binary.BigEndian.PutUint16(ret[4:6], uint16(headerLen)) // 写入head长度

    version := 6
    binary.BigEndian.PutUint16(ret[6:8], uint16(version)) // 写入版本号
    operation := 99
    binary.BigEndian.PutUint32(ret[8:12], uint32(operation)) // 写入操作码
    sequence := 8
    binary.BigEndian.PutUint32(ret[12:16], uint32(sequence)) // 客户序列号

    byteBody := []byte(body)
    copy(ret[16:], byteBody) // 写入body内容

    return ret

}
