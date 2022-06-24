package main

import (
	"bytes"
	"fmt"
	"image/jpeg"
	"time"

	"github.com/icza/mjpeg"
)

func main() {
	fmt.Println(time.Now())

	const (
		W          = 1920
		H          = 1080
		FRAME_RATE = 60
		SECONDS    = 1
	)

	animation := NewAnimation(W, H)

	aw, err := mjpeg.New("video.avi", W, H, FRAME_RATE)
	if err != nil {
		panic(err)
	}

	for i := 0; i < FRAME_RATE*SECONDS; i++ {
		fmt.Print(i, " ")
		buf := &bytes.Buffer{}
		img := animation.next()
		err = jpeg.Encode(buf, img, nil)
		if err != nil {
			panic(err)
		}
		err = aw.AddFrame(buf.Bytes())
		if err != nil {
			panic(err)
		}
	}

	err = aw.Close()
	if err != nil {
		panic(err)
	}

	fmt.Println(time.Now())
}
