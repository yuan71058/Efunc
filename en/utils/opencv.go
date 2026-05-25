//go:build opencv

// OpenCV 计算机视觉模块
// 基于 gocv.io/x/gocv 库，提供图像处理、边缘检测、形态学操作、模板匹配、
// 轮廓检测、特征检测、人脸检测、视频分析等功能。
// 适用于自动化测试、图像识别、目标检测等场景。
package utils

import (
	"image"
	"image/color"

	"gocv.io/x/gocv"
)

// ============================================================
// OpenCV 核心操作
// ============================================================

func OCV_Version() string                       { return gocv.Version() }
func OCV_CUDADeviceCount() int                  { return gocv.GetCudaEnabledDeviceCount() }

// ============================================================
// 图像读取与保存
// ============================================================

func OCV_IMRead(filePath string) (gocv.Mat, error) {
	return gocv.IMRead(filePath, gocv.IMReadColor)
}

func OCV_IMReadGray(filePath string) (gocv.Mat, error) {
	return gocv.IMRead(filePath, gocv.IMReadGrayScale)
}

func OCV_IMReadUnchanged(filePath string) (gocv.Mat, error) {
	return gocv.IMRead(filePath, gocv.IMReadUnchanged)
}

func OCV_IMWrite(mat gocv.Mat, filePath string) bool {
	return gocv.IMWrite(filePath, mat)
}

func OCV_IMDecode(data []byte) (gocv.Mat, error) {
	return gocv.IMDecode(data, gocv.IMReadColor)
}

func OCV_IMDecodeGray(data []byte) (gocv.Mat, error) {
	return gocv.IMDecode(data, gocv.IMReadGrayScale)
}

func OCV_IMEncode(mat gocv.Mat, ext string) ([]byte, error) {
	buf, err := gocv.IMEncode(ext, mat)
	if err != nil {
		return nil, err
	}
	return buf.GetBytes(), nil
}

// ============================================================
// Mat 信息与属性
// ============================================================

func OCV_Width(mat gocv.Mat) int     { return mat.Cols() }
func OCV_Height(mat gocv.Mat) int    { return mat.Rows() }
func OCV_Channels(mat gocv.Mat) int  { return mat.Channels() }
func OCV_Type(mat gocv.Mat) gocv.MatType { return mat.Type() }
func OCV_Total(mat gocv.Mat) int     { return mat.Total() }
func OCV_IsEmpty(mat gocv.Mat) bool  { return mat.Empty() }
func OCV_Size(mat gocv.Mat) (int, int) { return mat.Cols(), mat.Rows() }

// ============================================================
// 颜色空间转换
// ============================================================

func OCV_BGRToGray(mat gocv.Mat) gocv.Mat { return gocv.CvtColor(mat, gocv.ColorBGRToGray) }
func OCV_GrayToBGR(mat gocv.Mat) gocv.Mat { return gocv.CvtColor(mat, gocv.ColorGrayToBGR) }
func OCV_BGRToHSV(mat gocv.Mat) gocv.Mat  { return gocv.CvtColor(mat, gocv.ColorBGRToHSV) }
func OCV_HSVToBGR(mat gocv.Mat) gocv.Mat  { return gocv.CvtColor(mat, gocv.ColorHSVToBGR) }
func OCV_BGRToRGB(mat gocv.Mat) gocv.Mat  { return gocv.CvtColor(mat, gocv.ColorBGRToRGB) }
func OCV_RGBToBGR(mat gocv.Mat) gocv.Mat  { return gocv.CvtColor(mat, gocv.ColorRGBToBGR) }
func OCV_BGRToLab(mat gocv.Mat) gocv.Mat  { return gocv.CvtColor(mat, gocv.ColorBGRToLab) }
func OCV_BGRToYUV(mat gocv.Mat) gocv.Mat  { return gocv.CvtColor(mat, gocv.ColorBGRToYUV) }

// ============================================================
// 图像变换
// ============================================================

func OCV_Resize(mat gocv.Mat, width int, height int) gocv.Mat {
	return gocv.Resize(mat, image.Pt(width, height), 0, 0, gocv.InterpolationLinear)
}

func OCV_ResizeByRatio(mat gocv.Mat, fx float64, fy float64) gocv.Mat {
	return gocv.Resize(mat, image.Pt(0, 0), fx, fy, gocv.InterpolationLinear)
}

func OCV_Crop(mat gocv.Mat, left int, top int, width int, height int) gocv.Mat {
	return mat.Region(image.Rect(left, top, left+width, top+height))
}

