package ranker

import (
	"github.com/Telenav/osrm-backend/integration/oasis/osrmconnector"
	"github.com/Telenav/osrm-backend/integration/service/spatialindexer"
)

const (
	SimpleRanker    = "SimpleRanker"
	OSRMBasedRanker = "OSRMBasedRanker"
)

func CreateRanker(rankerType string, oc *osrmconnector.OSRMConnector) spatialindexer.Ranker {
	switch rankerType {
	case SimpleRanker:
		return newSimpleRanker()
	case OSRMBasedRanker:
		return newOsrmRanker(oc)
	default:
		return newSimpleRanker()
	}
}
