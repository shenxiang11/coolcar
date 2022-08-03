package dao

import (
	"context"
	"github.com/google/go-cmp/cmp"
	rentalpb "github.com/shenxiang11/coolcar/rental-service/gen/go/proto"
	"github.com/shenxiang11/coolcar/shared/id"
	"github.com/shenxiang11/coolcar/shared/mongo/objid"
	mongotesting "github.com/shenxiang11/coolcar/shared/testing"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/protobuf/testing/protocmp"
	"os"
	"testing"
)

func TestCreateTrip(t *testing.T) {
	c := context.Background()
	mc, err := mongotesting.NewClient(c)
	if err != nil {
		t.Fatalf("cannot connect mongodb: %v", err)
	}

	db := mc.Database("coolcar")
	err = mongotesting.SetupIndexes(c, db)
	if err != nil {
		t.Fatalf("cannot setup indexes: %v", err)
	}

	m := NewMongo(db)

	cases := []struct {
		name       string
		tripID     id.TripID
		accountID  string
		tripStatus rentalpb.TripStatus
		wantErr    bool
	}{
		{
			"finished",
			"5f8132eb00714bf62948905c",
			"account1",
			rentalpb.TripStatus_FINISHED,
			false,
		},
		{
			"another_finished",
			"5f8132eb00714bf62948905d",
			"account1",
			rentalpb.TripStatus_FINISHED,
			false,
		},
		{
			"in_progress",
			"5f8132eb00714bf62948905e",
			"account1",
			rentalpb.TripStatus_IN_PROGRESS,
			false,
		},
		{
			"another_in_progress",
			"5f8132eb00714bf62948905f",
			"account1",
			rentalpb.TripStatus_IN_PROGRESS,
			true,
		},
		{
			"in_progress_by_another_account",
			"5f8132eb00714bf629489060",
			"account2",
			rentalpb.TripStatus_IN_PROGRESS,
			false,
		},
	}

	for _, cc := range cases {
		m.GenIDFunc = func() primitive.ObjectID {
			return objid.MustFromID(cc.tripID)
		}
		tr, err := m.CreateTrip(c, &rentalpb.Trip{
			AccountId: cc.accountID,
			Status:    cc.tripStatus,
		})
		if cc.wantErr {
			if err == nil {
				t.Errorf("%s: error expected; got none", cc.name)
			}
			continue
		}
		if err != nil {
			t.Errorf("%s: error creating trip: %v", cc.name, err)
			continue
		}
		if tr.ID.Hex() != cc.tripID.String() {
			t.Errorf("%s: incorrect trip id: %q; got: %q", cc.name, cc.tripID, tr.ID.Hex())
		}
	}
}

func TestGetTrip(t *testing.T) {
	c := context.Background()
	mc, err := mongotesting.NewClient(c)
	if err != nil {
		t.Fatalf("cannot connect mongodb: %v", err)
	}

	m := NewMongo(mc.Database("coolcar"))
	acct := id.AccountID("account1")
	m.GenIDFunc = primitive.NewObjectID
	tr, err := m.CreateTrip(c, &rentalpb.Trip{
		AccountId: acct.String(),
		CarId:     "car1",
		Start: &rentalpb.LocationStatus{
			Location: &rentalpb.Location{
				Latitude:  30,
				Longitude: 120,
			},
			PoiName: "startpoint",
		},
		End: &rentalpb.LocationStatus{
			Location: &rentalpb.Location{
				Latitude:  35,
				Longitude: 120,
			},
			PoiName: "endpoint",
		},
		Status: rentalpb.TripStatus_FINISHED,
	})
	if err != nil {
		t.Fatalf("cannot get trip: %v", err)
	}

	got, err := m.GetTrip(c, objid.ToTripID(tr.ID).String(), acct)
	if err != nil {
		t.Errorf("cannot get trip: %v", err)
	}
	if diff := cmp.Diff(tr, got, protocmp.Transform()); diff != "" {
		t.Errorf("result differs; -want +got: %s", diff)
	}
}

