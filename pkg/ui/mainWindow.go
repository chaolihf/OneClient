package ui

import (
	_ "embed"
	"fmt"
	"os"
	"strings"
	"time"

	"com.chinatelecom.oneops.client/pkg/robot"
	"github.com/chaolihf/goey"
	"github.com/chaolihf/goey/base"
	"github.com/chaolihf/goey/loop"
	"github.com/chaolihf/goey/windows"
	"github.com/getlantern/systray"
	"github.com/lxn/win"
	"go.uber.org/zap"
)

//go:embed client.ico
var icoData []byte

var (
	logger       *zap.Logger
	runUILoop    bool = false
	window       *windows.Window
	addressValue string = "https://www.baidu.com/"
)

func ShowMain(rootLogger *zap.Logger) {
	logger = rootLogger
	systray.Run(onReady, onExit)
}

func onReady() {
	systray.SetIcon(icoData)
	systray.SetTitle("WindowsHelper")
	systray.SetTooltip("客户端助理")
	mainPageMenuItem := systray.AddMenuItem("首页", "首页")
	browserMenuItem := systray.AddMenuItem("安全浏览器", "浏览器")
	settingMenuItem := systray.AddMenuItem("设置", "打开设置")
	mockMenuItem := systray.AddMenuItem("拨测", "网络应用测试")
	testMenuItem := systray.AddMenuItem("测试", "测试")
	quitMenuItem := systray.AddMenuItem("退出", "完全退出应用")

	go func() {
		for {
			select {
			case <-quitMenuItem.ClickedCh:
				systray.Quit()
				ShutdownCef()
				os.Exit(0)
			case <-settingMenuItem.ClickedCh:
				showSettingWindow()
			case <-mockMenuItem.ClickedCh:
				runRobotTest()
			case <-mainPageMenuItem.ClickedCh:
				showMainWindow()
			case <-browserMenuItem.ClickedCh:
				showBrowserWindow()
			case <-testMenuItem.ClickedCh:
				TestMenuItem()
			}
		}
	}()
}

func showBrowserWindow() {
	go func() {
		if runUILoop {
			createBrowserWindow()
		} else {
			runUILoop = true
			err := loop.Run(createBrowserWindow)
			if err != nil {
				runUILoop = false
				fmt.Println("Error: ", err)
			}
		}

	}()
}

func createBrowserWindow() error {
	w, err := windows.NewWindow("Browser", createToolbar())
	if err != nil {
		return err
	}
	w.SetScroll(false, true)
	window = w

	createBrowser("aaa", "http://www.sina.com.cn", w.NativeHandle(), 0, 100, 800, 800)
	return nil
}

func GetMainWindowHandler() win.HWND {
	if window != nil {
		return window.NativeHandle()
	}
	return 0
}

func showSettingWindow() {
	go func() {
		if runUILoop {
			createWindow()
		} else {
			runUILoop = true
			err := loop.Run(createWindow)
			if err != nil {
				runUILoop = false
				fmt.Println("Error: ", err)
			}
		}
	}()
}
func runRobotTest() {
	robot.RunScript(logger)
}
func showMainWindow() {
	createBrowser("百度", "http://baidu.com", 0, 0, 0, 0, 0)
}

func onExit() {
	// clean up here
}

func createWindow() error {
	w, err := windows.NewWindow("Controls", renderWindow())
	if err != nil {
		return err
	}
	w.SetScroll(false, true)
	return nil
}

func createToolbar() base.Widget {
	browserTab := &goey.Tabs{
		Children: []goey.TabItem{
			{
				Caption: "Sina",
			},
			{
				Caption: "Baidu",
			},
		},
	}

	widget := &goey.VBox{
		Children: []base.Widget{
			browserTab,
			&goey.HBox{
				Children: []base.Widget{
					&goey.Button{Text: "Back", OnClick: func() {
						goBack()
					}},
					&goey.Button{Text: "Forward", OnClick: func() {
						goForward()
					}},
					&goey.Button{Text: "Reload", Default: true, OnClick: func() {
						goReload()
					}},
					&goey.Label{Text: "Address:"},
					&goey.Expand{
						Child: &goey.TextInput{
							Value:       addressValue,
							Placeholder: "Input Address",
							OnChange: func(v string) {
								addressValue = v
							},
							OnEnterKey: func(v string) {
								goNewAddress()
							},
						},
					},
					&goey.Button{Text: "转到", Default: true, OnClick: func() {
						window.Message("访问" + addressValue).WithInfo().WithTitle("提示").Show()
						goNewAddress()
					}},
				},
			},
		},
	}
	return widget
}

