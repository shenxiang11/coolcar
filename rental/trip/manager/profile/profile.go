package profile

import (
	"context"
	"github.com/shenxiang11/coolcar/shared/id"
)

type Manager struct {
}

func (m *Manager) Verify(ctx context.Context, aID id.AccountID) (id.IdentityID, error) {
	return id.IdentityID("identity1"), nil
}
