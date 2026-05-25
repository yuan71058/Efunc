// 图片处理工具
// 基于 disintegration/imaging 和 skip2/go-qrcode 库，提供图片的读取/保存、缩放/裁剪/旋转、
// 效果处理（灰度/模糊/锐化/色相调整）、水印叠加、图片拼接、二维码生成等功能。
package utils

import (
	"bytes"
	"encoding/base64"
	"image"
	"image/color"
	"image/draw"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"os"
	"strings"

	"github.com/disintegration/imaging"
	"github.com/skip2/go-qrcode"
)

// ============================================================
// 图片读取与保存
// ============================================================

// Image_Read 从文件读取图片，自动识别格式（PNG/JPEG/GIF/BMP/TIFF）。
//
// 参数:
//   - filePath: 图片文件的完整路径
//
// 返回:
//   - image.Image: 解码后的图片对象
//   - error: 读取失败时返回错误
func Image_Read(filePath string) (image.Image, error) {
	return imaging.Open(filePath)
}

// Image_ReadBase64 从 Base64 编码字符串读取图片，自动识别格式。
// 支持带 data:image/xxx;base64, 前缀的字符串。
//
// 参数:
//   - b64Text: Base64 编码的图片字符串
//
// 返回:
//   - image.Image: 解码后的图片对象
//   - error: 解码失败时返回错误
func Image_ReadBase64(b64Text string) (image.Image, error) {
	text := b64Text
	if idx := strings.Index(b64Text, ","); idx > 0 {
		text = b64Text[idx+1:]
	}
	data, err := base64.StdEncoding.DecodeString(text)
	if err != nil {
		return nil, err
	}
	return Image_FromBytes(data)
}

// Image_FromBytes 从字节切片读取图片，自动识别格式。
//
// 参数:
//   - data: 图片的字节数据
//
// 返回:
//   - image.Image: 解码后的图片对象
//   - error: 解码失败时返回错误
func Image_FromBytes(data []byte) (image.Image, error) {
	reader := bytes.NewReader(data)
	img, _, err := image.Decode(reader)
	return img, err
}

// Image_ReadFromReader 从 io.Reader 读取图片，自动识别格式。
//
// 参数:
//   - r: 实现 io.Reader 接口的读取器
//
// 返回:
//   - image.Image: 解码后的图片对象
//   - string: 识别到的图片格式（如 "png"、"jpeg"）
//   - error: 解码失败时返回错误
func Image_ReadFromReader(r io.Reader) (image.Image, string, error) {
	return image.Decode(r)
}

// Image_Save 保存图片到文件，根据文件扩展名自动选择格式。
// 支持的扩展名：.png、.jpg/.jpeg、.gif、.bmp、.tiff/.tif
//
// 参数:
//   - img: 要保存的图片对象
//   - filePath: 保存路径（含扩展名）
//
// 返回:
//   - error: 保存失败时返回错误
func Image_Save(img image.Image, filePath string) error {
	return imaging.Save(img, filePath)
}

// Image_SavePNG 将图片保存为 PNG 格式。
//
// 参数:
//   - img: 要保存的图片对象
//   - filePath: 保存路径
//
// 返回:
//   - error: 保存失败时返回错误
func Image_SavePNG(img image.Image, filePath string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()
	return png.Encode(file, img)
}

// Image_SaveJPEG 将图片保存为 JPEG 格式。
//
// 参数:
//   - img: 要保存的图片对象
//   - filePath: 保存路径
//   - quality: JPEG 压缩质量（1-100），推荐 75-85
//
// 返回:
//   - error: 保存失败时返回错误
func Image_SaveJPEG(img image.Image, filePath string, quality int) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()
	if quality < 1 {
		quality = 1
	}
	if quality > 100 {
		quality = 100
	}
	return jpeg.Encode(file, img, &jpeg.Options{Quality: quality})
}

// Image_SaveGIF 将图片保存为 GIF 格式。
//
// 参数:
//   - img: 要保存的图片对象
//   - filePath: 保存路径
//
// 返回:
//   - error: 保存失败时返回错误
func Image_SaveGIF(img image.Image, filePath string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()
	return gif.Encode(file, img, nil)
}

// ============================================================
// 图片信息获取
// ============================================================

