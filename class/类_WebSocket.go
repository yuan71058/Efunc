package class

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  4096,
	WriteBufferSize: 4096,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// L_WS服务端 WebSocket 服务端，支持多客户端并发连接、文本/二进制消息收发和广播。
// 基于 gorilla/websocket 库实现，可同时处理文本和二进制消息。
//
// 使用示例:
//
//	s := &class.L_WS服务端{}
//	s.S收到文本回调 = func(客户端ID string, 文本 string) { fmt.Println("收到:", 文本) }
//	s.Q启动(8080, "/ws")
type L_WS服务端 struct {
	mu             sync.Mutex
	server         *http.Server
	客户端列表     map[string]*websocket.Conn
	running        bool
	S收到文本回调  func(客户端ID string, 文本 string)
	S收到字节回调  func(客户端ID string, 数据 []byte)
	K客户端连接回调 func(客户端ID string)
	K客户端断开回调 func(客户端ID string)
}

// Q启动 启动 WebSocket 服务端，监听指定端口和路径。
// 客户端连接 ws://host:port/path 即可建立 WebSocket 连接。
//
// 参数:
//   - 端口: 监听端口号（如 8080）
//   - 路径: WebSocket 路径（如 "/ws"）
//
// 返回:
//   - error: 启动失败时返回错误
func (s *L_WS服务端) Q启动(端口 int, 路径 string) error {
	s.客户端列表 = make(map[string]*websocket.Conn)
	s.running = true

	mux := http.NewServeMux()
	mux.HandleFunc(路径, s.handleWS)
	mux.HandleFunc(路径+"/", s.handleWS)

	s.server = &http.Server{Addr: fmt.Sprintf(":%d", 端口), Handler: mux}

	go func() {
		if err := s.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			return
		}
	}()

	return nil
}

// Q启动带地址 启动 WebSocket 服务端，指定完整监听地址和路径。
//
// 参数:
//   - 地址: 监听地址（如 ":8080" 或 "0.0.0.0:8080"）
//   - 路径: WebSocket 路径（如 "/ws"）
//
// 返回:
//   - error: 启动失败时返回错误
func (s *L_WS服务端) Q启动带地址(地址 string, 路径 string) error {
	s.客户端列表 = make(map[string]*websocket.Conn)
	s.running = true

	mux := http.NewServeMux()
	mux.HandleFunc(路径, s.handleWS)
	mux.HandleFunc(路径+"/", s.handleWS)

	s.server = &http.Server{Addr: 地址, Handler: mux}

	go func() {
		if err := s.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			return
		}
	}()

	return nil
}

func (s *L_WS服务端) handleWS(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}

	客户端ID := conn.RemoteAddr().String()
	s.mu.Lock()
	s.客户端列表[客户端ID] = conn
	s.mu.Unlock()

	if s.K客户端连接回调 != nil {
		go s.K客户端连接回调(客户端ID)
	}

	defer func() {
		conn.Close()
		s.mu.Lock()
		delete(s.客户端列表, 客户端ID)
		s.mu.Unlock()
		if s.K客户端断开回调 != nil {
			go s.K客户端断开回调(客户端ID)
		}
	}()

	for s.running {
		msgType, data, err := conn.ReadMessage()
		if err != nil {
			return
		}
		switch msgType {
		case websocket.TextMessage:
			if s.S收到文本回调 != nil {
				go s.S收到文本回调(客户端ID, string(data))
			}
		case websocket.BinaryMessage:
			if s.S收到字节回调 != nil {
				go s.S收到字节回调(客户端ID, data)
			}
		}
	}
}

// T停止 停止 WebSocket 服务端，关闭所有连接。
func (s *L_WS服务端) T停止() {
	s.running = false
	s.mu.Lock()
	defer s.mu.Unlock()
	for _, conn := range s.客户端列表 {
		conn.Close()
	}
	s.客户端列表 = make(map[string]*websocket.Conn)
	if s.server != nil {
		s.server.Close()
	}
}

// F发送文本 向指定客户端发送文本消息。
//
// 参数:
//   - 客户端ID: 客户端的远程地址标识
//   - 文本: 要发送的文本
//
// 返回:
//   - error: 发送失败时返回错误
func (s *L_WS服务端) F发送文本(客户端ID string, 文本 string) error {
	s.mu.Lock()
	conn, ok := s.客户端列表[客户端ID]
	s.mu.Unlock()
	if !ok {
		return websocket.ErrCloseSent
	}
	return conn.WriteMessage(websocket.TextMessage, []byte(文本))
}

// F发送字节 向指定客户端发送二进制消息。
//
// 参数:
//   - 客户端ID: 客户端的远程地址标识
//   - 数据: 要发送的字节数据
//
// 返回:
//   - error: 发送失败时返回错误
func (s *L_WS服务端) F发送字节(客户端ID string, 数据 []byte) error {
	s.mu.Lock()
	conn, ok := s.客户端列表[客户端ID]
	s.mu.Unlock()
	if !ok {
		return websocket.ErrCloseSent
	}
	return conn.WriteMessage(websocket.BinaryMessage, 数据)
}

