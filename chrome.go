package chromedev

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

type ChromeVersion struct {
	Browser              string `json:"Browser"`
	ProtocolVersion      string `json:"Protocol-Version"`
	UserAgent            string `json:"User-Agent"`
	V8Version            string `json:"V8-Version"`
	WebKitVersion        string `json:"WebKit-Version"`
	WebSocketDebuggerUrl string `json:"webSocketDebuggerUrl"`
}

func Open(port uint) error {
	versionURL := fmt.Sprintf("http://localhost:%d/json/version", port)
	resp, err := http.Get(versionURL)
	if err == nil && resp.StatusCode == http.StatusOK {
		defer resp.Body.Close()
		var version ChromeVersion
		if err := json.NewDecoder(resp.Body).Decode(&version); err != nil {
			return fmt.Errorf("failed to decode json response: %w", err)
		}
		return dumpToFile(version)
	}

	chromePath := "/Applications/Google Chrome.app/Contents/MacOS/Google Chrome"
	cmd := exec.Command(chromePath, fmt.Sprintf("--remote-debugging-port=%d", port))
	if err := cmd.Start(); err != nil {
		return fmt.Errorf("failed to start chrome: %w", err)
	}

	// Wait for Chrome to start
	time.Sleep(2 * time.Second)

	resp, err = http.Get(versionURL)
	if err != nil || resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to get json version after starting chrome: %w", err)
	}
	defer resp.Body.Close()

	var version ChromeVersion
	if err := json.NewDecoder(resp.Body).Decode(&version); err != nil {
		return fmt.Errorf("failed to decode json response: %w", err)
	}

	return dumpToFile(version)
}

func dumpToFile(version ChromeVersion) error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("failed to get home directory: %w", err)
	}
	filePath := filepath.Join(homeDir, ".chromedev.json")
	data, err := json.MarshalIndent(version, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal json: %w", err)
	}
	if err := os.WriteFile(filePath, data, 0644); err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}
	return nil
}
