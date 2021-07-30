package view

import (
	ctrl2 "day7/internal/ctrl"
	"day7/internal/model"
	"fmt"
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/container"
	"fyne.io/fyne/widget"
)

func Init() {
	//应用程序对象
	myApp := app.New()
	//界面对象
	myWin := myApp.NewWindow("Chat")
	//连接状态
	model.ConnectStatus = widget.NewLabel("Connection status: disconnected")
	//用户名输入框
	model.UserEntry = widget.NewEntry()
	model.UserEntry.SetPlaceHolder("username")
	model.UserEntry.Resize(fyne.NewSize(200, 100))

	//服务地址输入框
	model.ServerEntry = widget.NewEntry()
	model.ServerEntry.SetPlaceHolder("url")
	model.ServerEntry.Resize(fyne.NewSize(200, 100))
	//发送信息输入框
	model.MessageEntry = widget.NewEntry()
	model.MessageEntry.Resize(fyne.NewSize(600, 200))

	//连接按钮
	model.ConnectButton = widget.NewButton("connect", func() {
		if model.UserEntry.Text == "" || model.ServerEntry.Text == "" {
			println("有参数为空")
			return
		}
		println(model.UserEntry.Text)
		println(model.ServerEntry.Text)
		ctrl2.Connect(model.UserEntry.Text, model.ServerEntry.Text)
	})
	//断开按钮
	model.DisConnectButton = widget.NewButton("disconnect", func() {
		fmt.Println(model.UserEntry.Text + "断开")
		ctrl2.DisConnect()
	})
	model.SendButton = widget.NewButton("send", func() {
		fmt.Println(model.UserEntry.Text + " send:" + model.MessageEntry.Text)
		ctrl2.Send()
		model.MessageEntry.SetText("") //输入框清空

	})
	model.SendButton.Resize(fyne.NewSize(100, 100))

	//user和server表单
	screen := widget.NewForm(
		&widget.FormItem{Text: "user:", Widget: model.UserEntry},
		&widget.FormItem{Text: "server:", Widget: model.ServerEntry},
	)
	//上层用户名盒子
	Box1 := widget.NewHBox(screen, model.ConnectButton, model.DisConnectButton, model.ConnectStatus)
	//中左的用户列表布局
	model.Left = widget.NewLabel("")
	//中右的聊天室信息布局
	model.Right = widget.NewLabel("")
	left := widget.NewCard("userlist:", "", model.Left)
	right := widget.NewCard("", "", model.Right)

	//输入框盒子
	Box3 := widget.NewHBox(model.MessageEntry, model.SendButton)
	Box3.Resize(fyne.NewSize(800, 200))

	Content := container.NewBorder(Box1, Box3, left, nil, right)
	myWin.Resize(fyne.NewSize(800, 800))
	myWin.SetContent(Content)
	myWin.ShowAndRun()
}
