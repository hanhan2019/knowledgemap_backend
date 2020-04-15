package model

const (
	Invitation_COLLECTION_NAME = "invitation"
)

type Invitation struct {
	InvitationCode string `bson:"_id" json:"_id"`
	ClassId        string `bson:"name" json:"name"`
	UserId         string `bson:"userid" json:"userid"`
	CreateTime     int64  `bson:"createtime" json:"createtime"`
	DropTime       int64  `bson:"droptime" json:"droptime"`
}
