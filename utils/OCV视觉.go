//go:build opencv

package utils

import (
	"fmt"
	"image"
	"image/color"

	"gocv.io/x/gocv"
)

// ============================================================
// OpenCV 核心操作
// ============================================================

// OCV_取版本 获取 OpenCV 版本号。
//
// 返回:
//   - string: OpenCV 版本号（如 "4.9.0"）
func OCV_取版本() string {
	return gocv.Version()
}

// OCV_取CUDA设备数 获取可用 CUDA 设备数量。
//
// 返回:
//   - int: CUDA 设备数量
func OCV_取CUDA设备数() int {
	return gocv.GetCudaEnabledDeviceCount()
}

// ============================================================
// 图像读取与保存
// ============================================================

// OCV_读取图片 从文件读取图片为 Mat 对象。
//
// 参数:
//   - 文件路径: 图片文件路径
//
// 返回:
//   - gocv.Mat: OpenCV Mat 对象
//   - error: 读取失败时返回错误
func OCV_读取图片(文件路径 string) (gocv.Mat, error) {
	return gocv.IMRead(文件路径, gocv.IMReadColor)
}

// OCV_读取图片灰度 从文件读取图片为灰度 Mat 对象。
//
// 参数:
//   - 文件路径: 图片文件路径
//
// 返回:
//   - gocv.Mat: 灰度 Mat 对象
//   - error: 读取失败时返回错误
func OCV_读取图片灰度(文件路径 string) (gocv.Mat, error) {
	return gocv.IMRead(文件路径, gocv.IMReadGrayScale)
}

// OCV_读取图片原色 从文件读取图片（保留原始通道和深度）。
//
// 参数:
//   - 文件路径: 图片文件路径
//
// 返回:
//   - gocv.Mat: 原始 Mat 对象
//   - error: 读取失败时返回错误
func OCV_读取图片原色(文件路径 string) (gocv.Mat, error) {
	return gocv.IMRead(文件路径, gocv.IMReadUnchanged)
}

// OCV_保存图片 将 Mat 对象保存为图片文件。
//
// 参数:
//   - 矩阵: OpenCV Mat 对象
//   - 文件路径: 保存路径（按扩展名自动选格式）
//
// 返回:
//   - bool: 保存成功返回 true
func OCV_保存图片(矩阵 gocv.Mat, 文件路径 string) bool {
	return gocv.IMWrite(文件路径, 矩阵)
}

// OCV_从字节读取 从字节切片读取图片为 Mat 对象。
//
// 参数:
//   - 数据: 图片的字节数据
//
// 返回:
//   - gocv.Mat: Mat 对象
//   - error: 解码失败时返回错误
func OCV_从字节读取(数据 []byte) (gocv.Mat, error) {
	return gocv.IMDecode(数据, gocv.IMReadColor)
}

// OCV_从字节读取灰度 从字节切片读取图片为灰度 Mat 对象。
//
// 参数:
//   - 数据: 图片的字节数据
//
// 返回:
//   - gocv.Mat: 灰度 Mat 对象
//   - error: 解码失败时返回错误
func OCV_从字节读取灰度(数据 []byte) (gocv.Mat, error) {
	return gocv.IMDecode(数据, gocv.IMReadGrayScale)
}

// OCV_到字节 将 Mat 对象编码为指定格式的字节切片。
//
// 参数:
//   - 矩阵: OpenCV Mat 对象
//   - 扩展名: 图片格式扩展名（如 ".png"、".jpg"）
//
// 返回:
//   - []byte: 编码后的字节数据
//   - error: 编码失败时返回错误
func OCV_到字节(矩阵 gocv.Mat, 扩展名 string) ([]byte, error) {
	buf, err := gocv.IMEncode(扩展名, 矩阵)
	if err != nil {
		return nil, err
	}
	return buf.GetBytes(), nil
}

// ============================================================
// Mat 信息与属性
// ============================================================

// OCV_取宽度 获取 Mat 的宽度（列数）。
//
// 参数:
//   - 矩阵: OpenCV Mat 对象
//
// 返回:
//   - int: 宽度像素数
func OCV_取宽度(矩阵 gocv.Mat) int {
	return 矩阵.Cols()
}

// OCV_取高度 获取 Mat 的高度（行数）。
//
// 参数:
//   - 矩阵: OpenCV Mat 对象
//
// 返回:
//   - int: 高度像素数
func OCV_取高度(矩阵 gocv.Mat) int {
	return 矩阵.Rows()
}

// OCV_取通道数 获取 Mat 的通道数。
//
// 参数:
//   - 矩阵: OpenCV Mat 对象
//
// 返回:
//   - int: 通道数（1=灰度, 3=BGR, 4=BGRA）
func OCV_取通道数(矩阵 gocv.Mat) int {
	return 矩阵.Channels()
}

// OCV_取类型 获取 Mat 的数据类型。
//
// 参数:
//   - 矩阵: OpenCV Mat 对象
//
// 返回:
//   - gocv.MatType: Mat 类型
func OCV_取类型(矩阵 gocv.Mat) gocv.MatType {
	return 矩阵.Type()
}

// OCV_取像素数 获取 Mat 的总像素数。
//
// 参数:
//   - 矩阵: OpenCV Mat 对象
//
// 返回:
//   - int: 总像素数
func OCV_取像素数(矩阵 gocv.Mat) int {
	return 矩阵.Total()
}

// OCV_是否为空 判断 Mat 是否为空。
//
// 参数:
//   - 矩阵: OpenCV Mat 对象
//
// 返回:
//   - bool: 为空返回 true
func OCV_是否为空(矩阵 gocv.Mat) bool {
	return 矩阵.Empty()
}

// OCV_取尺寸 获取 Mat 的尺寸（宽, 高）。
//
// 参数:
//   - 矩阵: OpenCV Mat 对象
//
// 返回:
//   - int: 宽度
//   - int: 高度
func OCV_取尺寸(矩阵 gocv.Mat) (int, int) {
	return 矩阵.Cols(), 矩阵.Rows()
}

// ============================================================
// 颜色空间转换
// ============================================================

// OCV_BGR转灰度 将 BGR 彩色图转换为灰度图。
//
// 参数:
//   - 矩阵: BGR 格式的 Mat 对象
//
// 返回:
//   - gocv.Mat: 灰度 Mat 对象
func OCV_BGR转灰度(矩阵 gocv.Mat) gocv.Mat {
	return gocv.CvtColor(矩阵, gocv.ColorBGRToGray)
}

// OCV_灰度转BGR 将灰度图转换为 BGR 彩色图。
//
// 参数:
//   - 矩阵: 灰度 Mat 对象
//
// 返回:
//   - gocv.Mat: BGR Mat 对象
func OCV_灰度转BGR(矩阵 gocv.Mat) gocv.Mat {
	return gocv.CvtColor(矩阵, gocv.ColorGrayToBGR)
}

// OCV_BGR转HSV 将 BGR 彩色图转换为 HSV 颜色空间。
//
// 参数:
//   - 矩阵: BGR 格式的 Mat 对象
//
// 返回:
//   - gocv.Mat: HSV Mat 对象
func OCV_BGR转HSV(矩阵 gocv.Mat) gocv.Mat {
	return gocv.CvtColor(矩阵, gocv.ColorBGRToHSV)
}

// OCV_HSV转BGR 将 HSV 颜色空间转换为 BGR 彩色图。
//
// 参数:
//   - 矩阵: HSV 格式的 Mat 对象
//
// 返回:
//   - gocv.Mat: BGR Mat 对象
func OCV_HSV转BGR(矩阵 gocv.Mat) gocv.Mat {
	return gocv.CvtColor(矩阵, gocv.ColorHSVToBGR)
}

