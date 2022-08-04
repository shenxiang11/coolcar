package trip

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/go-cmp/cmp"
	rentalpb "github.com/shenxiang11/coolcar/rental-service/gen/go/proto"
	"github.com/shenxiang11/coolcar/rental-service/trip/dao"
	"github.com/shenxiang11/coolcar/rental-service/trip/manager/poi"
	"github.com/shenxiang11/coolcar/shared/auth"
	"github.com/shenxiang11/coolcar/shared/id"
	"github.com/shenxiang11/coolcar/shared/mongo/objid"
	"github.com/shenxiang11/coolcar/shared/server"
	mongotesting "github.com/shenxiang11/coolcar/shared/testing"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func TestCreateTrip(t *testing.T) {
	c := context.Background()

	pm := &profileManager{}
	cm := &carManager{}
	s := newService(c, t, pm, cm)
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

	cases := []struct {
		name         string
		accountID    id.AccountID
		tripID       id.TripID
		profileErr   error
		carVerifyErr error
		carUnlockErr error
		want         string
		wantErr      bool
	}{
		{
			name:      "normal_create",
			accountID: "account1",
			tripID:    "5f8132eb12714bf629489054",
			want:      fmt.Sprintf(golden, "account1"),
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
	}

	for _, cc := range cases {
		t.Run(cc.name, func(t *testing.T) {
			s.Mongo.GenIDFunc = func() primitive.ObjectID {
				return objid.MustFromID(cc.tripID)
			}
			pm.err = cc.profileErr
			cm.unlockErr = cc.carUnlockErr
			cm.verifyErr = cc.carVerifyErr
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

func newService(c context.Context, t *testing.T, pm ProfileManager, cm CarManager) *Service {
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
		POIManager:     &poi.Manager{},
		DistanceCalc:   &distCalc{},
	}
}

func TestMain(m *testing.M) {
	os.Exit(mongotesting.RunWithMongoInDocker(m))
}
