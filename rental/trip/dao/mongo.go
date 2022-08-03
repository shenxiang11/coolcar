package dao

import (
	"context"
	"fmt"
	rentalpb "github.com/shenxiang11/coolcar/rental-service/gen/go/proto"
	"github.com/shenxiang11/coolcar/shared/id"
	mgo "github.com/shenxiang11/coolcar/shared/mongo"
	"github.com/shenxiang11/coolcar/shared/mongo/objid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

const (
	tripField      = "trip"
	accountIDField = tripField + ".accountid"
	statusField    = tripField + ".status"
)

type Mongo struct {
	col       *mongo.Collection
	GenIDFunc func() primitive.ObjectID
	NowFunc   func() int64
}

func NewMongo(db *mongo.Database) *Mongo {
	return &Mongo{
		col:       db.Collection("trip"),
		GenIDFunc: primitive.NewObjectID,
		NowFunc: func() int64 {
			return time.Now().UnixNano()
		},
	}
}

type TripRecord struct {
	mgo.IDField        `bson:"inline"`
	mgo.UpdatedAtField `bson:"inline"`
	Trip               *rentalpb.Trip `bson:"trip"`
}

func (m *Mongo) CreateTrip(c context.Context, trip *rentalpb.Trip) (*TripRecord, error) {
	r := &TripRecord{
		Trip: trip,
	}
	r.ID = m.GenIDFunc()
	r.UpdateAt = m.NowFunc()

	_, err := m.col.InsertOne(c, r)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (m *Mongo) GetTrip(c context.Context, id string, accountID id.AccountID) (*TripRecord, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("invalid id: %v", err)
	}

	res := m.col.FindOne(c, bson.M{
		mgo.IDFieldName: objID,
		accountIDField:  accountID,
	})

	if err := res.Err(); err != nil {
		return nil, err
	}

	var tr TripRecord
	err = res.Decode(&tr)
	if err != nil {
		return nil, fmt.Errorf("cannot decode: %v", err)
	}

	return &tr, nil
}

func (m *Mongo) GetTrips(c context.Context, accountID id.AccountID, status rentalpb.TripStatus) ([]*TripRecord, error) {
	filter := bson.M{
		accountIDField: accountID,
	}
	if status != rentalpb.TripStatus_NOT_SPECIFIED {
		filter[statusField] = status
	}

	res, err := m.col.Find(c, filter, options.Find().SetSort(bson.M{mgo.IDFieldName: -1}))
	if err != nil {
		return nil, err
	}

	var trips []*TripRecord
	for res.Next(c) {
		var trip TripRecord
		err := res.Decode(&trip)
		if err != nil {
			return nil, err
		}
		trips = append(trips, &trip)
	}

	return trips, nil
}

func (m *Mongo) UpdateTrip(c context.Context, tid id.TripID, aid id.AccountID, updateAt int64, trip *rentalpb.Trip) error {
	objID, err := objid.FromID(tid)
	if err != nil {
		return fmt.Errorf("invalid id: %v", err)
	}

	newUpdateAt := m.NowFunc()
	res, err := m.col.UpdateOne(c, bson.M{
		mgo.IDFieldName:        objID,
		accountIDField:         aid.String(),
		mgo.UpdatedAtFieldName: updateAt,
	}, mgo.Set(bson.M{
		tripField:              trip,
		mgo.UpdatedAtFieldName: newUpdateAt,
	}))

	if err != nil {
		return err
	}

	if res.MatchedCount == 0 {
		return mongo.ErrNoDocuments
	}

	return nil
}
