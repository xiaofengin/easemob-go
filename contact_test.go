package easemob_go

import (
	"context"
	"fmt"
	"os"
	"testing"
)

func TestClient_AddContact(t *testing.T) {
	client, err := New(os.Getenv("appkey"),
		os.Getenv("clientId"),
		os.Getenv("clientSecret"),
		"https://a1.easemob.com")
	if err != nil {
		return
	}
	ret, err := client.AddContact(context.Background(), "userID_1", "userID_2")
	if err != nil {
		return
	}
	fmt.Printf("数据的值：%v\n", ret.Entities)
}

func TestClient_DeleteContact(t *testing.T) {
	client, err := New(os.Getenv("appkey"),
		os.Getenv("clientId"),
		os.Getenv("clientSecret"),
		"https://a1.easemob.com")
	if err != nil {
		return
	}
	ret, err := client.DeleteContact(context.Background(), "userID_1", "userID_2")
	if err != nil {
		return
	}
	fmt.Printf("数据的值：%v\n", ret.Entities)
}

func TestClient_SetContactRemark(t *testing.T) {

	client, err := New(os.Getenv("appkey"),
		os.Getenv("clientId"),
		os.Getenv("clientSecret"),
		"https://a1.easemob.com")
	if err != nil {
		return
	}
	ret, err := client.SetContactRemark(context.Background(), "userID_1", "userID_2", "remark")
	if err != nil {
		return
	}
	fmt.Printf("数据的值：%v\n", ret.Status)
}

func TestClient_GetContactList(t *testing.T) {
	client, err := New(os.Getenv("appkey"),
		os.Getenv("clientId"),
		os.Getenv("clientSecret"),
		"https://a1.easemob.com")
	if err != nil {
		return
	}
	ret, err := client.GetContactList(context.Background(), "userID_1", "", "20", "true")
	if err != nil {
		return
	}
	fmt.Printf("数据的值：%v\n", ret.Data)
}
func TestClient_GetAllContactList(t *testing.T) {
	client, err := New(os.Getenv("appkey"),
		os.Getenv("clientId"),
		os.Getenv("clientSecret"),
		"https://a1.easemob.com")
	if err != nil {
		return
	}
	ret, err := client.GetAllContactList(context.Background(), "userID_1")
	if err != nil {
		return
	}
	fmt.Printf("数据的值：%v\n", ret.Data)
}

func TestClient_AddContactBlocks(t *testing.T) {
	client, err := New(os.Getenv("appkey"),
		os.Getenv("clientId"),
		os.Getenv("clientSecret"),
		"https://a1.easemob.com")
	if err != nil {
		return
	}
	ret, err := client.AddContactBlocks(context.Background(), "userID_1", []string{"userID_2", "userID_3"})
	if err != nil {
		return
	}
	fmt.Printf("数据的值：%v\n", ret.Data)
}

func TestClient_GetContactBlocksList(t *testing.T) {
	client, err := New(os.Getenv("appkey"),
		os.Getenv("clientId"),
		os.Getenv("clientSecret"),
		"https://a1.easemob.com")
	if err != nil {
		return
	}
	ret, err := client.GetContactBlocksList(context.Background(), "userID_1", "5", "")
	if err != nil {
		return
	}
	fmt.Printf("数据的值：%v\n", ret.Data)
}

func TestClient_DeleteContactBlocks(t *testing.T) {
	client, err := New(os.Getenv("appkey"),
		os.Getenv("clientId"),
		os.Getenv("clientSecret"),
		"https://a1.easemob.com")
	if err != nil {
		return
	}
	ret, err := client.DeleteContactBlocks(context.Background(), "userID_1", "userID_2")
	if err != nil {
		return
	}
	fmt.Printf("数据的值：%v\n", ret.Entities)
}
