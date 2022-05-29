package types

//Point contains the relevant information to a seek point
type Point struct {
	time      uint
	positions []TrackPosition
}

//TrackPosition contains relevant information about different tracks and their positions
type TrackPosition struct {
	track            uint
	clusterPosition  uint
	relativePosition uint
	duration         uint
	blockNumber      uint
	codecState       uint
	references       []Reference
}

//Reference is clusters containing referenced block data
type Reference struct {
	referenceTime uint
}
