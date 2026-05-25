//go:build opencv

package main

import (
	"fmt"
	"image"
	"image/color"

	. "github.com/yuan71058/Efunc/utils"

	"gocv.io/x/gocv"
)

func 示例_视觉() {
	fmt.Println("===== OCV 视觉（找图）=====")

	ver := OCV_取版本()
	fmt.Println("OpenCV版本:", ver)

	OCV_找图示例()
	OCV_找图特征匹配示例()
	OCV_找图多尺度示例()
	OCV_找图边缘示例()
}

func OCV_找图示例() {
	fmt.Println("--- OCV_找图 / OCV_找图中心 ---")

	源图 := gocv.NewMatWithSize(200, 300, gocv.MatTypeCV8UC3)
	defer 源图.Close()
	gocv.Rectangle(&源图, image.Rect(50, 50, 100, 100), color.RGBA{0, 255, 0, 255}, -1)

	模板 := gocv.NewMatWithSize(50, 50, gocv.MatTypeCV8UC3)
	defer 模板.Close()
	gocv.Rectangle(&模板, image.Rect(0, 0, 50, 50), color.RGBA{0, 255, 0, 255}, -1)

	位置, 置信度, err := OCV_找图(源图, 模板, 0.8)
	fmt.Println("OCV_找图:", 位置, "置信度:", 置信度, "错误:", err)

	中心, 中心置信度, _ := OCV_找图中心(源图, 模板, 0.8)
	fmt.Println("OCV_找图中心:", 中心, "置信度:", 中心置信度)
	fmt.Println()
}

func OCV_找图特征匹配示例() {
	fmt.Println("--- OCV_找图SIFT / ORB / AKAZE ---")

	源图 := gocv.NewMatWithSize(200, 300, gocv.MatTypeCV8UC3)
	defer 源图.Close()
	gocv.Rectangle(&源图, image.Rect(50, 50, 120, 120), color.RGBA{255, 0, 0, 255}, -1)
	gocv.Circle(&源图, image.Pt(150, 100), 30, color.RGBA{0, 0, 255, 255}, -1)

	模板 := gocv.NewMatWithSize(50, 50, gocv.MatTypeCV8UC3)
	defer 模板.Close()
	gocv.Rectangle(&模板, image.Rect(0, 0, 50, 50), color.RGBA{255, 0, 0, 255}, -1)

	矩形, 匹配数, err := OCV_找图ORB(源图, 模板, 4)
	fmt.Println("OCV_找图ORB:", 矩形, "匹配数:", len(匹配数), "错误:", err)

	矩形, 匹配数, err = OCV_找图AKAZE(源图, 模板, 4)
	fmt.Println("OCV_找图AKAZE:", 矩形, "匹配数:", len(匹配数), "错误:", err)

	矩形, 匹配数, err = OCV_找图SIFT(源图, 模板, 4)
	fmt.Println("OCV_找图SIFT:", 矩形, "匹配数:", len(匹配数), "错误:", err)
	fmt.Println()
}

func OCV_找图多尺度示例() {
	fmt.Println("--- OCV_找图多尺度 ---")

	源图 := gocv.NewMatWithSize(200, 300, gocv.MatTypeCV8UC3)
	defer 源图.Close()
	gocv.Rectangle(&源图, image.Rect(50, 50, 100, 100), color.RGBA{0, 255, 0, 255}, -1)

	模板 := gocv.NewMatWithSize(50, 50, gocv.MatTypeCV8UC3)
	defer 模板.Close()
	gocv.Rectangle(&模板, image.Rect(0, 0, 50, 50), color.RGBA{0, 255, 0, 255}, -1)

	位置, 置信度, 比例, err := OCV_找图多尺度(源图, 模板, 0.5, 2.0, 0.1, 0.8)
	fmt.Println("OCV_找图多尺度:", 位置, "置信度:", 置信度, "比例:", 比例, "错误:", err)
	fmt.Println()
}

func OCV_找图边缘示例() {
	fmt.Println("--- OCV_找图边缘 ---")

	源图 := gocv.NewMatWithSize(200, 300, gocv.MatTypeCV8UC3)
	defer 源图.Close()
	gocv.Rectangle(&源图, image.Rect(50, 50, 100, 100), color.RGBA{0, 255, 0, 255}, -1)

	模板 := gocv.NewMatWithSize(50, 50, gocv.MatTypeCV8UC3)
	defer 模板.Close()
	gocv.Rectangle(&模板, image.Rect(0, 0, 50, 50), color.RGBA{0, 255, 0, 255}, -1)

	位置, 置信度, err := OCV_找图边缘(源图, 模板, 50, 150, 0.5)
	fmt.Println("OCV_找图边缘:", 位置, "置信度:", 置信度, "错误:", err)
	fmt.Println()
}