module githubntf/common

go 1.12

require (
	github.com/GoogleCloudPlatform/berglas v0.2.0
	github.com/kelseyhightower/envconfig v1.4.0
	githubntf/repository-service-proto v0.0.0-00010101000000-000000000000
	gopkg.in/yaml.v2 v2.2.2
)

replace githubntf/repository-service-proto => ../repository-service-proto
