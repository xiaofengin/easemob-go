package easemob_go

import (
	"context"
	"fmt"
	"testing"
)

func TestClient_CreateGroup(t *testing.T) {
	client, err := New("easemob-demo#support",
		"YXA6DWY3t8VBQke0my0Q4RJdeQ",
		"YXA6Q3fF5ZwxbbXaAb3zOCYoQhKKAH4")
	if err != nil {
		return
	}
	p := CreateGroupParam{
		Groupname:   "testgroup",
		Description: "test",
		Public:      false,
		Maxusers:    300,
		Owner:       "wf1",
		Members:     []string{"wf2"},
	}
	ret, err := client.CreateGroup(context.Background(), &p)
	if err != nil {
		return
	}
	fmt.Printf("数据的值：%v\n", ret.Data)
	//233150354620422
}

func TestClient_BanGroup(t *testing.T) {
	client, err := New("easemob-demo#support",
		"YXA6DWY3t8VBQke0my0Q4RJdeQ",
		"YXA6Q3fF5ZwxbbXaAb3zOCYoQhKKAH4")
	if err != nil {
		return
	}
	ret, err := client.BanGroup(context.Background(), "233150354620422")
	if err != nil {
		return
	}
	fmt.Printf("数据的值：%v\n", ret.Data)
}
func TestClient_UnBanGroup(t *testing.T) {
	client, err := New("easemob-demo#support",
		"YXA6DWY3t8VBQke0my0Q4RJdeQ",
		"YXA6Q3fF5ZwxbbXaAb3zOCYoQhKKAH4")
	if err != nil {
		return
	}
	ret, err := client.UnBanGroup(context.Background(), "233150354620422")
	if err != nil {
		return
	}
	fmt.Printf("数据的值：%v\n", ret.Data)
}
func TestClient_UpdateGroup(t *testing.T) {
	client, err := New("easemob-demo#support",
		"YXA6DWY3t8VBQke0my0Q4RJdeQ",
		"YXA6Q3fF5ZwxbbXaAb3zOCYoQhKKAH4")
	if err != nil {
		return
	}
	p := UpdateGroupParam{
		Groupname:   "test groupname",
		Description: "updategroupinfo12311",
		Public:      true,
		Maxusers:    99,
	}
	ret, err := client.UpdateGroup(context.Background(), "233150354620422", &p)
	if err != nil {
		return
	}
	fmt.Printf("数据的值：%v\n", ret.Data)
}
func TestClient_GetAllGroup(t *testing.T) {
	client, err := New("easemob-demo#support",
		"YXA6DWY3t8VBQke0my0Q4RJdeQ",
		"YXA6Q3fF5ZwxbbXaAb3zOCYoQhKKAH4")
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
	client, err := New("easemob-demo#support",
		"YXA6DWY3t8VBQke0my0Q4RJdeQ",
		"YXA6Q3fF5ZwxbbXaAb3zOCYoQhKKAH4")
	if err != nil {
		return
	}
	ret, err := client.GetUserJoinedGroup(context.Background(), "233150354620422", "wf2")
	if err != nil {
		return
	}
	fmt.Printf("数据的值：%v\n", ret.Data)
}
func TestClient_GetGroupDetail(t *testing.T) {
	client, err := New("easemob-demo#support",
		"YXA6DWY3t8VBQke0my0Q4RJdeQ",
		"YXA6Q3fF5ZwxbbXaAb3zOCYoQhKKAH4")
	if err != nil {
		return
	}
	ret, err := client.GetGroupDetail(context.Background(), "233150354620422")
	if err != nil {
		return
	}
	fmt.Printf("数据的值：%v\n", ret.Data)
}
func TestClient_DeleteGroup(t *testing.T) {
	client, err := New("easemob-demo#support",
		"YXA6DWY3t8VBQke0my0Q4RJdeQ",
		"YXA6Q3fF5ZwxbbXaAb3zOCYoQhKKAH4")
	if err != nil {
		return
	}
	ret, err := client.DeleteGroup(context.Background(), "233150354620422")
	if err != nil {
		return
	}
	fmt.Printf("数据的值：%v\n", ret.Data)
}

