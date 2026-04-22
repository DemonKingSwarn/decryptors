package decryptor

type DecryptedSource struct {
	File  string `json:"file"`
	Type  string `json:"type"`
	Label string `json:"label"`
}

type DecryptedTrack struct {
	File  string `json:"file"`
	Kind  string `json:"kind"`
	Label string `json:"label"`
}

type DecryptResponse struct {
	Sources []DecryptedSource `json:"sources"`
	Tracks  []DecryptedTrack  `json:"tracks"`
}
