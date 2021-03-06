package libgocqhttp

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type File struct {
	Busid         int32  `json:"busid"`
	DeadTime      int64  `json:"dead_time"`
	DownloadTimes int32  `json:"download_times"`
	FileID        string `json:"file_id"`
	FileName      string `json:"file_name"`
	FileSize      int64  `json:"file_size"`
	GroupID       int32  `json:"group_id"`
	ModifyTime    int64  `json:"modify_time"`
	Uploader      int64  `json:"uploader"`
	UploaderName  string `json:"uploader_name"`
	UploadTime    int64  `json:"upload_time"`
}

type Folder struct {
	Creator        int64  `json:"creator"`
	CreatorName    string `json:"creator_name"`
	CreateTime     int64  `json:"create_time"`
	FolderID       string `json:"folder_id"`
	FolderName     string `json:"folder_name"`
	GroupID        int32  `json:"group_id"`
	TotalFileCount int32  `json:"total_file_count"`
}

type Client struct {
	AppId      int64  `json:"app_id"`
	DeviceKind string `json:"device_kind"`
	DeviceName string `json:"device_name"`
}

type _response struct {
	Data []struct {
		Age             int32  `json:"age"`
		Area            string `json:"area"`
		Card            string `json:"card"`
		CardChangeable  bool   `json:"card_changeable"`
		GroupId         int64  `json:"group_id"`
		JoinTime        int64  `json:"join_time"`
		LastSentTime    int32  `json:"last_sent_time""`
		Level           int32  `json:"level"`
		MessageId       int32  `json:"message_id"`
		Nickname        string `json:"nickname"`
		OperatorId      int64  `json:"operator_id"`
		OperatorNick    string `json:"operator_nick"`
		OperatorTime    int64  `json:"operator_time"`
		Remark          string `json:"remark"`
		Role            string `json:"role"`
		SenderId        int64  `json:"sender_id"`
		SenderNick      string `json:"sender_name"`
		SenderTime      int64  `json:"sender_time"`
		Sex             string `json:"sex"`
		ShunUpTimestamp int64  `json:"shun_up_timestamp"`
		Title           string `json:"title"`
		TitleExpireTime int64  `json:"title_expire_time"`
		Unfriendly      bool   `json:"unfriendly"`
		UserId          int64  `json:"user_id"`
	} `json:"data"`
	Message string `json:"msg"`
	Retcode int    `json:"retcode"`
	Status  string `json:"status"`
	Wording string `json:"wording"`
}

