package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"time"

	"github.com/PierreZ/fdb-prometheus-exporter/models"
	"github.com/apple/foundationdb/bindings/go/src/fdb"
	"github.com/pkg/errors"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	db      fdb.Database
	jsonKey = append([]byte{255, 255}, []byte("/status/json")...)
)

func main() {

	apiVersion, err := strconv.Atoi(getEnv("FDB_API_VERSION", "620"))
	if err != nil {
		log.Fatal("cannot parse FDB_API_VERSION from env")
	}
	// Different API versions may expose different runtime behaviors.
	fdb.MustAPIVersion(apiVersion)

	clusterFile := getEnv("FDB_CLUSTER_FILE", "/var/fdb/data/fdb.cluster")

	if _, exists := os.LookupEnv("FDB_CREATE_CLUSTER_FILE"); exists {
		createClusterFile()
	}

	fmt.Println("opening cluster file at", clusterFile)
	data, err := ioutil.ReadFile(clusterFile)
	if err != nil {
		log.Fatalf("cannot read cluster file: %+v", err)
	}
	fmt.Println(string(data))

	// Open the default database from the system cluster
	db = fdb.MustOpenDatabase(clusterFile)

	exportWorkload, err := strconv.ParseBool(getEnv("FDB_EXPORT_WORKLOAD", "true"))
	if err != nil {
		log.Fatal("cannot parse FDB_EXPORT_WORLOAD from env")
	}
	exportDatabaseStatus, err := strconv.ParseBool(getEnv("FDB_EXPORT_DATABASE_STATUS", "true"))
	if err != nil {
		log.Fatal("cannot parse FDB_EXPORT_DATABASE_STATUS from env")
	}

	exportConfiguration, err := strconv.ParseBool(getEnv("FDB_EXPORT_CONFIGURATION", "true"))
	if err != nil {
		log.Fatal("cannot parse FDB_EXPORT_CONFIGURATION from env")
	}
	exportProcesses, err := strconv.ParseBool(getEnv("FDB_EXPORT_PROCESSES", "true"))
	if err != nil {
		log.Fatal("cannot parse FDB_EXPORT_PROCESSES from env")
	}

	listenTo := getEnv("FDB_METRICS_LISTEN", ":8080")
	refreshEvery, err := strconv.Atoi(getEnv("FDB_METRICS_EVERY", "10"))
	if err != nil {
		log.Fatal("cannot parse FDB_METRICS_EVERY from env")
	}

	ticker := time.NewTicker(time.Duration(refreshEvery) * time.Second)
	go func() {
		for range ticker.C {
			//Call the periodic function here.
			models, err := retrieveMetrics()
			if err != nil {
				fmt.Errorf("cannot retrieve metrics from FDB: (%v)", err)
				continue
			}

			fmt.Println("retrieved data")

			if exportWorkload {
				models.ExportWorkload()
			}

			if exportDatabaseStatus {
				models.ExportDatabaseStatus()
			}

			if exportConfiguration {
				models.ExportConfiguration()
			}

			if exportProcesses {
				models.ExportProcesses()
			}
		}
	}()

	r := prometheus.NewRegistry()
	models.Register(r)

	// [...] update metrics within a goroutine
	handler := promhttp.HandlerFor(r, promhttp.HandlerOpts{})
	http.Handle("/metrics", handler)

	log.Fatal(http.ListenAndServe(listenTo, nil))
}

func retrieveMetrics() (*models.FDBStatus, error) {

	jsonRaw, err := db.ReadTransact(func(tr fdb.ReadTransaction) (interface{}, error) {
		return tr.Get(fdb.Key(jsonKey)).Get()
	})

	if err != nil {
		return nil, errors.Wrap(err, "cannot get status")
	}

	var status models.FDBStatus
	err = json.Unmarshal([]byte(jsonRaw.([]byte)), &status)
	if err != nil {
		return nil, errors.Wrap(err, "cannot decode json")
	}
	return &status, nil
}

// getEnv is wrapping os.getenv with a fallback
func getEnv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}

func createClusterFile() {
	cmd := exec.Command("/fdb.bash")

	fmt.Printf("Running command 'create_cluster_file' and waiting for it to finish...\n")
	stdoutStderr, err := cmd.CombinedOutput()
	fmt.Printf("%s\n", stdoutStderr)
	if err != nil {
		log.Fatal(err)
	}
}
