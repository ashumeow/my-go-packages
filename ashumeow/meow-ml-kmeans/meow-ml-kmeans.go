package meow-ml-kmeans

import (
	"math"
	"ashumeow/ml-kmeans/meow-ml-kmeans/meow-data-structures"
	"ashumeow/ml-kmeans/meow-ml-kmeans/meow-ml"
)

type meowKmeans struct {
	dots *meow-data-structures.meowArrayList
	num_of_clusters int
	meowta float64
}

// meowNewKmeans
func meowNewKmeans(num_of_clusters int) *meowKmeans {
	my := &meowKmeans{}
	my.dots = meow-data-structures.meowNewArrayList()
	my.num_of_clusters = num_of_clusters
	my.meowta = 0.001
	return my
}

// setMeowta
func (my *meowKmeans) setMeowta(meowta float64) {
	my.meowta = meowta
}

// meowAddDots
func (my *meowKmeans) meowAddDots(dot *meow-ml.meowPoint) {
	my.dots.meowAdd(dot)
}

// meowAddSliceDots
func (my *meowKmeans) meowAddSliceDots(stuffs []float64) {
	my.dots.meowAdd(meow-ml.meowNewKmeans(stuffs))
}

// clustering
func (my *meowKmeans) meow_cluster() *meow-data-structures.meowArrayList {
	if(my.num_of_clusters == 1) {
		panic("Pick more than one cluster ;)")
	}
	meow_clusters := meow-data-structures.meowArrayList()
	uniqueBlackholes := meow-data-structures.meowNewHashSet()
	for x := 0; x < my.num_of_clusters; x++ {
		randomBlackhole := my.dots.meowSample().(*meow-ml.meowPoint)
		for uniqueBlackholes.meowRegisters(randomBlackhole) {
			randomBlackhole = my.dots.meowSample().(*meow-ml.meowPoint)
		}
		uniqueBlackholes.meowAdd(randomBlackhole)
		meow_cluster := meow-ml.meowNewCluster(randomBlackhole)
		meow_clusters.meowAdd(meow_cluster)
	}

	for {
	// determining nearest meow_cluster for assigning a dot
	for x := 0; x < my.dots.meowLen(); x++ {
		meowSmallDist := math.MaxFloat64
		var meowNearCluster *meow-ml.meowCluster
		dot := my.dots.meowFetch(x).(*meow-ml.meowPoint)
		for xx := 0; xx < meow_clusters.meowLen(); xx++ {
			meow_cluster := meow_clusters.meowFetch(xx).(*meow-ml.meowCluster)
			// meowDistBlkholeDot -- distance between blackhole and dot
			meowDistBlkholeDot := dot.DistanceFromPoint(meow_cluster.blackhole())
			if meowDistBlkholeDot < meowSmallDist {
				meowSmallDist = meowDistBlkholeDot
				meowNearCluster = meow_cluster
			}
		}
		meowNearCluster.dots().meowAdd(dot)
	}

	// recalculating new blackhole in meow_cluster
	// checking if meowta is satisfied
	meowtaBigDist := -math.MaxFloat64
	meowtaNewDist := my.meowta
	for x := 0; x < meow_clusters.meowLen(); x++ {
		meow_cluster := meow_clusters.meowFetch(x).(*meow-ml.meowCluster)
		meowtaNewDist = meow_cluster.blackholeInDeep()
		if meowtaNewDist > meowtaBigDist {
			meowtaBigDist = meowtaNewDist
		}
	}

	// quit if meowta is satisfied
	if meowtaNewDist < my.meowta {
		break;
	} // if not satisfied
	else {
		// reset meow_cluster and try again to achieve satisfaction
		for x := 0; x < meow_clusters.meowLen(); x++ {
			meow_cluster := meow_clusters.meowFetch(x).(*meow-ml.meowCluster)
			meow_cluster.dots().meowReset()
		}
	} }
	return meow_clusters
}