type response struct {
	Data struct {
		Age             int32    `json:"age"`
		Area            string   `json:"area"`
		AppEnabled      bool     `json:"app_enabled"`
		AppFullName     string   `json:"app_full_name"`
		AppGood         bool     `json:"app_good"`
		AppInitialized  bool     `json:"app_initialized"`
		AppName         string   `json:"app_name"`
		AppVersion      string   `json:"app_version"`
		CanAtAll        bool     `json:"can_at_all"`
		Card            string   `json:"card"`
		CardChangeable  bool     `json:"card_changeable"`
		Clients         []Client `json:"clients"`
		CoolqDirectory  string   `json:"coolq_directory"`
		CoolqEdition    string   `json:"coolq_edition"`
		CreateTime      int64    `json:"create_time"`
		ExtName         string   `json:"ext_name"`
		File            string   `json:"file"`
		FileCount       int32    `json:"file_count"`
		Filename        string   `json:"filename"`
		Files           []File   `json:"files"`
		Folders         []Folder `json:"folders"`
		Go_cqhttp       bool     `json:"go-cqhttp"`
		Good            bool     `json:"good"`
		GroupCreateTime uint32   `json:"group_create_time"`
		GroupId         int64    `json:"group_id"`
		GroupLevel      uint32   `json:"group_level"`
		GroupMemo       string   `json:"group_memo"`
		GroupName       string   `json:"group_name"`
		InvitedRequests []struct {
			RequestID   int64  `json:"request_id"`
			InvitorUin  int    `json:"invitor_uin"`
			InvitorNick string `json:"invitor_nick"`
			GroupID     int    `json:"group_id"`
			GroupName   string `json:"group_name"`
			Checked     bool   `json:"checked"`
			Actor       int    `json:"actor"`
		} `json:"invited_requests"`
		JoinRequests []struct {
			RequestID     int64  `json:"request_id"`
			Message       string `json:"message"`
			RequesterUin  int64  `json:"requester_uin"`
			RequesterNick string `json:"requester_nick"`
			GroupID       int    `json:"group_id"`
			GroupName     string `json:"group_name"`
			ActionUinNick string `json:"action_uin_nick"`
			ActionUin     int    `json:"action_uin"`
			Checked       bool   `json:"checked"`
			Actor         int    `json:"actor"`
			Suspicious    bool   `json:"suspicious"`
		} `json:"join_requests"`
		JoinTime       int32   `json:"join_time"`
		LastSentTime   int32   `json:"last_sent_time"`
		Level          int32   `json:"level"`
		LevelSpeed     float64 `json:"level_speed"`
		LimitCount     int32   `json:"limit_count"`
		LoginDays      int32   `json:"login_days"`
		MasterId       int64   `json:"master_id"`
		MaxMemberCount int32   `json:"max_member_count"`
		MemberCount    int32   `json:"member_count"`
		MessageID      int     `json:"message_id"`
		Messages       []struct {
			Anonymous struct {
				Flag string `json:"flag"`
				Id   string `json:"id"`
				Name string `json:"name"`
			} `json:"anonymous"`
			Font        int    `json:"font"`
			GroupID     int    `json:"group_id"`
			Message     string `json:"message"`
			MessageID   int    `json:"message_id"`
			MessageSeq  int    `json:"message_seq"`
			MessageType string `json:"message_type"`
			PostType    string `json:"post_type"`
			RawMessage  string `json:"raw_message"`
			SelfID      int    `json:"self_id"`
			Sender      struct {
				Age      int    `json:"age"`
				Area     string `json:"area"`
				Card     string `json:"card"`
				Level    string `json:"level"`
				Nickname string `json:"nickname"`
				Role     string `json:"role"`
				Sex      string `json:"sex"`
				Title    string `json:"title"`
				UserID   int    `json:"user_id"`
			} `json:"sender"`
			SubType string `json:"sub_type"`
			Time    int    `json:"time"`
			UserID  int    `json:"user_id"`
		} `json:"messages"`
		NickName                 string   `json:"nickname"`
		Online                   bool     `json:"online"`
		PluginBuildConfiguration int      `json:"plugin_build_configuration"`
		PluginBuildNumber        int      `json:"plugin_build_number"`
		PluginsGood              bool     `json:"plugins_good"`
		PluginVersion            string   `json:"plugin_version"`
		Protocol                 int      `json:"protocol"`
		ProtocolVersion          string   `json:"protocol_version"`
		Qid                      string   `json:"qid"`
		Remark                   string   `json:"remark"`
		RemainAtAllCountForGroup int16    `json:"remain_at_all_count_for_group"`
		RemainAtAllCountForUin   int16    `json:"remain_at_all_count_for_uin"`
		Role                     string   `json:"role"`
		RuntimeOs                string   `json:"runtime_os"`
		RuntimeVersion           string   `json:"runtime_version"`
		Sex                      string   `json:"sex"`
		ShutUpTimestamp          int64    `json:"shut_up_timestamp"`
		Size                     int      `json:"size"`
		Slices                   []string `json:"slices"`
		Stat                     struct {
			DisconnectTimes uint32 `json:"DisconnectTimes"`
			LastMessageTime int64  `json:"LastMessageTime"`
			LostTimes       uint32 `json:"LostTimes"`
			MessageReceived uint64 `json:"MessageReceived"`
			MessageSent     uint64 `json:"MessageSent"`
			PacketReceived  uint64 `json:"PacketReceived"`
			PacketSent      uint64 `json:"PacketSent"`
			PacketLost      uint32 `json:"PacketLost"`
		} `json:"stat"`
		Texts []struct {
			Text        string `json:"text"`
			Confidence  int    `json:"confidence"`
			Coordinates []struct {
				X int `json:"x"`
				Y int `json:"y"`
			} `json:"coordinates"`
		} `json:"texts"`
		Title           string `json:"title"`
		TitleExpireTime int64  `json:"title_expire_time"`
		TotalSpace      int64  `json:"total_space"`
		Unfriendly      bool   `json:"unfriendly"`
		UsedSpace       int64  `json:"used_space"`
		Url             string `json:"url"`
		UserId          int64  `json:"user_id"`
		Variants        []struct {
			ModelShow string `json:"model_show"`
			NeedPay   bool   `json:"need_pay"`
		} `json:"variants"`
		Version        string `json:"version"`
		VipGrowthSpeed int64  `json:"vip_growth_speed"`
		VipGrowthTotal int64  `json:"vip_growth_total"`
		VipLevel       string `json:"vip_level"`
		Yes            bool   `json:"yes"`
	} `json:"data"`
	Message string `json:"msg"`
	Retcode int    `json:"retcode"`
	Status  string `json:"status"`
	Wording string `json:"wording"`
}

