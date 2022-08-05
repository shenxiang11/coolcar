package ai

import (
	"context"
	"fmt"
	rentalpb "github.com/shenxiang11/coolcar/rental-service/gen/go/proto"
	coolenvpb "github.com/shenxiang11/coolcar/shared/coolenv"
)

type Client struct {
	AIClient  coolenvpb.AIServiceClient
	UseRealAI bool
}

func (c *Client) Resolve(ctx context.Context, photo []byte) (*rentalpb.Identity, error) {
	i, err := c.AIClient.LicIdentity(ctx, &coolenvpb.IdentityRequest{
		Photo:  photo,
		RealAi: c.UseRealAI,
	})
	if err != nil {
		return nil, fmt.Errorf("cannnot resolve identity: %v", err)
	}
	return &rentalpb.Identity{
		LicNumber:       i.LicNumber,
		Name:            i.Name,
		Gender:          rentalpb.Gender(i.Gender),
		BirthDateMillis: i.BirthDateMillis,
	}, nil
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
