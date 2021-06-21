// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)
// Corbatto (luca@corbatto.de)

//Package avformat provides some generic global options, which can be set on all the muxers and demuxers.
//In addition each muxer or demuxer may support so-called private options, which are specific for that component.
//Supported formats (muxers and demuxers) provided by the libavformat library
package avformat

//#cgo CFLAGS: -Wno-deprecated-declarations
//#cgo pkg-config: libavformat libavcodec libavutil libavdevice libavfilter libswresample libswscale
//#include <stdio.h>
//#include <stdlib.h>
//#include <inttypes.h>
//#include <stdint.h>
//#include <string.h>
//#include <libavformat/avformat.h>
//#include <libavcodec/avcodec.h>
//#include <libavutil/avutil.h>
//#include <libavutil/opt.h>
//#include <libavutil/rational.h>
//#include <libavdevice/avdevice.h>
import "C"
import (
	"fmt"
	"unsafe"

	"github.com/wix-playground/ffgopeg/avutil"
)

type (
	ProbeData                C.struct_AVProbeData
	InputFormat              C.struct_AVInputFormat
	OutputFormat             C.struct_AVOutputFormat
	FormatContext            C.struct_AVFormatContext
	Frame                    C.struct_AVFrame
	CodecFormatContext       C.struct_AVCodecFormatContext
	CodecParserContext       C.struct_AVCodecParserContext
	Dictionary               C.struct_AVDictionary
	DictionaryEntry          C.struct_AVDictionaryEntry
	IndexEntry               C.struct_AVIndexEntry
	Stream                   C.struct_AVStream
	Program                  C.struct_AVProgram
	Chapter                  C.struct_AVChapter
	PacketList               C.struct_AVPacketList
	Packet                   C.struct_AVPacket
	CodecParserFormatContext C.struct_AVCodecParserFormatContext
	IOContext                C.struct_AVIOContext
	IOFormatContext          C.struct_AVIOFormatContext
	Rational                 C.struct_AVRational
	Codec                    C.struct_AVCodec
	CodecTag                 C.struct_AVCodecTag
	Class                    C.struct_AVClass
	FormatInternal           C.struct_AVFormatInternal
	IOInterruptCB            C.struct_AVIOInterruptCB
	PacketSideData           C.struct_AVPacketSideData
	FFFrac                   C.struct_FFFrac
	StreamParseType          C.enum_AVStreamParseType
	Discard                  C.enum_AVDiscard
	MediaType                C.enum_AVMediaType
	DurationEstimationMethod C.enum_AVDurationEstimationMethod
	PacketSideDataType       C.enum_AVPacketSideDataType
	CodecId                  C.enum_AVCodecID
)

type File C.FILE

// TODO: Find a nice way to convert os.File into C.FILE.

// GetPacket allocates and reads the payload of a packet and initialize its fields with default values.
//
// C-Function: av_get_packet
func (c *IOContext) GetPacket(size int) (*Packet, avutil.ReturnCode) {
	var pkt *Packet
	code := avutil.NewReturnCode(int(C.av_get_packet((*C.struct_AVIOContext)(c), (*C.struct_AVPacket)(pkt), C.int(size))))
	return pkt, code
}

// AppendPacket Read data and append it to the current content of the Packet.
//
// C-Function: av_append_packet
func (c *IOContext) AppendPacket(pkt *Packet, s int) int {
	return int(C.av_append_packet((*C.struct_AVIOContext)(c), (*C.struct_AVPacket)(pkt), C.int(s)))
}

func (c *IOContext) Size() int64 {
	if c == nil {
		return -1
	}

	return int64(C.avio_size((*C.struct_AVIOContext)(c)))
}

func NewRational(num, den int) (rational Rational) {
	rational.num = C.int(num)
	rational.den = C.int(den)
	return
}

func (r Rational) Num() int {
	return int(r.num)
}

func (r Rational) Den() int {
	return int(r.den)
}

func (r Rational) String() string {
	return fmt.Sprintf("%v/%v", r.num, r.den)
}

func (r Rational) RatioString() string {
	return fmt.Sprintf("%v:%v", r.num, r.den)
}

func (r Rational) Float32() float32 {
	return float32(r.num) / float32(r.den)
}

func (r Rational) Float64() float64 {
	return float64(r.num) / float64(r.den)
}

func (r Rational) Reduce(max int64) (result Rational) {
	//int av_reduce(int *dst_num, int *dst_den, int64_t num, int64_t den, int64_t max);
	C.av_reduce(&result.num, &result.den, C.int64_t(r.num), C.int64_t(r.den), C.int64_t(max))
	return
}

// Register registers the InputFormat.
//
// C-Function: av_register_input_format
func (f *InputFormat) Register() {
	C.av_register_input_format((*C.struct_AVInputFormat)(f))
}