func OCV_Rotate(mat gocv.Mat, angle int) gocv.Mat {
	switch angle {
	case 0:
		return gocv.Rotate(mat, gocv.Rotate90Clockwise)
	case 1:
		return gocv.Rotate(mat, gocv.Rotate180)
	case 2:
		return gocv.Rotate(mat, gocv.Rotate90CounterClockwise)
	default:
		return gocv.Rotate(mat, gocv.Rotate90Clockwise)
	}
}

func OCV_FlipH(mat gocv.Mat) gocv.Mat { return gocv.Flip(mat, 1) }
func OCV_FlipV(mat gocv.Mat) gocv.Mat { return gocv.Flip(mat, 0) }

func OCV_AffineTransform(mat gocv.Mat, src []image.Point, dst []image.Point) gocv.Mat {
	transform := gocv.GetAffineTransform(src, dst)
	defer transform.Close()
	return gocv.WarpAffine(mat, transform, image.Pt(mat.Cols(), mat.Rows()))
}

func OCV_PerspectiveTransform(mat gocv.Mat, src []image.Point, dst []image.Point) gocv.Mat {
	transform := gocv.GetPerspectiveTransform(src, dst)
	defer transform.Close()
	return gocv.WarpPerspective(mat, transform, image.Pt(mat.Cols(), mat.Rows()))
}

// ============================================================
// 图像滤波
// ============================================================

func OCV_GaussianBlur(mat gocv.Mat, kernelSize int, sigma float64) gocv.Mat {
	return gocv.GaussianBlur(mat, image.Pt(kernelSize, kernelSize), sigma, 0, gocv.BorderDefault)
}

func OCV_MedianBlur(mat gocv.Mat, kernelSize int) gocv.Mat {
	return gocv.MedianBlur(mat, kernelSize)
}

func OCV_BilateralFilter(mat gocv.Mat, diameter int, sigmaColor float64, sigmaSpace float64) gocv.Mat {
	return gocv.BilateralFilter(mat, diameter, sigmaColor, sigmaSpace)
}

func OCV_Blur(mat gocv.Mat, kernelSize int) gocv.Mat {
	return gocv.Blur(mat, image.Pt(kernelSize, kernelSize))
}

// ============================================================
// 形态学操作
// ============================================================

func OCV_Erode(mat gocv.Mat, kernelSize int) gocv.Mat {
	kernel := gocv.GetStructuringElement(gocv.MorphRect, image.Pt(kernelSize, kernelSize))
	defer kernel.Close()
	return gocv.Erode(mat, kernel)
}

func OCV_Dilate(mat gocv.Mat, kernelSize int) gocv.Mat {
	kernel := gocv.GetStructuringElement(gocv.MorphRect, image.Pt(kernelSize, kernelSize))
	defer kernel.Close()
	return gocv.Dilate(mat, kernel)
}

func OCV_MorphOpen(mat gocv.Mat, kernelSize int) gocv.Mat {
	kernel := gocv.GetStructuringElement(gocv.MorphRect, image.Pt(kernelSize, kernelSize))
	defer kernel.Close()
	return gocv.MorphologyEx(mat, gocv.MorphOpen, kernel)
}

func OCV_MorphClose(mat gocv.Mat, kernelSize int) gocv.Mat {
	kernel := gocv.GetStructuringElement(gocv.MorphRect, image.Pt(kernelSize, kernelSize))
	defer kernel.Close()
	return gocv.MorphologyEx(mat, gocv.MorphClose, kernel)
}

func OCV_MorphGradient(mat gocv.Mat, kernelSize int) gocv.Mat {
	kernel := gocv.GetStructuringElement(gocv.MorphRect, image.Pt(kernelSize, kernelSize))
	defer kernel.Close()
	return gocv.MorphologyEx(mat, gocv.MorphGradient, kernel)
}

func OCV_MorphTopHat(mat gocv.Mat, kernelSize int) gocv.Mat {
	kernel := gocv.GetStructuringElement(gocv.MorphRect, image.Pt(kernelSize, kernelSize))
	defer kernel.Close()
	return gocv.MorphologyEx(mat, gocv.MorphTopHat, kernel)
}

func OCV_MorphBlackHat(mat gocv.Mat, kernelSize int) gocv.Mat {
	kernel := gocv.GetStructuringElement(gocv.MorphRect, image.Pt(kernelSize, kernelSize))
	defer kernel.Close()
	return gocv.MorphologyEx(mat, gocv.MorphBlackHat, kernel)
}

// ============================================================
// 边缘检测
// ============================================================

func OCV_Canny(mat gocv.Mat, lowThreshold float64, highThreshold float64) gocv.Mat {
	return gocv.Canny(mat, lowThreshold, highThreshold)
}

