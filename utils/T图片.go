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

// T图片_读取 从文件读取图片，自动识别格式（PNG/JPEG/GIF/BMP/TIFF）。
//
// 参数:
//   - 文件路径: 图片文件的完整路径
//
// 返回:
//   - image.Image: 解码后的图片对象
//   - error: 读取失败时返回错误
func T图片_读取(文件路径 string) (image.Image, error) {
	return imaging.Open(文件路径)
}

// T图片_读取Base64 从 Base64 编码字符串读取图片，自动识别格式。
// 支持带 data:image/xxx;base64, 前缀的字符串。
//
// 参数:
//   - base64文本: Base64 编码的图片字符串
//
// 返回:
//   - image.Image: 解码后的图片对象
//   - error: 解码失败时返回错误
func T图片_读取Base64(base64文本 string) (image.Image, error) {
	文本 := base64文本
	如果 := strings.Index(base64文本, ",")
	if 如果 > 0 {
		文本 = base64文本[如果+1:]
	}
	数据, err := base64.StdEncoding.DecodeString(文本)
	if err != nil {
		return nil, err
	}
	return T图片_从字节读取(数据)
}

// T图片_从字节读取 从字节切片读取图片，自动识别格式。
//
// 参数:
//   - 数据: 图片的字节数据
//
// 返回:
//   - image.Image: 解码后的图片对象
//   - error: 解码失败时返回错误
func T图片_从字节读取(数据 []byte) (image.Image, error) {
	reader := bytes.NewReader(数据)
	img, _, err := image.Decode(reader)
	return img, err
}

// T图片_从读取器读取 从 io.Reader 读取图片，自动识别格式。
//
// 参数:
//   - 读取器: 实现 io.Reader 接口的读取器
//
// 返回:
//   - image.Image: 解码后的图片对象
//   - string: 识别到的图片格式（如 "png"、"jpeg"）
//   - error: 解码失败时返回错误
func T图片_从读取器读取(读取器 io.Reader) (image.Image, string, error) {
	return image.Decode(读取器)
}

// T图片_保存 保存图片到文件，根据文件扩展名自动选择格式。
// 支持的扩展名：.png、.jpg/.jpeg、.gif、.bmp、.tiff/.tif
//
// 参数:
//   - 图片: 要保存的图片对象
//   - 文件路径: 保存路径（含扩展名）
//
// 返回:
//   - error: 保存失败时返回错误
func T图片_保存(图片 image.Image, 文件路径 string) error {
	return imaging.Save(图片, 文件路径)
}

// T图片_保存PNG 将图片保存为 PNG 格式。
//
// 参数:
//   - 图片: 要保存的图片对象
//   - 文件路径: 保存路径
//
// 返回:
//   - error: 保存失败时返回错误
func T图片_保存PNG(图片 image.Image, 文件路径 string) error {
	file, err := os.Create(文件路径)
	if err != nil {
		return err
	}
	defer file.Close()
	return png.Encode(file, 图片)
}

// T图片_保存JPEG 将图片保存为 JPEG 格式。
//
// 参数:
//   - 图片: 要保存的图片对象
//   - 文件路径: 保存路径
//   - 质量: JPEG 压缩质量（1-100），推荐 75-85
//
// 返回:
//   - error: 保存失败时返回错误
func T图片_保存JPEG(图片 image.Image, 文件路径 string, 质量 int) error {
	file, err := os.Create(文件路径)
	if err != nil {
		return err
	}
	defer file.Close()
	if 质量 < 1 {
		质量 = 1
	}
	if 质量 > 100 {
		质量 = 100
	}
	return jpeg.Encode(file, 图片, &jpeg.Options{Quality: 质量})
}

// T图片_保存GIF 将图片保存为 GIF 格式。
//
// 参数:
//   - 图片: 要保存的图片对象
//   - 文件路径: 保存路径
//
// 返回:
//   - error: 保存失败时返回错误
func T图片_保存GIF(图片 image.Image, 文件路径 string) error {
	file, err := os.Create(文件路径)
	if err != nil {
		return err
	}
	defer file.Close()
	return gif.Encode(file, 图片, nil)
}

// ============================================================
// 图片信息获取
// ============================================================

