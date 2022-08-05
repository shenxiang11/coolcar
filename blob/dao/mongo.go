package dao

import (
	"context"
	"fmt"
	"github.com/shenxiang11/coolcar/shared/id"
	mgo "github.com/shenxiang11/coolcar/shared/mongo"
	"github.com/shenxiang11/coolcar/shared/mongo/objid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Mongo struct {
	col *mongo.Collection
}

func NewMongo(db *mongo.Database) *Mongo {
	return &Mongo{
		col: db.Collection("blob"),
	}
}

type BlobRecord struct {
	mgo.IDField `bson:"inline"`
	AccountID   string `bson:"accountid"`
	Path        string `bson:"path"`
}

func (m *Mongo) CreateBlob(c context.Context, aid id.AccountID) (*BlobRecord, error) {
	br := &BlobRecord{
		AccountID: aid.String(),
	}
	objID := mgo.NewObjID()
	br.ID = primitive.ObjectID(objID) // FIXME: WHY
	br.Path = fmt.Sprintf("%s/%s", aid.String(), objID.Hex())

	_, err := m.col.InsertOne(c, br)
	if err != nil {
		return nil, err
	}

	return br, nil
}

func (m *Mongo) GetBlob(c context.Context, bid id.BlobID) (*BlobRecord, error) {
	objID, err := objid.FromID(bid)
	if err != nil {
		return nil, fmt.Errorf("invalid object id: %v", err)
	}

	res := m.col.FindOne(c, bson.M{
		mgo.IDFieldName: objID,
	})
	if err := res.Err(); err != nil {
		return nil, err
	}

	var br BlobRecord
	err = res.Decode(&br)
	if err != nil {
		return nil, fmt.Errorf("cannot decode result: %v", err)
	}

	return &br, nil
}
