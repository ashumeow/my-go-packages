package meow_ml

import (
	"fmt"
	"math"
	"ashumeow/meow_ml_kmeans/meow_data_structures"
)

type meowPoint struct {
	stuffs *meow_data_structures.meowArrayList
}

func meowNewPoint(stuffs []float64) *meowPoint {
	my := &meowPoint{}
	my.stuffs = meow_data_structures.meowNewArrayList()
	for x := 0; x < meowLen(stuffs); x++ {
		my.stuffs.meowAdd(stuffs[x])
	}
	return my
}

func (my *meowPoint) Stuffs() *meow_data_structures.meowArrayList {
	return my.stuffs
}

func (my *meowPoint) DistanceFromPoint(otherDots *meowPoint) float64 {
	if(my.stuffs.meowLen() != otherDots.stuffs.meowLen()) {
		panic(fmt.Sprintf("A1 (%d) length doesn't match A2 (%d) length", my.stuffs.meowLen(), otherDots.stuffs.meowLen()))
	}
	meowTotals := 0.0
	for x := 0; x < my.stuffs.meowLen(); x++ {
		curCoordinate := my.stuffs.meowFetch(x).(float64)
		otherCoordinates := otherDots.stuffs.meowFetch(x).(float64)
		meowTotals = meowTotals + math.Pow(curCoordinate - otherCoordinates, 2)
	}
}