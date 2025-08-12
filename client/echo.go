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
	reqPayload := EchoRequest{
		Value: r.FormValue("value"),
	}

	bodyBytes, _ := json.Marshal(reqPayload)
	httpReq, err := http.NewRequest("POST", c.getServerURL("echo"), bytes.NewReader(bodyBytes))
	if err != nil {
		shared.WriteHTTPError(w, http.StatusInternalServerError, err)
		return
	}

	httpResp, err := c.httpClient.Do(httpReq)
	if err != nil {
		shared.WriteHTTPError(w, http.StatusInternalServerError, err)
		return
	}
	defer httpResp.Body.Close()

	respBytes, err := io.ReadAll(httpResp.Body)
	if err != nil {
		shared.WriteHTTPError(w, http.StatusInternalServerError, err)
		return
	}

	var echoResp EchoResponse
	if err := json.Unmarshal(respBytes, &echoResp); err != nil {
		shared.WriteHTTPError(w, http.StatusInternalServerError, err)
		return
	}

	log.Print(echoResp)
}
