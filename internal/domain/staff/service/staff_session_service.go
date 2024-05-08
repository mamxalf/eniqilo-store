package service

import (
	"context"
	"github.com/mamxalf/eniqilo-store/internal/domain/staff/request"
	"github.com/mamxalf/eniqilo-store/internal/domain/staff/response"
	"github.com/mamxalf/eniqilo-store/shared/failure"
	"github.com/mamxalf/eniqilo-store/shared/logger"
	"github.com/mamxalf/eniqilo-store/shared/token"
	"github.com/mamxalf/eniqilo-store/shared/utils"
	"net/http"
)

func (s *StaffServiceImpl) RegisterNewStaff(ctx context.Context, req request.RegisterRequest) (res response.SessionResponse, err error) {
	registerModel, err := req.ToModel()
	if err != nil {
		logger.ErrorInterfaceWithMessage(err, "failed parse request to model", "params", req)
		return
	}

	id, err := s.StaffRepository.Register(ctx, &registerModel)
	if err != nil {
		if failure.GetCode(err) == http.StatusConflict {
			logger.ErrorInterfaceWithMessage(err, "duplicate phone", "params", req)
			return
		}
		logger.ErrorInterfaceWithMessage(err, "failed register staff", "params", req)
		return
	}

	generateTokenParams := &token.GenerateTokenParams{
		AccessTokenSecret: s.Config.JwtSecret,
		AccessTokenExpiry: s.Config.JwtExpiry,
	}

	userData := &token.UserData{
		ID:          id.String(),
		Name:        req.Name,
		PhoneNumber: req.PhoneNumber,
	}

	generatedToken, err := token.GenerateToken(userData, generateTokenParams)
	if err != nil {
		logger.ErrorInterfaceWithMessage(err, "failed generate token", "generateTokenParams", generateTokenParams)
		return
	}

	// TODO: save user sessions if needed

	res = response.SessionResponse{
		UserID:      userData.ID,
		AccessToken: generatedToken.AccessToken,
		PhoneNumber: userData.PhoneNumber,
		Name:        userData.Name,
	}

	return
}

func (s *StaffServiceImpl) LoginStaff(ctx context.Context, req request.LoginRequest) (res response.SessionResponse, err error) {
	staff, err := s.StaffRepository.FindByPhone(ctx, req.PhoneNumber)
	if err != nil {
		logger.ErrorWithMessage(err, "Failed Get Staff Data ByPhone")
		return
	}

	err = utils.CheckPasswordHash(req.Password, staff.Password)
	if err != nil {
		err = failure.BadRequest(err)
		logger.ErrorWithMessage(err, "Invalid Password!")
		return
	}

	generateTokenParams := &token.GenerateTokenParams{
		AccessTokenSecret: s.Config.JwtSecret,
		AccessTokenExpiry: s.Config.JwtExpiry,
	}

	staffData := &token.UserData{
		ID:          staff.ID.String(),
		Name:        staff.Name,
		PhoneNumber: staff.Phone,
	}
	generatedToken, err := token.GenerateToken(staffData, generateTokenParams)
	if err != nil {
		logger.ErrorInterfaceWithMessage(err, "failed generate token", "generateTokenParams", generateTokenParams)
		return
	}

	res = response.SessionResponse{
		UserID:      staff.ID.String(),
		AccessToken: generatedToken.AccessToken,
		PhoneNumber: staff.Phone,
		Name:        staff.Name,
	}

	return
}
