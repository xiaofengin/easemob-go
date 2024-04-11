package easemob_go

import (
	"context"
	"fmt"
	"os"
	"testing"
)

func TestClient_GetUserToken(t *testing.T) {
	client, err := New(os.Getenv("appkey"),
		os.Getenv("clientId"),
		os.Getenv("clientSecret"),
		"https://a1.easemob.com")
	if err != nil {
		return
	}

	//通过用户 ID 和密码获取用户 token
	//data := TokenParam{
	//	GrantType: "password",
	//	Username:  "userID",
	//	Password:  "1",
	//	Ttl:       "1024000",
	//}

	//通过用户 ID 获取用户 token
	data := TokenParam{
		GrantType:      "inherit",
		Username:       "userID",
		AutoCreateUser: true,
		Ttl:            "1024000",
	}
	ret, err := client.GetUserToken(context.Background(), &data)
	if err != nil {
		return
	}
	fmt.Printf("数据的值：%v\n", ret.AccessToken)
}
func TestClient_UserRegister(t *testing.T) {
	client, err := New(os.Getenv("appkey"),
		os.Getenv("clientId"),
		os.Getenv("clientSecret"),
		"https://a1.easemob.com")
	if err != nil {
		return
	}
	user1 := UserRegisterParam{
		Username: "userID_1",
		Password: "1",
	}
	user2 := UserRegisterParam{
		Username: "userID_2",
		Password: "1",
	}
	users := []UserRegisterParam{user1, user2}
	ret, err := client.UserRegister(context.Background(), &users)
	if err != nil {
		return
	}
	fmt.Printf("数据的值：%v\n", ret.Entities)
}

func TestClient_GetUserDetail(t *testing.T) {
	client, err := New(os.Getenv("appkey"),
		os.Getenv("clientId"),
		os.Getenv("clientSecret"),
		"https://a1.easemob.com")
	if err != nil {
		return
	}
	ret, err := client.GetUserDetail(context.Background(), "userID_1")
	if err != nil {
		return
	}
	fmt.Printf("数据的值：%v\n", ret.Entities)
}

func TestClient_GetUserDetailList(t *testing.T) {
	client, err := New(os.Getenv("appkey"),
		os.Getenv("clientId"),
		os.Getenv("clientSecret"),
		"https://a1.easemob.com")
	if err != nil {
		return
	}
	ret, err := client.GetUserDetailList(context.Background(), "20", "")
	if err != nil {
		return
	}
	fmt.Printf("数据的值：%v\n", ret.Entities)
}

func TestClient_DeleteUser(t *testing.T) {
	client, err := New(os.Getenv("appkey"),
		os.Getenv("clientId"),
		os.Getenv("clientSecret"),
		"https://a1.easemob.com")
	if err != nil {
		return
	}
	ret, err := client.DeleteUser(context.Background(), "userID_1")
	if err != nil {
		return
	}
	fmt.Printf("数据的值：%v\n", ret.Entities)
}
func TestClient_UpdateUserPassword(t *testing.T) {
	client, err := New(os.Getenv("appkey"),
		os.Getenv("clientId"),
		os.Getenv("clientSecret"),
		"https://a1.easemob.com")
	if err != nil {
		return
	}
	ret, err := client.UpdateUserPassword(context.Background(), "userID_1", "1")
	if err != nil {
		return
	}
	fmt.Printf("数据的值：%v\n", ret.Action)
}

func TestClient_GetUserStatusList(t *testing.T) {
	client, err := New(os.Getenv("appkey"),
		os.Getenv("clientId"),
		os.Getenv("clientSecret"),
		"https://a1.easemob.com")
	if err != nil {
		return
	}
	ret, err := client.GetUserStatusList(context.Background(), []string{"userID_1", "userID_2"})
	if err != nil {
		return
	}
	fmt.Printf("数据的值：%v\n", ret.Data)
}

func TestClient_GetUserResources(t *testing.T) {

	client, err := New(os.Getenv("appkey"),
		os.Getenv("clientId"),
		os.Getenv("clientSecret"),
		"https://a1.easemob.com")
	if err != nil {
		return
	}
	ret, err := client.GetUserResources(context.Background(), "userID_1")
	if err != nil {
		return
	}
	fmt.Printf("数据的值：%v\n", ret.Data)
}

func TestClient_SetUserMutes(t *testing.T) {
	client, err := New(os.Getenv("appkey"),
		os.Getenv("clientId"),
		os.Getenv("clientSecret"),
		"https://a1.easemob.com")
	if err != nil {
		return
	}
	data := MutesParam{
		Username:  "userID_1",
		Chat:      0,
		GroupChat: 0,
		Chatroom:  0,
	}
	ret, err := client.SetUserMutes(context.Background(), &data)
	if err != nil {
		return
	}
	fmt.Printf("数据的值：%v\n", ret.Data)

}

func TestClient_GetUserMutes(t *testing.T) {

	client, err := New(os.Getenv("appkey"),
		os.Getenv("clientId"),
		os.Getenv("clientSecret"),
		"https://a1.easemob.com")
	if err != nil {
		return
	}
	ret, err := client.GetUserMutes(context.Background(), "userID_1")
	if err != nil {
		return
	}
	fmt.Printf("数据的值：%v\n", ret.Data)
}

