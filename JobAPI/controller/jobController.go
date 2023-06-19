package controller

import (
	"JobAPI/logger"
	"JobAPI/models"
	"JobAPI/service"
	"fmt"
)

type IJobController interface {
	GetJobs(filter *models.JobFilter) (*models.GetJobResponse, error)
}

type jobController struct {
	service service.IJobService
	logger  logger.ILogger
}

var jobControllerObj IJobController

func GetJobController() (IJobController, error) {
	if jobControllerObj == nil {
		return nil, fmt.Errorf("Job Controller not initilized")
	}
	return jobControllerObj, nil
}
func InitJobController(jobServiceObj service.IJobService, logObj logger.ILogger) IJobController {
	if jobControllerObj == nil {
		jobControllerObj = &jobController{
			service: jobServiceObj,
			logger:  logObj,
		}
	}
	return jobControllerObj
}

func (ctlr *jobController) GetJobs(filter *models.JobFilter) (*models.GetJobResponse, error) {
	return ctlr.service.GetJobs(filter, filter.PageSize, filter.PageNumber)
}
