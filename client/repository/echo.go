package repository

import (
	"chat/server"
	"chat/shared"
	"chat/shared/endpoints"
	"context"
)

func (r *Repository) Echo(ctx context.Context, payload string) (string, error) {
	eb := shared.NewErrorBuilder().Msg("failed to execute echo")

	req := server.EchoRequest{
		Value: payload,
	}

	resp, err := shared.DoHTTPRequest[server.EchoResponse](ctx, r.httpClient, "POST", r.getServerURL(endpoints.Echo), req)
	if err != nil {
		return "", eb.Cause(err).Err()
	}

	return resp.Value, nil
}
