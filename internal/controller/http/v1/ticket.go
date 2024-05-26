package v1

import (
	"SkyTicket/internal/usecase/ticket"
	"github.com/gin-gonic/gin"
)

// TicketController handles HTTP requests for ticket operations
type TicketController struct {
	uc ticket.UseCase
}

// NewTicketController creates a new TicketController
func NewTicketController(uc ticket.UseCase) *TicketController {
	return &TicketController{uc: uc}
}

func (c *TicketController) BookTicket(ctx *gin.Context) {
}
