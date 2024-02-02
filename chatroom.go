package easemob_go

import (
	"context"
	"net/http"
	"net/url"
	"path"
	"strings"
)

type CreateRoomParam struct {
	Name        string   `json:"name"`               //聊天室名称，最大长度为 128 个字符。
	Description string   `json:"description"`        //聊天室描述，最大长度为 512 字符。
	MaxUsers    int      `json:"maxusers,omitempty"` //聊天室最大成员数（包括聊天室所有者）。取值范围为 [1,10,000]，默认值为 1000。如需调整请联系商务。
	Owner       string   `json:"owner"`              //聊天室所有者。
	Members     []string `json:"members,omitempty"`  //聊天室普通成员和管理员的用户 ID 数组，不包含聊天室所有者的用户 ID。该数组可包含的元素数量不超过 maxusers 的值。若传该参数，确保至少设置一个数组元素。
	Custom      string   `json:"custom,omitempty"`   //聊天室扩展信息，例如可以给聊天室添加业务相关的标记，不要超过 1,024 个字符。
}
type UpdateRoomParam struct {
	Name        string `json:"name"`               //聊天室名称，最大长度为 128 个字符。
	Description string `json:"description"`        //聊天室描述，最大长度为 512 字符。
	MaxUsers    int    `json:"maxusers,omitempty"` //聊天室最大成员数（包括聊天室所有者）。取值范围为 [1,10,000]，默认值为 1000。如需调整请联系商务。
}

// AddChatRoomsSuperAdmin  添加超级管理员  在即时通讯应用中，仅聊天室超级管理员具有在客户端创建聊天室的权限。
func (c *Client) AddChatRoomsSuperAdmin(ctx context.Context, superAdmin string) (*ResultResponse, error) {
	var resp ResultResponse
	data := map[string]interface{}{"superadmin": superAdmin}
	err := c.makeRequest(ctx, http.MethodPost, "chatrooms/super_admin", nil, data, &resp)
	return &resp, err
}

// GetChatRoomsSuperAdmin  分页获取超级管理员列表
// pagesize 每页返回的超级管理员数量，默认值为 10。
// pagenum 当前页码。默认从第 1 页开始获取
func (c *Client) GetChatRoomsSuperAdmin(ctx context.Context, pagesize, pagenum string) (*ResultResponse, error) {
	var resp ResultResponse
	values := url.Values{}
	if pagenum != "" {
		values.Add("pagenum", pagenum)
	}
	if pagesize != "" {
		values.Add("pagesize", pagesize)
	}
	err := c.makeRequest(ctx, http.MethodGet, "chatrooms/super_admin", values, nil, &resp)
	return &resp, err
}

// DeleteChatRoomsSuperAdmin  撤销超级管理员
func (c *Client) DeleteChatRoomsSuperAdmin(ctx context.Context, superAdmin string) (*ResultResponse, error) {
	var resp ResultResponse
	p := path.Join("chatrooms/super_admin", url.PathEscape(superAdmin))
	err := c.makeRequest(ctx, http.MethodDelete, p, nil, nil, &resp)
	return &resp, err
}

// GetAppChatRoomsList  获取 app 中的聊天室
// limit 每次期望返回的聊天室数量。取值范围为 [1,100]，默认值为 10。该参数仅在分页获取时为必需。
// cursor 数据查询的起始位置。
func (c *Client) GetAppChatRoomsList(ctx context.Context, limit, cursor string) (*ResultResponse, error) {
	var resp ResultResponse
	values := url.Values{}
	if limit != "" {
		values.Add("limit", limit)
	}
	values.Add("cursor", cursor)
	err := c.makeRequest(ctx, http.MethodGet, "chatrooms", values, nil, &resp)
	return &resp, err
}

// GetUserJoinedChatRooms  获取用户加入的聊天室
// pagesize 每页返回的聊天室数量，取值范围为 [1,1000]，默认值为 1000。
// pagenum 当前页码。默认从第 1 页开始获取
func (c *Client) GetUserJoinedChatRooms(ctx context.Context, userID, pagesize, pagenum string) (*ResultResponse, error) {
	var resp ResultResponse
	values := url.Values{}
	if pagenum != "" {
		values.Add("pagenum", pagenum)
	}
	if pagesize != "" {
		values.Add("pagesize", pagesize)
	}
	p := path.Join("users", url.PathEscape(userID), "joined_chatrooms")
	err := c.makeRequest(ctx, http.MethodGet, p, values, nil, &resp)
	return &resp, err
}

// CreateChatRooms  创建聊天室
func (c *Client) CreateChatRooms(ctx context.Context, param *CreateRoomParam) (*ResultResponse, error) {
	var resp ResultResponse
	err := c.makeRequest(ctx, http.MethodPost, "chatrooms", nil, param, &resp)
	return &resp, err
}

