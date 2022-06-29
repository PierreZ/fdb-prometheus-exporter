package models

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	databaseStatus = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "fdb_database_status",
		Help: "state of the dabase",
	}, []string{"state"})

	numConnectedClients = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "fdb_connected_clients",
		Help: "number of connected clients"})

	coordinators = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "fdb_coordinators_reachable",
		Help: "coodrnators info",
	}, []string{"address"})

	movingData = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "fdb_data_move",
		Help: "repartitioning data",
	}, []string{"type"})
)

// ExportDatabaseStatus is exporting the status of the database
func (s FDBStatus) ExportDatabaseStatus() {

	databaseStatus.With(prometheus.Labels{
		"state": "available",
	}).Set(boolToNumber(s.Client.DatabaseStatus.Available))
	databaseStatus.With(prometheus.Labels{
		"state": "healthy",
	}).Set(boolToNumber(s.Client.DatabaseStatus.Healthy))
	databaseStatus.With(prometheus.Labels{
		"state": "quorum_reachable",
	}).Set(boolToNumber(s.Client.Coordinators.QuorumReachable))

	for _, coordinator := range s.Client.Coordinators.Coordinators {
		coordinators.With(prometheus.Labels{
			"address": coordinator.Address,
		}).Set(boolToNumber(coordinator.Reachable))
	}

	numConnectedClients.Set(s.Cluster.Clients.Count)

	movingData.With(prometheus.Labels{
		"type": "in_flight_bytes",
	}).Set(s.Cluster.Data.MovingData.InFlightBytes)

	movingData.With(prometheus.Labels{
		"type": "InQueueBytes",
	}).Set(s.Cluster.Data.MovingData.InQueueBytes)

}

func registerDatabaseStatus(r *prometheus.Registry) {
	r.MustRegister(databaseStatus)
	r.MustRegister(numConnectedClients)
	r.MustRegister(coordinators)
	r.MustRegister(movingData)
}
