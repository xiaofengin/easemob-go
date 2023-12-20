package easemob_go

import (
	"context"
	"net/http"
	"net/url"
	"path"
)

type PushParam struct {
	DeviceId     string `json:"device_id"`     //移动端设备标识，服务端用于识别设备，进行推送信息的绑定和解绑。
	DeviceToken  string `json:"device_token"`  //推送证书名称。传入的证书名称必需与你在环信控制台的添加推送证书页面上填写的证书名称一致，否则推送失败。 若 notifier_name 为空，表示解除当前设备与所有推送信息的绑定。
	NotifierName string `json:"notifier_name"` //推送证书名称。 传入的证书名称必需与你在环信控制台的添加推送证书页面上填写的证书名称一致，否则推送失败。 若 notifier_name 为空，表示解除当前设备与所有推送信息的绑定。
}
type OfflinePushParam struct {
	UserID string `json:"username"`
	//对象类型，即会话类型
	//- user：用户，表示单聊；
	//- chatgroup：群组，表示群聊。
	ChatType string `json:"chattype"`
	//对象名称：
	//-单聊时为对端用户的用户 ID；
	//-群聊时为群组 ID。
	//-如需设置 app 全局离线推送，chattype 需传 user，key 为当前用户 ID。
	Key string `json:"key"`
	//离线推送通知方式
	// - DEFAULT：指定的会话采用 app 的设置。该值仅对单聊或群聊会话有效，对 app 级别无效。
	// - ALL：接收全部离线消息的推送通知；
	// - AT：只接收提及当前用户的离线消息的推送通知。该字段推荐在群聊中使用。若提及一个或多个用户，需在创建消息时对 ext 字段传 "em_at_list":["user1", "user2" ...]；若提及所有人，对该字段传 "em_at_list":"all"。
	// - NONE：不接收离线消息的推送通知。
	// 若 app 和指定会话均设置了该参数，则该会话采用自身的设置，其他会话采用 app 的设置。
	Type string `json:"type"`
	// 每天触发离线推送免打扰的时间段，精确到分钟，格式为 HH:MM-HH:MM，例如 08:30-10:00。该时间为 24 小时制，免打扰时间段的开始时间和结束时间中的小时数和分钟数的取值范围分别为 [00,23] 和 [00,59]
	IgnoreInterval string `json:"ignoreInterval"`
	// 离线推送免打扰时长，单位为毫秒。该参数的取值范围为 [0,604800000]，0 表示该参数无效，604800000 表示免打扰模式持续 7 天
	IgnoreDuration int64 `json:"ignoreDuration"`
}

// PushBinding  绑定/解绑推送信息
// param 用户数据
func (c *Client) PushBinding(ctx context.Context, userID string, param *PushParam) (*ResultResponse, error) {
	var resp ResultResponse
	p := path.Join("users", url.PathEscape(userID), "push/binding")
	err := c.makeRequest(ctx, http.MethodPut, p, nil, param, &resp)
	return &resp, err
}

// PushBindingInfo  查询推送绑定信息
func (c *Client) PushBindingInfo(ctx context.Context, userID string) (*ResultResponse, error) {
	var resp ResultResponse
	p := path.Join("users", url.PathEscape(userID), "push/binding")
	err := c.makeRequest(ctx, http.MethodGet, p, nil, nil, &resp)
	return &resp, err
}

// UpdateUserPushNickname  设置离线推送时显示的昵称
func (c *Client) UpdateUserPushNickname(ctx context.Context, userID, nickname string) (*ResultResponse, error) {
	var resp ResultResponse
	data := map[string]string{"nickname": nickname}
	p := path.Join("users", url.PathEscape(userID))
	err := c.makeRequest(ctx, http.MethodPut, p, nil, data, &resp)
	return &resp, err
}

// SetNotificationDisplayStyle  设置离线推送通知的展示方式
// displayStyle （默认）0：推送标题为“您有一条新消息”，推送内容为“请点击查看”； 1：推送标题为“您有一条新消息”，推送内容为发送人昵称和离线消息的内容。
func (c *Client) SetNotificationDisplayStyle(ctx context.Context, userID, displayStyle string) (*ResultResponse, error) {
	var resp ResultResponse
	data := map[string]string{"notification_display_style": displayStyle}
	p := path.Join("users", url.PathEscape(userID))
	err := c.makeRequest(ctx, http.MethodPut, p, nil, data, &resp)
	return &resp, err
}

// SetOfflinePush  设置离线推送
// param 免打扰参数
func (c *Client) SetOfflinePush(ctx context.Context, param *OfflinePushParam) (*ResultResponse, error) {
	var resp ResultResponse
	data := map[string]interface{}{"type": param.Type, "ignoreInterval": param.IgnoreInterval, "ignoreDuration": param.IgnoreDuration}
	p := path.Join("users", url.PathEscape(param.UserID), "notification", url.PathEscape(param.ChatType), url.PathEscape(param.Key))
	err := c.makeRequest(ctx, http.MethodPut, p, nil, data, &resp)
	return &resp, err
}

// GetOfflinePush  查询离线推送设置
// UserID 用户ID
// ChatType 对象类型，即会话类型
// Key 对象名称：-单聊时为对端用户的用户 ID；-群聊时为群组 ID。
func (c *Client) GetOfflinePush(ctx context.Context, UserID, ChatType, Key string) (*ResultResponse, error) {
	var resp ResultResponse
	p := path.Join("users", url.PathEscape(UserID), "notification", url.PathEscape(ChatType), url.PathEscape(Key))
	err := c.makeRequest(ctx, http.MethodGet, p, nil, nil, &resp)
	return &resp, err
}