func TestClient_GetAppMutesList(t *testing.T) {
	client, err := New(os.Getenv("appkey"),
		os.Getenv("clientId"),
		os.Getenv("clientSecret"),
		"https://a1.easemob.com")
	if err != nil {
		return
	}
	ret, err := client.GetAppMutesList(context.Background())
	if err != nil {
		return
	}
	fmt.Printf("数据的值：%v\n", ret.Data)
}
func TestClient_CountMissedMessages(t *testing.T) {
	client, err := New(os.Getenv("appkey"),
		os.Getenv("clientId"),
		os.Getenv("clientSecret"),
		"https://a1.easemob.com")
	if err != nil {
		return
	}
	if err != nil {
		return
	}
	ret, err := client.CountMissedMessages(context.Background(), "userID_1")
	if err != nil {
		return
	}
	fmt.Printf("数据的值：%v\n", ret.Data)
}
func TestClient_IsMessageDeliveredToUser(t *testing.T) {
	client, err := New(os.Getenv("appkey"),
		os.Getenv("clientId"),
		os.Getenv("clientSecret"),
		"https://a1.easemob.com")
	if err != nil {
		return
	}
	ret, err := client.IsMessageDeliveredToUser(context.Background(), "userID_1", "messageID")
	if err != nil {
		return
	}
	fmt.Printf("数据的值：%v\n", ret.Data)
}

func TestClient_AccountBan(t *testing.T) {
	client, err := New(os.Getenv("appkey"),
		os.Getenv("clientId"),
		os.Getenv("clientSecret"),
		"https://a1.easemob.com")
	if err != nil {
		return
	}
	ret, err := client.AccountBan(context.Background(), "userID_1")
	if err != nil {
		return
	}
	fmt.Printf("数据的值：%v\n", ret.Entities)
}

func TestClient_AccountUnban(t *testing.T) {
	client, err := New(os.Getenv("appkey"),
		os.Getenv("clientId"),
		os.Getenv("clientSecret"),
		"https://a1.easemob.com")
	if err != nil {
		return
	}
	ret, err := client.AccountUnban(context.Background(), "userID_1")
	if err != nil {
		return
	}
	fmt.Printf("数据的值：%v\n", ret.Action)
}

func TestClient_AccountDisconnect(t *testing.T) {

	client, err := New(os.Getenv("appkey"),
		os.Getenv("clientId"),
		os.Getenv("clientSecret"),
		"https://a1.easemob.com")
	if err != nil {
		return
	}
	ret, err := client.AccountDisconnect(context.Background(), "userID_1")
	if err != nil {
		return
	}
	fmt.Printf("数据的值：%v\n", ret.Data)
}

func TestClient_SetUserMetadata(t *testing.T) {
	client, err := New(os.Getenv("appkey"),
		os.Getenv("clientId"),
		os.Getenv("clientSecret"),
		"https://a1.easemob.com")
	if err != nil {
		return
	}
	data := map[string]string{"ext": "ext", "nickname": "nickname", "avatarurl": "https://www.easemob.com/avatar.png"}
	ret, err := client.SetUserMetadata(context.Background(), "userID_1", data)
	if err != nil {
		return
	}
	fmt.Printf("数据的值：%v\n", ret.Data)
}

func TestClient_GetUserMetadata(t *testing.T) {
	client, err := New(os.Getenv("appkey"),
		os.Getenv("clientId"),
		os.Getenv("clientSecret"),
		"https://a1.easemob.com")
	if err != nil {
		return
	}
	ret, err := client.GetUserMetadata(context.Background(), "userID_1")
	if err != nil {
		return
	}
	fmt.Printf("数据的值：%v\n", ret.Data)
}

func TestClient_GetUserMetadataList(t *testing.T) {

	client, err := New(os.Getenv("appkey"),
		os.Getenv("clientId"),
		os.Getenv("clientSecret"),
		"https://a1.easemob.com")
	if err != nil {
		return
	}

	properties := []string{"avatarurl", "ext", "nickname"}
	targets := []string{"userID_1", "userID_2", "userID_3"}

	ret, err := client.GetUserMetadataList(context.Background(), targets, properties)
	if err != nil {
		return
	}
	fmt.Printf("数据的值：%v\n", ret.Data)
}
func TestClient_DeleteUserMetadata(t *testing.T) {
	client, err := New(os.Getenv("appkey"),
		os.Getenv("clientId"),
		os.Getenv("clientSecret"),
		"https://a1.easemob.com")
	if err != nil {
		return
	}
	ret, err := client.DeleteUserMetadata(context.Background(), "userID_1")
	if err != nil {
		return
	}
	fmt.Printf("数据的值：%v\n", ret.Data)
}
func TestClient_GetUserMetadataCapacity(t *testing.T) {
	client, err := New(os.Getenv("appkey"),
		os.Getenv("clientId"),
		os.Getenv("clientSecret"),
		"https://a1.easemob.com")
	if err != nil {
		return
	}

	ret, err := client.GetUserMetadataCapacity(context.Background())
	if err != nil {
		return
	}
	fmt.Printf("数据的值：%v\n", ret.Data)
}
