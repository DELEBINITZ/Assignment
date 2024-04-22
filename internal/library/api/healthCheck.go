package api

import (
	"net/http"

	"github.com/akshay/libraryAssign/internal/library/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h *Handler) HealthCheck(ctx *gin.Context) {
	ctx.Set("key", uuid.NewString())
	err := h.myService.HealthCheck(ctx)
	if err != nil {

		utils.RespondError(ctx, err, nil)
		return
	}
	utils.RespondJSON(ctx, http.StatusOK, nil)
}
