package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/cloudwego/eino-ext/components/model/ollama"
	"github.com/cloudwego/eino/schema"
	lark "github.com/larksuite/oapi-sdk-go/v3"
	larkcore "github.com/larksuite/oapi-sdk-go/v3/core"
	"github.com/larksuite/oapi-sdk-go/v3/event/dispatcher"
	larkim "github.com/larksuite/oapi-sdk-go/v3/service/im/v1"
	larkws "github.com/larksuite/oapi-sdk-go/v3/ws"
	"github.com/ollama/ollama/api"
)

var processedMessages = make(map[string]bool)

func main() {
	appID := os.Getenv("APP_ID")
	appSecret := os.Getenv("APP_SECRET")

	// 创建 LarkClient 对象，用于请求 OpenAPI
	client := lark.NewClient(appID, appSecret)

	// 初始化 Ollama 模型
	ctx := context.Background()
	model, err := ollama.NewChatModel(ctx, &ollama.ChatModelConfig{
		BaseURL: "http://localhost:11434", // Ollama 模型服务地址
		Timeout: 30 * time.Second,
		Model:   "llama2:7b", // 使用的模型名称
		Options: &api.Options{},
	})
	if err != nil {
		panic(err)
	}

	// 注册事件处理器
	eventHandler := dispatcher.NewEventDispatcher("", "").OnP2MessageReceiveV1(func(ctx context.Context, event *larkim.P2MessageReceiveV1) error {
		fmt.Printf("[OnP2MessageReceiveV1 access], data: %s\n", larkcore.Prettify(event))

		// 使用消息 ID 检查是否已处理过
		messageID := *event.Event.Message.MessageId
		if processedMessages[messageID] {
			return nil // 如果已处理过则直接返回
		}

		processedMessages[messageID] = true // 标记消息为已处理

		// 解析用户发送的消息
		content := *event.Event.Message.Content

		// 检查消息格式是否有效
		var respContent map[string]string
		err := json.Unmarshal([]byte(content), &respContent)
		if err != nil || *event.Event.Message.MessageType != "text" {
			respContent = map[string]string{
				"text": "解析消息失败，请发送文本消息\nparse message failed, please send text message",
			}
		} else {
			// 准备 Ollama 消息
			messages := []*schema.Message{
				schema.SystemMessage("直接给出问题的答案"),
				schema.UserMessage(respContent["text"]),
			}

			// 使用普通生成模式获取回复
			response, err := model.Generate(ctx, messages)
			if err != nil {
				fmt.Println("Error generating response from model:", err)
				respContent["text"] = "生成响应失败，请稍后再试。"
			} else {
				respContent["text"] = response.Content // 更新响应内容
			}
		}

		// 将内容序列化为 JSON 字符串，用于发送
		contentToSend := map[string]string{"text": respContent["text"]}
		contentJSON, err := json.Marshal(contentToSend) // 将内容序列化为 JSON 字符串
		if err != nil {
			fmt.Println("JSON marshal error:", err)
			return nil
		}

		// 发送消息
		if *event.Event.Message.ChatType == "p2p" {
			resp, err := client.Im.Message.Create(context.Background(), larkim.NewCreateMessageReqBuilder().
				ReceiveIdType(larkim.ReceiveIdTypeChatId).
				Body(larkim.NewCreateMessageReqBodyBuilder().
					MsgType(larkim.MsgTypeText).
					ReceiveId(*event.Event.Message.ChatId).
					Content(string(contentJSON)). // 确保是字符串格式的 JSON
					Build()).
				Build())

			if err != nil {
				fmt.Println("Error sending message:", err) // 打印具体错误
				return nil
			} else if !resp.Success() {
				fmt.Println("Failed to send message:", resp.Msg) // 打印信息
				return nil
			}
		} else {
			resp, err := client.Im.Message.Reply(context.Background(), larkim.NewReplyMessageReqBuilder().
				MessageId(*event.Event.Message.MessageId).
				Body(larkim.NewReplyMessageReqBodyBuilder().
					MsgType(larkim.MsgTypeText).
					Content(string(contentJSON)). // 确保是字符串格式的 JSON
					Build()).
				Build())

			if err != nil || !resp.Success() {
				fmt.Printf("logId: %s, error response: \n%s", resp.RequestId(), larkcore.Prettify(resp.CodeError))
				return nil
			}
		}

		return nil
	})

	// 启动长连接，并注册事件处理器
	cli := larkws.NewClient(appID, appSecret,
		larkws.WithEventHandler(eventHandler),
		larkws.WithLogLevel(larkcore.LogLevelDebug),
	)
	err = cli.Start(context.Background())
	if err != nil {
		panic(err)
	}
}
