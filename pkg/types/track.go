package types

//DisplayTrack is a top level element describing many entries
type DisplayTrack struct {
	Entries []*Entry `json:",omitempty"`
}

//Entry describes a track with all elements
type Entry struct {
	number                      uint              `json:",omitempty"`
	uid                         uint              `json:",omitempty"`
	entryType                   uint              `json:",omitempty"`
	flag                        Flag              `json:",omitempty"`
	minCache                    uint              `json:",omitempty"`
	maxCache                    uint              `json:",omitempty"`
	defaultDuration             uint              `json:",omitempty"`
	defaultDecodedFieldDuration uint              `json:",omitempty"`
	maxBlockAdditionalID        uint              `json:",omitempty"`
	name                        string            `json:",omitempty"`
	language                    string            `json:",omitempty"`
	languageOverride            string            `json:",omitempty"`
	codec                       Codec             `json:",omitempty"`
	overlayTracks               []uint            `json:",omitempty"`
	seekPreRoll                 uint              `json:",omitempty"`
	translations                []EntryTranslate  `json:",omitempty"`
	video                       Video             `json:",omitempty"`
	audio                       Audio             `json:",omitempty"`
	operation                   Operation         `json:",omitempty"`
	encodings                   []ContentEncoding `json:",omitempty"`
}

//Codec is a reference struct used to encapsulate some of the Entry data
type Codec struct {
	id                      uint   `json:",omitempty"`
	data                    []byte `json:",omitempty"`
	name                    string `json:",omitempty"`
	willTryDamaged          uint   `json:",omitempty"`
	builtInDelayNanoseconds uint   `json:",omitempty"`
}

//Flag is a reference struct used to encapsulate some of the Entry data
type Flag struct {
	enabled     uint `json:",omitempty"`
	flagDefault uint `json:",omitempty"`
	forced      uint `json:",omitempty"`
	lacing      uint `json:",omitempty"`
}

//EntryTranslate is a mapping between this track entry and the chapter codec data
type EntryTranslate struct {
	editionUids []uint `json:",omitempty"`
	codec       uint   `json:",omitempty"`
	trackID     []byte `json:",omitempty"`
}

//Video is all of the video settings
type Video struct {
	flagInterlaced     uint            `json:",omitempty"`
	fieldOrder         uint            `json:",omitempty"`
	stereo3DVideoWorks uint            `json:",omitempty"`
	alphaVideoMode     uint            `json:",omitempty"`
	pixel              Pixel           `json:",omitempty"`
	Display            DisplaySettings `json:",omitempty"`
	colorSpace         []byte          `json:",omitempty"`
	color              Color           `json:",omitempty"`
	projection         Projection      `json:",omitempty"`
}

//Pixel is all the pixel related settings
type Pixel struct {
	width  uint `json:",omitempty"`
	height uint `json:",omitempty"`
	crop   Crop `json:",omitempty"`
}

//Crop is all the cropping settings
type Crop struct {
	bottom uint `json:",omitempty"`
	top    uint `json:",omitempty"`
	left   uint `json:",omitempty"`
	right  uint `json:",omitempty"`
}

//DisplaySettings is all the display related settings
type DisplaySettings struct {
	width           uint `json:",omitempty"`
	height          uint `json:",omitempty"`
	unit            uint `json:",omitempty"`
	aspectRatioType uint `json:",omitempty"`
}

//Color is all the color settings
type Color struct {
	matrixCoefficients            uint          `json:",omitempty"`
	bitsPerChannel                uint          `json:",omitempty"`
	transferCharacteristics       uint          `json:",omitempty"`
	colorRange                    uint          `json:",omitempty"`
	primaries                     uint          `json:",omitempty"`
	maximumContentLightLevel      uint          `json:",omitempty"`
	maximumFrameAverageLightLevel uint          `json:",omitempty"`
	metadata                      ColorMetadata `json:",omitempty"`
	chroma                        Chroma        `json:",omitempty"`
}

//Chroma is the chroma settings
type Chroma struct {
	chroma SubSample `json:",omitempty"`
	cb     SubSample `json:",omitempty"`
	siting SubSample `json:",omitempty"`
}

//SubSample is a reference struct to contain horizontal and vertical settings
type SubSample struct {
	horizontal uint `json:",omitempty"`
	veritcal   uint `json:",omitempty"`
}

//ColorMetadata is the metadata of the color settings, includes things like RGB settings
type ColorMetadata struct {
	red          Chromaticity `json:",omitempty"`
	green        Chromaticity `json:",omitempty"`
	white        Chromaticity `json:",omitempty"`
	blue         Chromaticity `json:",omitempty"`
	maxLuminance uint         `json:",omitempty"`
	minLuminance uint         `json:",omitempty"`
}

//Chromaticity is just an easier way to encapsulate the metadata settings
type Chromaticity struct {
	x float32 `json:",omitempty"`
	y float32 `json:",omitempty"`
}

//Audio settings
type Audio struct {
	sampling SamplingFrequency `json:",omitempty"`
	channels uint              `json:",omitempty"`
	bitDepth uint              `json:",omitempty"`
}

//SamplingFrequency just encapsulates the input and output frequencies
type SamplingFrequency struct {
	input  float32 `json:",omitempty"`
	output float32 `json:",omitempty"`
}

//Operation is all the operations to be applied to create the track
type Operation struct {
	videoTracks []Plane    `json:",omitempty"`
	joinBlocks  JoinBlocks `json:",omitempty"`
}

//Plane contains a video plane track that need to be combined to create this 3D track
type Plane struct {
	trackType uint `json:",omitempty"`
	uid       uint `json:",omitempty"`
}

//JoinBlocks contains the list of all tracks whose Blocks need to be combined to create this virtual track
type JoinBlocks struct {
	uids []uint `json:",omitempty"`
}

//Projection Describes the video projection details. Used to render spherical, VR videos or flipping videos horizontally/vertically
type Projection struct {
	projectionType uint   `json:",omitempty"`
	pose           Pose   `json:",omitempty"`
	privateData    []byte `json:",omitempty"`
}

//Pose descripes all the pose details for projection
type Pose struct {
	yew   float32 `json:",omitempty"`
	pitch float32 `json:",omitempty"`
	roll  float32 `json:",omitempty"`
}

//ContentEncoding describes all of the encoding settings
type ContentEncoding struct {
	order        uint        `json:",omitempty"`
	scope        uint        `json:",omitempty"`
	encodingType uint        `json:",omitempty"`
	compression  Compression `json:",omitempty"`
	encryption   Encryption  `json:",omitempty"`
}

//Compression describes the compression settings
type Compression struct {
	algorithm uint   `json:",omitempty"`
	settings  []byte `json:",omitempty"`
}

//Encryption describes the encryption settings
type Encryption struct {
	algorithm              uint        `json:",omitempty"`
	keyID                  []byte      `json:",omitempty"`
	aesSettings            AesSettings `json:",omitempty"`
	signature              []byte      `json:",omitempty"`
	privateKeyID           []byte      `json:",omitempty"`
	signatureAlgorithm     uint        `json:",omitempty"`
	signatureHashAlgorithm uint        `json:",omitempty"`
}

//AesSettings describes the AES settings
type AesSettings struct {
	cipherMode uint `json:",omitempty"`
}