func OCV_Sobel(mat gocv.Mat, kernelSize int) gocv.Mat {
	result := gocv.NewMat()
	dx := gocv.Sobel(mat, gocv.MatTypeCV64F, 1, 0, kernelSize, 1, 0, gocv.BorderDefault)
	defer dx.Close()
	dy := gocv.Sobel(mat, gocv.MatTypeCV64F, 0, 1, kernelSize, 1, 0, gocv.BorderDefault)
	defer dy.Close()
	gocv.AddWeighted(dx, 0.5, dy, 0.5, 0, &result)
	return result
}

func OCV_Laplacian(mat gocv.Mat, kernelSize int) gocv.Mat {
	return gocv.Laplacian(mat, gocv.MatTypeCV64F, kernelSize, 1, 0, gocv.BorderDefault)
}

func OCV_Scharr(mat gocv.Mat) gocv.Mat {
	result := gocv.NewMat()
	dx := gocv.Scharr(mat, gocv.MatTypeCV64F, 1, 0, 1, 0, gocv.BorderDefault)
	defer dx.Close()
	dy := gocv.Scharr(mat, gocv.MatTypeCV64F, 0, 1, 1, 0, gocv.BorderDefault)
	defer dy.Close()
	gocv.AddWeighted(dx, 0.5, dy, 0.5, 0, &result)
	return result
}

// ============================================================
// 阈值处理
// ============================================================

func OCV_Threshold(mat gocv.Mat, threshold float64, maxVal float64) gocv.Mat {
	result := gocv.NewMat()
	gocv.Threshold(mat, &result, threshold, maxVal, gocv.ThresholdBinary)
	return result
}

func OCV_AdaptiveThreshold(mat gocv.Mat, maxVal float64, blockSize int, c float64) gocv.Mat {
	result := gocv.NewMat()
	gocv.AdaptiveThreshold(mat, &result, maxVal, gocv.AdaptiveThresholdGaussian, gocv.ThresholdBinary, blockSize, c)
	return result
}

func OCV_ThresholdInv(mat gocv.Mat, threshold float64, maxVal float64) gocv.Mat {
	result := gocv.NewMat()
	gocv.Threshold(mat, &result, threshold, maxVal, gocv.ThresholdBinaryInv)
	return result
}

func OCV_ThresholdOtsu(mat gocv.Mat, maxVal float64) gocv.Mat {
	result := gocv.NewMat()
	gocv.Threshold(mat, &result, 0, maxVal, gocv.ThresholdOtsu)
	return result
}

// ============================================================
// 轮廓检测
// ============================================================

func OCV_FindContours(mat gocv.Mat) [][]image.Point {
	return gocv.FindContours(mat, gocv.RetrievalExternal, gocv.ChainApproxSimple)
}

func OCV_FindAllContours(mat gocv.Mat) [][]image.Point {
	return gocv.FindContours(mat, gocv.RetrievalTree, gocv.ChainApproxSimple)
}

func OCV_DrawContours(mat gocv.Mat, contours [][]image.Point, c color.RGBA, thickness int) gocv.Mat {
	result := mat.Clone()
	gocv.DrawContours(&result, contours, -1, c, thickness)
	return result
}

func OCV_ContourArea(contour []image.Point) float64       { return gocv.ContourArea(contour) }
func OCV_ArcLength(contour []image.Point, closed bool) float64 { return gocv.ArcLength(contour, closed) }
func OCV_BoundingRect(contour []image.Point) image.Rectangle  { return gocv.BoundingRect(contour) }
func OCV_MinEnclosingCircle(contour []image.Point) (image.Point, float64) {
	return gocv.MinEnclosingCircle(contour)
}

// ============================================================
// 模板匹配
// ============================================================

// OCV_MatchTemplate 在源图像中搜索模板图像的最佳匹配位置。
// 使用 TM_CCOEFF_NORMED 方法，返回匹配分数和位置信息。
func OCV_MatchTemplate(src gocv.Mat, template gocv.Mat) (float64, image.Point, error) {
	if src.Empty() || template.Empty() {
		return 0, image.Point{}, nil
	}
	resultW := src.Cols() - template.Cols() + 1
	resultH := src.Rows() - template.Rows() + 1
	if resultW <= 0 || resultH <= 0 {
		return 0, image.Point{}, nil
	}
	result := gocv.NewMatWithSize(resultH, resultW, gocv.MatTypeCV32F)
	defer result.Close()
	gocv.MatchTemplate(src, template, &result, gocv.TmCcoeffNormed, gocv.NewMat())
	_, maxConf, _, maxLoc := gocv.MinMaxLoc(result)
	return maxConf, maxLoc, nil
}

