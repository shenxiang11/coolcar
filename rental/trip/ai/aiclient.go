package ai

import (
	"context"
	rentalpb "github.com/shenxiang11/coolcar/rental-service/gen/go/proto"
	coolenvpb "github.com/shenxiang11/coolcar/shared/coolenv"
)

type Client struct {
	AIClient coolenvpb.AIServiceClient
}

func (c *Client) DistanceKm(ctx context.Context, from *rentalpb.Location, to *rentalpb.Location) (float64, error) {
	resp, err := c.AIClient.MeasureDistance(ctx, &coolenvpb.MeasureDistanceRequest{
		From: &coolenvpb.Location{
			Latitude:  from.Latitude,
			Longitude: from.Longitude,
		},
		To: &coolenvpb.Location{
			Latitude:  to.Latitude,
			Longitude: to.Longitude,
		},
	})
	if err != nil {
		return 0, err
	}
	return resp.DistanceKm, nil
}
