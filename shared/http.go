package shared

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
)

type httpError struct {
	Message string `json:"message"`
}

func WriteHTTPError(w http.ResponseWriter, code int, err error) {
	body := httpError{
		Message: err.Error(),
	}

	bytes, _ := json.Marshal(body)

	w.WriteHeader(code)
	w.Write(bytes)
}

type validatable interface {
	Validate() error
}

func ParseHTTPRequest[T validatable](r *http.Request) (T, error) {
	eb := NewErrorBuilder().Msg("failed to parse http request")
	var val T

	bytes, err := io.ReadAll(r.Body)
	if err != nil {
		return val, eb.Cause(err).Err()
	}

	err = json.Unmarshal(bytes, &val)
	if err != nil {
		return val, eb.Cause(err).Err()
	}

	err = val.Validate()
	if err != nil {
		return val, eb.Cause(err).Err()
	}

	return val, nil
}

func DoHTTPRequest[Resp any](ctx context.Context, client *http.Client, method, url string, reqStruct any) (Resp, error) {
	eb := NewErrorBuilder().Msg("failed to do http request")
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
	defer CloseWithEB(resp.Body, eb)

	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return respStruct, eb.Cause(err).Err()
	}

	if resp.StatusCode >= 400 {
		var httpErr httpError
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