func comm(req string) response {
	raw, _ := http.Get(req)
	var tmp response
	_ = json.NewDecoder(raw.Body).Decode(&tmp)
	return tmp
}

// SendPrivateMsg ??????????????????
func SendPrivateMsg(Address string, UserId int64, GroupId int64, Message string, AutoEscape bool) response {
	var tmp = Address +
		"/send_private_msg?" +
		"user_id=" + strconv.FormatInt(UserId, 10)
	if GroupId != 0 {
		tmp = tmp +
			"&group_id=" + strconv.FormatInt(GroupId, 10)
	}
	tmp = tmp +
		"&message=" + Message
	if AutoEscape != false {
		tmp = tmp +
			"&auto_escape=" + "true"
	}
	return comm(tmp)
}

// SendGroupMsg ???????????????
func SendGroupMsg(Address string, GroupId int64, Message string, AutoEscape bool) response {
	var tmp = Address +
		"/send_group_msg?" +
		"user_id=" + strconv.FormatInt(GroupId, 10) +
		"&message=" + Message
	if AutoEscape != false {
		tmp = tmp +
			"&auto_escape=true"
	}
	return comm(tmp)
}

// SendGroupForwardMsg ?????????????????? ( ??? )
func SendGroupForwardMsg(Address string, GroupId int64, Messages string, AutoEscape bool) response {
	var tmp = Address +
		"/send_group_forward_msg?" +
		"group_id=" + strconv.FormatInt(GroupId, 10) +
		"&messages=" + Messages
	if AutoEscape != false {
		tmp = tmp +
			"&auto_escape=true"
	}
	return comm(tmp)
}

// SendMsg ????????????
func SendMsg(Address string, MessageType string, UserId int64, GroupId int64, Message string, AutoEscape bool) response {
	var tmp = Address +
		"/send_group_forward_msg?"
	if UserId != 0 || MessageType == "private" {
		tmp = tmp +
			"user_id=" + strconv.FormatInt(UserId, 10) +
			"&message=" + Message +
			"&auto_escape=" + strconv.FormatBool(AutoEscape)
	}
	if GroupId != 0 || MessageType == "group" {
		tmp = tmp +
			"group_id=" + strconv.FormatInt(GroupId, 10) +
			"&message=" + Message +
			"&auto_escape=" + strconv.FormatBool(AutoEscape)
	}
	return comm(tmp)
}

// DeleteMsg ????????????
func DeleteMsg(Address string, MessageId int32) response {
	var tmp = Address +
		"/delete_msg?" +
		"message_id=" + fmt.Sprint(MessageId)
	return comm(tmp)
}

// GetMsg ????????????
func GetMsg(Address string, MessageId int32) response {
	var tmp = Address +
		"/get_msg?" +
		"message_id=" + fmt.Sprint(MessageId)
	return comm(tmp)
}

