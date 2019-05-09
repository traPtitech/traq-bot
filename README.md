# traQ BOTサーバーライブラリ
traQ BOTサーバーを簡単に作るためのライブラリです。

`traQ--->BOTサーバー`へのイベントの受け取り部分を補助します。`BOTサーバー--->traQ`へのリクエストを行うための、traQ APIクライアントは含まれていません。

## サンプル
```go
package main

import (
	"github.com/traPtitech/traq-bot"
	"log"
	"os"
)

func main() {
	vt := os.Getenv("VERIFICATION_TOKEN")

	handlers := traqbot.EventHandlers{}
	handlers.SetMessageCreatedHandler(func(payload *traqbot.MessageCreatedPayload) {
		log.Println("=================================================")
		log.Printf("%sさんがメッセージを投稿しました。\n", payload.Message.User.DisplayName)
		log.Println("メッセージ：")
		log.Println(payload.Message.Text)
		log.Println("=================================================")
	})

	server := traqbot.NewBotServer(vt, handlers)
	log.Fatal(server.ListenAndServe(":80"))
}

```
