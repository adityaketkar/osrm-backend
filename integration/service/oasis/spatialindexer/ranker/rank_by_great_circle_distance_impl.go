package ranker

import (
	"github.com/Telenav/osrm-backend/integration/service/oasis/spatialindexer"
	"github.com/blevesearch/bleve/geo"
	"github.com/golang/glog"
)

// defaultSpeed is 22.2 meter/second
// 22.2 m/s = 50 miles/hour
const defaultSpeed = 22.2

func rankPointsByGreatCircleDistanceToCenter(center spatialindexer.Location, targets []*spatialindexer.PointInfo) []*spatialindexer.RankedPointInfo {
	if len(targets) == 0 {
		glog.Warningf("When try to rankPointsByGreatCircleDistanceToCenter, input array is empty, center = %+v\n", center)
		return nil
	}

	pointWithDistanceC := make(chan *spatialindexer.RankedPointInfo, len(targets))
	go func() {
		defer close(pointWithDistanceC)

		for _, p := range targets {
			// geo.Haversin's unit is kilometer, convert to meter
			length := geo.Haversin(center.Lon, center.Lat, p.Location.Lon, p.Location.Lat) * 1000

			pointWithDistanceC <- &spatialindexer.RankedPointInfo{
				PointInfo: spatialindexer.PointInfo{
					ID:       p.ID,
					Location: p.Location,
				},
				Distance: length,
				Duration: length / defaultSpeed,
			}
		}
	}()

	rankAgent := newRankAgent(len(targets))
	return rankAgent.RankByDistance(pointWithDistanceC)
}
