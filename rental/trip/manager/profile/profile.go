package profile

import (
	"context"
	"encoding/base64"
	"fmt"
	rentalpb "github.com/shenxiang11/coolcar/rental-service/gen/go/proto"
	"github.com/shenxiang11/coolcar/shared/id"
	"google.golang.org/protobuf/proto"
)

type Manager struct {
	Fetcher Fetcher
}

type Fetcher interface {
	GetProfile(c context.Context, req *rentalpb.GetProfileRequest) (*rentalpb.Profile, error)
}

func (m *Manager) Verify(ctx context.Context, aID id.AccountID) (id.IdentityID, error) {
	nilID := id.IdentityID("")
	p, err := m.Fetcher.GetProfile(ctx, &rentalpb.GetProfileRequest{})
	if err != nil {
		return nilID, fmt.Errorf("cannot get profile: %v", err)
	}

	if p.IdentityStatus != rentalpb.IdentityStatus_VERIFIED {
		return nilID, fmt.Errorf("invalid identity status")
	}

	b, err := proto.Marshal(p.Identity)
	if err != nil {
		return nilID, fmt.Errorf("cannnot marshal identity: %v", err)
	}

	return id.IdentityID(base64.StdEncoding.EncodeToString(b)), nil
}