// GetForwardMsg ????????????????????????
func GetForwardMsg(Address string, MessageId int32) response {
	var tmp = Address +
		"/get_forward_msg?" +
		"message_id=" + fmt.Sprint(MessageId)
	return comm(tmp)
}

// GetImage ??????????????????
func GetImage(Address string, file string) response {
	var tmp = Address +
		"/get_image?" +
		"file=" + file
	return comm(tmp)
}

// SetGroupKick ????????????
func SetGroupKick(Address string, GroupId int64, UserId int64, RejectAddRequest bool) response {
	var tmp = Address +
		"/set_group_kick?" +
		"group_id=" + strconv.FormatInt(GroupId, 10) +
		"&user_id=" + strconv.FormatInt(UserId, 10)
	if RejectAddRequest == true {
		tmp = tmp +
			"&reject_add_request=true"
	}
	return comm(tmp)
}

// SetGroupBan ??????????????????
func SetGroupBan(Address string, GroupId int64, UserId int64, Duration int) response {
	var tmp = Address +
		"/set_group_ban?" +
		"group_id=" + strconv.FormatInt(UserId, 10) +
		"&user_id=" + strconv.FormatInt(UserId, 10) +
		"&duration=" + strconv.Itoa(Duration)
	return comm(tmp)
}

// SetGroupAnonymousBan ????????????????????????
func SetGroupAnonymousBan(Address string, GroupId int64, Anonymous anonymous, Flag string, Duration int) response {
	var tmp = Address +
		"/set_group_anonymous_ban?" +
		"group_id=" + strconv.FormatInt(GroupId, 10)
	if Anonymous.Flag != "" && Anonymous.ID != 0 && Anonymous.Name != "" {
		tmp2, _ := json.Marshal(Anonymous)
		tmp = tmp +
			"&anonymous=" + string(tmp2)
	}
	tmp = tmp +
		"&flag=" + Flag +
		"&duration=" + strconv.Itoa(Duration)
	return comm(tmp)
}

// SetGroupWholeBan ??????????????????
func SetGroupWholeBan(Address string, GroupId int64, Enable bool) response {
	var tmp = Address +
		"/set_group_whole_ban?" +
		"group_id=" + strconv.FormatInt(GroupId, 10) +
		"&enable=" + strconv.FormatBool(Enable)
	return comm(tmp)
}

// SetGroupAdmin ?????????????????????
func SetGroupAdmin(Address string, GroupId int64, UserId int64, Enable bool) response {
	var tmp = Address +
		"/set_group_admin?" +
		"group_id=" + strconv.FormatInt(GroupId, 10) +
		"&user_id=" + strconv.FormatInt(UserId, 10) +
		"&enable=" + strconv.FormatBool(Enable)
	return comm(tmp)
}

// SetGroupAnonymous ????????????
func SetGroupAnonymous(Address string, GroupId int64, Enable bool) response {
	var tmp = Address +
		"/set_group_anonymous?" +
		"group_id=" + strconv.FormatInt(GroupId, 10) +
		"&enable=" + strconv.FormatBool(Enable)
	return comm(tmp)
}

// SetGroupCard ??????????????? ( ????????? )
func SetGroupCard(Address string, GroupId int64, UserId int64, card string) response {
	var tmp = Address +
		"/set_group_card?" +
		"group_id=" + strconv.FormatInt(GroupId, 10) +
		"&user_id=" + strconv.FormatInt(UserId, 10) +
		"&card=" + card
	return comm(tmp)
}

// SetGroupName ????????????
func SetGroupName(Address string, GroupId int64, GroupName string) response {
	var tmp = Address +
		"/set_group_card?" +
		"group_id=" + strconv.FormatInt(GroupId, 10) +
		"&group_name=" + GroupName
	return comm(tmp)
}

// SetGroupLeave ????????????
func SetGroupLeave(Address string, GroupId int64, IsDismiss bool) response {
	var tmp = Address +
		"/set_group_leave?" +
		"group_id=" + strconv.FormatInt(GroupId, 10) +
		"&is_dismiss=" + strconv.FormatBool(IsDismiss)
	return comm(tmp)
}

