package trip

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/go-cmp/cmp"
	rentalpb "github.com/shenxiang11/coolcar/rental-service/gen/go/proto"
	"github.com/shenxiang11/coolcar/rental-service/trip/dao"
	"github.com/shenxiang11/coolcar/shared/auth"
	"github.com/shenxiang11/coolcar/shared/id"
	"github.com/shenxiang11/coolcar/shared/mongo/objid"
	"github.com/shenxiang11/coolcar/shared/server"
	mongotesting "github.com/shenxiang11/coolcar/shared/testing"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"math/rand"
	"os"
	"testing"
)

type profileManager struct {
	iID id.IdentityID
	err error
}

func (p *profileManager) Verify(ctx context.Context, id id.AccountID) (id.IdentityID, error) {
	return p.iID, p.err
}

type carManager struct {
	verifyErr error
	unlockErr error
	lockErr   error
}

func (car *carManager) Verify(c context.Context, cid id.CarID, loc *rentalpb.Location) error {
	return car.verifyErr
}

func (car *carManager) Unlock(c context.Context, cid id.CarID, aid id.AccountID, tid id.TripID, avatarURL string) error {
	return car.unlockErr
}

func (car carManager) Lock(c context.Context, cid id.CarID) error {
	return car.lockErr
}

type distCalc struct {
}

func (d *distCalc) DistanceKm(ctx context.Context, from *rentalpb.Location, to *rentalpb.Location) (float64, error) {
	if from.Latitude == to.Latitude && from.Longitude == to.Longitude {
		return 0, nil
	}
	return 100, nil
}

type poiManager struct {
	resolveErr error
}

func (p *poiManager) Resolve(ctx context.Context, location *rentalpb.Location) (string, error) {
	if p.resolveErr != nil {
		return "", p.resolveErr
	}
	return "桥梁", nil
}

func TestCreateTrip(t *testing.T) {
	c := context.Background()

	pm := &profileManager{}
	cm := &carManager{}
	poi := &poiManager{}
	s := newService(c, t, pm, cm, poi)
	s.NowFun = func() int64 {
		return 1605695246
	}

	req := &rentalpb.CreateTripRequest{
		CarId: "car1",
		Start: &rentalpb.Location{
			Latitude:  32.123,
			Longitude: 114.2525,
		},
	}
	pm.iID = "identity1"
	golden := `{"account_id":%q,"car_id":"car1","start":{"location":{"latitude":32.123,"longitude":114.2525},"poi_name":"桥梁","timestamp_sec":1605695246},"current":{"location":{"latitude":32.123,"longitude":114.2525},"poi_name":"桥梁","timestamp_sec":1605695246},"status":1,"identity_id":"identity1"}`
	goldenWithNoPoi := `{"account_id":%q,"car_id":"car1","start":{"location":{"latitude":32.123,"longitude":114.2525},"timestamp_sec":1605695246},"current":{"location":{"latitude":32.123,"longitude":114.2525},"timestamp_sec":1605695246},"status":1,"identity_id":"identity1"}`

	cases := []struct {
		name          string
		accountID     id.AccountID
		tripID        id.TripID
		profileErr    error
		carVerifyErr  error
		carUnlockErr  error
		poiResolveErr error
		want          string
		wantErr       bool
	}{
		{
			name:      "normal_create",
			accountID: "account1",
			tripID:    "5f8132eb12714bf629489054",
			want:      fmt.Sprintf(golden, "account1"),
		},
		{
			name:      "repeat_create",
			accountID: "account1",
			tripID:    "5f8132eb12714bf629489054",
			wantErr:   true,
		},
		{
			name:       "profile_err",
			accountID:  "account2",
			tripID:     "5f8132eb12714bf629489055",
			profileErr: fmt.Errorf("profile"),
			wantErr:    true,
		},
		{
			name:         "car_verify_err",
			accountID:    "account3",
			tripID:       "5f8132eb12714bf629489056",
			carVerifyErr: fmt.Errorf("verify"),
			wantErr:      true,
		},
		{
			name:         "car_unlock_err",
			accountID:    "account4",
			tripID:       "5f8132eb12714bf629489057",
			carUnlockErr: fmt.Errorf("unlock"),
			want:         fmt.Sprintf(golden, "account4"),
		},
		{
			name:          "poi_resolve_err",
			accountID:     "account5",
			tripID:        "5f8132eb12714bf629489058",
			poiResolveErr: fmt.Errorf("poi_resolve_error"),
			want:          fmt.Sprintf(goldenWithNoPoi, "account5"),
		},
	}

	for _, cc := range cases {
		t.Run(cc.name, func(t *testing.T) {
			s.Mongo.GenIDFunc = func() primitive.ObjectID {
				return objid.MustFromID(cc.tripID)
			}
			pm.err = cc.profileErr
			cm.unlockErr = cc.carUnlockErr
			cm.verifyErr = cc.carVerifyErr
			poi.resolveErr = cc.poiResolveErr

			c := auth.ContextWithAccount(c, cc.accountID)
			res, err := s.CreateTrip(c, req)
			if cc.wantErr {
				if err == nil {
					t.Errorf("want error; got none")
				} else {
					return
				}
			}
			if err != nil {
				t.Errorf("errpr creating trip: %v", err)
				return
			}
			if res.Id != cc.tripID.String() {
				t.Errorf("incorrect id; want %q, got %q", cc.tripID, res.Id)
			}
			b, err := json.Marshal(res.Trip)
			if err != nil {
				t.Errorf("cannot marshall response: %v", err)
			}
			got := string(b)
			if cc.want != got {
				diff := cmp.Diff(cc.want, got)
				t.Errorf("incorrect response: -want +got: %s", diff)
			}
		})
	}
}

