package traqbot

import "time"

// BasePayload ベースペイロード
type BasePayload struct {
	// EventTime イベント発生日時
	EventTime time.Time `json:"eventTime"`
}

// UserPayload ユーザー情報ペイロード
type UserPayload struct {
	// ID ユーザーUUID
	ID string `json:"id"`
	// Name ユーザーのtraQ ID
	Name string `json:"name"`
	// DisplayName ユーザーの表示名
	DisplayName string `json:"displayName"`
	// IconID ユーザーアイコンのファイルUUID
	IconID string `json:"iconId"`
	// Bot ユーザーがBotかどうか
	Bot bool `json:"bot"`
}

// ChannelPayload チャンネル情報ペイロード
type ChannelPayload struct {
	// ID チャンネルUUID
	ID string `json:"id"`
	// Name チャンネル名
	Name string `json:"name"`
	// Path チャンネルパス
	Path string `json:"path"`
	// ParentID 親チャンネルのUUID
	//
	// ルートチャンネルの場合は"00000000-0000-0000-0000-000000000000"
	ParentID string `json:"parentId"`
	// Creator チャンネル作成者
	Creator UserPayload `json:"creator"`
	// CreatedAt チャンネル作成日時
	CreatedAt time.Time `json:"createdAt"`
	// UpdatedAt チャンネル更新日時
	UpdatedAt time.Time `json:"updatedAt"`
}

// MessagePayload メッセージ情報ペイロード
type MessagePayload struct {
	// ID メッセージUUID
	ID string `json:"id"`
	// User メッセージ投稿者
	User UserPayload `json:"user"`
	// ChannelID 投稿先チャンネルUUID
	ChannelID string `json:"channelId"`
	// Text 生メッセージ本文
	Text string `json:"text"`
	// PlainText メッセージ本文(埋め込み情報・改行なし)
	PlainText string `json:"plainText"`
	// Embedded メッセージ埋め込み情報の配列
	Embedded []EmbeddedInfoPayload `json:"embedded"`
	// CreatedAt メッセージ投稿日時
	CreatedAt time.Time `json:"createdAt"`
	// UpdatedAt メッセージ更新日時
	UpdatedAt time.Time `json:"updatedAt"`
}

// EmbeddedInfoPayload メッセージ埋め込み情報
type EmbeddedInfoPayload struct {
	// Raw 表示文字列
	Raw string `json:"raw"`
	// Type タイプ
	Type string `json:"type"`
	// ID 各種ID(タイプによる)
	ID string `json:"id"`
}

// MessageStampPayload メッセージスタンプ情報
type MessageStampPayload struct {
	// StampID スタンプUUID
	StampID string `json:"stampId"`
	// UserID スタンプを押したユーザーUUID
	UserID string `json:"userId"`
	// Count このユーザーによって押されたこのスタンプの数
	Count int `json:"count"`
	// CreatedAt 最初にスタンプが押された日時
	CreatedAt time.Time `json:"createdAt"`
	// UpdatedAt 最後にスタンプが押された日時
	UpdatedAt time.Time `json:"updatedAt"`
}