// SetGroupSpecialTitle ????????????????????????
func SetGroupSpecialTitle(Address string, GroupId int64, UserId int64, SpecialTitle string, Duration int) response {
	var tmp = Address +
		"/set_group_special_title?" +
		"group_id=" + strconv.FormatInt(GroupId, 10) +
		"&user_id=" + strconv.FormatInt(UserId, 10) +
		"&special_title=" + SpecialTitle +
		"&duration=" + strconv.Itoa(Duration)
	return comm(tmp)
}

// SetFriendAddRequest ?????????????????????
func SetFriendAddRequest(Address string, Flag string, Approve bool, Remark string) response {
	var tmp = Address +
		"/set_friend_add_request?" +
		"flag=" + Flag +
		"&approve=" + strconv.FormatBool(Approve)
	if Approve == true && Remark != "" {
		tmp = tmp +
			"&remark=" + Remark
	}
	return comm(tmp)
}

// SetGroupAddRequest ???????????????????????????
func SetGroupAddRequest(Address string, Flag string, SubType string, Approve bool, Reason string) response {
	var tmp = Address +
		"/set_group_add_request?" +
		"flag=" + Flag +
		"&type=" + SubType +
		"&approve=" + strconv.FormatBool(Approve)
	if Approve == false && Reason != "" {
		tmp = tmp +
			"&reason=" + Reason
	}
	return comm(tmp)
}

// GetLoginInfo ?????????????????????
func GetLoginInfo(Address string) response {
	var tmp = Address +
		"/get_login_info"
	return comm(tmp)
}

// QidianGetAccountInfo ????????????????????????
func QidianGetAccountInfo(Address string) response {
	var tmp = Address +
		"/qidian_get_account_info"
	return comm(tmp)
}

// GetStrangerInfo ?????????????????????
func GetStrangerInfo(Address string, UserId int64, NoCache bool) response {
	var tmp = Address +
		"/get_stranger_info?" +
		"user_id=" + strconv.FormatInt(UserId, 10)
	if NoCache == true {
		tmp = tmp +
			"&no_cache=" + strconv.FormatBool(NoCache)
	}
	return comm(tmp)
}

// GetFriendList ??????????????????
func GetFriendList(Address string) _response {
	var tmp = Address +
		"/get_friend_list"
	raw, _ := http.Get(tmp)
	var tmp2 _response
	_ = json.NewDecoder(raw.Body).Decode(&tmp)
	return tmp2
}

// DeleteFriend ????????????
func DeleteFriend(Address string, FriendId int64) response {
	var tmp = Address +
		"/delete_friend?" +
		"friend_id=" + strconv.FormatInt(FriendId, 10)
	return comm(tmp)
}

// GetGroupInfo ???????????????
func GetGroupInfo(Address string, GroupId int64, NoCache bool) response {
	var tmp = Address +
		"/get_group_info?" +
		"group_id=" + strconv.FormatInt(GroupId, 10)
	if NoCache == true {
		tmp = tmp +
			"&no_cache=true"
	}
	return comm(tmp)
}

// GetGroupList ???????????????
func GetGroupList(Address string) _response {
	var tmp = Address +
		"/get_group_list"
	raw, _ := http.Get(tmp)
	var tmp2 _response
	_ = json.NewDecoder(raw.Body).Decode(&tmp2)
	return tmp2
}

// GetGroupMemberInfo ?????????????????????
func GetGroupMemberInfo(Address string, GroupId int64, UserId int64, NoCache bool) response {
	var tmp = Address +
		"/get_group_member_info?" +
		"group_id=" + strconv.FormatInt(UserId, 10) +
		"&user_id=" + strconv.FormatInt(GroupId, 10)
	if NoCache == true {
		tmp = tmp +
			"&no_cache=" + strconv.FormatBool(NoCache)
	}
	return comm(tmp)
}