// T图片_取宽度 获取图片宽度（像素）。
//
// 参数:
//   - 图片: 图片对象
//
// 返回:
//   - int: 宽度像素数
func T图片_取宽度(图片 image.Image) int {
	return 图片.Bounds().Dx()
}

// T图片_取高度 获取图片高度（像素）。
//
// 参数:
//   - 图片: 图片对象
//
// 返回:
//   - int: 高度像素数
func T图片_取高度(图片 image.Image) int {
	return 图片.Bounds().Dy()
}

// T图片_取尺寸 获取图片尺寸（宽度, 高度）。
//
// 参数:
//   - 图片: 图片对象
//
// 返回:
//   - int: 宽度像素数
//   - int: 高度像素数
func T图片_取尺寸(图片 image.Image) (int, int) {
	return 图片.Bounds().Dx(), 图片.Bounds().Dy()
}

// T图片_取边界 获取图片边界矩形。
//
// 参数:
//   - 图片: 图片对象
//
// 返回:
//   - image.Rectangle: 图片边界矩形
func T图片_取边界(图片 image.Image) image.Rectangle {
	return 图片.Bounds()
}

// T图片_取像素颜色 获取指定坐标的像素颜色。
//
// 参数:
//   - 图片: 图片对象
//   - x: 像素 X 坐标（从左到右，从 0 开始）
//   - y: 像素 Y 坐标（从上到下，从 0 开始）
//
// 返回:
//   - color.Color: 像素颜色
func T图片_取像素颜色(图片 image.Image, x int, y int) color.Color {
	return 图片.At(x, y)
}

// T图片_取像素RGBA 获取指定坐标的 RGBA 颜色分量（0-255）。
//
// 参数:
//   - 图片: 图片对象
//   - x: 像素 X 坐标
//   - y: 像素 Y 坐标
//
// 返回:
//   - uint32: 红色分量（0-255）
//   - uint32: 绿色分量（0-255）
//   - uint32: 蓝色分量（0-255）
//   - uint32: 透明度分量（0-255）
func T图片_取像素RGBA(图片 image.Image, x int, y int) (uint32, uint32, uint32, uint32) {
	r, g, b, a := 图片.At(x, y).RGBA()
	return r >> 8, g >> 8, b >> 8, a >> 8
}

// T图片_转Base64 将图片转换为 Base64 编码字符串。
//
// 参数:
//   - 图片: 图片对象
//   - 格式: 图片格式，如 "png"、"jpeg"、"gif"
//
// 返回:
//   - string: Base64 编码字符串
//   - error: 编码失败时返回错误
func T图片_转Base64(图片 image.Image, 格式 string) (string, error) {
	var buf bytes.Buffer
	var err error
	switch strings.ToLower(格式) {
	case "jpeg", "jpg":
		err = jpeg.Encode(&buf, 图片, &jpeg.Options{Quality: 85})
	case "gif":
		err = gif.Encode(&buf, 图片, nil)
	default:
		err = png.Encode(&buf, 图片)
	}
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(buf.Bytes()), nil
}

// T图片_转DataURI 将图片转换为 Data URI 格式字符串。
// 可直接用于 HTML img 标签的 src 属性。
//
// 参数:
//   - 图片: 图片对象
//   - 格式: 图片格式，如 "png"、"jpeg"、"gif"
//
// 返回:
//   - string: Data URI 字符串，如 "data:image/png;base64,xxx"
//   - error: 编码失败时返回错误
func T图片_转DataURI(图片 image.Image, 格式 string) (string, error) {
	b64, err := T图片_转Base64(图片, 格式)
	if err != nil {
		return "", err
	}
	mime类型 := "image/png"
	switch strings.ToLower(格式) {
	case "jpeg", "jpg":
		mime类型 = "image/jpeg"
	case "gif":
		mime类型 = "image/gif"
	}
	return "data:" + mime类型 + ";base64," + b64, nil
}

