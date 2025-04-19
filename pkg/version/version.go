package version

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/logrusorgru/aurora"
)

const (
	gitHubAPIURL   = "https://api.github.com/repos/nimaism/trackit/releases/latest"
	currentVersion = "v0.1.0"
)

const banner = `
████████╗██████╗░░█████╗░░█████╗░██╗░░██╗██╗████████╗
╚══██╔══╝██╔══██╗██╔══██╗██╔══██╗██║░██╔╝██║╚══██╔══╝
░░░██║░░░██████╔╝███████║██║░░╚═╝█████═╝░██║░░░██║░░░
░░░██║░░░██╔══██╗██╔══██║██║░░██╗██╔═██╗░██║░░░██║░░░
░░░██║░░░██║░░██║██║░░██║╚█████╔╝██║░╚██╗██║░░░██║░░░
░░░╚═╝░░░╚═╝░░╚═╝╚═╝░░╚═╝░╚════╝░╚═╝░░╚═╝╚═╝░░░╚═╝░░░
        TrackIt ` + currentVersion + `
`

const newVersionBannerTemplate = `
███╗   ██╗███████╗██╗    ██╗    ██╗   ██╗██████╗ ██████╗  █████╗ ████████╗███████╗     █████╗ ██╗   ██╗ █████╗ ██╗██╗      █████╗ ██████╗ ██╗     ███████╗
████╗  ██║██╔════╝██║    ██║    ██║   ██║██╔══██╗██╔══██╗██╔══██╗╚══██╔══╝██╔════╝    ██╔══██╗██║   ██║██╔══██╗██║██║     ██╔══██╗██╔══██╗██║     ██╔════╝
██╔██╗ ██║█████╗  ██║ █╗ ██║    ██║   ██║██████╔╝██║  ██║███████║   ██║   █████╗      ███████║██║   ██║███████║██║██║     ███████║██████╔╝██║     █████╗
██║╚██╗██║██╔══╝  ██║███╗██║    ██║   ██║██╔═══╝ ██║  ██║██╔══██║   ██║   ██╔══╝      ██╔══██║╚██╗ ██╔╝██╔══██║██║██║     ██╔══██║██╔══██╗██║     ██╔══╝
██║ ╚████║███████╗╚███╔███╔╝    ╚██████╔╝██║     ██████╔╝██║  ██║   ██║   ███████╗    ██║  ██║ ╚████╔╝ ██║  ██║██║███████╗██║  ██║██████╔╝███████╗███████╗
╚═╝  ╚═══╝╚══════╝ ╚══╝╚══╝      ╚═════╝ ╚═╝     ╚═════╝ ╚═╝  ╚═╝   ╚══════╝    ╚═╝  ╚═╝  ╚═══╝  ╚═╝  ╚═╝╚═╝╚══════╝╚═╝  ╚═╝╚═════╝ ╚══════╝╚══════╝
 							go install github.com/nimaism/trackit/cmd/trackit@%s
`

func ShowVersion() {
	fmt.Println(aurora.Bold(banner))
}

func CheckLatestVersion() {
	latestVersion, err := fetchLatestVersion()
	if err != nil {
		log.Fatalf("Error checking for latest version: %v", err)
		return
	}

	if latestVersion != currentVersion {
		displayNewVersionBanner(latestVersion)
	}
}

func fetchLatestVersion() (string, error) {
	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Get(gitHubAPIURL)
	if err != nil {
		return "", fmt.Errorf("failed to fetch latest version: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var release Release
	if err := json.NewDecoder(resp.Body).Decode(&release); err != nil {
		return "", fmt.Errorf("failed to decode JSON response: %w", err)
	}

	if release.TagName == "" {
		return "", fmt.Errorf("empty tag name in release data")
	}

	return release.TagName, nil
}

func displayNewVersionBanner(latestVersion string) {
	banner := fmt.Sprintf(newVersionBannerTemplate, latestVersion)
	fmt.Println(aurora.Blue(banner))
}
