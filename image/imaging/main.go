package main

import (
	"image"
	"image/color"
	"log"

	"github.com/disintegration/imaging"
)

func main() {
	// Open a test image.
	src, err := imaging.Open("./flowers.png")
	if err != nil {
		log.Fatalf("failed to open image: %v", err)
	}

	// Crop the original image to 300x300px size using the center anchor.
	// 将图像裁剪为 300x300 像素 使用中心锚点
	src = imaging.CropAnchor(src, 300, 300, imaging.Center)

	// 将裁剪后的图像调整为宽度 200 像素，同时保持纵横比
	// Resize the cropped image to width = 200px preserving the aspect ratio.
	src = imaging.Resize(src, 200, 0, imaging.Lanczos)

	// Create a blurred version of the image.
	// 创建模糊版本的图像
	img1 := imaging.Blur(src, 5)

	// Create a grayscale version of the image with higher contrast and sharpness.
	img2 := imaging.Grayscale(src)          // 灰度处理
	img2 = imaging.AdjustContrast(img2, 20) // 对比度的调整值20
	img2 = imaging.Sharpen(img2, 2)         // 锐化的强度

	// Create an inverted version of the image.
	// 输入图像 src 的颜色进行反转。每个像素的颜色值都会被转换为其补色
	img3 := imaging.Invert(src)

	// Create an embossed version of the image using a convolution filter.
	img4 := imaging.Convolve3x3(
		src,
		[9]float64{
			-1, -1, 0,
			-1, 1, 1,
			0, 1, 1,
		},
		nil,
	)

	// Create a new image and paste the four produced images into it.
	dst := imaging.New(400, 400, color.NRGBA{})
	dst = imaging.Paste(dst, img1, image.Pt(0, 0))
	dst = imaging.Paste(dst, img2, image.Pt(0, 200))
	dst = imaging.Paste(dst, img3, image.Pt(200, 0))
	dst = imaging.Paste(dst, img4, image.Pt(200, 200))

	// Save the resulting image as JPEG.
	err = imaging.Save(dst, "./out.jpg")
	if err != nil {
		log.Fatalf("failed to save image: %v", err)
	}
}
