package stationfinder

import (
	"sync"

	"github.com/Telenav/osrm-backend/integration/pkg/api/search/nearbychargestation"
)

type basicFinder struct {
}

func (bf *basicFinder) iterateNearbyStations(stations []*nearbychargestation.Result, respLock *sync.RWMutex) <-chan ChargeStationInfo {
	if len(stations) == 0 {
		c := make(chan ChargeStationInfo)
		go func() {
			defer close(c)
		}()
		return c
	}

	if respLock != nil {
		respLock.RLock()
	}
	size := len(stations)
	results := make([]*nearbychargestation.Result, size)
	copy(results, stations)
	if respLock != nil {
		respLock.RUnlock()
	}

	c := make(chan ChargeStationInfo, size)
	go func() {
		defer close(c)
		for _, result := range results {
			if len(result.Place.Address) == 0 {
				continue
			}
			station := ChargeStationInfo{
				ID: result.ID,
				Location: StationCoordinate{
					Lat: result.Place.Address[0].GeoCoordinate.Latitude,
					Lon: result.Place.Address[0].GeoCoordinate.Longitude},
			}
			c <- station
		}
	}()

	return c
}