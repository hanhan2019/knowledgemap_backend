package dao

import (
	"context"
	"fmt"
	"knowledgemap_backend/library/database/mongo"
	"knowledgemap_backend/microservices/knowledgemap/knowledgemap/model"
	pmodel "knowledgemap_backend/microservices/knowledgemap/passport/model"

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

func (d *Dao) NewNode(ctx context.Context, node *model.Node) (*model.Node, error) {
	db := d.mdb.Copy()
	defer db.Session.Close()
	genIdCol := mongo.NewNumberCollection(db.C(pmodel.GEN_ID_COLLECTION_NAME))
	nodeId, _ := genIdCol.Inc(pmodel.GEN_NODE_ID_KEY_NAME)
	node.Kind = fmt.Sprintf("2-1-1-1-%v", nodeId)
	col := db.C(model.NODE_NAME)
	return node, col.Insert(node)
}

func (d *Dao) FillKnowledgeByID(ctx context.Context, id bson.ObjectId, node *model.Node) (err error) {
	db := d.mdb.Copy()
	defer db.Session.Close()
	err = db.C(model.NODE_NAME).FindId(id).One(node)
	return
}
