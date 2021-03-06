package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"image/color"
	"log"
	//"math"
	"math/rand"
	"time"
)

//Globals
const windowTitle = "goGraphDraw"
const scale = 1
const windowWidth = 1024
const windowHeight = 768
const assetsDir = "assets"

var sprite *ebiten.Image

type state struct {
	position []vec
	arcs     [][]int //index dominates next index
	disp     []vec   //displasement vector
}

type vec struct {
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
		//initGraph()
		mode = 1
	}

	if ebiten.IsDrawingSkipped() {
		return nil
	}
	//Do stuff goes here
	//physics
	screen.Fill(color.RGBA{0xff, 0xff, 0xff, 255})
	FruchtermanReingold()
	drawArc(screen)
	drawSprite(screen)
	//updatePosition()
	return nil
}

func main() {
	initGraph()

	if err := ebiten.Run(update, windowWidth, windowHeight, scale, windowTitle); err != nil {
		log.Fatal(err)

	}
}

func initGraph() {
	vertImg, _, err := ebitenutil.NewImageFromFile("assets/goBowl.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	vertNum := 8
	grState.arcs = make([][]int, vertNum, vertNum)
	grState.position = make([]vec, vertNum, vertNum)
	grState.disp = make([]vec, vertNum, vertNum)

	grState.arcs[0] = append(grState.arcs[0], 1, 2)
	grState.arcs[1] = append(grState.arcs[1], 2, 3, 6, 7)
	grState.arcs[2] = append(grState.arcs[2], 3)
	grState.arcs[3] = append(grState.arcs[3], 0)
	grState.arcs[4] = append(grState.arcs[4], 5)
	grState.arcs[5] = append(grState.arcs[5], 3)
	grState.arcs[6] = append(grState.arcs[6], 5)
	grState.arcs[7] = append(grState.arcs[7], 5)
	fmt.Println(grState.arcs)
	randPos()

	sprite = vertImg

}

func drawArc(screen *ebiten.Image) {
	for i := 0; i < len(grState.arcs); i++ {
		sizeX, _ := sprite.Size()
		offset := float64(sizeX) / 2
		var from vec
		var to vec
		from = grState.position[i]
		for j := 0; j < len(grState.arcs[i]); j++ {
			to = grState.position[grState.arcs[i][j]]
			ebitenutil.DrawLine(screen, from.x+offset, from.y+offset, to.x+offset, to.y+offset, color.Black)
		}

	}
}
func drawSprite(screen *ebiten.Image) {
	for i := 0; i < len(grState.position); i++ {
		opts := &ebiten.DrawImageOptions{}
		opts.GeoM.Translate(float64(grState.position[i].x), float64(grState.position[i].y))
		screen.DrawImage(sprite, opts)
	}
}

//random seeding
func randPos() {
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < len(grState.position); i++ {
		grState.position[i].x = float64(rand.Intn(windowWidth))
		grState.position[i].y = float64(rand.Intn(windowHeight))
	}
}
