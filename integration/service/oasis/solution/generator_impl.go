package solution

import (
	"fmt"

	"github.com/Telenav/osrm-backend/integration/api/oasis"
	"github.com/Telenav/osrm-backend/integration/service/oasis/internal/osrmhelper"
	"github.com/Telenav/osrm-backend/integration/service/oasis/internal/resourcemanager"
	"github.com/Telenav/osrm-backend/integration/service/oasis/solution/selectionstrategy"
)

type generatorImpl struct {
	resourceMgr *resourcemanager.ResourceMgr
}

// NewGeneratorImpl creates implementation of interface Generator
func NewGeneratorImpl(resourceMgr *resourcemanager.ResourceMgr) *generatorImpl {
	return &generatorImpl{
		resourceMgr: resourceMgr,
	}
}

// Generate calculates oasis solutions based on oasis request
func (g *generatorImpl) Generate(oasisReq *oasis.Request) (int, []*oasis.Solution, error) {
	// generate route response based on given oasis's orig/destination
	//. get the route from origin to destination as first step
	//. No fancy EV stuff till here
	routeResp, err := osrmhelper.RequestRoute4InputOrigDest(oasisReq, g.resourceMgr.OSRMConnector())
	if err != nil {
		return StatusFailedToCalculateRoute, nil, err
	}

	//. check whether orig and dest is reachable at all
	if len(routeResp.Routes) == 0 {
		return StatusOrigAndDestIsNotReachable, nil, nil
	}

	//. check whether vehicle has enough energy to complete the trip without charge
	b, energyLeft, err := selectionstrategy.HasEnoughEnergy(oasisReq.CurrRange, oasisReq.SafeLevel, routeResp)
	if err != nil {
		return StatusFailedToGenerateChargeResult, nil, err
	}
	//. if yes, return the route directly by generating a solution object
	if b {
		solutions, err := selectionstrategy.GenerateSolution4NoChargeNeeded(routeResp, energyLeft)
		return StatusNoNeedCharge, solutions, err
	}

	// check whether orig and dest reachable range covers shared stations
	//. ie. check if possible in single charge
	//! notice we only pass the first route's distance to HasEnoughEnergy (check implementation of GetOverlapChargeStations4OrigDest to know why)
	//. for now, we only assume best case scenario for above
	//. also observation: can't see any example where two routes are returned
	overlap := selectionstrategy.GetOverlapChargeStations4OrigDest(oasisReq, routeResp.Routes[0].Distance, g.resourceMgr)
	if len(overlap) > 0 {
		solutions, err := selectionstrategy.GenerateResponse4SingleChargeStation(nil, oasisReq, overlap, g.resourceMgr)
		return StatusChargeForSingleTime, solutions, err
	}

	//. need to charge multiple times
	// generate result for multiple charge
	solutions := selectionstrategy.GenerateSolutions4SearchAlongRoute(oasisReq, routeResp, g.resourceMgr.OSRMConnector(), g.resourceMgr.IteratorGenerator())
	if len(solutions) == 0 && err != nil {
		err = fmt.Errorf("Failed to generate charge solution, internal error.\n")
	}
	return StatusChargeForMultipleTime, solutions, err
}
