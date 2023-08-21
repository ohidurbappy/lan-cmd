package main

//go:generate go-winres simply --icon assets/server.png --manifest gui

import (
	_ "embed"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/getlantern/systray"
	"github.com/go-vgo/robotgo"
)

//go:embed assets/server.png
var iconServer []byte

func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / request\n")
	io.WriteString(w, "Lan Commander\n")
}

func main() {

	systray.Run(onReady, onExit)

}

func startServer() {
	// 	Spacebar or K: Play/Pause the video.
	// Left Arrow: Skip backward 5 seconds.
	// Right Arrow: Skip forward 5 seconds.
	// J: Rewind 10 seconds.
	// L: Fast forward 10 seconds.
	// Home: Go to the beginning of the video.
	// End: Go to the end of the video.
	// Number Keys (1-9): Jump to a specific percentage of the video's duration (10% increments).
	// M: Mute/unmute the video.
	// Up Arrow: Increase the volume.
	// Down Arrow: Decrease the volume.
	// F: Enter/exit full-screen mode.
	// C: Turn on/off captions or subtitles (if available).
	// Shift + N: Move to the next video in a playlist.
	// Shift + P: Move to the previous video in a playlist.
	http.HandleFunc("/", getRoot)
	http.HandleFunc("/space", func(w http.ResponseWriter, r *http.Request) {
		robotgo.KeyTap("space")
		io.WriteString(w, "space\n")
	})

	http.HandleFunc("/left", func(w http.ResponseWriter, r *http.Request) {
		robotgo.KeyTap("left")
		io.WriteString(w, "left\n")
	})

	http.HandleFunc("/right", func(w http.ResponseWriter, r *http.Request) {
		robotgo.KeyTap("right")
		io.WriteString(w, "right\n")
	})

	http.HandleFunc("/j", func(w http.ResponseWriter, r *http.Request) {
		robotgo.KeyTap("j")
		io.WriteString(w, "j\n")
	})

	http.HandleFunc("/l", func(w http.ResponseWriter, r *http.Request) {
		robotgo.KeyTap("l")
		io.WriteString(w, "l\n")
	})

	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		robotgo.KeyTap("home")
		io.WriteString(w, "home\n")
	})

	http.HandleFunc("/end", func(w http.ResponseWriter, r *http.Request) {
		robotgo.KeyTap("end")
		io.WriteString(w, "end\n")
	})

	http.HandleFunc("/m", func(w http.ResponseWriter, r *http.Request) {
		robotgo.KeyTap("m")
		io.WriteString(w, "m\n")
	})

	http.HandleFunc("/up", func(w http.ResponseWriter, r *http.Request) {
		robotgo.KeyTap("up")
		io.WriteString(w, "up\n")
	})

	http.HandleFunc("/down", func(w http.ResponseWriter, r *http.Request) {
		robotgo.KeyTap("down")
		io.WriteString(w, "down\n")
	})

	http.HandleFunc("/f", func(w http.ResponseWriter, r *http.Request) {
		robotgo.KeyTap("f")
		io.WriteString(w, "f\n")
	})

	http.HandleFunc("/c", func(w http.ResponseWriter, r *http.Request) {
		robotgo.KeyTap("c")
		io.WriteString(w, "c\n")
	})

	http.HandleFunc("/shift+n", func(w http.ResponseWriter, r *http.Request) {
		robotgo.KeyTap("shift+n")
		io.WriteString(w, "shift+n\n")
	})

	http.HandleFunc("/shift+p", func(w http.ResponseWriter, r *http.Request) {
		robotgo.KeyTap("shift+p")
		io.WriteString(w, "shift+p\n")
	})

	fmt.Printf("starting server on port 9431\n")

	err := http.ListenAndServe(":9431", nil)

	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}

func onReady() {
	systray.SetIcon(iconServer)
	systray.SetTitle("Lan Commander")
	systray.SetTooltip("Lan Commander")

	mQuit := systray.AddMenuItem("Quit", "Quit the whole app")
	go func() {
		<-mQuit.ClickedCh
		systray.Quit()
	}()

	startServer()
}

func onExit() {
	fmt.Println("Exiting...")
}