// OCV_BGR转RGB 将 BGR 格式转换为 RGB 格式。
//
// 参数:
//   - 矩阵: BGR 格式的 Mat 对象
//
// 返回:
//   - gocv.Mat: RGB Mat 对象
func OCV_BGR转RGB(矩阵 gocv.Mat) gocv.Mat {
	return gocv.CvtColor(矩阵, gocv.ColorBGRToRGB)
}

// OCV_RGB转BGR 将 RGB 格式转换为 BGR 格式。
//
// 参数:
//   - 矩阵: RGB 格式的 Mat 对象
//
// 返回:
//   - gocv.Mat: BGR Mat 对象
func OCV_RGB转BGR(矩阵 gocv.Mat) gocv.Mat {
	return gocv.CvtColor(矩阵, gocv.ColorRGBToBGR)
}

// OCV_BGR转Lab 将 BGR 格式转换为 Lab 颜色空间。
//
// 参数:
//   - 矩阵: BGR 格式的 Mat 对象
//
// 返回:
//   - gocv.Mat: Lab Mat 对象
func OCV_BGR转Lab(矩阵 gocv.Mat) gocv.Mat {
	return gocv.CvtColor(矩阵, gocv.ColorBGRToLab)
}

// OCV_BGR转YUV 将 BGR 格式转换为 YUV 颜色空间。
//
// 参数:
//   - 矩阵: BGR 格式的 Mat 对象
//
// 返回:
//   - gocv.Mat: YUV Mat 对象
func OCV_BGR转YUV(矩阵 gocv.Mat) gocv.Mat {
	return gocv.CvtColor(矩阵, gocv.ColorBGRToYUV)
}

// ============================================================
// 图像变换
// ============================================================

// OCV_缩放 将图片缩放到指定尺寸。
//
// 参数:
//   - 矩阵: 原始 Mat 对象
//   - 宽度: 目标宽度
//   - 高度: 目标高度
//
// 返回:
//   - gocv.Mat: 缩放后的 Mat 对象
func OCV_缩放(矩阵 gocv.Mat, 宽度 int, 高度 int) gocv.Mat {
	return gocv.Resize(矩阵, image.Pt(宽度, 高度), 0, 0, gocv.InterpolationLinear)
}

// OCV_缩放按比例 按比例缩放图片。
//
// 参数:
//   - 矩阵: 原始 Mat 对象
//   - 水平比例: 水平缩放比例（如 0.5 为缩小一半）
//   - 垂直比例: 垂直缩放比例
//
// 返回:
//   - gocv.Mat: 缩放后的 Mat 对象
func OCV_缩放按比例(矩阵 gocv.Mat, 水平比例 float64, 垂直比例 float64) gocv.Mat {
	return gocv.Resize(矩阵, image.Pt(0, 0), 水平比例, 垂直比例, gocv.InterpolationLinear)
}

// OCV_裁剪 裁剪图片到指定矩形区域。
//
// 参数:
//   - 矩阵: 原始 Mat 对象
//   - 左: 裁剪区域左边界
//   - 上: 裁剪区域上边界
//   - 宽度: 裁剪区域宽度
//   - 高度: 裁剪区域高度
//
// 返回:
//   - gocv.Mat: 裁剪后的 Mat 对象
func OCV_裁剪(矩阵 gocv.Mat, 左 int, 上 int, 宽度 int, 高度 int) gocv.Mat {
	区域 := image.Rect(左, 上, 左+宽度, 上+高度)
	return 矩阵.Region(区域)
}

// OCV_旋转 旋转图片（90° 倍数）。
//
// 参数:
//   - 矩阵: 原始 Mat 对象
//   - 角度: 旋转角度（0=90°顺时针, 1=180°, 2=270°顺时针）
//
// 返回:
//   - gocv.Mat: 旋转后的 Mat 对象
func OCV_旋转(矩阵 gocv.Mat, 角度 int) gocv.Mat {
	switch 角度 {
	case 0:
		return gocv.Rotate(矩阵, gocv.Rotate90Clockwise)
	case 1:
		return gocv.Rotate(矩阵, gocv.Rotate180)
	case 2:
		return gocv.Rotate(矩阵, gocv.Rotate90CounterClockwise)
	default:
		return gocv.Rotate(矩阵, gocv.Rotate90Clockwise)
	}
}

// OCV_水平翻转 水平翻转图片（左右镜像）。
//
// 参数:
//   - 矩阵: 原始 Mat 对象
//
// 返回:
//   - gocv.Mat: 翻转后的 Mat 对象
func OCV_水平翻转(矩阵 gocv.Mat) gocv.Mat {
	return gocv.Flip(矩阵, 1)
}

// OCV_垂直翻转 垂直翻转图片（上下镜像）。
//
// 参数:
//   - 矩阵: 原始 Mat 对象
//
// 返回:
//   - gocv.Mat: 翻转后的 Mat 对象
func OCV_垂直翻转(矩阵 gocv.Mat) gocv.Mat {
	return gocv.Flip(矩阵, 0)
}

// OCV_仿射变换 通过三个点进行仿射变换。
//
// 参数:
//   - 矩阵: 原始 Mat 对象
//   - 源点: 源图像的三个点 [][2]float64
//   - 目标点: 目标图像的三个点 [][2]float64
//
// 返回:
//   - gocv.Mat: 变换后的 Mat 对象
func OCV_仿射变换(矩阵 gocv.Mat, 源点 []image.Point, 目标点 []image.Point) gocv.Mat {
	变换矩阵 := gocv.GetAffineTransform(源点, 目标点)
	defer 变换矩阵.Close()
	return gocv.WarpAffine(矩阵, 变换矩阵, image.Pt(矩阵.Cols(), 矩阵.Rows()))
}

// OCV_透视变换 通过四个点进行透视变换。
//
// 参数:
//   - 矩阵: 原始 Mat 对象
//   - 源点: 源图像的四个点
//   - 目标点: 目标图像的四个点
//
// 返回:
//   - gocv.Mat: 变换后的 Mat 对象
func OCV_透视变换(矩阵 gocv.Mat, 源点 []image.Point, 目标点 []image.Point) gocv.Mat {
	变换矩阵 := gocv.GetPerspectiveTransform(源点, 目标点)
	defer 变换矩阵.Close()
	return gocv.WarpPerspective(矩阵, 变换矩阵, image.Pt(矩阵.Cols(), 矩阵.Rows()))
}

// ============================================================
// 图像滤波
// ============================================================

// OCV_高斯模糊 对图片进行高斯模糊处理。
//
// 参数:
//   - 矩阵: 原始 Mat 对象
//   - 核大小: 模糊核大小（必须为正奇数，如 5、7、15）
//   - 标准差: 高斯核标准差（0 表示自动计算）
//
// 返回:
//   - gocv.Mat: 模糊后的 Mat 对象
func OCV_高斯模糊(矩阵 gocv.Mat, 核大小 int, 标准差 float64) gocv.Mat {
	return gocv.GaussianBlur(矩阵, image.Pt(核大小, 核大小), 标准差, 0, gocv.BorderDefault)
}

// OCV_中值滤波 对图片进行中值滤波（去椒盐噪声）。
//
// 参数:
//   - 矩阵: 原始 Mat 对象
//   - 核大小: 滤波核大小（必须为正奇数）
//
// 返回:
//   - gocv.Mat: 滤波后的 Mat 对象
func OCV_中值滤波(矩阵 gocv.Mat, 核大小 int) gocv.Mat {
	return gocv.MedianBlur(矩阵, 核大小)
}

