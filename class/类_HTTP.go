package class

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"
)

// L_HTTP服务端 HTTP 服务端，基于标准库 net/http 封装，支持路由注册、中间件、静态文件服务。
// 提供中文 API 风格的 HTTP 服务器构建方法。
//
// 使用示例:
//
//	s := &class.L_HTTP服务端{}
//	s.T注册路由("GET", "/hello", func(w http.ResponseWriter, r *http.Request) {
//	    w.Write([]byte("Hello World"))
//	})
//	s.Q启动(8080)
type L_HTTP服务端 struct {
	mu       sync.Mutex
	server   *http.Server
	mux      *http.ServeMux
	running  bool
	中间件列表 []func(http.HandlerFunc) http.HandlerFunc
}

// Q启动 启动 HTTP 服务端，监听指定端口。
//
// 参数:
//   - 端口: 监听端口号（如 8080）
//
// 返回:
//   - error: 启动失败时返回错误
func (s *L_HTTP服务端) Q启动(端口 int) error {
	return s.Q启动带地址(fmt.Sprintf(":%d", 端口))
}

// Q启动带地址 启动 HTTP 服务端，指定完整监听地址。
//
// 参数:
//   - 地址: 监听地址（如 ":8080" 或 "0.0.0.0:8080"）
//
// 返回:
//   - error: 启动失败时返回错误
func (s *L_HTTP服务端) Q启动带地址(地址 string) error {
	if s.mux == nil {
		s.mux = http.NewServeMux()
	}
	s.running = true

	var handler http.Handler = s.mux
	for i := len(s.中间件列表) - 1; i >= 0; i-- {
		handler = s.中间件列表[i](handler.ServeHTTP)
	}

	s.server = &http.Server{Addr: 地址, Handler: handler}

	go func() {
		if err := s.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			return
		}
	}()

	return nil
}

// T停止 停止 HTTP 服务端，等待已有请求处理完毕（5秒超时）。
//
// 返回:
//   - error: 停止失败时返回错误
func (s *L_HTTP服务端) T停止() error {
	s.running = false
	if s.server != nil {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		return s.server.Shutdown(ctx)
	}
	return nil
}

// T注册路由 注册 HTTP 路由处理函数。
//
// 参数:
//   - 方法: HTTP 方法（如 "GET", "POST", "PUT", "DELETE" 等）
//   - 路径: URL 路径（如 "/api/hello"）
//   - 处理函数: 标准 http.HandlerFunc 处理函数
func (s *L_HTTP服务端) T注册路由(方法 string, 路径 string, 处理函数 http.HandlerFunc) {
	if s.mux == nil {
		s.mux = http.NewServeMux()
	}
	pattern := fmt.Sprintf("%s %s", 方法, 路径)
	s.mux.HandleFunc(pattern, 处理函数)
}

// T注册通用路由 注册不区分 HTTP 方法的通配路由。
// 所有 HTTP 方法匹配该路径时均触发处理函数。
//
// 参数:
//   - 路径: URL 路径（如 "/hello"）
//   - 处理函数: 标准 http.HandlerFunc 处理函数
func (s *L_HTTP服务端) T注册通用路由(路径 string, 处理函数 http.HandlerFunc) {
	if s.mux == nil {
		s.mux = http.NewServeMux()
	}
	s.mux.HandleFunc(路径, 处理函数)
}

// J静态文件服务 注册静态文件服务，将指定目录映射到 URL 路径前缀。
//
// 参数:
//   - URL前缀: URL 路径前缀（如 "/static/"）
//   - 本地目录: 本地文件目录路径（如 "./public"）
func (s *L_HTTP服务端) J静态文件服务(URL前缀 string, 本地目录 string) {
	if s.mux == nil {
		s.mux = http.NewServeMux()
	}
	fs := http.FileServer(http.Dir(本地目录))
	s.mux.Handle(URL前缀, http.StripPrefix(URL前缀, fs))
}

