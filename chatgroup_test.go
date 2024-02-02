package easemob_go

import (
	"context"
	"fmt"
	"os"
	"testing"
)

func TestClient_CreateGroup(t *testing.T) {
	client, err := New(os.Getenv("appkey"),
		os.Getenv("clientId"),
		os.Getenv("clientSecret"))
	if err != nil {
		return
	}
	p := CreateGroupParam{
		GroupName:   "testgroup",
		Description: "test",
		Public:      false,
		MaxUsers:    300,
		Owner:       "userID_1",
		Members:     []string{"userID_2"},
	}
	ret, err := client.CreateGroup(context.Background(), &p)
	if err != nil {
		return
	}
	fmt.Printf("数据的值：%v\n", ret.Data)
	//233150354620422
}

func TestClient_BanGroup(t *testing.T) {
	client, err := New(os.Getenv("appkey"),
		os.Getenv("clientId"),
		os.Getenv("clientSecret"))
	if err != nil {
		return
	}
	ret, err := client.BanGroup(context.Background(), "groupID")
	if err != nil {
		return
	}
	fmt.Printf("数据的值：%v\n", ret.Data)
}
func TestClient_UnBanGroup(t *testing.T) {
	client, err := New(os.Getenv("appkey"),
		os.Getenv("clientId"),
		os.Getenv("clientSecret"))
	if err != nil {
		return
	}
	ret, err := client.UnBanGroup(context.Background(), "groupID")
	if err != nil {
		return
	}
	fmt.Printf("数据的值：%v\n", ret.Data)
}
func TestClient_UpdateGroup(t *testing.T) {
	client, err := New(os.Getenv("appkey"),
		os.Getenv("clientId"),
		os.Getenv("clientSecret"))
	if err != nil {
		return
	}
	p := UpdateGroupParam{
		GroupName:   "test groupname",
		Description: "updategroupinfo12311",
		Public:      true,
		MaxUsers:    99,
	}
	ret, err := client.UpdateGroup(context.Background(), "groupID", &p)
	if err != nil {
		return
	}
	fmt.Printf("数据的值：%v\n", ret.Data)
}
func TestClient_GetAllGroup(t *testing.T) {
	client, err := New(os.Getenv("appkey"),
		os.Getenv("clientId"),
		os.Getenv("clientSecret"))
	if err != nil {
		return
	}
	ret, err := client.GetAllGroup(context.Background(), "", "5")
	if err != nil {
		return
	}
	fmt.Printf("数据的值：%v\n    下一次的cursor:%v\n", ret.Data, ret.Cursor)
}

func TestClient_GetUserJoinedGroup(t *testing.T) {
	client, err := New(os.Getenv("appkey"),
		os.Getenv("clientId"),
		os.Getenv("clientSecret"))
	if err != nil {
		return
	}
	ret, err := client.GetUserJoinedGroup(context.Background(), "groupID", "userID")
	if err != nil {
		return
	}
	fmt.Printf("数据的值：%v\n", ret.Data)
}
func TestClient_GetGroupDetail(t *testing.T) {
	client, err := New(os.Getenv("appkey"),
		os.Getenv("clientId"),
		os.Getenv("clientSecret"))
	if err != nil {
		return
	}
	ret, err := client.GetGroupDetail(context.Background(), "groupID")
	if err != nil {
		return
	}
	fmt.Printf("数据的值：%v\n", ret.Data)
}
func TestClient_DeleteGroup(t *testing.T) {
	client, err := New(os.Getenv("appkey"),
		os.Getenv("clientId"),
		os.Getenv("clientSecret"))
	if err != nil {
		return
	}
	ret, err := client.DeleteGroup(context.Background(), "groupID")
	if err != nil {
		return
	}
	fmt.Printf("数据的值：%v\n", ret.Data)
}

func TestClient_GetGroupAnnouncement(t *testing.T) {
	client, err := New(os.Getenv("appkey"),
		os.Getenv("clientId"),
		os.Getenv("clientSecret"))
	if err != nil {
		return
	}
	ret, err := client.GetGroupAnnouncement(context.Background(), "groupID")
	if err != nil {
		return
	}
	fmt.Printf("数据的值：%v\n", ret.Data)
}
func TestClient_UpdateGroupAnnouncement(t *testing.T) {

	client, err := New(os.Getenv("appkey"),
		os.Getenv("clientId"),
		os.Getenv("clientSecret"))
	if err != nil {
		return
	}
	ret, err := client.UpdateGroupAnnouncement(context.Background(), "groupID", "群组公告…")
	if err != nil {
		return
	}
	fmt.Printf("数据的值：%v\n", ret.Data)
}

