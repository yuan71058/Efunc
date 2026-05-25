// 字节缓冲对象池
// 基于 valyala/bytebufferpool 库，提供高性能字节缓冲区复用。
// 适用于高频字节拼接场景，有效减少 GC 压力。
package utils

import (
	"github.com/valyala/bytebufferpool"
)

// ByteBuffer_Get 从字节缓冲池中获取一个字节缓冲区。
// 获取的缓冲区是空的（长度为0），可以安全写入。
// 使用完毕后必须调用 ByteBuffer_Put 归还，避免内存泄漏。
//
// 返回:
//   - *bytebufferpool.ByteBuffer: 空的字节缓冲区指针
func ByteBuffer_Get() *bytebufferpool.ByteBuffer {
	return bytebufferpool.Get()
}

// ByteBuffer_Put 将字节缓冲区归还到池中。
// 归还后缓冲区内容会被自动清空，不可再使用该指针。
// 每次获取后必须调用此函数归还，否则会导致内存泄漏。
//
// 参数:
//   - buf: 之前通过 ByteBuffer_Get 获取的缓冲区指针
func ByteBuffer_Put(buf *bytebufferpool.ByteBuffer) {
	bytebufferpool.Put(buf)
}

// ByteBuffer_GetBytes 从对象池获取缓冲区，写入数据后取出字节集副本，再归还缓冲区。
// 适用于临时需要字节集切片的场景，减少 GC 压力。
//
// 参数:
//   - data: 要写入缓冲区的数据
//
// 返回:
//   - []byte: 数据的字节集副本
func ByteBuffer_GetBytes(data []byte) []byte {
	buf := bytebufferpool.Get()
	buf.Set(data)
	result := make([]byte, len(buf.B))
	copy(result, buf.B)
	bytebufferpool.Put(buf)
	return result
}

// ByteBuffer_GetString 从对象池获取缓冲区，写入字符串后取出文本副本，再归还缓冲区。
//
// 参数:
//   - s: 要写入缓冲区的字符串
//
// 返回:
//   - string: 文本副本
func ByteBuffer_GetString(s string) string {
	buf := bytebufferpool.Get()
	buf.SetString(s)
	result := string(buf.B)
	bytebufferpool.Put(buf)
	return result
}