func Image_Width(img image.Image) int           { return img.Bounds().Dx() }
func Image_Height(img image.Image) int          { return img.Bounds().Dy() }
func Image_Size(img image.Image) (int, int)     { return img.Bounds().Dx(), img.Bounds().Dy() }
func Image_Bounds(img image.Image) image.Rectangle { return img.Bounds() }

// Image_GetPixel 获取指定坐标的像素颜色。
func Image_GetPixel(img image.Image, x int, y int) color.Color {
	return img.At(x, y)
}

// Image_GetPixelRGBA 获取指定坐标的 RGBA 颜色分量（0-255）。
func Image_GetPixelRGBA(img image.Image, x int, y int) (uint32, uint32, uint32, uint32) {
	r, g, b, a := img.At(x, y).RGBA()
	return r >> 8, g >> 8, b >> 8, a >> 8
}

// Image_ToBase64 将图片转换为 Base64 编码字符串。
func Image_ToBase64(img image.Image, format string) (string, error) {
	var buf bytes.Buffer
	var err error
	switch strings.ToLower(format) {
	case "jpeg", "jpg":
		err = jpeg.Encode(&buf, img, &jpeg.Options{Quality: 85})
	case "gif":
		err = gif.Encode(&buf, img, nil)
	default:
		err = png.Encode(&buf, img)
	}
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(buf.Bytes()), nil
}

// Image_ToDataURI 将图片转换为 Data URI 格式字符串。
func Image_ToDataURI(img image.Image, format string) (string, error) {
	b64, err := Image_ToBase64(img, format)
	if err != nil {
		return "", err
	}
	mimeType := "image/png"
	switch strings.ToLower(format) {
	case "jpeg", "jpg":
		mimeType = "image/jpeg"
	case "gif":
		mimeType = "image/gif"
	}
	return "data:" + mimeType + ";base64," + b64, nil
}

// Image_ToBytes 将图片编码为字节切片。
func Image_ToBytes(img image.Image, format string) ([]byte, error) {
	var buf bytes.Buffer
	var err error
	switch strings.ToLower(format) {
	case "jpeg", "jpg":
		err = jpeg.Encode(&buf, img, &jpeg.Options{Quality: 85})
	case "gif":
		err = gif.Encode(&buf, img, nil)
	default:
		err = png.Encode(&buf, img)
	}
	return buf.Bytes(), err
}

// ============================================================
// 图片变换（缩放/裁剪/旋转/翻转）
// ============================================================

func Image_Resize(img image.Image, width int, height int) image.Image {
	if width <= 0 && height <= 0 {
		return img
	}
	origW, origH := Image_Size(img)
	if width <= 0 {
		ratio := float64(height) / float64(origH)
		width = int(float64(origW) * ratio)
	}
	if height <= 0 {
		ratio := float64(width) / float64(origW)
		height = int(float64(origH) * ratio)
	}
	return imaging.Resize(img, width, height, imaging.Lanczos)
}
func Image_ResizeToWidth(img image.Image, width int) image.Image {
	return imaging.Resize(img, width, 0, imaging.Lanczos)
}
func Image_ResizeToHeight(img image.Image, height int) image.Image {
	return imaging.Resize(img, 0, height, imaging.Lanczos)
}
func Image_Thumbnail(img image.Image, width int, height int) image.Image {
	return imaging.Thumbnail(img, width, height, imaging.Lanczos)
}
func Image_Crop(img image.Image, left int, top int, right int, bottom int) image.Image {
	return imaging.Crop(img, image.Rect(left, top, right, bottom))
}
func Image_CropCenter(img image.Image, width int, height int) image.Image {
	return imaging.CropCenter(img, width, height)
}
func Image_Rotate(img image.Image, angle float64) image.Image {
	return imaging.Rotate(img, angle, color.Transparent)
}
func Image_Rotate90(img image.Image) image.Image   { return imaging.Rotate90(img) }
func Image_Rotate180(img image.Image) image.Image  { return imaging.Rotate180(img) }
func Image_Rotate270(img image.Image) image.Image  { return imaging.Rotate270(img) }
func Image_FlipH(img image.Image) image.Image      { return imaging.FlipH(img) }
func Image_FlipV(img image.Image) image.Image      { return imaging.FlipV(img) }

// ============================================================
// 图片效果
// ============================================================