// T图片_转字节 将图片编码为字节切片。
//
// 参数:
//   - 图片: 图片对象
//   - 格式: 图片格式，如 "png"、"jpeg"、"gif"
//
// 返回:
//   - []byte: 编码后的字节数据
//   - error: 编码失败时返回错误
func T图片_转字节(图片 image.Image, 格式 string) ([]byte, error) {
	var buf bytes.Buffer
	var err error
	switch strings.ToLower(格式) {
	case "jpeg", "jpg":
		err = jpeg.Encode(&buf, 图片, &jpeg.Options{Quality: 85})
	case "gif":
		err = gif.Encode(&buf, 图片, nil)
	default:
		err = png.Encode(&buf, 图片)
	}
	return buf.Bytes(), err
}

// ============================================================
// 图片变换（缩放/裁剪/旋转/翻转）
// ============================================================

// T图片_缩放 将图片缩放到指定尺寸。
// 使用 Lanczos 滤镜进行高质量缩放。
//
// 参数:
//   - 图片: 原始图片
//   - 宽度: 目标宽度（像素），传 0 则按高度等比缩放
//   - 高度: 目标高度（像素），传 0 则按宽度等比缩放
//
// 返回:
//   - image.Image: 缩放后的图片
func T图片_缩放(图片 image.Image, 宽度 int, 高度 int) image.Image {
	if 宽度 <= 0 && 高度 <= 0 {
		return 图片
	}
	原宽, 原高 := T图片_取尺寸(图片)
	if 宽度 <= 0 {
		比例 := float64(高度) / float64(原高)
		宽度 = int(float64(原宽) * 比例)
	}
	if 高度 <= 0 {
		比例 := float64(宽度) / float64(原宽)
		高度 = int(float64(原高) * 比例)
	}
	return imaging.Resize(图片, 宽度, 高度, imaging.Lanczos)
}

// T图片_缩放到宽度 将图片等比缩放到指定宽度。
//
// 参数:
//   - 图片: 原始图片
//   - 宽度: 目标宽度（像素）
//
// 返回:
//   - image.Image: 缩放后的图片
func T图片_缩放到宽度(图片 image.Image, 宽度 int) image.Image {
	return imaging.Resize(图片, 宽度, 0, imaging.Lanczos)
}

// T图片_缩放到高度 将图片等比缩放到指定高度。
//
// 参数:
//   - 图片: 原始图片
//   - 高度: 目标高度（像素）
//
// 返回:
//   - image.Image: 缩放后的图片
func T图片_缩放到高度(图片 image.Image, 高度 int) image.Image {
	return imaging.Resize(图片, 0, 高度, imaging.Lanczos)
}

// T图片_缩略图 生成缩略图，将图片等比缩放并裁剪到指定尺寸。
// 与 T图片_缩放 不同，此函数会裁剪多余部分以确保输出尺寸精确。
//
// 参数:
//   - 图片: 原始图片
//   - 宽度: 缩略图宽度（像素）
//   - 高度: 缩略图高度（像素）
//
// 返回:
//   - image.Image: 缩略图
func T图片_缩略图(图片 image.Image, 宽度 int, 高度 int) image.Image {
	return imaging.Thumbnail(图片, 宽度, 高度, imaging.Lanczos)
}

// T图片_裁剪 裁剪图片到指定矩形区域。
//
// 参数:
//   - 图片: 原始图片
//   - 左: 裁剪区域左边界 X 坐标
//   - 上: 裁剪区域上边界 Y 坐标
//   - 右: 裁剪区域右边界 X 坐标
//   - 下: 裁剪区域下边界 Y 坐标
//
// 返回:
//   - image.Image: 裁剪后的图片
func T图片_裁剪(图片 image.Image, 左 int, 上 int, 右 int, 下 int) image.Image {
	return imaging.Crop(图片, image.Rect(左, 上, 右, 下))
}

// T图片_居中裁剪 从图片中心裁剪出指定尺寸的区域。
//
// 参数:
//   - 图片: 原始图片
//   - 宽度: 裁剪宽度（像素）
//   - 高度: 裁剪高度（像素）
//
// 返回:
//   - image.Image: 裁剪后的图片
func T图片_居中裁剪(图片 image.Image, 宽度 int, 高度 int) image.Image {
	return imaging.CropCenter(图片, 宽度, 高度)
}

// T图片_旋转 将图片顺时针旋转指定角度。
//
// 参数:
//   - 图片: 原始图片
//   - 角度: 顺时针旋转角度（如 90、180、270）
//
// 返回:
//   - image.Image: 旋转后的图片
func T图片_旋转(图片 image.Image, 角度 float64) image.Image {
	return imaging.Rotate(图片, 角度, color.Transparent)
}

