// begin
//    specify the number k of clustering to assign.
//    randomly initialize k centroids.
//    repeat
//       expectation: Assign each point to its closest centroid.
//       maximization: Compute the new centroid (mean) of each cluster.
//    until The centroid position do not change.
// end
// Clustering with Constrained Problem for cluster result to have an equal number of member cluster.

package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

	"github.com/muesli/clusters"
	"github.com/muesli/kmeans"
	"github.com/muesli/kmeans/plotter"
)

var d clusters.Observations

func main() {
	// set up a random two-dimensional data set (float64 values between 0.0 and 1.0)
	setupData("Average_Daily_Traffic_Counts.csv")
	// Partition the data points into 20 clusters
	km, _ := kmeans.NewWithOptions(0.01, plotter.SimplePlotter{})
	clusters, _ := km.Partition(d, 20)

	for _, c := range clusters {
		fmt.Printf("Centered at x: %.2f y: %.2f\n", c.Center[0], c.Center[1])
		fmt.Printf("Matching data points: %+v\n", c.Observations)
		fmt.Printf("total: %d\n\n", len(c.Observations))
	}
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
		lat, _ := strconv.ParseFloat(csvData[i][6], 64)
		lon, _ := strconv.ParseFloat(csvData[i][7], 64)
		d = append(d, clusters.Coordinates{
			lon,
			lat,
		})

	}
}
