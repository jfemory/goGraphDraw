package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"image/color"
	"log"
	"math"
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
	spring()
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
	vertNum := 4
	grState.arcs = make([][]int, vertNum, vertNum)
	grState.position = make([]vec, vertNum, vertNum)
	grState.disp = make([]vec, vertNum, vertNum)

	grState.arcs[0] = append(grState.arcs[0], 1)
	grState.arcs[1] = append(grState.arcs[1], 2)
	grState.arcs[2] = append(grState.arcs[2], 3)
	grState.arcs[3] = append(grState.arcs[3], 0)

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

//Spring

func spring() {
	areaX := windowWidth
	areaY := windowHeight
	k := math.Sqrt(float64((areaX * areaY)) / float64(len(grState.position))) //optimal pairwise distance
	//zero displacement vectors
	for i := 0; i < len(grState.disp); i++ {
		grState.disp[i] = vec{0, 0}

	}

	for i := 0; i < len(grState.position); i++ {
		for j := 0; j < len(grState.position); j++ {
			if i != j {
				fReplusiveSpring(i, j, k) //two indices and optimal pairwise distance
			}
		}
	}
	for i := 0; i < len(grState.arcs); i++ {
		if len(grState.arcs[i]) > 0 {
			for j := 0; j < len(grState.arcs[i]); j++ {
				fAttractiveSpring(i, grState.arcs[i][j], k) //two indices and optimal pairwise distance
			}
		}
	}

	//update Position
	for i := 0; i < len(grState.position); i++ {
		modD := modDelta(grState.disp[i])
		grState.position[i].x += (grState.disp[i].x / modD)
		grState.position[i].y += (grState.disp[i].y / modD)

		grState.position[i].x = math.Min(math.Max(0, grState.position[i].x), float64(areaX))
		grState.position[i].y = math.Min(math.Max(0, grState.position[i].y), float64(areaY))
	}
}

func fReplusiveSpring(i, j int, k float64) {
	var iDelta vec
	iDelta.x = grState.position[i].x - grState.position[j].x
	iDelta.y = grState.position[i].y - grState.position[j].y

	displace(iDelta, i, k)
}

func fAttractiveSpring(i, j int, k float64) {
	var delta vec
	fmt.Println(i, j)
	fmt.Println(grState.position[i])
	fmt.Println(grState.position[j])
	delta.x = grState.position[i].x - grState.position[j].x
	delta.y = grState.position[i].y - grState.position[j].y
	fmt.Println("delta")
	fmt.Println(delta)
	modD := modDelta(delta)

	dispX := (delta.x / modD) * ((modD * modD) / k)
	dispY := (delta.y / modD) * ((modD * modD) / k)
	fmt.Println("DispX, Y")
	fmt.Println(dispX)
	fmt.Println(dispY)
	grState.disp[i].x -= dispX
	grState.disp[i].y -= dispY
	grState.disp[j].x += dispX
	grState.disp[j].y += dispY
}

func displace(delta vec, index int, k float64) {
	modD := modDelta(delta)
	dispX := (delta.x / modD) * ((k * k) / modD)
	dispY := (delta.y / modD) * ((k * k) / modD)
	grState.disp[index].x += dispX
	grState.disp[index].y += dispY
}

func modDelta(delta vec) float64 {
	xSquared := delta.x * delta.x
	ySquared := delta.y * delta.y
	out := math.Sqrt(xSquared + ySquared)
	if math.IsNaN(out) == true {
		log.Fatal("NaN")
	}
	return out
}

//random seeding
func randPos() {
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < len(grState.position); i++ {
		grState.position[i].x = float64(rand.Intn(windowWidth))
		grState.position[i].y = float64(rand.Intn(windowHeight))
	}
}