// T图片_旋转90 顺时针旋转 90 度。
//
// 参数:
//   - 图片: 原始图片
//
// 返回:
//   - image.Image: 旋转后的图片
func T图片_旋转90(图片 image.Image) image.Image {
	return imaging.Rotate90(图片)
}

// T图片_旋转180 旋转 180 度。
//
// 参数:
//   - 图片: 原始图片
//
// 返回:
//   - image.Image: 旋转后的图片
func T图片_旋转180(图片 image.Image) image.Image {
	return imaging.Rotate180(图片)
}

// T图片_旋转270 顺时针旋转 270 度（即逆时针 90 度）。
//
// 参数:
//   - 图片: 原始图片
//
// 返回:
//   - image.Image: 旋转后的图片
func T图片_旋转270(图片 image.Image) image.Image {
	return imaging.Rotate270(图片)
}

// T图片_水平翻转 将图片左右镜像翻转。
//
// 参数:
//   - 图片: 原始图片
//
// 返回:
//   - image.Image: 翻转后的图片
func T图片_水平翻转(图片 image.Image) image.Image {
	return imaging.FlipH(图片)
}

// T图片_垂直翻转 将图片上下镜像翻转。
//
// 参数:
//   - 图片: 原始图片
//
// 返回:
//   - image.Image: 翻转后的图片
func T图片_垂直翻转(图片 image.Image) image.Image {
	return imaging.FlipV(图片)
}

// ============================================================
// 图片效果（灰度/亮度/对比度/模糊/锐化）
// ============================================================

// T图片_灰度化 将图片转换为灰度图。
//
// 参数:
//   - 图片: 原始图片
//
// 返回:
//   - image.Image: 灰度图
func T图片_灰度化(图片 image.Image) image.Image {
	return imaging.Grayscale(图片)
}

// T图片_反色 将图片颜色取反（底片效果）。
//
// 参数:
//   - 图片: 原始图片
//
// 返回:
//   - image.Image: 反色后的图片
func T图片_反色(图片 image.Image) image.Image {
	return imaging.Invert(图片)
}

// T图片_调整亮度 调整图片亮度。
//
// 参数:
//   - 图片: 原始图片
//   - 亮度: 亮度调整值，范围 -100 到 100；负值变暗，正值变亮
//
// 返回:
//   - image.Image: 调整后的图片
func T图片_调整亮度(图片 image.Image, 亮度 int) image.Image {
	return imaging.AdjustBrightness(图片, float64(亮度)/100.0)
}

// T图片_调整对比度 调整图片对比度。
//
// 参数:
//   - 图片: 原始图片
//   - 对比度: 对比度调整值，范围 -100 到 100；负值降低，正值增强
//
// 返回:
//   - image.Image: 调整后的图片
func T图片_调整对比度(图片 image.Image, 对比度 int) image.Image {
	return imaging.AdjustContrast(图片, float64(对比度)/100.0)
}

// T图片_调整饱和度 调整图片饱和度。
//
// 参数:
//   - 图片: 原始图片
//   - 饱和度: 饱和度调整值，范围 -100 到 100；负值降低，正值增强
//
// 返回:
//   - image.Image: 调整后的图片
func T图片_调整饱和度(图片 image.Image, 饱和度 int) image.Image {
	return imaging.AdjustSaturation(图片, float64(饱和度)/100.0)
}

// T图片_调整色相 调整图片色相。
// 通过 HSL 色彩空间进行色相偏移。
//
// 参数:
//   - 图片: 原始图片
//   - 色相: 色相偏移值，范围 -180 到 180
//
// 返回:
//   - image.Image: 调整后的图片
func T图片_调整色相(图片 image.Image, 色相 int) image.Image {
	bounds := 图片.Bounds()
	结果 := image.NewRGBA(bounds)
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, a := 图片.At(x, y).RGBA()
			h, s, l := rgbToHsl(uint8(r>>8), uint8(g>>8), uint8(b>>8))
			h += float64(色相) / 360.0
			for h < 0 {
				h += 1.0
			}
			for h > 1.0 {
				h -= 1.0
			}
			nr, ng, nb := hslToRgb(h, s, l)
			结果.SetRGBA(x, y, color.RGBA{R: nr, G: ng, B: nb, A: uint8(a >> 8)})
		}
	}
	return 结果
}

