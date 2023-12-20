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
	"path/filepath"
	"strings"
)

type CreateGroupParam struct {
	Groupname         string   `json:"groupname,omitempty"`           //群组名称，最大长度为 128 字符。
	Description       string   `json:"description,omitempty"`         //群组描述，最大长度为 512 字符。
	Public            bool     `json:"public"`                        //是否是公开群。公开群可以被搜索到，用户可以申请加入公开群；私有群无法被搜索到，因此需要群主或群管理员添加，用户才可以加入。
	Scale             string   `json:"scale,omitempty"`               //群组规模，取决于群成员总数 maxusers 参数。（默认）normal：普通群，即群成员总数不超过 3000。large：大型群，群成员总数超过 3000。
	Maxusers          int      `json:"maxusers,omitempty"`            //群组最大成员数（包括群主）。对于普通群，该参数的默认值为 200，大型群为 1000
	Membersonly       bool     `json:"membersonly,omitempty"`         //用户申请入群是否需要群主或者群管理员审批。（默认）false
	InviteNeedConfirm bool     `json:"invite_need_confirm,omitempty"` //邀请用户入群时是否需要被邀用户同意。（默认）true：
	Owner             string   `json:"owner"`                         //群主的用户 ID。
	Members           []string `json:"members,omitempty"`             //群成员的用户 ID 数组，不包含群主的用户 ID。该数组可包含的元素数量不超过 maxusers 的值。
	Custom            string   `json:"custom,omitempty"`              //群主的用户 ID。
}
type UpdateGroupParam struct {
	Groupname         string `json:"groupname,omitempty"`           //群组名称，最大长度为 128 字符。
	Description       string `json:"description,omitempty"`         //群组描述，最大长度为 512 字符。
	Public            bool   `json:"public,omitempty"`              //是否是公开群。公开群可以被搜索到，用户可以申请加入公开群；私有群无法被搜索到，因此需要群主或群管理员添加，用户才可以加入。
	Maxusers          int    `json:"maxusers,omitempty"`            //群组最大成员数（包括群主）。对于普通群，该参数的默认值为 200，大型群为 1000
	Membersonly       bool   `json:"membersonly,omitempty"`         //用户申请入群是否需要群主或者群管理员审批。（默认）false
	InviteNeedConfirm bool   `json:"invite_need_confirm,omitempty"` //邀请用户入群时是否需要被邀用户同意。（默认）true：
	Custom            string `json:"custom,omitempty"`              //群主的用户 ID。
}
type ThreadParam struct {
	GroupId string `json:"group_id"`
	Name    string `json:"name"`
	Owner   string `json:"owner"`
	MsgId   string `json:"msg_id"`
}

// CreateGroup  创建群组
func (c *Client) CreateGroup(ctx context.Context, param *CreateGroupParam) (*ResultResponse, error) {
	var resp ResultResponse
	err := c.makeRequest(ctx, http.MethodPost, "chatgroups", nil, param, &resp)
	return &resp, err
}

// BanGroup  封禁群组
func (c *Client) BanGroup(ctx context.Context, groupID string) (*ResultResponse, error) {
	var resp ResultResponse
	p := path.Join("chatgroups", url.PathEscape(groupID), "disable")
	err := c.makeRequest(ctx, http.MethodPost, p, nil, nil, &resp)
	return &resp, err
}

// UnBanGroup  解封群组
func (c *Client) UnBanGroup(ctx context.Context, groupID string) (*ResultResponse, error) {
	var resp ResultResponse
	p := path.Join("chatgroups", url.PathEscape(groupID), "enable")
	err := c.makeRequest(ctx, http.MethodPost, p, nil, nil, &resp)
	return &resp, err
}

// UpdateGroup  修改群组信息
func (c *Client) UpdateGroup(ctx context.Context, groupID string, param *UpdateGroupParam) (*ResultResponse, error) {
	var resp ResultResponse
	p := path.Join("chatgroups", url.PathEscape(groupID))
	err := c.makeRequest(ctx, http.MethodPut, p, nil, param, &resp)
	return &resp, err
}

