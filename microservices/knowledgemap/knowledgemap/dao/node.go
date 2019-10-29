package dao

import (
	"context"

	"knowledgemap_backend/microservices/knowledgemap/knowledgemap/model"

	"gopkg.in/mgo.v2/bson"
)

func (d *Dao) QueryRelationByNodeID(ctx context.Context, id bson.ObjectId, relations *[]*model.Relation) error {
	db := d.mdb.Copy()
	defer db.Session.Close()
	if err := db.C(model.RELATION_NAME).FindId(id).All(*relations); err != nil {
		return err
	}
	return nil
}
