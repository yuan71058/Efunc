//go:build opencv

package main

import (
	"fmt"
	"image"
	"image/color"

	. "github.com/yuan71058/Efunc/utils"

	"gocv.io/x/gocv"
)

// 示例_视觉 OpenCV 计算机视觉功能演示。
// 需要 opencv 编译标签和 gocv 依赖。
// 运行：go run -tags opencv ./cmd/demo
func 示例_视觉() {
	fmt.Println("===== OCV 视觉（找图）=====")

	ver := OCV_取版本()
	fmt.Println("OpenCV版本:", ver)

	OCV_找图示例()
	OCV_找图特征匹配示例()
	OCV_找图多尺度示例()
	OCV_找图边缘示例()
}

// OCV_找图示例 演示模板匹配（Template Matching）。
// 在源图中查找与模板匹配的位置，返回坐标和置信度。
func OCV_找图示例() {
	fmt.Println("--- OCV_找图 / OCV_找图中心 ---")

	// 创建源图：300x200，包含一个绿色矩形区域
	源图 := gocv.NewMatWithSize(200, 300, gocv.MatTypeCV8UC3)
	defer 源图.Close()
	gocv.Rectangle(&源图, image.Rect(50, 50, 100, 100), color.RGBA{0, 255, 0, 255}, -1)

	// 创建模板：50x50，与源图中的矩形一致
	模板 := gocv.NewMatWithSize(50, 50, gocv.MatTypeCV8UC3)
	defer 模板.Close()
	gocv.Rectangle(&模板, image.Rect(0, 0, 50, 50), color.RGBA{0, 255, 0, 255}, -1)

	// 模板匹配，置信度阈值 0.8
	位置, 置信度, err := OCV_找图(源图, 模板, 0.8)
	fmt.Println("OCV_找图:", 位置, "置信度:", 置信度, "错误:", err)

	// 返回匹配区域中心点
	中心, 中心置信度, _ := OCV_找图中心(源图, 模板, 0.8)
	fmt.Println("OCV_找图中心:", 中心, "置信度:", 中心置信度)
	fmt.Println()
}

// OCV_找图特征匹配示例 演示特征点匹配（SIFT/ORB/AKAZE）。
// 适用于旋转、缩放后的图像匹配，比模板匹配更鲁棒。
func OCV_找图特征匹配示例() {
	fmt.Println("--- OCV_找图SIFT / ORB / AKAZE ---")

	// 创建源图：包含矩形和圆形
	源图 := gocv.NewMatWithSize(200, 300, gocv.MatTypeCV8UC3)
	defer 源图.Close()
	gocv.Rectangle(&源图, image.Rect(50, 50, 120, 120), color.RGBA{255, 0, 0, 255}, -1)
	gocv.Circle(&源图, image.Pt(150, 100), 30, color.RGBA{0, 0, 255, 255}, -1)

	// 创建模板：包含与源图中匹配的矩形
	模板 := gocv.NewMatWithSize(50, 50, gocv.MatTypeCV8UC3)
	defer 模板.Close()
	gocv.Rectangle(&模板, image.Rect(0, 0, 50, 50), color.RGBA{255, 0, 0, 255}, -1)

	// ORB 特征匹配
	矩形, 匹配数, err := OCV_找图ORB(源图, 模板, 4)
	fmt.Println("OCV_找图ORB:", 矩形, "匹配数:", len(匹配数), "错误:", err)

	// AKAZE 特征匹配
	矩形, 匹配数, err = OCV_找图AKAZE(源图, 模板, 4)
	fmt.Println("OCV_找图AKAZE:", 矩形, "匹配数:", len(匹配数), "错误:", err)

	// SIFT 特征匹配
	矩形, 匹配数, err = OCV_找图SIFT(源图, 模板, 4)
	fmt.Println("OCV_找图SIFT:", 矩形, "匹配数:", len(匹配数), "错误:", err)
	fmt.Println()
}

// OCV_找图多尺度示例 演示多尺度模板匹配。
// 在不同缩放比例下搜索模板，返回最佳匹配位置、置信度和缩放比例。
func OCV_找图多尺度示例() {
	fmt.Println("--- OCV_找图多尺度 ---")

	源图 := gocv.NewMatWithSize(200, 300, gocv.MatTypeCV8UC3)
	defer 源图.Close()
	gocv.Rectangle(&源图, image.Rect(50, 50, 100, 100), color.RGBA{0, 255, 0, 255}, -1)

	模板 := gocv.NewMatWithSize(50, 50, gocv.MatTypeCV8UC3)
	defer 模板.Close()
	gocv.Rectangle(&模板, image.Rect(0, 0, 50, 50), color.RGBA{0, 255, 0, 255}, -1)

	// 缩放范围 0.5 ~ 2.0，步长 0.1
	位置, 置信度, 比例, err := OCV_找图多尺度(源图, 模板, 0.5, 2.0, 0.1, 0.8)
	fmt.Println("OCV_找图多尺度:", 位置, "置信度:", 置信度, "比例:", 比例, "错误:", err)
	fmt.Println()
}

// OCV_找图边缘示例 演示基于边缘检测的模板匹配。
// 使用 Canny 边缘检测后再进行模板匹配，对光照变化不敏感。
func OCV_找图边缘示例() {
	fmt.Println("--- OCV_找图边缘 ---")

	源图 := gocv.NewMatWithSize(200, 300, gocv.MatTypeCV8UC3)
	defer 源图.Close()
	gocv.Rectangle(&源图, image.Rect(50, 50, 100, 100), color.RGBA{0, 255, 0, 255}, -1)

	模板 := gocv.NewMatWithSize(50, 50, gocv.MatTypeCV8UC3)
	defer 模板.Close()
	gocv.Rectangle(&模板, image.Rect(0, 0, 50, 50), color.RGBA{0, 255, 0, 255}, -1)

	// Canny 边缘检测阈值：低 50，高 150
	位置, 置信度, err := OCV_找图边缘(源图, 模板, 50, 150, 0.5)
	fmt.Println("OCV_找图边缘:", 位置, "置信度:", 置信度, "错误:", err)
	fmt.Println()
}