// G广播文本 向所有已连接客户端广播文本消息。
//
// 参数:
//   - 文本: 要广播的文本
func (s *L_WS服务端) G广播文本(文本 string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for _, conn := range s.客户端列表 {
		conn.WriteMessage(websocket.TextMessage, []byte(文本))
	}
}

// G广播字节 向所有已连接客户端广播二进制消息。
//
// 参数:
//   - 数据: 要广播的字节数据
func (s *L_WS服务端) G广播字节(数据 []byte) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for _, conn := range s.客户端列表 {
		conn.WriteMessage(websocket.BinaryMessage, 数据)
	}
}

// Q取客户端数量 获取当前连接的客户端数量。
//
// 返回:
//   - int: 客户端数量
func (s *L_WS服务端) Q取客户端数量() int {
	s.mu.Lock()
	defer s.mu.Unlock()
	return len(s.客户端列表)
}

// Q取客户端列表 获取所有已连接客户端的 ID 列表。
//
// 返回:
//   - []string: 客户端 ID 列表
func (s *L_WS服务端) Q取客户端列表() []string {
	s.mu.Lock()
	defer s.mu.Unlock()
	list := make([]string, 0, len(s.客户端列表))
	for id := range s.客户端列表 {
		list = append(list, id)
	}
	return list
}

// L_WS客户端 WebSocket 客户端，支持连接服务端、收发文本和二进制消息。
//
// 使用示例:
//
//	c := &class.L_WS客户端{}
//	c.S收到文本回调 = func(文本 string) { fmt.Println(文本) }
//	c.L连接("ws://127.0.0.1:8080/ws")
//	c.F发送文本("hello")
type L_WS客户端 struct {
	conn           *websocket.Conn
	running        bool
	mu             sync.Mutex
	S收到文本回调  func(文本 string)
	S收到字节回调  func(数据 []byte)
	D断开回调      func()
}

// L连接 连接到 WebSocket 服务端。
// 连接成功后启动接收 goroutine，通过回调函数通知上层。
//
// 参数:
//   - URL: WebSocket 地址（如 "ws://127.0.0.1:8080/ws"）
//
// 返回:
//   - error: 连接失败时返回错误
func (c *L_WS客户端) L连接(URL string) error {
	conn, _, err := websocket.DefaultDialer.Dial(URL, nil)
	if err != nil {
		return err
	}
	c.conn = conn
	c.running = true
	go c.receive()
	return nil
}

func (c *L_WS客户端) receive() {
	defer func() {
		c.running = false
		c.conn.Close()
		if c.D断开回调 != nil {
			c.D断开回调()
		}
	}()

	for c.running {
		msgType, data, err := c.conn.ReadMessage()
		if err != nil {
			return
		}
		switch msgType {
		case websocket.TextMessage:
			if c.S收到文本回调 != nil {
				c.S收到文本回调(string(data))
			}
		case websocket.BinaryMessage:
			if c.S收到字节回调 != nil {
				c.S收到字节回调(data)
			}
		}
	}
}

// D断开 断开与 WebSocket 服务端的连接。
func (c *L_WS客户端) D断开() {
	c.running = false
	if c.conn != nil {
		c.conn.Close()
	}
}

// F发送文本 向服务端发送文本消息。
//
// 参数:
//   - 文本: 要发送的文本
//
// 返回:
//   - error: 发送失败时返回错误
func (c *L_WS客户端) F发送文本(文本 string) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.conn == nil {
		return websocket.ErrCloseSent
	}
	return c.conn.WriteMessage(websocket.TextMessage, []byte(文本))
}

// F发送字节 向服务端发送二进制消息。
//
// 参数:
//   - 数据: 要发送的字节数据
//
// 返回:
//   - error: 发送失败时返回错误
func (c *L_WS客户端) F发送字节(数据 []byte) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.conn == nil {
		return websocket.ErrCloseSent
	}
	return c.conn.WriteMessage(websocket.BinaryMessage, 数据)
}

// F发送JSON 向服务端发送 JSON 格式消息（自动序列化）。
//
// 参数:
//   - 数据: 要发送的数据（将自动 JSON 序列化）
//
// 返回:
//   - error: 发送失败时返回错误
func (c *L_WS客户端) F发送JSON(数据 interface{}) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.conn == nil {
		return websocket.ErrCloseSent
	}
	return c.conn.WriteJSON(数据)
}

// S是否已连接 检查是否已连接到服务端。
//
// 返回:
//   - bool: true 表示已连接
func (c *L_WS客户端) S是否已连接() bool {
	return c.conn != nil && c.running
}

// Q取远程地址 获取远程服务端地址。
//
// 返回:
//   - string: 远程地址
func (c *L_WS客户端) Q取远程地址() string {
	if c.conn != nil {
		return c.conn.RemoteAddr().String()
	}
	return ""
}