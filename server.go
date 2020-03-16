package traqbot

import (
	"encoding/json"
	"net/http"
	"strings"
)

const (
	contentTypeHeader   = "Content-Type"
	mimeApplicationJson = "application/json"
	botTokenHeader      = "X-TRAQ-BOT-TOKEN"
	botEventHeader      = "X-TRAQ-BOT-EVENT"
)

// EventHandler イベントハンドラ
type EventHandler func(event string, payload interface{})

// EventHandlers イベントハンドラマップ
type EventHandlers map[string]EventHandler

// BotServer BOTサーバー
type BotServer struct {
	verificationToken string
	handlers          EventHandlers
}

// NewBotServer 新しくBOTサーバーを生成します
func NewBotServer(verification string, handlers EventHandlers) *BotServer {
	return &BotServer{
		verificationToken: verification,
		handlers:          handlers,
	}
}

// ListenAndServe BOTサーバーを起動します
func (bs *BotServer) ListenAndServe(addr string) error {
	mux := http.NewServeMux()
	mux.Handle("/", bs)
	server := &http.Server{Addr: addr, Handler: mux}
	return server.ListenAndServe()
}

// ServeHTTP BOTサーバーハンドラ
func (bs *BotServer) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	// VerificationTokenチェック
	if req.Header.Get(botTokenHeader) != bs.verificationToken {
		rw.WriteHeader(http.StatusForbidden)
		return
	}

	// Eventヘッダチェック
	event := req.Header.Get(botEventHeader)
	if len(event) == 0 {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	// RequestがJSONかどうか
	if !strings.HasPrefix(req.Header.Get(contentTypeHeader), mimeApplicationJson) {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	var payload interface{}
	switch event {
	case Ping:
		payload = &PingPayload{}
	case Joined:
		payload = &JoinedPayload{}
	case Left:
		payload = &LeftPayload{}
	case MessageCreated:
		payload = &MessageCreatedPayload{}
	case DirectMessageCreated:
		payload = &DirectMessageCreatedPayload{}
	case ChannelCreated:
		payload = &ChannelCreatedPayload{}
	case ChannelTopicChanged:
		payload = &ChannelTopicChangedPayload{}
	case UserCreated:
		payload = &UserCreatedPayload{}
	case StampCreated:
		payload = &StampCreatedPayload{}
	default:
		rw.WriteHeader(http.StatusNotImplemented)
		return
	}

	if err := json.NewDecoder(req.Body).Decode(payload); err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	rw.WriteHeader(http.StatusNoContent)

	h, ok := bs.handlers[event]
	if ok && h != nil {
		h(event, payload)
	}
}
