package main

import (
	"github.com/influxdata/influxdb/client/v2"
	"github.com/jinzhu/gorm"
	"github.com/jonlil/nibe-go"
	"github.com/jonlil/nibestats/database"
	"github.com/jonlil/nibestats/models"
	"github.com/jonlil/nibestats/utils"
	"log"
	"strconv"
	"time"
)

var (
	db *gorm.DB
)

func fetchUserParameters(api *nibe.API, system *nibe.System) []nibe.Parameter {
	parameters, _ := api.GetParameters(system, []string{
		"40004",
		"40033",
	})

	return parameters
}

func getSystems(api *nibe.API) []nibe.System {
	systems, err := api.GetSystems()
	if err != nil {
		return nil
	}
	return *systems
}

func getUsers() []models.AccessToken {
	tokens := []models.AccessToken{}
	db.Find(&tokens)
	return tokens
}

func run(c client.Client, bp client.BatchPoints) error {
	for _, at := range getUsers() {
		api := nibe.NewAPI(at.Token)
		for _, system := range getSystems(api) {
			log.Println(at.UserID)
			parameters := fetchUserParameters(api, &system)
			tags := map[string]string{
				"user": strconv.FormatInt(at.UserID, 10),
			}

			fields := map[string]interface{}{
				"indoor":  parameters[1].RawValue,
				"outdoor": parameters[0].RawValue,
			}

			pt, err := client.NewPoint("temperature", tags, fields, time.Now())
			if err != nil {
				log.Fatal(err)
			}
			bp.AddPoint(pt)
		}
	}
	// Write the batch
	if err := c.Write(bp); err != nil {
		log.Fatal(err)
	}

	return nil
}

func main() {
	db = database.Open()

	c, _ := client.NewHTTPClient(client.HTTPConfig{
		Addr: utils.GetEnv("INFLUX_DB_HOST", "http://localhost:8086"),
	})

	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  "nibestats",
		Precision: "m",
	})

	if err != nil {
		log.Fatal(err)
	}

	for {
		time.Sleep(time.Duration(1) * time.Minute)
		run(c, bp)
	}

	// Close client resources
	if err := c.Close(); err != nil {
		log.Fatal(err)
	}
}
