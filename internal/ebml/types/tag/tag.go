package tag

type Tag struct {
	target Target
	tags   []SimpleTag
}

type Target struct {
	logicalLevelValue uint
	logicalLevel      uint
	trackUids         []uint
	editionUids       []uint
	chapterUids       []uint
	attachementUids   []uint
}

type SimpleTag struct {
	child          *SimpleTag
	name           string
	language       string
	languageIETF   string
	defaultLaunage uint
	tagString      string
	binary         []byte
}
