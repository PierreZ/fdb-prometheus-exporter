package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/apple/foundationdb/bindings/go/src/fdb"
	"github.com/pkg/errors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var db fdb.Database

func main() {

	apiVersion, err := strconv.Atoi(getEnv("FDB_API_VERSION", "620"))
	if err != nil {
		log.Fatal("cannot parse FDB_API_VERSION from env")
	}
	// Different API versions may expose different runtime behaviors.
	fdb.MustAPIVersion(apiVersion)

	// Open the default database from the system cluster
	db = fdb.MustOpenDatabase(getEnv("FDB_CLUSTER_FILE", "/var/fdb/data/fdb.cluster"))

	exportWorkload, err := strconv.ParseBool(getEnv("FDB_EXPORT_WORKLOAD", "true"))
	if err != nil {
		log.Fatal("cannot parse FDB_EXPORT_WORLOAD from env")
	}

	ticker := time.NewTicker(2 * time.Second)
	go func() {
		for t := range ticker.C {
			//Call the periodic function here.
			status, err := retrieveMetrics(t)
			if err != nil {
				fmt.Errorf("cannot retrieve metrics from FDB: (%v)", err)
			}

			if exportWorkload {
				status.exportWorkload()
			}
		}
	}()

	// Expose the registered metrics via HTTP.
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func retrieveMetrics(tick time.Time) (*FDBStatus, error) {
	fmt.Println("refreshing metrics", tick)
	status, err := db.Transact(func(tr fdb.Transaction) (interface{}, error) {
		keyCode := []byte("\xff\xff/status/json")
		var k fdb.Key
		k = keyCode
		resp, err := tr.Get(k).Get()
		if err != nil {
			return nil, errors.Wrap(err, "cannot retrieve key")
		}
		if len(resp) == 0 {
			return nil, errors.Wrap(err, "no key for status")
		}

		status := FDBStatus{}
		err = json.Unmarshal(resp, &status)
		if err != nil {
			return nil, errors.Wrap(err, "cannot unmarshal key")
		}

		return &status, nil
	})

	if err != nil {
		return nil, err
	}

	fdbStatus := status.(FDBStatus)
	return &fdbStatus, nil
}

// getEnv is wrapping os.getenv with a fallback
func getEnv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}
