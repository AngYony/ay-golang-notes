package trip

import (
	"context"
	trippb "coolcar/proto/gen/go"
)

type Service struct {
}

func (*Service) GetTrip(c context.Context,
	req *trippb.GetTripRequest) (*trippb.GetTripResponse, error) {
	return &trippb.GetTripResponse{
		Id: req.Id,
		Trip: &trippb.Trip{
			Start:       "abc",
			End:         "def",
			DurationSec: 3600,
			FeeCent:     10000,
			StartPos: &trippb.Location{
				Latitude:  35,
				Longitude: 100,
			},
			EndPos: &trippb.Location{
				Latitude:  40,
				Longitude: 130,
			},
			PathLocations: []*trippb.Location{
				{
					Latitude:  50,
					Longitude: 100,
				},
				{
					Latitude:  66,
					Longitude: 77,
				},
			},

			Status: trippb.TripStatus_IN_PROGRESS,
		},
	}, nil
}
