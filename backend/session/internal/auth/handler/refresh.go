package handler

import (
	"context"

	"github.com/cloudnativedaysjp/cnd-handson-app/backend/session/internal/auth/service"
	sessionpb "github.com/cloudnativedaysjp/cnd-handson-app/backend/session/proto"
	"github.com/google/uuid"
)

// RefreshTokenServiceHandler リフレッシュトークンサービスのハンドラー
type RefreshTokenServiceHandler struct {
	// Dependency injection for storage or other services
	sessionpb.UnimplementedRefreshTokenServiceServer
}

// GenerateRefreshToken リフレッシュトークンの生成
func (h *RefreshTokenServiceHandler) GenerateRefreshToken(ctx context.Context, req *sessionpb.GenerateRefreshTokenRequest) (*sessionpb.RefreshTokenResponse, error) {
	refreshToken, expiresAt, err := service.GenerateRefreshToken(uuid.MustParse(req.UserId))
	if err != nil {
		return nil, err
	}
	return &sessionpb.RefreshTokenResponse{
		RefreshToken: refreshToken,
		ExpiresAt:    expiresAt,
	}, nil
}

// ValidateRefreshToken リフレッシュトークンの検証
func (h *RefreshTokenServiceHandler) ValidateRefreshToken(ctx context.Context, req *sessionpb.ValidateRefreshTokenRequest) (*sessionpb.ValidateRefreshTokenResponse, error) {
	// リフレッシュトークンを検証
	err := service.ValidateRefreshToken(uuid.MustParse(req.UserId), req.RefreshToken)
	if err != nil {
		return &sessionpb.ValidateRefreshTokenResponse{
			Valid: false,
		}, err
	}

	return &sessionpb.ValidateRefreshTokenResponse{
		Valid: true,
	}, nil
}

// RevokeRefreshToken リフレッシュトークンの無効化
func (h *RefreshTokenServiceHandler) RevokeRefreshToken(ctx context.Context, req *sessionpb.RevokeRefreshTokenRequest) (*sessionpb.RevokeRefreshTokenResponse, error) {
	err := service.RevokeRefreshToken(uuid.MustParse(req.UserId), req.RefreshToken)
	if err != nil {
		return &sessionpb.RevokeRefreshTokenResponse{
			Success: false,
		}, err
	}
	return &sessionpb.RevokeRefreshTokenResponse{
		Success: true,
	}, nil
}
