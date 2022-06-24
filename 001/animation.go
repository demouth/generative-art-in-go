package main

import (
	"image"
	"math"
	"math/rand"

	"github.com/fogleman/gg"
)

type Particle struct {
	x  float64
	y  float64
	ax float64
	ay float64
	r  float64
}

func (p *Particle) reset(a Animation) {
	p.r = 2
	pow := rand.Float64()
	angle := math.Pi * 2 * rand.Float64()
	// 中央に集中するように初期位置を決める
	p.x = float64(a.width / 2)
	p.y = float64(a.height / 2)
	p.x += math.Cos(angle) * (float64(a.width) / 2) * pow * 1.3  // 画面の端に行き渡るようにちょっとだけ定数を掛ける
	p.y += math.Sin(angle) * (float64(a.height) / 2) * pow * 1.3 // 画面の端に行き渡るようにちょっとだけ定数を掛ける
	p.ax = (0.5 - rand.Float64()) * 2 * 1.5
	p.ay = (0.5 - rand.Float64()) * 2 * 1.5
}
func (p *Particle) draw(c *gg.Context) {
	c.SetRGB(1, 1, 1)
	c.DrawCircle(p.x, p.y, 1)
	c.Fill()
}
func (p *Particle) move() {
	p.x = p.x + p.ax
	p.y = p.y + p.ay
}
func (p *Particle) outOfBounds(width, height int) bool {
	if p.x < 0 {
		return true
	} else if p.y < 0 {
		return true
	} else if p.x > float64(width) {
		return true
	} else if p.y > float64(height) {
		return true
	}
	return false
}

type Animation struct {
	particles []Particle
	width     int
	height    int
}

func NewAnimation(width, height int) *Animation {
	const NUM_PARTICLES = 20000
	a := &Animation{}
	a.width = width
	a.height = height
	a.particles = make([]Particle, NUM_PARTICLES)
	for i := 0; i < len(a.particles); i++ {
		temp := Particle{}
		temp.reset(*a)

		a.particles[i] = temp
	}
	return a
}

func (a *Animation) next() image.Image {
	c := gg.NewContext(a.width, a.height)

	// particleを動かす
	for i, p := range a.particles {
		p.move()
		if p.outOfBounds(a.width, a.height) {
			p.reset(*a)
		}
		a.particles[i] = p
	}

	// ジョイントをつなげる
	c.SetLineWidth(1)
	threshold := float64(40)
	for x := 0; x < len(a.particles); x++ {
		for y := x + 1; y < len(a.particles); y++ {
			xp := a.particles[x]
			yp := a.particles[y]
			diffX := xp.x - yp.x
			diffY := xp.y - yp.y
			diff := math.Sqrt(diffX*diffX + diffY*diffY)
			if diff < threshold {
				// 色を決める
				per := 1 - diff/threshold
				c.SetRGBA(1, 1, 1, per)

				a.particles[x] = xp
				a.particles[y] = yp
				c.DrawLine(xp.x, xp.y, yp.x, yp.y)
				c.Stroke()
			}
		}
	}

	// パーティクルを描画
	for _, p := range a.particles {
		p.draw(c)
	}

	return c.Image()
}
