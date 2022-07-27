package dao

import (
	"context"
	mgo "github.com/shenxiang11/coolcar/shared/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const openIDField = "open_id"

type Mongo struct {
	col   *mongo.Collection
	genID func() primitive.ObjectID
}

func NewMongo(db *mongo.Database) *Mongo {
	return &Mongo{
		col:   db.Collection("account"),
		genID: primitive.NewObjectID,
	}
}

func NewMongoWithMock(db *mongo.Database, genIDFunc func() primitive.ObjectID) *Mongo {
	return &Mongo{
		col:   db.Collection("account"),
		genID: genIDFunc,
	}
}

func (m *Mongo) ResolveAccountID(c context.Context, openID string) (string, error) {
	insertID := m.genID()
	res := m.col.FindOneAndUpdate(c, bson.M{
		openIDField: openID,
	}, mgo.SetOnInsert(bson.M{
		mgo.IDFieldName: insertID,
		openIDField:     openID,
	}), options.FindOneAndUpdate().SetUpsert(true).SetReturnDocument(options.After))

	if err := res.Err(); err != nil {
		return "", err
	}

	var row mgo.IDField
	err := res.Decode(&row)
	if err != nil {
		return "", err
	}

	return row.ID.Hex(), nil
}
