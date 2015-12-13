package ytdl

// FormatKey is a string type containg a key in a video format map
type FormatKey string

// Available format Keys
const (
	FormatExtensionKey     FormatKey = "ext"
	FormatResolutionKey    FormatKey = "res"
	FormatVideoEncodingKey FormatKey = "videnc"
	FormatAudioEncodingKey FormatKey = "audenc"
	FormatProfileKey       FormatKey = "prof"
)

// Format is a map type for formats
type Format map[FormatKey]interface{}

// FORMATS is a map of all itags and their formats
var FORMATS = map[int]Format{
	5: {
		FormatExtensionKey:     "flv",
		FormatResolutionKey:    "240p",
		FormatVideoEncodingKey: "Sorenson H.283",
		FormatProfileKey:       nil,
		FormatAudioEncodingKey: "mp3",
	},
	6: {
		FormatExtensionKey:     "flv",
		FormatResolutionKey:    "270p",
		FormatVideoEncodingKey: "Sorenson H.263",
		FormatProfileKey:       nil,
		FormatAudioEncodingKey: "mp3",
	},
	13: {
		FormatExtensionKey:     "3gp",
		FormatResolutionKey:    nil,
		FormatVideoEncodingKey: "MPEG-4 Visual",
		FormatProfileKey:       nil,
		FormatAudioEncodingKey: "aac",
	},
	17: {
		FormatExtensionKey:     "3gp",
		FormatResolutionKey:    "144p",
		FormatVideoEncodingKey: "MPEG-4 Visual",
		FormatAudioEncodingKey: "aac",
		FormatProfileKey:       nil,
	},
	18: {
		FormatExtensionKey:     "mp4",
		FormatResolutionKey:    "360p",
		FormatVideoEncodingKey: "H.264",
		FormatProfileKey:       "baseline",
		FormatAudioEncodingKey: "aac",
	},
	22: {
		FormatExtensionKey:     "mp4",
		FormatResolutionKey:    "720p",
		FormatVideoEncodingKey: "H.264",
		FormatProfileKey:       "high",
		FormatAudioEncodingKey: "aac",
	},
	34: {
		FormatExtensionKey:     "flv",
		FormatResolutionKey:    "480p",
		FormatVideoEncodingKey: "H.264",
		FormatProfileKey:       "main",
		FormatAudioEncodingKey: "aac",
	},
	35: {
		FormatExtensionKey:     "flv",
		FormatResolutionKey:    "360p",
		FormatVideoEncodingKey: "H.264",
		FormatProfileKey:       "main",
		FormatAudioEncodingKey: "aac",
	},
	36: {
		FormatExtensionKey:     "3gp",
		FormatResolutionKey:    "240p",
		FormatVideoEncodingKey: "MPEG-4 Visual",
		FormatProfileKey:       "simple",
		FormatAudioEncodingKey: "aac",
	},
	37: {
		FormatExtensionKey:     "mp4",
		FormatResolutionKey:    "1080p",
		FormatVideoEncodingKey: "H.264",
		FormatProfileKey:       "high",
		FormatAudioEncodingKey: "aac",
	},
	38: {
		FormatExtensionKey:     "mp4",
		FormatResolutionKey:    "3072p",
		FormatVideoEncodingKey: "H.264",
		FormatProfileKey:       "high",
		FormatAudioEncodingKey: "aac",
	},
	43: {
		FormatExtensionKey:     "webm",
		FormatResolutionKey:    "360p",
		FormatVideoEncodingKey: "VP8",
		FormatProfileKey:       nil,
		FormatAudioEncodingKey: "vorbis",
	},
	44: {
		FormatExtensionKey:     "webm",
		FormatResolutionKey:    "480p",
		FormatVideoEncodingKey: "VP8",
		FormatProfileKey:       nil,
		FormatAudioEncodingKey: "vorbis",
	},
	45: {
		FormatExtensionKey:     "webm",
		FormatResolutionKey:    "720p",
		FormatVideoEncodingKey: "VP8",
		FormatProfileKey:       nil,
		FormatAudioEncodingKey: "vorbis",
	},
	46: {
		FormatExtensionKey:     "webm",
		FormatResolutionKey:    "1080p",
		FormatVideoEncodingKey: "VP8",
		FormatProfileKey:       nil,
		FormatAudioEncodingKey: "vorbis",
	},
	82: {
		FormatExtensionKey:     "mp4",
		FormatResolutionKey:    "360p",
		FormatVideoEncodingKey: "H.264",
		FormatProfileKey:       "3d",
	},
	83: {
		FormatExtensionKey:     "mp4",
		FormatResolutionKey:    "240p",
		FormatVideoEncodingKey: "H.264",
		FormatProfileKey:       "3d",
		FormatAudioEncodingKey: "aac",
	},
	84: {
		FormatExtensionKey:     "mp4",
		FormatResolutionKey:    "720p",
		FormatVideoEncodingKey: "H.264",
		FormatProfileKey:       "3d",
		FormatAudioEncodingKey: "aac",
	},
	85: {
		FormatExtensionKey:     "mp4",
		FormatResolutionKey:    "1080p",
		FormatVideoEncodingKey: "H.264",
		FormatProfileKey:       "3d",
		FormatAudioEncodingKey: "aac",
	},
	100: {
		FormatExtensionKey:     "webm",
		FormatResolutionKey:    "360p",
		FormatVideoEncodingKey: "VP8",
		FormatProfileKey:       "3d",
		FormatAudioEncodingKey: "vorbis",
	},
	101: {
		FormatExtensionKey:     "webm",
		FormatResolutionKey:    "360p",
		FormatVideoEncodingKey: "VP8",
		FormatProfileKey:       "3d",
		FormatAudioEncodingKey: "vorbis",
	},
	102: {
		FormatExtensionKey:     "webm",
		FormatResolutionKey:    "720p",
		FormatVideoEncodingKey: "VP8",
		FormatProfileKey:       "3d",
		FormatAudioEncodingKey: "vorbis",
	},
	// DASH (video only)
	133: {
		FormatExtensionKey:     "mp4",
		FormatResolutionKey:    "240p",
		FormatVideoEncodingKey: "H.264",
		FormatProfileKey:       "main",
		FormatAudioEncodingKey: nil,
	},
	134: {
		FormatExtensionKey:     "mp4",
		FormatResolutionKey:    "360p",
		FormatVideoEncodingKey: "H.264",
		FormatProfileKey:       "main",
		FormatAudioEncodingKey: nil,
	},
	135: {
		FormatExtensionKey:     "mp4",
		FormatResolutionKey:    "480p",
		FormatVideoEncodingKey: "H.264",
		FormatProfileKey:       "main",
		FormatAudioEncodingKey: nil,
	},
	136: {
		FormatExtensionKey:     "mp4",
		FormatResolutionKey:    "720p",
		FormatVideoEncodingKey: "H.264",
		FormatProfileKey:       "main",
		FormatAudioEncodingKey: nil,
	},
	137: {
		FormatExtensionKey:     "mp4",
		FormatResolutionKey:    "1080p",
		FormatVideoEncodingKey: "H.264",
		FormatProfileKey:       "high",
		FormatAudioEncodingKey: nil,
	},
	138: {
		FormatExtensionKey:     "mp4",
		FormatResolutionKey:    "2160p",
		FormatVideoEncodingKey: "H.264",
		FormatProfileKey:       "high",
		FormatAudioEncodingKey: nil,
	},
	160: {
		FormatExtensionKey:     "mp4",
		FormatResolutionKey:    "144p",
		FormatVideoEncodingKey: "H.264",
		FormatProfileKey:       "main",
		FormatAudioEncodingKey: nil,
	},
	242: {
		FormatExtensionKey:     "webm",
		FormatResolutionKey:    "240p",
		FormatVideoEncodingKey: "VP9",
		FormatProfileKey:       nil,
		FormatAudioEncodingKey: nil,
	},
	243: {
		FormatExtensionKey:     "webm",
		FormatResolutionKey:    "360p",
		FormatVideoEncodingKey: "VP9",
		FormatProfileKey:       nil,
		FormatAudioEncodingKey: nil,
	},
	244: {
		FormatExtensionKey:     "webm",
		FormatResolutionKey:    "480p",
		FormatVideoEncodingKey: "VP9",
		FormatProfileKey:       nil,
		FormatAudioEncodingKey: nil,
	},
	247: {
		FormatExtensionKey:     "webm",
		FormatResolutionKey:    "720p",
		FormatVideoEncodingKey: "VP9",
		FormatProfileKey:       nil,
		FormatAudioEncodingKey: nil,
	},
	248: {
		FormatExtensionKey:     "webm",
		FormatResolutionKey:    "1080p",
		FormatVideoEncodingKey: "VP9",
		FormatProfileKey:       nil,
		FormatAudioEncodingKey: nil,
	},
	264: {
		FormatExtensionKey:     "mp4",
		FormatResolutionKey:    "1440p",
		FormatVideoEncodingKey: "H.264",
		FormatProfileKey:       "high",
		FormatAudioEncodingKey: nil,
	},
	266: {
		FormatExtensionKey:     "mp4",
		FormatResolutionKey:    "2160p",
		FormatVideoEncodingKey: "H.264",
		FormatProfileKey:       "high",
		FormatAudioEncodingKey: nil,
	},
	271: {
		FormatExtensionKey:     "webm",
		FormatResolutionKey:    "1440p",
		FormatVideoEncodingKey: "VP9",
		FormatProfileKey:       nil,
		FormatAudioEncodingKey: nil,
	},
	272: {
		FormatExtensionKey:     "webm",
		FormatResolutionKey:    "2160p",
		FormatVideoEncodingKey: "VP9",
		FormatProfileKey:       nil,
		FormatAudioEncodingKey: nil,
	},
	278: {
		FormatExtensionKey:     "webm",
		FormatResolutionKey:    "144p",
		FormatVideoEncodingKey: "VP9",
		FormatProfileKey:       nil,
		FormatAudioEncodingKey: nil,
	},
	298: {
		FormatExtensionKey:     "mp4",
		FormatResolutionKey:    "720p",
		FormatVideoEncodingKey: "H.264",
		FormatProfileKey:       "main",
		FormatAudioEncodingKey: nil,
	},
	299: {
		FormatExtensionKey:     "mp4",
		FormatResolutionKey:    "1080p",
		FormatVideoEncodingKey: "H.264",
		FormatProfileKey:       "high",
		FormatAudioEncodingKey: nil,
	},
	302: {
		FormatExtensionKey:     "webm",
		FormatResolutionKey:    "720p",
		FormatVideoEncodingKey: "VP9",
		FormatProfileKey:       nil,
		FormatAudioEncodingKey: nil,
	},
	303: {
		FormatExtensionKey:     "webm",
		FormatResolutionKey:    "1080p",
		FormatVideoEncodingKey: "VP9",
		FormatProfileKey:       nil,
		FormatAudioEncodingKey: nil,
	},
	// DASH (audio only)
	139: {
		FormatExtensionKey:     "mp4",
		FormatResolutionKey:    nil,
		FormatVideoEncodingKey: nil,
		FormatProfileKey:       nil,
		FormatAudioEncodingKey: "aac",
	},
	140: {
		FormatExtensionKey:     "mp4",
		FormatResolutionKey:    nil,
		FormatVideoEncodingKey: nil,
		FormatProfileKey:       nil,
		FormatAudioEncodingKey: "aac",
	},
	141: {
		FormatExtensionKey:     "mp4",
		FormatResolutionKey:    nil,
		FormatVideoEncodingKey: nil,
		FormatProfileKey:       nil,
		FormatAudioEncodingKey: "aac",
	},
	171: {
		FormatExtensionKey:     "webm",
		FormatResolutionKey:    nil,
		FormatVideoEncodingKey: nil,
		FormatProfileKey:       nil,
		FormatAudioEncodingKey: "vorbis",
	},
	172: {
		FormatExtensionKey:     "webm",
		FormatResolutionKey:    nil,
		FormatVideoEncodingKey: nil,
		FormatProfileKey:       nil,
		FormatAudioEncodingKey: "vorbis",
	},
	249: {
		FormatExtensionKey:     "webm",
		FormatResolutionKey:    nil,
		FormatVideoEncodingKey: nil,
		FormatProfileKey:       nil,
		FormatAudioEncodingKey: "opus",
	},
	250: {
		FormatExtensionKey:     "webm",
		FormatResolutionKey:    nil,
		FormatVideoEncodingKey: nil,
		FormatProfileKey:       nil,
		FormatAudioEncodingKey: "opus",
	},
	251: {
		FormatExtensionKey:     "webm",
		FormatResolutionKey:    nil,
		FormatVideoEncodingKey: nil,
		FormatProfileKey:       nil,
		FormatAudioEncodingKey: "opus",
	},
	// Live streaming
	92: {
		FormatExtensionKey:     "ts",
		FormatResolutionKey:    "240p",
		FormatVideoEncodingKey: "H.264",
		FormatProfileKey:       "main",
		FormatAudioEncodingKey: "aac",
	},
	93: {
		FormatExtensionKey:     "ts",
		FormatResolutionKey:    "480p",
		FormatVideoEncodingKey: "H.264",
		FormatProfileKey:       "main",
		FormatAudioEncodingKey: "aac",
	},
	94: {
		FormatExtensionKey:     "ts",
		FormatResolutionKey:    "720p",
		FormatVideoEncodingKey: "H.264",
		FormatProfileKey:       "main",
		FormatAudioEncodingKey: "aac",
	},
	95: {
		FormatExtensionKey:     "ts",
		FormatResolutionKey:    "1080p",
		FormatVideoEncodingKey: "H.264",
		FormatProfileKey:       "main",
		FormatAudioEncodingKey: "aac",
	},
	96: {
		FormatExtensionKey:     "ts",
		FormatResolutionKey:    "720p",
		FormatVideoEncodingKey: "H.264",
		FormatProfileKey:       "main",
		FormatAudioEncodingKey: "aac",
	},
	120: {
		FormatExtensionKey:     "flv",
		FormatResolutionKey:    "720p",
		FormatVideoEncodingKey: "H.264",
		FormatProfileKey:       "Main@L3.1",
		FormatAudioEncodingKey: "aac",
	},
	127: {
		FormatExtensionKey:     "ts",
		FormatResolutionKey:    nil,
		FormatVideoEncodingKey: nil,
		FormatProfileKey:       nil,
		FormatAudioEncodingKey: "aac",
	},
	128: {
		FormatExtensionKey:     "ts",
		FormatResolutionKey:    nil,
		FormatVideoEncodingKey: nil,
		FormatProfileKey:       nil,
		FormatAudioEncodingKey: "aac",
	},
	132: {
		FormatExtensionKey:     "ts",
		FormatResolutionKey:    "240p",
		FormatVideoEncodingKey: "H.264",
		FormatProfileKey:       "baseline",
		FormatAudioEncodingKey: "aac",
	},
	151: {
		FormatExtensionKey:     "ts",
		FormatResolutionKey:    "720p",
		FormatVideoEncodingKey: "H.264",
		FormatProfileKey:       "baseline",
		FormatAudioEncodingKey: "aac",
	},
}