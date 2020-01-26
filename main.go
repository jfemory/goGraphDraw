package main

import (
	"github.com/hajimehoshi/ebiten"
	"image/color"
	"log"
)

//Globals
const windowTitle = "goGraphDraw"
const scale = 3
const windowWidth = 480
const windowHeight = 320
const assetsDir = "assets"

type state struct {
	verts []int
}

type vertex struct {
	xPos float64
	yPos float64
	arcs []arc
}

type arc struct {
	from int
	to   int
}

var grState state
var mode = 0

//position holds positions, ordered x, y, as pairs of entries in the array
//var position [maxEntity * 2]float32

//update is the main loop of the ebiten engine. Core loop is here.
func update(screen *ebiten.Image) error {
	//gameMode 0 initializes assets
	if mode == -1 {
		panic("BadState")
	}
	if mode == 0 {

	}

	if ebiten.IsDrawingSkipped() {
		return nil
	}
	//Do stuff goes here
	//physics
	screen.Fill(color.RGBA{0xff, 0xff, 0xff, 255})
	updatePosition()
	return nil
}

func main() {
	if err := ebiten.Run(update, windowWidth, windowHeight, scale, windowTitle); err != nil {
		log.Fatal(err)

	}
}

func updatePosition() {

}

func initGraph() {

}
