// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)
// Corbatto (luca@corbatto.de)

package avformat

//#cgo CFLAGS: -Wno-deprecated-declarations
//#cgo pkg-config: libavformat
//#include <libavformat/avformat.h>
import "C"
import (
	"unsafe"

	"github.com/wix-playground/ffgopeg/avcodec"
	"github.com/wix-playground/ffgopeg/avutil"
)

// ProbeScore returns the probe score.
//
// C-Function: av_format_get_probe_score
func (c *FormatContext) ProbeScore() int {
	return int(C.av_format_get_probe_score((*C.struct_AVFormatContext)(c)))
}

// VideoCodec returns the VideoCodec.
//
// C-Function: av_format_get_video_codec
func (c *FormatContext) VideoCodec() *Codec {
	return (*Codec)(C.av_format_get_video_codec((*C.struct_AVFormatContext)(c)))
}

// SetVideoCodec sets the video codec.
//
// C-Function: av_format_set_video_codec
func (c *FormatContext) SetVideoCodec(codec *Codec) {
	C.av_format_set_video_codec((*C.struct_AVFormatContext)(c), (*C.struct_AVCodec)(codec))
}

// AudioCodec returns the audio codec.
//
// C-Function: av_format_get_audio_codec
func (c *FormatContext) AudioCodec() *Codec {
	return (*Codec)(C.av_format_get_audio_codec((*C.struct_AVFormatContext)(c)))
}

// SetAudioCodec sets the audio codec.
//
// C-Function: av_format_set_audio_codec
func (c *FormatContext) SetAudioCodec(codec *Codec) {
	C.av_format_set_audio_codec((*C.struct_AVFormatContext)(c), (*C.struct_AVCodec)(codec))
}

// SubtitleCodec returns the subtitle codec.
//
// C-Function: av_format_get_subtitle_codec
func (c *FormatContext) SubtitleCodec() *Codec {
	return (*Codec)(C.av_format_get_subtitle_codec((*C.struct_AVFormatContext)(c)))
}

// SetSubtitleCodec sets the subtitle codec.
//
// C-Function: av_format_set_subtitle_codec
func (c *FormatContext) SetSubtitleCodec(codec *Codec) {
	C.av_format_set_subtitle_codec((*C.struct_AVFormatContext)(c), (*C.struct_AVCodec)(codec))
}

// MetadataHeaderPadding returns metadata header padding size.
//
// C-Function: av_format_get_metadata_header_padding
func (c *FormatContext) MetadataHeaderPadding() int {
	return int(C.av_format_get_metadata_header_padding((*C.struct_AVFormatContext)(c)))
}

// SetMetadataHeaderPadding sets metadata header padding size.
//
// C-Function: av_format_set_metadata_header_padding
func (c *FormatContext) SetMetadataHeaderPadding(size int) {
	C.av_format_set_metadata_header_padding((*C.struct_AVFormatContext)(c), C.int(size))
}

// InjectGlobalSideData will cause global side data to be injected in the next packet of each stream as well as after any subsequent seek.
//
// C-Function: av_format_inject_global_side_data
func (c *FormatContext) InjectGlobalSideData() {
	C.av_format_inject_global_side_data((*C.struct_AVFormatContext)(c))
}

// DurationEstimationMethod returns the method used to set ctx->duration.
//
// C-Function: av_fmt_ctx_get_duration_estimation_method
func (c *FormatContext) DurationEstimationMethod() DurationEstimationMethod {
	return (DurationEstimationMethod)(C.av_fmt_ctx_get_duration_estimation_method((*C.struct_AVFormatContext)(c)))
}

// Free frees the Context and all its streams.
//
// C-Function: avformat_free_context
func (c *FormatContext) Free() {
	C.avformat_free_context((*C.struct_AVFormatContext)(c))
}

// NewStream adds a new stream to a media file.
//
// C-Function: avformat_new_stream
func (c *FormatContext) NewStream(codec *Codec) *Stream {
	return (*Stream)(C.avformat_new_stream((*C.struct_AVFormatContext)(c), (*C.struct_AVCodec)(codec)))
}

// NewProgram creates a new Program.
//
// C-Function: av_new_program
func (c *FormatContext) NewProgram(id int) *Program {
	return (*Program)(C.av_new_program((*C.struct_AVFormatContext)(c), C.int(id)))
}