// GetChatRooms  查询聊天室详情
func (c *Client) GetChatRooms(ctx context.Context, chatroomID string) (*ResultResponse, error) {
	var resp ResultResponse
	p := path.Join("chatrooms", url.PathEscape(chatroomID))
	err := c.makeRequest(ctx, http.MethodGet, p, nil, nil, &resp)
	return &resp, err
}

// UpdateChatRooms  修改聊天室信息
func (c *Client) UpdateChatRooms(ctx context.Context, chatroomID string, param *UpdateRoomParam) (*ResultResponse, error) {
	var resp ResultResponse
	p := path.Join("chatrooms", url.PathEscape(chatroomID))
	err := c.makeRequest(ctx, http.MethodPut, p, nil, param, &resp)
	return &resp, err
}

// DeleteChatRooms  删除聊天室
func (c *Client) DeleteChatRooms(ctx context.Context, chatroomID string) (*ResultResponse, error) {
	var resp ResultResponse
	p := path.Join("chatrooms", url.PathEscape(chatroomID))
	err := c.makeRequest(ctx, http.MethodDelete, p, nil, nil, &resp)
	return &resp, err
}

// GetRoomAnnouncement  获取聊天室公告
func (c *Client) GetRoomAnnouncement(ctx context.Context, chatroomID string) (*ResultResponse, error) {
	var resp ResultResponse
	p := path.Join("chatrooms", url.PathEscape(chatroomID), "announcement")
	err := c.makeRequest(ctx, http.MethodGet, p, nil, nil, &resp)
	return &resp, err
}

// UpdateRoomAnnouncement  修改群组公告
func (c *Client) UpdateRoomAnnouncement(ctx context.Context, chatroomID, announcement string) (*ResultResponse, error) {
	var resp ResultResponse
	data := map[string]string{"announcement": announcement}
	p := path.Join("chatrooms", url.PathEscape(chatroomID), "announcement")
	err := c.makeRequest(ctx, http.MethodPost, p, nil, data, &resp)
	return &resp, err
}

// SetRoomMetadata  设置聊天室自定义属性
// chatroomID 聊天室id
// userID  用户ID
// autoDelete 当前成员退出聊天室时是否自动删除该自定义属性。（默认）'DELETE'：是； 'NO_DELETE'：否。
// metaData 聊天室的自定义属性，存储为键值对（key-value）集合，即 Map<String,String>。该集合中最多可包含 10 个键值对，在每个键值对中，key 为属性名称，最多可包含 128 个字符；value 为属性值，不能超过 4096 个字符。每个聊天室最多可有 100 个自定义属性，每个应用的聊天室自定义属性总大小为 10 GB。
func (c *Client) SetRoomMetadata(ctx context.Context, chatroomID, userID, autoDelete string, metaData map[string]string) (*ResultResponse, error) {
	var resp ResultResponse
	data := map[string]interface{}{"metaData": metaData}
	if autoDelete != "" {
		data["autoDelete"] = autoDelete
	}
	p := path.Join("metadata/chatroom", url.PathEscape(chatroomID), "user", url.PathEscape(userID))
	err := c.makeRequest(ctx, http.MethodPut, p, nil, data, &resp)
	return &resp, err
}

// GetRoomMetadata  获取聊天室自定义属性
// chatroomID 聊天室id
// userID  用户ID
// keys 聊天室自定义属性名称列表。
func (c *Client) GetRoomMetadata(ctx context.Context, chatroomID string, keys []string) (*ResultResponse, error) {
	var resp ResultResponse
	data := map[string]interface{}{"keys": keys}
	p := path.Join("metadata/chatroom", url.PathEscape(chatroomID))
	err := c.makeRequest(ctx, http.MethodPost, p, nil, data, &resp)
	return &resp, err
}

// DeleteRoomMetadata  删除聊天室自定义属性
// chatroomID 聊天室id
// userID  用户ID
func (c *Client) DeleteRoomMetadata(ctx context.Context, chatroomID, userID string, keys []string) (*ResultResponse, error) {
	var resp ResultResponse
	data := map[string]interface{}{"keys": keys}
	p := path.Join("metadata/chatroom", url.PathEscape(chatroomID), "user", url.PathEscape(userID))
	err := c.makeRequest(ctx, http.MethodDelete, p, nil, data, &resp)
	return &resp, err
}