func TestClient_GetGroupShareFile(t *testing.T) {
	client, err := New(os.Getenv("appkey"),
		os.Getenv("clientId"),
		os.Getenv("clientSecret"))
	if err != nil {
		return
	}
	ret, err := client.GetGroupShareFile(context.Background(), "groupID", "5", "1")
	if err != nil {
		return
	}
	fmt.Printf("数据的值：%v\n", ret.Data)
}
func TestClient_UpdateGroupShareFile(t *testing.T) {
	client, err := New(os.Getenv("appkey"),
		os.Getenv("clientId"),
		os.Getenv("clientSecret"))
	if err != nil {
		return
	}
	ret, err := client.UpdateGroupShareFile(context.Background(), "groupID", "./README.md")
	if err != nil {
		return
	}
	fmt.Printf("数据的值：%v\n", ret.Data)
}

func TestClient_DownloadGroupShareFile(t *testing.T) {
	client, err := New(os.Getenv("appkey"),
		os.Getenv("clientId"),
		os.Getenv("clientSecret"))
	if err != nil {
		return
	}
	ret, err := client.DownloadGroupShareFile(context.Background(), "groupID", "./examples/README.md", "fileID")
	if err != nil {
		return
	}
	fmt.Printf("数据的值：%v\n", ret.Data)
}
func TestClient_DeleteGroupShareFile(t *testing.T) {
	client, err := New(os.Getenv("appkey"),
		os.Getenv("clientId"),
		os.Getenv("clientSecret"))
	if err != nil {
		return
	}
	ret, err := client.DeleteGroupShareFile(context.Background(), "groupID", "fileID")
	if err != nil {
		return
	}
	fmt.Printf("数据的值：%v\n", ret.Data)
}

func TestClient_GetGroupMember(t *testing.T) {

	client, err := New(os.Getenv("appkey"),
		os.Getenv("clientId"),
		os.Getenv("clientSecret"))
	if err != nil {
		return
	}
	ret, err := client.GetGroupMember(context.Background(), "groupID", "5", "1")
	if err != nil {
		return
	}
	fmt.Printf("数据的值：%v\n", ret.Data)
}

func TestClient_AddGroupMembers(t *testing.T) {
	client, err := New(os.Getenv("appkey"),
		os.Getenv("clientId"),
		os.Getenv("clientSecret"))
	if err != nil {
		return
	}
	ret, err := client.AddGroupMembers(context.Background(), "groupID", []string{"userID_1", "userID_2", "userID_3"})
	if err != nil {
		return
	}
	fmt.Printf("数据的值：%v\n", ret.Data)
}
func TestClient_DeleteGroupMembers(t *testing.T) {
	client, err := New(os.Getenv("appkey"),
		os.Getenv("clientId"),
		os.Getenv("clientSecret"))
	if err != nil {
		return
	}
	ret, err := client.DeleteGroupMembers(context.Background(), "groupID", []string{"userID_1", "userID_2"})
	if err != nil {
		return
	}
	fmt.Printf("数据的值：%v\n", ret.Data)
}

func TestClient_SetGroupMemberAttributes(t *testing.T) {
	client, err := New(os.Getenv("appkey"),
		os.Getenv("clientId"),
		os.Getenv("clientSecret"))
	if err != nil {
		return
	}
	ret, err := client.SetGroupMemberAttributes(context.Background(), "groupID", "userID", map[string]string{"key2": "value2"})
	if err != nil {
		return
	}
	fmt.Printf("数据的值：%v\n", ret.Data)
}

func TestClient_GetGroupMemberAttributes(t *testing.T) {
	client, err := New(os.Getenv("appkey"),
		os.Getenv("clientId"),
		os.Getenv("clientSecret"))
	if err != nil {
		return
	}
	ret, err := client.GetGroupMemberAttributes(context.Background(), "groupID", "userID")
	if err != nil {
		return
	}
	fmt.Printf("数据的值：%v\n", ret.Data)
}

func TestClient_GetGroupMembersAttributesData(t *testing.T) {
	client, err := New(os.Getenv("appkey"),
		os.Getenv("clientId"),
		os.Getenv("clientSecret"))
	if err != nil {
		return
	}
	ret, err := client.GetGroupMembersAttributesData(context.Background(), "groupID", []string{"userID_1", "userID_2"}, []string{"key1", "key2"})
	if err != nil {
		return
	}
	fmt.Printf("数据的值：%v\n", ret.Data)
}
func TestClient_GetGroupAdmin(t *testing.T) {
	client, err := New(os.Getenv("appkey"),
		os.Getenv("clientId"),
		os.Getenv("clientSecret"))
	if err != nil {
		return
	}
	ret, err := client.GetGroupAdmin(context.Background(), "groupID")
	if err != nil {
		return
	}
	fmt.Printf("数据的值：%v\n", ret.Data)
}

