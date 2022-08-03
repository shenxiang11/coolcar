package dao

import (
	"context"
	"fmt"
	rentalpb "github.com/shenxiang11/coolcar/rental/gen/go/proto"
	"github.com/shenxiang11/coolcar/shared/id"
	mgo "github.com/shenxiang11/coolcar/shared/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

const (
	tripField      = "trip"
	accountIDField = tripField + ".accountid"
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
