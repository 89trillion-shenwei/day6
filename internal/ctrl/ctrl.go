package ctrl

import (
	"day7/internal"
	"day7/internal/Service"
	mes "day7/internal/message"
	"day7/internal/model"
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

// Connect 连接请求
func Connect(username, url string) {
	if model.UserEntry.Text == "" {
		println("用户名为空")
		internal.ReturnErr(internal.NoNameError("username si empty"))
		return
	}
	if model.ServerEntry.Text == "" {
		println("地址为空")
		internal.ReturnErr(internal.NoAdressError("websocket is empty"))
	}
	if !Service.Check(model.Left.Text, username) {
		internal.ReturnErr(internal.ServerError("the name is the same as others"))
		return
	}
	if model.ConnectStatus.Text == "Connection status: connected" {
		internal.ReturnErr(internal.ServerError("user has connected"))
		return
	}
	Header := http.Header{}
	Header.Add("username", username)
	ws, _, err := websocket.DefaultDialer.Dial(url, Header)
	if err != nil {
		fmt.Println("websocket连接失败，" + err.Error())
		internal.ReturnErr(internal.ServerError("websocket connect failed"))
		return
	}
	//发送广播
	msg1 := mes.Msg{
		UserName: username,
		Msg:      username + "login",
		MsgType:  "talk",
		List:     nil,
	}
	//发送userlist请求
	msg2 := mes.Msg{
		UserName: username,
		Msg:      "",
		MsgType:  "userlist",
		List:     nil,
	}
	Service.UploadData(msg1, *ws)
	Service.UploadData(msg2, *ws)
	client := model.GetClient()
	client.Conn = ws
	client.UserName = username
	model.ConnectStatus.SetText("Connection status: connected")
	model.Right.SetText("")
	//开启发送和接收
	go client.ReceiveMsg()
	go client.SendMsg()
}

func DisConnect() {
	if model.ConnectStatus.Text == "Connection status: disconnected" {
		return
	}
	client := model.GetClient()
	exitMsg1 := mes.Msg{
		UserName: client.UserName,
		MsgType:  "exit",
	}
	exitMsg2 := mes.Msg{
		UserName: client.UserName,
		Msg:      client.UserName + " exit",
		MsgType:  "talk",
	}
	exitMsg3 := mes.Msg{
		UserName: client.UserName,
		MsgType:  "userlist",
	}
	model.ConnectStatus.SetText("Connection status: disconnected")
	client.Send <- Service.Struct2proto(exitMsg3)
	client.Send <- Service.Struct2proto(exitMsg2)
	client.Send <- Service.Struct2proto(exitMsg1)
}

func Send() {
	if model.ConnectStatus.Text == "Connection status: disconnected" {
		model.Right.SetText("")
		internal.ReturnErr(internal.ServerError("No user connect"))
		return
	}
	str := model.MessageEntry.Text
	fmt.Println(str)
	client := model.GetClient()
	sendMsg := mes.Msg{
		UserName: client.UserName,
		Msg:      str,
		MsgType:  "talk",
		List:     nil,
	}
	client.Send <- Service.Struct2proto(sendMsg)
}

func ClearEntry() {
	model.MessageEntry.SetText("") //输入框清空
}
