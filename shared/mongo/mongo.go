package mgo

import (
	"fmt"
	"github.com/shenxiang11/coolcar/shared/mongo/objid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const IDFieldName = "_id"

type IDField struct {
	ID primitive.ObjectID `bson:"_id"`
}

var NewObjID = primitive.NewObjectID

func MockNewObjIdWithValue(id fmt.Stringer) func() primitive.ObjectID {
	return func() primitive.ObjectID {
		return objid.MustFromID(id)
	}
}

func Set(v any) bson.M {
	return bson.M{
		"$set": v,
	}
}

func SetOnInsert(v any) bson.M {
	return bson.M{
		"$setOnInsert": v,
	}
}