// GetAllGroup  获取 App 中的群组
// limit 每次期望返回的群组数量。取值范围为 [1,1000]，默认值为 10。
// cursor 数据查询的起始位置。
func (c *Client) GetAllGroup(ctx context.Context, cursor, limit string) (*ResultResponse, error) {
	var resp ResultResponse
	values := url.Values{}
	values.Add("limit", limit)
	values.Add("cursor", cursor)
	err := c.makeRequest(ctx, http.MethodGet, "chatgroups", values, nil, &resp)
	return &resp, err
}

// GetUserJoinedGroup  查看指定用户是否已加入群组
func (c *Client) GetUserJoinedGroup(ctx context.Context, groupID, userID string) (*ResultResponse, error) {
	var resp ResultResponse
	p := path.Join("chatgroups", url.PathEscape(groupID), "user", url.PathEscape(userID), "is_joined")
	err := c.makeRequest(ctx, http.MethodGet, p, nil, nil, &resp)
	return &resp, err
}

// GetGroupDetail  获取群组详情
func (c *Client) GetGroupDetail(ctx context.Context, groupID string) (*ResultResponse, error) {
	var resp ResultResponse
	p := path.Join("chatgroups", url.PathEscape(groupID))
	err := c.makeRequest(ctx, http.MethodGet, p, nil, nil, &resp)
	return &resp, err
}

// DeleteGroup  删除群组
func (c *Client) DeleteGroup(ctx context.Context, groupID string) (*ResultResponse, error) {
	var resp ResultResponse
	p := path.Join("chatgroups", url.PathEscape(groupID))
	err := c.makeRequest(ctx, http.MethodDelete, p, nil, nil, &resp)
	return &resp, err
}

// GetGroupAnnouncement  获取群组公告
func (c *Client) GetGroupAnnouncement(ctx context.Context, groupID string) (*ResultResponse, error) {
	var resp ResultResponse
	p := path.Join("chatgroups", url.PathEscape(groupID), "announcement")
	err := c.makeRequest(ctx, http.MethodGet, p, nil, nil, &resp)
	return &resp, err
}

// UpdateGroupAnnouncement  修改群组公告
func (c *Client) UpdateGroupAnnouncement(ctx context.Context, groupID, announcement string) (*ResultResponse, error) {
	var resp ResultResponse
	data := map[string]string{"announcement": announcement}
	p := path.Join("chatgroups", url.PathEscape(groupID), "announcement")
	err := c.makeRequest(ctx, http.MethodPost, p, nil, data, &resp)
	return &resp, err
}

// GetGroupShareFile  获取群组共享文件
// pagesize 每页期望返回的共享文件数。取值范围为 [1,1000]，默认为 1000。
// pagenum 当前页码。默认从第 1 页开始获取
func (c *Client) GetGroupShareFile(ctx context.Context, groupID, pagesize, pagenum string) (*ResultResponse, error) {
	var resp ResultResponse
	values := url.Values{}
	values.Add("pagesize", pagesize)
	values.Add("pagenum", pagenum)
	p := path.Join("chatgroups", url.PathEscape(groupID), "share_files")
	err := c.makeRequest(ctx, http.MethodGet, p, values, nil, &resp)
	return &resp, err
}

// UpdateGroupShareFile  上传群组共享文件
func (c *Client) UpdateGroupShareFile(ctx context.Context, groupID, filePath string) (*ResultResponse, error) {
	var resp ResultResponse
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", filepath.Base(file.Name()))
	_, err = io.Copy(part, file)
	writer.Close()
	p := path.Join("chatgroups", url.PathEscape(groupID), "share_files")
	err = c.uploadingFile(ctx, http.MethodPost, p, writer.FormDataContentType(), body, &resp)
	return &resp, err
}

