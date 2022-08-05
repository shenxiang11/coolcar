package dao

import (
	"fmt"
	mgo "github.com/shenxiang11/coolcar/shared/mongo"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testing"
)

func TestPlayground(t *testing.T) {
	br := &BlobRecord{
		AccountID: "",
		Path:      "",
	}

	objID := mgo.NewObjID()
	fmt.Println(objID)
	fmt.Println(br.ID)

	fmt.Printf("TYPE: %T\nTYPE: %T", objID, br.ID)

	br.ID = primitive.ObjectID(objID)
	fmt.Println(br.ID)
}
