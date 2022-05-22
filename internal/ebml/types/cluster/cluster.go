package cluster

type Cluster struct {
	timestamp    uint
	track        Track
	position     uint
	previousSize uint
	blocks       []byte
	groups       []BlockGroup
}

type Track struct {
	numbers []uint
}

type BlockGroup struct {
	block             []byte
	addition          BlockAddition
	duration          uint
	referencePriority uint
	referenceBlocks   []int
	codecState        []byte
	discardPadding    int
}

type BlockAddition struct {
	more []BlockMore
}

type BlockMore struct {
	id   uint
	data []byte
}