/**
 * navigate to new address
 */
func goNewAddress() {
	formalUrl := strings.ToLower(addressValue)
	if !strings.HasPrefix(formalUrl, "http://") && !strings.HasPrefix(formalUrl, "https://") {
		addressValue = "http://" + addressValue
	}
	loadUrl(addressValue)
}

func renderWindow() base.Widget {
	widget := &goey.Tabs{
		Insets: goey.DefaultInsets(),
		Children: []goey.TabItem{
			{
				Caption: "Input",
				Child: &goey.VBox{
					Children: []base.Widget{
						&goey.Label{Text: "Text input:"},
						&goey.TextInput{
							Value: "Some input...", Placeholder: "Type some text here.  And some more.  And something really long.",
							OnChange: func(v string) { println("text input ", v) }, OnEnterKey: func(v string) { println("t1* ", v) },
						},
						&goey.Label{Text: "Password input:"},
						&goey.TextInput{
							Value: "", Placeholder: "Don't share", Password: true,
							OnChange: func(v string) { println("password input ", v) },
						},
						&goey.Label{Text: "Integer input:"},
						&goey.IntInput{
							Value: 3, Placeholder: "Please enter a number",
							Min: -100, Max: 100,
							OnChange: func(v int64) { println("int input ", v) },
						},
						&goey.Label{Text: "Date input:"},
						&goey.DateInput{
							Value:    time.Now().Add(24 * time.Hour),
							OnChange: func(v time.Time) { println("date input: ", v.String()) },
						},
						&goey.Label{Text: "Select input (combobox):"},
						&goey.SelectInput{
							Items:    []string{"Choice 1", "Choice 2", "Choice 3"},
							OnChange: func(v int) { println("select input: ", v) },
						},
						&goey.Label{Text: "Number input:"},
						&goey.Slider{
							Value: 25, Min: 0, Max: 100,
							OnChange: func(v float64) { println("slider input: ", v) },
						},
						&goey.HR{},
						&goey.Expand{Child: &goey.TextArea{
							Value: "", Placeholder: "Room to write",
							OnChange: func(v string) { println("text area: ", v) },
						}},
					},
				},
			},
			{
				Caption: "Buttons",
				Child: &goey.VBox{
					Children: []base.Widget{
						&goey.HBox{Children: []base.Widget{
							&goey.Button{Text: "Left 1", Default: true},
							&goey.Button{Text: "Left 2"},
						}},
						&goey.HBox{
							Children: []base.Widget{
								&goey.Button{Text: "Center"},
							},
							AlignMain: goey.MainCenter,
						},
						&goey.HBox{
							Children: []base.Widget{
								&goey.Button{Text: "D1"},
								&goey.Button{Text: "D2", Disabled: true},
								&goey.Button{Text: "D3"},
							},
							AlignMain: goey.MainEnd,
						},
						&goey.HR{},
						&goey.Label{Text: "Check boxes:"},
						&goey.Checkbox{
							Value: true, Text: "Please click on the checkbox A",
							OnChange: func(v bool) { println("check box input: ", v) },
						},
						&goey.Checkbox{
							Text:     "Please click on the checkbox B",
							OnChange: func(v bool) { println("check box input: ", v) },
						},
						&goey.Checkbox{
							Text:     "Please click on the checkbox C",
							Disabled: true,
						},
					},
				},
			},
			{
				Caption: "Lorem",
				Child: &goey.VBox{
					Children: []base.Widget{
						&goey.P{Text: "wss", Align: goey.JustifyFull},
						&goey.P{Text: "This is a paragraph, but without much text.", Align: goey.JustifyLeft},
						&goey.P{Text: "This is a paragraph, but without much text.", Align: goey.JustifyCenter},
						&goey.P{Text: "This is a paragraph, but without much text.", Align: goey.JustifyRight},
					},
					AlignMain: goey.MainCenter,
				},
			},
		},
	}
	return &goey.Padding{
		Insets: goey.DefaultInsets(),
		Child:  widget,
	}
}
