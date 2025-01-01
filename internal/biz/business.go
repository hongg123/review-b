package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

//var (
//	// ErrUserNotFound is user not found.
//	ErrUserNotFound = errors.NotFound(v1.ErrorReason_USER_NOT_FOUND.String(), "user not found")
//)

// Greeter is a Greeter model.
type Business struct {
	Hello string
}

// GreeterRepo is a Greater repo.
type BusinessRepo interface {
	Reply(context.Context, *ReplyParam) (int64, error)
}

// GreeterUsecase is a Greeter usecase.
type BusinessUsecase struct {
	repo BusinessRepo
	log  *log.Helper
}

// NewGreeterUsecase new a Greeter usecase.
func NewBusinessUsecase(repo BusinessRepo, logger log.Logger) *BusinessUsecase {
	return &BusinessUsecase{repo: repo, log: log.NewHelper(logger)}

}

// CreateGreeter creates a Greeter, and returns the new Greeter.
func (uc *BusinessUsecase) CreateReply(ctx context.Context, param *ReplyParam) (int64, error) {
	uc.log.WithContext(ctx).Infof("CreateReply: %v", param)

	return uc.repo.Reply(ctx, param)
}
