package easemob_go

import (
	"bytes"
	"context"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path"
)

type entities struct {
	Uuid        string `json:"uuid"`
	Type        string `json:"type"`
	ShareSecret string `json:"share-secret"`
}
type UploadingResponse struct {
	Entities []entities `json:"entities"`
	Response
}
type ChannelParam struct {
	Channel    string `json:"channel"`     //要删除的会话 ID
	Type       string `json:"type"`        //会话类型。chat：单聊会话；groupchat：群聊会话。
	DeleteRoam bool   `json:"delete_roam"` //是否删除该会话在服务端的漫游消息。
}
type MessageDownloadParam struct {
	Dir  string `json:"dir"`
	Time string `json:"time"`
}
type MsgRecallParam struct {
	MsgId    string `json:"msg_id"`
	To       string `json:"to"`
	From     string `json:"from"`
	ChatType string `json:"chat_type"`
	Force    bool   `json:"force"`
}

// DeleteChannel 单向删除会话
func (c *Client) DeleteChannel(ctx context.Context, userID string, param *ChannelParam) (*ResultResponse, error) {

	var resp ResultResponse
	p := path.Join("users", url.PathEscape(userID), "user_channel")
	err := c.makeRequest(ctx, http.MethodDelete, p, nil, param, &resp)
	return &resp, err
}

// GetHistoryAsUri 获取历史消息记录
func (c *Client) GetHistoryAsUri(ctx context.Context, time string) (*ResultResponse, error) {
	var resp ResultResponse
	p := path.Join("chatmessages", url.PathEscape(time))
	err := c.makeRequest(ctx, http.MethodGet, p, nil, nil, &resp)
	return &resp, err
}

// ImportChatMessage 导入单聊消息
func (c *Client) ImportChatMessage(ctx context.Context, msg *ImportMsgModel) (*ResultResponse, error) {
	var resp ResultResponse
	err := c.makeRequest(ctx, http.MethodPost, "messages/users/import", nil, msg, &resp)
	return &resp, err
}

// ImportGroupMessage 导入群聊消息
func (c *Client) ImportGroupMessage(ctx context.Context, msg *ImportMsgModel) (*ResultResponse, error) {
	var resp ResultResponse
	err := c.makeRequest(ctx, http.MethodPost, "messages/chatgroups/import", nil, msg, &resp)
	return &resp, err
}

// RecallMsg 撤回消息
func (c *Client) RecallMsg(ctx context.Context, param *MsgRecallParam) (*ResultResponse, error) {
	var resp ResultResponse
	err := c.makeRequest(ctx, http.MethodPost, "messages/msg_recall", nil, param, &resp)
	return &resp, err
}

// SendChatMessage 发送单聊消息
func (c *Client) SendChatMessage(ctx context.Context, msg *MsgModel) (*ResultResponse, error) {
	var resp ResultResponse

	err := c.makeRequest(ctx, http.MethodPost, "messages/users?useMsgId=true", nil, msg, &resp)
	return &resp, err
}

// SendGroupsMessage 发送群聊消息
func (c *Client) SendGroupsMessage(ctx context.Context, msg *MsgModel) (*ResultResponse, error) {
	var resp ResultResponse

	err := c.makeRequest(ctx, http.MethodPost, "messages/chatgroups?useMsgId=true", nil, msg, &resp)
	return &resp, err
}

// SendRoomsMessage 发送聊天室消息
func (c *Client) SendRoomsMessage(ctx context.Context, msg *MsgModel) (*ResultResponse, error) {
	var resp ResultResponse

	err := c.makeRequest(ctx, http.MethodPost, "messages/chatrooms?useMsgId=true", nil, msg, &resp)
	return &resp, err
}

// UploadingChatFile 上传文件到环信
// restrictAccess:是否限制访问该文件：
// - true：是。用户需要通过响应 body 中获取的文件访问密钥（share-secret）才能下载该文件。
// - false：否。表示不限制访问。用户可以直接下载该文件。
func (c *Client) UploadingChatFile(ctx context.Context, filePath string) (*UploadingResponse, error) {
	var resp UploadingResponse
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", file.Name())
	_, err = io.Copy(part, file)
	writer.Close()
	err = c.uploadingFile(ctx, http.MethodPost, "chatfiles", writer.FormDataContentType(), body, &resp)
	return &resp, err
}

// DownloadChatFile 下载文件
// shareSecret:文件访问密钥。若上传文件时限制了访问，下载该文件时则需要该访问密钥。成功上传文件后，从 文件上传 的响应 body 中获取该密钥。
// fileUuid 服务器为文件生成的 UUID。
func (c *Client) DownloadChatFile(ctx context.Context, fileUuid, filePath string) error {
	p := path.Join("chatfiles", url.PathEscape(fileUuid))
	err := c.downloadFile(ctx, http.MethodGet, p, filePath)
	return err
}
