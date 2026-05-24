package utils

import (
	"path/filepath"

	"github.com/fsnotify/fsnotify"
)

// F文件监控_创建 创建一个新的文件系统监控器。
// 基于 fsnotify 库实现，可监控文件和目录的创建、修改、删除、重命名事件。
// 创建后需调用 F文件监控_添加目录 或 F文件监控_添加文件 添加监控目标，
// 再调用 F文件监控_开始 启动事件监听。
//
// 返回:
//   - *fsnotify.Watcher: 监控器对象指针
//   - error: 创建失败时返回错误
func F文件监控_创建() (*fsnotify.Watcher, error) {
	return fsnotify.NewWatcher()
}

// F文件监控_添加目录 添加一个目录到监控列表。
// 监控目录下的所有文件和子目录的变化。
//
// 参数:
//   - 监控器: 监控器对象指针
//   - 目录路径: 要监控的目录路径
//
// 返回:
//   - error: 添加失败时返回错误
func F文件监控_添加目录(监控器 *fsnotify.Watcher, 目录路径 string) error {
	return 监控器.Add(目录路径)
}

// F文件监控_添加文件 添加一个文件到监控列表。
//
// 参数:
//   - 监控器: 监控器对象指针
//   - 文件路径: 要监控的文件路径
//
// 返回:
//   - error: 添加失败时返回错误
func F文件监控_添加文件(监控器 *fsnotify.Watcher, 文件路径 string) error {
	return 监控器.Add(文件路径)
}

// F文件监控_移除 从监控列表中移除一个文件或目录。
//
// 参数:
//   - 监控器: 监控器对象指针
//   - 路径: 要移除的文件或目录路径
//
// 返回:
//   - error: 移除失败时返回错误
func F文件监控_移除(监控器 *fsnotify.Watcher, 路径 string) error {
	return 监控器.Remove(路径)
}

// F文件监控_关闭 关闭监控器，释放资源。
// 关闭后监控器不再接收任何事件。
//
// 参数:
//   - 监控器: 监控器对象指针
func F文件监控_关闭(监控器 *fsnotify.Watcher) {
	监控器.Close()
}

// F文件监控_取事件通道 获取监控器的事件通道。
// 通过此通道接收文件系统变化事件，用于自定义事件处理逻辑。
//
// 参数:
//   - 监控器: 监控器对象指针
//
// 返回:
//   - <-chan fsnotify.Event: 只读事件通道
func F文件监控_取事件通道(监控器 *fsnotify.Watcher) <-chan fsnotify.Event {
	return 监控器.Events
}

// F文件监控_取错误通道 获取监控器的错误通道。
// 通过此通道接收监控过程中发生的错误。
//
// 参数:
//   - 监控器: 监控器对象指针
//
// 返回:
//   - <-chan error: 只读错误通道
func F文件监控_取错误通道(监控器 *fsnotify.Watcher) <-chan error {
	return 监控器.Errors
}

// F文件监控_开始 开始监听文件系统事件。
// 阻塞式监听，当文件发生变化时调用回调函数。
// 回调函数接收事件对象，包含事件类型和文件路径。
// 常见事件类型: Create(创建), Write(修改), Remove(删除), Rename(重命名), Chmod(权限变更)。
//
// 参数:
//   - 监控器: 监控器对象指针
//   - 回调函数: 事件回调，参数为 fsnotify.Event
func F文件监控_开始(监控器 *fsnotify.Watcher, 回调函数 func(event fsnotify.Event)) {
	for {
		select {
		case event, ok := <-监控器.Events:
			if !ok {
				return
			}
			回调函数(event)
		case _, ok := <-监控器.Errors:
			if !ok {
				return
			}
		}
	}
}

// F文件监控_监控目录变化 便捷函数：监控指定目录的文件变化。
// 自动创建监控器、添加目录，并在检测到变化时调用回调。
// 返回停止函数，调用即可停止监控。
//
// 参数:
//   - 目录路径: 要监控的目录路径
//   - 回调函数: 文件变化时的回调函数
//
// 返回:
//   - func(): 停止监控的函数
//   - error: 创建监控器或添加目录失败时返回错误
func F文件监控_监控目录变化(目录路径 string, 回调函数 func(event fsnotify.Event)) (func(), error) {
	监控器, err := fsnotify.NewWatcher()
	if err != nil {
		return nil, err
	}

	absPath, _ := filepath.Abs(目录路径)
	err = 监控器.Add(absPath)
	if err != nil {
		监控器.Close()
		return nil, err
	}

	done := make(chan struct{})
	go func() {
		for {
			select {
			case <-done:
				return
			case event, ok := <-监控器.Events:
				if !ok {
					return
				}
				回调函数(event)
			case _, ok := <-监控器.Errors:
				if !ok {
					return
				}
			}
		}
	}()

	return func() {
		close(done)
		监控器.Close()
	}, nil
}
