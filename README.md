# easemob-go

# 版本说明
- 该项目 是对环信 IM REST API 的封装，这样做是为了节省服务器端开发者对接环信 API 的时间
- 该项目提供了用户、消息、群组、聊天室等资源的操作管理能力。

# Rest Api文档
- [官方文档](https://docs-im-beta.easemob.com/document/server-side/overview.html)

# 使用方法
- go.mod 文件引入： github.com/xiaofengin/easemob-go

有效的环信即时通讯 IM 开发者账号和 App Key、Client ID、ClientSecret、BasePath，登录 [环信管理后台](https://console.easemob.com/user/login) 到“应用列表” → 点击“查看”即可获取到 App Key、Client ID、ClientSecret，到"即时通讯" → 点击"服务概览"获取到 "Rest api" 的服务器域名。
#
```go

package main

import (
	"context"
	"fmt"
	IMSDK "github.com/xiaofengin/easemob-go"
)

func main() {
	client, err := IMSDK.New("appkey",
		"clientId",
		"clientSecret",
		"domainURL")
	if err != nil {
		return
	}
	var tos []string
	tos = append(tos, "环信用户ID")
	m := IMSDK.CreateTextMsg("hello word", tos)
	//m := CreateImageMsg("图片地址URL", "1.png", tos)
	//m := CreateAudioMsg("语音地址URL", "", tos, 3)
	//m := CreateVideoMsg("视频地址URL", "视频缩略图地址URL", "视频名.mp4", tos)
	//m := CreateFileMsg("文件地址URL", "文件名.pdf", tos)
	//m := CreateLocMsg("39.938881", "116.340836", "北京西直门外大街", tos)
	//m := CreateCmdMsg("cmd_action", tos)
	//m := CreateCustomMsg("custom_event", map[string]string{"ext_key1":"ext_value1"}, tos)
	m.Ext = map[string]interface{}{"key1": "value1", "key2": "value2"}
	ret, err := client.SendChatMessage(context.Background(), m)
	if err != nil {
		return
	}
	fmt.Printf("数据的值：%v\n", ret.Data)
}

```

# SDK功能清单
- user 用户信息模块  https://github.com/xiaofengin/easemob-go/blob/main/user_test.go
- push 推送信息模块  https://github.com/xiaofengin/easemob-go/blob/main/push_test.go
- message 消息模块  https://github.com/xiaofengin/easemob-go/blob/main/message_test.go
- contact 好友模块  https://github.com/xiaofengin/easemob-go/blob/main/contact_test.go
- chatroom 聊天室模块  https://github.com/xiaofengin/easemob-go/blob/main/chatroom_test.go
- chatgroup 群组模块  https://github.com/xiaofengin/easemob-go/blob/main/chatgroup_test.go