package router

import (
	event 	    "event-management-service/event/handler"
	workshop    "event-management-service/workshop/handler"
	reservation "event-management-service/reservation/handler"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Route(eventDB *gorm.DB) *gin.Engine{
	eventHandler := event.NewEventHandler(eventDB)
	workshopHandler := workshop.NewWorkshopHandler(eventDB)
	reservationHandler := reservation.NewReservationHandler(eventDB)
	
	router := gin.Default()
	baseRouter := router.Group("/api")

	eventRouter := baseRouter.Group("/events")
	eventRouter.GET("", func(ctx *gin.Context){ eventHandler.List(ctx)})
	eventRouter.GET("/:id", func(ctx *gin.Context){ eventHandler.Detail(ctx)})

	workshopRouter := baseRouter.Group("/workshops")
	workshopRouter.GET("/:eventID", func(ctx *gin.Context){ workshopHandler.List(ctx)})
	workshopRouter.GET("/detail/:id", func(ctx *gin.Context){ workshopHandler.Detail(ctx)})

	reservationRouter := baseRouter.Group("/reservation")
	reservationRouter.POST("/create/:workshoID", func(ctx *gin.Context){ reservationHandler.Create(ctx)})
	
	return router
}