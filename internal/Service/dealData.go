package Service

import (
	mes "day7/internal/message"
	"fmt"

	"github.com/golang/protobuf/proto"
	"github.com/gorilla/websocket"
)

// Struct2proto 解析为proto格式
func Struct2proto(msg mes.Msg) []byte {
	byts, err := proto.Marshal(&msg)
	if err != nil {
		fmt.Println("解析为proto格式失败")
	}
	return byts
}

// Proto2Struct proto转结构体
func Proto2Struct(byts []byte) mes.Msg {
	msg := mes.Msg{}
	err := proto.Unmarshal(byts, &msg)
	if err != nil {
		return mes.Msg{}
		fmt.Println("解析为结构体失败")
	}
	return msg

}

// UploadData 上传数据
func UploadData(msg mes.Msg, conn websocket.Conn) {
	err := conn.WriteMessage(websocket.TextMessage, Struct2proto(msg))
	if err != nil {
		fmt.Println("上传登录信息失败，" + err.Error())
	}
	fmt.Println(msg)
}

// DownloadData 获取数据
func DownloadData(conn *websocket.Conn) (mes.Msg, error) {
	//var byts = make([]byte, 1024)
	_, byts, err := conn.ReadMessage()
	if err != nil {
		fmt.Println("读取信息失败" + err.Error())
		return mes.Msg{}, err
	}
	msg := &mes.Msg{}
	err1 := proto.Unmarshal(byts, msg)
	if err1 != nil {
		fmt.Println(err1.Error())
	}
	return *msg, nil
}
