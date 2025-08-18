package http

import (
	"bytes"
	"chat/shared/errs"
	"chat/shared/utils"
	"context"
	"encoding/json"
	"io"
	"net/http"
)

func DoRequest[Resp any](ctx context.Context, client *http.Client, method, url string, reqStruct any) (Resp, error) {
	eb := errs.B().Msg("failed to do http request")
	var respStruct Resp

	var body io.Reader
	if reqStruct != nil && method != http.MethodGet && method != http.MethodHead {
		reqBytes, err := json.Marshal(reqStruct)
		if err != nil {
			return respStruct, eb.Cause(err).Err()
		}
		body = bytes.NewReader(reqBytes)
	}

	req, err := http.NewRequestWithContext(ctx, method, url, body)
	if err != nil {
		return respStruct, eb.Cause(err).Err()
	}
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	resp, err := client.Do(req)
	if err != nil {
		return respStruct, eb.Cause(err).Err()
	}
	defer utils.CloseWithEB(resp.Body, eb)

	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return respStruct, eb.Cause(err).Err()
	}

	if resp.StatusCode >= 400 {
		var httpErr response
		_ = json.Unmarshal(respBytes, &httpErr)
		if httpErr.Message != "" {
			return respStruct, eb.Causef("%v", httpErr.Message).Err()
		}

		return respStruct, eb.Causef(
			"received error code: %v, body: %s",
			resp.StatusCode,
			string(respBytes),
		).Err()
	}

	if len(respBytes) > 0 {
		if err := json.Unmarshal(respBytes, &respStruct); err != nil {
			return respStruct, eb.Cause(err).Err()
		}
	}

	return respStruct, nil
}
