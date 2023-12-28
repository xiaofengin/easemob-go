package easemob_go

import (
	"context"
	"errors"
	"net/http"
	"net/url"
	"path"
)

type ActionResponse struct {
	Action    string `json:"action"`
	Timestamp int64  `json:"timestamp"`
	Duration  int    `json:"duration"`
}
type UserRegisterParam struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// MutesParam
// > 0：该用户 ID 的单聊禁言时长。
// = 0：取消该用户的单聊禁言。
// = -1：该用户被设置永久单聊禁言。
type MutesParam struct {
	Username  string `json:"username"`
	Chat      int    `json:"chat"`
	GroupChat int    `json:"groupchat"`
	Chatroom  int    `json:"chatroom"`
}

type TokenParam struct {
	// 授权方式。
	//- 若值为 password，通过用户 ID 和密码获取 token，需设置 username 和 password 参数。
	//- 若值为 inherit，通过用户 ID 获取 token，只需设置 username 参数。在该请求中，该参数需设置为 inherit
	GrantType string `json:"grant_type"`
	//用户 ID。
	Username string `json:"username"`
	//用户的登录密码。
	Password string `json:"password,omitempty"`
	//当用户不存在时，是否自动创建用户。自动创建用户时，需保证授权方式（grant_type）必须为 inherit，API 请求 header 中使用 App token 进行鉴权。
	AutoCreateUser bool `json:"autoCreateUser,omitempty"`
	//否	token 有效期，单位为秒。
	//- 若传入该参数，token 有效期以传入的值为准。
	//- 若不传该参数，以 环信即时通讯云控制台的用户认证页面的 token 有效期的设置为准。
	//- 若设置为 0，则 token 永久有效
	Ttl string `json:"ttl,omitempty"`
}
type TokenResponse struct {
	AccessToken string      `json:"access_token"`
	ExpiresIn   int         `json:"expires_in"`
	User        interface{} `json:"user"`
}

// GetUserToken 获取用户token
func (c *Client) GetUserToken(ctx context.Context, param *TokenParam) (*TokenResponse, error) {
	var resp TokenResponse
	err := c.makeRequest(ctx, http.MethodPost, "token", nil, param, &resp)
	return &resp, err
}

// UserRegister  批量注册用户  单次请求最多可注册 60 个用户 ID。
// users 用户数据
func (c *Client) UserRegister(ctx context.Context, users *[]UserRegisterParam) (*ResultResponse, error) {
	var resp ResultResponse
	err := c.makeRequest(ctx, http.MethodPost, "users", nil, users, &resp)
	return &resp, err
}

// GetUserDetail  获取用户详情
// userID 用户id
func (c *Client) GetUserDetail(ctx context.Context, userID string) (*ResultResponse, error) {
	var resp ResultResponse
	p := path.Join("users", url.PathEscape(userID))
	err := c.makeRequest(ctx, http.MethodGet, p, nil, nil, &resp)
	return &resp, err
}

// GetUserDetailList  批量获取用户详情
// limit 请求查询用户的数量。取值范围为 [1,100]，默认值为 10。若实际用户数量超过 100，返回 100 个用户。
// cursor 开始获取数据的游标位置，用于分页显示用户列表。第一次发起批量查询用户请求时若不设置 cursor，请求成功后会获得第一页用户列表。从响应 body 中获取 cursor，并在下一次请求的 URL 中传入该 cursor，直到响应 body 中不再有 cursor 字段，则表示已查询到 app 中所有用户。
func (c *Client) GetUserDetailList(ctx context.Context, limit, cursor string) (*ResultResponse, error) {
	var resp ResultResponse
	values := url.Values{}
	values.Add("limit", limit)
	values.Add("cursor", cursor)
	err := c.makeRequest(ctx, http.MethodGet, "users", values, nil, &resp)
	return &resp, err
}

// DeleteUser  删除单个用户
// userID 用户id
func (c *Client) DeleteUser(ctx context.Context, userID string) (*ResultResponse, error) {
	var resp ResultResponse
	p := path.Join("users", url.PathEscape(userID))
	err := c.makeRequest(ctx, http.MethodDelete, p, nil, nil, &resp)
	return &resp, err
}

// UpdateUserPassword  修改用户密码
// userID 用户id
// password 新密码
func (c *Client) UpdateUserPassword(ctx context.Context, userID, password string) (*ActionResponse, error) {
	var resp ActionResponse
	p := path.Join("users", url.PathEscape(userID), "password")
	data := map[string]string{"newpassword": password}
	err := c.makeRequest(ctx, http.MethodPut, p, nil, data, &resp)
	return &resp, err
}

// GetUserStatusList 批量获取用户在线状态
// users 用户数据
func (c *Client) GetUserStatusList(ctx context.Context, users []string) (*ResultResponse, error) {
	var resp ResultResponse
	data := map[string]interface{}{"usernames": users}
	err := c.makeRequest(ctx, http.MethodPost, "users/batch/status", nil, data, &resp)
	return &resp, err
}

// GetUserResources 获取指定账号的在线登录设备列表
// userID 用户id
func (c *Client) GetUserResources(ctx context.Context, userID string) (*ResultResponse, error) {
	var resp ResultResponse
	p := path.Join("users", url.PathEscape(userID), "resources")
	err := c.makeRequest(ctx, http.MethodGet, p, nil, nil, &resp)
	return &resp, err
}