// Z中间件 添加全局中间件函数。
// 中间件按添加顺序执行。
//
// 参数:
//   - 中间件函数: 接受 http.HandlerFunc 返回 http.HandlerFunc 的包装函数
func (s *L_HTTP服务端) Z中间件(中间件函数 func(http.HandlerFunc) http.HandlerFunc) {
	s.中间件列表 = append(s.中间件列表, 中间件函数)
}

// Z中间件CORS 快捷添加 CORS 跨域中间件，允许所有来源访问。
// 支持 GET/POST/PUT/DELETE/OPTIONS/PATCH 方法及常用请求头。
func (s *L_HTTP服务端) Z中间件CORS() {
	s.Z中间件(func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS, PATCH")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With")
			if r.Method == "OPTIONS" {
				w.WriteHeader(http.StatusNoContent)
				return
			}
			next(w, r)
		}
	})
}

// Z中间件日志 快捷添加请求日志中间件。
// 记录每个请求的方法、路径、耗时。
func (s *L_HTTP服务端) Z中间件日志() {
	s.Z中间件(func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			next(w, r)
			fmt.Printf("[HTTP] %s %s %v\n", r.Method, r.URL.Path, time.Since(start))
		}
	})
}

// Q取启动地址 获取服务监听地址。
//
// 返回:
//   - string: 监听地址
func (s *L_HTTP服务端) Q取启动地址() string {
	if s.server != nil {
		return s.server.Addr
	}
	return ""
}

// S是否运行中 检查服务是否正在运行。
//
// 返回:
//   - bool: true 表示运行中
func (s *L_HTTP服务端) S是否运行中() bool {
	return s.running
}

// F响应JSON 便捷方法：向客户端写入 JSON 响应。
// 自动设置 Content-Type 为 application/json。
//
// 参数:
//   - w: http.ResponseWriter
//   - 状态码: HTTP 状态码（如 200）
//   - 数据: 要序列化为 JSON 的数据
func F响应JSON(w http.ResponseWriter, 状态码 int, 数据 interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(状态码)
	json.NewEncoder(w).Encode(数据)
}

// F响应文本 便捷方法：向客户端写入纯文本响应。
//
// 参数:
//   - w: http.ResponseWriter
//   - 状态码: HTTP 状态码
//   - 文本: 要返回的文本内容
func F响应文本(w http.ResponseWriter, 状态码 int, 文本 string) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(状态码)
	w.Write([]byte(文本))
}

// F响应HTML 便捷方法：向客户端写入 HTML 响应。
//
// 参数:
//   - w: http.ResponseWriter
//   - 状态码: HTTP 状态码
//   - html: HTML 内容
func F响应HTML(w http.ResponseWriter, 状态码 int, html string) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(状态码)
	w.Write([]byte(html))
}

// F取查询参数 从 HTTP 请求中获取 URL 查询参数值。
//
// 参数:
//   - r: *http.Request
//   - 参数名: 查询参数名称
//
// 返回:
//   - string: 参数值，不存在则返回空字符串
func F取查询参数(r *http.Request, 参数名 string) string {
	return r.URL.Query().Get(参数名)
}

// F取POST参数 从 HTTP 请求中解析并获取 POST 表单参数值。
//
// 参数:
//   - r: *http.Request
//   - 参数名: POST 参数名称
//
// 返回:
//   - string: 参数值，解析失败或不存在则返回空字符串
func F取POST参数(r *http.Request, 参数名 string) string {
	r.ParseForm()
	return r.FormValue(参数名)
}

// F解析JSON请求体 从 HTTP 请求体中解析 JSON 到目标对象。
//
// 参数:
//   - r: *http.Request
//   - 目标: 用于接收解析结果的对象指针
//
// 返回:
//   - error: 解析失败时返回错误
func F解析JSON请求体(r *http.Request, 目标 interface{}) error {
	return json.NewDecoder(r.Body).Decode(目标)
}