// GetGroupMemberList ?????????????????????
func GetGroupMemberList(Address string, GroupId int64) _response {
	var tmp = Address +
		"/get_group_member_list?" +
		"group_id=" + strconv.FormatInt(GroupId, 10)
	raw, _ := http.Get(tmp)
	var tmp2 _response
	_ = json.NewDecoder(raw.Body).Decode(&tmp)
	return tmp2
}

// GetGroupHonorInfo ?????????????????????
func GetGroupHonorInfo(Address string, GroupId int64, Type string) {

}

// CanSendImage ??????????????????????????????
func CanSendImage(Address string) response {
	var tmp = Address +
		"/can_send_image"
	return comm(tmp)
}

// CanSendRecord ??????????????????????????????
func CanSendRecord(Address string) response {
	var tmp = Address +
		"/can_send_record"
	return comm(tmp)
}

// GetVersionInfo ??????????????????
func GetVersionInfo(Address string) response {
	var tmp = Address +
		"/get_version_info"
	return comm(tmp)
}

// SetRestart ?????? go-cqhttp
func SetRestart(Address string) response {
	var tmp = Address +
		"/set_restart"
	return comm(tmp)
}

// SetGroupPortrait ???????????????
func SetGroupPortrait(Address string, GroupId int64, File string, Cache int) response {
	var tmp = Address +
		"/set_group_portrait?" +
		"group_id" + strconv.FormatInt(GroupId, 10) +
		"&file=" + File
	if Cache == 0 {
		tmp = tmp +
			"&cache=0"
	}
	return comm(tmp)
}

// GetWordSlices ?????????????????? ( ?????? API )
func GetWordSlice(Address string, Content string) response {
	var tmp = Address +
		"/.get_word_slices?" +
		"content=" + Content
	return comm(tmp)
}

// OCRImage ?????? OCR
func OcrImage(Address string, Image string) response {
	var tmp = Address +
		"/ocr_image?" +
		"image=" + Image
	return comm(tmp)
}

// GetGroupSystemMessage ?????????????????????
func GetGroupSystemMessage(Address string) response {
	var tmp = Address +
		"/get_group_system_msg"
	return comm(tmp)
}

// UploadGroupFile ???????????????
func UploadGroupFile(Address string, GroupId int64, File string, Name string, Folder string) response {
	var tmp = Address +
		"/upload_group_file?" +
		"group_id=" + strconv.FormatInt(GroupId, 10) +
		"&file=" + File +
		"&name=" + Name +
		"&folder=" + Folder
	return comm(tmp)
}

// GetGroupFileSystemInfo ???????????????????????????
func GetGroupFileSystemInfo(Address string, GroupId int64) response {
	var tmp = Address +
		"/get_group_file_system_info?" +
		"group_id=" + strconv.FormatInt(GroupId, 10)
	return comm(tmp)
}

// GetGroupRootFiles ??????????????????????????????
func GetGroupRootFiles(Address string, GroupId int64) response {
	var tmp = Address +
		"/get_group_root_files?" +
		"group_id=" + strconv.FormatInt(GroupId, 10)
	return comm(tmp)
}

// GetGroupFilesByFolder ??????????????????????????????
func GetGroupFilesByFolder(Address string, GroupId int64, FolderId string) response {
	var tmp = Address +
		"/get_group_files_by_folder?" +
		"group_id=" + strconv.FormatInt(GroupId, 10) +
		"&folder_id=" + FolderId
	return comm(tmp)
}

// GetGroupFilesUrl ???????????????????????????
func GetGroupFilesUrl(Address string, GroupId int64, FileId string, Busid int32) response {
	var tmp = Address +
		"/get_group_file_url?" +
		"group_id=" + strconv.FormatInt(GroupId, 10) +
		"&file_id=" + FileId +
		"&busid=" + fmt.Sprint(Busid)
	return comm(tmp)
}

// GetStatus ????????????
func GetStatus(Address string) response {
	var tmp = Address +
		"/get_status"
	return comm(tmp)
}

// GetGroupAtAllRemain ????????? @???????????? ????????????
func GetGroupAtAllRemain(Address string, GroupId int64) response {
	var tmp = Address +
		"/get_group_at_all_remain?" +
		"group_id=" + strconv.FormatInt(GroupId, 10)
	return comm(tmp)
}

