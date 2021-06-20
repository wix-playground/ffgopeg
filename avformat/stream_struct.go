// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)
// Corbatto (luca@corbatto.de)

package avformat

//#cgo pkg-config: libavformat
//#include <libavformat/avformat.h>
//#include <libavutil/rational.h>
import "C"
import (
	"unsafe"

	"github.com/wix-playground/ffgopeg/avcodec"
)

// C-Variable: AVStream::codec
func (s *Stream) CodecContext() *avcodec.CodecContext {
	return (*avcodec.CodecContext)(unsafe.Pointer(s.codec))
}

// C-Variable: AVStream::metadata
func (s *Stream) Metadata() *Dictionary {
	return (*Dictionary)(unsafe.Pointer(s.metadata))
}

// C-Variable: AVStream::index_entries
func (s *Stream) IndexEntries() *IndexEntry {
	return (*IndexEntry)(unsafe.Pointer(s.index_entries))
}

// C-Variable: AVStream::attached_pic
func (s *Stream) AttachedPic() Packet {
	return Packet(s.attached_pic)
}

// C-Variable: AVStream::avg_frame_rate
func (s *Stream) AvgFrameRate() Rational {
	return Rational(s.avg_frame_rate)
}

// C-Variable: AVStream::r_frame_rate
func (s *Stream) RFrameRate() Rational {
	return Rational(s.r_frame_rate)
}

// C-Variable: AVStream::sample_aspect_ratio
func (s *Stream) SampleAspectRatio() Rational {
	return Rational(s.sample_aspect_ratio)
}

// C-Variable: AVStream::time_base
func (s *Stream) TimeBase() Rational {
	return Rational(s.time_base)
}

// C-Variable: AVStream::discard
func (s *Stream) Discard() Discard {
	return Discard(s.discard)
}

// C-Variable: AVStream::need_parsing
func (s *Stream) NeedParsing() StreamParseType {
	return StreamParseType(s.need_parsing)
}

// C-Variable: AVStream::codec_info_nb_frames
func (s *Stream) CodecInfoNbFrames() int {
	return int(s.codec_info_nb_frames)
}

// C-Variable: AVStream::disposition
func (s *Stream) Disposition() Disposition {
	return Disposition(s.disposition)
}

// C-Variable: AVStream::event_flags
func (s *Stream) EventFlags() int {
	return int(s.event_flags)
}

// C-Variable: AVStream::id
func (s *Stream) Id() int {
	return int(s.id)
}

// C-Variable: AVStream::index
func (s *Stream) Index() int {
	return int(s.index)
}

// C-Variable: AVStream::last_
func (s *Stream) LastIpDuration() int {
	return int(s.last_IP_duration)
}

// C-Variable: AVStream::nb_index_entries
func (s *Stream) NbIndexEntries() int {
	return int(s.nb_index_entries)
}

// C-Variable: AVStream::nb_side_data
func (s *Stream) NbSideData() int {
	return int(s.nb_side_data)
}

// C-Variable: AVStream::probe_packets
func (s *Stream) ProbePackets() int {
	return int(s.probe_packets)
}

//
// C-Variable: AVStream::stream_identifier
func (s *Stream) StreamIdentifier() int {
	return int(s.stream_identifier)
}

//
// C-Variable: AVStream::duration
func (s *Stream) Duration() int64 {
	return int64(s.duration)
}

//
// C-Variable: AVStream::last_
func (s *Stream) LastIpPts() int64 {
	return int64(s.last_IP_pts)
}

//
// C-Variable: AVStream::nb_frames
func (s *Stream) NbFrames() int64 {
	return int64(s.nb_frames)
}

//
// C-Variable: AVStream::start_time
func (s *Stream) StartTime() int64 {
	return int64(s.start_time)
}

//
// C-Variable: AVStream::parser
func (s *Stream) Parser() *CodecParserContext {
	return (*CodecParserContext)(unsafe.Pointer(s.parser))
}

//
// C-Variable: AVStream::index_entries_allocated_size
func (s *Stream) IndexEntriesAllocatedSize() uint {
	return uint(s.index_entries_allocated_size)
}

//
// C-Variable: AVStream::codecpar
func (s *Stream) CodecPar() *avcodec.CodecParameters {
	return (*avcodec.CodecParameters)(unsafe.Pointer(s.codecpar))
}
