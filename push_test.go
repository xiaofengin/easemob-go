package easemob_go

import (
	"context"
	"fmt"
	"os"
	"testing"
)

func TestClient_PushBinding(t *testing.T) {

	client, err := New(os.Getenv("appkey"),
		os.Getenv("clientId"),
		os.Getenv("clientSecret"),
		"https://a1.easemob.com")
	if err != nil {
		return
	}
	data := PushParam{
		DeviceId:     "8ce08cad-9369-4bdd-86c8-695a0d247cda",
		DeviceToken:  "ffffff",
		NotifierName: "notifierName",
	}
	ret, err := client.PushBinding(context.Background(), "userID_1", &data)
	if err != nil {
		return
	}
	fmt.Printf("数据的值：%v\n", ret.Entities)
}

func TestClient_PushBindingInfo(t *testing.T) {
	client, err := New(os.Getenv("appkey"),
		os.Getenv("clientId"),
		os.Getenv("clientSecret"),
		"https://a1.easemob.com")
	if err != nil {
		return
	}
	ret, err := client.PushBindingInfo(context.Background(), "userID_1")
	if err != nil {
		return
	}
	fmt.Printf("数据的值：%v\n", ret.Entities)
}

func TestClient_UpdateUserPushNickname(t *testing.T) {
	client, err := New(os.Getenv("appkey"),
		os.Getenv("clientId"),
		os.Getenv("clientSecret"),
		"https://a1.easemob.com")
	if err != nil {
		return
	}
	ret, err := client.UpdateUserPushNickname(context.Background(), "userID_1", "nickName")
	if err != nil {
		return
	}
	fmt.Printf("数据的值：%v\n", ret.Entities)
}

func TestClient_SetNotificationDisplayStyle(t *testing.T) {
	client, err := New(os.Getenv("appkey"),
		os.Getenv("clientId"),
		os.Getenv("clientSecret"),
		"https://a1.easemob.com")
	if err != nil {
		return
	}
	ret, err := client.SetNotificationDisplayStyle(context.Background(), "userID_1", "1")
	if err != nil {
		return
	}
	fmt.Printf("数据的值：%v\n", ret.Entities)
}

func TestClient_SetOfflinePush(t *testing.T) {
	client, err := New(os.Getenv("appkey"),
		os.Getenv("clientId"),
		os.Getenv("clientSecret"),
		"https://a1.easemob.com")
	if err != nil {
		return
	}
	data := OfflinePushParam{
		UserID:         "userID_1",
		ChatType:       "user",
		Key:            "userID_2",
		Type:           "NONE",
		IgnoreInterval: "21:30-08:00",
		IgnoreDuration: 86400000,
	}
	ret, err := client.SetOfflinePush(context.Background(), &data)
	if err != nil {
		return
	}
	fmt.Printf("数据的值：%v\n", ret.Data)
}

func TestClient_GetOfflinePush(t *testing.T) {
	client, err := New(os.Getenv("appkey"),
		os.Getenv("clientId"),
		os.Getenv("clientSecret"),
		"https://a1.easemob.com")
	if err != nil {
		return
	}
	ret, err := client.GetOfflinePush(context.Background(), "userID_1", "user", "userID_2")
	if err != nil {
		return
	}
	fmt.Printf("数据的值：%v\n", ret.Data)
}
