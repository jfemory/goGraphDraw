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
const scale = 1
const windowWidth = 1024
const windowHeight = 768
const assetsDir = "assets"

var sprite *ebiten.Image

type state struct {
	position []pos
	arcs     [][]int //index dominates next index
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
	drawArc(screen)
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
	grState.arcs = make([][]int, vertNum, vertNum)
	grState.position = make([]pos, vertNum, vertNum)

	grState.arcs[0] = append(grState.arcs[0], 1, 2)
	grState.arcs[1] = append(grState.arcs[1], 2)
	grState.arcs[2] = append(grState.arcs[2], 3)
	grState.arcs[3] = append(grState.arcs[3], 0)

	grState.position[0].x = 151
	grState.position[0].y = 201

	grState.position[1].x = 500
	grState.position[1].y = 500

	grState.position[2].x = 250
	grState.position[2].y = 350

	grState.position[3].x = 300
	grState.position[3].y = 250
	sprite = vertImg

}

func drawArc(screen *ebiten.Image) {
	for i := 0; i < len(grState.arcs); i++ {
		sizeX, _ := sprite.Size()
		offset := float64(sizeX) / 2
		fmt.Println(offset)
		var from pos
		var to pos
		from = grState.position[i]
		for j := 0; j < len(grState.arcs[i]); j++ {
			to = grState.position[grState.arcs[0][1]]
		}
		ebitenutil.DrawLine(screen, from.x+offset, from.y+offset, to.x+offset, to.y+offset, color.Black)
	}
}
func drawSprite(screen *ebiten.Image) {
	for i := 0; i < len(grState.position); i++ {
		opts := &ebiten.DrawImageOptions{}
		opts.GeoM.Translate(float64(grState.position[i].x), float64(grState.position[i].y))
		screen.DrawImage(sprite, opts)
	}
}