// rgbToHsl 将 RGB 颜色转换为 HSL。
func rgbToHsl(r, g, b uint8) (float64, float64, float64) {
	rf := float64(r) / 255.0
	gf := float64(g) / 255.0
	bf := float64(b) / 255.0
	max := rf
	if gf > max {
		max = gf
	}
	if bf > max {
		max = bf
	}
	min := rf
	if gf < min {
		min = gf
	}
	if bf < min {
		min = bf
	}
	l := (max + min) / 2.0
	var h, s float64
	if max == min {
		return 0, 0, l
	}
	d := max - min
	if l > 0.5 {
		s = d / (2.0 - max - min)
	} else {
		s = d / (max + min)
	}
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

// hslToRgb 将 HSL 颜色转换为 RGB。
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
	r := hueToRgb(p, q, h+1.0/3.0)
	g := hueToRgb(p, q, h)
	b := hueToRgb(p, q, h-1.0/3.0)
	return uint8(r * 255.0), uint8(g * 255.0), uint8(b * 255.0)
}

// hueToRgb 色相到 RGB 分量转换辅助函数。
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

// T图片_模糊 对图片进行高斯模糊处理。
//
// 参数:
//   - 图片: 原始图片
//   - 半径: 模糊半径（像素），值越大越模糊；推荐 1-20
//
// 返回:
//   - image.Image: 模糊后的图片
func T图片_模糊(图片 image.Image, 半径 float64) image.Image {
	return imaging.Blur(图片, 半径)
}

// T图片_锐化 对图片进行锐化处理。
//
// 参数:
//   - 图片: 原始图片
//   - 强度: 锐化强度，范围 0-1；值越大锐化越明显
//
// 返回:
//   - image.Image: 锐化后的图片
func T图片_锐化(图片 image.Image, 强度 float64) image.Image {
	return imaging.Sharpen(图片, 强度)
}

// T图片_伽马校正 对图片进行伽马校正。
//
// 参数:
//   - 图片: 原始图片
//   - 伽马值: 伽马校正值；1.0 为原始值，<1 变亮，>1 变暗
//
// 返回:
//   - image.Image: 校正后的图片
func T图片_伽马校正(图片 image.Image, 伽马值 float64) image.Image {
	return imaging.AdjustGamma(图片, 伽马值)
}

// ============================================================
// 图片合成与水印
// ============================================================

// T图片_添加水印 在图片上添加半透明水印。
// 水印图片将被缩放到指定尺寸后叠加到目标位置。
//
// 参数:
//   - 底图: 底层图片
//   - 水印图: 水印图片
//   - 位置: 水印位置（"左上"/"右上"/"左下"/"右下"/"居中"）
//   - 偏移X: 水印相对位置的 X 偏移（像素）
//   - 偏移Y: 水印相对位置的 Y 偏移（像素）
//   - 透明度: 水印透明度（0-1），0 完全透明，1 完全不透明
//
// 返回:
//   - image.Image: 添加水印后的图片
func T图片_添加水印(底图 image.Image, 水印图 image.Image, 位置 string, 偏移X int, 偏移Y int, 透明度 float64) image.Image {
	底宽 := T图片_取宽度(底图)
	底高 := T图片_取高度(底图)
	印宽 := T图片_取宽度(水印图)
	印高 := T图片_取高度(水印图)

	水印图 = T图片_设置透明度(水印图, 透明度)

	var x, y int
	switch 位置 {
	case "左上":
		x, y = 偏移X, 偏移Y
	case "右上":
		x, y = 底宽-印宽-偏移X, 偏移Y
	case "左下":
		x, y = 偏移X, 底高-印高-偏移Y
	case "右下":
		x, y = 底宽-印宽-偏移X, 底高-印高-偏移Y
	case "居中":
		x, y = (底宽-印宽)/2+偏移X, (底高-印高)/2+偏移Y
	default:
		x, y = 偏移X, 偏移Y
	}

	return imaging.Overlay(底图, 水印图, image.Pt(x, y), 1.0)
}

