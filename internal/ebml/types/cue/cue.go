package cue

type Cue struct {
	points []Point
}

type Point struct {
	time      uint
	positions []TrackPosition
}

type TrackPosition struct {
	track            uint
	clusterPosition  uint
	relativePosition uint
	duration         uint
	blockNumber      uint
	codecState       uint
	references       []Reference
}

type Reference struct {
	referenceTime uint
}
