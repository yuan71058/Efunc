package class

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

var wsUpgrader = websocket.Upgrader{
	ReadBufferSize:  4096,
	WriteBufferSize: 4096,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// WSServer WebSocket 服务端，支持多客户端并发连接、文本/二进制消息收发和广播。
// 基于 gorilla/websocket 库实现，可同时处理文本和二进制消息。
//
// 使用示例:
//
//	s := &class.WSServer{}
//	s.OnReceiveText = func(clientID string, text string) { fmt.Println("收到:", text) }
//	s.Start(8080, "/ws")
type WSServer struct {
	mu             sync.Mutex
	server         *http.Server
	clients        map[string]*websocket.Conn
	running        bool
	OnReceiveText  func(clientID string, text string)
	OnReceiveBytes func(clientID string, data []byte)
	OnConnect      func(clientID string)
	OnDisconnect   func(clientID string)
}

// Start 启动 WebSocket 服务端，监听指定端口和路径。
// 客户端连接 ws://host:port/path 即可建立 WebSocket 连接。
//
// 参数:
//   - port: 监听端口号（如 8080）
//   - path: WebSocket 路径（如 "/ws"）
//
// 返回:
//   - error: 启动失败时返回错误
func (s *WSServer) Start(port int, path string) error {
	s.clients = make(map[string]*websocket.Conn)
	s.running = true

	mux := http.NewServeMux()
	mux.HandleFunc(path, s.handleWS)
	mux.HandleFunc(path+"/", s.handleWS)

	s.server = &http.Server{Addr: fmt.Sprintf(":%d", port), Handler: mux}

	go func() {
		if err := s.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			return
		}
	}()

	return nil
}

// StartWithAddr 启动 WebSocket 服务端，指定完整监听地址和路径。
//
// 参数:
//   - addr: 监听地址（如 ":8080" 或 "0.0.0.0:8080"）
//   - path: WebSocket 路径（如 "/ws"）
//
// 返回:
//   - error: 启动失败时返回错误
func (s *WSServer) StartWithAddr(addr string, path string) error {
	s.clients = make(map[string]*websocket.Conn)
	s.running = true

	mux := http.NewServeMux()
	mux.HandleFunc(path, s.handleWS)
	mux.HandleFunc(path+"/", s.handleWS)

	s.server = &http.Server{Addr: addr, Handler: mux}

	go func() {
		if err := s.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			return
		}
	}()

	return nil
}

func (s *WSServer) handleWS(w http.ResponseWriter, r *http.Request) {
	conn, err := wsUpgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}

	clientID := conn.RemoteAddr().String()
	s.mu.Lock()
	s.clients[clientID] = conn
	s.mu.Unlock()

	if s.OnConnect != nil {
		go s.OnConnect(clientID)
	}

	defer func() {
		conn.Close()
		s.mu.Lock()
		delete(s.clients, clientID)
		s.mu.Unlock()
		if s.OnDisconnect != nil {
			go s.OnDisconnect(clientID)
		}
	}()

	for s.running {
		msgType, data, err := conn.ReadMessage()
		if err != nil {
			return
		}
		switch msgType {
		case websocket.TextMessage:
			if s.OnReceiveText != nil {
				go s.OnReceiveText(clientID, string(data))
			}
		case websocket.BinaryMessage:
			if s.OnReceiveBytes != nil {
				go s.OnReceiveBytes(clientID, data)
			}
		}
	}
}

// Stop 停止 WebSocket 服务端，关闭所有连接。
func (s *WSServer) Stop() {
	s.running = false
	s.mu.Lock()
	defer s.mu.Unlock()
	for _, conn := range s.clients {
		conn.Close()
	}
	s.clients = make(map[string]*websocket.Conn)
	if s.server != nil {
		s.server.Close()
	}
}

// SendText 向指定客户端发送文本消息。
//
// 参数:
//   - clientID: 客户端的远程地址标识
//   - text: 要发送的文本
//
// 返回:
//   - error: 发送失败时返回错误
func (s *WSServer) SendText(clientID string, text string) error {
	s.mu.Lock()
	conn, ok := s.clients[clientID]
	s.mu.Unlock()
	if !ok {
		return websocket.ErrCloseSent
	}
	return conn.WriteMessage(websocket.TextMessage, []byte(text))
}

