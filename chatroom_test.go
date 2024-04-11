package easemob_go

import (
	"context"
	"fmt"
	"os"
	"testing"
)

func TestClient_AddChatRoomsSuperAdmin(t *testing.T) {

	client, err := New(os.Getenv("appkey"),
		os.Getenv("clientId"),
		os.Getenv("clientSecret"),
		"https://a1.easemob.com")
	if err != nil {
		return
	}

	ret, err := client.AddChatRoomsSuperAdmin(context.Background(), "userID_1")
	if err != nil {
		return
	}
	fmt.Printf("数据的值：%v\n", ret.Data)
}

func TestClient_GetChatRoomsSuperAdmin(t *testing.T) {
	client, err := New(os.Getenv("appkey"),
		os.Getenv("clientId"),
		os.Getenv("clientSecret"),
		"https://a1.easemob.com")
	if err != nil {
		return
	}

	ret, err := client.GetChatRoomsSuperAdmin(context.Background(), "", "")
	if err != nil {
		return
	}
	fmt.Printf("数据的值：%v\n", ret.Data)
}

func TestClient_DeleteChatRoomsSuperAdmin(t *testing.T) {
	client, err := New(os.Getenv("appkey"),
		os.Getenv("clientId"),
		os.Getenv("clientSecret"),
		"https://a1.easemob.com")
	if err != nil {
		return
	}

	ret, err := client.DeleteChatRoomsSuperAdmin(context.Background(), "userID_1")
	if err != nil {
		return
	}
	fmt.Printf("数据的值：%v\n", ret.Data)
}

func TestClient_GetAppChatRoomsList(t *testing.T) {
	client, err := New(os.Getenv("appkey"),
		os.Getenv("clientId"),
		os.Getenv("clientSecret"),
		"https://a1.easemob.com")
	if err != nil {
		return
	}

	ret, err := client.GetAppChatRoomsList(context.Background(), "", "")
	if err != nil {
		return
	}
	fmt.Printf("数据的值：%v\n", ret.Data)

}

func TestClient_GetUserJoinedChatRooms(t *testing.T) {
	client, err := New(os.Getenv("appkey"),
		os.Getenv("clientId"),
		os.Getenv("clientSecret"),
		"https://a1.easemob.com")
	if err != nil {
		return
	}

	ret, err := client.GetUserJoinedChatRooms(context.Background(), "userID_1", "", "")
	if err != nil {
		return
	}
	fmt.Printf("数据的值：%v\n", ret.Data)
}
func TestClient_CreateChatRooms(t *testing.T) {
	client, err := New(os.Getenv("appkey"),
		os.Getenv("clientId"),
		os.Getenv("clientSecret"),
		"https://a1.easemob.com")
	if err != nil {
		return
	}
	data := CreateRoomParam{
		Name:        "testchatroom",
		Description: "test",
		MaxUsers:    300,
		Owner:       "userID_1",
		Members:     []string{"userID_2"},
		Custom:      "",
	}
	ret, err := client.CreateChatRooms(context.Background(), &data)
	if err != nil {
		return
	}
	fmt.Printf("数据的值：%v\n", ret.Data)
}

func TestClient_GetChatRooms(t *testing.T) {
	client, err := New(os.Getenv("appkey"),
		os.Getenv("clientId"),
		os.Getenv("clientSecret"),
		"https://a1.easemob.com")
	if err != nil {
		return
	}
	ret, err := client.GetChatRooms(context.Background(), "roomID")
	if err != nil {
		return
	}
	fmt.Printf("数据的值：%v\n", ret.Data)
}

func TestClient_UpdateChatRooms(t *testing.T) {
	client, err := New(os.Getenv("appkey"),
		os.Getenv("clientId"),
		os.Getenv("clientSecret"),
		"https://a1.easemob.com")
	if err != nil {
		return
	}
	data := UpdateRoomParam{
		Name:        "Update 123",
		Description: "jjjjj",
		MaxUsers:    10,
	}
	ret, err := client.UpdateChatRooms(context.Background(), "roomID", &data)
	if err != nil {
		return
	}
	fmt.Printf("数据的值：%v\n", ret.Data)
}

