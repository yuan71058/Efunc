// 文件监控工具
// 基于 fsnotify 库，提供文件和目录的实时变化监控功能。
// 支持监控创建(Create)、修改(Write)、删除(Remove)、重命名(Rename)、权限变更(Chmod)等事件。
package utils

import (
	"path/filepath"

	"github.com/fsnotify/fsnotify"
)

// FileWatcher_New 创建一个新的文件系统监控器。
// 基于 fsnotify 库实现，可监控文件和目录的创建、修改、删除、重命名事件。
// 创建后需调用 FileWatcher_AddDir 或 FileWatcher_AddFile 添加监控目标，
// 再调用 FileWatcher_Start 启动事件监听。
//
// 返回:
//   - *fsnotify.Watcher: 监控器对象指针
//   - error: 创建失败时返回错误
func FileWatcher_New() (*fsnotify.Watcher, error) {
	return fsnotify.NewWatcher()
}

// FileWatcher_AddDir 添加一个目录到监控列表。
// 监控目录下的所有文件和子目录的变化。
//
// 参数:
//   - w: 监控器对象指针
//   - dirPath: 要监控的目录路径
//
// 返回:
//   - error: 添加失败时返回错误
func FileWatcher_AddDir(w *fsnotify.Watcher, dirPath string) error {
	return w.Add(dirPath)
}

// FileWatcher_AddFile 添加一个文件到监控列表。
//
// 参数:
//   - w: 监控器对象指针
//   - filePath: 要监控的文件路径
//
// 返回:
//   - error: 添加失败时返回错误
func FileWatcher_AddFile(w *fsnotify.Watcher, filePath string) error {
	return w.Add(filePath)
}

// FileWatcher_Remove 从监控列表中移除一个文件或目录。
//
// 参数:
//   - w: 监控器对象指针
//   - path: 要移除的文件或目录路径
//
// 返回:
//   - error: 移除失败时返回错误
func FileWatcher_Remove(w *fsnotify.Watcher, path string) error {
	return w.Remove(path)
}

// FileWatcher_Close 关闭监控器，释放资源。
// 关闭后监控器不再接收任何事件。
//
// 参数:
//   - w: 监控器对象指针
func FileWatcher_Close(w *fsnotify.Watcher) {
	w.Close()
}

// FileWatcher_Events 获取监控器的事件通道。
// 通过此通道接收文件系统变化事件，用于自定义事件处理逻辑。
//
// 参数:
//   - w: 监控器对象指针
//
// 返回:
//   - <-chan fsnotify.Event: 只读事件通道
func FileWatcher_Events(w *fsnotify.Watcher) <-chan fsnotify.Event {
	return w.Events
}

// FileWatcher_Errors 获取监控器的错误通道。
// 通过此通道接收监控过程中发生的错误。
//
// 参数:
//   - w: 监控器对象指针
//
// 返回:
//   - <-chan error: 只读错误通道
func FileWatcher_Errors(w *fsnotify.Watcher) <-chan error {
	return w.Errors
}

// FileWatcher_Start 开始监听文件系统事件。
// 阻塞式监听，当文件发生变化时调用回调函数。
// 回调函数接收事件对象，包含事件类型和文件路径。
// 常见事件类型: Create(创建), Write(修改), Remove(删除), Rename(重命名), Chmod(权限变更)。
//
// 参数:
//   - w: 监控器对象指针
//   - fn: 事件回调，参数为 fsnotify.Event
func FileWatcher_Start(w *fsnotify.Watcher, fn func(event fsnotify.Event)) {
	for {
		select {
		case event, ok := <-w.Events:
			if !ok {
				return
			}
			fn(event)
		case _, ok := <-w.Errors:
			if !ok {
				return
			}
		}
	}
}

// FileWatcher_WatchDir 便捷函数：监控指定目录的文件变化。
// 自动创建监控器、添加目录，并在检测到变化时调用回调。
// 返回停止函数，调用即可停止监控。
//
// 参数:
//   - dirPath: 要监控的目录路径
//   - fn: 文件变化时的回调函数
//
// 返回:
//   - func(): 停止监控的函数
//   - error: 创建监控器或添加目录失败时返回错误
func FileWatcher_WatchDir(dirPath string, fn func(event fsnotify.Event)) (func(), error) {
	w, err := fsnotify.NewWatcher()
	if err != nil {
		return nil, err
	}

	absPath, _ := filepath.Abs(dirPath)
	err = w.Add(absPath)
	if err != nil {
		w.Close()
		return nil, err
	}

	done := make(chan struct{})
	go func() {
		for {
			select {
			case <-done:
				return
			case event, ok := <-w.Events:
				if !ok {
					return
				}
				fn(event)
			case _, ok := <-w.Errors:
				if !ok {
					return
				}
			}
		}
	}()

	return func() {
		close(done)
		w.Close()
	}, nil
}