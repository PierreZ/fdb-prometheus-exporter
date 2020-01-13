package main

import (
	"bytes"
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
	db                   fdb.Database
	jsonKey              = append([]byte{255, 255}, []byte("/status/json")...)
	transactionHistogram = prometheus.NewHistogram(
		prometheus.HistogramOpts{
			Name:    "fdb_retrieve_metrics_duration",
			Help:    "retrieve metrics duration distribution",
			Buckets: []float64{1, 2, 5, 10, 20, 60},
		})
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
		}
	}()

	r := prometheus.NewRegistry()
	r.MustRegister(transactionHistogram)
	models.Register(r)

	// [...] update metrics within a goroutine
	handler := promhttp.HandlerFor(r, promhttp.HandlerOpts{})
	http.Handle("/metrics", handler)

	log.Fatal(http.ListenAndServe(listenTo, nil))
}

func retrieveMetrics() (*models.FDBStatus, error) {
	fmt.Println("refreshing metrics")

	start := time.Now()
	jsonRaw, err := db.ReadTransact(func(tr fdb.ReadTransaction) (interface{}, error) {
		return tr.Get(fdb.Key(jsonKey)).Get()
	})
	duration := time.Since(start)

	if err != nil {
		return nil, errors.Wrap(err, "cannot get status")
	}

	transactionHistogram.Observe(duration.Seconds())

	var status models.FDBStatus
	err = json.NewDecoder(bytes.NewReader(jsonRaw.([]byte))).Decode(&status)
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
	cmd := exec.Command("/create_cluster_file.bash")

	fmt.Printf("Running command 'create_cluster_file' and waiting for it to finish...\n")
	stdoutStderr, err := cmd.CombinedOutput()
	fmt.Printf("%s\n", stdoutStderr)
	if err != nil {
		log.Fatal(err)
	}
}