// FindStreamInfo reads packets of a media file to get stream information.
//
// C-Function: avformat_find_stream_info
func (c *FormatContext) FindStreamInfo(d **Dictionary) avutil.ReturnCode {
	return avutil.NewReturnCode(int(C.avformat_find_stream_info((*C.struct_AVFormatContext)(c), (**C.struct_AVDictionary)(unsafe.Pointer(d)))))
}

func (c *FormatContext) findProgramFromStream(l *Program, su int) *Program {
	return (*Program)(C.av_find_program_from_stream((*C.struct_AVFormatContext)(c), (*C.struct_AVProgram)(l), C.int(su)))
}

// IterateProgramsFromStream returns a channel which can be used to iterate over the Programs from the given Stream.
//
// Usage:
//
//     for i, prog := range fmtCtxt.IterateProgramsFromStream(0) {
//         // ...
//     }
//
// C-Function: av_find_program_from_stream
func (c *FormatContext) IterateProgramsFromStream(streamIndex int) <-chan *Program {
	ch := make(chan *Program)
	v := c.findProgramFromStream(nil, streamIndex)
	go func() {
		for v != nil {
			ch <- v
			v = c.findProgramFromStream(v, streamIndex)
		}
		close(ch)
	}()

	return ch
}

// FindBestStream finds the "best" stream in the file.
//
// C-Function: av_find_best_stream
func FindBestStream(ic *FormatContext, t avutil.MediaType, ws, rs int, c **Codec, f int) int {
	return int(C.av_find_best_stream((*C.struct_AVFormatContext)(ic), (C.enum_AVMediaType)(t), C.int(ws), C.int(rs), (**C.struct_AVCodec)(unsafe.Pointer(c)), C.int(f)))
}

// ReadFrame returns the next frame of a stream.
//
// C-Function: av_read_frame
func (c *FormatContext) ReadFrame(pkt *avcodec.Packet) avutil.ReturnCode {
	return avutil.NewReturnCode(int(C.av_read_frame((*C.struct_AVFormatContext)(unsafe.Pointer(c)), (*C.struct_AVPacket)(unsafe.Pointer(pkt)))))
}

// SeekFrame seeks to the keyframe at timestamp.
//
// C-Function: av_seek_frame
func (c *FormatContext) SeekFrame(stream int, timestamp int64, flags int) avutil.ReturnCode {
	return avutil.NewReturnCode(int(C.av_seek_frame((*C.struct_AVFormatContext)(c), C.int(stream), C.int64_t(timestamp), C.int(flags))))
}

// SeekFile seeks to timestamp.
//
// C-Function: avformat_seek_file
func (c *FormatContext) SeekFile(stream int, minTS, timestamp, maxTS int64, flags int) int {
	return int(C.avformat_seek_file((*C.struct_AVFormatContext)(c), C.int(stream), C.int64_t(minTS), C.int64_t(timestamp), C.int64_t(maxTS), C.int(flags)))
}

// ReadPlay starts playing a network-based stream (e.g.
//
// C-Function: av_read_play
func (c *FormatContext) ReadPlay() int {
	return int(C.av_read_play((*C.struct_AVFormatContext)(c)))
}

// ReadPause pauses a network-based stream (e.g.
//
// C-Function: av_read_pause
func (c *FormatContext) ReadPause() int {
	return int(C.av_read_pause((*C.struct_AVFormatContext)(c)))
}

// Close closes an opened input Context.
//
// C-Function: avformat_close_input
func (c *FormatContext) Close() {
	C.avformat_close_input((**C.struct_AVFormatContext)(unsafe.Pointer(&c)))
}

// WriteHeader allocates the stream private data and write the stream header to an output media file.
//
// C-Function: avformat_write_header
func (c *FormatContext) WriteHeader(o **Dictionary) int {
	return int(C.avformat_write_header((*C.struct_AVFormatContext)(c), (**C.struct_AVDictionary)(unsafe.Pointer(o))))
}

// WriteFrame writes a packet to an output media file.
//
// C-Function: av_write_frame
func (c *FormatContext) WriteFrame(pkt *Packet) int {
	return int(C.av_write_frame((*C.struct_AVFormatContext)(c), (*C.struct_AVPacket)(pkt)))
}

