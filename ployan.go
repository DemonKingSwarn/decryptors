package decryptor

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
)

func DecryptPloyan(urlStr string, client *http.Client) (string, []string, string, error) {
	req, _ := http.NewRequest("GET", urlStr, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36")
	resp, err := client.Do(req)
	if err != nil {
		return "", nil, "", err
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)

	re := regexp.MustCompile(`https://[^"'<> ]+/master\.m3u8`)
	if m := re.Find(body); len(m) > 0 {
		return string(m), nil, "https://ployan.live/", nil
	}

	return "", nil, "", fmt.Errorf("could not find ployan m3u8")
}