// Name returns the short format name
func (f *InputFormat) Name() string {
	return C.GoString((*C.struct_AVInputFormat)(f).name)
}

// LongName returns the human readable format name
func (f *InputFormat) LongName() string {
	return C.GoString((*C.struct_AVInputFormat)(f).long_name)
}

// Register registers the OutputFormat.
//
// C-Function: av_register_output_format
func (f *OutputFormat) Register() {
	C.av_register_output_format((*C.struct_AVOutputFormat)(f))
}

func inputFormatNext(f *InputFormat) *InputFormat {
	return (*InputFormat)(C.av_iformat_next((*C.struct_AVInputFormat)(f)))
}

// InputFormats returns a channel which can be used to iterate over the input formats.
//
// Usage:
//
//     for in := range avformat.InputFormats() {
//         // ...
//     }
//
// C-Function: av_iformat_next
func InputFormats() <-chan *InputFormat {
	ch := make(chan *InputFormat)

	var e *InputFormat
	go func() {
		for {
			e = inputFormatNext(e)
			if e == nil {
				break
			}
			ch <- e
		}
		close(ch)
	}()

	return ch
}

func outputFormatNext(f *OutputFormat) *OutputFormat {
	return (*OutputFormat)(C.av_oformat_next((*C.struct_AVOutputFormat)(f)))
}

// OutputFormats returns a channel which can be used to iterate over the output formats.
//
// Usage:
//
//     for out := range avformat.OutputFormats() {
//         // ...
//     }
//
// C-Function: av_oformat_next
func OutputFormats() <-chan *OutputFormat {
	ch := make(chan *OutputFormat)

	var e *OutputFormat
	go func() {
		for {
			e = outputFormatNext(e)
			if e == nil {
				break
			}
			ch <- e
		}
		close(ch)
	}()

	return ch
}

// Version returns the LIBAVFORMAT_VERSION_INT constant.
//
// C-Function: avformat_version
func Version() uint {
	return uint(C.avformat_version())
}

// Configuration returns the libavformat build-time configuration.
//
// C-Function: avformat_configuration
func Configuration() string {
	return C.GoString(C.avformat_configuration())
}

// License returns the libavformat license.
//
// C-Function: avformat_license
func License() string {
	return C.GoString(C.avformat_license())
}

// RegisterAll initializes libavformat and register all the muxers, demuxers and protocols.
//
// C-Function: av_register_all
func RegisterAll() {
	C.av_register_all()
}

// NetworkInit does global initialization of network components.
//
// This is optional, but recommended, since it avoids the overhead of implicitly doing the setup for each session.
//
// Calling this function will become mandatory if using network protocols at some major version bump.
//
// C-Function: avformat_network_init
func NetworkInit() avutil.ReturnCode {
	return avutil.NewReturnCode(int(C.avformat_network_init()))
}

// NetworkDeinit undoes the initialization done by avformat_network_init.
//
// C-Function: avformat_network_deinit
func NetworkDeinit() int {
	return int(C.avformat_network_deinit())
}

// NewFormatContext allocates an FormatContext.
//
// C-Function: avformat_alloc_context
func NewFormatContext() *FormatContext {
	return (*FormatContext)(C.avformat_alloc_context())
}

// GetClass gets the Class for FormatContext.
//
// C-Function: avformat_get_class
func GetClass() *Class {
	return (*Class)(C.avformat_get_class())
}

// SideData returns side information from stream.
//
// C-Function: av_stream_get_side_data
func (s *Stream) SideData(t PacketSideDataType) []byte {
	var size int
	arr := (*[1 << 30]byte)(unsafe.Pointer(C.av_stream_get_side_data((*C.struct_AVStream)(s), (C.enum_AVPacketSideDataType)(t), (*C.int)(unsafe.Pointer(&size)))))
	return arr[:size:size]
}

// NewOutputFormatContext allocates a FormatContext for an output format.
//
// C-Function: avformat_alloc_output_context2
func NewOutputFormatContext(outFormat *OutputFormat, formatName, fileName string) (*FormatContext, avutil.ReturnCode) {
	var ctx *FormatContext
	code := avutil.NewReturnCode(int(C.avformat_alloc_output_context2((**C.struct_AVFormatContext)(unsafe.Pointer(&ctx)), (*C.struct_AVOutputFormat)(outFormat), C.CString(formatName), C.CString(fileName))))
	return ctx, code
}

// FindInputFormat finds InputFormat based on the short name of the input format.
//
// C-Function: av_find_input_format
func FindInputFormat(shortName string) *InputFormat {
	return (*InputFormat)(C.av_find_input_format(C.CString(shortName)))
}

