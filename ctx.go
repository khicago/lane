package lane

import (
	"context"
	"errors"
)

var (
	ErrLaneNotExist     = errors.New("target lane are not exist")
	ErrPayloadTypeError = errors.New("lane payload type error")
)

type contextKey struct {
	Name string
}

func ExtractLaneFromContext(name string, ctx context.Context) (Lane, error) {
	ctxKey := contextKey{Name: name}
	val := ctx.Value(ctxKey)
	if val == nil {
		return nil, ErrLaneNotExist
	}

	p, ok := val.(*PayloadTable)
	if !ok {
		return nil, ErrPayloadTypeError
	}
	l := New(name)
	l.(*lane).Payloads = p

	return l, nil
}