// InterleavedWriteFrame writes a packet to an output media file ensuring correct interleaving.
//
// C-Function: av_interleaved_write_frame
func (c *FormatContext) InterleavedWriteFrame(pkt *Packet) int {
	return int(C.av_interleaved_write_frame((*C.struct_AVFormatContext)(c), (*C.struct_AVPacket)(pkt)))
}

// WriteUncodedFrame writes an uncoded frame to an output media file.
//
// C-Function: av_write_uncoded_frame
func (c *FormatContext) WriteUncodedFrame(si int, f *Frame) int {
	return int(C.av_write_uncoded_frame((*C.struct_AVFormatContext)(c), C.int(si), (*C.struct_AVFrame)(f)))
}

// InterleavedWriteUncodedFrame writes a uncoded frame to an output media file.
//
// C-Function: av_interleaved_write_uncoded_frame
func (c *FormatContext) InterleavedWriteUncodedFrame(stream int, frame *Frame) int {
	return int(C.av_interleaved_write_uncoded_frame((*C.struct_AVFormatContext)(c), C.int(stream), (*C.struct_AVFrame)(frame)))
}

// WriteUncodedFrameQuery tests whether a muxer supports uncoded frame.
//
// C-Function: av_write_uncoded_frame_query
func (c *FormatContext) WriteUncodedFrameQuery(stream int) int {
	return int(C.av_write_uncoded_frame_query((*C.struct_AVFormatContext)(c), C.int(stream)))
}

// WriteTrailer writes the stream trailer to an output media file and free the file private data.
//
// C-Function: av_write_trailer
func (c *FormatContext) WriteTrailer() int {
	return int(C.av_write_trailer((*C.struct_AVFormatContext)(c)))
}

// OutputTimestamp returns timing information for the data currently output.
//
// C-Function: av_get_output_timestamp
func (c *FormatContext) OutputTimestamp(stream int) (dts, wall int, err avutil.ReturnCode) {
	err = avutil.NewReturnCode(int(C.av_get_output_timestamp((*C.struct_AVFormatContext)(c), C.int(stream), (*C.int64_t)(unsafe.Pointer(&dts)), (*C.int64_t)(unsafe.Pointer(&wall)))))
	return
}

// FindDefaultStreamIndex returns the default stream index.
//
// C-Function: av_find_default_stream_index
func (c *FormatContext) FindDefaultStreamIndex() avutil.ReturnCode {
	return avutil.NewReturnCode(int(C.av_find_default_stream_index((*C.struct_AVFormatContext)(c))))
}

// DumpFormat prints detailed information about the input or output format, such as duration, bitrate, streams, container, programs, metadata, side data, codec and time base.
//
// C-Function: av_dump_format
func (c *FormatContext) DumpFormat(index int, url string, isOutput bool) {
	io := 0
	if isOutput {
		io = 1
	}
	C.av_dump_format((*C.struct_AVFormatContext)(unsafe.Pointer(c)), C.int(index), C.CString(url), C.int(io))
}

// GuessSampleAspectRatio guesses the sample aspect ratio of a frame, based on both the stream and the frame aspect ratio.
//
// C-Function: av_guess_sample_aspect_ratio
func (c *FormatContext) GuessSampleAspectRatio(st *Stream, fr *Frame) Rational {
	return (Rational)(C.av_guess_sample_aspect_ratio((*C.struct_AVFormatContext)(c), (*C.struct_AVStream)(st), (*C.struct_AVFrame)(fr)))
}

// GuessFrameRate guesses the frame rate, based on both the container and codec information.
//
// C-Function: av_guess_frame_rate
func (c *FormatContext) GuessFrameRate(st *Stream, fr *Frame) Rational {
	return (Rational)(C.av_guess_frame_rate((*C.struct_AVFormatContext)(c), (*C.struct_AVStream)(st), (*C.struct_AVFrame)(fr)))
}

// MatchStreamSpecifier checks if the stream st contained in s is matched by the stream specifier spec.
//
// C-Function: avformat_match_stream_specifier
func (c *FormatContext) MatchStreamSpecifier(st *Stream, spec string) int {
	return int(C.avformat_match_stream_specifier((*C.struct_AVFormatContext)(c), (*C.struct_AVStream)(st), C.CString(spec)))
}

// QueueAttachedPictures is undocumented.
//
// C-Function: avformat_queue_attached_pictures
func (c *FormatContext) QueueAttachedPictures() int {
	return int(C.avformat_queue_attached_pictures((*C.struct_AVFormatContext)(c)))
}
