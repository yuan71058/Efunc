// 邮件发送工具
// 基于 jordan-wright/email 库，支持 SMTP 协议发送邮件。
// 支持 HTML 正文、附件、TLS 加密连接等功能。
package utils

import (
	"crypto/tls"
	"fmt"
	"net/smtp"
	"strings"

	"github.com/jordan-wright/email"
)

// Email_Send 发送邮件，支持 HTML 正文和附件。
// 使用 SMTP 协议发送，支持 TLS 加密连接。
//
// 参数:
//   - serverAddr: SMTP 服务器地址，如 "smtp.qq.com:587"
//   - from: 发件人邮箱地址
//   - password: 发件人邮箱的授权码（非登录密码）
//   - to: 收件人邮箱地址，多个用逗号分隔
//   - subject: 邮件主题
//   - textBody: 邮件正文（纯文本）
//   - htmlBody: HTML 格式的邮件正文，为空时使用纯文本正文
//   - attachments: 附件文件路径列表，无附件传 nil
//
// 返回:
//   - error: 发送失败时返回错误
func Email_Send(serverAddr string, from string, password string, to string, subject string, textBody string, htmlBody string, attachments []string) error {
	e := email.NewEmail()
	e.From = from
	e.To = strings.Split(to, ",")
	e.Subject = subject

	if htmlBody != "" {
		e.HTML = []byte(htmlBody)
	} else {
		e.Text = []byte(textBody)
	}

	for _, path := range attachments {
		if path != "" {
			_, err := e.AttachFile(path)
			if err != nil {
				return fmt.Errorf("添加附件失败: %v", err)
			}
		}
	}

	host := strings.Split(serverAddr, ":")[0]
	auth := smtp.PlainAuth("", from, password, host)

	return e.Send(serverAddr, auth)
}

// Email_SendTLS 使用 TLS 加密连接发送邮件。适用于 465 端口的 SMTP 服务器。
func Email_SendTLS(serverAddr string, from string, password string, to string, subject string, textBody string, htmlBody string, attachments []string) error {
	e := email.NewEmail()
	e.From = from
	e.To = strings.Split(to, ",")
	e.Subject = subject

	if htmlBody != "" {
		e.HTML = []byte(htmlBody)
	} else {
		e.Text = []byte(textBody)
	}

	for _, path := range attachments {
		if path != "" {
			_, err := e.AttachFile(path)
			if err != nil {
				return fmt.Errorf("添加附件失败: %v", err)
			}
		}
	}

	host := strings.Split(serverAddr, ":")[0]
	auth := smtp.PlainAuth("", from, password, host)

	tlsConfig := &tls.Config{
		ServerName: host,
	}

	return e.SendWithTLS(serverAddr, auth, tlsConfig)
}

// Email_SendSimple 发送一封简单的纯文本邮件，无附件。
func Email_SendSimple(serverAddr string, from string, password string, to string, subject string, textBody string) error {
	return Email_Send(serverAddr, from, password, to, subject, textBody, "", nil)
}