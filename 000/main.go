package main

import (
	"bytes"
	"image/jpeg"
	"math/rand"

	"github.com/fogleman/gg"
	"github.com/icza/mjpeg"
)

func main() {
	example1()
	example2()
	example3()
}
func example1() {
	c := gg.NewContext(1000, 500)
	c.SetRGB(1, 0, 0)
	c.DrawCircle(500, 250, 10)
	c.Fill()
	c.SavePNG("image.png")
}
func example2() {
	const (
		W          = 1000 // 解像度 : 高さ
		H          = 500  // 解像度 : 幅
		FRAME_RATE = 1    // 1秒間のフレーム数
	)
	aw, _ := mjpeg.New("single_frame.avi", W, H, FRAME_RATE)
	c := gg.NewContext(W, H)
	c.SetRGB(1, 0, 0)
	c.DrawCircle(rand.Float64()*W, rand.Float64()*H, 50)
	c.Fill()
	buf := &bytes.Buffer{}
	_ = jpeg.Encode(buf, c.Image(), nil)
	_ = aw.AddFrame(buf.Bytes())
	_ = aw.Close()
}
func example3() {
	const (
		W          = 1000 // 解像度 : 高さ
		H          = 500  // 解像度 : 幅
		FRAME_RATE = 1    // 1秒間のフレーム数
		SECONDS    = 10   // 動画の長さ(秒)
	)
	aw, _ := mjpeg.New("multi_frames.avi", W, H, FRAME_RATE)
	for i := 0; i < FRAME_RATE*SECONDS; i++ {
		c := gg.NewContext(W, H)
		c.SetRGB(1, 0, 0)
		c.DrawCircle(rand.Float64()*W, rand.Float64()*H, 50)
		c.Fill()
		buf := &bytes.Buffer{}
		_ = jpeg.Encode(buf, c.Image(), nil)
		_ = aw.AddFrame(buf.Bytes())
	}
	_ = aw.Close()
}
