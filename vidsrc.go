package decryptor

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"
)

func DecryptVidsrc(urlStr string, client *http.Client) (string, []string, string, error) {
	req, _ := http.NewRequest("GET", urlStr, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36")
	resp, err := client.Do(req)
	if err != nil {
		return "", nil, "", err
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)

	reCloud := regexp.MustCompile(`src="//cloudnestra\.com/rcp/([^"]+)"`)
	match := reCloud.FindSubmatch(body)
	if len(match) < 2 {
		return "", nil, "", fmt.Errorf("could not find cloudnestra iframe")
	}
	hash := string(match[1])
	cloudURL := "https://cloudnestra.com/rcp/" + hash

	req, _ = http.NewRequest("GET", cloudURL, nil)
	req.Header.Set("Referer", urlStr)
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36")
	resp, err = client.Do(req)
	if err != nil {
		return "", nil, "", err
	}
	defer resp.Body.Close()
	body, _ = io.ReadAll(resp.Body)

	rePro := regexp.MustCompile(`src:\s*'/prorcp/([^']+)'`)
	match = rePro.FindSubmatch(body)
	if len(match) < 2 {
		return "", nil, "", fmt.Errorf("could not find prorcp iframe")
	}
	proURL := "https://cloudnestra.com/prorcp/" + string(match[1])

	req, _ = http.NewRequest("GET", proURL, nil)
	req.Header.Set("Referer", "https://cloudnestra.com/")
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36")
	resp, err = client.Do(req)
	if err != nil {
		return "", nil, "", err
	}
	defer resp.Body.Close()
	body, _ = io.ReadAll(resp.Body)

	reFile := regexp.MustCompile(`file:\s*"(https://[^"]+)"`)
	match = reFile.FindSubmatch(body)
	if len(match) < 2 {
		return "", nil, "", fmt.Errorf("could not find m3u8 file")
	}
	rawM3u8 := string(match[1])
	finalURL := rawM3u8
	if !strings.HasSuffix(finalURL, ".m3u8") {
		return "", nil, "", fmt.Errorf("extracted url is not m3u8: %s", finalURL)
	}

	var subs []string
	parsedURL, _ := http.NewRequest("GET", urlStr, nil)
	_ = parsedURL
	parsed, _ := http.NewRequest("GET", urlStr, nil)
	_ = parsed
	return finalURL, subs, "https://cloudnestra.com/", nil
}

func parseCloudnestraSubtitles(_ string, _ *http.Client) ([]string, error) { return nil, nil }