// OCV_双边滤波 对图片进行双边滤波（保边去噪）。
//
// 参数:
//   - 矩阵: 原始 Mat 对象
//   - 直径: 每个像素邻域的直径
//   - 颜色空间: 颜色空间的标准差
//   - 坐标空间: 坐标空间的标准差
//
// 返回:
//   - gocv.Mat: 滤波后的 Mat 对象
func OCV_双边滤波(矩阵 gocv.Mat, 直径 int, 颜色空间 float64, 坐标空间 float64) gocv.Mat {
	return gocv.BilateralFilter(矩阵, 直径, 颜色空间, 坐标空间)
}

// OCV_方框滤波 对图片进行方框滤波。
//
// 参数:
//   - 矩阵: 原始 Mat 对象
//   - 核大小: 滤波核大小
//
// 返回:
//   - gocv.Mat: 滤波后的 Mat 对象
func OCV_方框滤波(矩阵 gocv.Mat, 核大小 int) gocv.Mat {
	return gocv.Blur(矩阵, image.Pt(核大小, 核大小))
}

// ============================================================
// 形态学操作
// ============================================================

// OCV_腐蚀 对图片进行腐蚀操作（膨胀暗区域）。
//
// 参数:
//   - 矩阵: 原始 Mat 对象
//   - 核大小: 腐蚀核大小
//
// 返回:
//   - gocv.Mat: 腐蚀后的 Mat 对象
func OCV_腐蚀(矩阵 gocv.Mat, 核大小 int) gocv.Mat {
	核 := gocv.GetStructuringElement(gocv.MorphRect, image.Pt(核大小, 核大小))
	defer 核.Close()
	return gocv.Erode(矩阵, 核)
}

// OCV_膨胀 对图片进行膨胀操作（膨胀亮区域）。
//
// 参数:
//   - 矩阵: 原始 Mat 对象
//   - 核大小: 膨胀核大小
//
// 返回:
//   - gocv.Mat: 膨胀后的 Mat 对象
func OCV_膨胀(矩阵 gocv.Mat, 核大小 int) gocv.Mat {
	核 := gocv.GetStructuringElement(gocv.MorphRect, image.Pt(核大小, 核大小))
	defer 核.Close()
	return gocv.Dilate(矩阵, 核)
}

// OCV_开运算 对图片进行形态学开运算（先腐蚀后膨胀，去小噪点）。
//
// 参数:
//   - 矩阵: 原始 Mat 对象
//   - 核大小: 核大小
//
// 返回:
//   - gocv.Mat: 处理后的 Mat 对象
func OCV_开运算(矩阵 gocv.Mat, 核大小 int) gocv.Mat {
	核 := gocv.GetStructuringElement(gocv.MorphRect, image.Pt(核大小, 核大小))
	defer 核.Close()
	return gocv.MorphologyEx(矩阵, gocv.MorphOpen, 核)
}

// OCV_闭运算 对图片进行形态学闭运算（先膨胀后腐蚀，填小孔洞）。
//
// 参数:
//   - 矩阵: 原始 Mat 对象
//   - 核大小: 核大小
//
// 返回:
//   - gocv.Mat: 处理后的 Mat 对象
func OCV_闭运算(矩阵 gocv.Mat, 核大小 int) gocv.Mat {
	核 := gocv.GetStructuringElement(gocv.MorphRect, image.Pt(核大小, 核大小))
	defer 核.Close()
	return gocv.MorphologyEx(矩阵, gocv.MorphClose, 核)
}

// OCV_形态学梯度 对图片进行形态学梯度运算（膨胀减腐蚀，提取边缘）。
//
// 参数:
//   - 矩阵: 原始 Mat 对象
//   - 核大小: 核大小
//
// 返回:
//   - gocv.Mat: 处理后的 Mat 对象
func OCV_形态学梯度(矩阵 gocv.Mat, 核大小 int) gocv.Mat {
	核 := gocv.GetStructuringElement(gocv.MorphRect, image.Pt(核大小, 核大小))
	defer 核.Close()
	return gocv.MorphologyEx(矩阵, gocv.MorphGradient, 核)
}

// OCV_顶帽变换 对图片进行顶帽变换（原图减开运算，提取小亮区域）。
//
// 参数:
//   - 矩阵: 原始 Mat 对象
//   - 核大小: 核大小
//
// 返回:
//   - gocv.Mat: 处理后的 Mat 对象
func OCV_顶帽变换(矩阵 gocv.Mat, 核大小 int) gocv.Mat {
	核 := gocv.GetStructuringElement(gocv.MorphRect, image.Pt(核大小, 核大小))
	defer 核.Close()
	return gocv.MorphologyEx(矩阵, gocv.MorphTopHat, 核)
}

// OCV_黑帽变换 对图片进行黑帽变换（闭运算减原图，提取小暗区域）。
//
// 参数:
//   - 矩阵: 原始 Mat 对象
//   - 核大小: 核大小
//
// 返回:
//   - gocv.Mat: 处理后的 Mat 对象
func OCV_黑帽变换(矩阵 gocv.Mat, 核大小 int) gocv.Mat {
	核 := gocv.GetStructuringElement(gocv.MorphRect, image.Pt(核大小, 核大小))
	defer 核.Close()
	return gocv.MorphologyEx(矩阵, gocv.MorphBlackHat, 核)
}

// ============================================================
// 边缘检测
// ============================================================

// OCV_Canny边缘检测 使用 Canny 算法进行边缘检测。
//
// 参数:
//   - 矩阵: 原始 Mat 对象（建议先转灰度）
//   - 低阈值: 低阈值（如 50）
//   - 高阈值: 高阈值（如 150）
//
// 返回:
//   - gocv.Mat: 边缘检测结果 Mat 对象
func OCV_Canny边缘检测(矩阵 gocv.Mat, 低阈值 float64, 高阈值 float64) gocv.Mat {
	return gocv.Canny(矩阵, 低阈值, 高阈值)
}

// OCV_Sobel边缘检测 使用 Sobel 算子进行边缘检测。
//
// 参数:
//   - 矩阵: 原始 Mat 对象
//   - 核大小: Sobel 核大小（1、3、5 或 7）
//
// 返回:
//   - gocv.Mat: 边缘检测结果 Mat 对象
func OCV_Sobel边缘检测(矩阵 gocv.Mat, 核大小 int) gocv.Mat {
	结果 := gocv.NewMat()
	dx := gocv.Sobel(矩阵, gocv.MatTypeCV64F, 1, 0, 核大小, 1, 0, gocv.BorderDefault)
	defer dx.Close()
	dy := gocv.Sobel(矩阵, gocv.MatTypeCV64F, 0, 1, 核大小, 1, 0, gocv.BorderDefault)
	defer dy.Close()
	gocv.AddWeighted(dx, 0.5, dy, 0.5, 0, &结果)
	return 结果
}

// OCV_Laplacian边缘检测 使用 Laplacian 算子进行边缘检测。
//
// 参数:
//   - 矩阵: 原始 Mat 对象
//   - 核大小: 核大小
//
// 返回:
//   - gocv.Mat: 边缘检测结果 Mat 对象
func OCV_Laplacian边缘检测(矩阵 gocv.Mat, 核大小 int) gocv.Mat {
	return gocv.Laplacian(矩阵, gocv.MatTypeCV64F, 核大小, 1, 0, gocv.BorderDefault)
}

// OCV_Scharr边缘检测 使用 Scharr 算子进行边缘检测。
//
// 参数:
//   - 矩阵: 原始 Mat 对象
//
// 返回:
//   - gocv.Mat: 边缘检测结果 Mat 对象
func OCV_Scharr边缘检测(矩阵 gocv.Mat) gocv.Mat {
	结果 := gocv.NewMat()
	dx := gocv.Scharr(矩阵, gocv.MatTypeCV64F, 1, 0, 1, 0, gocv.BorderDefault)
	defer dx.Close()
	dy := gocv.Scharr(矩阵, gocv.MatTypeCV64F, 0, 1, 1, 0, gocv.BorderDefault)
	defer dy.Close()
	gocv.AddWeighted(dx, 0.5, dy, 0.5, 0, &结果)
	return 结果
}

