// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)
// Corbatto (luca@corbatto.de)

package avformat

//#cgo CFLAGS: -Wno-deprecated-declarations
//#cgo pkg-config: libavutil
//#include <libavutil/dict.h>
//#include <stdlib.h>
import "C"

// Count returns the number of entries in dictionary.
//
// C-Function: av_dict_count
func (d *Dictionary) Count() int {
	return int(C.av_dict_count((*C.struct_AVDictionary)(d)))
}

// Get returns a dictionary entry with matching key.
//
// C-Function: av_dict_count
func (d *Dictionary) Get(key string) *DictionaryEntry {
	var entry *DictionaryEntry
	entry = (*DictionaryEntry)(C.av_dict_get((*C.struct_AVDictionary)(d), C.CString(key), (*C.struct_AVDictionaryEntry)(entry), C.AV_DICT_MATCH_CASE))

	return entry
}

var emptyString = C.CString("")

func (d *Dictionary) ToMap() (m map[string]string) {
	var entry *DictionaryEntry
	for {
		entry = (*DictionaryEntry)(C.av_dict_get((*C.struct_AVDictionary)(d), emptyString,
			(*C.struct_AVDictionaryEntry)(entry), C.AV_DICT_IGNORE_SUFFIX))

		if entry == nil {
			break
		}

		if m == nil {
			m = map[string]string{}
		}

		m[entry.Key()] = entry.Value()
	}
	return
}

// Key returns the entry value.
//
// C-Variable: AVDictionaryEntry::value
func (d *DictionaryEntry) Key() string {
	return C.GoString(d.key)
}

// Value returns the entry value.
//
// C-Variable: AVDictionaryEntry::value
func (d *DictionaryEntry) Value() string {
	return C.GoString(d.value)
}
