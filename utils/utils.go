package Utils

import (
	"os"
	"os/exec"
	"runtime"
)

func Copy(url string) {
	// usual clipboards commands
	clipboardCommands := []string{"xclip", "xsel", "pbpaste", "getclip"}

	// execute each clipboard command
	for _, command := range clipboardCommands {
		echo := exec.Command("echo", url)

		cmd := exec.Command(command)
		cmd.Stdin, _ = echo.StdoutPipe()
		cmd.Stdout = os.Stdout
		cmd.Start()
		echo.Run()
		cmd.Wait()
	}
}

func OpenBrowser(url string) {
        os := runtime.GOOS
        switch {
        case os == "windows":
                exec.Command("cmd", "/c", "start", url).Run()
        case os == "darwin":
                exec.Command("open", url).Run()
        case os == "linux":
                exec.Command("xdg-open", url).Run()
        }
}
