package easemob_go

type ImportMsgModel struct {
	Target       string                 `json:"target"`
	Type         string                 `json:"type"`
	Body         map[string]interface{} `json:"body"`
	From         string                 `json:"from"`
	IsAckRead    bool                   `json:"is_ack_read"`
	MsgTimestamp int64                  `json:"msg_timestamp"`
	NeedDownload bool                   `json:"need_download"`
}

type MsgModel struct {
	From       string                 `json:"from"`
	To         []string               `json:"to"`
	Type       string                 `json:"type"`
	Body       map[string]interface{} `json:"body"`
	Ext        map[string]interface{} `json:"ext"`
	SyncDevice bool                   `json:"sync_device"`
	RouteType  string                 `json:"routetype"`
}

// CreateTextMsg  创建文本消息
func CreateTextMsg(text string, to []string) *MsgModel {
	b := map[string]interface{}{"msg": text}
	return &MsgModel{
		From:       "admin",
		To:         to,
		Type:       "txt",
		Body:       b,
		Ext:        nil,
		SyncDevice: true,
		RouteType:  "",
	}
}

// CreateImageMsg 创建图片消息
func CreateImageMsg(url, fileName string, to []string) *MsgModel {
	b := map[string]interface{}{"url": url, "filename": fileName}
	return &MsgModel{
		From:       "admin",
		To:         to,
		Type:       "img",
		Body:       b,
		Ext:        nil,
		SyncDevice: true,
		RouteType:  "",
	}
}

// CreateAudioMsg 创建音频消息
func CreateAudioMsg(url, fileName string, to []string, length int) *MsgModel {
	b := map[string]interface{}{"url": url, "filename": fileName, "length": length}
	return &MsgModel{
		From:       "admin",
		To:         to,
		Type:       "audio",
		Body:       b,
		Ext:        nil,
		SyncDevice: true,
		RouteType:  "",
	}
}

// CreateVideoMsg 创建视频消息
func CreateVideoMsg(url, thumb, fileName string, to []string) *MsgModel {
	b := map[string]interface{}{"url": url, "thumb": thumb, "filename": fileName}
	return &MsgModel{
		From:       "admin",
		To:         to,
		Type:       "video",
		Body:       b,
		Ext:        nil,
		SyncDevice: true,
		RouteType:  "",
	}
}

// CreateFileMsg 创建文件消息
func CreateFileMsg(url, fileName string, to []string) *MsgModel {
	b := map[string]interface{}{"url": url, "filename": fileName}
	return &MsgModel{
		From:       "admin",
		To:         to,
		Type:       "file",
		Body:       b,
		Ext:        nil,
		SyncDevice: true,
		RouteType:  "",
	}
}

// CreateLocMsg 创建位置消息
func CreateLocMsg(lat, lng, addr string, to []string) *MsgModel {
	b := map[string]interface{}{"lat": lat, "lng": lng, "addr": addr}
	return &MsgModel{
		From:       "admin",
		To:         to,
		Type:       "loc",
		Body:       b,
		Ext:        nil,
		SyncDevice: true,
		RouteType:  "",
	}
}

// CreateCmdMsg 创建cmd消息
func CreateCmdMsg(action string, to []string) *MsgModel {
	b := map[string]interface{}{"action": action}
	return &MsgModel{
		From:       "admin",
		To:         to,
		Type:       "cmd",
		Body:       b,
		Ext:        nil,
		SyncDevice: true,
		RouteType:  "",
	}
}

// CreateCustomMsg 创建自定义消息
func CreateCustomMsg(customEvent string, customExts map[string]string, to []string) *MsgModel {
	b := map[string]interface{}{"customEvent": customEvent, "customExts": customExts}
	return &MsgModel{
		From:       "admin",
		To:         to,
		Type:       "custom",
		Body:       b,
		Ext:        nil,
		SyncDevice: true,
		RouteType:  "",
	}
}
