package traqbot

// SetHandler イベントハンドラをセットします
//
// ハンドラにnilを指定すると、ハンドラは削除されます
func (hs EventHandlers) SetHandler(event string, h EventHandler) {
	if h == nil {
		delete(hs, event)
	} else {
		hs[event] = h
	}
}

// SetPingHandler イベントハンドラをセットします
func (hs EventHandlers) SetPingHandler(h func(payload *PingPayload)) {
	hs.SetHandler(Ping, func(event string, payload interface{}) { h(payload.(*PingPayload)) })
}

// SetJoinedHandler イベントハンドラをセットします
func (hs EventHandlers) SetJoinedHandler(h func(payload *JoinedPayload)) {
	hs.SetHandler(Joined, func(event string, payload interface{}) { h(payload.(*JoinedPayload)) })
}

// SetLeftHandler イベントハンドラをセットします
func (hs EventHandlers) SetLeftHandler(h func(payload *LeftPayload)) {
	hs.SetHandler(Left, func(event string, payload interface{}) { h(payload.(*LeftPayload)) })
}

// SetMessageCreatedHandler イベントハンドラをセットします
func (hs EventHandlers) SetMessageCreatedHandler(h func(payload *MessageCreatedPayload)) {
	hs.SetHandler(MessageCreated, func(event string, payload interface{}) { h(payload.(*MessageCreatedPayload)) })
}

// SetDirectMessageCreatedHandler イベントハンドラをセットします
func (hs EventHandlers) SetDirectMessageCreatedHandler(h func(payload *DirectMessageCreatedPayload)) {
	hs.SetHandler(DirectMessageCreated, func(event string, payload interface{}) { h(payload.(*DirectMessageCreatedPayload)) })
}

// SetChannelCreatedHandler イベントハンドラをセットします
func (hs EventHandlers) SetChannelCreatedHandler(h func(payload *ChannelCreatedPayload)) {
	hs.SetHandler(ChannelCreated, func(event string, payload interface{}) { h(payload.(*ChannelCreatedPayload)) })
}

// SetChannelTopicChangedHandler イベントハンドラをセットします
func (hs EventHandlers) SetChannelTopicChangedHandler(h func(payload *ChannelTopicChangedPayload)) {
	hs.SetHandler(ChannelTopicChanged, func(event string, payload interface{}) { h(payload.(*ChannelTopicChangedPayload)) })
}

// SetUserCreatedHandler イベントハンドラをセットします
func (hs EventHandlers) SetUserCreatedHandler(h func(payload *UserCreatedPayload)) {
	hs.SetHandler(UserCreated, func(event string, payload interface{}) { h(payload.(*UserCreatedPayload)) })
}

func (hs EventHandlers) SetStampCreatedHandler(h func(payload *StampCreatedPayload)) {
	hs.SetHandler(StampCreated, func(event string, payload interface{}) { h(payload.(*StampCreatedPayload)) })
}
