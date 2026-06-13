package main

import (
	"embed"
	"fmt"
	"io/fs"
	"net"
	"net/http"
	"os/exec"
	"runtime"
	"time"
)

//go:embed web
var webFiles embed.FS

// Serve はWeb版（アイコンで操作する画面）をローカルで開く。
// ブラウザを自動で開き、サーバを起動したまま待機する。
func Serve() {
	sub, err := fs.Sub(webFiles, "web")
	if err != nil {
		fmt.Println("Web画面の読み込みに失敗:", err)
		return
	}

	// 空きポートを確保
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		fmt.Println("サーバの起動に失敗:", err)
		return
	}
	url := fmt.Sprintf("http://%s/", ln.Addr().String())

	http.Handle("/", http.FileServer(http.FS(sub)))
	fmt.Printf("Web版を起動しました → %s\n（終了するにはこのウィンドウを閉じるか Ctrl+C）\n", url)

	go func() {
		time.Sleep(400 * time.Millisecond)
		openBrowser(url)
	}()
	if err := http.Serve(ln, nil); err != nil {
		fmt.Println("サーバ終了:", err)
	}
}

// openBrowser は既定のブラウザでURLを開く（Windows/macOS/Linux対応）。
func openBrowser(url string) {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("rundll32", "url.dll,FileProtocolHandler", url)
	case "darwin":
		cmd = exec.Command("open", url)
	default:
		cmd = exec.Command("xdg-open", url)
	}
	if err := cmd.Start(); err != nil {
		fmt.Println("ブラウザを自動で開けませんでした。上のURLを手動で開いてください。")
	}
}
