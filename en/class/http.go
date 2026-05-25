package class

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"
)

// HTTPServer HTTP 服务端，基于标准库 net/http 封装，支持路由注册、中间件、静态文件服务。
// 提供中文 API 风格的 HTTP 服务器构建方法。
//
// 使用示例:
//
//	s := &class.HTTPServer{}
//	s.RegisterRoute("GET", "/hello", func(w http.ResponseWriter, r *http.Request) {
//	    w.Write([]byte("Hello World"))
//	})
//	s.Start(8080)
type HTTPServer struct {
	mu          sync.Mutex
	server      *http.Server
	mux         *http.ServeMux
	running     bool
	middlewares []func(http.HandlerFunc) http.HandlerFunc
}

// Start 启动 HTTP 服务端，监听指定端口。
//
// 参数:
//   - port: 监听端口号（如 8080）
//
// 返回:
//   - error: 启动失败时返回错误
func (s *HTTPServer) Start(port int) error {
	return s.StartWithAddr(fmt.Sprintf(":%d", port))
}

// StartWithAddr 启动 HTTP 服务端，指定完整监听地址。
//
// 参数:
//   - addr: 监听地址（如 ":8080" 或 "0.0.0.0:8080"）
//
// 返回:
//   - error: 启动失败时返回错误
func (s *HTTPServer) StartWithAddr(addr string) error {
	if s.mux == nil {
		s.mux = http.NewServeMux()
	}
	s.running = true

	var handler http.Handler = s.mux
	for i := len(s.middlewares) - 1; i >= 0; i-- {
		handler = s.middlewares[i](handler.ServeHTTP)
	}

	s.server = &http.Server{Addr: addr, Handler: handler}

	go func() {
		if err := s.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			return
		}
	}()

	return nil
}

// Stop 停止 HTTP 服务端，等待已有请求处理完毕（5秒超时）。
//
// 返回:
//   - error: 停止失败时返回错误
func (s *HTTPServer) Stop() error {
	s.running = false
	if s.server != nil {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		return s.server.Shutdown(ctx)
	}
	return nil
}

// RegisterRoute 注册 HTTP 路由处理函数。
//
// 参数:
//   - method: HTTP 方法（如 "GET", "POST", "PUT", "DELETE" 等）
//   - path: URL 路径（如 "/api/hello"）
//   - handler: 标准 http.HandlerFunc 处理函数
func (s *HTTPServer) RegisterRoute(method string, path string, handler http.HandlerFunc) {
	if s.mux == nil {
		s.mux = http.NewServeMux()
	}
	pattern := fmt.Sprintf("%s %s", method, path)
	s.mux.HandleFunc(pattern, handler)
}

// RegisterAnyRoute 注册不区分 HTTP 方法的通配路由。
// 所有 HTTP 方法匹配该路径时均触发处理函数。
//
// 参数:
//   - path: URL 路径（如 "/hello"）
//   - handler: 标准 http.HandlerFunc 处理函数
func (s *HTTPServer) RegisterAnyRoute(path string, handler http.HandlerFunc) {
	if s.mux == nil {
		s.mux = http.NewServeMux()
	}
	s.mux.HandleFunc(path, handler)
}

// StaticFileServe 注册静态文件服务，将指定目录映射到 URL 路径前缀。
//
// 参数:
//   - urlPrefix: URL 路径前缀（如 "/static/"）
//   - localDir: 本地文件目录路径（如 "./public"）
func (s *HTTPServer) StaticFileServe(urlPrefix string, localDir string) {
	if s.mux == nil {
		s.mux = http.NewServeMux()
	}
	fs := http.FileServer(http.Dir(localDir))
	s.mux.Handle(urlPrefix, http.StripPrefix(urlPrefix, fs))
}

// UseMiddleware 添加全局中间件函数。
// 中间件按添加顺序执行。
//
// 参数:
//   - mw: 接受 http.HandlerFunc 返回 http.HandlerFunc 的包装函数
func (s *HTTPServer) UseMiddleware(mw func(http.HandlerFunc) http.HandlerFunc) {
	s.middlewares = append(s.middlewares, mw)
}

// UseCORS 快捷添加 CORS 跨域中间件，允许所有来源访问。
// 支持 GET/POST/PUT/DELETE/OPTIONS/PATCH 方法及常用请求头。
func (s *HTTPServer) UseCORS() {
	s.UseMiddleware(func(next http.HandlerFunc) http.HandlerFunc {
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

// UseLogger 快捷添加请求日志中间件。
// 记录每个请求的方法、路径、耗时。
func (s *HTTPServer) UseLogger() {
	s.UseMiddleware(func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			next(w, r)
			fmt.Printf("[HTTP] %s %s %v\n", r.Method, r.URL.Path, time.Since(start))
		}
	})
}

// ListenAddr 获取服务监听地址。
//
// 返回:
//   - string: 监听地址
func (s *HTTPServer) ListenAddr() string {
	if s.server != nil {
		return s.server.Addr
	}
	return ""
}

// IsRunning 检查服务是否正在运行。
//
// 返回:
//   - bool: true 表示运行中
func (s *HTTPServer) IsRunning() bool {
	return s.running
}

// RespondJSON 便捷方法：向客户端写入 JSON 响应。
// 自动设置 Content-Type 为 application/json。
//
// 参数:
//   - w: http.ResponseWriter
//   - statusCode: HTTP 状态码（如 200）
//   - data: 要序列化为 JSON 的数据
func RespondJSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}

// RespondText 便捷方法：向客户端写入纯文本响应。
//
// 参数:
//   - w: http.ResponseWriter
//   - statusCode: HTTP 状态码
//   - text: 要返回的文本内容
func RespondText(w http.ResponseWriter, statusCode int, text string) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(statusCode)
	w.Write([]byte(text))
}

// RespondHTML 便捷方法：向客户端写入 HTML 响应。
//
// 参数:
//   - w: http.ResponseWriter
//   - statusCode: HTTP 状态码
//   - html: HTML 内容
func RespondHTML(w http.ResponseWriter, statusCode int, html string) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(statusCode)
	w.Write([]byte(html))
}

// GetQueryParam 从 HTTP 请求中获取 URL 查询参数值。
//
// 参数:
//   - r: *http.Request
//   - key: 查询参数名称
//
// 返回:
//   - string: 参数值，不存在则返回空字符串
func GetQueryParam(r *http.Request, key string) string {
	return r.URL.Query().Get(key)
}

// GetPostParam 从 HTTP 请求中解析并获取 POST 表单参数值。
//
// 参数:
//   - r: *http.Request
//   - key: POST 参数名称
//
// 返回:
//   - string: 参数值，解析失败或不存在则返回空字符串
func GetPostParam(r *http.Request, key string) string {
	r.ParseForm()
	return r.FormValue(key)
}

// ParseJSONBody 从 HTTP 请求体中解析 JSON 到目标对象。
//
// 参数:
//   - r: *http.Request
//   - target: 用于接收解析结果的对象指针
//
// 返回:
//   - error: 解析失败时返回错误
func ParseJSONBody(r *http.Request, target interface{}) error {
	return json.NewDecoder(r.Body).Decode(target)
}