// HandleQuickOperation ??????????????????????????? ( ?????? API )
func HandleQuickOperation(Address string, Context string, Operation string) response {
	var tmp = Address +
		"/.handle_quick_operation?" +
		"context=" + Context +
		"&operation=" + Operation
	return comm(tmp)
}

// GetVipInfo ??????VIP??????
func GetVipInfo(Address string, UserId int64) response {
	var tmp = Address +
		"/_get_vip_info?" +
		"user_id=" + strconv.FormatInt(UserId, 10)
	return comm(tmp)
}

// SendGroupNotice ???????????????
func SendGroupNotice(Address string, GroupId int64, Content string, Image string) response {
	var tmp = Address +
		"/_send_group_notice?" +
		"group_id=" + strconv.FormatInt(GroupId, 10) +
		"&content=" + Content
	if Image != "" {
		tmp = tmp +
			"&image=" + Image
	}
	return comm(tmp)
}

// ReloadEventFilter ?????????????????????
func ReloadEventFilter(Address string, File string) response {
	var tmp = Address +
		"/reload_event_filter?" +
		"file=" + File
	return comm(tmp)
}

// DownloadFile ???????????????????????????
func DownloadFile(Address string, Url string, ThreadCount int32, Headers string) response {
	var tmp = Address +
		"/download_file?" +
		"url=" + Url
	if ThreadCount != 0 {
		tmp = tmp +
			"&thread_count=" + fmt.Sprint(ThreadCount)
	}
	if Headers != "" {
		tmp = tmp +
			"&headers=" + Headers
	}
	return comm(tmp)
}

// GetOnlineClients ???????????????????????????????????????
func GetOnlineClients(Address string, NoCache bool) response {
	var tmp = Address +
		"/get_online_clients?" +
		"no_cache=" + strconv.FormatBool(NoCache)
	return comm(tmp)
}

// GetGroupMsgHistory ???????????????????????????
func GetGroupMsgHistory(Address string, GroupId int64, MessageSeq int64) response {
	var tmp = Address +
		"/get_group_msg_history?" +
		"group_id=" + strconv.FormatInt(GroupId, 10)
	if MessageSeq != 0 {
		tmp = tmp +
			"&message_seq=" + strconv.FormatInt(MessageSeq, 10)
	}
	return comm(tmp)
}

// SetEssenceMsg ??????????????????
func SetEssenceMsg(Address string, MessageId int32) response {
	var tmp = Address +
		"/set_essence_msg?" +
		"message_id=" + fmt.Sprint(MessageId)
	return comm(tmp)
}

// DeleteEssenceMsg ??????????????????
func DeleteEssenceMsg(Address string, MessageId int32) response {
	var tmp = Address +
		"/delete_essence_msg?" +
		"message_id=" + fmt.Sprint(MessageId)
	return comm(tmp)
}

// GetEssenceMsgList ????????????????????????
func GetEssenceMsgList(Address string, GroupId int64) _response {
	var tmp = Address +
		"/get_essence_msg_list?" +
		"group_id=" + strconv.FormatInt(GroupId, 10)
	raw, _ := http.Get(tmp)
	var tmp2 _response
	_ = json.NewDecoder(raw.Body).Decode(&tmp2)
	return tmp2
}

// CheckUrlSafely ?????????????????????
func CheckUrlSafely(Address string, Url string) response {
	var tmp = Address +
		"/check_url_safely?" +
		"url=" + Url
	return comm(tmp)
}

// GetModelShow ??????????????????
func GetModelShow(Address string, Model string) response {
	var tmp = Address +
		"/_get_model_show?" +
		"model=" + Model
	return comm(tmp)
}

// SetModelShow ??????????????????
func SetModelShow(Address string, Model string, ModelShow string) response {
	var tmp = Address +
		"/_set_model_show?" +
		"model=" + Model +
		"&model_show=" + ModelShow
	return comm(tmp)
}
