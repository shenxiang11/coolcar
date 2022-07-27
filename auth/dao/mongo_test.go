package dao

import (
	"context"
	"github.com/shenxiang11/coolcar/shared/id"
	mgo "github.com/shenxiang11/coolcar/shared/mongo"
	"github.com/shenxiang11/coolcar/shared/mongo/objid"
	mongotesting "github.com/shenxiang11/coolcar/shared/testing"
	"go.mongodb.org/mongo-driver/bson"
	"os"
	"testing"
)

func TestResoleAccountID(t *testing.T) {
	c := context.Background()
	mc, err := mongotesting.NewClient(c)
	if err != nil {
		t.Fatalf("cannot connect mongodb: %v", err)
	}

	m := NewMongo(mc.Database("coolcar"))
	m.genID = mgo.MockNewObjIdWithValue(id.AccountID("62e0b23a691951fd24073fa8"))
	_, err = m.col.InsertMany(c, []any{
		bson.M{
			mgo.IDFieldName: objid.MustFromID(id.AccountID("5f7c245ab0361e00ffb9fd6f")),
			openIDField:     "openid_1",
		},
		bson.M{
			mgo.IDFieldName: objid.MustFromID(id.AccountID("5f7c245ab0361e00ffb9fd70")),
			openIDField:     "openid_2",
		},
	})
	if err != nil {
		t.Fatalf("cannot insert initial values: %v", err)
	}

	cases := []struct {
		name   string
		openID string
		want   string
	}{
		{
			"existing_user",
			"openid_1",
			"5f7c245ab0361e00ffb9fd6f",
		},
		{
			"another_existing_user",
			"openid_2",
			"5f7c245ab0361e00ffb9fd70",
		},
		{
			"new_user",
			"openid_3",
			"62e0b23a691951fd24073fa8",
		},
	}

	for _, cc := range cases {
		t.Run(cc.name, func(t *testing.T) {
			id, err := m.ResolveAccountID(context.Background(), cc.openID)
			if err != nil {
				t.Errorf("failed resolve account id for %q: %v", cc.openID, err)
			}
			if id != cc.want {
				t.Errorf("resolve acount id, want %q, got %q", cc.want, id)
			}
		})
	}
}

func TestMain(m *testing.M) {
	os.Exit(mongotesting.RunWithMongoInDocker(m))
}
