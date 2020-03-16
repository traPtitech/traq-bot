package traqbot

// PingPayload PINGイベントペイロード
type PingPayload struct {
	BasePayload
}

// JoinedPayload JOINEDイベントペイロード
type JoinedPayload struct {
	BasePayload
	// Channel 参加したチャンネル
	Channel ChannelPayload `json:"channel"`
}

// LeftPayload LEFTイベントペイロード
type LeftPayload struct {
	BasePayload
	// Channel 退出したチャンネル
	Channel ChannelPayload `json:"channel"`
}

// MessageCreatedPayload MESSAGE_CREATEDイベントペイロード
type MessageCreatedPayload struct {
	BasePayload
	// Message 投稿されたメッセージ
	Message MessagePayload `json:"message"`
}

// DirectMessageCreatedPayload DIRECT_MESSAGE_CREATEDイベントペイロード
type DirectMessageCreatedPayload struct {
	BasePayload
	// Message 投稿されたメッセージ
	Message MessagePayload `json:"message"`
}

// ChannelCreatedPayload CHANNEL_CREATEDイベントペイロード
type ChannelCreatedPayload struct {
	BasePayload
	// Channel 作成されたチャンネル
	Channel ChannelPayload `json:"channel"`
}

// ChannelTopicChangedPayload CHANNEL_TOPIC_CHANGEDイベントペイロード
type ChannelTopicChangedPayload struct {
	BasePayload
	// Channel 変更されたチャンネル
	Channel ChannelPayload `json:"channel"`
	// Topic 変更後のトピック
	Topic string `json:"topic"`
	// Updater トピック更新者
	Updater UserPayload `json:"updater"`
}

// UserCreatedPayload USER_CREATEDイベントペイロード
type UserCreatedPayload struct {
	BasePayload
	// User 作成されたユーザー
	User UserPayload `json:"user"`
}

// StampCreatedPayload STAMP_CREATEDイベントペイロード
type StampCreatedPayload struct {
	BasePayload
	ID      string      `json:"id"`
	Name    string      `json:"name"`
	FileID  string      `json:"fileId"`
	Creator UserPayload `json:"creator"`
}