func TestClient_DeleteChatRooms(t *testing.T) {
	client, err := New(os.Getenv("appkey"),
		os.Getenv("clientId"),
		os.Getenv("clientSecret"),
		"https://a1.easemob.com")
	if err != nil {
		return
	}

	ret, err := client.DeleteChatRooms(context.Background(), "roomID")
	if err != nil {
		return
	}
	fmt.Printf("数据的值：%v\n", ret.Data)
}

//234349519765506

func TestClient_GetRoomAnnouncement(t *testing.T) {
	client, err := New(os.Getenv("appkey"),
		os.Getenv("clientId"),
		os.Getenv("clientSecret"),
		"https://a1.easemob.com")
	if err != nil {
		return
	}

	ret, err := client.GetRoomAnnouncement(context.Background(), "roomID")
	if err != nil {
		return
	}
	fmt.Printf("数据的值：%v\n", ret.Data)
}

func TestClient_UpdateRoomAnnouncement(t *testing.T) {
	client, err := New(os.Getenv("appkey"),
		os.Getenv("clientId"),
		os.Getenv("clientSecret"),
		"https://a1.easemob.com")
	if err != nil {
		return
	}

	ret, err := client.UpdateRoomAnnouncement(context.Background(), "roomID", "Announcement")
	if err != nil {
		return
	}
	fmt.Printf("数据的值：%v\n", ret.Data)
}

func TestClient_SetRoomMetadata(t *testing.T) {
	client, err := New(os.Getenv("appkey"),
		os.Getenv("clientId"),
		os.Getenv("clientSecret"),
		"https://a1.easemob.com")
	if err != nil {
		return
	}

	data := map[string]string{
		"key1": "value1",
		"key2": "value2",
	}
	ret, err := client.SetRoomMetadata(context.Background(), "roomID", "userID_1", "", data)
	if err != nil {
		return
	}
	fmt.Printf("数据的值：%v\n", ret.Data)
}
func TestClient_GetRoomMetadata(t *testing.T) {

	client, err := New(os.Getenv("appkey"),
		os.Getenv("clientId"),
		os.Getenv("clientSecret"),
		"https://a1.easemob.com")
	if err != nil {
		return
	}

	ret, err := client.GetRoomMetadata(context.Background(), "roomID", []string{"userID_1", "userID_2"})
	if err != nil {
		return
	}
	fmt.Printf("数据的值：%v\n", ret.Data)
}

func TestClient_DeleteRoomMetadata(t *testing.T) {
	client, err := New(os.Getenv("appkey"),
		os.Getenv("clientId"),
		os.Getenv("clientSecret"),
		"https://a1.easemob.com")
	if err != nil {
		return
	}

	ret, err := client.DeleteRoomMetadata(context.Background(), "roomID", "userID_1", []string{"key1"})
	if err != nil {
		return
	}
	fmt.Printf("数据的值：%v\n", ret.Data)
}

func TestClient_SetRoomMetadataForced(t *testing.T) {
	client, err := New(os.Getenv("appkey"),
		os.Getenv("clientId"),
		os.Getenv("clientSecret"),
		"https://a1.easemob.com")
	if err != nil {
		return
	}

	data := map[string]string{
		"key1": "value11",
		"key2": "value22",
	}
	ret, err := client.SetRoomMetadataForced(context.Background(), "roomID", "userID_1", "", data)
	if err != nil {
		return
	}
	fmt.Printf("数据的值：%v\n", ret.Data)
}

func TestClient_DeleteRoomMetadataForced(t *testing.T) {
	client, err := New(os.Getenv("appkey"),
		os.Getenv("clientId"),
		os.Getenv("clientSecret"),
		"https://a1.easemob.com")
	if err != nil {
		return
	}

	ret, err := client.DeleteRoomMetadataForced(context.Background(), "roomID", "userID_1", []string{"key3"})
	if err != nil {
		return
	}
	fmt.Printf("数据的值：%v\n", ret.Data)
}

