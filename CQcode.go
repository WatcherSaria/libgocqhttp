package robot

import (
	"fmt"
	"strconv"
)

func face(id int) string {
	return "[CQ:face,id=" + strconv.Itoa(id) + "]"
}

func record(file string, magic bool, cache bool, proxy bool, timeout int) string {
	var temp string = "[CQ:record"
	temp = temp + ",file=" + file
	if magic == true {
		temp = temp + ",magic=1"
	}
	if cache == false {
		temp = temp + ",cache=0"
	}
	if proxy == false {
		temp = temp + ",proxy=0"
	}
	if timeout != 0 {
		temp = temp + ",timeout=" +strconv.Itoa(timeout)
	}
	return temp + "]"
}

func video(file string, cover string, c int) string {
	var temp string = "[CQ:video"
	temp = temp + ",file=" + file
	if cover != "" {
		temp = temp + ",cover=" + cover
	}
	if c != 0 {
		temp = temp + ",c=" + strconv.Itoa(c)
	}
	return temp + "]"
}

func at(qq int64, name string) string {
	var temp string = "[CQ:at"
	temp = temp + ",qq=" + strconv.FormatInt(qq, 10)
	if name != "" {
		temp = temp + ",name=" + name
	}
	return temp + "]"
}

//func rps()  {
//该 CQcode 暂未被 go-cqhttp 支持, 您可以向go-cqhttp项目提交 pr 以使该 CQcode 被支持:https://github.com/Mrs4s/go-cqhttp/compare
//}

//func dice()  {
//该 CQcode 暂未被 go-cqhttp 支持, 您可以向go-cqhttp项目提交 pr 以使该 CQcode 被支持:https://github.com/Mrs4s/go-cqhttp/compare
//}

//func shake()  {
//该 CQcode 暂未被 go-cqhttp 支持, 您可以向go-cqhttp项目提交 pr 以使该 CQcode 被支持:https://github.com/Mrs4s/go-cqhttp/compare
//}

//func anonymous()  {
//该 CQcode 暂未被 go-cqhttp 支持, 您可以向go-cqhttp项目提交 pr 以使该 CQcode 被支持:https://github.com/Mrs4s/go-cqhttp/compare
//}

//func contact() {
//该 CQcode 暂未被 go-cqhttp 支持, 您可以向go-cqhttp项目提交 pr 以使该 CQcode 被支持:https://github.com/Mrs4s/go-cqhttp/compare
//}

//func location() {
//该 CQcode 暂未被 go-cqhttp 支持, 您可以向go-cqhttp项目提交 pr 以使该 CQcode 被支持:https://github.com/Mrs4s/go-cqhttp/compare
//}

func music(_type string, id int64, url string, audio string, title string, content string, image string) string {
	var temp string = "[CQ:music"
	if _type == "qq" || _type == "163" || _type == "xm" {
		temp = temp + ",type=" + _type
		temp = temp + ",id=" + strconv.FormatInt(id, 10)
	} else {
		temp = temp + ",type=custom"
		temp = temp + ",url=" + url
		audio = temp + ",audio=" + audio
		title = temp + ",title=" + title
		if content == "" {
			temp = temp + ",content=" + content
		}
		if image == "" {
			temp = temp + ",image=" + image
		}
	}
	return temp + "]"
}

func image(file string, _type string, subType string, url string, cache bool, id int, c int) string {
	var temp string = "[CQ:image"
	temp = temp + ",file=" + file
	temp = temp + ",url=" + url
	if _type != "" {
		temp = temp + ",type=" + _type
	}
	if subType != "" {
		temp = temp + ",subType=" + subType
	}
	if cache == false {
		temp = temp + ",cache=0"
	}
	if id != 0 {
		temp = temp + ",id=" + strconv.Itoa(id)
	}
	if c != 0 {
		temp = temp + ",c=" + strconv.Itoa(c)
	}
	return temp + "]"
}

func reply(id int, text string, qq int64, time int64, seq int64) string {
	var temp string = "[CQ:reply"
	temp = temp + ",id=" + strconv.Itoa(id)
	if text != "" {
		temp = temp + ",text=" + text
	}
	if qq != 0 {
		temp = temp + ",qq=" + strconv.FormatInt(qq, 10)
	}
	if time != 0 {
		temp = temp + ",time=" + strconv.FormatInt(time, 10)
	}
	if seq != 0 {
		temp = temp + ",seq=" + strconv.FormatInt(seq, 10)
	}
	return temp + "]"
}

func redbag(title string) string {
	return "CQ:redbag,title=" + title + "]"
}

func poke(qq int64) string {
	return "CQ:poke,qq=" + strconv.FormatInt(qq, 10) + "]"
}

func gift(qq int64, id int) string {
	return "CQ:gift,qq=" + strconv.FormatInt(qq, 10) + ",id=" + strconv.Itoa(id) + "]"
}

func forward(id string) string {
	return "[CQ:forward,id=" + id + "]"
}

//func node(id int32, name string, uin int64, content string, seq string) string {
//因无具体实例所以摸了
//}

func xml(data string, resid int32) string {
	return "[CQ:xml,data=" + data + "]"
}

func _json(data string, resid int32) string {
	var temp string = "[CQ:json"
	temp = temp + ",data=" + data
	if resid != 0 {
		temp = temp + ",resid=" + fmt.Sprint(resid)
	}
	return temp + "]"
}

func cardimage(file string, minwidth int64, minheight int64, maxwidth int64, maxheight int64, source string, icon string) string {
	var temp string = "[CQ:cardimage"
	temp = temp + ",file=" + file
	if minwidth != 0 {
		temp = temp + ",minwidth=" + strconv.FormatInt(minwidth, 10)
	}
	if minheight != 0 {
		temp = temp + ",minheight=" + strconv.FormatInt(minheight, 10)
	}
	if maxwidth != 0 {
		temp = temp + ",maxwidth=" + strconv.FormatInt(maxwidth, 10)
	}
	if maxheight != 0 {
		temp = temp + ",maxheight=" + strconv.FormatInt(maxheight, 10)
	}
	if source != "" {
		temp = temp + ",source=" + source
	}
	if icon != "" {
		temp = temp + ",icon=" + icon
	}
	return temp + "]"
}

func tts(text string) string {
	return "[CQ:tts,text=" + text + "]"
}