// ============================================================
// 阈值处理
// ============================================================

// OCV_二值化 对灰度图进行二值化处理。
//
// 参数:
//   - 矩阵: 灰度 Mat 对象
//   - 阈值: 二值化阈值（0-255）
//   - 最大值: 最大值（通常为 255）
//
// 返回:
//   - gocv.Mat: 二值化后的 Mat 对象
func OCV_二值化(矩阵 gocv.Mat, 阈值 float64, 最大值 float64) gocv.Mat {
	结果 := gocv.NewMat()
	gocv.Threshold(矩阵, &结果, 阈值, 最大值, gocv.ThresholdBinary)
	return 结果
}

// OCV_自适应二值化 使用自适应阈值进行二值化。
//
// 参数:
//   - 矩阵: 灰度 Mat 对象
//   - 最大值: 最大值
//   - 核大小: 计算阈值的邻域大小（必须为正奇数）
//   - 常数: 从均值/加权均值中减去的常数
//
// 返回:
//   - gocv.Mat: 二值化后的 Mat 对象
func OCV_自适应二值化(矩阵 gocv.Mat, 最大值 float64, 核大小 int, 常数 float64) gocv.Mat {
	结果 := gocv.NewMat()
	gocv.AdaptiveThreshold(矩阵, &结果, 最大值, gocv.AdaptiveThresholdGaussian, gocv.ThresholdBinary, 核大小, 常数)
	return 结果
}

// OCV_反二值化 对灰度图进行反二值化处理。
//
// 参数:
//   - 矩阵: 灰度 Mat 对象
//   - 阈值: 二值化阈值
//   - 最大值: 最大值
//
// 返回:
//   - gocv.Mat: 反二值化后的 Mat 对象
func OCV_反二值化(矩阵 gocv.Mat, 阈值 float64, 最大值 float64) gocv.Mat {
	结果 := gocv.NewMat()
	gocv.Threshold(矩阵, &结果, 阈值, 最大值, gocv.ThresholdBinaryInv)
	return 结果
}

// OCV_OTSU二值化 使用 OTSU 算法自动计算阈值进行二值化。
//
// 参数:
//   - 矩阵: 灰度 Mat 对象
//   - 最大值: 最大值
//
// 返回:
//   - gocv.Mat: 二值化后的 Mat 对象
func OCV_OTSU二值化(矩阵 gocv.Mat, 最大值 float64) gocv.Mat {
	结果 := gocv.NewMat()
	gocv.Threshold(矩阵, &结果, 0, 最大值, gocv.ThresholdOtsu)
	return 结果
}

// ============================================================
// 轮廓检测
// ============================================================

// OCV_查找轮廓 查找图片中的轮廓。
//
// 参数:
//   - 矩阵: 二值化 Mat 对象
//
// 返回:
//   - [][]image.Point: 轮廓点集
func OCV_查找轮廓(矩阵 gocv.Mat) [][]image.Point {
	return gocv.FindContours(矩阵, gocv.RetrievalExternal, gocv.ChainApproxSimple)
}

// OCV_查找全部轮廓 查找图片中所有层级的轮廓。
//
// 参数:
//   - 矩阵: 二值化 Mat 对象
//
// 返回:
//   - [][]image.Point: 所有轮廓点集
func OCV_查找全部轮廓(矩阵 gocv.Mat) [][]image.Point {
	return gocv.FindContours(矩阵, gocv.RetrievalTree, gocv.ChainApproxSimple)
}

// OCV_绘制轮廓 在图片上绘制轮廓。
//
// 参数:
//   - 矩阵: 目标 Mat 对象
//   - 轮廓: 轮廓点集
//   - 颜色: 绘制颜色
//   - 线宽: 线条宽度
//
// 返回:
//   - gocv.Mat: 绘制后的 Mat 对象
func OCV_绘制轮廓(矩阵 gocv.Mat, 轮廓 [][]image.Point, 颜色 color.RGBA, 线宽 int) gocv.Mat {
	结果 := 矩阵.Clone()
	gocv.DrawContours(&结果, 轮廓, -1, 颜色, 线宽)
	return 结果
}

// OCV_轮廓面积 计算轮廓的面积。
//
// 参数:
//   - 轮廓: 轮廓点集
//
// 返回:
//   - float64: 轮廓面积
func OCV_轮廓面积(轮廓 []image.Point) float64 {
	return gocv.ContourArea(轮廓)
}

// OCV_轮廓周长 计算轮廓的周长。
//
// 参数:
//   - 轮廓: 轮廓点集
//   - 闭合: 轮廓是否闭合
//
// 返回:
//   - float64: 轮廓周长
func OCV_轮廓周长(轮廓 []image.Point, 闭合 bool) float64 {
	return gocv.ArcLength(轮廓, 闭合)
}

// OCV_轮廓外接矩形 获取轮廓的最小外接矩形。
//
// 参数:
//   - 轮廓: 轮廓点集
//
// 返回:
//   - image.Rectangle: 外接矩形
func OCV_轮廓外接矩形(轮廓 []image.Point) image.Rectangle {
	return gocv.BoundingRect(轮廓)
}

// OCV_轮廓最小外接圆 获取轮廓的最小外接圆。
//
// 参数:
//   - 轮廓: 轮廓点集
//
// 返回:
//   - image.Point: 圆心
//   - float64: 半径
func OCV_轮廓最小外接圆(轮廓 []image.Point) (image.Point, float64) {
	return gocv.MinEnclosingCircle(轮廓)
}

// ============================================================
// 特征检测
// ============================================================

// OCV_Harris角点检测 使用 Harris 算法进行角点检测。
//
// 参数:
//   - 矩阵: 灰度 Mat 对象
//   - 核大小: 邻域大小
//   - 标准差: Sobel 标准差
//   - K值: Harris 检测器自由参数
//
// 返回:
//   - gocv.Mat: 角点检测结果 Mat 对象
func OCV_Harris角点检测(矩阵 gocv.Mat, 核大小 int, 标准差 float64, K值 float64) gocv.Mat {
	return gocv.CornerHarris(矩阵, 核大小, int(标准差), K值, gocv.BorderDefault)
}

// OCV_良好角点 使用 Shi-Tomasi 算法检测良好角点。
//
// 参数:
//   - 矩阵: 灰度 Mat 对象
//   - 最大数量: 最大角点数量
//   - 质量: 质量水平（0-1）
//   - 最小距离: 角点间最小距离
//
// 返回:
//   - []image.Point: 角点坐标列表
func OCV_良好角点(矩阵 gocv.Mat, 最大数量 int, 质量 float64, 最小距离 float64) []image.Point {
	return gocv.GoodFeaturesToTrack(矩阵, 最大数量, 质量, 最小距离)
}

// OCV_FAST角点 使用 FAST 算法检测角点。
//
// 参数:
//   - 矩阵: 灰度 Mat 对象
//   - 阈值: FAST 阈值
//
// 返回:
//   - []image.Point: 角点坐标列表
func OCV_FAST角点(矩阵 gocv.Mat, 阈值 int) []image.Point {
	return gocv.FAST(矩阵, 阈值, true)
}

// ============================================================
// 模板匹配（找图）
// ============================================================