// OCV_MatchTemplateAll 在源图像中搜索所有匹配位置（置信度高于阈值）。
func OCV_MatchTemplateAll(src gocv.Mat, template gocv.Mat, threshold float64) ([]image.Point, []float64) {
	if src.Empty() || template.Empty() {
		return nil, nil
	}
	resultW := src.Cols() - template.Cols() + 1
	resultH := src.Rows() - template.Rows() + 1
	if resultW <= 0 || resultH <= 0 {
		return nil, nil
	}
	result := gocv.NewMatWithSize(resultH, resultW, gocv.MatTypeCV32F)
	defer result.Close()
	gocv.MatchTemplate(src, template, &result, gocv.TmCcoeffNormed, gocv.NewMat())

	var points []image.Point
	var confidences []float64
	for y := 0; y < result.Rows(); y++ {
		for x := 0; x < result.Cols(); x++ {
			if result.GetFloatAt(y, x) >= float32(threshold) {
				points = append(points, image.Pt(x, y))
				confidences = append(confidences, float64(result.GetFloatAt(y, x)))
			}
		}
	}
	return points, confidences
}

// ============================================================
// 人脸检测
// ============================================================

// OCV_FaceDetect 使用 Haar 级联分类器进行人脸检测。
func OCV_FaceDetect(mat gocv.Mat, cascadePath string) ([]image.Rectangle, error) {
	classifier := gocv.NewCascadeClassifier()
	defer classifier.Close()

	if !classifier.Load(cascadePath) {
		return nil, nil
	}
	return classifier.DetectMultiScale(mat), nil
}

// OCV_FaceDetectWithParams 带参数的人脸检测。
func OCV_FaceDetectWithParams(mat gocv.Mat, cascadePath string, scaleFactor float64, minNeighbors int, minSize image.Point) ([]image.Rectangle, error) {
	classifier := gocv.NewCascadeClassifier()
	defer classifier.Close()
	if !classifier.Load(cascadePath) {
		return nil, nil
	}
	return classifier.DetectMultiScaleWithParams(mat, scaleFactor, minNeighbors, 0, minSize), nil
}

// ============================================================
// 视频捕获
// ============================================================

// OCV_VideoCapture_Open 打开视频捕获设备或视频文件。
func OCV_VideoCapture_Open(source int) (*gocv.VideoCapture, error) {
	vc, err := gocv.OpenVideoCapture(source)
	if err != nil {
		return nil, err
	}
	return vc, nil
}

// OCV_VideoCapture_OpenFile 打开视频文件进行读取。
func OCV_VideoCapture_OpenFile(filePath string) (*gocv.VideoCapture, error) {
	vc, err := gocv.VideoCaptureFile(filePath)
	if err != nil {
		return nil, err
	}
	return vc, nil
}

// OCV_VideoCapture_Read 读取视频的下一帧。
func OCV_VideoCapture_Read(vc *gocv.VideoCapture) (gocv.Mat, bool) {
	mat := gocv.NewMat()
	ok := vc.Read(&mat)
	return mat, ok
}

// OCV_VideoCapture_Close 释放视频捕获资源。
func OCV_VideoCapture_Close(vc *gocv.VideoCapture) { vc.Close() }

// ============================================================
// 窗口显示（调试用）
// ============================================================

// OCV_Window_New 创建显示窗口。
func OCV_Window_New(name string) *gocv.Window { return gocv.NewWindow(name) }

// OCV_Window_Show 在窗口中显示 Mat 图像。
func OCV_Window_Show(w *gocv.Window, mat gocv.Mat) { w.IMShow(mat) }

// OCV_Window_Close 关闭窗口。
func OCV_Window_Close(w *gocv.Window) { w.Close() }

// OCV_WaitKey 等待按键并返回按键码。参数 0 表示无限等待。
func OCV_WaitKey(ms int) int { return gocv.WaitKey(ms) }

// ============================================================
// 绘制
// ============================================================

// OCV_DrawRectangle 在 Mat 上绘制矩形。
func OCV_DrawRectangle(mat *gocv.Mat, rect image.Rectangle, c color.RGBA, thickness int) {
	gocv.Rectangle(mat, rect, c, thickness)
}

// OCV_DrawCircle 在 Mat 上绘制圆形。
func OCV_DrawCircle(mat *gocv.Mat, center image.Point, radius int, c color.RGBA, thickness int) {
	gocv.Circle(mat, center, radius, c, thickness)
}

// OCV_DrawLine 在 Mat 上绘制直线。
func OCV_DrawLine(mat *gocv.Mat, pt1 image.Point, pt2 image.Point, c color.RGBA, thickness int) {
	gocv.Line(mat, pt1, pt2, c, thickness)
}

// OCV_DrawText 在 Mat 上绘制文字。
func OCV_DrawText(mat *gocv.Mat, text string, origin image.Point, scale float64, c color.RGBA, thickness int) {
	gocv.PutText(mat, text, origin, gocv.FontHersheyPlain, scale, c, thickness)
}