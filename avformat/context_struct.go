// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)
// Corbatto (luca@corbatto.de)

package avformat

//#cgo CFLAGS: -Wno-deprecated-declarations
//#cgo pkg-config: libavformat
//#include <libavformat/avformat.h>
import "C"
import "unsafe"

// Chapters returns chapters.
//
// C-Variable: AVFormatContext::chapters
func (c *FormatContext) Chapters() []*Chapter {
	return (*[1 << 30]*Chapter)(unsafe.Pointer(c.chapters))[:c.NbChapters():c.NbChapters()]
}

//
// C-Variable: AVFormatContext::metadata
func (c *FormatContext) Metadata() *Dictionary {
	return (*Dictionary)(unsafe.Pointer(c.metadata))
}

//
// C-Variable: AVFormatContext::internal
func (c *FormatContext) Internal() *FormatInternal {
	return (*FormatInternal)(unsafe.Pointer(c.internal))
}

//
// C-Variable: AVFormatContext::pb
func (c *FormatContext) Pb() *IOContext {
	return (*IOContext)(unsafe.Pointer(c.pb))
}

//
// C-Variable: AVFormatContext::interrupt_callback
func (c *FormatContext) InterruptCallback() IOInterruptCB {
	return IOInterruptCB(c.interrupt_callback)
}

//
// C-Variable: AVFormatContext::programs
func (c *FormatContext) Programs() **Program {
	return (**Program)(unsafe.Pointer(c.programs))
}

//
// C-Variable: AVFormatContext::streams
func (c *FormatContext) Streams() []*Stream {
	length := c.NbStreams()
	return (*[1 << 30]*Stream)(unsafe.Pointer(c.streams))[:length:length]
}

//
// C-Variable: AVFormatContext::filename
func (c *FormatContext) Filename() string {
	return C.GoString((*C.char)(unsafe.Pointer(&c.filename[0])))
}

// //
// C-Variable: AVFormatContext::codec_whitelist
func (c *FormatContext) CodecWhitelist() string {
	return C.GoString(c.codec_whitelist)
}

// //
// C-Variable: AVFormatContext::format_whitelist
func (c *FormatContext) FormatWhitelist() string {
	return C.GoString(c.format_whitelist)
}

//
// C-Variable: AVFormatContext::audio_codec_id
func (c *FormatContext) AudioCodecId() CodecId {
	return CodecId(c.audio_codec_id)
}

//
// C-Variable: AVFormatContext::subtitle_codec_id
func (c *FormatContext) SubtitleCodecId() CodecId {
	return CodecId(c.subtitle_codec_id)
}

//
// C-Variable: AVFormatContext::video_codec_id
func (c *FormatContext) VideoCodecId() CodecId {
	return CodecId(c.video_codec_id)
}

//
// C-Variable: AVFormatContext::audio_preload
func (c *FormatContext) AudioPreload() int {
	return int(c.audio_preload)
}

//
// C-Variable: AVFormatContext::avio_flags
func (c *FormatContext) AvioFlags() int {
	return int(c.avio_flags)
}

//
// C-Variable: AVFormatContext::avoid_negative_ts
func (c *FormatContext) AvoidNegativeTs() int {
	return int(c.avoid_negative_ts)
}

//
// C-Variable: AVFormatContext::bit_rate
func (c *FormatContext) BitRate() int {
	return int(c.bit_rate)
}

//
// C-Variable: AVFormatContext::ctx_flags
func (c *FormatContext) CtxFlags() int {
	return int(c.ctx_flags)
}

//
// C-Variable: AVFormatContext::debug
func (c *FormatContext) Debug() int {
	return int(c.debug)
}

//
// C-Variable: AVFormatContext::error_recognition
func (c *FormatContext) ErrorRecognition() int {
	return int(c.error_recognition)
}

//
// C-Variable: AVFormatContext::event_flags
func (c *FormatContext) EventFlags() int {
	return int(c.event_flags)
}

//
// C-Variable: AVFormatContext::flags
func (c *FormatContext) Flags() int {
	return int(c.flags)
}

//
// C-Variable: AVFormatContext::flush_packets
func (c *FormatContext) FlushPackets() int {
	return int(c.flush_packets)
}

//
// C-Variable: AVFormatContext::format_probesize
func (c *FormatContext) FormatProbesize() int {
	return int(c.format_probesize)
}

//
// C-Variable: AVFormatContext::fps_probe_size
func (c *FormatContext) FpsProbeSize() int {
	return int(c.fps_probe_size)
}

