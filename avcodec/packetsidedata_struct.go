// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)
// Corbatto (luca@corbatto.de)

package avcodec

import "unsafe"

// Data returns the data.
//
// C-Field: AVPacketSideData::data
func (d *PacketSideData) Data() []uint8 {
	return (*[1 << 30]uint8)(unsafe.Pointer(d.data))[:d.Size():d.Size()]
}

// Size returns size of the data.
//
// C-Field: AVPacketSideData::size
func (d *PacketSideData) Size() int {
	return int(d.size)
}

// Type returns the type of the side data.
func (d *PacketSideData) Type() PacketSideDataType {
	return PacketSideDataType(d._type)
}
