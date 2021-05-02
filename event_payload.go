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

// MessageUpdatedPayload MESSAGE_UPDATEDイベントペイロード
type MessageUpdatedPayload struct {
	BasePayload
	// Message 更新されたメッセージ
	Message MessagePayload `json:"message"`
}

// MessageDeletedPayload MESSAGE_DELETEDイベントペイロード
type MessageDeletedPayload struct {
	BasePayload
	// Message 削除されたメッセージ
	Message struct {
		// ID メッセージUUID
		ID string `json:"id"`
		// ChannelID 投稿先チャンネルUUID
		ChannelID string `json:"channelId"`
	} `json:"message"`
}

// BotMessageStampsUpdatedPayload BOT_MESSAGE_STAMPS_UPDATEDイベントペイロード
type BotMessageStampsUpdatedPayload struct {
	BasePayload
	// MessageID スタンプの更新があったメッセージUUID
	MessageID string `json:"messageId"`
	// Stamps メッセージに現在ついている全てのスタンプ
	Stamps []MessageStampPayload `json:"stamps"`
}

// DirectMessageCreatedPayload DIRECT_MESSAGE_CREATEDイベントペイロード
type DirectMessageCreatedPayload struct {
	BasePayload
	// Message 投稿されたメッセージ
	Message MessagePayload `json:"message"`
}

// DirectMessageUpdatedPayload DIRECT_MESSAGE_UPDATEDイベントペイロード
type DirectMessageUpdatedPayload struct {
	BasePayload
	// Message 更新されたメッセージ
	Message MessagePayload `json:"message"`
}

// DirectMessageDeletedPayload DIRECT_MESSAGE_DELETEDイベントペイロード
type DirectMessageDeletedPayload struct {
	BasePayload
	// Message 削除されたメッセージ
	Message struct {
		// ID メッセージUUID
		ID string `json:"id"`
		// UserID DMの宛先ユーザーUUID
		UserID string `json:"userId"`
		// ChannelID 投稿先チャンネルUUID
		ChannelID string `json:"channelId"`
	} `json:"message"`
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
	// ID スタンプUUID
	ID string `json:"id"`
	// Name スタンプ名
	Name string `json:"name"`
	// FileID スタンプ画像ファイルUUID
	FileID string `json:"fileId"`
	// Creator スタンプを作成したユーザー
	Creator UserPayload `json:"creator"`
}

// TagAddedPayload TAG_ADDEDイベントペイロード
type TagAddedPayload struct {
	BasePayload
	// TagID タグUUID
	TagID string `json:"tagId"`
	// Tag タグ名
	Tag string `json:"tag"`
}

// TagRemovedPayload TAG_REMOVEDイベントペイロード
type TagRemovedPayload struct {
	BasePayload
	// TagID タグUUID
	TagID string `json:"tagId"`
	// Tag タグ名
	Tag string `json:"tag"`
}
