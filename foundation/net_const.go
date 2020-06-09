package foundation

import (
	"context"
)

const (
	key_ctx_requestion = "key_ctx_requestion"
)


/**
 * 服务信息
 */
type Requisition struct {
	RequestId       uint64        `json:"requestId"`
	UserId          uint64        `json:"userId"`
	UserName        string        `json:"userName"`
	RoleId          uint64        `json:"roleId"`
}

type GinContext interface {
	Set(key string, value interface{})
	Get(key string) (value interface{}, exists bool)
}

/**
 * 请求信息创建，并绑定至context上
 */
func NewRequestion(ctx GinContext) *Requisition {
	rq := &Requisition{
		RequestId:       0,
		UserId:          0,
		UserName:        "",
	}
	ctx.Set(key_ctx_requestion, rq)
	return rq
}

/**
 * 创建包含请求信息实例的context，并绑定至context上
 */
func NewRequestionContext(ctx context.Context) context.Context {
	rq := &Requisition{
		RequestId:       0,
		UserId:          0,
		UserName:        "",
	}
	return context.WithValue(ctx, key_ctx_requestion, rq)
}

/**
 * 信息获取，获取传输链上context绑定的用户请求调用信息
 */
func GetRequisition(ctx context.Context) *Requisition {
	obj := ctx.Value(key_ctx_requestion)
	if info, ok := obj.(*Requisition); ok {
		return  info
	}
	return nil
}

/**
 * 信息校验，token绑定的用户信息同参数传入信息校验，信息不一致说明恶意用户传他人数据渗透
 */
func CheckRequisitionInfo(ctx context.Context, userId uint64, userName string) bool {
	info := GetRequisition(ctx)
	if info == nil {
		return false
	}
	if info.UserId != userId || info.UserName != userName {
		return false
	}
	return true
}