// SetRoomMetadataForced  强制设置聊天室自定义属性
// chatroomID 聊天室id
// userID  用户ID
// autoDelete 当前成员退出聊天室时是否自动删除该自定义属性。（默认）'DELETE'：是； 'NO_DELETE'：否。
// metaData 聊天室的自定义属性，存储为键值对（key-value）集合，即 Map<String,String>。该集合中最多可包含 10 个键值对，在每个键值对中，key 为属性名称，最多可包含 128 个字符；value 为属性值，不能超过 4096 个字符。每个聊天室最多可有 100 个自定义属性，每个应用的聊天室自定义属性总大小为 10 GB。
func (c *Client) SetRoomMetadataForced(ctx context.Context, chatroomID, userID, autoDelete string, metaData map[string]string) (*ResultResponse, error) {
	var resp ResultResponse
	data := map[string]interface{}{"metaData": metaData}
	if autoDelete != "" {
		data["autoDelete"] = autoDelete
	}
	p := path.Join("metadata/chatroom", url.PathEscape(chatroomID), "user", url.PathEscape(userID), "forced")
	err := c.makeRequest(ctx, http.MethodPut, p, nil, data, &resp)
	return &resp, err
}

// DeleteRoomMetadataForced  强制删除聊天室自定义属性 该方法每次最多可删除 10 个自定义属性。
// chatroomID 聊天室id
// userID  用户ID
func (c *Client) DeleteRoomMetadataForced(ctx context.Context, chatroomID, userID string, keys []string) (*ResultResponse, error) {
	var resp ResultResponse
	data := map[string]interface{}{"keys": keys}
	p := path.Join("metadata/chatroom", url.PathEscape(chatroomID), "user", url.PathEscape(userID), "forced")
	err := c.makeRequest(ctx, http.MethodDelete, p, nil, data, &resp)
	return &resp, err
}

// GetRoomMember  分页获取聊天室成员列表
// pagesize 每页期望返回的共享文件数。取值范围为 [1,1000]，默认为 1000。
// pagenum 当前页码。默认从第 1 页开始获取
func (c *Client) GetRoomMember(ctx context.Context, chatroomID, pagesize, pagenum string) (*ResultResponse, error) {
	var resp ResultResponse
	values := url.Values{}
	values.Add("pagesize", pagesize)
	values.Add("pagenum", pagenum)
	p := path.Join("chatrooms", url.PathEscape(chatroomID), "users")
	err := c.makeRequest(ctx, http.MethodGet, p, values, nil, &resp)
	return &resp, err
}

// AddRoomMembers  批量添加聊天室成员
// usernames 一次性最多可添加 60 位用户。
func (c *Client) AddRoomMembers(ctx context.Context, chatroomID string, usernames []string) (*ResultResponse, error) {
	var resp ResultResponse
	data := map[string]interface{}{"usernames": usernames}
	p := path.Join("chatrooms", url.PathEscape(chatroomID), "users")
	err := c.makeRequest(ctx, http.MethodPost, p, nil, data, &resp)
	return &resp, err
}

// DeleteRoomMembers  批量移除聊天室成员
// usernames 每次最多可传 100 个用户 ID
func (c *Client) DeleteRoomMembers(ctx context.Context, chatroomID string, usernames []string) (*ResultResponse, error) {
	var resp ResultResponse
	p := path.Join("chatrooms", url.PathEscape(chatroomID), "users", strings.Join(usernames, ","))
	err := c.makeRequest(ctx, http.MethodDelete, p, nil, nil, &resp)
	return &resp, err
}

// GetRoomAdmin 获取聊天室管理员列表
func (c *Client) GetRoomAdmin(ctx context.Context, chatroomID string) (*ResultResponse, error) {
	var resp ResultResponse
	p := path.Join("chatrooms", url.PathEscape(chatroomID), "admin")
	err := c.makeRequest(ctx, http.MethodGet, p, nil, nil, &resp)
	return &resp, err
}

// AddRoomAdmin 添加聊天室管理员
// admin 要添加的新管理员的用户 ID。
func (c *Client) AddRoomAdmin(ctx context.Context, chatroomID, userID string) (*ResultResponse, error) {
	var resp ResultResponse
	p := path.Join("chatrooms", url.PathEscape(chatroomID), "admin")
	data := map[string]string{"newadmin": userID}
	err := c.makeRequest(ctx, http.MethodPost, p, nil, data, &resp)
	return &resp, err
}

// DeleteRoomAdmin 移除聊天室管理员
func (c *Client) DeleteRoomAdmin(ctx context.Context, chatroomID, userID string) (*ResultResponse, error) {
	var resp ResultResponse
	p := path.Join("chatrooms", url.PathEscape(chatroomID), "admin", url.PathEscape(userID))
	err := c.makeRequest(ctx, http.MethodDelete, p, nil, nil, &resp)
	return &resp, err
}

// GetRoomBlocks 查询聊天室黑名单
func (c *Client) GetRoomBlocks(ctx context.Context, chatroomID string) (*ResultResponse, error) {
	var resp ResultResponse
	p := path.Join("chatrooms", url.PathEscape(chatroomID), "blocks/users")
	err := c.makeRequest(ctx, http.MethodGet, p, nil, nil, &resp)
	return &resp, err
}

