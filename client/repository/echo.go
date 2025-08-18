package repository

import (
	"chat/server"
	"chat/shared/endpoints"
	"chat/shared/errs"
	"chat/shared/http"
	"context"
)

func (r *Repository) Echo(ctx context.Context, payload string) (string, error) {
	eb := errs.B().Msg("failed to execute echo")

	req := server.EchoRequest{
		Value: payload,
	}

	resp, err := http.DoRequest[server.EchoResponse](ctx, r.httpClient, "POST", r.getServerURL(endpoints.Echo), req)
	if err != nil {
		return "", eb.Cause(err).Err()
	}

	return resp.Value, nil
}