func TestClient_AddGroupAdmin(t *testing.T) {

	client, err := New(os.Getenv("appkey"),
		os.Getenv("clientId"),
		os.Getenv("clientSecret"))
	if err != nil {
		return
	}
	ret, err := client.AddGroupAdmin(context.Background(), "groupID", "userID")
	if err != nil {
		return
	}
	fmt.Printf("数据的值：%v\n", ret.Data)
}

func TestClient_DeleteGroupAdmin(t *testing.T) {
	client, err := New(os.Getenv("appkey"),
		os.Getenv("clientId"),
		os.Getenv("clientSecret"))
	if err != nil {
		return
	}
	ret, err := client.DeleteGroupAdmin(context.Background(), "groupID", "userID")
	if err != nil {
		return
	}
	fmt.Printf("数据的值：%v\n", ret.Data)
}

func TestClient_TransferGroupAdmin(t *testing.T) {
	client, err := New(os.Getenv("appkey"),
		os.Getenv("clientId"),
		os.Getenv("clientSecret"))
	if err != nil {
		return
	}
	ret, err := client.TransferGroupAdmin(context.Background(), "groupID", "userID")
	if err != nil {
		return
	}
	fmt.Printf("数据的值：%v\n", ret.Data)
}
func TestClient_GetGroupBlocks(t *testing.T) {

	client, err := New(os.Getenv("appkey"),
		os.Getenv("clientId"),
		os.Getenv("clientSecret"))
	if err != nil {
		return
	}
	ret, err := client.GetGroupBlocks(context.Background(), "groupID")
	if err != nil {
		return
	}
	fmt.Printf("数据的值：%v\n", ret.Data)
}
func TestClient_AddGroupBlocks(t *testing.T) {
	client, err := New(os.Getenv("appkey"),
		os.Getenv("clientId"),
		os.Getenv("clientSecret"))
	if err != nil {
		return
	}
	ret, err := client.AddGroupBlocks(context.Background(), "groupID", []string{"userID_1", "userID_2"})
	if err != nil {
		return
	}
	fmt.Printf("数据的值：%v\n", ret.Data)
}

func TestClient_DeleteGroupBlocks(t *testing.T) {
	client, err := New(os.Getenv("appkey"),
		os.Getenv("clientId"),
		os.Getenv("clientSecret"))
	if err != nil {
		return
	}
	ret, err := client.DeleteGroupBlocks(context.Background(), "groupID", []string{"userID_1", "userID_2"})
	if err != nil {
		return
	}
	fmt.Printf("数据的值：%v\n", ret.Data)
}

func TestClient_GetGroupWhite(t *testing.T) {
	client, err := New(os.Getenv("appkey"),
		os.Getenv("clientId"),
		os.Getenv("clientSecret"))
	if err != nil {
		return
	}
	ret, err := client.GetGroupWhite(context.Background(), "groupID")
	if err != nil {
		return
	}
	fmt.Printf("数据的值：%v\n", ret.Data)
}

func TestClient_AddGroupWhites(t *testing.T) {

	client, err := New(os.Getenv("appkey"),
		os.Getenv("clientId"),
		os.Getenv("clientSecret"))
	if err != nil {
		return
	}
	ret, err := client.AddGroupWhites(context.Background(), "groupID", []string{"userID"})
	if err != nil {
		return
	}
	fmt.Printf("数据的值：%v\n", ret.Data)
}

func TestClient_DeleteGroupWhite(t *testing.T) {
	client, err := New(os.Getenv("appkey"),
		os.Getenv("clientId"),
		os.Getenv("clientSecret"))
	if err != nil {
		return
	}
	ret, err := client.DeleteGroupWhite(context.Background(), "groupID", "userID")
	if err != nil {
		return
	}
	fmt.Printf("数据的值：%v\n", ret.Data)
}

func TestClient_GetGroupMute(t *testing.T) {
	client, err := New(os.Getenv("appkey"),
		os.Getenv("clientId"),
		os.Getenv("clientSecret"))
	if err != nil {
		return
	}
	ret, err := client.GetGroupMute(context.Background(), "groupID")
	if err != nil {
		return
	}
	fmt.Printf("数据的值：%v\n", ret.Data)
}
func TestClient_AddGroupMute(t *testing.T) {
	client, err := New(os.Getenv("appkey"),
		os.Getenv("clientId"),
		os.Getenv("clientSecret"))
	if err != nil {
		return
	}
	ret, err := client.AddGroupMute(context.Background(), "groupID", 1000000, []string{"userID_1", "userID_2"})
	if err != nil {
		return
	}
	fmt.Printf("数据的值：%v\n", ret.Data)
}

func TestClient_AllGroupMute(t *testing.T) {
	client, err := New(os.Getenv("appkey"),
		os.Getenv("clientId"),
		os.Getenv("clientSecret"))
	if err != nil {
		return
	}
	ret, err := client.AllGroupMute(context.Background(), "groupID")
	if err != nil {
		return
	}
	fmt.Printf("数据的值：%v\n", ret.Data)
}

