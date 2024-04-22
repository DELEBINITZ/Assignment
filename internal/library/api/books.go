package api

import (
	"encoding/json"
	"net/http"

	"github.com/akshay/libraryAssign/internal/library/exceptions"
	"github.com/akshay/libraryAssign/internal/library/middleware/auth"
	"github.com/akshay/libraryAssign/internal/library/models/dto"
	"github.com/akshay/libraryAssign/internal/library/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func (h *Handler) GetBooks(ctx *gin.Context) {
	resp, err := h.myService.GetBooks(ctx)
	if err != nil {
		utils.RespondError(ctx, err, utils.ConstructErrorAPIResp(err))
		return
	}

	utils.RespondJSON(ctx, http.StatusOK, utils.ConstructAPIResp("OK", resp))
}

func (h *Handler) AddBooks(ctx *gin.Context) {
	claims, err := auth.GetClaimsFromContext(ctx)
	if err != nil {
		err = exceptions.GenerateNewServerError(exceptions.NoClaimsFound, err, "no claims found", http.StatusUnauthorized)
		utils.RespondError(ctx, err, utils.ConstructErrorAPIResp(err))
		return
	}

	if claims.Role != dto.ADMIN {
		utils.RespondJSON(ctx, http.StatusForbidden, utils.ConstructAPIResp("insufficient permission", nil))
		return
	}

	r := ctx.Request.Body
	req := &dto.BooksList{}

	err = json.NewDecoder(r).Decode(req)
	if err != nil {
		err = exceptions.GenerateNewServerError(exceptions.InvalidRequestStruct, err, "invalid request", http.StatusBadRequest)
		utils.RespondError(ctx, err, utils.ConstructErrorAPIResp(err))
		return
	}

	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		err = exceptions.GenerateNewServerError(exceptions.InvalidRequestStruct, err, "invalid request", http.StatusBadRequest)
		utils.RespondError(ctx, err, utils.ConstructErrorAPIResp(err))
		return
	}

	err = h.myService.AddBooks(ctx, req)
	if err != nil {
		utils.RespondError(ctx, err, utils.ConstructErrorAPIResp(err))
		return
	}

	utils.RespondJSON(ctx, http.StatusOK, utils.ConstructAPIResp("OK", nil))
}

func (h *Handler) DeleteBook(ctx *gin.Context) {
	claims, err := auth.GetClaimsFromContext(ctx)
	if err != nil {
		err = exceptions.GenerateNewServerError(exceptions.NoClaimsFound, err, "no claims found", http.StatusUnauthorized)
		utils.RespondError(ctx, err, utils.ConstructErrorAPIResp(err))
		return
	}

	if claims.Role != dto.ADMIN {
		utils.RespondJSON(ctx, http.StatusForbidden, utils.ConstructAPIResp("insufficient permission", nil))
		return
	}

	r := ctx.Request.Body
	req := &dto.BookName{}

	err = json.NewDecoder(r).Decode(req)
	if err != nil {
		err = exceptions.GenerateNewServerError(exceptions.InvalidRequestStruct, err, "invalid request", http.StatusBadRequest)
		utils.RespondError(ctx, err, utils.ConstructErrorAPIResp(err))
		return
	}

	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		err = exceptions.GenerateNewServerError(exceptions.InvalidRequestStruct, err, "invalid request", http.StatusBadRequest)
		utils.RespondError(ctx, err, utils.ConstructErrorAPIResp(err))
		return
	}

	err = h.myService.DeleteBook(ctx, req)
	if err != nil {
		utils.RespondError(ctx, err, utils.ConstructErrorAPIResp(err))
		return
	}

	utils.RespondJSON(ctx, http.StatusOK, utils.ConstructAPIResp("OK", nil))
}