func TestTripLifecycle(t *testing.T) {
	c := auth.ContextWithAccount(context.Background(), id.AccountID("account_for_lifecycle"))
	s := newService(c, t, &profileManager{}, &carManager{}, &poiManager{})

	tid := id.TripID("5f8132eb22714bf629489056")
	s.Mongo.GenIDFunc = func() primitive.ObjectID {
		return objid.MustFromID(tid)
	}
	cases := []struct {
		name    string
		now     int64
		op      func() (*rentalpb.Trip, error)
		want    string
		wantErr bool
	}{
		{
			name: "create_trip",
			now:  10000,
			op: func() (*rentalpb.Trip, error) {
				e, err := s.CreateTrip(c, &rentalpb.CreateTripRequest{
					Start: &rentalpb.Location{
						Latitude:  32.123,
						Longitude: 114.2525,
					},
					CarId: "Car1",
				})
				if err != nil {
					return nil, err
				}
				return e.Trip, nil
			},
			want: `{"account_id":"account_for_lifecycle","car_id":"Car1","start":{"location":{"latitude":32.123,"longitude":114.2525},"poi_name":"桥梁","timestamp_sec":10000},"current":{"location":{"latitude":32.123,"longitude":114.2525},"poi_name":"桥梁","timestamp_sec":10000},"status":1}`,
		},
		{
			name: "update_trip",
			now:  20000,
			op: func() (*rentalpb.Trip, error) {
				return s.UpdateTrip(c, &rentalpb.UpdateTripRequest{
					Id: tid.String(),
					Current: &rentalpb.Location{
						Latitude:  28.234234,
						Longitude: 123.243255,
					},
				})
			},
			want: `{"account_id":"account_for_lifecycle","car_id":"Car1","start":{"location":{"latitude":32.123,"longitude":114.2525},"poi_name":"桥梁","timestamp_sec":10000},"current":{"location":{"latitude":28.234234,"longitude":123.243255},"fee_cent":7968,"km_driven":100,"poi_name":"桥梁","timestamp_sec":20000},"status":1}`,
		},
		{
			name: "finish_trip",
			now:  30000,
			op: func() (*rentalpb.Trip, error) {
				return s.UpdateTrip(c, &rentalpb.UpdateTripRequest{
					Id:      tid.String(),
					EndTrip: true,
				})
			},
			want: `{"account_id":"account_for_lifecycle","car_id":"Car1","start":{"location":{"latitude":32.123,"longitude":114.2525},"poi_name":"桥梁","timestamp_sec":10000},"current":{"location":{"latitude":28.234234,"longitude":123.243255},"fee_cent":11825,"km_driven":100,"poi_name":"桥梁","timestamp_sec":30000},"end":{"location":{"latitude":28.234234,"longitude":123.243255},"fee_cent":11825,"km_driven":100,"poi_name":"桥梁","timestamp_sec":30000},"status":2}`,
		},
		{
			name: "query_trip",
			now:  40000,
			op: func() (*rentalpb.Trip, error) {
				return s.GetTrip(c, &rentalpb.GetTripRequest{
					Id: tid.String(),
				})
			},
			want: `{"account_id":"account_for_lifecycle","car_id":"Car1","start":{"location":{"latitude":32.123,"longitude":114.2525},"poi_name":"桥梁","timestamp_sec":10000},"current":{"location":{"latitude":28.234234,"longitude":123.243255},"fee_cent":11825,"km_driven":100,"poi_name":"桥梁","timestamp_sec":30000},"end":{"location":{"latitude":28.234234,"longitude":123.243255},"fee_cent":11825,"km_driven":100,"poi_name":"桥梁","timestamp_sec":30000},"status":2}`,
		},
		{
			name: "update_after_finished",
			now:  50000,
			op: func() (*rentalpb.Trip, error) {
				return s.UpdateTrip(c, &rentalpb.UpdateTripRequest{
					Id: tid.String(),
				})
			},
			wantErr: true,
		},
	}

	rand.Seed(1345)
	for _, cc := range cases {
		s.NowFun = func() int64 {
			return cc.now
		}
		trip, err := cc.op()
		if cc.wantErr {
			if err == nil {
				t.Errorf("%s: want error; got none", cc.name)
			} else {
				continue
			}
		}
		if err != nil {
			t.Errorf("%s: operation failed: %v", cc.name, err)
			continue
		}
		b, err := json.Marshal(trip)
		if err != nil {
			t.Errorf("%s: failed marshalling response: %v", cc.name, err)
		}
		got := string(b)
		if cc.want != got {
			diff := cmp.Diff(cc.want, got)
			fmt.Println(got)
			t.Errorf("%s: incorrect response; -want +got: %s", cc.name, diff)
		}
	}
}

func newService(c context.Context, t *testing.T, pm ProfileManager, cm CarManager, poim POIManager) *Service {
	mc, err := mongotesting.NewClient(c)
	if err != nil {
		t.Fatalf("cannot create mongo client: %v", err)
	}

	logger, err := server.NewZapLogger()
	if err != nil {
		t.Fatalf("cannot create logger: %v", err)
	}

	db := mc.Database("coolcar")
	mongotesting.SetupIndexes(c, db)

	return &Service{
		Mongo:          dao.NewMongo(db),
		Logger:         logger,
		ProfileManager: pm,
		CarManager:     cm,
		POIManager:     poim,
		DistanceCalc:   &distCalc{},
	}
}

func TestMain(m *testing.M) {
	os.Exit(mongotesting.RunWithMongoInDocker(m))
}
