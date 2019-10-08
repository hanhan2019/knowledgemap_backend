package mongo

import (
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Number struct {
	Key   string `bson:"_id"`
	Value int64  `bson:"v"`
}

//////////////      new version
type NumberClient struct {
	*mgo.Collection
}

func NewNumberCollection(c *mgo.Collection) *NumberClient {
	_n := new(NumberClient)
	_n.Collection = c
	return _n
}

func (_n *NumberClient) Inc(key string) (int64, error) {
	var n Number
	n.Key = key
	Change := mgo.Change{
		Update: bson.M{
			"$inc": bson.M{
				"v": 1,
			},
		},
		Upsert:    true,
		ReturnNew: true,
	}
	_, err := _n.FindId(key).Apply(
		Change,
		&n,
	)
	return n.Value, err
}

func (_n *NumberClient) Get(key string) (int64, error) {
	var n Number
	cont := bson.M{
		"_id": key,
	}
	err := _n.Find(cont).One(&n)
	return n.Value, err
}

func (_n *NumberClient) Set(key string, value int64) (retn int64, err error) {
	var n Number
	n.Key = key
	Change := mgo.Change{
		Update: bson.M{
			"$set": bson.M{
				"v": value,
			},
		},
		Upsert:    true,
		ReturnNew: true,
	}
	_, err = _n.FindId(key).Apply(
		Change,
		&n,
	)
	if err != nil {
		return
	}

	retn = n.Value
	return
}
