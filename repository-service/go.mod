module githubntf

go 1.12

require (
	github.com/99designs/gqlgen v0.9.1
	github.com/GoogleCloudPlatform/berglas v0.2.0
	github.com/go-stack/stack v1.8.0 // indirect
	github.com/golang/protobuf v1.3.2
	github.com/golang/snappy v0.0.1 // indirect
	github.com/google/go-github/v27 v27.0.4
	github.com/google/wire v0.3.0
	github.com/kelseyhightower/envconfig v1.4.0
	github.com/tidwall/pretty v1.0.0 // indirect
	github.com/twitchtv/twirp v5.8.0+incompatible
	github.com/vektah/gqlparser v1.1.2
	github.com/xdg/scram v0.0.0-20180814205039-7eeb5667e42c // indirect
	github.com/xdg/stringprep v1.0.0 // indirect
	githubntf/common v0.0.0
	go.mongodb.org/mongo-driver v1.0.4
	gopkg.in/yaml.v2 v2.2.2
)

replace githubntf/common => ../common
