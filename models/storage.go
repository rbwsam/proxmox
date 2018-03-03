package models

type Storage struct {
	// Required
	Storage  string `json:"storage"`
	Type     string `json:"type"`
	// Optional
	Content  string `json:"content,omitempty"`
	ThinPool string `json:"thinpool,omitempty"`
	Digest   string `json:"digest,omitempty"`
	VGName   string `json:"vgname,omitempty"`
	Path 	 string `json:"path,omitempty"`
}