// OCV_找图 在源图中查找模板图的位置（使用归一化相关系数匹配法）。
//
// 参数:
//   - 源图: 源图像 Mat 对象（大图）
//   - 模板: 模板图像 Mat 对象（小图）
//   - 相似度: 最低相似度阈值（0-1，默认 0.8；低于此值视为未找到）
//
// 返回:
//   - image.Point: 匹配区域左上角坐标，未找到返回 image.ZP
//   - float64: 匹配相似度（0-1），越高越相似
//   - error: 模板尺寸大于源图时返回错误
func OCV_找图(源图 gocv.Mat, 模板 gocv.Mat, 相似度 float64) (image.Point, float64, error) {
	return OCV_找图带方法(源图, 模板, 相似度, gocv.TmCcoeffNormed)
}

// OCV_找图带方法 使用指定匹配方法在源图中查找模板图。
//
// 参数:
//   - 源图: 源图像 Mat 对象（大图）
//   - 模板: 模板图像 Mat 对象（小图）
//   - 相似度: 最低相似度阈值（0-1）
//   - 方法: 匹配方法
//     gocv.TmSqdiffNormed   - 归一化平方差（值越小越匹配，相似度自动转为 1-值）
//     gocv.TmCcorrNormed    - 归一化相关匹配
//     gocv.TmCcoeffNormed   - 归一化相关系数匹配（推荐，抗光照变化）
//
// 返回:
//   - image.Point: 匹配区域左上角坐标
//   - float64: 匹配相似度（0-1）
//   - error: 模板尺寸大于源图时返回错误
func OCV_找图带方法(源图 gocv.Mat, 模板 gocv.Mat, 相似度 float64, 方法 gocv.TemplateMatchMode) (image.Point, float64, error) {
	if 模板.Rows() > 源图.Rows() || 模板.Cols() > 源图.Cols() {
		return image.ZP, 0, fmt.Errorf("模板尺寸(%dx%d)不能大于源图(%dx%d)", 模板.Cols(), 模板.Rows(), 源图.Cols(), 源图.Rows())
	}

	结果 := gocv.NewMat()
	defer 结果.Close()

	gocv.MatchTemplate(源图, 模板, &结果, 方法, gocv.NewMat())

	_, 最大置信度, _, 最大位置 := gocv.MinMaxLoc(结果)

	最终置信度 := 最大置信度
	if 方法 == gocv.TmSqdiff || 方法 == gocv.TmSqdiffNormed {
		最小置信度, _, 最小位置, _ := gocv.MinMaxLoc(结果)
		最终置信度 = 1 - 最小置信度
		if 最终置信度 < 相似度 {
			return image.ZP, 最终置信度, nil
		}
		return 最小位置, 最终置信度, nil
	}

	if 最终置信度 < 相似度 {
		return image.ZP, 最终置信度, nil
	}
	return 最大位置, 最终置信度, nil
}

// OCV_找图中心 在源图中查找模板图，返回匹配区域的中心点坐标。
//
// 参数:
//   - 源图: 源图像 Mat 对象（大图）
//   - 模板: 模板图像 Mat 对象（小图）
//   - 相似度: 最低相似度阈值（0-1）
//
// 返回:
//   - image.Point: 匹配区域中心点坐标，未找到返回 image.ZP
//   - float64: 匹配相似度（0-1）
//   - error: 匹配失败时返回错误
func OCV_找图中心(源图 gocv.Mat, 模板 gocv.Mat, 相似度 float64) (image.Point, float64, error) {
	位置, 置信度, err := OCV_找图(源图, 模板, 相似度)
	if err != nil {
		return image.ZP, 0, err
	}
	if 位置 == image.ZP && 置信度 < 相似度 {
		return image.ZP, 0, nil
	}
	中心X := 位置.X + 模板.Cols()/2
	中心Y := 位置.Y + 模板.Rows()/2
	return image.Pt(中心X, 中心Y), 置信度, nil
}

// OCV_找图全部 在源图中查找所有匹配位置（相似度 >= 阈值）。
//
// 参数:
//   - 源图: 源图像 Mat 对象（大图）
//   - 模板: 模板图像 Mat 对象（小图）
//   - 相似度: 最低相似度阈值（0-1）
//   - 最小间距: 相邻匹配点最小间距（像素，避免重复匹配）
//
// 返回:
//   - []image.Point: 所有匹配区域左上角坐标列表
//   - []float64: 对应的相似度列表
//   - error: 模板尺寸大于源图时返回错误
func OCV_找图全部(源图 gocv.Mat, 模板 gocv.Mat, 相似度 float64, 最小间距 int) ([]image.Point, []float64, error) {
	if 模板.Rows() > 源图.Rows() || 模板.Cols() > 源图.Cols() {
		return nil, nil, fmt.Errorf("模板尺寸(%dx%d)不能大于源图(%dx%d)", 模板.Cols(), 模板.Rows(), 源图.Cols(), 源图.Rows())
	}

	结果 := gocv.NewMat()
	defer 结果.Close()

	gocv.MatchTemplate(源图, 模板, &结果, gocv.TmCcoeffNormed, gocv.NewMat())

	var 位置列表 []image.Point
	var 置信度列表 []float64

	// 获取结果矩阵的数据类型 (32FC1)
	行 := 结果.Rows()
	列 := 结果.Cols()

	for {
		_, 最大置信度, _, 最大位置 := gocv.MinMaxLoc(结果)
		if 最大置信度 < 相似度 {
			break
		}

		位置列表 = append(位置列表, 最大位置)
		置信度列表 = append(置信度列表, 最大置信度)

		// 屏蔽已匹配区域，防止重复匹配
		// 以最大值点为中心，绘制矩形屏蔽区域
		左 := 最大位置.X - 最小间距
		if 左 < 0 {
			左 = 0
		}
		上 := 最大位置.Y - 最小间距
		if 上 < 0 {
			上 = 0
		}
		右 := 最大位置.X + 最小间距
		if 右 > 列 {
			右 = 列
		}
		下 := 最大位置.Y + 最小间距
		if 下 > 行 {
			下 = 行
		}
		gocv.Rectangle(&结果, image.Rect(左, 上, 右, 下), color.RGBA{0, 0, 0, 255}, -1)
	}

	return 位置列表, 置信度列表, nil
}

// OCV_找图区域 在源图的指定区域内查找模板图。
//
// 参数:
//   - 源图: 源图像 Mat 对象（大图）
//   - 模板: 模板图像 Mat 对象（小图）
//   - 左: 搜索区域左边界
//   - 上: 搜索区域上边界
//   - 宽度: 搜索区域宽度
//   - 高度: 搜索区域高度
//   - 相似度: 最低相似度阈值（0-1）
//
// 返回:
//   - image.Point: 匹配区域左上角在源图中的坐标
//   - float64: 匹配相似度（0-1）
//   - error: 匹配失败时返回错误
func OCV_找图区域(源图 gocv.Mat, 模板 gocv.Mat, 左 int, 上 int, 宽度 int, 高度 int, 相似度 float64) (image.Point, float64, error) {
	区域 := 源图.Region(image.Rect(左, 上, 左+宽度, 上+高度))
	defer 区域.Close()

	局部位置, 置信度, err := OCV_找图(区域, 模板, 相似度)
	if err != nil {
		return image.ZP, 0, err
	}
	if 局部位置 == image.ZP && 置信度 < 相似度 {
		return image.ZP, 置信度, nil
	}

	// 转换为源图坐标
	全局X := 左 + 局部位置.X
	全局Y := 上 + 局部位置.Y
	return image.Pt(全局X, 全局Y), 置信度, nil
}

