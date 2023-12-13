package rumblelivestreamlib

import (
	"bufio"
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/tylertravisty/go-utils/random"
)

type ChatInfo struct {
	UrlPrefix string
	ChatID    string
	ChannelID int
}

func (ci *ChatInfo) Url() string {
	return fmt.Sprintf("%s/chat/%s/message", ci.UrlPrefix, ci.ChatID)
}

func (c *Client) streamChatInfo() (*ChatInfo, error) {
	if c.StreamUrl == "" {
		return nil, fmt.Errorf("stream url is empty")
	}

	resp, err := c.getWebpage(c.StreamUrl)
	if err != nil {
		return nil, fmt.Errorf("error getting stream webpage: %v", err)
	}
	defer resp.Body.Close()

	r := bufio.NewReader(resp.Body)
	line, _, err := r.ReadLine()
	var lineS string
	for err == nil {
		lineS = string(line)
		if strings.Contains(lineS, "RumbleChat(") {
			start := strings.Index(lineS, "RumbleChat(") + len("RumbleChat(")
			end := strings.Index(lineS[start:], ");")
			argsS := strings.ReplaceAll(lineS[start:start+end], ", ", ",")
			argsS = strings.Replace(argsS, "[", "\"[", 1)
			n := strings.LastIndex(argsS, "]")
			argsS = argsS[:n] + "]\"" + argsS[n+1:]
			c := csv.NewReader(strings.NewReader(argsS))
			args, err := c.ReadAll()
			if err != nil {
				return nil, fmt.Errorf("error parsing csv: %v", err)
			}
			info := args[0]
			channelID, err := strconv.Atoi(info[5])
			if err != nil {
				return nil, fmt.Errorf("error converting channel ID argument string to int: %v", err)
			}
			return &ChatInfo{info[0], info[1], channelID}, nil
		}
		line, _, err = r.ReadLine()
	}
	if err != nil {
		return nil, fmt.Errorf("error reading line from stream webpage: %v", err)
	}

	return nil, fmt.Errorf("did not find RumbleChat function call")
}

type ChatMessage struct {
	Text string `json:"text"`
}

type ChatData struct {
	RequestID string      `json:"request_id"`
	Message   ChatMessage `json:"message"`
	Rant      *string     `json:"rant"`
	ChannelID *int        `json:"channel_id"`
}

type ChatRequest struct {
	Data ChatData `json:"data"`
}

func (c *Client) Chat(message string) error {
	if c.httpClient == nil {
		return pkgErr("", fmt.Errorf("http client is nil"))
	}

	chatInfo, err := c.streamChatInfo()
	if err != nil {
		return pkgErr("error getting stream chat info", err)
	}

	requestID, err := random.String(32)
	if err != nil {
		return pkgErr("error generating request ID", err)
	}
	body := ChatRequest{
		Data: ChatData{
			RequestID: requestID,
			Message: ChatMessage{
				Text: message,
			},
			Rant:      nil,
			ChannelID: &chatInfo.ChannelID,
		},
	}
	bodyB, err := json.Marshal(body)
	if err != nil {
		return pkgErr("error marshaling request body into json", err)
	}

	resp, err := c.httpClient.Post(chatInfo.Url(), "application/json", bytes.NewReader(bodyB))
	if err != nil {
		return pkgErr("http Post request returned error", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("http Post response status not %s: %s", http.StatusText(http.StatusOK), resp.Status)
	}

	return nil
}
