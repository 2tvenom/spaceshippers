package main

import "math/rand"
import "github.com/bennicholls/burl-E/burl"

type StarField struct {
	field         []int
	starFrequency int
	starShift     int
	view          *burl.TileView
	dirty         bool
}

//initializes a starfield twice the width of the screen
func NewStarField(starFrequency int, v *burl.TileView) (sf StarField) {
	sf.view = v
	w, h := v.Dims()
	sf.field = make([]int, w*h*2)
	sf.starFrequency = starFrequency
	sf.dirty = true
	for i := 0; i < len(sf.field); i++ {
		if rand.Intn(sf.starFrequency) == 0 {
			sf.field[i] = 1
		}
	}
	return
}

//moves the "camera" on the stars.
func (sf *StarField) Shift() {
	w, _ := sf.view.Dims()
	sf.starShift, _ = burl.ModularClamp(sf.starShift+1, 0, (w*2)-1)
	sf.dirty = true
}

//Draws the starfield, offset by the current starShift value.
func (sf *StarField) Draw() {
	if sf.dirty {
		sf.view.Reset()
		w, h := sf.view.Dims()
		for i := 0; i < w*h; i++ {
			if sf.field[(i/w)*w*2+(i%w+sf.starShift)%(w*2)] != 0 {
				sf.view.Draw(i%w, i/w, burl.GLYPH_ASTERISK, burl.COL_DARKGREY, burl.COL_BLACK)
			}
		}
		sf.dirty = false
	}
}