//// DownloadGroupShareFile  下载群组共享文件
//// fileID 群组共享文件ID
//func (c *Client) DownloadGroupShareFile(ctx context.Context, groupID, fileID string) (*ResultResponse, error) {
//	var resp ResultResponse
//	p := path.Join("chatgroups", url.PathEscape(groupID), "share_files", url.PathEscape(fileID))
//	err := c.makeRequest(ctx, http.MethodGet, p, nil, nil, &resp)
//	return &resp, err
//}

// DeleteGroupShareFile  下载群组共享文件
// fileID 群组共享文件ID
func (c *Client) DeleteGroupShareFile(ctx context.Context, groupID, fileID string) (*ResultResponse, error) {
	var resp ResultResponse
	p := path.Join("chatgroups", url.PathEscape(groupID), "share_files", url.PathEscape(fileID))
	err := c.makeRequest(ctx, http.MethodDelete, p, nil, nil, &resp)
	return &resp, err
}

// GetGroupMember  分页获取群成员列表
// pagesize 每页期望返回的共享文件数。取值范围为 [1,1000]，默认为 1000。
// pagenum 当前页码。默认从第 1 页开始获取
func (c *Client) GetGroupMember(ctx context.Context, groupID, pagesize, pagenum string) (*ResultResponse, error) {
	var resp ResultResponse
	values := url.Values{}
	values.Add("pagesize", pagesize)
	values.Add("pagenum", pagenum)
	p := path.Join("chatgroups", url.PathEscape(groupID), "users")
	err := c.makeRequest(ctx, http.MethodGet, p, values, nil, &resp)
	return &resp, err
}

// AddGroupMembers  批量添加群组成员
func (c *Client) AddGroupMembers(ctx context.Context, groupID string, usernames []string) (*ResultResponse, error) {
	var resp ResultResponse
	data := map[string]interface{}{"usernames": usernames}
	p := path.Join("chatgroups", url.PathEscape(groupID), "users")
	err := c.makeRequest(ctx, http.MethodPost, p, nil, data, &resp)
	return &resp, err
}

// DeleteGroupMembers  批量移除群组成员
func (c *Client) DeleteGroupMembers(ctx context.Context, groupID string, usernames []string) (*ResultResponse, error) {
	var resp ResultResponse
	p := path.Join("chatgroups", url.PathEscape(groupID), "users", strings.Join(usernames, ","))
	err := c.makeRequest(ctx, http.MethodDelete, p, nil, nil, &resp)
	return &resp, err
}

// SetGroupMemberAttributes 设置群成员自定义属性
func (c *Client) SetGroupMemberAttributes(ctx context.Context, groupID, userID string, metaData interface{}) (*ResultResponse, error) {
	var resp ResultResponse
	p := path.Join("metadata/chatgroup", url.PathEscape(groupID), "user", url.PathEscape(userID))
	data := map[string]interface{}{"metaData": metaData}
	err := c.makeRequest(ctx, http.MethodPut, p, nil, data, &resp)
	return &resp, err
}

// GetGroupMemberAttributes 获取单个群成员的所有自定义属性
func (c *Client) GetGroupMemberAttributes(ctx context.Context, groupID, userID string) (*ResultResponse, error) {
	var resp ResultResponse
	p := path.Join("metadata/chatgroup", url.PathEscape(groupID), "user", url.PathEscape(userID))
	err := c.makeRequest(ctx, http.MethodGet, p, nil, nil, &resp)
	return &resp, err
}

// GetGroupMembersAttributesData 根据属性 key 获取多个群成员的自定义属性
// targets 要获取自定义属性的群成员的用户 ID。一次最多可传 10 个用户 ID。
// properties 要获取自定义属性的 key 的数组。若该参数设置为空数组或不传，则获取这些群成员的所有自定义属性。
func (c *Client) GetGroupMembersAttributesData(ctx context.Context, groupID string, targets, properties []string) (*ResultResponse, error) {
	var resp ResultResponse
	p := path.Join("metadata/chatgroup", url.PathEscape(groupID), "get")
	data := map[string]interface{}{"targets": targets, "properties": properties}
	err := c.makeRequest(ctx, http.MethodPost, p, nil, data, &resp)
	return &resp, err
}

