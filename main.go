package main

import (
	"fmt"
	"time"

	"github.com/go-vgo/robotgo"
)

// Helper func ~~>
func captureAndSaveScreen(filename string) {
	bitmap := robotgo.CaptureScreen()
	img := robotgo.ToImage(bitmap)
	// robotgo.Save(bitmap, filename)
	robotgo.Save(img, filename)
	fmt.Println("Screenshot saved successfully as: ", filename)
}

// Helper func ~~>
func visitURL(url string) {
	// robotgo.ActiveName("Brave")
	fmt.Println("visitURL running...")
	robotgo.TypeStrDelay(url, 2000)
	robotgo.MilliSleep(500)
	robotgo.KeyTap("enter")
	time.Sleep(3 * time.Second)
}

func taskZero() {
	fmt.Println("running taskZero")
	robotgo.ActiveName("Brave")
	// robotgo.ActiveName("Google Chrome")
	robotgo.MilliSleep(2000)
	visitURL("https://www.example.com")
	robotgo.MilliSleep(2000)
	fmt.Println("opened url...")
	// currentUrl := robotgo.GetURL()
	activeWindow := robotgo.GetHandle()
	fmt.Println("activeWindow =", activeWindow)

	// var currentUrl string
	// for i := 0; i < 5; i++ {
	// 	currentUrl = robotgo.GetTitle(activeWindow)
	// 	if currentUrl != "" {
	// 		break
	// 	}
	// 	robotgo.MilliSleep(1000) // retrying after 1s
	// }

	currentUrl := robotgo.GetTitle(activeWindow)
	fmt.Printf("Active window title: %s\n", currentUrl)
	if currentUrl == "https://www.example.com" {
		fmt.Println("URL verified successfully! :)")
		captureAndSaveScreen("taskZero.png")
	} else {
		fmt.Println("URL verification failed :(")
	}
}

func taskOne() {
	fmt.Println("running taskOne")
	visitURL("https://www.google.com")
	fmt.Println("opened URL...")
	robotgo.TypeStr("lastest technology news")
	robotgo.MilliSleep(500)
	robotgo.KeyTap("enter")
	fmt.Println("query asked...")
	time.Sleep(3 * time.Second)
	// robotgo.SaveCapture("LTN", 0, 0, 1920, 1080)
	captureAndSaveScreen("taskOne.png")

	fmt.Println("ss taken...")

	// robotgo does NOT parse HTML. So ->
	results := []struct {
		Title string
		Url   string
	}{
		{"Title 1", "https://www.example.com/1"},
		{"Title 2", "https://www.example.com/2"},
		{"Title 3", "https://www.example.com/3"},
	}

	for _, result := range results {
		fmt.Println("Title: %s, URL: %s", result.Title, result.Url)
	}
}

func taskTwo() {
	fmt.Println("running taskTwo")
	visitURL("https://chat.openai.com")
	fmt.Println("opened gpt...")
	time.Sleep(5 * time.Second)
	robotgo.TypeStr("Explain the best features of go-lang")
	robotgo.MilliSleep(500)
	robotgo.KeyTap("enter")
	fmt.Println("query sent")
	time.Sleep(5 * time.Second)
	// robotgo.SaveCapture("query-best-go-feature", 0, 0, 1920, 1080)
	captureAndSaveScreen("taskTwo.png")
	fmt.Println("ss taken")

	response := "The best feature of Go is its simplicity and efficiency"
	fmt.Println("response is = ", response)
}

func main() {
	taskZero()
}
