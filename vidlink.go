package decryptor

import (
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"
)

func DecryptVidlink(urlStr string, client *http.Client) (string, []string, string, error) {
	re := regexp.MustCompile(`/(movie|tv)/([^/?#]+)`)
	matches := re.FindStringSubmatch(urlStr)
	if len(matches) < 3 {
		return "", nil, "", fmt.Errorf("could not parse vidlink url")
	}

	tmdbID := matches[2]
	subURL := fmt.Sprintf("https://vidlink.pro/api/subtitles/%s", tmdbID)

	req, _ := http.NewRequest("GET", subURL, nil)
	resp, err := client.Do(req)

	var subs []string
	if err == nil && resp.StatusCode == 200 {
		var tracks []struct {
			URL   string `json:"url"`
			Label string `json:"label"`
		}
		if err := json.NewDecoder(resp.Body).Decode(&tracks); err == nil {
			for _, t := range tracks {
				label := strings.ToLower(t.Label)
				if strings.Contains(label, "english") || strings.Contains(label, " eng") || label == "eng" {
					subs = append(subs, t.URL)
				}
			}
		}
		resp.Body.Close()
	}

	videoLink, _, referer, err := DecryptVidlinkStream(urlStr, tmdbID, client)
	return videoLink, subs, referer, err
}

func DecryptVidlinkStream(urlStr, tmdbID string, client *http.Client) (string, []string, string, error) {
	keyHex := "2de6e6ea13a9df9503b11a6117fd7e51941e04a0c223dfeacfe8a1dbb6c52783"
	key, _ := hex.DecodeString(keyHex)

	encryptedID := aesEncrypt(tmdbID, key)
	encodedID := base64.StdEncoding.EncodeToString([]byte(encryptedID))

	apiURL := fmt.Sprintf("https://vidlink.pro/api/b/movie/%s", encodedID)

	req, _ := http.NewRequest("GET", apiURL, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36")
	req.Header.Set("Referer", urlStr)

	resp, err := client.Do(req)
	if err != nil {
		return "", nil, "", err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	parts := strings.Split(string(body), ":")
	if len(parts) != 2 {
		return DecryptGeneric(urlStr, client)
	}

	return parts[0], nil, parts[1], nil
}
