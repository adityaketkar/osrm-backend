package stationgraph

import "github.com/Telenav/osrm-backend/integration/service/oasis/chargingstrategy"

// IGraph defines interface used for Graph
type IGraph interface {
	// Node returns node object by its nodeID
	Node(id nodeID) *node

	// AdjacentNodes returns a group of node ids which connect with given node id
	// The connectivity between nodes is build during running time.
	AdjacentNodes(id nodeID) []nodeID

	// Edge returns edge information between given two nodes
	Edge(from, to nodeID) *edge

	// StartNodeID returns start node's ID for given graph
	StartNodeID() nodeID

	// EndNodeID returns end node's ID for given graph
	EndNodeID() nodeID

	// ChargeStrategy returns charge strategy used for graph construction
	ChargeStrategy() chargingstrategy.Strategy
}

// IStationInfo defines station related information
type IStationInfo interface {
	StationID(id nodeID) string

	StationLocation(id nodeID) *locationInfo
}