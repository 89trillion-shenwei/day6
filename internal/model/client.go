package model

import (
	"day7/internal/Service"
	"fmt"
	"strings"
	"time"

	"github.com/gorilla/websocket"
)

type Client struct {
	Conn     *websocket.Conn //连接
	UserName string          //用户名
	Send     chan []byte     //发送管道
	Receive  chan []byte     //接收管道

}

var client = &Client{
	Conn:     nil,
	UserName: "",
	Send:     make(chan []byte, 0),
	Receive:  make(chan []byte, 0),
}

const (
	writeWait      = 360 * time.Second
	pongWait       = 360 * time.Second
	pingPeriod     = (pongWait * 9) / 10
	maxMessageSize = 512
)

func GetClient() *Client {
	return client
}

// ReceiveMsg 处理接收到的信息
func (client *Client) ReceiveMsg() {
	client.Conn.SetPongHandler(func(string) error {
		fmt.Println("心跳检测" + time.Now().Format("2006-01-02 15:04:05"))
		return nil
	})
	for {
		if client.Conn == nil {
			break
		}
		//得到数据
		msg, err := Service.DownloadData(client.Conn)
		if err != nil {
			fmt.Println(err.Error())
			break
		}
		switch msg.MsgType {
		case "talk":
			str := "\n" + msg.UserName + "send:" + msg.Msg
			Right.SetText(Right.Text + str)
		case "userlist":
			str := ""
			strs := strings.Split(msg.Msg, " ")
			for i := 0; i < len(strs); i++ {
				str += "\n"
				str += strs[i]
			}
			//更新userlist
			msg.Msg = str
			client.Send <- Service.Struct2proto(msg)
			Left.SetText(str)
		case "exit":
			if msg.UserName == client.UserName {
				Left.SetText("")
				Right.SetText("")
			}
			break
		}
	}
}

// SendMsg 发送消息
func (client *Client) SendMsg() {
	ticker := time.NewTicker(pingPeriod)
	for {
		if client.Conn == nil {
			break
		}
		select {
		case msg := <-client.Send:
			client.Conn.WriteMessage(websocket.TextMessage, msg)
			/*Msg:=Service.Proto2Struct(msg)
			if Msg.MsgType=="exit"{
				//close(client.Send)
				//close(client.Receive)
				client.Conn.Close()
				client.Conn=nil
			}*/
		case <-ticker.C:
			client.Conn.SetWriteDeadline(time.Now().Add(writeWait))
			err := client.Conn.WriteMessage(websocket.PingMessage, nil)
			if err != nil {
				return
			}
		}

	}
}
