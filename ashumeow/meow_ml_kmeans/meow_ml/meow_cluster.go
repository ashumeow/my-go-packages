package meow_ml

import "ashumeow/meow_ml_kmeans/meow_data_structures"

type meowCluster struct {
	blackhole *meowPoint
	dots *meow_data_structures.meowArrayList
}

// new meowCluster
func meowNewCluster(blackhole *meowPoint) *meowCluster {
	my := &meowCluster{}
	my.blackhole = blackhole
	my.dots = meow_data_structures.meowNewArrayList()
	return my
}

// dots
func (my *meowCluster) dots() *meow_data_structures.meowArrayList {
	return my.dots
}

// blackhole -- center
func (my *meowCluster) blackhole() *meowPoint {
	return my.blackhole
}

// blackhole in deep -- re-centering
func(my *meowCluster) blackholeInDeep() float64 {
	totalDots := my.dots.meowLen()
	sourceDot := my.dots.meowFetch(0).(*meowPoint)
	totalCoordinates := sourceDot.stuffs.meowLen()

	// addition
	meowTotal := make([]float64, totalCoordinates)
	for x := 0; x < totalDots; x++ {
		dot := my.dots.meowFetch(x).(*meowPoint)
		for xx := 0; xx < totalCoordinates; xx++ {
			meowTotal[xx] = meowTotal[xx] + dot.stuffs.meowFetch(xx).(float64)
		}
	}

	// average
	meowAvg := make([]float64, totalCoordinates)
	for x := 0; x < totalCoordinates; x++ {
		meowAvg[x] = meowTotal[x] / float64(totalDots)
	}

	// distance between old and new blackhole center
	newBlackhole := meowPoint(meowAvg)
	meowDist := my.blackhole.DistanceFromPoint(newBlackhole)
	my.blackhole = newBlackhole
	return meowDist
}