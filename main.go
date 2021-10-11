// begin
//    specify the number k of clustering to assign.
//    randomly initialize k centroids.
//    repeat
//       expectation: Assign each point to its closest centroid.
//       maximization: Compute the new centroid (mean) of each cluster.
//    until The centroid position do not change.
// end
// Clustering with Constrained Problem for cluster result to have an equal number of member cluster.
// must learn weighted clustering

package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

var d Observations
var vol []int

func main() {
	//setup data
	setupData("Traffic4.csv")
	// Partition the data points into 20 clusters
	// km, _ := kmeans.NewWithOptions(0.01, plotter.SimplePlotter{})
	km, _ := NewWithOptions(0.01, SimplePlotter{})
	clusters, _ := km.Partition(d, 20)

	for _, c := range clusters {
		fmt.Printf("Centered at x: %.2f y: %.2f\n", c.Center[0], c.Center[1])
		fmt.Printf("Matching data points: %+v\n", c.Observations)
		fmt.Printf("total: %d\n\n", len(c.Observations))
	}
	fmt.Println(len(clusters))
	fmt.Println(sum(vol))
}

func setupData(file string) {
	f, err := os.Open(file)
	if err != nil {
		return
	}
	csvReader := csv.NewReader(f)
	csvData, _ := csvReader.ReadAll()

	//read without header
	for i := 1; i < len(csvData); i++ {
		val, _ := strconv.Atoi(csvData[i][3])
		vol = append(vol, val)
		for j := 0; j < val; j++ {
			lat, _ := strconv.ParseFloat(csvData[i][1], 64)
			lng, _ := strconv.ParseFloat(csvData[i][2], 64)
			d = append(d, Coordinates{
				lng,
				lat,
			})
		}

	}
}

func sum(arr []int) int {
	var res int
	res = 0
	for i := 0; i < len(arr); i++ {
		res += arr[i]
	}
	return res
}
