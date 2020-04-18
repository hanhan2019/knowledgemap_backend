package dao

import (
	"context"
	"knowledgemap_backend/microservices/knowledgemap/class/api"
	"knowledgemap_backend/microservices/knowledgemap/class/model"
	"time"
)

func (d *Dao) newInvitaion(ctx context.Context, invitation *model.Invitation) (*model.Invitation, error) {
	db := d.mdb.Copy()
	defer db.Session.Close()
	col := db.C(model.Invitation_COLLECTION_NAME)
	return invitation, col.Insert(invitation)
}

func (d *Dao) NewInvitation(ctx context.Context, req *api.InvitationReq) (*model.Invitation, error) {
	invitaion := createDefaultInvitation(req)
	return d.newInvitaion(ctx, invitaion)
}

func createDefaultInvitation(req *api.InvitationReq) *model.Invitation {
	invitation := new(model.Invitation)
	invitation.InvitationCode = req.Invitaioncode
	invitation.ClassId = req.Classid
	invitation.UserId = req.Userid
	invitation.CreateTime = time.Now().Unix()
	return invitation
}

func (d *Dao) StopInvitaion(ctx context.Context, invitation string) error {
	db := d.mdb.Copy()
	defer db.Session.Close()
	col := db.C(model.Invitation_COLLECTION_NAME)
	return col.RemoveId(invitation)
}

func (d *Dao) FillInvitaion(ctx context.Context, invitationCode string, invitation *model.Invitation) (err error) {
	db := d.mdb.Copy()
	defer db.Session.Close()
	if invitation == nil {
		invitation = &model.Invitation{}
	}
	err = db.C(model.Invitation_COLLECTION_NAME).FindId(invitationCode).One(invitation)
	return
}
