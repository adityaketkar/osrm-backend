package s2indexer

import (
	"github.com/Telenav/osrm-backend/integration/service/spatialindexer"
	"github.com/Telenav/osrm-backend/integration/service/spatialindexer/poiloader"
	"github.com/golang/geo/s2"
	"github.com/golang/glog"
)

// S2Indexer provide spatial index ability based on google s2
type S2Indexer struct {
	cellID2PointIDs  map[s2.CellID][]spatialindexer.PointID
	pointID2Location map[spatialindexer.PointID]spatialindexer.Location
}

// NewS2Indexer generates spatial indexer based on google s2
func NewS2Indexer() *S2Indexer {
	return &S2Indexer{}
}

// Build constructs S2 indexer
func (indexer *S2Indexer) Build(filePath string) *S2Indexer {
	records, err := poiloader.LoadData(filePath)
	if err != nil || len(records) == 0 {
		return nil
	}

	var pointInfos []spatialindexer.PointInfo

	for _, record := range records {
		pointInfo := spatialindexer.PointInfo{
			ID: elementID2PointID(record.ID),
			Location: spatialindexer.Location{
				Lat: record.Lat,
				Lon: record.Lon,
			},
		}
		pointInfos = append(pointInfos, pointInfo)

		indexer.pointID2Location[elementID2PointID(record.ID)] = spatialindexer.Location{
			Lat: record.Lat,
			Lon: record.Lon,
		}
	}

	indexer.cellID2PointIDs = build(pointInfos, minS2Level, maxS2Level)
	return indexer
}

// Load S2Indexer's data from contents recorded in folder
func (indexer *S2Indexer) Load(folderPath string) *S2Indexer {
	if err := deSerializeS2Indexer(indexer, folderPath); err != nil {
		glog.Errorf("Load S2Indexer's data from folder %s failed, err=%v\n", folderPath, err)
		return nil
	}
	return indexer
}

// Dump S2Indexer's content into folderPath
func (indexer *S2Indexer) Dump(folderPath string) {
	if err := serializeS2Indexer(indexer, folderPath); err != nil {
		glog.Errorf("Dump S2Indexer's data to folder %s failed, err=%v\n", folderPath, err)
	}
}

// IteratePoints returns PointInfo in channel
// It implements interface of PointsIterator
func (indexer *S2Indexer) IteratePoints() <-chan spatialindexer.PointInfo {
	pointsC := make(chan spatialindexer.PointInfo, len(indexer.pointID2Location))
	go func() {
		for pointID, location := range indexer.pointID2Location {
			pointsC <- spatialindexer.PointInfo{
				ID:       pointID,
				Location: location,
			}
		}
		close(pointsC)
	}()

	return pointsC
}

// FindNearByPointIDs returns nearby points for given center and conditions
func (indexer *S2Indexer) FindNearByPointIDs(center spatialindexer.Location, radius float64, limitCount int) []*spatialindexer.PointInfo {
	if !indexer.isInitialized() {
		glog.Warning("S2Indexer is empty, try to Build() with correct input file first.\n")
		return nil
	}

	results := queryNearByPoints(indexer, center, radius)
	if limitCount != spatialindexer.UnlimitedCount && len(results) > limitCount {
		results = results[:limitCount]
	}

	return results
}

func (indexer S2Indexer) getPointLocationByPointID(id spatialindexer.PointID) (spatialindexer.Location, bool) {
	location, ok := indexer.pointID2Location[id]
	return location, ok
}

func (indexer S2Indexer) getPointIDsByS2CellID(cellid s2.CellID) ([]spatialindexer.PointID, bool) {
	pointIDs, ok := indexer.cellID2PointIDs[cellid]
	return pointIDs, ok
}

func (indexer S2Indexer) isInitialized() bool {
	return len(indexer.cellID2PointIDs) != 0 && len(indexer.pointID2Location) != 0
}

func elementID2PointID(id int64) spatialindexer.PointID {
	return (spatialindexer.PointID)(id)
}
