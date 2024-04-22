package service

import (
	"github.com/akshay/libraryAssign/internal/library/middleware/auth"
	"github.com/akshay/libraryAssign/internal/library/models/dto"
	"github.com/akshay/libraryAssign/internal/library/utils"
	"github.com/gin-gonic/gin"
)

func (s *Service) Login(ctx *gin.Context, req *dto.LoginReq) (*dto.LoginRes, error) {
	role := dto.REGULAR
	if utils.IsAdmin(req.Email) {
		role = dto.ADMIN
	}

	accessToken, err := auth.GenerateToken(req.Email, role)
	if err != nil {
		return nil, err
	}

	res := &dto.LoginRes{
		Email:       req.Email,
		Status:      "success",
		AccessToken: accessToken,
	}
	return res, nil
}