// GetGroupAdmin 获取群管理员列表
func (c *Client) GetGroupAdmin(ctx context.Context, groupID string) (*ResultResponse, error) {
	var resp ResultResponse
	p := path.Join("chatgroups", url.PathEscape(groupID), "admin")
	err := c.makeRequest(ctx, http.MethodGet, p, nil, nil, &resp)
	return &resp, err
}

// AddGroupAdmin 添加群管理员
// admin 要添加的新管理员的用户 ID。
func (c *Client) AddGroupAdmin(ctx context.Context, groupID, userID string) (*ResultResponse, error) {
	var resp ResultResponse
	p := path.Join("chatgroups", url.PathEscape(groupID), "admin")
	data := map[string]string{"newadmin": userID}
	err := c.makeRequest(ctx, http.MethodPost, p, nil, data, &resp)
	return &resp, err
}

// DeleteGroupAdmin 移除群管理员
func (c *Client) DeleteGroupAdmin(ctx context.Context, groupID, userID string) (*ResultResponse, error) {
	var resp ResultResponse
	p := path.Join("chatgroups", url.PathEscape(groupID), "admin", url.PathEscape(userID))
	err := c.makeRequest(ctx, http.MethodDelete, p, nil, nil, &resp)
	return &resp, err
}

// TransferGroupAdmin 转让群组
func (c *Client) TransferGroupAdmin(ctx context.Context, groupID, userID string) (*ResultResponse, error) {
	var resp ResultResponse
	p := path.Join("chatgroups", url.PathEscape(groupID))
	data := map[string]string{"newowner": userID}
	err := c.makeRequest(ctx, http.MethodPut, p, nil, data, &resp)
	return &resp, err
}

// GetGroupBlocks 查询群组黑名单
func (c *Client) GetGroupBlocks(ctx context.Context, groupID string) (*ResultResponse, error) {
	var resp ResultResponse
	p := path.Join("chatgroups", url.PathEscape(groupID), "blocks/users")
	err := c.makeRequest(ctx, http.MethodGet, p, nil, nil, &resp)
	return &resp, err
}

// AddGroupBlocks 批量添加用户至群组黑名单
func (c *Client) AddGroupBlocks(ctx context.Context, groupID string, usernames []string) (*ResultResponse, error) {
	var resp ResultResponse
	p := path.Join("chatgroups", url.PathEscape(groupID), "blocks/users")
	data := map[string]interface{}{"usernames": usernames}
	err := c.makeRequest(ctx, http.MethodPost, p, nil, data, &resp)
	return &resp, err
}

// DeleteGroupBlocks 从群组黑名单批量移除用户 你一次最多可移除 60 个用户。
func (c *Client) DeleteGroupBlocks(ctx context.Context, groupID string, usernames []string) (*ResultResponse, error) {
	var resp ResultResponse
	p := path.Join("chatgroups", url.PathEscape(groupID), "blocks/users", strings.Join(usernames, ","))
	err := c.makeRequest(ctx, http.MethodDelete, p, nil, nil, &resp)
	return &resp, err
}

// GetGroupWhite 查询群组白名单
func (c *Client) GetGroupWhite(ctx context.Context, groupID string) (*ResultResponse, error) {
	var resp ResultResponse
	p := path.Join("chatgroups", url.PathEscape(groupID), "white/users")
	err := c.makeRequest(ctx, http.MethodGet, p, nil, nil, &resp)
	return &resp, err
}

// AddGroupWhites 批量添加用户至群组白名单
func (c *Client) AddGroupWhites(ctx context.Context, groupID string, usernames []string) (*ResultResponse, error) {
	var resp ResultResponse
	p := path.Join("chatgroups", url.PathEscape(groupID), "white/users")
	data := map[string]interface{}{"usernames": usernames}
	err := c.makeRequest(ctx, http.MethodPost, p, nil, data, &resp)
	return &resp, err
}

