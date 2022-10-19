package osrmtype

import (
	"reflect"
	"testing"
)

func TestEdgeBasedEdgesWrite(t *testing.T) {

	cases := []struct {
		p []byte
		EdgeBasedEdges
	}{
		{nil, nil},
		{[]byte{}, nil},
		{
			[]byte{
				0xFB, 0x1F, 0x04, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x2F, 0x00, 0x00, 0x00, 0xC3, 0xA2, 0x8D, 0x42, 0x2F, 0x00, 0x00, // insufficient data
			},
			nil,
		},
		{
			[]byte{
				0xFB, 0x1F, 0x04, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x2F, 0x00, 0x00, 0x00, 0xC3, 0xA2, 0x8D, 0x42, 0x2F, 0x00, 0x00, 0x40,
			},
			EdgeBasedEdges{
				EdgeBasedEdge{
					Source: 270331,
					Target: 0,
					EdgeData: EdgeData{
						TurnID:   0,
						Weight:   47,
						Distance: 70.817894,
						Duration: 47,
						Forward:  true,
						Backward: false,
					},
				},
			},
		},
		{
			[]byte{
				0xFB, 0x1F, 0x04, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x2F, 0x00, 0x00, 0x00, 0xC3, 0xA2, 0x8D, 0x42, 0x2F, 0x00, 0x00, 0x40,
				0xFB, 0x1F, 0x04, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x2F, 0x00, 0x00, 0x00, 0xC3, 0xA2, 0x8D, 0x42, 0x2F, 0x00, 0x00, // insufficient data
			},
			EdgeBasedEdges{
				EdgeBasedEdge{
					Source: 270331,
					Target: 0,
					EdgeData: EdgeData{
						TurnID:   0,
						Weight:   47,
						Distance: 70.817894,
						Duration: 47,
						Forward:  true,
						Backward: false,
					},
				},
			},
		},
		{
			[]byte{
				0xFB, 0x1F, 0x04, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x2F, 0x00, 0x00, 0x00, 0xC3, 0xA2, 0x8D, 0x42, 0x2F, 0x00, 0x00, 0x40,
				0xFB, // insufficient data
			},
			EdgeBasedEdges{
				EdgeBasedEdge{
					Source: 270331,
					Target: 0,
					EdgeData: EdgeData{
						TurnID:   0,
						Weight:   47,
						Distance: 70.817894,
						Duration: 47,
						Forward:  true,
						Backward: false,
					},
				},
			},
		},
	}

	for _, c := range cases {
		var edges EdgeBasedEdges
		writeLen, err := edges.Write(c.p)
		if err != nil {
			t.Error(err)
		}
		if len(c.p)-writeLen != len(c.p)%edgeBasedEdgeBytes {
			t.Errorf("len(p) %d but write len %d", len(c.p), writeLen)
		}
		if !reflect.DeepEqual(edges, c.EdgeBasedEdges) {
			t.Errorf("parse %v expect %v but got %v", c.p, c.EdgeBasedEdges, edges)
		}
	}
}