package easemob_go

import (
	"context"
	"net/http"
	"net/url"
	"path"
)

type RemarkResponse struct {
	Action    string `json:"action"`
	Duration  int    `json:"duration"`
	Status    string `json:"status"`
	Timestamp int64  `json:"timestamp"`
	Uri       string `json:"uri"`
}

// AddContact  添加好友
// ownerID 当前用户的用户 ID。
// friendID 要添加的用户 ID。
func (c *Client) AddContact(ctx context.Context, ownerID, friendID string) (*ResultResponse, error) {
	var resp ResultResponse
	p := path.Join("users", url.PathEscape(ownerID), "contacts/users", url.PathEscape(friendID))
	err := c.makeRequest(ctx, http.MethodPost, p, nil, nil, &resp)
	return &resp, err
}

// DeleteContact  移除好友
// ownerID 发起操作的用户 ID。
// friendID 被移除好友的用户 ID。
func (c *Client) DeleteContact(ctx context.Context, ownerID, friendID string) (*ResultResponse, error) {
	var resp ResultResponse
	p := path.Join("users", url.PathEscape(ownerID), "contacts/users", url.PathEscape(friendID))
	err := c.makeRequest(ctx, http.MethodDelete, p, nil, nil, &resp)
	return &resp, err
}

// SetContactRemark  设置好友备注
// ownerID 当前用户的用户 ID。
// friendID 要设置备注的用户 ID。
// remark 备注名
func (c *Client) SetContactRemark(ctx context.Context, ownerID, friendID, remark string) (*RemarkResponse, error) {
	var resp RemarkResponse
	p := path.Join("user", url.PathEscape(ownerID), "contacts/users", url.PathEscape(friendID))
	data := map[string]string{"remark": remark}
	err := c.makeRequest(ctx, http.MethodPut, p, nil, data, &resp)
	return &resp, err
}

// GetContactList 获取好友列表
func (c *Client) GetContactList(ctx context.Context, ownerID string) (*ResultResponse, error) {
	var resp ResultResponse
	p := path.Join("users", url.PathEscape(ownerID), "contacts/users")
	err := c.makeRequest(ctx, http.MethodGet, p, nil, nil, &resp)
	return &resp, err
}

// AddContactBlocks 添加用户至黑名单
// ownerID 当前用户的用户 ID。
// usernames 要加入黑名单的用户 ID，例如 ["user1", "user2"]。
func (c *Client) AddContactBlocks(ctx context.Context, ownerID string, usernames []string) (*ResultResponse, error) {
	var resp ResultResponse
	p := path.Join("users", url.PathEscape(ownerID), "blocks/users")
	data := map[string]interface{}{"usernames": usernames}
	err := c.makeRequest(ctx, http.MethodPost, p, nil, data, &resp)
	return &resp, err
}

// GetContactBlocksList 获取黑名单列表
// ownerID 当前用户的用户 ID。
// pageSize 每次期望返回的黑名单用户的数量。取值范围为 [1,50]。该参数仅在分页获取时为必需。
// cursor 数据查询的起始位置。该参数仅在分页获取时为必需。
func (c *Client) GetContactBlocksList(ctx context.Context, ownerID, pageSize, cursor string) (*ResultResponse, error) {
	var resp ResultResponse
	values := url.Values{}
	values.Add("pageSize", pageSize)
	values.Add("cursor", cursor)
	p := path.Join("users", url.PathEscape(ownerID), "blocks/users")
	err := c.makeRequest(ctx, http.MethodGet, p, values, nil, &resp)
	return &resp, err
}

// DeleteContactBlocks 从黑名单中移除用户
// ownerID 当前用户的用户 ID。
// blocked_username 要移出黑名单的用户 ID。
func (c *Client) DeleteContactBlocks(ctx context.Context, ownerID, blockedUsername string) (*ResultResponse, error) {
	var resp ResultResponse
	p := path.Join("users", url.PathEscape(ownerID), "blocks/users", url.PathEscape(blockedUsername))
	err := c.makeRequest(ctx, http.MethodDelete, p, nil, nil, &resp)
	return &resp, err
}