// OCV_找图掩码 带掩码的模板匹配（支持透明图匹配，忽略掩码中的黑色区域）。
//
// 参数:
//   - 源图: 源图像 Mat 对象（大图）
//   - 模板: 模板图像 Mat 对象（小图，通常为 BGRA 带 alpha 通道）
//   - 掩码: 掩码 Mat 对象（模板中需要参与匹配的白色区域）
//   - 相似度: 最低相似度阈值（0-1）
//
// 返回:
//   - image.Point: 匹配区域左上角坐标
//   - float64: 匹配相似度（0-1）
//   - error: 匹配失败时返回错误
func OCV_找图掩码(源图 gocv.Mat, 模板 gocv.Mat, 掩码 gocv.Mat, 相似度 float64) (image.Point, float64, error) {
	if 模板.Rows() > 源图.Rows() || 模板.Cols() > 源图.Cols() {
		return image.ZP, 0, fmt.Errorf("模板尺寸(%dx%d)不能大于源图(%dx%d)", 模板.Cols(), 模板.Rows(), 源图.Cols(), 源图.Rows())
	}

	结果 := gocv.NewMat()
	defer 结果.Close()

	gocv.MatchTemplate(源图, 模板, &结果, gocv.TmCcoeffNormed, 掩码)

	_, 最大置信度, _, 最大位置 := gocv.MinMaxLoc(结果)
	if 最大置信度 < 相似度 {
		return image.ZP, 最大置信度, nil
	}
	return 最大位置, 最大置信度, nil
}

// ============================================================
// 特征匹配（高级找图，抗缩放/旋转/光照变化）
// ============================================================

// OCV_找图SIFT 使用 SIFT 特征匹配在源图中查找模板（抗缩放、旋转、视角变化）。
//
// 参数:
//   - 源图: 源图像 Mat 对象
//   - 模板: 模板图像 Mat 对象
//   - 最小匹配点数: 最少有效匹配点数（如 10，低于此值视为未找到）
//
// 返回:
//   - image.Rectangle: 匹配区域外接矩形，未找到返回 image.ZR
//   - int: 有效匹配点数量
//   - error: 匹配失败时返回错误
func OCV_找图SIFT(源图 gocv.Mat, 模板 gocv.Mat, 最小匹配点数 int) (image.Rectangle, []gocv.DMatch, error) {
	sift := gocv.NewSIFT()
	defer sift.Close()

	return OCV_特征匹配(sift, 源图, 模板, 最小匹配点数, gocv.NormL2, 0.75)
}

// OCV_找图ORB 使用 ORB 特征匹配在源图中查找模板（快速，免费，适合实时场景）。
//
// 参数:
//   - 源图: 源图像 Mat 对象
//   - 模板: 模板图像 Mat 对象
//   - 最小匹配点数: 最少有效匹配点数（如 10）
//
// 返回:
//   - image.Rectangle: 匹配区域外接矩形
//   - int: 有效匹配点数量
//   - error: 匹配失败时返回错误
func OCV_找图ORB(源图 gocv.Mat, 模板 gocv.Mat, 最小匹配点数 int) (image.Rectangle, []gocv.DMatch, error) {
	orb := gocv.NewORB()
	defer orb.Close()

	return OCV_特征匹配(orb, 源图, 模板, 最小匹配点数, gocv.NormHamming, 0.75)
}

// OCV_找图AKAZE 使用 AKAZE 特征匹配在源图中查找模板（非线性尺度空间，适合模糊/压缩图片）。
//
// 参数:
//   - 源图: 源图像 Mat 对象
//   - 模板: 模板图像 Mat 对象
//   - 最小匹配点数: 最少有效匹配点数（如 10）
//
// 返回:
//   - image.Rectangle: 匹配区域外接矩形
//   - int: 有效匹配点数量
//   - error: 匹配失败时返回错误
func OCV_找图AKAZE(源图 gocv.Mat, 模板 gocv.Mat, 最小匹配点数 int) (image.Rectangle, []gocv.DMatch, error) {
	akaze := gocv.NewAKAZE()
	defer akaze.Close()

	return OCV_特征匹配(akaze, 源图, 模板, 最小匹配点数, gocv.NormL2, 0.75)
}

// OCV_特征匹配 通用特征匹配引擎（SIFT/ORB/AKAZE 内部调用此函数）。
//
// 参数:
//   - 检测器: 特征检测器（实现 Feature2D 接口）
//   - 源图: 源图像 Mat
//   - 模板: 模板图像 Mat
//   - 最小匹配点数: 最少有效匹配点数
//   - 范数类型: 距离范数（gocv.NormL2 用于 SIFT/AKAZE，gocv.NormHamming 用于 ORB）
//   - 距离比例: Lowe's ratio test 阈值（0.7-0.8，越严格匹配越少但越准确）
//
// 返回:
//   - image.Rectangle: 匹配区域外接矩形
//   - int: 有效匹配点数量
//   - error: 匹配失败时返回错误
func OCV_特征匹配(检测器 gocv.Feature2D, 源图 gocv.Mat, 模板 gocv.Mat, 最小匹配点数 int, 范数类型 gocv.NormType, 距离比例 float64) (image.Rectangle, []gocv.DMatch, error) {
	if 模板.Rows() > 源图.Rows() || 模板.Cols() > 源图.Cols() {
		return image.ZR, nil, fmt.Errorf("模板尺寸(%dx%d)不能大于源图(%dx%d)", 模板.Cols(), 模板.Rows(), 源图.Cols(), 源图.Rows())
	}

	源关键点, 源描述符 := 检测器.DetectAndCompute(源图, gocv.NewMat())
	defer 源描述符.Close()
	模板关键点, 模板描述符 := 检测器.DetectAndCompute(模板, gocv.NewMat())
	defer 模板描述符.Close()

	if len(源关键点) < 2 || len(模板关键点) < 2 {
		return image.ZR, nil, fmt.Errorf("检测到的特征点不足（源图:%d, 模板:%d）", len(源关键点), len(模板关键点))
	}

	匹配器 := gocv.NewBFMatcherWithParams(范数类型, false)
	defer 匹配器.Close()

	knn匹配 := 匹配器.KnnMatch(模板描述符, 源描述符, 2)

	var 好匹配 []gocv.DMatch
	for _, 匹配对 := range knn匹配 {
		if len(匹配对) < 2 {
			continue
		}
		if 匹配对[0].Distance < 距离比例*匹配对[1].Distance {
			好匹配 = append(好匹配, 匹配对[0])
		}
	}

	if len(好匹配) < 最小匹配点数 {
		return image.ZR, 好匹配, nil
	}

	// 从好匹配中提取源图像点坐标，计算外接矩形
	var 源点 []image.Point
	for _, m := range 好匹配 {
		源点 = append(源点, image.Pt(int(源关键点[m.TrainIdx].X), int(源关键点[m.TrainIdx].Y)))
	}

	最小X, 最小Y := 源点[0].X, 源点[0].Y
	最大X, 最大Y := 源点[0].X, 源点[0].Y
	for _, p := range 源点 {
		if p.X < 最小X {
			最小X = p.X
		}
		if p.Y < 最小Y {
			最小Y = p.Y
		}
		if p.X > 最大X {
			最大X = p.X
		}
		if p.Y > 最大Y {
			最大Y = p.Y
		}
	}

	// 扩展边界以包含整个模板区域
	矩形 := image.Rect(最小X, 最小Y, 最大X+模板.Cols()/2, 最大Y+模板.Rows()/2)
	if 矩形.Max.X > 源图.Cols() {
		矩形.Max.X = 源图.Cols()
	}
	if 矩形.Max.Y > 源图.Rows() {
		矩形.Max.Y = 源图.Rows()
	}

	return 矩形, 好匹配, nil
}

