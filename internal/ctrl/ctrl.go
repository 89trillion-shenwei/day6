package ctrl

import (
	"day7/internal/Service"
	mes "day7/internal/message"
	"day7/internal/model"
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

// Connect 连接请求
func Connect(username, url string) {
	if !Service.Check(model.Left.Text, username) {
		fmt.Println("重名")
		return
	}
	model.ConnectStatus.SetText("Connection status: connected")
	Header := http.Header{}
	Header.Add("username", username)
	ws, _, err := websocket.DefaultDialer.Dial(url, Header)
	if err != nil {
		fmt.Println("websocket连接失败，" + err.Error())
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
	//开启发送和接收
	go client.ReceiveMsg()
	go client.SendMsg()
}

func DisConnect() {
	model.ConnectStatus.SetText("Connection status: disconnected")
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

	client.Send <- Service.Struct2proto(exitMsg3)
	client.Send <- Service.Struct2proto(exitMsg2)
	client.Send <- Service.Struct2proto(exitMsg1)
}

func Send() {
	if model.ConnectStatus.Text == "Connection status: disconnected" {
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
