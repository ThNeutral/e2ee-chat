package ctxinjector

import "context"

type errorKey struct{}

func InjectError(ctx context.Context, err error) context.Context {
	return context.WithValue(ctx, errorKey{}, err)
}

func GetError(ctx context.Context) error {
	err, ok := ctx.Value(errorKey{}).(error)
	if !ok {
		return nil
	}

	return err
}