func TestGetTrips(t *testing.T) {
	rows := []struct {
		id        id.TripID
		accountId id.AccountID
		status    rentalpb.TripStatus
	}{
		{
			"5f8132eb10714bf629489051",
			"account_id_for_get_trips",
			rentalpb.TripStatus_FINISHED,
		},
		{
			"5f8132eb10714bf629489052",
			"account_id_for_get_trips",
			rentalpb.TripStatus_FINISHED,
		},
		{
			"5f8132eb10714bf629489053",
			"account_id_for_get_trips",
			rentalpb.TripStatus_FINISHED,
		},
		{
			"5f8132eb10714bf629489054",
			"account_id_for_get_trips",
			rentalpb.TripStatus_IN_PROGRESS,
		},
		{
			"5f8132eb10714bf629489055",
			"account_id_for_get_trips_1",
			rentalpb.TripStatus_IN_PROGRESS,
		},
	}

	c := context.Background()
	mc, err := mongotesting.NewClient(c)
	if err != nil {
		t.Fatalf("cannot connect mongodb: %v", err)
	}

	m := NewMongo(mc.Database("coolcar"))

	for _, r := range rows {
		m.GenIDFunc = func() primitive.ObjectID {
			return objid.MustFromID(r.id)
		}
		_, err := m.CreateTrip(c, &rentalpb.Trip{AccountId: r.accountId.String(), Status: r.status})
		if err != nil {
			t.Fatalf("cannot create rows: %v", err)
		}
	}

	cases := []struct {
		name       string
		accountID  id.AccountID
		status     rentalpb.TripStatus
		wantCount  int
		wantOnlyID string
	}{
		{
			name:      "get_all",
			accountID: "account_id_for_get_trips",
			status:    rentalpb.TripStatus_NOT_SPECIFIED,
			wantCount: 4,
		},
		{
			name:       "get_in_progress",
			accountID:  "account_id_for_get_trips",
			status:     rentalpb.TripStatus_IN_PROGRESS,
			wantCount:  1,
			wantOnlyID: "5f8132eb10714bf629489054",
		},
	}

	for _, cc := range cases {
		t.Run(cc.name, func(t *testing.T) {
			res, err := m.GetTrips(context.Background(), cc.accountID, cc.status)
			if err != nil {
				t.Errorf("cannot get trips: %v", err)
			}

			if cc.wantCount != len(res) {
				t.Errorf("incorrect result count; wnat: %d, got: %d", cc.wantCount, len(res))
			}

			if cc.wantOnlyID != "" && len(res) > 0 {
				if cc.wantOnlyID != res[0].ID.Hex() {
					t.Errorf("only_id incorrect; want: %q, got %q", cc.wantOnlyID, res[0].ID.Hex())
				}
			}
		})
	}
}

func TestUpdateTrip(t *testing.T) {
	c := context.Background()
	mc, err := mongotesting.NewClient(c)
	if err != nil {
		t.Fatalf("cannot connect mongodb: %v", err)
	}

	m := NewMongo(mc.Database("coolcar"))
	tid := id.TripID("5f8132eb12714bf629489054")
	aid := id.AccountID("account_for_update")

	var now int64 = 10000
	m.GenIDFunc = func() primitive.ObjectID {
		return objid.MustFromID(tid)
	}
	m.NowFunc = func() int64 {
		return now
	}

	tr, err := m.CreateTrip(c, &rentalpb.Trip{
		AccountId: aid.String(),
		Status:    rentalpb.TripStatus_IN_PROGRESS,
		Start:     &rentalpb.LocationStatus{PoiName: "start_poi"},
	})
	if err != nil {
		t.Fatalf("cannot create trip: %v", err)
	}
	if tr.UpdateAt != 10000 {
		t.Fatalf("wrong updateat; want: %d, got: %d", now, tr.UpdateAt)
	}

	update := &rentalpb.Trip{
		AccountId: aid.String(),
		Start:     &rentalpb.LocationStatus{PoiName: "start_poi_updated"},
		Status:    rentalpb.TripStatus_IN_PROGRESS,
	}

	cases := []struct {
		name         string
		now          int64
		withUpdateAt int64
		wantErr      bool
	}{
		{
			"normal_update",
			20000,
			10000,
			false,
		},
		{
			"update_with_stal_timestamp",
			30000,
			10000,
			true,
		},
		{
			"update_with_refetch",
			40000,
			20000,
			false,
		},
	}

	for _, cc := range cases {
		now = cc.now
		err := m.UpdateTrip(c, tid, aid, cc.withUpdateAt, update)
		if cc.wantErr {
			if err == nil {
				t.Errorf("%s: want err; got nonw", cc.name)
			} else {
				continue
			}
		} else {
			if err != nil {
				t.Errorf("%s: cannot update: %v", cc.name, err)
			}
		}
		updatedTrip, err := m.GetTrip(c, tid.String(), aid)
		if err != nil {
			t.Errorf("%s: cannot get trip after update: %v", cc.name, err)
		}
		if cc.now != updatedTrip.UpdateAt {
			t.Errorf("%s: incorrect updateat: want %d, got %d", cc.name, cc.now, updatedTrip.UpdatedAtField)
		}
	}
}

func TestMain(m *testing.M) {
	os.Exit(mongotesting.RunWithMongoInDocker(m))
}
