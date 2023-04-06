package frontend

import (
	"mopoen-remake/controller/helper"
	"mopoen-remake/controller/request"
	svc "mopoen-remake/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ISensorController interface {
	Create(ctx *gin.Context)
	Delete(ctx *gin.Context)
	Get(ctx *gin.Context)
	GetAll(ctx *gin.Context)
}

type SensorController struct {
	service svc.IServices
}

func NewSensorController(service svc.IServices) SensorController {
	return SensorController{service: service}
}

func (ctr SensorController) CreateNewTipeSensor(ctx *gin.Context) {
	req := request.PostSensor{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	err := ctr.service.CreateTipeSensor(ctx, req.Tipe, req.Satuan)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "created",
	})
}

func (ctr SensorController) GetAllTipeSensor(ctx *gin.Context) {
	sensors, err := ctr.service.GetAllTipeSensor(ctx)
	if err != nil {
		helper.RespBadRequest(ctx, err.Error())
		return
	}

	helper.RespStatusOk(ctx, sensors)
}

// GetTipeSensorByID is a function to get tipe sensor by id
func (ctr SensorController) GetTipeSensorByID(ctx *gin.Context) {
	var idSensor request.GetSensor
	if err := ctx.ShouldBindUri(&idSensor); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	tipeSensor, err := ctr.service.GetTipeSensor(ctx, idSensor.Id)
	if err != nil {
		helper.RespCatchSqlErr(ctx, err)
		return
	}

	helper.RespStatusOk(ctx, tipeSensor)
}

func (ctr SensorController) DeleteTipeSensor(ctx *gin.Context) {
	var req request.DeleteSensor
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	ctr.service.DeleteTipeSensor(ctx, req.Id)
}
