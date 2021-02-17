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

func (d Disposition) ToMap() map[string]int {
	m := make(map[string]int)
	d.addMapEntry(m, "default", AV_DISPOSITION_DEFAULT)
	d.addMapEntry(m, "dub", AV_DISPOSITION_DUB)
	d.addMapEntry(m, "original", AV_DISPOSITION_ORIGINAL)
	d.addMapEntry(m, "comment", AV_DISPOSITION_COMMENT)
	d.addMapEntry(m, "lyrics", AV_DISPOSITION_LYRICS)
	d.addMapEntry(m, "karaoke", AV_DISPOSITION_KARAOKE)
	d.addMapEntry(m, "forced", AV_DISPOSITION_FORCED)
	d.addMapEntry(m, "hearing_impaired", AV_DISPOSITION_HEARING_IMPAIRED)
	d.addMapEntry(m, "visual_impaired", AV_DISPOSITION_VISUAL_IMPAIRED)
	d.addMapEntry(m, "clean_effects", AV_DISPOSITION_CLEAN_EFFECTS)
	d.addMapEntry(m, "attached_pic", AV_DISPOSITION_ATTACHED_PIC)
	d.addMapEntry(m, "timed_thumbnails", AV_DISPOSITION_TIMED_THUMBNAILS)

	return m
}

func (d Disposition) addMapEntry(m map[string]int, name string, mask Disposition) {
	value := 0
	if d&mask != 0 {
		value = 1
	}
	m[name] = value
}
