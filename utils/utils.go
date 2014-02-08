package Utils

import (
	"os"
	"os/exec"
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
	cmd := exec.Command("xdg-open", url)
	cmd.Start()
}
