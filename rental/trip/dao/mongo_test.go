package dao

import (
	"context"
	rentalpb "github.com/shenxiang11/coolcar/rental/gen/go/proto"
	mongotesting "github.com/shenxiang11/coolcar/shared/testing"
	"os"
	"testing"
)

func TestCreateTrip(t *testing.T) {
	c := context.Background()
	mc, err := mongotesting.NewClient(c)
	if err != nil {
		t.Fatalf("cannot connect mongodb: %v", err)
	}

	m := NewMongo(mc.Database("coolcar"))

	tr, err := m.CreateTrip(c, &rentalpb.Trip{
		AccountId: "account1",
		CarId:     "car1",
		Start: &rentalpb.LocationStatus{
			Location: &rentalpb.Location{
				Latitude:  30,
				Longitude: 120,
			},
			FeeCent:      0,
			KmDriven:     0,
			PoiName:      "startpoint",
			TimestampSec: 0,
		},
		Current:    nil,
		End:        nil,
		Status:     rentalpb.TripStatus_FINISHED,
		IdentityId: "",
	})

	if err != nil {
		t.Errorf("cannot create trip: %v", err)
	}

	got, err := m.GetTrip(c, tr.ID.Hex(), "account1")
	if err != nil {
		t.Errorf("cannot get trip: %v", err)
	}

	t.Errorf("got trip: %+v", got)
}

func TestMain(m *testing.M) {
	os.Exit(mongotesting.RunWithMongoInDocker(m))
}
