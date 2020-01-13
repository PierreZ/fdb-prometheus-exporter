package models

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	databaseStatus = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "fdb_database_status",
		Help: "state of the dabase",
	}, []string{"state"})
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
}

func registerDatabaseStatus(r *prometheus.Registry) {
	r.MustRegister(databaseStatus)
}