func TestClient_GetGroupAnnouncement(t *testing.T) {
	client, err := New("easemob-demo#support",
		"YXA6DWY3t8VBQke0my0Q4RJdeQ",
		"YXA6Q3fF5ZwxbbXaAb3zOCYoQhKKAH4")
	if err != nil {
		return
	}
	ret, err := client.GetGroupAnnouncement(context.Background(), "233164044828677")
	if err != nil {
		return
	}
	fmt.Printf("数据的值：%v\n", ret.Data)
}
func TestClient_UpdateGroupAnnouncement(t *testing.T) {

	client, err := New("easemob-demo#support",
		"YXA6DWY3t8VBQke0my0Q4RJdeQ",
		"YXA6Q3fF5ZwxbbXaAb3zOCYoQhKKAH4")
	if err != nil {
		return
	}
	ret, err := client.UpdateGroupAnnouncement(context.Background(), "233164044828677", "群组公告…")
	if err != nil {
		return
	}
	fmt.Printf("数据的值：%v\n", ret.Data)
}

func TestClient_GetGroupShareFile(t *testing.T) {
	client, err := New("easemob-demo#support",
		"YXA6DWY3t8VBQke0my0Q4RJdeQ",
		"YXA6Q3fF5ZwxbbXaAb3zOCYoQhKKAH4")
	if err != nil {
		return
	}
	ret, err := client.GetGroupShareFile(context.Background(), "233164044828677", "5", "1")
	if err != nil {
		return
	}
	fmt.Printf("数据的值：%v\n", ret.Data)
}
func TestClient_UpdateGroupShareFile(t *testing.T) {
	client, err := New("easemob-demo#support",
		"YXA6DWY3t8VBQke0my0Q4RJdeQ",
		"YXA6Q3fF5ZwxbbXaAb3zOCYoQhKKAH4")
	if err != nil {
		return
	}
	ret, err := client.UpdateGroupShareFile(context.Background(), "233164044828677", "./README.md")
	if err != nil {
		return
	}
	fmt.Printf("数据的值：%v\n", ret.Data)
}

// func TestClient_DownloadGroupShareFile(t *testing.T) {
//
//		client, err := New("easemob-demo#support",
//			"YXA6DWY3t8VBQke0my0Q4RJdeQ",
//			"YXA6Q3fF5ZwxbbXaAb3zOCYoQhKKAH4")
//		if err != nil {
//			return
//		}
//		ret, err := client.DownloadGroupShareFile(context.Background(), "233164044828677", "9c742d30-9345-11ee-8c6a-95c432617ad3")
//		if err != nil {
//			return
//		}
//		fmt.Printf("数据的值：%v\n", ret.Data)
//	}
func TestClient_DeleteGroupShareFile(t *testing.T) {
	client, err := New("easemob-demo#support",
		"YXA6DWY3t8VBQke0my0Q4RJdeQ",
		"YXA6Q3fF5ZwxbbXaAb3zOCYoQhKKAH4")
	if err != nil {
		return
	}
	ret, err := client.DeleteGroupShareFile(context.Background(), "233164044828677", "9b33b360-9344-11ee-baf5-0b51cff18cba")
	if err != nil {
		return
	}
	fmt.Printf("数据的值：%v\n", ret.Data)
}

func TestClient_GetGroupMember(t *testing.T) {

	client, err := New("easemob-demo#support",
		"YXA6DWY3t8VBQke0my0Q4RJdeQ",
		"YXA6Q3fF5ZwxbbXaAb3zOCYoQhKKAH4")
	if err != nil {
		return
	}
	ret, err := client.GetGroupMember(context.Background(), "233164044828677", "5", "1")
	if err != nil {
		return
	}
	fmt.Printf("数据的值：%v\n", ret.Data)
}

func TestClient_AddGroupMembers(t *testing.T) {
	client, err := New("easemob-demo#support",
		"YXA6DWY3t8VBQke0my0Q4RJdeQ",
		"YXA6Q3fF5ZwxbbXaAb3zOCYoQhKKAH4")
	if err != nil {
		return
	}
	ret, err := client.AddGroupMembers(context.Background(), "233164044828677", []string{"wf2", "wf3", "wf4"})
	if err != nil {
		return
	}
	fmt.Printf("数据的值：%v\n", ret.Data)
}
func TestClient_DeleteGroupMembers(t *testing.T) {
	client, err := New("easemob-demo#support",
		"YXA6DWY3t8VBQke0my0Q4RJdeQ",
		"YXA6Q3fF5ZwxbbXaAb3zOCYoQhKKAH4")
	if err != nil {
		return
	}
	ret, err := client.DeleteGroupMembers(context.Background(), "233164044828677", []string{"wf2", "wf4"})
	if err != nil {
		return
	}
	fmt.Printf("数据的值：%v\n", ret.Data)
}

