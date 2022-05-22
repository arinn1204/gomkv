package track

type Track struct {
	entries []Entry
}

type Entry struct {
	number                      uint
	uid                         uint
	entryType                   uint
	flag                        Flag
	minCache                    uint
	maxCache                    uint
	defaultDuration             uint
	defaultDecodedFieldDuration uint
	maxBlockAdditionalId        uint
	name                        string
	language                    string
	languageOverride            string
	codec                       Codec
	overlayTracks               []uint
	seekPreRoll                 uint
	translations                []Translate
	video                       Video
	audio                       Audio
	operation                   Operation
	encodings                   []ContentEncoding
}

type Codec struct {
	id                      uint
	data                    []byte
	name                    string
	willTryDamaged          uint
	builtInDelayNanoseconds uint
}

type Flag struct {
	enabled     uint
	flagDefault uint
	forced      uint
	lacing      uint
}

type Translate struct {
	editionUids []uint
	codec       uint
	trackId     []byte
}

type Video struct {
	flagInterlaced     uint
	fieldOrder         uint
	stereo3DVideoWorks uint
	alphaVideoMode     uint
	pixel              Pixel
	Display            Display
	colorSpace         []byte
	color              Color
	projection         Projection
}

type Pixel struct {
	width  uint
	height uint
	crop   Crop
}

type Crop struct {
	bottom uint
	top    uint
	left   uint
	right  uint
}

type Display struct {
	width           uint
	height          uint
	unit            uint
	aspectRatioType uint
}

type Color struct {
	matrixCoefficients            uint
	bitsPerChannel                uint
	transferCharacteristics       uint
	colorRange                    uint
	primaries                     uint
	maximumContentLightLevel      uint
	maximumFrameAverageLightLevel uint
	metadata                      ColorMetadata
	chroma                        Chroma
}
type Chroma struct {
	chroma SubSample
	cb     SubSample
	siting SubSample
}

type SubSample struct {
	horizontal uint
	veritcal   uint
}

type ColorMetadata struct {
	red          Chromaticity
	green        Chromaticity
	white        Chromaticity
	blue         Chromaticity
	maxLuminance uint
	minLuminance uint
}

type Chromaticity struct {
	x float32
	y float32
}

type Audio struct {
	sampling SamplingFrequency
	channels uint
	bitDepth uint
}

type SamplingFrequency struct {
	input  float32
	output float32
}

type Operation struct {
	videoTracks []Plane
	joinBlocks  JoinBlocks
}

type Plane struct {
	trackType uint
	uid       uint
}
type JoinBlocks struct {
	uids []uint
}

type Projection struct {
	projectionType uint
	pose           Pose
	privateData    []byte
}

type Pose struct {
	yew   float32
	pitch float32
	roll  float32
}

type ContentEncoding struct {
	order        uint
	scope        uint
	encodingType uint
	compression  Compression
	encryption   Encryption
}
type Compression struct {
	algorithm uint
	settings  []byte
}

type Encryption struct {
	algorithm              uint
	keyId                  []byte
	aesSettings            AesSettings
	signature              []byte
	privateKeyId           []byte
	signatureAlgorithm     uint
	signatureHashAlgorithm uint
}
type AesSettings struct {
	cipherMode uint
}
