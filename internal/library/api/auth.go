package api

import (
	"encoding/json"
	"net/http"

	"github.com/akshay/libraryAssign/internal/library/exceptions"
	"github.com/akshay/libraryAssign/internal/library/models/dto"
	"github.com/akshay/libraryAssign/internal/library/utils"
	"github.com/gin-gonic/gin"
)

func (h *Handler) Login(ctx *gin.Context) {
	r := ctx.Request.Body
	req := &dto.LoginReq{}
	err := json.NewDecoder(r).Decode(req)
	if err != nil {
		err = exceptions.GenerateNewServerError(exceptions.InvalidRequestStruct, err, "invalid request", http.StatusBadRequest)
		utils.RespondError(ctx, err, utils.ConstructErrorAPIResp(err))
		return
	}
	resp, err := h.myService.Login(ctx, req)
	if err != nil {
		utils.RespondError(ctx, err, utils.ConstructErrorAPIResp(err))
		return
	}

	utils.RespondJSON(ctx, http.StatusOK, utils.ConstructAPIResp("OK", resp))
}
