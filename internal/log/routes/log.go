package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/viniqrz/gin-crud/internal/log/repository"
	"github.com/viniqrz/gin-crud/internal/log/use_cases/create_log"
	"github.com/viniqrz/gin-crud/internal/log/use_cases/delete_log"
	"github.com/viniqrz/gin-crud/internal/log/use_cases/find_all_logs"
)

func GetLogRoutes(rg *gin.RouterGroup, db *sql.DB) {
	logRouter := rg.Group("/log")
	logRepository := repository.NewLogPSQLRepository(db)

	createLogUseCase := create_log.NewCreateLogUseCase(logRepository)
	createLogUseCaseController := create_log.NewCreateLogUseCaseController(createLogUseCase)

	findAllLogsUseCase := find_all_logs.NewFindAllLogsUseCase(logRepository)
	findAllLogsUseCaseController := find_all_logs.NewFindAllLogsUseCaseController(findAllLogsUseCase)

	deleteLogUseCase := delete_log.NewDeleteLogUseCase(logRepository)
	deleteLogUseCaseController := delete_log.NewDeleteLogUseCaseController(deleteLogUseCase)

	logRouter.POST("/", HandleRoute(createLogUseCaseController.Execute))
	logRouter.GET("/", HandleRoute(findAllLogsUseCaseController.Execute))
	logRouter.DELETE("/:id", HandleRoute(deleteLogUseCaseController.Execute))
}