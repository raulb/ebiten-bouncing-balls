// Copyright 2018 The Ebiten Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	screenWidth  = 1792
	screenHeight = 1120
	doorWidth    = 10
	doorHeight   = 120
)

var door = ebiten.NewImage(doorWidth, doorHeight)
var wall = ebiten.NewImage(doorWidth, screenHeight)

var blackColor = color.RGBA{0, 0, 0, 0xff}
var redColor = color.RGBA{0xff, 0, 0, 0xff}
var blueColor = color.RGBA{0, 0, 0xff, 0xff}

func init() {
	// draw a bar in red, the one that's moving is in black which should be the same background color
	door.Fill(blueColor)
	wall.Fill(redColor)
}

type Game struct {
	x float64
	y float64
}

func (g *Game) Update() error {
	_, dy := ebiten.Wheel()
	//g.x += dx
	g.y += dy
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

	wallOps := &ebiten.DrawImageOptions{}
	wallOps.GeoM.Translate(0, 0)
	wallOps.GeoM.Translate(screenWidth/2, 0)

	screen.DrawImage(wall, wallOps)

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(g.x, g.y)
	op.GeoM.Translate(screenWidth/2, screenHeight/2)
	screen.DrawImage(door, op)

	ebitenutil.DebugPrint(screen,
		fmt.Sprintf("Move the red point by mouse wheel\n(%0.2f, %0.2f)", g.x, g.y))
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	g := &Game{x: 0.0, y: 0.0}

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Wheel (Ebitengine Demo)")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
