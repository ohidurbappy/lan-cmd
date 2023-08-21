### Lan Commander

A simple LAN commanding tool

#### The problem

I was watching a video tutorial on youtube on my secondary PC and coding on my macbook. And everytime I needed to pause the video I had to switch my keyboard and mouse to the secondary PC. I wanted to be able to pause the video from my macbook.

#### The solution

I created a simple tool that allows me to send commands to my secondary PC from my macbook. I can now pause the video from my macbook.
The server app is installed on the windows PC and the client app is installed on the macbook.
Both app stays on the system tray and the client app has a menu bar icon.
Whenever I need to play, pause, stop, or skip the video I just press some hotkeys on my macbook and the command is sent to the server app on the windows PC.

#### How it works

The server app is a simple http server that listens for commands on a specific port. 
The client app scans the local network for the server app and connects to it.
The client app then sends the commands to the server app via http requests.


Install icnsify


```
brew tap jackmordaunt/homebrew-tap # Ensure tap is added first.
brew install icnsify
```

Install windows build tool chain when working on linux/mac

```
brew install mingw-w64
```