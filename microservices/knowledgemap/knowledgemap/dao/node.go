package dao

import (
	"context"
	"knowledgemap_backend/microservices/knowledgemap/knowledgemap/model"

	"gopkg.in/mgo.v2/bson"
)

func (d *Dao) QueryNodeInfoByNodeID(ctx context.Context, id bson.ObjectId, node *model.Node) error {
	db := d.mdb.Copy()
	defer db.Session.Close()
	if err := db.C(model.NODE_NAME).FindId(id).One(node); err != nil {
		return err
	}
	return nil
}
