package controllers

import "github.com/jcocozza/cassidy-wails/internal/database"

type Controllers struct {
	UserHandler *UserHandler
	ActivityHandler *ActivityHandler
	ActivityTypeHandler *ActivityTypeHandler
	EquipmentHandler *EquipmentHandler
	MicrocycleHandler *MicrocycleHandler
	MiscHandler *MiscHandler
}

func NewControllers(db database.DbOperations) *Controllers {
	return &Controllers{
		UserHandler: NewUserHandler(db),
		ActivityHandler: NewActivityHandler(db),
		ActivityTypeHandler: NewActivityTypeHandler(db),
		EquipmentHandler: NewEquipmentHandler(db),
		MicrocycleHandler: NewMicrocycleHandler(db),
		MiscHandler: NewMiscHandler(db),
	}
}