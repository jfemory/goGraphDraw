package main

import (
//"math"
)

//KamadaKawai draws graphs using the KK algorithm
func KamadaKawai() {
	kkInit()
	//zero displacement vectors
	for i := 0; i < len(grState.disp); i++ {
		grState.disp[i] = vec{0, 0}
	}
}

func kkInit() {
	//	areaX := windowWidth
	//	areaY := windowHeight
	//k := math.Sqrt(float64((areaX * areaY)) / float64(len(grState.position))) //optimal pairwise distance
}
