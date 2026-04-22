package decryptor

import (
	"net/http"
	"strings"
)

func DecryptStream(embedLink string, client *http.Client) (string, []string, string, error) {
	if strings.Contains(embedLink, "ployan.live") {
		return DecryptPloyan(embedLink, client)
	}
	if strings.Contains(embedLink, "vidsrc.xyz") ||
		strings.Contains(embedLink, "vidsrc.me") ||
		strings.Contains(embedLink, "vidsrc.to") ||
		strings.Contains(embedLink, "vidsrc.in") ||
		strings.Contains(embedLink, "vidsrc.pm") ||
		strings.Contains(embedLink, "vidsrc.net") ||
		strings.Contains(embedLink, "vidsrc.rip") ||
		strings.Contains(embedLink, "vidsrc.icu") {
		return DecryptVidsrc(embedLink, client)
	}

	if strings.Contains(embedLink, "vidlink.pro") {
		return DecryptVidlink(embedLink, client)
	}

	if strings.Contains(embedLink, "embed.su") {
		return DecryptEmbedSu(embedLink, client)
	}

	if strings.Contains(embedLink, "multiembed.mov") || strings.Contains(embedLink, "superembeds") {
		return DecryptMultiembed(embedLink, client)
	}

	if strings.Contains(embedLink, "videostr.net") ||
		strings.Contains(embedLink, "streameeeeee.site") ||
		strings.Contains(embedLink, "streamaaa.top") ||
		strings.Contains(embedLink, "megacloud.") {
		return DecryptMegacloud(embedLink, client)
	}

	return DecryptGeneric(embedLink, client)
}
