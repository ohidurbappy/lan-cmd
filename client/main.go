package main

import (
	_ "embed"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"

	"github.com/getlantern/systray"
	"golang.design/x/hotkey"

	// "os"
	"strconv"
	"strings"
	"sync"
	"time"
)

//go:embed assets/link-red.png
var iconLinkRed []byte

//go:embed assets/link-green.png
var iconLinkGreen []byte

var targetHostIp string = ""

const targetAppPort = 9431

var unreachable = 0

func main() {

	systray.Run(onReady, onExit)

}

func refreshTargetHost() {
	localIP, _ := getLocalSubnet() // Get the local subnet dynamically

	// remove the last octet from the localIP
	localIP = strings.TrimSuffix(localIP, "0")

	var wg sync.WaitGroup

	newIp := ""

	for i := 1; i <= 254; i++ {
		targetIP := localIP + strconv.Itoa(i)
		wg.Add(1)
		go func(ip string) {
			defer wg.Done()

			if scanTargetPort(ip) != "" {
				newIp = ip
			}
		}(targetIP)

	}

	wg.Wait()

	if newIp != "" {
		unreachable = 0
		targetHostIp = newIp
	} else if unreachable > 2 {
		targetHostIp = ""
		unreachable++
	} else {
		unreachable = 0
	}
}

func getLocalSubnet() (string, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "", err
	}

	for _, addr := range addrs {
		ip, _, err := net.ParseCIDR(addr.String())
		if err == nil && !ip.IsLoopback() && ip.To4() != nil {
			mask := net.IPMask(ip.DefaultMask())
			subnet := ip.Mask(mask)
			return subnet.String(), nil
		}
	}

	return "", fmt.Errorf("unable to determine local subnet")
}

// func hostIsUp(ip string) bool {
// 	cmd := exec.Command("ping", "-c", "1", ip) // Adjust for your operating system
// 	err := cmd.Run()
// 	return err == nil
// }

// func scanPorts(ip string) {
// 	startPort := 1
// 	endPort := 100 // Adjust the range of ports you want to scan

// 	fmt.Printf("Scanning %s\n", ip)

// 	for port := startPort; port <= endPort; port++ {
// 		target := ip + ":" + strconv.Itoa(port)
// 		conn, err := net.DialTimeout("tcp", target, 1*time.Second)
// 		if err == nil {
// 			conn.Close()
// 			fmt.Printf("Port %d is open on %s\n", port, ip)
// 		}
// 	}
// }

func scanTargetPort(ip string) string {
	// println("Scanning IP: ", ip)
	target := ip + ":" + strconv.Itoa(targetAppPort)
	conn, err := net.DialTimeout("tcp", target, 1*time.Second)
	if err == nil {
		conn.Close()
		fmt.Printf("Port %d is open on %s\n", targetAppPort, ip)
		return ip
	}

	return ""
}

func sendGetRequest(ip string, port int64, path string) {
	url := fmt.Sprintf("http://%s:%d%s", ip, port, path)
	// fmt.Println("URL:>", url)

	resp, err := http.Get(url)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	if the := resp.StatusCode; the == 200 {
		if string(responseBody) == "ok" {
			println("ok")
		}
	}

	// fmt.Println("Response:")
	// fmt.Println(string(responseBody))

	// client := &http.Client{}
	// req, err := http.NewRequest("GET", url, nil)
	// if err != nil {
	// 		fmt.Println(err)
	// 		return
	// }
	// req.Header.Set("Content-Type", "application/json")
	// req.Header.Set("Accept", "application/json")

	// resp, err := client.Do(req)
	// if err != nil {
	// 		fmt.Println(err)
	// 		return
	// }
	// defer resp.Body.Close()

	// fmt.Println("response Status:", resp.Status)
	// fmt.Println("response Headers:", resp.Header)

}

// func getIcon(s string) []byte {
// 	b, err := os.ReadFile(s)
// 	if err != nil {
// 		fmt.Print(err)
// 	}
// 	return b
// }

// // a function to get the binary path
// func getBinaryPath() string {
// 	ex, err := os.Executable()
// 	if err != nil {
// 		panic(err)
// 	}
// 	return ex
// }

// systray
func onReady() {
	systray.SetIcon(iconLinkRed)
	// systray.SetTitle("Lan Commander")
	systray.SetTooltip("Lan Commander")
	mQuit := systray.AddMenuItem("Quit", "Quit the whole app")
	go func() {
		<-mQuit.ClickedCh
		systray.Quit()
	}()

	mPlay := systray.AddMenuItem("Play", "Play")
	go func() {
		<-mPlay.ClickedCh
		sendGetRequest(targetHostIp, targetAppPort, "/toggle")
	}()

	mPause := systray.AddMenuItem("Pause", "Pause")
	go func() {
		<-mPause.ClickedCh
		sendGetRequest(targetHostIp, targetAppPort, "/toggle")
	}()

	refreshTargetHost()
	if targetHostIp != "" {
		systray.SetIcon(iconLinkGreen)
	} else {
		systray.SetIcon(iconLinkRed)
	}
	// run a function every 10 seconds
	ticker := time.NewTicker(10 * time.Second)
	go func() {
		for range ticker.C {
			refreshTargetHost()
			if targetHostIp != "" {
				systray.SetIcon(iconLinkGreen)
			} else {
				systray.SetIcon(iconLinkRed)
			}
		}
	}()

	// register hotkeys
	// mainthread.Init(regHk)
	regHk()

}

func regHk() {

	// register hotkeys
	hk := hotkey.New([]hotkey.Modifier{hotkey.ModCtrl, hotkey.ModShift}, hotkey.KeySpace)
	hk1 := hotkey.New([]hotkey.Modifier{hotkey.ModCtrl, hotkey.ModShift}, hotkey.KeyRight)
	hk2 := hotkey.New([]hotkey.Modifier{hotkey.ModCtrl, hotkey.ModShift}, hotkey.KeyLeft)
	hk3 := hotkey.New([]hotkey.Modifier{hotkey.ModCtrl, hotkey.ModShift}, hotkey.KeyUp)
	hk4 := hotkey.New([]hotkey.Modifier{hotkey.ModCtrl, hotkey.ModShift}, hotkey.KeyDown)

	// register hotkeys
	_ = hk1.Register()
	_ = hk2.Register()
	_ = hk3.Register()
	_ = hk4.Register()

	err := hk.Register()

	if err != nil {
		log.Fatal(err)
	}

	go func() {
		for {
			select {
			case <-hk.Keyup():
				sendGetRequest(targetHostIp, targetAppPort, "/space")
			case <-hk1.Keyup():
				sendGetRequest(targetHostIp, targetAppPort, "/right")
			case <-hk2.Keyup():
				sendGetRequest(targetHostIp, targetAppPort, "/left")
			case <-hk3.Keyup():
				sendGetRequest(targetHostIp, targetAppPort, "/up")
			case <-hk4.Keyup():
				sendGetRequest(targetHostIp, targetAppPort, "/down")

			}
		}
	}()

}

func onExit() {

}