// AddRoomBlocks 批量添加用户至聊天室黑名单
func (c *Client) AddRoomBlocks(ctx context.Context, chatroomID string, usernames []string) (*ResultResponse, error) {
	var resp ResultResponse
	p := path.Join("chatrooms", url.PathEscape(chatroomID), "blocks/users")
	data := map[string]interface{}{"usernames": usernames}
	err := c.makeRequest(ctx, http.MethodPost, p, nil, data, &resp)
	return &resp, err
}

// DeleteRoomBlocks 从聊天室黑名单批量移除用户 你一次最多可移除 60 个用户。
func (c *Client) DeleteRoomBlocks(ctx context.Context, chatroomID string, usernames []string) (*ResultResponse, error) {
	var resp ResultResponse
	p := path.Join("chatrooms", url.PathEscape(chatroomID), "blocks/users", strings.Join(usernames, ","))
	err := c.makeRequest(ctx, http.MethodDelete, p, nil, nil, &resp)
	return &resp, err
}

// GetRoomWhite 查询聊天室白名单
func (c *Client) GetRoomWhite(ctx context.Context, chatroomID string) (*ResultResponse, error) {
	var resp ResultResponse
	p := path.Join("chatrooms", url.PathEscape(chatroomID), "white/users")
	err := c.makeRequest(ctx, http.MethodGet, p, nil, nil, &resp)
	return &resp, err
}

// AddRoomWhites 批量添加用户至聊天室白名单
func (c *Client) AddRoomWhites(ctx context.Context, chatroomID string, usernames []string) (*ResultResponse, error) {
	var resp ResultResponse
	p := path.Join("chatrooms", url.PathEscape(chatroomID), "white/users")
	data := map[string]interface{}{"usernames": usernames}
	err := c.makeRequest(ctx, http.MethodPost, p, nil, data, &resp)
	return &resp, err
}

// DeleteRoomWhite 从聊天室白名单移除用户
func (c *Client) DeleteRoomWhite(ctx context.Context, chatroomID, userID string) (*ResultResponse, error) {
	var resp ResultResponse
	p := path.Join("chatrooms", url.PathEscape(chatroomID), "white/users", url.PathEscape(userID))
	err := c.makeRequest(ctx, http.MethodDelete, p, nil, nil, &resp)
	return &resp, err
}

// GetRoomMute 获取当前聊天室的禁言用户列表。
func (c *Client) GetRoomMute(ctx context.Context, chatroomID string) (*ResultResponse, error) {
	var resp ResultResponse
	p := path.Join("chatrooms", url.PathEscape(chatroomID), "mute")
	err := c.makeRequest(ctx, http.MethodGet, p, nil, nil, &resp)
	return &resp, err
}

// AddRoomMute 禁言聊天室成员
// duration 禁言时长，从当前时间开始计算。单位为毫秒。-1 表示永久禁言。
// usernames 要添加到禁言列表的用户 ID 列表，每次最多可添加 60 个。
func (c *Client) AddRoomMute(ctx context.Context, chatroomID string, duration int, usernames []string) (*ResultResponse, error) {
	var resp ResultResponse
	p := path.Join("chatrooms", url.PathEscape(chatroomID), "mute")
	data := map[string]interface{}{"mute_duration": duration, "usernames": usernames}
	err := c.makeRequest(ctx, http.MethodPost, p, nil, data, &resp)
	return &resp, err
}

// AllRoomMute 禁言聊天室全体成员
func (c *Client) AllRoomMute(ctx context.Context, chatroomID string) (*ResultResponse, error) {
	var resp ResultResponse
	p := path.Join("chatrooms", url.PathEscape(chatroomID), "ban")
	err := c.makeRequest(ctx, http.MethodPost, p, nil, nil, &resp)
	return &resp, err
}

// DeleteRoomMute 解除聊天室禁言成员
func (c *Client) DeleteRoomMute(ctx context.Context, chatroomID string, usernames []string) (*ResultResponse, error) {
	var resp ResultResponse
	p := path.Join("chatrooms", url.PathEscape(chatroomID), "mute", strings.Join(usernames, ","))
	err := c.makeRequest(ctx, http.MethodDelete, p, nil, nil, &resp)
	return &resp, err
}

// RemoveAllRoomMute 解除聊天室全员禁言
func (c *Client) RemoveAllRoomMute(ctx context.Context, chatroomID string) (*ResultResponse, error) {
	var resp ResultResponse
	p := path.Join("chatrooms", url.PathEscape(chatroomID), "ban")
	err := c.makeRequest(ctx, http.MethodDelete, p, nil, nil, &resp)
	return &resp, err
}
