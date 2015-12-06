package Utils

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func Copy(url string) {
	// usual clipboards commands
	clipboardCommands := []string{"xclip", "xsel", "pbcopy", "getclip"}

	// execute each clipboard command
	for _, command := range clipboardCommands {
		_, err := exec.LookPath(command)
		if err != nil {
			continue // we do not have access to this command to execute
		}

		cmd := exec.Command(command)
		w, err := cmd.StdinPipe()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Could not copy to clipboard: %v", err)
			os.Exit(1)
		}

		if err := cmd.Start(); err != nil {
			fmt.Fprintf(os.Stderr, "Could not start clipboard command %q: %v", command, err)
		}

		if _, err := w.Write([]byte(url)); err != nil {
			fmt.Fprintf(os.Stderr, "Could not copy to clipboard: %v", err)
			os.Exit(1)
		}

		w.Close()

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