func TestClient_SetGroupMemberAttributes(t *testing.T) {
	client, err := New("easemob-demo#support",
		"YXA6DWY3t8VBQke0my0Q4RJdeQ",
		"YXA6Q3fF5ZwxbbXaAb3zOCYoQhKKAH4")
	if err != nil {
		return
	}
	ret, err := client.SetGroupMemberAttributes(context.Background(), "233164044828677", "wf1", map[string]string{"key2": "value2"})
	if err != nil {
		return
	}
	fmt.Printf("数据的值：%v\n", ret.Data)
}

func TestClient_GetGroupMemberAttributes(t *testing.T) {
	client, err := New("easemob-demo#support",
		"YXA6DWY3t8VBQke0my0Q4RJdeQ",
		"YXA6Q3fF5ZwxbbXaAb3zOCYoQhKKAH4")
	if err != nil {
		return
	}
	ret, err := client.GetGroupMemberAttributes(context.Background(), "233164044828677", "wf3")
	if err != nil {
		return
	}
	fmt.Printf("数据的值：%v\n", ret.Data)
}

func TestClient_GetGroupMembersAttributesData(t *testing.T) {
	client, err := New("easemob-demo#support",
		"YXA6DWY3t8VBQke0my0Q4RJdeQ",
		"YXA6Q3fF5ZwxbbXaAb3zOCYoQhKKAH4")
	if err != nil {
		return
	}
	ret, err := client.GetGroupMembersAttributesData(context.Background(), "233164044828677", []string{"wf1", "wf3"}, []string{"key1", "key2"})
	if err != nil {
		return
	}
	fmt.Printf("数据的值：%v\n", ret.Data)
}
func TestClient_GetGroupAdmin(t *testing.T) {
	client, err := New("easemob-demo#support",
		"YXA6DWY3t8VBQke0my0Q4RJdeQ",
		"YXA6Q3fF5ZwxbbXaAb3zOCYoQhKKAH4")
	if err != nil {
		return
	}
	ret, err := client.GetGroupAdmin(context.Background(), "233164044828677")
	if err != nil {
		return
	}
	fmt.Printf("数据的值：%v\n", ret.Data)
}

func TestClient_AddGroupAdmin(t *testing.T) {

	client, err := New("easemob-demo#support",
		"YXA6DWY3t8VBQke0my0Q4RJdeQ",
		"YXA6Q3fF5ZwxbbXaAb3zOCYoQhKKAH4")
	if err != nil {
		return
	}
	ret, err := client.AddGroupAdmin(context.Background(), "233164044828677", "wf4")
	if err != nil {
		return
	}
	fmt.Printf("数据的值：%v\n", ret.Data)
}

func TestClient_DeleteGroupAdmin(t *testing.T) {
	client, err := New("easemob-demo#support",
		"YXA6DWY3t8VBQke0my0Q4RJdeQ",
		"YXA6Q3fF5ZwxbbXaAb3zOCYoQhKKAH4")
	if err != nil {
		return
	}
	ret, err := client.DeleteGroupAdmin(context.Background(), "233164044828677", "wf4")
	if err != nil {
		return
	}
	fmt.Printf("数据的值：%v\n", ret.Data)
}

func TestClient_TransferGroupAdmin(t *testing.T) {
	client, err := New("easemob-demo#support",
		"YXA6DWY3t8VBQke0my0Q4RJdeQ",
		"YXA6Q3fF5ZwxbbXaAb3zOCYoQhKKAH4")
	if err != nil {
		return
	}
	ret, err := client.TransferGroupAdmin(context.Background(), "233164044828677", "wf1")
	if err != nil {
		return
	}
	fmt.Printf("数据的值：%v\n", ret.Data)
}
func TestClient_GetGroupBlocks(t *testing.T) {

	client, err := New("easemob-demo#support",
		"YXA6DWY3t8VBQke0my0Q4RJdeQ",
		"YXA6Q3fF5ZwxbbXaAb3zOCYoQhKKAH4")
	if err != nil {
		return
	}
	ret, err := client.GetGroupBlocks(context.Background(), "233164044828677")
	if err != nil {
		return
	}
	fmt.Printf("数据的值：%v\n", ret.Data)
}
func TestClient_AddGroupBlocks(t *testing.T) {
	client, err := New("easemob-demo#support",
		"YXA6DWY3t8VBQke0my0Q4RJdeQ",
		"YXA6Q3fF5ZwxbbXaAb3zOCYoQhKKAH4")
	if err != nil {
		return
	}
	ret, err := client.AddGroupBlocks(context.Background(), "233164044828677", []string{"wf1", "wf3"})
	if err != nil {
		return
	}
	fmt.Printf("数据的值：%v\n", ret.Data)
}

