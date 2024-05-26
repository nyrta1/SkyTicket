package main

import (
	v1 "SkyTicket/internal/controller/http/v1"
	"SkyTicket/internal/domain/entity"
	"SkyTicket/internal/usecase/ticket"
	"SkyTicket/pkg/logger"
	"github.com/gin-gonic/gin"
)

func main() {
	log := logger.NewLogger()
	data := entity.NewTicketData()
	uc := ticket.NewTicketUseCase(data)
	ctrl := v1.NewTicketController(uc)

	r := gin.Default()
	r.POST("/v1/tickets", ctrl.BookTicket)

	if err := r.Run(); err != nil {
		log.Fatalf("could not run server: %v", err)
	}
}
