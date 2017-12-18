package main

import (
	"log"
	"time"
	"math/rand"
	"github.com/influxdata/influxdb/client/v2"
)


const (
	MyDB = "sdp_test"
	username = "user"
	password = "pass"
)

func main() {

	// Create a new HTTPClient
	influxCnct, err := client.NewHTTPClient(client.HTTPConfig{
		Addr:     "host",
		Username: username,
		Password: password,
	})
	if err != nil {
		log.Fatal(err)
	}

	// Create a new point batch
	influxPoint, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  MyDB,
		Precision: "s",
	})
	if err != nil {
		log.Fatal(err)
	}

	for  {

		sampleSize := 1000

		for i := 0; i < sampleSize; i++ {

				regions := []string{"us-west1", "us-west2", "us-west3", "us-east1"}
				host := []string{"us-west1-host1", "us-west2-host2", "us-west3-hostc", "us-east1-host"}
				cluster := []string{"cluster1", "cluster2", "clusterc", "cluster3"}

				// Create a point and add to batch
				tagsTask := map[string]string{
					"cluster": cluster[rand.Intn(len(cluster))],
					"host": host[rand.Intn(len(host))],
					"region": regions[rand.Intn(len(regions))],
				}
				fieldsTask := map[string]interface{}{
					"trace":   "field_test",
					"metric": rand.Intn(100),
				}

				ptTask, err := client.NewPoint("sdp_test", tagsTask, fieldsTask, time.Now())
				if err != nil {
					log.Fatal(err)
				}
				influxPoint.AddPoint(ptTask)


		}

		// Write the batch
		if err := influxCnct.Write(influxPoint); err != nil {
			log.Fatal(err)
		}

	}


}
