package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jcocozza/cassidy-wails/internal/database"
	"github.com/jcocozza/cassidy-wails/internal/model"
	conversionrepo "github.com/jcocozza/cassidy-wails/internal/repository/conversionRepo"
	equipmentrepo "github.com/jcocozza/cassidy-wails/internal/repository/equipmentRepo"
	ctxutil "github.com/jcocozza/cassidy-wails/internal/utils/ctxUtil"
)

type EquipmentHandler struct {
	EquipmentRepository  equipmentrepo.EquipmentRepository
	ConversionRepository conversionrepo.MeasurementRepository
}

func NewEquipmentHandler(db database.DbOperations) *EquipmentHandler {
	return &EquipmentHandler{
		EquipmentRepository:  equipmentrepo.NewIEquipmentRepository(db),
		ConversionRepository: conversionrepo.NewIMeasurementRepository(),
	}
}

// Create equipment
func (eh *EquipmentHandler) CreateEquipment(ctx *gin.Context) {
	userUuid := ctxutil.GetUserUuidFromHeader(ctx)

	var createRequest model.Equipment
	if err := ctx.ShouldBindJSON(&createRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err0 := eh.ConversionRepository.ConvertEquipment(conversionrepo.Incoming, &createRequest)
	if err0 != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "equipment failed to convert incoming: " + err0.Error()})
		return
	}

	createRequest.UserUuid = userUuid
	err1 := createRequest.Validate()
	if err1 != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err1.Error()})
		return
	}
	id, err2 := eh.EquipmentRepository.Create(&createRequest)
	if err2 != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err2.Error()})
		return
	}
	createRequest.SetId(id)

	err3 := eh.ConversionRepository.ConvertEquipment(conversionrepo.Outgoing, &createRequest)
	if err3 != nil {
		ctx.JSON(http.StatusBadRequest, "equipment failed to convert outgoing: "+err3.Error())
		return
	}
	ctx.JSON(http.StatusCreated, createRequest)
}

// Update Equipment
//
// @param: id
func (eh *EquipmentHandler) UpdateEquipment(ctx *gin.Context) {
	equipmentIdStr := ctx.Param("id")
	_, err := strconv.Atoi(equipmentIdStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid equipment id"})
	}

	var updateRequest model.Equipment
	if err := ctx.ShouldBindJSON(&updateRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err0 := eh.ConversionRepository.ConvertEquipment(conversionrepo.Incoming, &updateRequest)
	if err0 != nil {
		ctx.JSON(http.StatusBadRequest, "equipment failed to convert: "+err0.Error())
		return
	}

	err1 := eh.EquipmentRepository.Update(&updateRequest)
	if err1 != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "failed to update equipment:" + err1.Error()})
		return
	}

	err2 := eh.ConversionRepository.ConvertEquipment(conversionrepo.Outgoing, &updateRequest)
	if err2 != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "failed to convert equipment:" + err2.Error()})
		return
	}
	ctx.JSON(http.StatusAccepted, updateRequest)
}

// Delete Equipment
//
// @param: id
func (eh *EquipmentHandler) DeleteEquipment(ctx *gin.Context) {
	equipmentIdStr := ctx.Param("id")
	equipmentId, err := strconv.Atoi(equipmentIdStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid equipment id"})
	}

	err1 := eh.EquipmentRepository.Delete(equipmentId)
	if err1 != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete equipment"})
		return
	}
	ctx.JSON(http.StatusAccepted, nil)
}

// List all equipment
func (eh *EquipmentHandler) List(ctx *gin.Context) {
	userUuid := ctxutil.GetUserUuidFromHeader(ctx)

	equipmentList, err := eh.EquipmentRepository.List(userUuid)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get equipment list:" + fmt.Sprint(err)})
		return
	}

	for _, equipment := range equipmentList {
		err0 := eh.ConversionRepository.ConvertEquipment(conversionrepo.Outgoing, equipment)
		if err0 != nil {
			ctx.JSON(http.StatusBadRequest, "equipment failed to convert")
			return
		}
	}

	ctx.JSON(http.StatusOK, equipmentList)
}

// List equipment types
func (eh *EquipmentHandler) ListEquipmentTypes(ctx *gin.Context) {
	equipmentTypesList, err := eh.EquipmentRepository.ListEquipmentTypes()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get equipment type list: " + err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, equipmentTypesList)
}

// Create activity equipment
func (eh *EquipmentHandler) CreateActivityEquipment(ctx *gin.Context) {
	var createRequest model.ActivityEquipment
	if err := ctx.ShouldBindJSON(&createRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err0 := eh.ConversionRepository.ConvertActivityEquipment(conversionrepo.Incoming, &createRequest)
	if err0 != nil {
		ctx.JSON(http.StatusBadRequest, "activity equipment failed to convert")
		return
	}

	id, err := eh.EquipmentRepository.CreateActivityEquipment(&createRequest)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get create activity equipment:" + fmt.Sprint(err)})
		return
	}
	createRequest.SetId(id)

	err3 := eh.ConversionRepository.ConvertActivityEquipment(conversionrepo.Outgoing, &createRequest)
	if err3 != nil {
		ctx.JSON(http.StatusBadRequest, "activity equipment failed to convert")
		return
	}
	ctx.JSON(http.StatusCreated, createRequest)
}

// Update Activity Equipment
// Update the mileage of an activity equipment (a patch for anything else in that doesn't make sense)
//
// @param: id
func (eh *EquipmentHandler) UpdateActivityEquipment(ctx *gin.Context) {
	activityEquipmentIdStr := ctx.Param("id")
	_, err := strconv.Atoi(activityEquipmentIdStr)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid activity equipment id"})
		return
	}
	var updateRequest model.ActivityEquipment
	if err := ctx.ShouldBindJSON(&updateRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err0 := eh.ConversionRepository.ConvertActivityEquipment(conversionrepo.Incoming, &updateRequest)
	if err0 != nil {
		ctx.JSON(http.StatusBadRequest, "activity equipment failed to convert")
		return
	}

	err1 := eh.EquipmentRepository.UpdateActivityEquipment(&updateRequest)
	if err1 != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "failed to update activity equipment"})
		return
	}

	err3 := eh.ConversionRepository.ConvertActivityEquipment(conversionrepo.Outgoing, &updateRequest)
	if err3 != nil {
		ctx.JSON(http.StatusBadRequest, "activity equipment failed to convert")
		return
	}
	ctx.JSON(http.StatusAccepted, updateRequest)
}

// Delete Activity Equipment
//
// @param: id
func (eh *EquipmentHandler) DeleteActivityEquipment(ctx *gin.Context) {
	activityEquipmentIdStr := ctx.Param("id")
	activityEquipmentId, err := strconv.Atoi(activityEquipmentIdStr)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid activity equipment id"})
		return
	}

	err1 := eh.EquipmentRepository.DeleteActivityEquipment(activityEquipmentId)
	if err1 != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete activity equipment"})
		return
	}
	ctx.JSON(http.StatusAccepted, nil)

}
