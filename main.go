package main

import (
	"image/color"
	"qingmusic/theme"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

func main() {
	a := app.New()
	w := a.NewWindow("轻音乐")

	a.Settings().SetTheme(theme.CustomTheme)

	text1 := canvas.NewText("我的收藏", color.Black)
	text2 := canvas.NewText("本地下载", color.Black)
	text3 := canvas.NewText("最近播放", color.Black)
	text4 := canvas.NewText("全部音乐", color.Black)

	text1.TextSize = 20
	text2.TextSize = 20
	text3.TextSize = 20
	text4.TextSize = 20

	// 创建绿色背景矩形
	topBg := canvas.NewRectangle(color.NRGBA{G: 255, A: 255})
	topBg.SetMinSize(fyne.NewSize(0, 100))

	rightBg := canvas.NewRectangle(color.NRGBA{B: 255, A: 255})
	rightBg.SetMinSize(fyne.NewSize(100, 100))

	topContent := container.New(layout.NewStackLayout(), topBg)

	leftContent := container.New(layout.NewVBoxLayout(), text1, text2, text3, text4)

	rightContent := container.New(layout.NewStackLayout(), rightBg)

	// 底部布局
	bottomContent := container.NewBorder(
		nil, nil, leftContent, nil,
		rightContent,
	)

	w.SetContent(container.New(layout.NewVBoxLayout(), topContent, bottomContent))

	w.Resize(fyne.Size{Width: 1750, Height: 1150})

	w.ShowAndRun()
}
