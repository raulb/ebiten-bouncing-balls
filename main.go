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

// TODO
// Make balls rounded
// make a bunch of balls
// Make balls move

const (
	screenWidth  = 1152
	screenHeight = 720
	doorWidth    = 10
	doorHeight   = 120

	ballWidth  = 20
	ballHeight = 20

	frameOX     = 0
	frameOY     = 32
	frameWidth  = 32
	frameHeight = 32
	frameCount  = 8

	stepCount = 2
)

var door = ebiten.NewImage(doorWidth, doorHeight)
var wall = ebiten.NewImage(doorWidth, screenHeight)
var ball = ebiten.NewImage(ballWidth, ballHeight)

var blackColor = color.RGBA{0, 0, 0, 0xff}
var redColor = color.RGBA{0xff, 0, 0, 0xff}
var blueColor = color.RGBA{0, 0, 0xff, 0xff}

func init() {
	// draw a bar in red, the one that's moving is in black which should be the same background color
	door.Fill(blackColor)
	wall.Fill(redColor)
	ball.Fill(blueColor)
}

type Game struct {
	x                  float64
	y                  float64
	ballX              float64
	ballY              float64
	directionXPositive bool
	directionYPositive bool

	ballFrameCount int
	count          int
}

func (g *Game) Update() error {
	_, dy := ebiten.Wheel()
	//g.x += dx
	g.y += dy

	// Update ball position based on direction
	if g.directionXPositive {
		g.ballX += stepCount
	} else {
		g.ballX -= stepCount
	}

	if g.directionYPositive {
		g.ballY += stepCount
	} else {
		g.ballY -= stepCount
	}

	// Check for collisions with the screen boundaries and reverse direction if necessary
	if g.ballX > screenWidth-ballWidth {
		g.ballX = screenWidth - ballWidth
		g.directionXPositive = false
	} else if g.ballX < 0 {
		g.ballX = 0
		g.directionXPositive = true
	}

	if g.ballY > screenHeight-ballHeight {
		g.ballY = screenHeight - ballHeight
		g.directionYPositive = false
	} else if g.ballY < 0 {
		g.ballY = 0
		g.directionYPositive = true
	}

	return nil
}

func (g *Game) drawWall(screen *ebiten.Image) {
	wallOps := &ebiten.DrawImageOptions{}
	wallOps.GeoM.Translate(0, 0)
	wallOps.GeoM.Translate(screenWidth/2, 0)

	screen.DrawImage(wall, wallOps)
}

func (g *Game) drawDoor(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(g.x, g.y)
	op.GeoM.Translate(screenWidth/2, screenHeight/2)
	screen.DrawImage(door, op)
}

func (g *Game) drawBall(screen *ebiten.Image) {

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(g.ballY, g.ballY)
	op.GeoM.Translate(20, 20)

	if g.ballFrameCount > 10 {
		g.ballFrameCount = 0
	}

	g.ballFrameCount++
	screen.DrawImage(ball, op)
	//op.GeoM.Translate(-float64(frameWidth)/2, -float64(frameHeight)/2)
	//op.GeoM.Translate(screenWidth/2, screenHeight/2)
	//i := (g.count / 5) % frameCount
	//sx, sy := frameOX+i*frameWidth, frameOY

	//screen.DrawImage(ball.SubImage(image.Rect(sx, sy, sx+frameWidth, sy+frameHeight)).(*ebiten.Image), op)

}

func (g *Game) Draw(screen *ebiten.Image) {
	g.drawWall(screen)
	g.drawDoor(screen)

	g.drawBall(screen)

	//ebitenutil.DebugPrint(screen,
	//	fmt.Sprintf("Move the red point by mouse wheel\n(%0.2f, %0.2f)", g.x, g.y))
	ebitenutil.DebugPrint(screen,
		fmt.Sprintf("Move the red point by mouse wheel\n(%0.2f, %0.2f, %d, %d, directionX: %b, directionY: %b)",
			g.ballX, g.ballX, screenWidth, screenHeight, g.directionXPositive, g.directionYPositive,
		))
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	g := &Game{x: 0.0, y: 0.0}

	g.directionXPositive = true
	g.directionYPositive = true

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Wheel (Ebitengine Demo)")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
