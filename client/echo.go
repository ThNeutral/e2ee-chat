package client

import (
	"bytes"
	"chat/shared"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type EchoRequest struct {
	Value string `json:"value"`
}

type EchoResponse struct {
	Value string `json:"value"`
}

func (c *Client) HandleEcho(w http.ResponseWriter, r *http.Request) {
	eb := shared.NewErrorBuilder().Msg("failed to handle echo")

	reqPayload := EchoRequest{
		Value: r.FormValue("value"),
	}

	bodyBytes, _ := json.Marshal(reqPayload)
	httpReq, err := http.NewRequest("POST", c.getServerURL("echo"), bytes.NewReader(bodyBytes))
	if err != nil {
		shared.WriteHTTPError(w, http.StatusInternalServerError, eb.Cause(err).Err())
		return
	}

	httpResp, err := c.httpClient.Do(httpReq)
	if err != nil {
		shared.WriteHTTPError(w, http.StatusInternalServerError, eb.Cause(err).Err())
		return
	}
	defer httpResp.Body.Close()

	respBytes, err := io.ReadAll(httpResp.Body)
	if err != nil {
		shared.WriteHTTPError(w, http.StatusInternalServerError, eb.Cause(err).Err())
		return
	}

	var echoResp EchoResponse
	if err := json.Unmarshal(respBytes, &echoResp); err != nil {
		shared.WriteHTTPError(w, http.StatusInternalServerError, eb.Cause(err).Err())
		return
	}

	log.Print(echoResp)
}
