package utils

import (
	"crypto/tls"
	"fmt"
	"net/smtp"
	"strings"

	"github.com/jordan-wright/email"
)

// E邮件_发送 发送邮件，支持 HTML 正文和附件。
// 使用 SMTP 协议发送，支持 TLS 加密连接。
//
// 参数:
//   - 服务器地址: SMTP 服务器地址，如 "smtp.qq.com:587"
//   - 发件人邮箱: 发件人邮箱地址
//   - 密码: 发件人邮箱的授权码（非登录密码）
//   - 收件人: 收件人邮箱地址，多个用逗号分隔
//   - 主题: 邮件主题
//   - 正文: 邮件正文（纯文本）
//   - HTML正文: HTML 格式的邮件正文，为空时使用纯文本正文
//   - 附件路径: 附件文件路径列表，无附件传 nil
//
// 返回:
//   - error: 发送失败时返回错误
func E邮件_发送(服务器地址 string, 发件人邮箱 string, 密码 string, 收件人 string, 主题 string, 正文 string, HTML正文 string, 附件路径 []string) error {
	e := email.NewEmail()
	e.From = 发件人邮箱
	e.To = strings.Split(收件人, ",")
	e.Subject = 主题

	if HTML正文 != "" {
		e.HTML = []byte(HTML正文)
	} else {
		e.Text = []byte(正文)
	}

	for _, 路径 := range 附件路径 {
		if 路径 != "" {
			_, 错误 := e.AttachFile(路径)
			if 错误 != nil {
				return fmt.Errorf("添加附件失败: %v", 错误)
			}
		}
	}

	主机 := strings.Split(服务器地址, ":")[0]
	认证 := smtp.PlainAuth("", 发件人邮箱, 密码, 主机)

	return e.Send(服务器地址, 认证)
}

// E邮件_发送TLS 使用 TLS 加密连接发送邮件。
// 适用于 465 端口的 SMTP 服务器。
//
// 参数:
//   - 服务器地址: SMTP 服务器地址，如 "smtp.qq.com:465"
//   - 发件人邮箱: 发件人邮箱地址
//   - 密码: 发件人邮箱的授权码
//   - 收件人: 收件人邮箱地址，多个用逗号分隔
//   - 主题: 邮件主题
//   - 正文: 邮件正文（纯文本）
//   - HTML正文: HTML 格式的邮件正文
//   - 附件路径: 附件文件路径列表
//
// 返回:
//   - error: 发送失败时返回错误
func E邮件_发送TLS(服务器地址 string, 发件人邮箱 string, 密码 string, 收件人 string, 主题 string, 正文 string, HTML正文 string, 附件路径 []string) error {
	e := email.NewEmail()
	e.From = 发件人邮箱
	e.To = strings.Split(收件人, ",")
	e.Subject = 主题

	if HTML正文 != "" {
		e.HTML = []byte(HTML正文)
	} else {
		e.Text = []byte(正文)
	}

	for _, 路径 := range 附件路径 {
		if 路径 != "" {
			_, 错误 := e.AttachFile(路径)
			if 错误 != nil {
				return fmt.Errorf("添加附件失败: %v", 错误)
			}
		}
	}

	主机 := strings.Split(服务器地址, ":")[0]
	认证 := smtp.PlainAuth("", 发件人邮箱, 密码, 主机)

	tls配置 := &tls.Config{
		ServerName: 主机,
	}

	return e.SendWithTLS(服务器地址, 认证, tls配置)
}

// E邮件_发送简单邮件 发送一封简单的纯文本邮件，无附件。
//
// 参数:
//   - 服务器地址: SMTP 服务器地址，如 "smtp.qq.com:587"
//   - 发件人邮箱: 发件人邮箱地址
//   - 密码: 发件人邮箱的授权码
//   - 收件人: 收件人邮箱地址
//   - 主题: 邮件主题
//   - 正文: 邮件正文
//
// 返回:
//   - error: 发送失败时返回错误
func E邮件_发送简单邮件(服务器地址 string, 发件人邮箱 string, 密码 string, 收件人 string, 主题 string, 正文 string) error {
	return E邮件_发送(服务器地址, 发件人邮箱, 密码, 收件人, 主题, 正文, "", nil)
}