// T图片_叠加 将上层图片叠加到底层图片的指定位置。
//
// 参数:
//   - 底图: 底层图片
//   - 上层图: 上层图片
//   - x: 叠加位置 X 坐标
//   - y: 叠加位置 Y 坐标
//
// 返回:
//   - image.Image: 叠加后的图片
func T图片_叠加(底图 image.Image, 上层图 image.Image, x int, y int) image.Image {
	return imaging.Overlay(底图, 上层图, image.Pt(x, y), 1.0)
}

// T图片_叠加带透明度 将上层图片以指定透明度叠加到底层图片。
//
// 参数:
//   - 底图: 底层图片
//   - 上层图: 上层图片
//   - x: 叠加位置 X 坐标
//   - y: 叠加位置 Y 坐标
//   - 透明度: 叠加透明度（0-1）
//
// 返回:
//   - image.Image: 叠加后的图片
func T图片_叠加带透明度(底图 image.Image, 上层图 image.Image, x int, y int, 透明度 float64) image.Image {
	return imaging.Overlay(底图, 上层图, image.Pt(x, y), 透明度)
}

// T图片_拼接水平 将多张图片水平拼接成一张。
// 所有图片按第一张的高度等比缩放后拼接。
//
// 参数:
//   - 图片列表: 要拼接的图片切片
//
// 返回:
//   - image.Image: 拼接后的图片
func T图片_拼接水平(图片列表 []image.Image) image.Image {
	if len(图片列表) == 0 {
		return nil
	}
	if len(图片列表) == 1 {
		return 图片列表[0]
	}

	目标高度 := T图片_取高度(图片列表[0])
	总宽度 := 0
	缩放列表 := make([]image.Image, len(图片列表))
	for i, img := range 图片列表 {
		缩放图 := T图片_缩放到高度(img, 目标高度)
		缩放列表[i] = 缩放图
		总宽度 += T图片_取宽度(缩放图)
	}

	画布 := image.NewRGBA(image.Rect(0, 0, 总宽度, 目标高度))
	当前X := 0
	for _, img := range 缩放列表 {
		draw.Draw(画布, image.Rect(当前X, 0, 当前X+T图片_取宽度(img), 目标高度), img, image.Pt(0, 0), draw.Over)
		当前X += T图片_取宽度(img)
	}
	return 画布
}

// T图片_拼接垂直 将多张图片垂直拼接成一张。
// 所有图片按第一张的宽度等比缩放后拼接。
//
// 参数:
//   - 图片列表: 要拼接的图片切片
//
// 返回:
//   - image.Image: 拼接后的图片
func T图片_拼接垂直(图片列表 []image.Image) image.Image {
	if len(图片列表) == 0 {
		return nil
	}
	if len(图片列表) == 1 {
		return 图片列表[0]
	}

	目标宽度 := T图片_取宽度(图片列表[0])
	总高度 := 0
	缩放列表 := make([]image.Image, len(图片列表))
	for i, img := range 图片列表 {
		缩放图 := T图片_缩放到宽度(img, 目标宽度)
		缩放列表[i] = 缩放图
		总高度 += T图片_取高度(缩放图)
	}

	画布 := image.NewRGBA(image.Rect(0, 0, 目标宽度, 总高度))
	当前Y := 0
	for _, img := range 缩放列表 {
		draw.Draw(画布, image.Rect(0, 当前Y, 目标宽度, 当前Y+T图片_取高度(img)), img, image.Pt(0, 0), draw.Over)
		当前Y += T图片_取高度(img)
	}
	return 画布
}

// ============================================================
// 图片创建
// ============================================================

// T图片_创建纯色图 创建指定尺寸和颜色的纯色图片。
//
// 参数:
//   - 宽度: 图片宽度（像素）
//   - 高度: 图片高度（像素）
//   - 颜色值: 填充颜色
//
// 返回:
//   - image.Image: 纯色图片
func T图片_创建纯色图(宽度 int, 高度 int, 颜色值 color.Color) image.Image {
	img := image.NewRGBA(image.Rect(0, 0, 宽度, 高度))
	draw.Draw(img, img.Bounds(), &image.Uniform{颜色值}, image.Pt(0, 0), draw.Src)
	return img
}

