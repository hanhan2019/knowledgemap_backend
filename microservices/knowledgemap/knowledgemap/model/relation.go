package model

import "gopkg.in/mgo.v2/bson"

const (
	RELATION_NAME = "relation"
)

type RelationKind string

const (
	RelationContain RelationKind = "contains"
	RelationDerive  RelationKind = "derives"
	RelationAffect  RelationKind = "affects"
)

type Relation struct {
	ID           bson.ObjectId `bson:"_id" json:"_id"`
	MainNodeID   bson.ObjectId `bson:"mainnodeid" json:"mainnodeid"`
	ObjectNodeID bson.ObjectId `bson:"objectnodeid" json:"objectnodeid"`
	Relation     RelationKind  `bson:"relation" json:"relation"`
}