// ProbeInputFormat guesses the file format.
//
// C-Function: av_probe_input_format
func ProbeInputFormat(pd *ProbeData, isOpened bool) *InputFormat {
	var iOpened int
	if isOpened {
		iOpened = 1
	} else {
		iOpened = 0
	}
	return (*InputFormat)(C.av_probe_input_format((*C.struct_AVProbeData)(pd), C.int(iOpened)))
}

// ProbeInputFormat2 guesses the file format.
//
// C-Function: av_probe_input_format2
func ProbeInputFormat2(pd *ProbeData, isOpened bool, scoreMax int) (*InputFormat, int) {
	var iOpened int
	if isOpened {
		iOpened = 1
	} else {
		iOpened = 0
	}

	inputFormat := (*InputFormat)(C.av_probe_input_format2((*C.struct_AVProbeData)(pd), C.int(iOpened), (*C.int)(unsafe.Pointer(&scoreMax))))
	return inputFormat, scoreMax
}

// ProbeInputFormat3 guesses the file format.
//
// C-Function: av_probe_input_format3
func ProbeInputFormat3(pd *ProbeData, isOpened bool) (*InputFormat, int) {
	var iOpened int
	if isOpened {
		iOpened = 1
	} else {
		iOpened = 0
	}
	var score int

	inputFormat := (*InputFormat)(C.av_probe_input_format3((*C.struct_AVProbeData)(pd), C.int(iOpened), (*C.int)(unsafe.Pointer(&score))))
	return inputFormat, score
}

// ProbeInputBuffer probes a bytestream to determine the input format.
//
// C-Function: av_probe_input_buffer2
func (c *IOContext) ProbeInputBuffer(url string, offset, maxProbeSize uint) (*InputFormat, avutil.ReturnCode) {
	var inputFormat *InputFormat
	var l int // TODO: this doesn't make any sense
	code := avutil.NewReturnCode(int(C.av_probe_input_buffer2((*C.struct_AVIOContext)(c), (**C.struct_AVInputFormat)(unsafe.Pointer(&inputFormat)), C.CString(url), unsafe.Pointer(&l), C.uint(offset), C.uint(maxProbeSize))))
	return inputFormat, code
}

// OpenInput opens an input stream and reads the header.
//
// C-Function: avformat_open_input
func OpenInput(url string, fmt *InputFormat, options **Dictionary) (*FormatContext, avutil.ReturnCode) {
	var ctx *FormatContext
	err := avutil.NewReturnCode(int(C.avformat_open_input((**C.struct_AVFormatContext)(unsafe.Pointer(&ctx)), C.CString(url), (*C.struct_AVInputFormat)(fmt), (**C.struct_AVDictionary)(unsafe.Pointer(options)))))
	return ctx, err
}

// GuessFormat Return the output format in the list of registered output formats which best matches the provided parameters, or return NULL if there is no match.
//
// C-Function: av_guess_format
func GuessFormat(sn, f, mt string) *OutputFormat {
	return (*OutputFormat)(C.av_guess_format(C.CString(sn), C.CString(f), C.CString(mt)))
}

// GuessCodec Guess the codec ID based upon muxer and filename.
//
// C-Function: av_guess_codec
func GuessCodec(fmt *OutputFormat, sn, f, mt string, t avutil.MediaType) CodecId {
	return (CodecId)(C.av_guess_codec((*C.struct_AVOutputFormat)(fmt), C.CString(sn), C.CString(f), C.CString(mt), (C.enum_AVMediaType)(t)))
}

// HexDump Send a nice hexadecimal dump of a buffer to the specified file stream.
//
// C-Function: av_hex_dump
func HexDump(f *File, b *uint8, s int) {
	C.av_hex_dump((*C.FILE)(f), (*C.uint8_t)(b), C.int(s))
}

// HexDumpLog Send a nice hexadecimal dump of a buffer to the log.
//
// C-Function: av_hex_dump_log
func HexDumpLog(a, l int, b *uint8, s int) {
	C.av_hex_dump_log(unsafe.Pointer(&a), C.int(l), (*C.uint8_t)(b), C.int(s))
}

// PktDump2 Send a nice dump of a packet to the specified file stream.
//
// C-Function: av_pkt_dump2
func PktDump2(f *File, pkt *Packet, dp int, st *Stream) {
	C.av_pkt_dump2((*C.FILE)(f), (*C.struct_AVPacket)(pkt), C.int(dp), (*C.struct_AVStream)(st))
}

// PktDumpLog2 Send a nice dump of a packet to the log.
//
// C-Function: av_pkt_dump_log2
func PktDumpLog2(a int, l int, pkt *Packet, dp int, st *Stream) {
	C.av_pkt_dump_log2(unsafe.Pointer(&a), C.int(l), (*C.struct_AVPacket)(pkt), C.int(dp), (*C.struct_AVStream)(st))
}

