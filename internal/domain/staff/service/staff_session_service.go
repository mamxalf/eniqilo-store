package service

import (
	"context"
	"github.com/mamxalf/eniqilo-store/internal/domain/staff/request"
	"github.com/mamxalf/eniqilo-store/internal/domain/staff/response"
	"github.com/mamxalf/eniqilo-store/shared/failure"
	"github.com/mamxalf/eniqilo-store/shared/logger"
	"github.com/mamxalf/eniqilo-store/shared/token"
	"github.com/rs/zerolog/log"
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
		log.Err(err).Msg("[Login - Service] Generate Token Error")
		return
	}

	// TODO: save user sessions if needed

	res = response.SessionResponse{
		AccessToken: generatedToken.AccessToken,
		PhoneNumber: req.PhoneNumber,
		Name:        req.Name,
	}

	return
}
