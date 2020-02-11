package main

import (
	"log"
	"math"
)

//FruchtermanReingold slkjd
func FruchtermanReingold() {
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

		grState.position[i].x = math.Min(math.Max(40, grState.position[i].x), float64(areaX-40))
		grState.position[i].y = math.Min(math.Max(40, grState.position[i].y), float64(areaY-40))
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

	delta.x = grState.position[i].x - grState.position[j].x
	delta.y = grState.position[i].y - grState.position[j].y

	modD := modDelta(delta)

	dispX := (delta.x / modD) * ((modD * modD) / k)
	dispY := (delta.y / modD) * ((modD * modD) / k)
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
