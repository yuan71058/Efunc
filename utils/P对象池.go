package utils

import (
	"github.com/valyala/bytebufferpool"
)

// P对象池_获取 从字节缓冲池中获取一个字节缓冲区。
// 获取的缓冲区是空的（长度为0），可以安全写入。
// 使用完毕后必须调用 P对象池_放回 归还，避免内存泄漏。
//
// 返回:
//   - *bytebufferpool.ByteBuffer: 空的字节缓冲区指针
func P对象池_获取() *bytebufferpool.ByteBuffer {
	return bytebufferpool.Get()
}

// P对象池_放回 将字节缓冲区归还到池中。
// 归还后缓冲区内容会被自动清空，不可再使用该指针。
// 每次获取后必须调用此函数归还，否则会导致内存泄漏。
//
// 参数:
//   - 缓冲区: 之前通过 P对象池_获取 获取的缓冲区指针
func P对象池_放回(缓冲区 *bytebufferpool.ByteBuffer) {
	bytebufferpool.Put(缓冲区)
}

// P对象池_取字节集 从对象池获取缓冲区，写入数据后取出字节集副本，再归还缓冲区。
// 适用于临时需要字节集切片的场景，减少 GC 压力。
//
// 参数:
//   - 数据: 要写入缓冲区的数据
//
// 返回:
//   - []byte: 数据的字节集副本
func P对象池_取字节集(数据 []byte) []byte {
	buf := bytebufferpool.Get()
	buf.Set(数据)
	result := make([]byte, len(buf.B))
	copy(result, buf.B)
	bytebufferpool.Put(buf)
	return result
}

// P对象池_取文本 从对象池获取缓冲区，写入字符串后取出文本副本，再归还缓冲区。
//
// 参数:
//   - 文本: 要写入缓冲区的字符串
//
// 返回:
//   - string: 文本副本
func P对象池_取文本(文本 string) string {
	buf := bytebufferpool.Get()
	buf.SetString(文本)
	result := string(buf.B)
	bytebufferpool.Put(buf)
	return result
}