func TestClient_GetRoomMember(t *testing.T) {
	client, err := New(os.Getenv("appkey"),
		os.Getenv("clientId"),
		os.Getenv("clientSecret"),
		"https://a1.easemob.com")
	if err != nil {
		return
	}

	ret, err := client.GetRoomMember(context.Background(), "roomID", "", "")
	if err != nil {
		return
	}
	fmt.Printf("数据的值：%v\n", ret.Data)
}

func TestClient_AddRoomMembers(t *testing.T) {
	client, err := New(os.Getenv("appkey"),
		os.Getenv("clientId"),
		os.Getenv("clientSecret"),
		"https://a1.easemob.com")
	if err != nil {
		return
	}

	ret, err := client.AddRoomMembers(context.Background(), "roomID", []string{"userID_1"})
	if err != nil {
		return
	}
	fmt.Printf("数据的值：%v\n", ret.Data)
}

func TestClient_DeleteRoomMembers(t *testing.T) {
	client, err := New(os.Getenv("appkey"),
		os.Getenv("clientId"),
		os.Getenv("clientSecret"),
		"https://a1.easemob.com")
	if err != nil {
		return
	}

	ret, err := client.DeleteRoomMembers(context.Background(), "roomID", []string{"userID_1", "userID_2"})
	if err != nil {
		return
	}
	fmt.Printf("数据的值：%v\n", ret.Data)
}

func TestClient_GetRoomAdmin(t *testing.T) {
	client, err := New(os.Getenv("appkey"),
		os.Getenv("clientId"),
		os.Getenv("clientSecret"),
		"https://a1.easemob.com")
	if err != nil {
		return
	}

	ret, err := client.GetRoomAdmin(context.Background(), "roomID")
	if err != nil {
		return
	}
	fmt.Printf("数据的值：%v\n", ret.Data)
}
func TestClient_AddRoomAdmin(t *testing.T) {
	client, err := New(os.Getenv("appkey"),
		os.Getenv("clientId"),
		os.Getenv("clientSecret"),
		"https://a1.easemob.com")
	if err != nil {
		return
	}

	ret, err := client.AddRoomAdmin(context.Background(), "roomID", "userID_1")
	if err != nil {
		return
	}
	fmt.Printf("数据的值：%v\n", ret.Data)
}
func TestClient_DeleteRoomAdmin(t *testing.T) {
	client, err := New(os.Getenv("appkey"),
		os.Getenv("clientId"),
		os.Getenv("clientSecret"),
		"https://a1.easemob.com")
	if err != nil {
		return
	}

	ret, err := client.DeleteRoomAdmin(context.Background(), "roomID", "userID_1")
	if err != nil {
		return
	}
	fmt.Printf("数据的值：%v\n", ret.Data)
}
func TestClient_GetRoomBlocks(t *testing.T) {
	client, err := New(os.Getenv("appkey"),
		os.Getenv("clientId"),
		os.Getenv("clientSecret"),
		"https://a1.easemob.com")
	if err != nil {
		return
	}

	ret, err := client.GetRoomBlocks(context.Background(), "roomID")
	if err != nil {
		return
	}
	fmt.Printf("数据的值：%v\n", ret.Data)
}

func TestClient_AddRoomBlocks(t *testing.T) {
	client, err := New(os.Getenv("appkey"),
		os.Getenv("clientId"),
		os.Getenv("clientSecret"),
		"https://a1.easemob.com")
	if err != nil {
		return
	}

	ret, err := client.AddRoomBlocks(context.Background(), "roomID", []string{"userID_1"})
	if err != nil {
		return
	}
	fmt.Printf("数据的值：%v\n", ret.Data)
}