// OCV_找图多尺度 多尺度模板匹配（模板以不同比例缩放后逐一匹配，抗缩放变化）。
//
// 参数:
//   - 源图: 源图像 Mat 对象
//   - 模板: 模板图像 Mat 对象
//   - 最小比例: 最小缩放比例（如 0.5 = 50%）
//   - 最大比例: 最大缩放比例（如 1.5 = 150%）
//   - 步长: 缩放步长（如 0.1，越小越精细但越慢）
//   - 相似度: 最低相似度阈值（0-1）
//
// 返回:
//   - image.Point: 最佳匹配区域左上角坐标
//   - float64: 最佳匹配相似度
//   - float64: 最佳匹配时的缩放比例
//   - error: 匹配失败时返回错误
func OCV_找图多尺度(源图 gocv.Mat, 模板 gocv.Mat, 最小比例 float64, 最大比例 float64, 步长 float64, 相似度 float64) (image.Point, float64, float64, error) {
	if 模板.Rows() > 源图.Rows() || 模板.Cols() > 源图.Cols() {
		return image.ZP, 0, 0, fmt.Errorf("模板尺寸(%dx%d)不能大于源图(%dx%d)", 模板.Cols(), 模板.Rows(), 源图.Cols(), 源图.Rows())
	}

	最佳位置 := image.ZP
	最佳置信度 := 0.0
	最佳比例 := 1.0

	for 比例 := 最小比例; 比例 <= 最大比例; 比例 += 步长 {
		新宽度 := int(float64(模板.Cols()) * 比例)
		新高度 := int(float64(模板.Rows()) * 比例)
		if 新宽度 < 5 || 新高度 < 5 {
			continue
		}
		if 新宽度 > 源图.Cols() || 新高度 > 源图.Rows() {
			continue
		}

		缩放模板 := gocv.Resize(模板, image.Pt(新宽度, 新高度), 0, 0, gocv.InterpolationLinear)

		位置, 置信度, _ := OCV_找图(源图, 缩放模板, 相似度)
		缩放模板.Close()

		if 置信度 > 最佳置信度 {
			最佳置信度 = 置信度
			最佳位置 = 位置
			最佳比例 = 比例
		}
	}

	if 最佳置信度 < 相似度 {
		return image.ZP, 最佳置信度, 最佳比例, nil
	}
	return 最佳位置, 最佳置信度, 最佳比例, nil
}

// OCV_找图边缘 边缘模板匹配（先 Canny 提取边缘再匹配，对光照变化和颜色差异不敏感）。
//
// 参数:
//   - 源图: 源图像 Mat 对象（自动转灰度+Canny）
//   - 模板: 模板图像 Mat 对象（自动转灰度+Canny）
//   - 低阈值: Canny 低阈值（如 50）
//   - 高阈值: Canny 高阈值（如 150）
//   - 相似度: 最低相似度阈值（0-1）
//
// 返回:
//   - image.Point: 匹配区域左上角坐标
//   - float64: 匹配相似度（0-1）
//   - error: 匹配失败时返回错误
func OCV_找图边缘(源图 gocv.Mat, 模板 gocv.Mat, 低阈值 float64, 高阈值 float64, 相似度 float64) (image.Point, float64, error) {
	源图灰度 := gocv.CvtColor(源图, gocv.ColorBGRToGray)
	defer 源图灰度.Close()
	模板灰度 := gocv.CvtColor(模板, gocv.ColorBGRToGray)
	defer 模板灰度.Close()

	源边缘 := gocv.Canny(源图灰度, 低阈值, 高阈值)
	defer 源边缘.Close()
	模板边缘 := gocv.Canny(模板灰度, 低阈值, 高阈值)
	defer 模板边缘.Close()

	return OCV_找图(源边缘, 模板边缘, 相似度)
}

// ============================================================
// 直方图
// ============================================================

// OCV_计算直方图 计算灰度图直方图。
//
// 参数:
//   - 矩阵: 灰度 Mat 对象
//   - 区间数: 直方图区间数（如 256）
//
// 返回:
//   - gocv.Mat: 直方图 Mat 对象
func OCV_计算直方图(矩阵 gocv.Mat, 区间数 int) gocv.Mat {
	return gocv.CalcHist([]gocv.Mat{矩阵}, []int{0}, gocv.NewMat(), []int{区间数}, []float64{0, 256}, false)
}

// OCV_直方图均衡化 对灰度图进行直方图均衡化。
//
// 参数:
//   - 矩阵: 灰度 Mat 对象
//
// 返回:
//   - gocv.Mat: 均衡化后的 Mat 对象
func OCV_直方图均衡化(矩阵 gocv.Mat) gocv.Mat {
	return gocv.EqualizeHist(矩阵)
}

// OCV_CLAHE均衡化 使用 CLAHE 算法进行自适应直方图均衡化。
//
// 参数:
//   - 矩阵: 灰度 Mat 对象
//   - 限幅: 限幅阈值（如 2.0）
//   - 网格大小: 网格大小（如 8）
//
// 返回:
//   - gocv.Mat: 均衡化后的 Mat 对象
func OCV_CLAHE均衡化(矩阵 gocv.Mat, 限幅 float64, 网格大小 int) gocv.Mat {
	结果 := gocv.NewMat()
	clahe := gocv.NewCLAHEWithParams(限幅, image.Pt(网格大小, 网格大小))
	defer clahe.Close()
	clahe.Apply(矩阵, &结果)
	return 结果
}

// ============================================================
// 绘图
// ============================================================

// OCV_画线 在图片上画直线。
//
// 参数:
//   - 矩阵: 目标 Mat 对象
//   - 起点: 线段起点
//   - 终点: 线段终点
//   - 颜色: 线条颜色
//   - 线宽: 线条宽度
//
// 返回:
//   - gocv.Mat: 绘制后的 Mat 对象
func OCV_画线(矩阵 gocv.Mat, 起点 image.Point, 终点 image.Point, 颜色 color.RGBA, 线宽 int) gocv.Mat {
	结果 := 矩阵.Clone()
	gocv.Line(&结果, 起点, 终点, 颜色, 线宽)
	return 结果
}

// OCV_画矩形 在图片上画矩形。
//
// 参数:
//   - 矩阵: 目标 Mat 对象
//   - 矩形: 矩形区域
//   - 颜色: 线条颜色
//   - 线宽: 线条宽度（-1 为填充）
//
// 返回:
//   - gocv.Mat: 绘制后的 Mat 对象
func OCV_画矩形(矩阵 gocv.Mat, 矩形 image.Rectangle, 颜色 color.RGBA, 线宽 int) gocv.Mat {
	结果 := 矩阵.Clone()
	gocv.Rectangle(&结果, 矩形, 颜色, 线宽)
	return 结果
}

// OCV_画圆 在图片上画圆。
//
// 参数:
//   - 矩阵: 目标 Mat 对象
//   - 圆心: 圆心坐标
//   - 半径: 圆的半径
//   - 颜色: 线条颜色
//   - 线宽: 线条宽度（-1 为填充）
//
// 返回:
//   - gocv.Mat: 绘制后的 Mat 对象
func OCV_画圆(矩阵 gocv.Mat, 圆心 image.Point, 半径 int, 颜色 color.RGBA, 线宽 int) gocv.Mat {
	结果 := 矩阵.Clone()
	gocv.Circle(&结果, 圆心, 半径, 颜色, 线宽)
	return 结果
}

