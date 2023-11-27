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

// CreateTextMsg  Creating text message
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

// CreateImageMsg Creating image message
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

// CreateAudioMsg Creating audio message
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

// CreateVideoMsg Creating video message
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

// CreateFileMsg Creating file message
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

// CreateLocMsg Creating location message
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

// CreateCmdMsg Creating cmd message
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

// CreateCustomMsg Creating custom message
func CreateCustomMsg(customEvent string, to []string) *MsgModel {
	b := map[string]interface{}{"customEvent": customEvent}
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