func TestClient_DeleteGroupBlocks(t *testing.T) {
	client, err := New("easemob-demo#support",
		"YXA6DWY3t8VBQke0my0Q4RJdeQ",
		"YXA6Q3fF5ZwxbbXaAb3zOCYoQhKKAH4")
	if err != nil {
		return
	}
	ret, err := client.DeleteGroupBlocks(context.Background(), "233164044828677", []string{"wf1", "wf3"})
	if err != nil {
		return
	}
	fmt.Printf("数据的值：%v\n", ret.Data)
}

func TestClient_GetGroupWhite(t *testing.T) {
	client, err := New("easemob-demo#support",
		"YXA6DWY3t8VBQke0my0Q4RJdeQ",
		"YXA6Q3fF5ZwxbbXaAb3zOCYoQhKKAH4")
	if err != nil {
		return
	}
	ret, err := client.GetGroupWhite(context.Background(), "233164044828677")
	if err != nil {
		return
	}
	fmt.Printf("数据的值：%v\n", ret.Data)
}

func TestClient_AddGroupWhites(t *testing.T) {

	client, err := New("easemob-demo#support",
		"YXA6DWY3t8VBQke0my0Q4RJdeQ",
		"YXA6Q3fF5ZwxbbXaAb3zOCYoQhKKAH4")
	if err != nil {
		return
	}
	ret, err := client.AddGroupWhites(context.Background(), "233164044828677", []string{"wf4"})
	if err != nil {
		return
	}
	fmt.Printf("数据的值：%v\n", ret.Data)
}

func TestClient_DeleteGroupWhite(t *testing.T) {
	client, err := New("easemob-demo#support",
		"YXA6DWY3t8VBQke0my0Q4RJdeQ",
		"YXA6Q3fF5ZwxbbXaAb3zOCYoQhKKAH4")
	if err != nil {
		return
	}
	ret, err := client.DeleteGroupWhite(context.Background(), "233164044828677", "wf4")
	if err != nil {
		return
	}
	fmt.Printf("数据的值：%v\n", ret.Data)
}

func TestClient_GetGroupMute(t *testing.T) {
	client, err := New("easemob-demo#support",
		"YXA6DWY3t8VBQke0my0Q4RJdeQ",
		"YXA6Q3fF5ZwxbbXaAb3zOCYoQhKKAH4")
	if err != nil {
		return
	}
	ret, err := client.GetGroupMute(context.Background(), "233164044828677")
	if err != nil {
		return
	}
	fmt.Printf("数据的值：%v\n", ret.Data)
}
func TestClient_AddGroupMute(t *testing.T) {
	client, err := New("easemob-demo#support",
		"YXA6DWY3t8VBQke0my0Q4RJdeQ",
		"YXA6Q3fF5ZwxbbXaAb3zOCYoQhKKAH4")
	if err != nil {
		return
	}
	ret, err := client.AddGroupMute(context.Background(), "233164044828677", 1000000, []string{"wf2", "wf3"})
	if err != nil {
		return
	}
	fmt.Printf("数据的值：%v\n", ret.Data)
}

func TestClient_AllGroupMute(t *testing.T) {
	client, err := New("easemob-demo#support",
		"YXA6DWY3t8VBQke0my0Q4RJdeQ",
		"YXA6Q3fF5ZwxbbXaAb3zOCYoQhKKAH4")
	if err != nil {
		return
	}
	ret, err := client.AllGroupMute(context.Background(), "233164044828677")
	if err != nil {
		return
	}
	fmt.Printf("数据的值：%v\n", ret.Data)
}

func TestClient_DeleteGroupMute(t *testing.T) {
	client, err := New("easemob-demo#support",
		"YXA6DWY3t8VBQke0my0Q4RJdeQ",
		"YXA6Q3fF5ZwxbbXaAb3zOCYoQhKKAH4")
	if err != nil {
		return
	}
	ret, err := client.DeleteGroupMute(context.Background(), "233164044828677", []string{"wf2"})
	if err != nil {
		return
	}
	fmt.Printf("数据的值：%v\n", ret.Data)
}

func TestClient_RemoveGroupMute(t *testing.T) {
	client, err := New("easemob-demo#support",
		"YXA6DWY3t8VBQke0my0Q4RJdeQ",
		"YXA6Q3fF5ZwxbbXaAb3zOCYoQhKKAH4")
	if err != nil {
		return
	}
	ret, err := client.RemoveGroupMute(context.Background(), "233164044828677")
	if err != nil {
		return
	}
	fmt.Printf("数据的值：%v\n", ret.Data)
}

