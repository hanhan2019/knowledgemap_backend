package model

import (
	"gopkg.in/mgo.v2/bson"
)

const (
	NODE_NAME = "node"
)

type NodeType string

const (
	NodeConcept    NodeType = "Concept"
	NodeCollection NodeType = "Collection"
)

type Node struct {
	ID    bson.ObjectId `bson:"_id" json:"_id"` //仅为数据库存储的唯一id
	Label []LableInfo   `bson:"babel" json:"label"`
	Kind  string        `bson:"id" json:"id"` //对应jsonld 里的实体id
	Type  NodeType      `bson:"type" json:"type"`
}

type LableInfo struct {
	Language string `bson:"language" json:"language"`
	Value    string `bson:"value" json:"value"`
}
