package ctxinjector

import "context"

type mapKey struct{}

func InjectMap(ctx context.Context, m map[any]any) context.Context {
	return context.WithValue(ctx, mapKey{}, m)
}

func GetMap(ctx context.Context) map[any]any {
	m, ok := ctx.Value(mapKey{}).(map[any]any)
	if !ok {
		return nil
	}

	return m
}
