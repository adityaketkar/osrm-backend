package entrypoint

import (
	"net/http"

	"github.com/Telenav/osrm-backend/integration/api/oasis"
	"github.com/Telenav/osrm-backend/integration/service/oasis/internal/resourcemanager"
	"github.com/Telenav/osrm-backend/integration/service/oasis/solution"
	"github.com/golang/glog"
)

// httpHandler handles oasis request and provide response
type httpHandler struct {
	resourceMgr *resourcemanager.ResourceMgr
}

// NewHttpHandler creates new Handler object
//. Oasis is only a service to determine charging stations. All other data like path costs and charging
//. station information are provided by other services
func NewHttpHandler(osrmBackend, finderType, searchEndpoint, apiKey, apiSignature, dataFolderPath string) (*httpHandler, error) {
	resourceMgr, err := resourcemanager.NewResourceMgr(osrmBackend, finderType, searchEndpoint, apiKey, apiSignature, dataFolderPath)
	if err != nil {
		glog.Errorf("Failed to create Handler due to error %+v.\n", err)
		return nil, err
	}

	return &httpHandler{
		resourceMgr: resourceMgr,
	}, nil
}

func (h *httpHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	glog.Infof("Handle incoming request %s from remote addr %s", req.RequestURI, req.RemoteAddr)

	// parse oasis request
	oasisReq, err := oasis.ParseRequestURL(req.URL)
	if err != nil || len(oasisReq.Coordinates) != 2 {
		generateResponse4BadRequest(w, err)
		return
	}

	// Calculate optimal charge solution
	//. layer 1 entrypoint passes the data to layer 2 solution
	statusCode, solutions, err := solution.NewGeneratorImpl(h.resourceMgr).Generate(oasisReq)
	if err != nil {
		generateResponseWhenMetError(w, statusCode, err)
	} else if statusCode == solution.StatusOrigAndDestIsNotReachable {
		generateResponse4UnableReachDestination(w, statusCode)
	} else {
		generateResponse(w, statusCode, solutions)
	}

}
