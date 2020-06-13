package requisition

import (
	"context"
	"github.com/sean-tech/gokit/foundation"
)

func NewError(err error, code int, ctx context.Context) *foundation.Error {

	var msg string
	req := GetRequisition(ctx)
	if req != nil {
		msg = Msg(req.Lang, code)
	} else {
		msg = "msg not found in request context"
	}
	return foundation.NewError(err, code, msg)
}