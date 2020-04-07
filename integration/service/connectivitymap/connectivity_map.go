package connectivitymap

import (
	"github.com/Telenav/osrm-backend/integration/service/spatialindexer"
	"github.com/golang/glog"
)

// IDAndDistance wraps ID and distance information
type IDAndDistance struct {
	ID       spatialindexer.PointID
	Distance float64
}

// ID2NearByIDsMap is a mapping between ID and its nearby IDs
type ID2NearByIDsMap map[spatialindexer.PointID][]IDAndDistance

// Connectivity Map used to query connectivity for given placeID
type ConnectivityMap struct {
	id2nearByIDs ID2NearByIDsMap
	maxRange     float64
	statistic    *statistic
}

// NewConnectivityMap creates ConnectivityMap
func NewConnectivityMap(maxRange float64) *ConnectivityMap {
	return &ConnectivityMap{
		maxRange:  maxRange,
		statistic: newStatistic(),
	}
}

// Build creates ConnectivityMap
func (cm *ConnectivityMap) Build() {
	glog.Info("Successfully finished GenerateConnectivityMap\n")
}

// Dump dump ConnectivityMap's content to given folderPath
func (cm *ConnectivityMap) Dump(folderPath string) {
}

// Load rebuild ConnectivityMap from dumpped data in given folderPath
func (cm *ConnectivityMap) Load(folderPath string) {
}

// QueryConnectivity answers connectivity query for given placeInfo
func (cm *ConnectivityMap) QueryConnectivity(placeInfo spatialindexer.PointInfo, limitDistance float64) {
	// for each everything recorded in data, apply limit option on that
}

// MaxRange tells the value used to pre-process place data.
// MaxRange means the maximum distance in meters could be reached from current location.
func (cm *ConnectivityMap) MaxRange() float64 {
	return cm.maxRange
}
