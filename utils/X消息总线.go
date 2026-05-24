package utils

import (
	messagebus "github.com/vardius/message-bus"
)

// X消息_创建 创建一个新的消息总线。
// 消息总线支持发布-订阅模式，用于 goroutine 间的异步通信。
// 订阅者通过主题（topic）接收消息，发布者向主题发送消息。
//
// 参数:
//   - 最大订阅者数: 每个主题允许的最大订阅者数量，0 表示不限制
//
// 返回:
//   - messagebus.MessageBus: 消息总线实例
func X消息_创建(最大订阅者数 int) messagebus.MessageBus {
	return messagebus.New(最大订阅者数)
}

// X消息_发布 向指定主题发布一条消息。
// 所有订阅了该主题的订阅者都会收到此消息。
// 发布操作是异步的，不会等待订阅者处理完成。
//
// 参数:
//   - 总线: 消息总线实例
//   - 主题: 消息主题名称
//   - 消息: 要发布的消息内容（可变参数）
func X消息_发布(总线 messagebus.MessageBus, 主题 string, 消息 ...interface{}) {
	总线.Publish(主题, 消息...)
}

// X消息_订阅 订阅指定主题的消息。
// 当有消息发布到该主题时，回调函数会被异步调用。
//
// 参数:
//   - 总线: 消息总线实例
//   - 主题: 要订阅的主题名称
//   - 回调函数: 接收消息的回调函数
//
// 返回:
//   - error: 订阅失败时返回错误
func X消息_订阅(总线 messagebus.MessageBus, 主题 string, 回调函数 interface{}) error {
	return 总线.Subscribe(主题, 回调函数)
}

// X消息_取消订阅 取消订阅指定主题。
// 取消后不再接收该主题的消息。
//
// 参数:
//   - 总线: 消息总线实例
//   - 主题: 要取消订阅的主题名称
//   - 回调函数: 之前订阅时使用的回调函数
//
// 返回:
//   - error: 取消订阅失败时返回错误
func X消息_取消订阅(总线 messagebus.MessageBus, 主题 string, 回调函数 interface{}) error {
	return 总线.Unsubscribe(主题, 回调函数)
}

// X消息_关闭主题 关闭指定主题，释放相关资源。
// 关闭后该主题的所有订阅者将被移除。
//
// 参数:
//   - 总线: 消息总线实例
//   - 主题: 要关闭的主题名称
func X消息_关闭主题(总线 messagebus.MessageBus, 主题 string) {
	总线.Close(主题)
}
