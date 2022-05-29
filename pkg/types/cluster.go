package types

//Cluster is the top level element that defines the block (data)
type Cluster struct {
	timestamp    uint
	track        Track
	position     uint
	previousSize uint
	blocks       []byte
	groups       []BlockGroup
}

//Track is the list of tracks not used in this part of the stream
type Track struct {
	numbers []uint
}

//BlockGroup is a basic container of information about this block of data
type BlockGroup struct {
	block             []byte
	addition          BlockAddition
	duration          uint
	referencePriority uint
	referenceBlocks   []int
	codecState        []byte
	discardPadding    int
}

//BlockAddition contains additional blocks to complete the main one
type BlockAddition struct {
	more []BlockMore
}

//BlockMore contains more block data
type BlockMore struct {
	id   uint
	data []byte
}