// T图片_创建透明图 创建指定尺寸的透明图片。
//
// 参数:
//   - 宽度: 图片宽度（像素）
//   - 高度: 图片高度（像素）
//
// 返回:
//   - image.Image: 透明图片
func T图片_创建透明图(宽度 int, 高度 int) image.Image {
	return image.NewRGBA(image.Rect(0, 0, 宽度, 高度))
}

// T图片_设置透明度 设置图片整体透明度。
//
// 参数:
//   - 图片: 原始图片
//   - 透明度: 透明度（0-1），0 完全透明，1 完全不透明
//
// 返回:
//   - image.Image: 调整后的图片
func T图片_设置透明度(图片 image.Image, 透明度 float64) image.Image {
	if 透明度 < 0 {
		透明度 = 0
	}
	if 透明度 > 1 {
		透明度 = 1
	}
	bounds := 图片.Bounds()
	结果 := image.NewRGBA(bounds)
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, a := 图片.At(x, y).RGBA()
			newA := uint32(float64(a) * 透明度)
			结果.SetRGBA(x, y, color.RGBA{
				R: uint8(r >> 8),
				G: uint8(g >> 8),
				B: uint8(b >> 8),
				A: uint8(newA >> 8),
			})
		}
	}
	return 结果
}

// ============================================================
// 二维码生成
// ============================================================

// T图片_生成二维码 生成二维码图片并保存到文件。
//
// 参数:
//   - 内容: 二维码编码内容
//   - 文件路径: 保存路径（PNG 格式）
//   - 尺寸: 二维码图片尺寸（像素），推荐 256
//
// 返回:
//   - error: 生成失败时返回错误
func T图片_生成二维码(内容 string, 文件路径 string, 尺寸 int) error {
	return qrcode.WriteFile(内容, qrcode.Medium, 尺寸, 文件路径)
}

// T图片_生成二维码base64 生成指定内容的二维码图片，并返回 Base64 编码字符串。
// 二维码大小为 256x256 像素，容错等级为 Medium。
// 返回的 Base64 字符串可直接用于 HTML img 标签的 src 属性。
//
// 参数:
//   - 内容: 二维码中编码的内容文本
//
// 返回:
//   - string: Base64 编码的 PNG 图片字符串；生成失败返回空串
func T图片_生成二维码base64(内容 string) string {
	局_二维码base64 := ""
	png, err := qrcode.Encode(内容, qrcode.Medium, 256)
	if err == nil {
		局_二维码base64 = base64.StdEncoding.EncodeToString(png)
	}
	return 局_二维码base64
}

// T图片_生成二维码自定义 生成自定义参数的二维码图片对象。
//
// 参数:
//   - 内容: 二维码编码内容
//   - 容错等级: 容错级别（1=Low, 2=Medium, 3=Quartile, 4=High）
//   - 尺寸: 二维码图片尺寸（像素）
//
// 返回:
//   - image.Image: 二维码图片对象
//   - error: 生成失败时返回错误
func T图片_生成二维码自定义(内容 string, 容错等级 int, 尺寸 int) (image.Image, error) {
	var level qrcode.RecoveryLevel
	switch {
	case 容错等级 <= 1:
		level = qrcode.Low
	case 容错等级 == 2:
		level = qrcode.Medium
	case 容错等级 == 3:
		level = qrcode.High
	default:
		level = qrcode.Highest
	}
	q, err := qrcode.New(内容, level)
	if err != nil {
		return nil, err
	}
	q.DisableBorder = false
	return q.Image(尺寸), nil
}

// T图片_生成二维码到写入器 将二维码图片写入 io.Writer。
//
// 参数:
//   - 内容: 二维码编码内容
//   - 写入器: 实现 io.Writer 接口的对象
//   - 尺寸: 二维码图片尺寸（像素）
//
// 返回:
//   - error: 生成失败时返回错误
func T图片_生成二维码到写入器(内容 string, 写入器 io.Writer, 尺寸 int) error {
	q, err := qrcode.New(内容, qrcode.Medium)
	if err != nil {
		return err
	}
	return png.Encode(写入器, q.Image(尺寸))
}