// DeleteGroupWhite 从群组白名单移除用户
func (c *Client) DeleteGroupWhite(ctx context.Context, groupID, userID string) (*ResultResponse, error) {
	var resp ResultResponse
	p := path.Join("chatgroups", url.PathEscape(groupID), "white/users", url.PathEscape(userID))
	err := c.makeRequest(ctx, http.MethodDelete, p, nil, nil, &resp)
	return &resp, err
}

// GetGroupMute 获取禁言列表
func (c *Client) GetGroupMute(ctx context.Context, groupID string) (*ResultResponse, error) {
	var resp ResultResponse
	p := path.Join("chatgroups", url.PathEscape(groupID), "mute")
	err := c.makeRequest(ctx, http.MethodGet, p, nil, nil, &resp)
	return &resp, err
}

// AddGroupMute 禁言指定群成员
// duration 禁言时长，单位为毫秒。
// usernames 要添加到禁言列表的用户 ID 列表，每次最多可添加 60 个。
func (c *Client) AddGroupMute(ctx context.Context, groupID string, duration int, usernames []string) (*ResultResponse, error) {
	var resp ResultResponse
	p := path.Join("chatgroups", url.PathEscape(groupID), "mute")
	data := map[string]interface{}{"mute_duration": duration, "usernames": usernames}
	err := c.makeRequest(ctx, http.MethodPost, p, nil, data, &resp)
	return &resp, err
}

// AllGroupMute 禁言全体群成员
func (c *Client) AllGroupMute(ctx context.Context, groupID string) (*ResultResponse, error) {
	var resp ResultResponse
	p := path.Join("chatgroups", url.PathEscape(groupID), "ban")
	err := c.makeRequest(ctx, http.MethodPost, p, nil, nil, &resp)
	return &resp, err
}

// DeleteGroupMute 解除成员禁言
func (c *Client) DeleteGroupMute(ctx context.Context, groupID string, usernames []string) (*ResultResponse, error) {
	var resp ResultResponse
	p := path.Join("chatgroups", url.PathEscape(groupID), "mute", strings.Join(usernames, ","))
	err := c.makeRequest(ctx, http.MethodDelete, p, nil, nil, &resp)
	return &resp, err
}

// RemoveAllGroupMute 解除全员禁言
func (c *Client) RemoveAllGroupMute(ctx context.Context, groupID string) (*ResultResponse, error) {
	var resp ResultResponse
	p := path.Join("chatgroups", url.PathEscape(groupID), "ban")
	err := c.makeRequest(ctx, http.MethodDelete, p, nil, nil, &resp)
	return &resp, err
}

// GetAllThread 获取 app 中的子区
// limit 每次期望返回的子区数量，取值范围为 [1,50]，默认值为 50。
// cursor 数据查询的起始位置。
// sort 获取的子区的排序顺序： - asc：按子区创建时间的正序； - （默认）desc：按子区创建时间的倒序。
func (c *Client) GetAllThread(ctx context.Context, limit, cursor, sort string) (*ResultResponse, error) {
	var resp ResultResponse
	values := url.Values{}
	values.Add("limit", limit)
	values.Add("cursor", cursor)
	values.Add("sort", sort)
	err := c.makeRequest(ctx, http.MethodGet, "thread", values, nil, &resp)
	return &resp, err
}

// GetUserJoinedThread 获取单个用户加入的所有子区（分页获取）
// limit 每次期望返回的子区数量，取值范围为 [1,50]，默认值为 50。
// cursor 数据查询的起始位置。
// sort 获取的子区的排序顺序： - asc：按子区创建时间的正序； - （默认）desc：按子区创建时间的倒序。
func (c *Client) GetUserJoinedThread(ctx context.Context, userID, limit, cursor, sort string) (*ResultResponse, error) {
	var resp ResultResponse
	values := url.Values{}
	values.Add("limit", limit)
	values.Add("cursor", cursor)
	values.Add("sort", sort)
	p := path.Join("threads/user", url.PathEscape(userID))
	err := c.makeRequest(ctx, http.MethodGet, p, values, nil, &resp)
	return &resp, err
}

