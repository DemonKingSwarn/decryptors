package decryptor

import (
	"fmt"
	"net/http"
)

func DecryptMegacloud(urlStr string, client *http.Client) (string, []string, string, error) {
	return DecryptGeneric(urlStr, client)
}

func DecryptEmbedSu(urlStr string, client *http.Client) (string, []string, string, error) {
	return DecryptGeneric(urlStr, client)
}

func DecryptMultiembed(urlStr string, client *http.Client) (string, []string, string, error) {
	return DecryptGeneric(urlStr, client)
}

func DecryptGeneric(urlStr string, client *http.Client) (string, []string, string, error) {
	return urlStr, nil, "", nil
}

func debugDecrypt(name string, v any) error {
	return fmt.Errorf("%s: %#v", name, v)
}
