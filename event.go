package libgocqhttp

type anonymous struct {
	Flag string `json:"flag"`
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type file struct {
	Busid int64  `json:"busid"`
	ID    string `json:"id"`
	Name  string `json:"name"`
	Size  int64  `json:"size"`
	URL   string `json:"url"`
}

type sender struct {
	Age      int32  `json:"age"`
	Area     string `json:"area"`
	Level    string `json:"level"`
	Nickname string `json:"nickname"`
	Sex      string `json:"sex"`
	UserID   int64  `json:"user_id"`
}

type Event struct {
	Anonymous   anonymous `json:"anonymous"`
	CardNew     string    `json:"card_new"`
	CardOld     string    `json:"card_old"`
	Client      string    `json:"client"`
	Comment     string    `json:"comment"`
	Duration    int64     `json:"duration"`
	File        file      `json:"file"`
	Flag        string    `json:"flag"`
	Font        int32     `json:"font"`
	GroupId     int64     `json:"group_id"`
	Message     string    `json:"message"`
	MessageID   int32     `json:"message_id"`
	MessageSeq  int32     `json:"message_seq"`
	MessageType string    `json:"message_type"`
	NoticeType  string    `json:"notice_type"`
	Online      bool      `json:"online"`
	OperatorID  int64     `json:"operator_id"`
	PostType    string    `json:"post_type"`
	RawMessage  string    `json:"raw_message"`
	RequestType string    `json:"request_type"`
	SelfID      int64     `json:"self_id"`
	Sender      sender    `json:"sender"`
	SenderID    int64     `json:"sender_id"`
	SubType     string    `json:"sub_type"`
	TargetID    int64     `json:"target_id"`
	TempSource  int       `json:"temp_source"`
	Time        int64     `json:"time"`
	Title       string    `json:"title"`
	UserID      int64     `json:"user_id"`
}
