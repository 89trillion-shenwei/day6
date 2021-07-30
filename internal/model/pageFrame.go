package model

import (
	"fyne.io/fyne/widget"
)

//用户名输入框
var UserEntry *widget.Entry

//服务地址输入框
var ServerEntry *widget.Entry

//发送信息输入框
var MessageEntry *widget.Entry

//连接状态
var ConnectStatus *widget.Label

//连接按钮
var ConnectButton *widget.Button

//断开按钮
var DisConnectButton *widget.Button

//发送按钮
var SendButton *widget.Button

//用户列表
var Left *widget.Label

//消息列表
var Right *widget.Label