func TestClient_GetAllThread(t *testing.T) {
	client, err := New("easemob-demo#support",
		"YXA6DWY3t8VBQke0my0Q4RJdeQ",
		"YXA6Q3fF5ZwxbbXaAb3zOCYoQhKKAH4")
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
	client, err := New("easemob-demo#support",
		"YXA6DWY3t8VBQke0my0Q4RJdeQ",
		"YXA6Q3fF5ZwxbbXaAb3zOCYoQhKKAH4")
	if err != nil {
		return
	}
	ret, err := client.GetUserJoinedThread(context.Background(), "wf1", "50", "", "")
	if err != nil {
		return
	}
	fmt.Printf("数据的值：%v\n", ret.Entities)
}

func TestClient_GetUserJoinedGroupThread(t *testing.T) {
	client, err := New("easemob-demo#support",
		"YXA6DWY3t8VBQke0my0Q4RJdeQ",
		"YXA6Q3fF5ZwxbbXaAb3zOCYoQhKKAH4")
	if err != nil {
		return
	}
	ret, err := client.GetUserJoinedGroupThread(context.Background(), "233164044828677", "wf1", "50", "", "")
	if err != nil {
		return
	}
	fmt.Printf("数据的值：%v\n", ret.Entities)
}

func TestClient_CreateThread(t *testing.T) {
	client, err := New("easemob-demo#support",
		"YXA6DWY3t8VBQke0my0Q4RJdeQ",
		"YXA6Q3fF5ZwxbbXaAb3zOCYoQhKKAH4")
	if err != nil {
		return
	}
	p := ThreadParam{
		GroupId: "233164044828677",
		Name:    "1",
		Owner:   "wf1",
		MsgId:   "1220920785539974764",
	}
	ret, err := client.CreateThread(context.Background(), &p)
	if err != nil {
		return
	}
	fmt.Printf("数据的值：%v\n", ret.Data)
}

func TestClient_UpdateThread(t *testing.T) {
	client, err := New("easemob-demo#support",
		"YXA6DWY3t8VBQke0my0Q4RJdeQ",
		"YXA6Q3fF5ZwxbbXaAb3zOCYoQhKKAH4")
	if err != nil {
		return
	}
	ret, err := client.UpdateThread(context.Background(), "test1234", "233239246602251")
	if err != nil {
		return
	}
	fmt.Printf("数据的值：%v\n", ret.Data)
}
func TestClient_DeleteThread(t *testing.T) {
	client, err := New("easemob-demo#support",
		"YXA6DWY3t8VBQke0my0Q4RJdeQ",
		"YXA6Q3fF5ZwxbbXaAb3zOCYoQhKKAH4")
	if err != nil {
		return
	}
	ret, err := client.DeleteThread(context.Background(), "233239246602251")
	if err != nil {
		return
	}
	fmt.Printf("数据的值：%v\n", ret.Data)
}

func TestClient_GetThreadMember(t *testing.T) {
	client, err := New("easemob-demo#support",
		"YXA6DWY3t8VBQke0my0Q4RJdeQ",
		"YXA6Q3fF5ZwxbbXaAb3zOCYoQhKKAH4")
	if err != nil {
		return
	}
	ret, err := client.GetThreadMember(context.Background(), "233240905449473", "", "")
	if err != nil {
		return
	}
	fmt.Printf("数据的值：%v\n", ret.Data)
}

func TestClient_AddThreadMember(t *testing.T) {
	client, err := New("easemob-demo#support",
		"YXA6DWY3t8VBQke0my0Q4RJdeQ",
		"YXA6Q3fF5ZwxbbXaAb3zOCYoQhKKAH4")
	if err != nil {
		return
	}
	ret, err := client.AddThreadMember(context.Background(), "233240905449473", []string{"wf2", "wf3"})
	if err != nil {
		return
	}
	fmt.Printf("数据的值：%v\n", ret.Data)
}

func TestClient_DeleteThreadMember(t *testing.T) {
	client, err := New("easemob-demo#support",
		"YXA6DWY3t8VBQke0my0Q4RJdeQ",
		"YXA6Q3fF5ZwxbbXaAb3zOCYoQhKKAH4")
	if err != nil {
		return
	}
	ret, err := client.DeleteThreadMember(context.Background(), "233240905449473", []string{"wf2", "wf3"})
	if err != nil {
		return
	}
	fmt.Printf("数据的值：%v\n", ret.Entities)
}
