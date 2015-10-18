package meow_ml_kmeans

import (
	"math"
	"ashumeow/meow_ml_kmeans/meow_data_structures"
	"ashumeow/meow_ml_kmeans/meow_ml"
)

type meowKmeans struct {
	dots *meow_data_structures.meowArrayList
	num_of_clusters int
	meowta float64
}

// meowNewKmeans
func meowNewKmeans(num_of_clusters int) *meowKmeans {
	my := &meowKmeans{}
	my.dots = meow_data_structures.meowNewArrayList()
	my.num_of_clusters = num_of_clusters
	my.meowta = 0.001
	return my
}

// setMeowta
func (my *meowKmeans) setMeowta(meowta float64) {
	my.meowta = meowta
}

// meowAddDots
func (my *meowKmeans) meowAddDots(dot *meow_ml.meowPoint) {
	my.dots.meowAdd(dot)
}

// meowAddSliceDots
func (my *meowKmeans) meowAddSliceDots(stuffs []float64) {
	my.dots.meowAdd(meow_ml.meowNewKmeans(stuffs))
}

// clustering
func (my *meowKmeans) meow_cluster() *meow_data_structures.meowArrayList {
	if(my.num_of_clusters == 1) {
		panic("Pick more than one cluster ;)")
	}
	meow_clusters := meow_data_structures.meowArrayList()
	uniqueBlackholes := meow_data_structures.meowNewHashSet()
	for x := 0; x < my.num_of_clusters; x++ {
		randomBlackhole := my.dots.meowSample().(*meow_ml.meowPoint)
		for uniqueBlackholes.meowRegisters(randomBlackhole) {
			randomBlackhole = my.dots.meowSample().(*meow_ml.meowPoint)
		}
		uniqueBlackholes.meowAdd(randomBlackhole)
		meow_cluster := meow_ml.meowNewCluster(randomBlackhole)
		meow_clusters.meowAdd(meow_cluster)
	}

	for {
	// determining nearest meow_cluster for assigning a dot
	for x := 0; x < my.dots.meowLen(); x++ {
		meowSmallDist := math.MaxFloat64
		var meowNearCluster *meow_ml.meowCluster
		dot := my.dots.meowFetch(x).(*meow_ml.meowPoint)
		for xx := 0; xx < meow_clusters.meowLen(); xx++ {
			meow_cluster := meow_clusters.meowFetch(xx).(*meow_ml.meowCluster)
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
		meow_cluster := meow_clusters.meowFetch(x).(*meow_ml.meowCluster)
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
			meow_cluster := meow_clusters.meowFetch(x).(*meow_ml.meowCluster)
			meow_cluster.dots().meowReset()
		}
	} }
	return meow_clusters
}