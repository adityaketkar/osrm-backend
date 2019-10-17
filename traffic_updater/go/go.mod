module github.com/Telenav/osrm-backend/traffic_updater/go

go 1.12

require (
	github.com/golang/protobuf v1.3.2
	github.com/golang/snappy v0.0.1
	github.com/qedus/osmpbf v1.1.0
	google.golang.org/grpc v1.22.0
)

replace github.com/Telenav/osrm-backend/traffic_updater/go/grpc/proxy => ./grpc/proxy