func TestClient_DeleteGroupMute(t *testing.T) {
	client, err := New(os.Getenv("appkey"),
		os.Getenv("clientId"),
		os.Getenv("clientSecret"))
	if err != nil {
		return
	}
	ret, err := client.DeleteGroupMute(context.Background(), "groupID", []string{"userID"})
	if err != nil {
		return
	}
	fmt.Printf("数据的值：%v\n", ret.Data)
}

func TestClient_RemoveGroupMute(t *testing.T) {
	client, err := New(os.Getenv("appkey"),
		os.Getenv("clientId"),
		os.Getenv("clientSecret"))
	if err != nil {
		return
	}
	ret, err := client.RemoveAllGroupMute(context.Background(), "groupID")
	if err != nil {
		return
	}
	fmt.Printf("数据的值：%v\n", ret.Data)
}

func TestClient_GetAllThread(t *testing.T) {
	client, err := New(os.Getenv("appkey"),
		os.Getenv("clientId"),
		os.Getenv("clientSecret"))
	if err != nil {
		return
	}
	ret, err := client.GetAllThread(context.Background(), "50", "", "")
	if err != nil {
		return
	}
	fmt.Printf("数据的值：%v\n", ret.Entities)
}

func TestClient_GetUserJoinedThread(t *testing.T) {
	client, err := New(os.Getenv("appkey"),
		os.Getenv("clientId"),
		os.Getenv("clientSecret"))
	if err != nil {
		return
	}
	ret, err := client.GetUserJoinedThread(context.Background(), "userID", "50", "", "")
	if err != nil {
		return
	}
	fmt.Printf("数据的值：%v\n", ret.Entities)
}

func TestClient_GetUserJoinedGroupThread(t *testing.T) {
	client, err := New(os.Getenv("appkey"),
		os.Getenv("clientId"),
		os.Getenv("clientSecret"))
	if err != nil {
		return
	}
	ret, err := client.GetUserJoinedGroupThread(context.Background(), "groupID", "userID", "50", "", "")
	if err != nil {
		return
	}
	fmt.Printf("数据的值：%v\n", ret.Entities)
}

func TestClient_CreateThread(t *testing.T) {
	client, err := New(os.Getenv("appkey"),
		os.Getenv("clientId"),
		os.Getenv("clientSecret"))
	if err != nil {
		return
	}
	p := ThreadParam{
		GroupId: "groupID",
		Name:    "1",
		Owner:   "wf1",
		MsgId:   "MsgId",
	}
	ret, err := client.CreateThread(context.Background(), &p)
	if err != nil {
		return
	}
	fmt.Printf("数据的值：%v\n", ret.Data)
}

func TestClient_UpdateThread(t *testing.T) {
	client, err := New(os.Getenv("appkey"),
		os.Getenv("clientId"),
		os.Getenv("clientSecret"))
	if err != nil {
		return
	}
	ret, err := client.UpdateThread(context.Background(), "test1234", "threadId")
	if err != nil {
		return
	}
	fmt.Printf("数据的值：%v\n", ret.Data)
}
func TestClient_DeleteThread(t *testing.T) {
	client, err := New(os.Getenv("appkey"),
		os.Getenv("clientId"),
		os.Getenv("clientSecret"))
	if err != nil {
		return
	}
	ret, err := client.DeleteThread(context.Background(), "threadId")
	if err != nil {
		return
	}
	fmt.Printf("数据的值：%v\n", ret.Data)
}

func TestClient_GetThreadMember(t *testing.T) {
	client, err := New(os.Getenv("appkey"),
		os.Getenv("clientId"),
		os.Getenv("clientSecret"))
	if err != nil {
		return
	}
	ret, err := client.GetThreadMember(context.Background(), "threadId", "", "")
	if err != nil {
		return
	}
	fmt.Printf("数据的值：%v\n", ret.Data)
}

func TestClient_AddThreadMember(t *testing.T) {
	client, err := New(os.Getenv("appkey"),
		os.Getenv("clientId"),
		os.Getenv("clientSecret"))
	if err != nil {
		return
	}
	ret, err := client.AddThreadMember(context.Background(), "threadId", []string{"userID_1", "userID_2"})
	if err != nil {
		return
	}
	fmt.Printf("数据的值：%v\n", ret.Data)
}

func TestClient_DeleteThreadMember(t *testing.T) {
	client, err := New(os.Getenv("appkey"),
		os.Getenv("clientId"),
		os.Getenv("clientSecret"))
	if err != nil {
		return
	}
	ret, err := client.DeleteThreadMember(context.Background(), "threadId", []string{"userID_1", "userID_2"})
	if err != nil {
		return
	}
	fmt.Printf("数据的值：%v\n", ret.Entities)
}