// SendBytes 向指定客户端发送二进制消息。
//
// 参数:
//   - clientID: 客户端的远程地址标识
//   - data: 要发送的字节数据
//
// 返回:
//   - error: 发送失败时返回错误
func (s *WSServer) SendBytes(clientID string, data []byte) error {
	s.mu.Lock()
	conn, ok := s.clients[clientID]
	s.mu.Unlock()
	if !ok {
		return websocket.ErrCloseSent
	}
	return conn.WriteMessage(websocket.BinaryMessage, data)
}

// BroadcastText 向所有已连接客户端广播文本消息。
//
// 参数:
//   - text: 要广播的文本
func (s *WSServer) BroadcastText(text string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for _, conn := range s.clients {
		conn.WriteMessage(websocket.TextMessage, []byte(text))
	}
}

// BroadcastBytes 向所有已连接客户端广播二进制消息。
//
// 参数:
//   - data: 要广播的字节数据
func (s *WSServer) BroadcastBytes(data []byte) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for _, conn := range s.clients {
		conn.WriteMessage(websocket.BinaryMessage, data)
	}
}

// ClientCount 获取当前连接的客户端数量。
//
// 返回:
//   - int: 客户端数量
func (s *WSServer) ClientCount() int {
	s.mu.Lock()
	defer s.mu.Unlock()
	return len(s.clients)
}

// ClientList 获取所有已连接客户端的 ID 列表。
//
// 返回:
//   - []string: 客户端 ID 列表
func (s *WSServer) ClientList() []string {
	s.mu.Lock()
	defer s.mu.Unlock()
	list := make([]string, 0, len(s.clients))
	for id := range s.clients {
		list = append(list, id)
	}
	return list
}

// WSClient WebSocket 客户端，支持连接服务端、收发文本和二进制消息。
//
// 使用示例:
//
//	c := &class.WSClient{}
//	c.OnReceiveText = func(text string) { fmt.Println(text) }
//	c.Connect("ws://127.0.0.1:8080/ws")
//	c.SendText("hello")
type WSClient struct {
	conn           *websocket.Conn
	running        bool
	mu             sync.Mutex
	OnReceiveText  func(text string)
	OnReceiveBytes func(data []byte)
	OnDisconnect   func()
}

// Connect 连接到 WebSocket 服务端。
// 连接成功后启动接收 goroutine，通过回调函数通知上层。
//
// 参数:
//   - url: WebSocket URL（如 "ws://127.0.0.1:8080/ws"）
//
// 返回:
//   - error: 连接失败时返回错误
func (c *WSClient) Connect(url string) error {
	conn, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		return err
	}
	c.conn = conn
	c.running = true
	go c.receive()
	return nil
}

func (c *WSClient) receive() {
	defer func() {
		c.running = false
		c.conn.Close()
		if c.OnDisconnect != nil {
			c.OnDisconnect()
		}
	}()

	for c.running {
		msgType, data, err := c.conn.ReadMessage()
		if err != nil {
			return
		}
		switch msgType {
		case websocket.TextMessage:
			if c.OnReceiveText != nil {
				go c.OnReceiveText(string(data))
			}
		case websocket.BinaryMessage:
			if c.OnReceiveBytes != nil {
				go c.OnReceiveBytes(data)
			}
		}
	}
}

// Disconnect 断开与服务端的连接。
func (c *WSClient) Disconnect() {
	c.running = false
	if c.conn != nil {
		c.conn.Close()
	}
}

// SendText 向服务端发送文本消息。
//
// 参数:
//   - text: 要发送的文本
//
// 返回:
//   - error: 发送失败时返回错误
func (c *WSClient) SendText(text string) error {
	return c.conn.WriteMessage(websocket.TextMessage, []byte(text))
}

// SendBytes 向服务端发送二进制消息。
//
// 参数:
//   - data: 要发送的字节数据
//
// 返回:
//   - error: 发送失败时返回错误
func (c *WSClient) SendBytes(data []byte) error {
	return c.conn.WriteMessage(websocket.BinaryMessage, data)
}

// SendJSON 向服务端发送 JSON 消息。
//
// 参数:
//   - v: 要序列化为 JSON 的对象
//
// 返回:
//   - error: 发送失败时返回错误
func (c *WSClient) SendJSON(v interface{}) error {
	return c.conn.WriteJSON(v)
}

// IsConnected 检查客户端是否已连接。
//
// 返回:
//   - bool: true 表示已连接
func (c *WSClient) IsConnected() bool {
	return c.running && c.conn != nil
}