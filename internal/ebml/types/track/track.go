package track

//Track is a top level element describing many entries
type Track struct {
	entries []Entry
}

//Entry describes a track with all elements
type Entry struct {
	number                      uint
	uid                         uint
	entryType                   uint
	flag                        Flag
	minCache                    uint
	maxCache                    uint
	defaultDuration             uint
	defaultDecodedFieldDuration uint
	maxBlockAdditionalID        uint
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

//Codec is a reference struct used to encapsulate some of the Entry data
type Codec struct {
	id                      uint
	data                    []byte
	name                    string
	willTryDamaged          uint
	builtInDelayNanoseconds uint
}

//Flag is a reference struct used to encapsulate some of the Entry data
type Flag struct {
	enabled     uint
	flagDefault uint
	forced      uint
	lacing      uint
}

//Translate is a mapping between this track entry and the chapter codec data
type Translate struct {
	editionUids []uint
	codec       uint
	trackID     []byte
}

//Video is all of the video settings
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

//Pixel is all the pixel related settings
type Pixel struct {
	width  uint
	height uint
	crop   Crop
}

//Crop is all the cropping settings
type Crop struct {
	bottom uint
	top    uint
	left   uint
	right  uint
}

//Display is all the display related settings
type Display struct {
	width           uint
	height          uint
	unit            uint
	aspectRatioType uint
}

//Color is all the color settings
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

//Chroma is the chroma settings
type Chroma struct {
	chroma SubSample
	cb     SubSample
	siting SubSample
}

//SubSample is a reference struct to contain horizontal and vertical settings
type SubSample struct {
	horizontal uint
	veritcal   uint
}

//ColorMetadata is the metadata of the color settings, includes things like RGB settings
type ColorMetadata struct {
	red          Chromaticity
	green        Chromaticity
	white        Chromaticity
	blue         Chromaticity
	maxLuminance uint
	minLuminance uint
}

//Chromaticity is just an easier way to encapsulate the metadata settings
type Chromaticity struct {
	x float32
	y float32
}

//Audio settings
type Audio struct {
	sampling SamplingFrequency
	channels uint
	bitDepth uint
}

//SamplingFrequency just encapsulates the input and output frequencies
type SamplingFrequency struct {
	input  float32
	output float32
}

//Operation is all the operations to be applied to create the track
type Operation struct {
	videoTracks []Plane
	joinBlocks  JoinBlocks
}

//Plane contains a video plane track that need to be combined to create this 3D track
type Plane struct {
	trackType uint
	uid       uint
}

//JoinBlocks contains the list of all tracks whose Blocks need to be combined to create this virtual track
type JoinBlocks struct {
	uids []uint
}

//Projection Describes the video projection details. Used to render spherical, VR videos or flipping videos horizontally/vertically
type Projection struct {
	projectionType uint
	pose           Pose
	privateData    []byte
}

//Pose descripes all the pose details for projection
type Pose struct {
	yew   float32
	pitch float32
	roll  float32
}

//ContentEncoding describes all of the encoding settings
type ContentEncoding struct {
	order        uint
	scope        uint
	encodingType uint
	compression  Compression
	encryption   Encryption
}

//Compression describes the compression settings
type Compression struct {
	algorithm uint
	settings  []byte
}

//Encryption describes the encryption settings
type Encryption struct {
	algorithm              uint
	keyID                  []byte
	aesSettings            AesSettings
	signature              []byte
	privateKeyID           []byte
	signatureAlgorithm     uint
	signatureHashAlgorithm uint
}

//AesSettings describes the AES settings
type AesSettings struct {
	cipherMode uint
}
