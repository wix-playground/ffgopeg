package avformat

//#cgo CFLAGS: -Wno-deprecated-declarations
//#cgo pkg-config: libavcodec
//#include <libavformat/avformat.h>
import "C"

type Disposition int

const (
	AV_DISPOSITION_DEFAULT          = Disposition(C.AV_DISPOSITION_DEFAULT)
	AV_DISPOSITION_DUB              = Disposition(C.AV_DISPOSITION_DUB)
	AV_DISPOSITION_ORIGINAL         = Disposition(C.AV_DISPOSITION_ORIGINAL)
	AV_DISPOSITION_COMMENT          = Disposition(C.AV_DISPOSITION_COMMENT)
	AV_DISPOSITION_LYRICS           = Disposition(C.AV_DISPOSITION_LYRICS)
	AV_DISPOSITION_KARAOKE          = Disposition(C.AV_DISPOSITION_KARAOKE)
	AV_DISPOSITION_FORCED           = Disposition(C.AV_DISPOSITION_FORCED)
	AV_DISPOSITION_HEARING_IMPAIRED = Disposition(C.AV_DISPOSITION_HEARING_IMPAIRED)
	AV_DISPOSITION_VISUAL_IMPAIRED  = Disposition(C.AV_DISPOSITION_VISUAL_IMPAIRED)
	AV_DISPOSITION_CLEAN_EFFECTS    = Disposition(C.AV_DISPOSITION_CLEAN_EFFECTS)
	AV_DISPOSITION_ATTACHED_PIC     = Disposition(C.AV_DISPOSITION_ATTACHED_PIC)
	AV_DISPOSITION_TIMED_THUMBNAILS = Disposition(C.AV_DISPOSITION_TIMED_THUMBNAILS)
)

var dispositionMaskNames = map[Disposition]string{
	AV_DISPOSITION_DEFAULT:          "default",
	AV_DISPOSITION_DUB:              "dub",
	AV_DISPOSITION_ORIGINAL:         "original",
	AV_DISPOSITION_COMMENT:          "comment",
	AV_DISPOSITION_LYRICS:           "lyrics",
	AV_DISPOSITION_KARAOKE:          "karaoke",
	AV_DISPOSITION_FORCED:           "forced",
	AV_DISPOSITION_HEARING_IMPAIRED: "hearing_impaired",
	AV_DISPOSITION_VISUAL_IMPAIRED:  "visual_impaired",
	AV_DISPOSITION_CLEAN_EFFECTS:    "clean_effects",
	AV_DISPOSITION_ATTACHED_PIC:     "attached_pic",
	AV_DISPOSITION_TIMED_THUMBNAILS: "timed_thumbnails",
}

func (d Disposition) Strings() (strings []string) {
	for mask, name := range dispositionMaskNames {
		if d&mask != 0 {
			strings = append(strings, name)
		}
	}

	return strings
}

func (d Disposition) addString(l []string, name string, mask Disposition) []string {
	if d&mask != 0 {
		l = append(l, name)
	}
	return l
}
