package types

//DisplayTrack is a top level element describing many entries
type DisplayTrack struct {
	Entries []*Entry `json:",omitempty"`
}

//Entry describes a track with all elements
type Entry struct {
	Name         string `json:",omitempty"`
	Language     string `json:",omitempty"`
	LanguageIETF string `json:",omitempty"`
	CodecID      uint   `json:",omitempty"`
	CodecName    string `json:",omitempty"`
	Video        *Video `json:",omitempty"`
	Audio        *Audio `json:",omitempty"`
}

//Video is all of the video settings
type Video struct {
	PixelWidth      uint `json:",omitempty"`
	PixelHeight     uint `json:",omitempty"`
	DisplayWidth    uint `json:",omitempty"`
	DisplaylHeight  uint `json:",omitempty"`
	DisplaylUnit    uint `json:",omitempty"`
	AspectRatioType uint `json:",omitempty"`
}

//Audio settings
type Audio struct {
	SamplingFrequency       float32 `json:",omitempty"`
	OutputSamplingFrequency float32 `json:",omitempty"`
	Channels                uint    `json:",omitempty"`
	BitDepth                uint    `json:",omitempty"`
}