// CodecGetId enum CodecId av_codec_get_id (const struct CodecTag *const *tags, unsigned int tag)
//  Get the CodecId for the given codec tag tag.
//
// C-Function: av_codec_get_id
func CodecGetId(t **CodecTag, tag uint) CodecId {
	return (CodecId)(C.av_codec_get_id((**C.struct_AVCodecTag)(unsafe.Pointer(t)), C.uint(tag)))
}

// CodecGetTag Get the codec tag for the given codec id id.
//
// C-Function: av_codec_get_tag
func CodecGetTag(t **CodecTag, id CodecId) uint {
	return uint(C.av_codec_get_tag((**C.struct_AVCodecTag)(unsafe.Pointer(t)), (C.enum_AVCodecID)(id)))
}

// CodecGetTag2 Get the codec tag for the given codec id.
//
// C-Function: av_codec_get_tag2
func CodecGetTag2(t **CodecTag, id CodecId, tag *uint) int {
	return int(C.av_codec_get_tag2((**C.struct_AVCodecTag)(unsafe.Pointer(t)), (C.enum_AVCodecID)(id), (*C.uint)(unsafe.Pointer(tag))))
}

// IndexSearchTimestamp Get the index for a specific timestamp.
//
// C-Function: av_index_search_timestamp
func IndexSearchTimestamp(st *Stream, t int64, f int) int {
	return int(C.av_index_search_timestamp((*C.struct_AVStream)(st), C.int64_t(t), C.int(f)))
}

// AddIndexEntry Add an index entry into a sorted list.
//
// C-Function: av_add_index_entry
func AddIndexEntry(st *Stream, pos, t, int64, s, d, f int) int {
	return int(C.av_add_index_entry((*C.struct_AVStream)(st), C.int64_t(pos), C.int64_t(t), C.int(s), C.int(d), C.int(f)))
}

// URLSplit Split a URL string into components.
//
// C-Function: av_url_split
func URLSplit(p string, ps int, a string, as int, h string, hs int, pp *int, path string, psize int, url string) {
	C.av_url_split(C.CString(p), C.int(ps), C.CString(a), C.int(as), C.CString(h), C.int(hs), (*C.int)(unsafe.Pointer(pp)), C.CString(path), C.int(psize), C.CString(url))
}

// GetFrameFilename int av_get_frame_filename (char *buf, int buf_size, const char *path, int number)
// Return in 'buf' the path with 'd' replaced by a number.
//
// C-Function: av_get_frame_filename
func GetFrameFilename(b string, bs int, pa string, n int) int {
	return int(C.av_get_frame_filename(C.CString(b), C.int(bs), C.CString(pa), C.int(n)))
}

// FilenameNumberTest Check whether filename actually is a numbered sequence generator.
//
// C-Function: av_filename_number_test
func FilenameNumberTest(f string) int {
	return int(C.av_filename_number_test(C.CString(f)))
}

// SPDCreate Generate an SDP for an RTP session.
//
// C-Function: av_sdp_create
func SPDCreate(ac **FormatContext, nf int, b string, s int) int {
	return int(C.av_sdp_create((**C.struct_AVFormatContext)(unsafe.Pointer(ac)), C.int(nf), C.CString(b), C.int(s)))
}

// MatchExt int av_match_ext (const char *filename, const char *extensions)
// Return a positive value if the given filename has one of the given extensions, 0 otherwise.
//
// C-Function: av_match_ext
func MatchExt(f, e string) int {
	return int(C.av_match_ext(C.CString(f), C.CString(e)))
}

// QueryCodec Test if the given container can store a codec.
//
// C-Function: avformat_query_codec
func QueryCodec(o *OutputFormat, cd CodecId, sc int) int {
	return int(C.avformat_query_codec((*C.struct_AVOutputFormat)(o), (C.enum_AVCodecID)(cd), C.int(sc)))
}

// GetRiffVideoTags
//
//C-Function: avformat_get_riff_video_tags
func GetRiffVideoTags() *CodecTag {
	return (*CodecTag)(C.avformat_get_riff_video_tags())
}

// GetRiffAudioTags struct CodecTag * avformat_get_riff_audio_tags (void)
//
// C-Function: avformat_get_riff_audio_tags
func GetRiffAudioTags() *CodecTag {
	return (*CodecTag)(C.avformat_get_riff_audio_tags())
}

// GetMovVideoTags
//
//C-Function: avformat_get_mov_video_tags
func GetMovVideoTags() *CodecTag {
	return (*CodecTag)(C.avformat_get_mov_video_tags())
}

// GetMovAudioTags
//
//C-Function: avformat_get_mov_audio_tags
func GetMovAudioTags() *CodecTag {
	return (*CodecTag)(C.avformat_get_mov_audio_tags())
}