// SetUserMutes 设置用户全局禁言
// param 禁言数据
func (c *Client) SetUserMutes(ctx context.Context, param *MutesParam) (*ResultResponse, error) {
	var resp ResultResponse
	err := c.makeRequest(ctx, http.MethodPost, "mutes", nil, param, &resp)
	return &resp, err
}

// GetUserMutes 查询单个用户 ID 全局禁言
// userID 用户id
func (c *Client) GetUserMutes(ctx context.Context, userID string) (*ResultResponse, error) {
	var resp ResultResponse
	p := path.Join("mutes", url.PathEscape(userID))
	err := c.makeRequest(ctx, http.MethodGet, p, nil, nil, &resp)
	return &resp, err
}

// GetAppMutesList 查询 app 下的所有全局禁言的用户
func (c *Client) GetAppMutesList(ctx context.Context) (*ResultResponse, error) {
	var resp ResultResponse
	err := c.makeRequest(ctx, http.MethodGet, "mutes", nil, nil, &resp)
	return &resp, err
}

// CountMissedMessages  获取用户离线消息数量
func (c *Client) CountMissedMessages(ctx context.Context, userID string) (*ResultResponse, error) {
	if len(userID) == 0 {
		return nil, errors.New("userID is nil")
	}
	p := path.Join("users", url.PathEscape(userID), "offline_msg_count")

	var resp ResultResponse
	err := c.makeRequest(ctx, http.MethodGet, p, nil, nil, &resp)
	return &resp, err
}

// IsMessageDeliveredToUser 查询离线消息的状态，如是否已下发
func (c *Client) IsMessageDeliveredToUser(ctx context.Context, toUser, messageId string) (*ResultResponse, error) {
	var resp ResultResponse
	p := path.Join("users", url.PathEscape(toUser), "offline_msg_status", url.PathEscape(messageId))
	err := c.makeRequest(ctx, http.MethodGet, p, nil, nil, &resp)
	return &resp, err
}

// AccountBan  账号封禁
// userID 用户id
func (c *Client) AccountBan(ctx context.Context, userID string) (*ResultResponse, error) {
	var resp ResultResponse
	p := path.Join("users", url.PathEscape(userID), "deactivate")
	err := c.makeRequest(ctx, http.MethodPost, p, nil, nil, &resp)
	return &resp, err
}

// AccountUnban  账号解禁
// userID 用户id
func (c *Client) AccountUnban(ctx context.Context, userID string) (*ActionResponse, error) {
	var resp ActionResponse
	p := path.Join("users", url.PathEscape(userID), "activate")
	err := c.makeRequest(ctx, http.MethodPost, p, nil, nil, &resp)
	return &resp, err
}

// AccountDisconnect  强制下线
// userID 用户id
func (c *Client) AccountDisconnect(ctx context.Context, userID string) (*ResultResponse, error) {
	var resp ResultResponse
	p := path.Join("users", url.PathEscape(userID), "disconnect")
	err := c.makeRequest(ctx, http.MethodGet, p, nil, nil, &resp)
	return &resp, err
}

// SetUserMetadata  设置用户属性
// metadata 用户属性 包含 nickname avatarurl phone mail gender sign  birth ext
func (c *Client) SetUserMetadata(ctx context.Context, userID string, metadata map[string]string) (*ResultResponse, error) {
	var resp ResultResponse
	values := url.Values{}
	for key, value := range metadata {
		values.Add(key, value)
	}
	p := path.Join("metadata/user", url.PathEscape(userID))
	err := c.makeMetadataRequest(ctx, http.MethodPut, p, values, "application/x-www-form-urlencoded", &resp)
	return &resp, err
}

// GetUserMetadata  获取用户属性
func (c *Client) GetUserMetadata(ctx context.Context, userID string) (*ResultResponse, error) {
	var resp ResultResponse
	p := path.Join("metadata/user", url.PathEscape(userID))
	err := c.makeRequest(ctx, http.MethodGet, p, nil, nil, &resp)
	return &resp, err
}

// GetUserMetadataList  批量获取用户属性
// targets 用户 ID 列表，最多可传 100 个用户 ID。
// properties 属性名列表，查询结果只返回该列表中包含的属性，不在该列表中的属性将被忽略。
func (c *Client) GetUserMetadataList(ctx context.Context, targets, properties []string) (*ResultResponse, error) {
	var resp ResultResponse
	data := map[string]interface{}{"targets": targets, "properties": properties}
	err := c.makeRequest(ctx, http.MethodPost, "metadata/user/get", nil, data, &resp)
	return &resp, err
}

// DeleteUserMetadata  删除用户属性
func (c *Client) DeleteUserMetadata(ctx context.Context, userID string) (*ResultResponse, error) {
	var resp ResultResponse
	p := path.Join("metadata/user", url.PathEscape(userID))
	err := c.makeRequest(ctx, http.MethodDelete, p, nil, nil, &resp)
	return &resp, err
}

// GetUserMetadataCapacity  获取 app 下用户属性总大小
func (c *Client) GetUserMetadataCapacity(ctx context.Context) (*ResultResponse, error) {
	var resp ResultResponse
	err := c.makeRequest(ctx, http.MethodGet, "metadata/user/capacity", nil, nil, &resp)
	return &resp, err
}
