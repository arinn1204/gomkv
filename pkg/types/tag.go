package types

//Tag is a metadata descriptor
type Tag struct {
	Targets *Target      `json:",omitempty"`
	Tags    []*SimpleTag `json:",omitempty"`
}

//Target describes what element the tag is describing
//If target is empty then it applies to the whole segment
type Target struct {
	TargetTypeValue  uint   `json:",omitempty"`
	TargetType       string `json:",omitempty"`
	TagTrackUID      uint   `json:",omitempty"`
	TagEditionUID    uint   `json:",omitempty"`
	TagAttachmentUID uint   `json:",omitempty"`
}

//SimpleTag contains general information about the target
type SimpleTag struct {
	Child           *SimpleTag `json:",omitempty"`
	TagName         string     `json:",omitempty"`
	TagLanguage     string     `json:",omitempty"`
	TagLanguageIETF string     `json:",omitempty"`
	TagDefault      uint       `json:",omitempty"`
	TagDefaultBogus string     `json:",omitempty"`
	TagString       string     `json:",omitempty"`
}
