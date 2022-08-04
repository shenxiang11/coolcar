package car

import (
	"context"
	rentalpb "github.com/shenxiang11/coolcar/rental-service/gen/go/proto"
	"github.com/shenxiang11/coolcar/shared/id"
)

type Manager struct {
}

func (m *Manager) Verify(c context.Context, cid id.CarID, loc *rentalpb.Location) error {
	return nil
}

func (m *Manager) Unlock(c context.Context, cid id.CarID, aid id.AccountID, tid id.TripID, avatarURL string) error {
	return nil
}

func (m *Manager) Lock(c context.Context, cid id.CarID) error {
	return nil
}
