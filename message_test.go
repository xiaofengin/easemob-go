package easemob_go

import (
	"context"
	"fmt"
	"testing"
)

func TestClient_DeleteChannel(t *testing.T) {

	client, err := New("appkey",
		"clientId",
		"clientSecret")
	if err != nil {
		return
	}
	p := &ChannelParam{
		Channel:    "会话ID",
		Type:       "chat",
		DeleteRoam: true,
	}
	ret, err := client.DeleteChannel(context.Background(), "用户ID", p)
	if err != nil {
		return
	}
	fmt.Printf("数据的值：%v\n", ret.Data)
}
func TestClient_GetHistoryAsUri(t *testing.T) {
	client, err := New("appkey",
		"clientId",
		"clientSecret")
	if err != nil {
		return
	}
	ret, err := client.GetHistoryAsUri(context.Background(), "2023111014")
	if err != nil {
		return
	}
	fmt.Printf("数据的值：%v\n", ret.Data)
}
func TestClient_ImportChatMessage(t *testing.T) {
	client, err := New("appkey",
		"clientId",
		"clientSecret")
	if err != nil {
		return
	}
	b := map[string]interface{}{
		"msg": "import message",
	}
	m := ImportMsgModel{
		Target:       "接收方ID",
		Type:         "txt",
		Body:         b,
		From:         "发送方id",
		IsAckRead:    false,
		MsgTimestamp: 0,
		NeedDownload: false,
	}
	ret, err := client.ImportChatMessage(context.Background(), &m)
	if err != nil {
		return
	}
	fmt.Printf("数据的值：%v\n", ret.Data)

}
func TestClient_ImportGroupMessage(t *testing.T) {
	client, err := New("appkey",
		"clientId",
		"clientSecret")
	if err != nil {
		return
	}
	b := map[string]interface{}{
		"msg": "import message",
	}
	m := ImportMsgModel{
		Target:       "群id",
		Type:         "txt",
		Body:         b,
		From:         "发送方id",
		IsAckRead:    false,
		MsgTimestamp: 0,
		NeedDownload: false,
	}
	ret, err := client.ImportGroupMessage(context.Background(), &m)
	if err != nil {
		return
	}
	fmt.Printf("数据的值：%v\n", ret.Data)
}
func TestClient_RecallMsg(t *testing.T) {
	client, err := New("appkey",
		"clientId",
		"clientSecret")
	if err != nil {
		return
	}
	p := MsgRecallParam{
		MsgId:    "消息ID",
		To:       "接收方ID",
		From:     "发送方ID",
		ChatType: "chat",
		Force:    true,
	}
	ret, err := client.RecallMsg(context.Background(), &p)
	if err != nil {
		return
	}
	fmt.Printf("数据的值：%v\n", ret.Data)
}
func TestClient_SendChatMessage(t *testing.T) {

	client, err := New("appkey",
		"clientId",
		"clientSecret")
	if err != nil {
		return
	}
	var tos []string
	tos = append(tos, "环信用户ID")
	m := CreateTextMsg("hello word", tos)
	//m := CreateImageMsg("图片地址URL", "1.png", tos)
	//m := CreateAudioMsg("语音地址URL", "", tos, 3)
	//m := CreateVideoMsg("视频地址URL", "视频缩略图地址URL", "视频名.mp4", tos)
	//m := CreateFileMsg("文件地址URL", "文件名.pdf", tos)
	//m := CreateLocMsg("39.938881", "116.340836", "北京西直门外大街", tos)
	//m := CreateCmdMsg("cmd_action", tos)
	//m := CreateCustomMsg("custom_event", map[string]string{"ext_key1": "ext_value1"}, tos)
	m.Ext = map[string]interface{}{"s": "s", "f": 6}
	ret, err := client.SendChatMessage(context.Background(), m)
	if err != nil {
		return
	}
	fmt.Printf("数据的值：%v\n", ret.Data)
}
func TestClient_SendGroupsMessage(t *testing.T) {
	client, err := New("appkey",
		"clientId",
		"clientSecret")
	if err != nil {
		return
	}
	var tos []string
	tos = append(tos, "群ID")
	m := CreateTextMsg("hello word", tos)
	//m := CreateImageMsg("图片地址URL", "1.png", tos)
	//m := CreateAudioMsg("语音地址URL", "", tos, 3)
	//m := CreateVideoMsg("视频地址URL", "视频缩略图地址URL", "视频名.mp4", tos)
	//m := CreateFileMsg("文件地址URL", "文件名.pdf", tos)
	//m := CreateLocMsg("39.938881", "116.340836", "北京西直门外大街", tos)
	//m := CreateCmdMsg("cmd_action", tos)
	//m := CreateCustomMsg("custom_event", map[string]string{"ext_key1":"ext_value1"}, tos)
	m.Ext = map[string]interface{}{"s": "s", "f": 6}
	ret, err := client.SendGroupsMessage(context.Background(), m)
	if err != nil {
		return
	}
	fmt.Printf("数据的值：%v\n", ret.Data)
}
func TestClient_SendRoomsMessage(t *testing.T) {
	client, err := New("appkey",
		"clientId",
		"clientSecret")
	if err != nil {
		return
	}
	var tos []string
	tos = append(tos, "聊天室ID")
	m := CreateTextMsg("hello word", tos)
	//m := CreateImageMsg("图片地址URL", "1.png", tos)
	//m := CreateAudioMsg("语音地址URL", "", tos, 3)
	//m := CreateVideoMsg("视频地址URL", "视频缩略图地址URL", "视频名.mp4", tos)
	//m := CreateFileMsg("文件地址URL", "文件名.pdf", tos)
	//m := CreateLocMsg("39.938881", "116.340836", "北京西直门外大街", tos)
	//m := CreateCmdMsg("cmd_action", tos)
	//m := CreateCustomMsg("custom_event", map[string]string{"ext_key1":"ext_value1"}, tos)
	m.Ext = map[string]interface{}{"s": "s", "f": 6}
	ret, err := client.SendRoomsMessage(context.Background(), m)
	if err != nil {
		return
	}
	fmt.Printf("数据的值：%v\n", ret.Data)
}
func TestClient_UploadingChatFile(t *testing.T) {
	client, err := New("appkey",
		"clientId",
		"clientSecret")
	if err != nil {
		return
	}
	ret, err := client.UploadingChatFile(context.Background(), "./README.md")
	if err != nil {
		return
	}
	fmt.Printf("数据的值：%v\n", ret.Entities[0].Uuid)
}
func TestClient_DownloadChatFile(t *testing.T) {
	client, err := New("appkey",
		"clientId",
		"clientSecret")
	if err != nil {
		return
	}
	err = client.DownloadChatFile(context.Background(), "4a07fb50-a55e-11ee-85b6-af93291f8570", "./examples/file.md")
	if err != nil {
		return
	}

}
