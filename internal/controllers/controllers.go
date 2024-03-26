package controllers

import (
	"github.com/jcocozza/cassidy-wails/internal/database"
	"github.com/jcocozza/cassidy-wails/internal/model"
)

type Controllers struct {
	UserHandler *UserHandler
	ActivityHandler *ActivityHandler
	ActivityTypeHandler *ActivityTypeHandler
	EquipmentHandler *EquipmentHandler
	MicrocycleHandler *MicrocycleHandler
	MiscHandler *MiscHandler
}
func NewControllers(db database.DbOperations, user *model.User) *Controllers {
	return &Controllers{
		UserHandler: NewUserHandler(db, user),
		ActivityHandler: NewActivityHandler(db, user),
		ActivityTypeHandler: NewActivityTypeHandler(db),
		EquipmentHandler: NewEquipmentHandler(db, user),
		MicrocycleHandler: NewMicrocycleHandler(db, user),
		MiscHandler: NewMiscHandler(db, user),
	}
}
func (c *Controllers) SetUser(user *model.User) {
	c.UserHandler.User = user
	c.ActivityHandler.User = user
	c.EquipmentHandler.User = user
	c.MicrocycleHandler.User = user
	c.MiscHandler.User = user
}