//
// C-Variable: AVFormatContext::io_repositioned
func (c *FormatContext) IoRepositioned() int {
	return int(c.io_repositioned)
}

//
// C-Variable: AVFormatContext::keylen
func (c *FormatContext) Keylen() int {
	return int(c.keylen)
}

//
// C-Variable: AVFormatContext::max_chunk_duration
func (c *FormatContext) MaxChunkDuration() int {
	return int(c.max_chunk_duration)
}

//
// C-Variable: AVFormatContext::max_chunk_size
func (c *FormatContext) MaxChunkSize() int {
	return int(c.max_chunk_size)
}

//
// C-Variable: AVFormatContext::max_delay
func (c *FormatContext) MaxDelay() int {
	return int(c.max_delay)
}

//
// C-Variable: AVFormatContext::max_ts_probe
func (c *FormatContext) MaxTsProbe() int {
	return int(c.max_ts_probe)
}

//
// C-Variable: AVFormatContext::seek2any
func (c *FormatContext) Seek2any() int {
	return int(c.seek2any)
}

//
// C-Variable: AVFormatContext::strict_std_compliance
func (c *FormatContext) StrictStdCompliance() int {
	return int(c.strict_std_compliance)
}

//
// C-Variable: AVFormatContext::ts_id
func (c *FormatContext) TsId() int {
	return int(c.ts_id)
}

//
// C-Variable: AVFormatContext::use_wallclock_as_timestamps
func (c *FormatContext) UseWallclockAsTimestamps() int {
	return int(c.use_wallclock_as_timestamps)
}

//
// C-Variable: AVFormatContext::duration
func (c *FormatContext) Duration() int64 {
	return int64(c.duration)
}

//
// C-Variable: AVFormatContext::max_analyze_duration
func (c *FormatContext) MaxAnalyzeDuration2() int64 {
	return int64(c.max_analyze_duration)
}

//
// C-Variable: AVFormatContext::max_interleave_delta
func (c *FormatContext) MaxInterleaveDelta() int64 {
	return int64(c.max_interleave_delta)
}

//
// C-Variable: AVFormatContext::output_ts_offset
func (c *FormatContext) OutputTsOffset() int64 {
	return int64(c.output_ts_offset)
}

//
// C-Variable: AVFormatContext::probesize
func (c *FormatContext) Probesize2() int64 {
	return int64(c.probesize)
}

//
// C-Variable: AVFormatContext::skip_initial_bytes
func (c *FormatContext) SkipInitialBytes() int64 {
	return int64(c.skip_initial_bytes)
}

//
// C-Variable: AVFormatContext::start_time
func (c *FormatContext) StartTime() int64 {
	return int64(c.start_time)
}

//
// C-Variable: AVFormatContext::start_time_realtime
func (c *FormatContext) StartTimeRealtime() int64 {
	return int64(c.start_time_realtime)
}

//
// C-Variable: AVFormatContext::iformat
func (c *FormatContext) Iformat() *InputFormat {
	return (*InputFormat)(unsafe.Pointer(c.iformat))
}

//
// C-Variable: AVFormatContext::oformat
func (c *FormatContext) Oformat() *OutputFormat {
	return (*OutputFormat)(unsafe.Pointer(c.oformat))
}

// //
// C-Variable: AVFormatContext::dump_separator
func (c *FormatContext) DumpSeparator() uint8 {
	return uint8(*c.dump_separator)
}

//
// C-Variable: AVFormatContext::correct_ts_overflow
func (c *FormatContext) CorrectTsOverflow() int {
	return int(c.correct_ts_overflow)
}

//
// C-Variable: AVFormatContext::max_index_size
func (c *FormatContext) MaxIndexSize() uint {
	return uint(c.max_index_size)
}

//
// C-Variable: AVFormatContext::max_picture_buffer
func (c *FormatContext) MaxPictureBuffer() uint {
	return uint(c.max_picture_buffer)
}

//
// C-Variable: AVFormatContext::nb_chapters
func (c *FormatContext) NbChapters() uint {
	return uint(c.nb_chapters)
}

//
// C-Variable: AVFormatContext::nb_programs
func (c *FormatContext) NbPrograms() uint {
	return uint(c.nb_programs)
}

//
// C-Variable: AVFormatContext::nb_streams
func (c *FormatContext) NbStreams() uint {
	return uint(c.nb_streams)
}

//
// C-Variable: AVFormatContext::packet_size
func (c *FormatContext) PacketSize() uint {
	return uint(c.packet_size)
}

//
// C-Variable: AVFormatContext::probesize
func (c *FormatContext) Probesize() uint {
	return uint(c.probesize)
}