// OCV_画文字 在图片上绘制文字。
//
// 参数:
//   - 矩阵: 目标 Mat 对象
//   - 文字: 要绘制的文字
//   - 位置: 文字左下角坐标
//   - 字体: 字体类型（gocv.FontHersheySimplex 等）
//   - 大小: 字体大小
//   - 颜色: 文字颜色
//   - 线宽: 线条宽度
//
// 返回:
//   - gocv.Mat: 绘制后的 Mat 对象
func OCV_画文字(矩阵 gocv.Mat, 文字 string, 位置 image.Point, 字体 gocv.HersheyFont, 大小 float64, 颜色 color.RGBA, 线宽 int) gocv.Mat {
	结果 := 矩阵.Clone()
	gocv.PutText(&结果, 文字, 位置, 字体, 大小, 颜色, 线宽)
	return 结果
}

// OCV_画椭圆 在图片上画椭圆。
//
// 参数:
//   - 矩阵: 目标 Mat 对象
//   - 圆心: 椭圆中心
//   - 长轴: 长轴半径
//   - 短轴: 短轴半径
//   - 旋转角: 旋转角度
//   - 起始角: 起始角度
//   - 终止角: 终止角度
//   - 颜色: 线条颜色
//   - 线宽: 线条宽度（-1 为填充）
//
// 返回:
//   - gocv.Mat: 绘制后的 Mat 对象
func OCV_画椭圆(矩阵 gocv.Mat, 圆心 image.Point, 长轴 int, 短轴 int, 旋转角 float64, 起始角 float64, 终止角 float64, 颜色 color.RGBA, 线宽 int) gocv.Mat {
	结果 := 矩阵.Clone()
	gocv.Ellipse(&结果, 圆心, image.Pt(长轴, 短轴), 旋转角, 起始角, 终止角, 颜色, 线宽)
	return 结果
}

// ============================================================
// 图像运算
// ============================================================

// OCV_加法 两张图片逐像素相加。
//
// 参数:
//   - 矩阵1: 第一张图片
//   - 矩阵2: 第二张图片
//
// 返回:
//   - gocv.Mat: 相加结果
func OCV_加法(矩阵1 gocv.Mat, 矩阵2 gocv.Mat) gocv.Mat {
	return gocv.Add(矩阵1, 矩阵2)
}

// OCV_加权加法 两张图片按权重相加。
//
// 参数:
//   - 矩阵1: 第一张图片
//   - 权重1: 第一张图片权重
//   - 矩阵2: 第二张图片
//   - 权重2: 第二张图片权重
//   - 伽马值: 标量偏移
//
// 返回:
//   - gocv.Mat: 加权相加结果
func OCV_加权加法(矩阵1 gocv.Mat, 权重1 float64, 矩阵2 gocv.Mat, 权重2 float64, 伽马值 float64) gocv.Mat {
	结果 := gocv.NewMat()
	gocv.AddWeighted(矩阵1, 权重1, 矩阵2, 权重2, 伽马值, &结果)
	return 结果
}

// OCV_减法 两张图片逐像素相减。
//
// 参数:
//   - 矩阵1: 被减图
//   - 矩阵2: 减图
//
// 返回:
//   - gocv.Mat: 相减结果
func OCV_减法(矩阵1 gocv.Mat, 矩阵2 gocv.Mat) gocv.Mat {
	return gocv.Subtract(矩阵1, 矩阵2)
}

// OCV_按位与 两张图片进行按位与运算。
//
// 参数:
//   - 矩阵1: 第一张图片
//   - 矩阵2: 第二张图片
//
// 返回:
//   - gocv.Mat: 按位与结果
func OCV_按位与(矩阵1 gocv.Mat, 矩阵2 gocv.Mat) gocv.Mat {
	return gocv.BitwiseAnd(矩阵1, 矩阵2)
}

// OCV_按位或 两张图片进行按位或运算。
//
// 参数:
//   - 矩阵1: 第一张图片
//   - 矩阵2: 第二张图片
//
// 返回:
//   - gocv.Mat: 按位或结果
func OCV_按位或(矩阵1 gocv.Mat, 矩阵2 gocv.Mat) gocv.Mat {
	return gocv.BitwiseOr(矩阵1, 矩阵2)
}

// OCV_按位异或 两张图片进行按位异或运算。
//
// 参数:
//   - 矩阵1: 第一张图片
//   - 矩阵2: 第二张图片
//
// 返回:
//   - gocv.Mat: 按位异或结果
func OCV_按位异或(矩阵1 gocv.Mat, 矩阵2 gocv.Mat) gocv.Mat {
	return gocv.BitwiseXor(矩阵1, 矩阵2)
}

// OCV_按位取反 对图片进行按位取反运算。
//
// 参数:
//   - 矩阵: 原始图片
//
// 返回:
//   - gocv.Mat: 取反结果
func OCV_按位取反(矩阵 gocv.Mat) gocv.Mat {
	return gocv.BitwiseNot(矩阵)
}

// ============================================================
// 视频与摄像头
// ============================================================

// OCV_打开摄像头 打开默认摄像头。
//
// 参数:
//   - 设备ID: 摄像头设备编号（0 为默认）
//
// 返回:
//   - *gocv.VideoCapture: 视频捕获对象
//   - error: 打开失败时返回错误
func OCV_打开摄像头(设备ID int) (*gocv.VideoCapture, error) {
	return gocv.OpenVideoCapture(设备ID)
}

// OCV_打开视频文件 打开视频文件。
//
// 参数:
//   - 文件路径: 视频文件路径
//
// 返回:
//   - *gocv.VideoCapture: 视频捕获对象
//   - error: 打开失败时返回错误
func OCV_打开视频文件(文件路径 string) (*gocv.VideoCapture, error) {
	return gocv.OpenVideoCapture(文件路径)
}

// OCV_读取帧 从视频捕获对象读取一帧。
//
// 参数:
//   - 捕获器: 视频捕获对象
//
// 返回:
//   - gocv.Mat: 帧图像
//   - bool: 读取成功返回 true
func OCV_读取帧(捕获器 *gocv.VideoCapture) (gocv.Mat, bool) {
	帧 := gocv.NewMat()
	ok := 捕获器.Read(&帧)
	return 帧, ok
}

// OCV_创建视频写入器 创建视频写入器。
//
// 参数:
//   - 文件路径: 输出视频文件路径
//   - 编码器: 四字符编码器代码（如 "MJPG"）
//   - 帧率: 视频帧率
//   - 宽度: 视频宽度
//   - 高度: 视频高度
//
// 返回:
//   - *gocv.VideoWriter: 视频写入器对象
//   - error: 创建失败时返回错误
func OCV_创建视频写入器(文件路径 string, 编码器 string, 帧率 float64, 宽度 int, 高度 int) (*gocv.VideoWriter, error) {
	return gocv.VideoWriterFile(文件路径, 编码器, 帧率, 宽度, 高度, true)
}

// OCV_写入帧 向视频写入器写入一帧。
//
// 参数:
//   - 写入器: 视频写入器对象
//   - 帧: 帧图像
func OCV_写入帧(写入器 *gocv.VideoWriter, 帧 gocv.Mat) {
	写入器.Write(帧)
}

// ============================================================
// 图像转换
// ============================================================

// OCV_Mat转Image 将 gocv.Mat 转换为 Go 标准 image.Image。
//
// 参数:
//   - 矩阵: OpenCV Mat 对象
//
// 返回:
//   - image.Image: Go 标准图像对象
func OCV_Mat转Image(矩阵 gocv.Mat) image.Image {
	return 矩阵.ToImage()
}

// OCV_Image转Mat 将 Go 标准 image.Image 謧换为 gocv.Mat。
//
// 参数:
//   - 图片: Go 标准图像对象
//
// 返回:
//   - gocv.Mat: OpenCV Mat 对象
//   - error: 转换失败时返回错误
func OCV_Image转Mat(图片 image.Image) (gocv.Mat, error) {
	return gocv.ImageToMatRGB(图片)
}