// GetUserJoinedGroupThread 获取单个用户在指定群组中加入的所有子区 (分页获取)
// limit 每次期望返回的子区数量，取值范围为 [1,50]，默认值为 50。
// cursor 数据查询的起始位置。
// sort 获取的子区的排序顺序： - asc：按子区创建时间的正序； - （默认）desc：按子区创建时间的倒序。
func (c *Client) GetUserJoinedGroupThread(ctx context.Context, groupID, userID, limit, cursor, sort string) (*ResultResponse, error) {
	var resp ResultResponse
	values := url.Values{}
	values.Add("limit", limit)
	values.Add("cursor", cursor)
	values.Add("sort", sort)
	p := path.Join("threads/chatgroups", url.PathEscape(groupID), "user", url.PathEscape(userID))
	err := c.makeRequest(ctx, http.MethodGet, p, values, nil, &resp)
	return &resp, err
}

// CreateThread 创建子区
func (c *Client) CreateThread(ctx context.Context, param *ThreadParam) (*ResultResponse, error) {
	var resp ResultResponse
	err := c.makeRequest(ctx, http.MethodPost, "thread", nil, param, &resp)
	return &resp, err
}

// UpdateThread 修改子区
// name 要修改的子区的名称。修改后的子区名称不能超过 64 个字符。
func (c *Client) UpdateThread(ctx context.Context, name, threadId string) (*ResultResponse, error) {
	var resp ResultResponse
	p := path.Join("thread", url.PathEscape(threadId))
	data := map[string]string{"name": name}
	err := c.makeRequest(ctx, http.MethodPut, p, nil, data, &resp)
	return &resp, err
}

// DeleteThread 删除子区
// name 要修改的子区的名称。修改后的子区名称不能超过 64 个字符。
func (c *Client) DeleteThread(ctx context.Context, threadId string) (*ResultResponse, error) {
	var resp ResultResponse
	p := path.Join("thread", url.PathEscape(threadId))
	err := c.makeRequest(ctx, http.MethodDelete, p, nil, nil, &resp)
	return &resp, err
}

// GetThreadMember 获取子区成员列表(分页)
// limit 每次期望返回的子区数量，取值范围为 [1,50]，默认值为 50。
// cursor 数据查询的起始位置。
func (c *Client) GetThreadMember(ctx context.Context, threadId, limit, cursor string) (*ResultResponse, error) {
	var resp ResultResponse
	values := url.Values{}
	values.Add("limit", limit)
	values.Add("cursor", cursor)
	p := path.Join("thread", url.PathEscape(threadId), "users")
	err := c.makeRequest(ctx, http.MethodGet, p, values, nil, &resp)
	return &resp, err
}

// AddThreadMember 用户批量加入子区
func (c *Client) AddThreadMember(ctx context.Context, threadId string, usernames []string) (*ResultResponse, error) {
	var resp ResultResponse
	p := path.Join("thread", url.PathEscape(threadId), "users")
	data := map[string]interface{}{"usernames": usernames}
	err := c.makeRequest(ctx, http.MethodPost, p, nil, data, &resp)
	return &resp, err
}

// DeleteThreadMember 批量踢出子区成员
func (c *Client) DeleteThreadMember(ctx context.Context, threadId string, usernames []string) (*ResultResponse, error) {
	var resp ResultResponse
	p := path.Join("thread", url.PathEscape(threadId), "users")
	data := map[string]interface{}{"usernames": usernames}
	err := c.makeRequest(ctx, http.MethodDelete, p, nil, data, &resp)
	return &resp, err
}
