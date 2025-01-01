package data

import (
	"context"
	v1 "review-b/api/review/v1"
	"review-b/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type businessRepo struct {
	data *Data
	log  *log.Helper
}

// NewGreeterRepo .
func NewBusinessRepo(data *Data, logger log.Logger) biz.BusinessRepo {
	return &businessRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *businessRepo) Reply(ctx context.Context, param *biz.ReplyParam) (int64, error) {
	r.log.WithContext(ctx).Infof("[data] Reply, param:%v", param)
	// 需要通过rpc调用其他服务获取数据
	ret, err := r.data.rc.ReplyReview(ctx, &v1.ReplyReviewRequest{
		ReviewID:  param.ReviewID,
		StoreID:   param.StoreID,
		Content:   param.Content,
		PicInfo:   param.PicInfo,
		VideoInfo: param.VideoInfo,
	})
	r.log.WithContext(ctx).Debugf("[data] ReplyReview return, ret:%v err:%v", ret, err)
	if err != nil {
		return 0, err
	}
	return ret.GetReplyID(), nil
}