func Image_Grayscale(img image.Image) image.Image                                { return imaging.Grayscale(img) }
func Image_Invert(img image.Image) image.Image                                   { return imaging.Invert(img) }
func Image_AdjustBrightness(img image.Image, brightness int) image.Image        { return imaging.AdjustBrightness(img, float64(brightness)/100.0) }
func Image_AdjustContrast(img image.Image, contrast int) image.Image            { return imaging.AdjustContrast(img, float64(contrast)/100.0) }
func Image_AdjustSaturation(img image.Image, saturation int) image.Image        { return imaging.AdjustSaturation(img, float64(saturation)/100.0) }
func Image_Blur(img image.Image, radius float64) image.Image                    { return imaging.Blur(img, radius) }
func Image_Sharpen(img image.Image, strength float64) image.Image               { return imaging.Sharpen(img, strength) }
func Image_AdjustGamma(img image.Image, gamma float64) image.Image              { return imaging.AdjustGamma(img, gamma) }

// Image_AdjustHue 调整图片色相。
func Image_AdjustHue(img image.Image, hue int) image.Image {
	bounds := img.Bounds()
	result := image.NewRGBA(bounds)
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, a := img.At(x, y).RGBA()
			h, s, l := rgbToHsl(uint8(r>>8), uint8(g>>8), uint8(b>>8))
			h += float64(hue) / 360.0
			for h < 0 {
				h += 1.0
			}
			for h > 1.0 {
				h -= 1.0
			}
			nr, ng, nb := hslToRgb(h, s, l)
			result.SetRGBA(x, y, color.RGBA{R: nr, G: ng, B: nb, A: uint8(a >> 8)})
		}
	}
	return result
}

func rgbToHsl(r, g, b uint8) (float64, float64, float64) {
	rf, gf, bf := float64(r)/255.0, float64(g)/255.0, float64(b)/255.0
	max, min := rf, rf
	if gf > max {
		max = gf
	}
	if bf > max {
		max = bf
	}
	if gf < min {
		min = gf
	}
	if bf < min {
		min = bf
	}
	l := (max + min) / 2.0
	if max == min {
		return 0, 0, l
	}
	d := max - min
	var s float64
	if l > 0.5 {
		s = d / (2.0 - max - min)
	} else {
		s = d / (max + min)
	}
	var h float64
	switch max {
	case rf:
		h = (gf - bf) / d
		if gf < bf {
			h += 6.0
		}
	case gf:
		h = (bf-rf)/d + 2.0
	case bf:
		h = (rf-gf)/d + 4.0
	}
	h /= 6.0
	return h, s, l
}

func hslToRgb(h, s, l float64) (uint8, uint8, uint8) {
	if s == 0 {
		v := uint8(l * 255.0)
		return v, v, v
	}
	var q float64
	if l < 0.5 {
		q = l * (1.0 + s)
	} else {
		q = l + s - l*s
	}
	p := 2.0*l - q
	return uint8(hueToRgb(p, q, h+1.0/3.0) * 255.0),
		uint8(hueToRgb(p, q, h) * 255.0),
		uint8(hueToRgb(p, q, h-1.0/3.0) * 255.0)
}

func hueToRgb(p, q, t float64) float64 {
	if t < 0 {
		t += 1.0
	}
	if t > 1 {
		t -= 1.0
	}
	if t < 1.0/6.0 {
		return p + (q-p)*6.0*t
	}
	if t < 1.0/2.0 {
		return q
	}
	if t < 2.0/3.0 {
		return p + (q-p)*(2.0/3.0-t)*6.0
	}
	return p
}

// ============================================================
// 图片合成与水印
// ============================================================

// Image_Watermark 在图片上添加半透明水印。
func Image_Watermark(bg image.Image, watermark image.Image, pos string, offsetX int, offsetY int, opacity float64) image.Image {
	bgW := Image_Width(bg)
	bgH := Image_Height(bg)
	wmW := Image_Width(watermark)
	wmH := Image_Height(watermark)

	watermark = Image_SetOpacity(watermark, opacity)

	var x, y int
	switch pos {
	case "左上":
		x, y = offsetX, offsetY
	case "右上":
		x, y = bgW-wmW-offsetX, offsetY
	case "左下":
		x, y = offsetX, bgH-wmH-offsetY
	case "右下":
		x, y = bgW-wmW-offsetX, bgH-wmH-offsetY
	case "居中":
		x, y = (bgW-wmW)/2+offsetX, (bgH-wmH)/2+offsetY
	default:
		x, y = offsetX, offsetY
	}

	return imaging.Overlay(bg, watermark, image.Pt(x, y), 1.0)
}

