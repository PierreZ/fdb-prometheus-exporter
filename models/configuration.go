package models

import "github.com/prometheus/client_golang/prometheus"

var (
	configurationStatus = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "fdb_database_configuration",
		Help: "configuration",
	}, []string{"component"})

	generation = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "fdb_generation",
		Help: "generation of the fdb",
	})
)

// ExportConfiguration is exporting the configuration
func (s FDBStatus) ExportConfiguration() {

	configurationStatus.With(prometheus.Labels{
		"component": "coordinators_count",
	}).Set(float64(s.Cluster.Configuration.CoordinatorsCount))
	configurationStatus.With(prometheus.Labels{
		"component": "excluded_servers",
	}).Set(float64(len(s.Cluster.Configuration.ExcludedServers)))

	configurationStatus.With(prometheus.Labels{
		"component": "proxies",
	}).Set(float64(s.Cluster.Configuration.Proxies))

	configurationStatus.With(prometheus.Labels{
		"component": "logs",
	}).Set(float64(s.Cluster.Configuration.Logs))

	configurationStatus.With(prometheus.Labels{
		"component": "log_spill",
	}).Set(float64(s.Cluster.Configuration.LogSpill))

	configurationStatus.With(prometheus.Labels{
		"component": "resolvers",
	}).Set(float64(s.Cluster.Configuration.Resolvers))

	configurationStatus.With(prometheus.Labels{
		"component": "usable_regions",
	}).Set(float64(s.Cluster.Configuration.UsableRegions))

	generation.Set(float64(s.Cluster.Generation))
}

func registerConfiguration(r *prometheus.Registry) {
	r.MustRegister(configurationStatus)
	r.MustRegister(generation)
}
