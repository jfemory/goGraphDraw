package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"image/color"
	"log"
)

//Globals
const windowTitle = "goGraphDraw"
const scale = 3
const windowWidth = 480
const windowHeight = 320
const assetsDir = "assets"

var sprite *ebiten.Image

type state struct {
	position []pos
	arcs     [][]arc
}

type pos struct {
	x float64
	y float64
}
type arc struct {
	dominates int
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
		initGraph()
		fmt.Println(grState)
		mode = 1
	}

	if ebiten.IsDrawingSkipped() {
		return nil
	}
	//Do stuff goes here
	//physics
	screen.Fill(color.RGBA{0xff, 0xff, 0xff, 255})
	drawSprite(screen)
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
	vertImg, _, err := ebitenutil.NewImageFromFile("assets/goBowl.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	vertNum := 4
	grState.arcs = make([][]arc, vertNum, vertNum)
	grState.position = make([]pos, vertNum, vertNum)

	grState.arcs[0] = append(grState.arcs[0], arc{1}, arc{2})
	grState.arcs[1] = append(grState.arcs[1], arc{2})
	grState.arcs[2] = append(grState.arcs[2], arc{3})
	grState.arcs[3] = append(grState.arcs[3], arc{0})

	grState.position[0].x = 15.1
	grState.position[0].y = 20.1

	grState.position[1].x = 50
	grState.position[1].y = 50

	grState.position[2].x = 25
	grState.position[2].y = 35

	grState.position[3].x = 30
	grState.position[3].y = 25
	sprite = vertImg

}

func drawSprite(screen *ebiten.Image) {
	for i := 0; i < len(grState.position); i++ {
		opts := &ebiten.DrawImageOptions{}
		opts.GeoM.Translate(float64(grState.position[i].x), float64(grState.position[i].y))
		screen.DrawImage(sprite, opts)
	}
}