func Image_Overlay(bg image.Image, fg image.Image, x int, y int) image.Image {
	return imaging.Overlay(bg, fg, image.Pt(x, y), 1.0)
}

func Image_OverlayWithOpacity(bg image.Image, fg image.Image, x int, y int, opacity float64) image.Image {
	return imaging.Overlay(bg, fg, image.Pt(x, y), opacity)
}

// Image_ConcatHorizontal 水平拼接多张图片。
func Image_ConcatHorizontal(imgs []image.Image) image.Image {
	if len(imgs) == 0 {
		return nil
	}
	if len(imgs) == 1 {
		return imgs[0]
	}
	targetH := Image_Height(imgs[0])
	totalW := 0
	scaled := make([]image.Image, len(imgs))
	for i, img := range imgs {
		scaled[i] = Image_ResizeToHeight(img, targetH)
		totalW += Image_Width(scaled[i])
	}
	canvas := image.NewRGBA(image.Rect(0, 0, totalW, targetH))
	curX := 0
	for _, img := range scaled {
		w := Image_Width(img)
		draw.Draw(canvas, image.Rect(curX, 0, curX+w, targetH), img, image.Pt(0, 0), draw.Over)
		curX += w
	}
	return canvas
}

// Image_ConcatVertical 垂直拼接多张图片。
func Image_ConcatVertical(imgs []image.Image) image.Image {
	if len(imgs) == 0 {
		return nil
	}
	if len(imgs) == 1 {
		return imgs[0]
	}
	targetW := Image_Width(imgs[0])
	totalH := 0
	scaled := make([]image.Image, len(imgs))
	for i, img := range imgs {
		scaled[i] = Image_ResizeToWidth(img, targetW)
		totalH += Image_Height(scaled[i])
	}
	canvas := image.NewRGBA(image.Rect(0, 0, targetW, totalH))
	curY := 0
	for _, img := range scaled {
		h := Image_Height(img)
		draw.Draw(canvas, image.Rect(0, curY, targetW, curY+h), img, image.Pt(0, 0), draw.Over)
		curY += h
	}
	return canvas
}

// ============================================================
// 图片创建
// ============================================================

func Image_NewSolid(width int, height int, c color.Color) image.Image {
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	draw.Draw(img, img.Bounds(), &image.Uniform{c}, image.Pt(0, 0), draw.Src)
	return img
}
func Image_NewTransparent(width int, height int) image.Image {
	return image.NewRGBA(image.Rect(0, 0, width, height))
}

// Image_SetOpacity 设置图片整体透明度。
func Image_SetOpacity(img image.Image, opacity float64) image.Image {
	if opacity < 0 {
		opacity = 0
	}
	if opacity > 1 {
		opacity = 1
	}
	bounds := img.Bounds()
	result := image.NewRGBA(bounds)
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, a := img.At(x, y).RGBA()
			newA := uint32(float64(a) * opacity)
			result.SetRGBA(x, y, color.RGBA{
				R: uint8(r >> 8), G: uint8(g >> 8), B: uint8(b >> 8), A: uint8(newA >> 8),
			})
		}
	}
	return result
}

// ============================================================
// 二维码生成
// ============================================================

func Image_QRCodeWrite(content string, filePath string, size int) error {
	return qrcode.WriteFile(content, qrcode.Medium, size, filePath)
}

func Image_QRCodeBase64(content string) string {
	pngBytes, err := qrcode.Encode(content, qrcode.Medium, 256)
	if err == nil {
		return base64.StdEncoding.EncodeToString(pngBytes)
	}
	return ""
}

func Image_QRCodeCustom(content string, recoveryLevel int, size int) (image.Image, error) {
	var level qrcode.RecoveryLevel
	switch {
	case recoveryLevel <= 1:
		level = qrcode.Low
	case recoveryLevel == 2:
		level = qrcode.Medium
	case recoveryLevel == 3:
		level = qrcode.High
	default:
		level = qrcode.Highest
	}
	q, err := qrcode.New(content, level)
	if err != nil {
		return nil, err
	}
	q.DisableBorder = false
	return q.Image(size), nil
}

func Image_QRCodeToWriter(content string, w io.Writer, size int) error {
	q, err := qrcode.New(content, qrcode.Medium)
	if err != nil {
		return err
	}
	return png.Encode(w, q.Image(size))
}