func TestClient_DeleteRoomBlocks(t *testing.T) {
	client, err := New(os.Getenv("appkey"),
		os.Getenv("clientId"),
		os.Getenv("clientSecret"),
		"https://a1.easemob.com")
	if err != nil {
		return
	}

	ret, err := client.DeleteRoomBlocks(context.Background(), "roomID", []string{"userID_1"})
	if err != nil {
		return
	}
	fmt.Printf("数据的值：%v\n", ret.Data)
}

func TestClient_GetRoomWhite(t *testing.T) {
	client, err := New(os.Getenv("appkey"),
		os.Getenv("clientId"),
		os.Getenv("clientSecret"),
		"https://a1.easemob.com")
	if err != nil {
		return
	}

	ret, err := client.GetRoomWhite(context.Background(), "roomID")
	if err != nil {
		return
	}
	fmt.Printf("数据的值：%v\n", ret.Data)
}
func TestClient_AddRoomWhites(t *testing.T) {
	client, err := New(os.Getenv("appkey"),
		os.Getenv("clientId"),
		os.Getenv("clientSecret"),
		"https://a1.easemob.com")
	if err != nil {
		return
	}

	ret, err := client.AddRoomWhites(context.Background(), "roomID", []string{"userID_1"})
	if err != nil {
		return
	}
	fmt.Printf("数据的值：%v\n", ret.Data)
}
func TestClient_DeleteRoomWhite(t *testing.T) {
	client, err := New(os.Getenv("appkey"),
		os.Getenv("clientId"),
		os.Getenv("clientSecret"),
		"https://a1.easemob.com")
	if err != nil {
		return
	}

	ret, err := client.DeleteRoomWhite(context.Background(), "roomID", "userID_1")
	if err != nil {
		return
	}
	fmt.Printf("数据的值：%v\n", ret.Data)
}

func TestClient_GetRoomMute(t *testing.T) {
	client, err := New(os.Getenv("appkey"),
		os.Getenv("clientId"),
		os.Getenv("clientSecret"),
		"https://a1.easemob.com")
	if err != nil {
		return
	}

	ret, err := client.GetRoomMute(context.Background(), "roomID")
	if err != nil {
		return
	}
	fmt.Printf("数据的值：%v\n", ret.Data)
}

func TestClient_AddRoomMute(t *testing.T) {
	client, err := New(os.Getenv("appkey"),
		os.Getenv("clientId"),
		os.Getenv("clientSecret"),
		"https://a1.easemob.com")
	if err != nil {
		return
	}

	ret, err := client.AddRoomMute(context.Background(), "roomID", -1, []string{"userID_1", "userID_2"})
	if err != nil {
		return
	}
	fmt.Printf("数据的值：%v\n", ret.Data)
}

func TestClient_AllRoomMute(t *testing.T) {
	client, err := New(os.Getenv("appkey"),
		os.Getenv("clientId"),
		os.Getenv("clientSecret"),
		"https://a1.easemob.com")
	if err != nil {
		return
	}

	ret, err := client.AllRoomMute(context.Background(), "roomID")
	if err != nil {
		return
	}
	fmt.Printf("数据的值：%v\n", ret.Data)
}
func TestClient_DeleteRoomMute(t *testing.T) {
	client, err := New(os.Getenv("appkey"),
		os.Getenv("clientId"),
		os.Getenv("clientSecret"),
		"https://a1.easemob.com")
	if err != nil {
		return
	}

	ret, err := client.DeleteRoomMute(context.Background(), "roomID", []string{"userID_1", "userID_2"})
	if err != nil {
		return
	}
	fmt.Printf("数据的值：%v\n", ret.Data)
}
func TestClient_RemoveAllRoomMute(t *testing.T) {
	client, err := New(os.Getenv("appkey"),
		os.Getenv("clientId"),
		os.Getenv("clientSecret"),
		"https://a1.easemob.com")
	if err != nil {
		return
	}

	ret, err := client.RemoveAllRoomMute(context.Background(), "roomID")
	if err != nil {
		return
	}
	fmt.Printf("数据的值：%v\n", ret.Data)
}
