// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)
// Corbatto (luca@corbatto.de)

package avcodec

//#cgo CFLAGS: -Wno-deprecated-declarations
//#cgo pkg-config: libavcodec
//#include <libavcodec/avcodec.h>
import "C"
import (
	"github.com/wix-playground/ffgopeg/avutil"
)

// ActiveThreadType returns which multithreading methods are used by the codec.
//
// C-Field: AVCodecContext::active_thread_type
func (c *CodecContext) ActiveThreadType() int {
	return int(c.active_thread_type)
}

// BQuantFactor returns the qscale factor between IP and B-frames.
//
// C-Field: AVCodecContext::b_quant_factor
func (c *CodecContext) BQuantFactor() float32 {
	return float32(c.b_quant_factor)
}

// SetBQuantFactor sets the qscale factor between IP and B-frames.
// If > 0 then the last P-frame quantizer will be used (q = lastp_q*factor+offset).
// If < 0 then normal ratecontrol will be done (q = -normal_q*factor+offset)
//
// C-Field: AVCodecContext::b_quant_factor
func (c *CodecContext) SetBQuantFactor(v float32) {
	c.b_quant_factor = C.float(v)
}

// BQuantOffset return the qscale offset between IP and B-frames.
//
// C-Field: AVCodecContext::b_quant_offset
func (c *CodecContext) BQuantOffset() float32 {
	return float32(c.b_quant_offset)
}

// SetBQuantOffset return the qscale offset between IP and B-frames.
//
// C-Field: AVCodecContext::b_quant_offset
func (c *CodecContext) SetBQuantOffset(v float32) {
	c.b_quant_offset = C.float(v)
}

// BidirRefine returns something undocumented...
//
// C-Field: AVCodecContext::bidir_refine
func (c *CodecContext) BidirRefine() int {
	return int(c.bidir_refine)
}

// SetBidirRefine sets something undocumented...
//
// C-Field: AVCodecContext::bidir_refine
func (c *CodecContext) SetBidirRefine(v int) {
	c.bidir_refine = C.int(v)
}

// BitRate returns the average bit rate.
//
// C-Field: AVCodecContext::bit_rate
func (c *CodecContext) BitRate() int {
	return int(c.bit_rate)
}

// SetBitRate sets the average bit rate.
//
// C-Field: AVCodecContext::bit_rate
func (c *CodecContext) SetBitRate(v int) {
	c.bit_rate = C.int64_t(v)
}

// BitRateTolerance returns the number of bits the bitstream is allowed to diverge from the reference.
//
// C-Field: AVCodecContext::bit_rate_tolerance
func (c *CodecContext) BitRateTolerance() int {
	return int(c.bit_rate_tolerance)
}

// SetBitRateTolerance sets the number of bits the bitstream is allowed to diverge from the reference.
// The reference can be CBR (for CBR pass 1) or VBR (for VBR pass 2)
//
// C-Field: AVCodecContext::bit_rate_tolerance
func (c *CodecContext) SetBitRateTolerance(v int) {
	c.bit_rate_tolerance = C.int(v)
}

// BitsPerCodedSample returns the bits per sample/pixel from the demuxer (needed for huffyuv).
//
// C-Field: AVCodecContext::bits_per_coded_sample
func (c *CodecContext) BitsPerCodedSample() int {
	return int(c.bits_per_coded_sample)
}

// SetBitsPerCodedSample set the bits per sample/pixel from the demuxer (needed for huffyuv).
//
// C-Field: AVCodecContext::bits_per_coded_sample
func (c *CodecContext) SetBitsPerCodedSample(v int) {
	c.bits_per_coded_sample = C.int(v)
}

// BitsPerRawSample returns the bits per sample/pixel of internal libavcodec pixel/sample format.
//
// C-Field: AVCodecContext::bits_per_raw_sample
func (c *CodecContext) BitsPerRawSample() int {
	return int(c.bits_per_raw_sample)
}

// SetBitsPerRawSample sets the bits per sample/pixel of internal libavcodec pixel/sample format.
//
// C-Field: AVCodecContext::bits_per_raw_sample
func (c *CodecContext) SetBitsPerRawSample(v int) {
	c.bits_per_raw_sample = C.int(v)
}

// BlockAlign returns the number of bytes per packet if constant and known or 0.
// Used by some WAV based audio codecs.
//
// C-Field: AVCodecContext::block_align
func (c *CodecContext) BlockAlign() int {
	return int(c.block_align)
}

// Channels returns the number of audio channels.
//
// C-Field: AVCodecContext::channels
func (c *CodecContext) Channels() int {
	return int(c.channels)
}

// SetChannels sets the number of audio channels.
//
// C-Field: AVCodecContext::channels
func (c *CodecContext) SetChannels(channels int) {
	c.channels = C.int(channels)
}

// ChannelLayout returns the audio channel layout.
//
// C-Field: AVCodecContext::channel_layout
func (c *CodecContext) ChannelLayout() uint64 {
	return uint64(c.channel_layout)
}

// SetChannelLayout sets the audio channel layout.
//
// C-Field: AVCodecContext::channel_layout
func (c *CodecContext) SetChannelLayout(layout uint64) {
	c.channel_layout = C.uint64_t(layout)
}

// CodedHeight returns the bitstream height.
//
// C-Field: AVCodecContext::coded_height
func (c *CodecContext) CodedHeight() int {
	return int(c.coded_height)
}

// SetCodedHeight sets the bitstream height.
//
// C-Field: AVCodecContext::coded_height
func (c *CodecContext) SetCodedHeight(v int) {
	c.coded_height = C.int(v)
}

// CodedWidth returns the bitstream width.
// It may be different from the raw width. When the decoded frame is cropped before being output or lowres is enabled.
//
// C-Field: AVCodecContext::coded_width
func (c *CodecContext) CodedWidth() int {
	return int(c.coded_width)
}

// SetCodedWidth sets the bitstream width.
// It may be different from the raw width. When the decoded frame is cropped before being output or lowres is enabled.
//
// C-Field: AVCodecContext::coded_width
func (c *CodecContext) SetCodedWidth(width int) {
	c.coded_width = C.int(width)
}

// Height returns the raw image height.
//
// C-Field: AVCodecContext::height
func (c *CodecContext) Height() int {
	return int(c.height)
}

// SetHeight sets the raw image height.
//
// C-Field: AVCodecContext::coded_height
func (c *CodecContext) SetHeight(height int) {
	c.height = C.int(height)
}

// Width returns the raw image width.
//
// C-Field: AVCodecContext::width
func (c *CodecContext) Width() int {
	return int(c.width)
}

// SetWidth sets the raw image width.
//
// C-Field: AVCodecContext::width
func (c *CodecContext) SetWidth(v int) {
	c.width = C.int(v)
}

// CompressionLevel returns the compression level.
//
// C-Field: AVCodecContext::compression_level
func (c *CodecContext) CompressionLevel() int {
	return int(c.compression_level)
}

// SetCompressionLevel sets the compression level.
//
// C-Field: AVCodecContext::compression_level
func (c *CodecContext) SetCompressionLevel(v int) {
	c.compression_level = C.int(v)
}

// Cutoff returns the audio cutoff bandwidth (0 means "automatic")
//
// C-Field: AVCodecContext::cutoff
func (c *CodecContext) Cutoff() int {
	return int(c.cutoff)
}

// SetCutoff sets the audio cutoff bandwidth (0 means "automatic")
//
// C-Field: AVCodecContext::cutoff
func (c *CodecContext) SetCutoff(v int) {
	c.cutoff = C.int(v)
}

// DarkMasking returns the darkness masking (0 means "disabled")
//
// C-Field: AVCodecContext::dark_masking
func (c *CodecContext) DarkMasking() float32 {
	return float32(c.dark_masking)
}

// SetDarkMasking sets the darkness masking (0 means "disabled")
//
// C-Field: AVCodecContext::dark_masking
func (c *CodecContext) SetDarkMasking(v float32) {
	c.dark_masking = C.float(v)
}

// DctAlgo returns the DCT algorithm, see FF_DCT_*.
//
// C-Field: AVCodecContext::dct_algo
func (c *CodecContext) DctAlgo() int {
	return int(c.dct_algo)
}

// SetDctAlgo sets the DCT algorithm, see FF_DCT_*.
//
// C-Field: AVCodecContext::dct_algo
func (c *CodecContext) SetDctAlgo(v int) {
	c.dct_algo = C.int(v)
}

// Debug returns... the documentation only says "debug"
//
// C-Field: AVCodecContext::debug
func (c *CodecContext) Debug() int {
	return int(c.debug)
}

// SetDebug sets... the documentation only says "debug"
//
// C-Field: AVCodecContext::debug
func (c *CodecContext) SetDebug(v int) {
	c.debug = C.int(v)
}

// DebugMv returns... the documentation only says "debug"
// Code outside libavcodec should access this field using AVOptions.
//
// C-Field: AVCodecContext::debug_mv
func (c *CodecContext) DebugMv() int {
	return int(c.debug_mv)
}

// SetDebugMv sets... the documentation only says "debug"
// Code outside libavcodec should access this field using AVOptions.
//
// C-Field: AVCodecContext::debug_mv
func (c *CodecContext) SetDebugMv(v int) {
	c.debug_mv = C.int(v)
}

// Delay returns the codec delay.
//
// Encoding: Number of frames delay there will be from the encoder input to the decoder output. (we assume the decoder matches the spec) Decoding: Number of frames delay in addition to what a standard decoder as specified in the spec would produce.
//
// Video: Number of frames the decoded output will be delayed relative to the encoded input.
//
// Audio: For encoding, this field is unused (see initial_padding).
//
// For decoding, this is the number of samples the decoder needs to output before the decoder's output is valid. When seeking, you should start decoding this many samples prior to your desired seek point.
//
// C-Field: AVCodecContext::delay
func (c *CodecContext) Delay() int {
	return int(c.delay)
}

// SetDelay sets the codec delay.
//
// Encoding: Number of frames delay there will be from the encoder input to the decoder output. (we assume the decoder matches the spec) Decoding: Number of frames delay in addition to what a standard decoder as specified in the spec would produce.
//
// Video: Number of frames the decoded output will be delayed relative to the encoded input.
//
// Audio: For encoding, this field is unused (see initial_padding).
//
// For decoding, this is the number of samples the decoder needs to output before the decoder's output is valid. When seeking, you should start decoding this many samples prior to your desired seek point.
//
// C-Field: AVCodecContext::delay
func (c *CodecContext) SetDelay(v int) {
	c.delay = C.int(v)
}

// DiaSize returns ME diamod size and shape.
//
// C-Field: AVCodecContext::dia_size
func (c *CodecContext) DiaSize() int {
	return int(c.dia_size)
}

// SetDiaSize sets ME diamod size and shape.
//
// C-Field: AVCodecContext::dia_size
func (c *CodecContext) SetDiaSize(v int) {
	c.dia_size = C.int(v)
}

// ErrRecognition returns the error recognition.
// It may misdetect some more or less valid parts as errors.
//
// C-Field: AVCodecContext::err_recognition
func (c *CodecContext) ErrRecognition() int {
	return int(c.err_recognition)
}

// SetErrRecognition sets the error recognition.
// It may misdetect some more or less valid parts as errors.
//
// C-Field: AVCodecContext::err_recognition
func (c *CodecContext) SetErrRecognition(v int) {
	c.err_recognition = C.int(v)
}

// ErrorConcealment returns error concealment flags.
//
// C-Field: AVCodecContext::error_concealment
func (c *CodecContext) ErrorConcealment() int {
	return int(c.error_concealment)
}

// SetErrorConcealment sets error concealment flags.
//
// C-Field: AVCodecContext::error_concealment
func (c *CodecContext) SetErrorConcealment(v int) {
	c.error_concealment = C.int(v)
}

// ExtradataSize returns the extradata size.
//
// C-Field: AVCodecContext::extradata_size
func (c *CodecContext) ExtradataSize() int {
	return int(c.extradata_size)
}

// Flags returns the flags AV_CODEC_FLAG_*.
//
// C-Field: AVCodecContext::flags
func (c *CodecContext) Flags() int {
	return int(c.flags)
}

// SetFlags set the flags AV_CODEC_FLAG_*.
//
// C-Field: AVCodecContext::flags
func (c *CodecContext) SetFlags() int {
	return int(c.flags)
}

// Flags2 returns the flags AV_CODEC_FLAG2_*.
//
// C-Field: AVCodecContext::flags2
func (c *CodecContext) Flags2() int {
	return int(c.flags2)
}

// SetFlags2 sets the flags AV_CODEC_FLAG2_*.
//
// C-Field: AVCodecContext::flags2
func (c *CodecContext) SetFlags2(v int) {
	c.flags2 = C.int(v)
}

// FrameNumber returns the Frame counter, set by libavcodec.
//
// Decoding: Total number of frames returned from the decoder so far.
// Encoding: Total number of frames passed to the encoder so far.
//
// C-Field: AVCodecContext::frame_number
func (c *CodecContext) FrameNumber() int {
	return int(c.frame_number)
}

// FrameSize returns the number of samples per channel in an audio frame.
//
// Encoding: Set by libavcodec in avcodec_open2(). Each submitted frame except the last must contain exactly frame_size samples per channel. May be 0 when the codec has AV_CODEC_CAP_VARIABLE_FRAME_SIZE set, then the frame size is not restricted.
// Decoding: May be set by some decoders to indicate constant frame size.
//
// C-Field: AVCodecContext::frame_size
func (c *CodecContext) FrameSize() int {
	return int(c.frame_size)
}

// GlobalQuality returns the global quality for codecs which cannot change it per frame.
// This should be proportional to MPEG-1/2/4 qscale.
func (c *CodecContext) GlobalQuality() int {
	return int(c.global_quality)
}

// SetGlobalQuality sets the global quality for codecs which cannot change it per frame.
// This should be proportional to MPEG-1/2/4 qscale.
func (c *CodecContext) SetGlobalQuality(v int) {
	c.global_quality = C.int(v)
}

// GopSize returns the number of pictures in a group of pictures, or 0 for intra_only.
//
// C-Field: AVCodecContext::gop_size
func (c *CodecContext) GopSize() int {
	return int(c.gop_size)
}

// SetGopSize sets the number of pictures in a group of pictures, or 0 for intra_only.
//
// C-Field: AVCodecContext::gop_size
func (c *CodecContext) SetGopSize(v int) {
	c.gop_size = C.int(v)
}

// HasBFrames returns the size of the reordering buffer in the decoder.
//
// For MPEG-2 it is 1 IPB or 0 low delay IP.
//
// C-Field: AVCodecContext::has_b_frames
func (c *CodecContext) HasBFrames() int {
	return int(c.has_b_frames)
}

// IQuantFactor returns the qscale factor between P- and I-frames.
// If > 0 then the last P-frame quantizer will be used (q = lastp_q * factor + offset).
// If < 0 then normal ratecontrol will be done (q = -normal_q * factor + offset).
//
// C-Field: AVCodecContext::i_quant_factor
func (c *CodecContext) IQuantFactor() float32 {
	return float32(c.i_quant_factor)
}

// SetIQuantFactor sets the qscale factor between P- and I-frames.
// If > 0 then the last P-frame quantizer will be used (q = lastp_q * factor + offset).
// If < 0 then normal ratecontrol will be done (q = -normal_q * factor + offset).
//
// C-Field: AVCodecContext::i_quant_factor
func (c *CodecContext) SetIQuantFactor(v float32) {
	c.i_quant_factor = C.float(v)
}

// IQuantOffset returns the qscale offset between P and I-frames.
//
// C-Field: AVCodecContext::i_quant_offset
func (c *CodecContext) IQuantOffset() float32 {
	return float32(c.i_quant_offset)
}

// SetIQuantOffset sets the qscale offset between P and I-frames.
//
// C-Field: AVCodecContext::i_quant_offset
func (c *CodecContext) SetIQuantOffset(v float32) {
	c.i_quant_offset = C.float(v)
}

// IdctAlgo returns the IDCT algorithm, see FF_IDCT_*.
//
// C-Field: AVCodecContext::idct_algo
func (c *CodecContext) IdctAlgo() int {
	return int(c.idct_algo)
}

// SetIdctAlgo sets the IDCT algorithm, see FF_IDCT_*.
//
// C-Field: AVCodecContext::idct_algo
func (c *CodecContext) SetIdctAlgo(v int) {
	c.idct_algo = C.int(v)
}

// IldctCmp returns the interlaced DCT comparison function.
//
// C-Field: AVCodecContext::ildct_cmp
func (c *CodecContext) IldctCmp() int {
	return int(c.ildct_cmp)
}

// SetIldctCmp sets the interlaced DCT comparison function.
//
// C-Field: AVCodecContext::ildct_cmp
func (c *CodecContext) SetIldctCmp(v int) {
	c.ildct_cmp = C.int(v)
}

// IntraDcPrecision returns the precision of the intra DC coefficient - 8.
//
// C-Field: AVCodecContext::intra_dc_precision
func (c *CodecContext) IntraDcPrecision() int {
	return int(c.intra_dc_precision)
}

// SetIntraDcPrecision sets the precision of the intra DC coefficient - 8.
//
// C-Field: AVCodecContext::intra_dc_precision
func (c *CodecContext) SetIntraDcPrecision(v int) {
	c.intra_dc_precision = C.int(v)
}

// KeyintMin returns the minimum GOP size.
//
// C-Field: AVCodecContext::keyint_min
func (c *CodecContext) KeyintMin() int {
	return int(c.keyint_min)
}

// SetKeyintMin sets the minimum GOP size.
//
// C-Field: AVCodecContext::keyint_min
func (c *CodecContext) SetKeyintMin(v int) {
	c.keyint_min = C.int(v)
}

// LastPredictorCount returns the amount of previous MV predictors (2a+1 x 2a+1)
//
// C-Field: AVCodecContext::last_predictor_count
func (c *CodecContext) LastPredictorCount() int {
	return int(c.last_predictor_count)
}

// SetLastPredictorCount sets the amount of previous MV predictors (2a+1 x 2a+1)
//
// C-Field: AVCodecContext::last_predictor_count
func (c *CodecContext) SetLastPredictorCount(v int) {
	c.last_predictor_count = C.int(v)
}

// Level returns the level.
//
// C-Field: AVCodecContext::level
func (c *CodecContext) Level() int {
	return int(c.level)
}

// SetLevel sets the level.
//
// C-Field: AVCodecContext::level
func (c *CodecContext) SetLevel(v int) {
	c.level = C.int(v)
}

// LogLevelOffset returns something undocumented...
//
// C-Field: AVCodecContext::log_level_offset
func (c *CodecContext) LogLevelOffset() int {
	return int(c.log_level_offset)
}

// SetLogLevelOffset sets something undocumented...
//
// C-Field: AVCodecContext::log_level_offset
func (c *CodecContext) SetLogLevelOffset(v int) {
	c.log_level_offset = C.int(v)
}

// LumiMasking retunrs the level of luminance masking.
// 0 means disabled.
//
// C-Field: AVCodecContext::lumi_masking
func (c *CodecContext) LumiMasking() float32 {
	return float32(c.lumi_masking)
}

// SetLumiMasking retunrs the level of luminance masking.
// 0 means disabled.
//
// C-Field: AVCodecContext::lumi_masking
func (c *CodecContext) SetLumiMasking(v float32) {
	c.lumi_masking = C.float(v)
}

// MaxBFrames returns the maximum nomber of B-frames between non-B-frames.
//
// Note: The output will be delayed by max_b_frames+1 relative to the input.
//
// C-Field: AVCodecContext::max_b_frames
func (c *CodecContext) MaxBFrames() int {
	return int(c.max_b_frames)
}

// SetMaxBFrames sets the maximum nomber of B-frames between non-B-frames.
//
// Note: The output will be delayed by max_b_frames+1 relative to the input.
//
// C-Field: AVCodecContext::max_b_frames
func (c *CodecContext) SetMaxBFrames(v int) {
	c.max_b_frames = C.int(v)
}

// MaxQDiff returns the maximum quantizer difference between frames.
//
// C-Field: AVCodecContext::max_qdiff
func (c *CodecContext) MaxQDiff() int {
	return int(c.max_qdiff)
}

// SetMaxQDiff sets the maximum quantizer difference between frames.
//
// C-Field: AVCodecContext::max_qdiff
func (c *CodecContext) SetMaxQDiff(v int) {
	c.max_qdiff = C.int(v)
}

// MbCmp returns the macroblock comparison function.
// (not supported yet)
//
// C-Field: AVCodecContext::mb_cmp
func (c *CodecContext) MbCmp() int {
	return int(c.mb_cmp)
}

// SetMbCmp sets the macroblock comparison function.
// (not supported yet)
//
// C-Field: AVCodecContext::mb_cmp
func (c *CodecContext) SetMbCmp(v int) {
	c.mb_cmp = C.int(v)
}

// MbDecision returns the macroblock decision mode.
//
// C-Field: AVCodecContext::mb_decision
func (c *CodecContext) MbDecision() int {
	return int(c.mb_decision)
}

// SetMbDecision sets the macroblock decision mode.
//
// C-Field: AVCodecContext::mb_decision
func (c *CodecContext) SetMbDecision(v int) {
	c.mb_decision = C.int(v)
}

// MbLMax returns the maximum MB Lagrange multiplier.
//
// C-Field: AVCodecContext::mb_lmax
func (c *CodecContext) MbLMax() int {
	return int(c.mb_lmax)
}

// SetMbLMax returns the maximum MB Lagrange multiplier.
//
// C-Field: AVCodecContext::mb_lmax
func (c *CodecContext) SetMbLMax(v int) {
	c.mb_lmax = C.int(v)
}

// MbLMin returns the minimum MB Lagrange multiplier.
//
// C-Field: AVCodecContext::mb_lmin
func (c *CodecContext) MbLMin() int {
	return int(c.mb_lmin)
}

// SetMbLMin returns the minimum MB Lagrange multiplier.
//
// C-Field: AVCodecContext::mb_lmin
func (c *CodecContext) SetMbLMin(v int) {
	c.mb_lmin = C.int(v)
}

// MeCmp returns the motion estimation comparison function.
//
// C-Field: AVCodecContext::me_cmp
func (c *CodecContext) MeCmp() int {
	return int(c.me_cmp)
}

// SetMeCmp sets the motion estimation comparison function.
//
// C-Field: AVCodecContext::me_cmp
func (c *CodecContext) SetMeCmp(v int) {
	c.me_cmp = C.int(v)
}

// MePreCmp returns the motion estimation prepass comparison function.
//
// C-Field: AVCodecContext::me_pre_cmp
func (c *CodecContext) MePreCmp() int {
	return int(c.me_pre_cmp)
}

// SetMePreCmp sets the motion estimation prepass comparison function.
//
// C-Field: AVCodecContext::me_pre_cmp
func (c *CodecContext) SetMePreCmp(v int) {
	c.me_pre_cmp = C.int(v)
}

// MeRange returns the maximum motion estimation search range in subpel units.
// If 0 then no limit.
//
// C-Field: AVCodecContext::me_range
func (c *CodecContext) MeRange() int {
	return int(c.me_range)
}

// SetMeRange sets the maximum motion estimation search range in subpel units.
// If 0 then no limit.
//
// C-Field: AVCodecContext::me_range
func (c *CodecContext) SetMeRange(v int) {
	c.me_range = C.int(v)
}

// MeSubCmp returns the subpixel motion comparison function.
//
// C-Field: AVCodecContext::me_sub_cmp
func (c *CodecContext) MeSubCmp() int {
	return int(c.me_sub_cmp)
}

// SetMeSubCmp sets the subpixel motion comparison function.
//
// C-Field: AVCodecContext::me_sub_cmp
func (c *CodecContext) SetMeSubCmp(v int) {
	c.me_sub_cmp = C.int(v)
}

// MeSubpelQuality returns the subpel ME quality.
//
// C-Field: AVCodecContext::me_subpel_quality
func (c *CodecContext) MeSubpelQuality() int {
	return int(c.me_subpel_quality)
}

// SetMeSubpelQuality sets the subpel ME quality.
//
// C-Field: AVCodecContext::me_subpel_quality
func (c *CodecContext) SetMeSubpelQuality(v int) {
	c.me_subpel_quality = C.int(v)
}

// Mv0Threshold returns the mv0 threshold.
// Value depends on the compare function used for follpel ME.
//
// C-Field: AVCodecContext::mv0_threshold
func (c *CodecContext) Mv0Threshold() int {
	return int(c.mv0_threshold)
}

// SetMv0Threshold sets the mv0 threshold.
// Value depends on the compare function used for follpel ME.
//
// C-Field: AVCodecContext::mv0_threshold
func (c *CodecContext) SetMv0Threshold(v int) {
	c.mv0_threshold = C.int(v)
}

// NsseWeight returns the noise vs sse weight for the nsse comparison function.
//
// C-Field: AVCodecContext::nsse_weight
func (c *CodecContext) NsseWeight() int {
	return int(c.nsse_weight)
}

// PMasking returns the spatial complexity masking. 0 equals disabled.
//
// C-Field: AVCodecContext::p_masking
func (c *CodecContext) PMasking() float32 {
	return float32(c.p_masking)
}

// SetPMasking sets the spatial complexity masking. 0 equals disabled.
//
// C-Field: AVCodecContext::p_masking
func (c *CodecContext) SetPMasking(v float32) {
	c.p_masking = C.float(v)
}

// PreDiaSize returns the ME prepass diamond size & shape.
//
// C-Field: AVCodecContext::pre_dia_size
func (c *CodecContext) PreDiaSize() int {
	return int(c.pre_dia_size)
}

// SetPreDiaSize sets the ME prepass diamond size & shape.
//
// C-Field: AVCodecContext::pre_dia_size
func (c *CodecContext) SetPreDiaSize(v int) {
	c.pre_dia_size = C.int(v)
}

// Profile returns somthing undocumented...
//
// C-Field: AVCodecContext::profile
func (c *CodecContext) Profile() int {
	return int(c.profile)
}

// SetProfile sets somthing undocumented...
//
// C-Field: AVCodecContext::profile
func (c *CodecContext) SetProfile(v int) {
	c.profile = C.int(v)
}

// QBlur returns the amount of qscale smoothing over time.
// (0.0 - 1.0)
//
// C-Field: AVCodecContext::qblur
func (c *CodecContext) QBlur() float32 {
	return float32(c.qblur)
}

// SetQBlur sets the amount of qscale smoothing over time.
// (0.0 - 1.0)
//
// C-Field: AVCodecContext::qblur
func (c *CodecContext) SetQBlur(v float32) {
	c.qblur = C.float(v)
}

// QCompress returns the amount of qscale change between easy & hard scenes. (0.0 - 1.0)
//
// C-Field: AVCodecContext::qcompress
func (c *CodecContext) QCompress() float32 {
	return float32(c.qcompress)
}

// SetQCompress sets the amount of qscale change between easy & hard scenes. (0.0 - 1.0)
//
// C-Field: AVCodecContext::qcompress
func (c *CodecContext) SetQCompress(v float32) {
	c.qcompress = C.float(v)
}

// QMax returns the maximum quantizer.
//
// C-Field: AVCodecContext::qmax
func (c *CodecContext) QMax() int {
	return int(c.qmax)
}

// SetQMax sets the maximum quantizer.
//
// C-Field: AVCodecContext::qmax
func (c *CodecContext) SetQMax(v int) {
	c.qmax = C.int(v)
}

// QMin returns the minimum quantizer.
//
// C-Field: AVCodecContext::qmin
func (c *CodecContext) QMin() int {
	return int(c.qmin)
}

// SetQMin sets the minimum quantizer.
//
// C-Field: AVCodecContext::qmin
func (c *CodecContext) SetQMin(v int) {
	c.qmin = C.int(v)
}

// RcBufferSize returns the decoder bitstream buffer size.
//
// C-Field: AVCodecContext::rc_buffer_size
func (c *CodecContext) RcBufferSize() int {
	return int(c.rc_buffer_size)
}

// SetRcBufferSize sets the decoder bitstream buffer size.
//
// C-Field: AVCodecContext::rc_buffer_size
func (c *CodecContext) SetRcBufferSize(v int) {
	c.rc_buffer_size = C.int(v)
}

// RcInitialBufferOccupancy returns the number of bits which should be loaded into the rx buffer before decoding starts.
//
// C-Field: AVCodecContext::rc_initial_buffer_occupancy
func (c *CodecContext) RcInitialBufferOccupancy() int {
	return int(c.rc_initial_buffer_occupancy)
}

// SetRcInitialBufferOccupancy sets the number of bits which should be loaded into the rx buffer before decoding starts.
//
// C-Field: AVCodecContext::rc_initial_buffer_occupancy
func (c *CodecContext) SetRcInitialBufferOccupancy(v int) {
	c.rc_initial_buffer_occupancy = C.int(v)
}

// RcMaxAvailableVbvUse returns the ratecontrol attempt to use, at maximum, of what can be used without an underflow.
//
// C-Field: AVCodecContext::rc_max_available_vbv_use
func (c *CodecContext) RcMaxAvailableVbvUse() float32 {
	return float32(c.rc_max_available_vbv_use)
}

// SetRcMaxAvailableVbvUse sets the ratecontrol attempt to use, at maximum, of what can be used without an underflow.
//
// C-Field: AVCodecContext::rc_max_available_vbv_use
func (c *CodecContext) SetRcMaxAvailableVbvUse(v float32) {
	c.rc_max_available_vbv_use = C.float(v)
}

// RcMaxRate returns the maximum bitrate.
//
// C-Field: AVCodecContext::rc_max_rate
func (c *CodecContext) RcMaxRate() int64 {
	return int64(c.rc_max_rate)
}

// SetRcMaxRate sets the maximum bitrate.
//
// C-Field: AVCodecContext::rc_max_rate
func (c *CodecContext) SetRcMaxRate(v int64) {
	c.rc_max_rate = C.int64_t(v)
}

// RcMinRate returns the minimum bitrate.
//
// C-Field: AVCodecContext::rc_min_rate
func (c *CodecContext) RcMinRate() int64 {
	return int64(c.rc_min_rate)
}

// SetRcMinRate sets the minimum bitrate.
//
// C-Field: AVCodecContext::rc_min_rate
func (c *CodecContext) SetRcMinRate(v int64) {
	c.rc_min_rate = C.int64_t(v)
}

// RcMinVbvOverflowUse returns the ratecontrol attempt to use, at least, times the amount needed to prevent a vbv overflow.
//
// C-Field: AVCodecContext::rc_min_vbv_overflow_use
func (c *CodecContext) RcMinVbvOverflowUse() float32 {
	return float32(c.rc_min_vbv_overflow_use)
}

// SetRcMinVbvOverflowUse sets the ratecontrol attempt to use, at least, times the amount needed to prevent a vbv overflow.
//
// C-Field: AVCodecContext::rc_min_vbv_overflow_use
func (c *CodecContext) SetRcMinVbvOverflowUse(v float32) {
	c.rc_min_vbv_overflow_use = C.float(v)
}

// RcOverrideCount returns ratecontrol override, see RcOverride.
//
// C-Field: AVCodecContext::rc_override_count
func (c *CodecContext) RcOverrideCount() int {
	return int(c.rc_override_count)
}

// SetRcOverrideCount sets ratecontrol override, see RcOverride.
//
// C-Field: AVCodecContext::rc_override_count
func (c *CodecContext) SetRcOverrideCount(v int) {
	c.rc_override_count = C.int(v)
}

// RefcountedFrames returns the number of references of audio and video frames.
//
// If non-zero, the decoded audio and video frames returned from avcodec_decode_video2() and avcodec_decode_audio4() are reference-counted and are valid indefinitely.
//
// The caller must free them with av_frame_unref() when they are not needed anymore. Otherwise, the decoded frames must not be freed by the caller and are only valid until the next decode call.
//
// This is always automatically enabled if avcodec_receive_frame() is used.
//
// C-Field: AVCodecContext::refcounted_frames
func (c *CodecContext) RefcountedFrames() int {
	return int(c.refcounted_frames)
}

// Refs returns the number reference frames.
//
// C-Field: AVCodecContext::refs
func (c *CodecContext) Refs() int {
	return int(c.refs)
}

// SetRefs sets the number reference frames.
//
// C-Field: AVCodecContext::refs
func (c *CodecContext) SetRefs(v int) {
	c.refs = C.int(v)
}

// SampleRate returns the number of samples per second.
//
// C-Field: AVCodecContext::sample_rate
func (c *CodecContext) SampleRate() int {
	return int(c.sample_rate)
}

// SetSampleRate sets the number of samples per second.
//
// C-Field: AVCodecContext::sample_rate
func (c *CodecContext) SetSampleRate(sr int) {
	c.sample_rate = C.int(sr)
}

// SkipBottom returns the number of macroblocks rows at the bottom which are skipped.
//
// C-Field: AVCodecContext::skip_bottom
func (c *CodecContext) SkipBottom() int {
	return int(c.skip_bottom)
}

// SetSkipBottom sets the number of macroblocks rows at the bottom which are skipped.
//
// C-Field: AVCodecContext::skip_bottom
func (c *CodecContext) SetSkipBottom(v int) {
	c.skip_bottom = C.int(v)
}

// SkipTop returns the number of macroblock rows at the top which are skipped.
//
// C-Field: AVCodecContext::skip_top
func (c *CodecContext) SkipTop() int {
	return int(c.skip_top)
}

// SetSkipTop sets the number of macroblock rows at the top which are skipped.
//
// C-Field: AVCodecContext::skip_top
func (c *CodecContext) SetSkipTop(v int) {
	c.skip_top = C.int(v)
}

// SliceCount returns the slice count.
//
// C-Field: AVCodecContext::slice_count
func (c *CodecContext) SliceCount() int {
	return int(c.slice_count)
}

// SetSliceCount sets the slice count.
//
// C-Field: AVCodecContext::slice_count
func (c *CodecContext) SetSliceCount(v int) {
	c.slice_count = C.int(v)
}

// SliceFlags returns the slice flags.
//
// C-Field: AVCodecContext::slice_flags
func (c *CodecContext) SliceFlags() int {
	return int(c.slice_flags)
}

// SetSliceFlags sets the slice flags.
//
// C-Field: AVCodecContext::slice_flags
func (c *CodecContext) SetSliceFlags(v int) {
	c.slice_flags = C.int(v)
}

// Slices returns the number of slices.
//
// Indicates number of picture subdicisions. Used for parallelized decoding.
//
// C-Field: AVCodecContext::slices
func (c *CodecContext) Slices() int {
	return int(c.slices)
}

// SetSlices sets the number of slices.
//
// Indicates number of picture subdicisions. Used for parallelized decoding.
//
// C-Field: AVCodecContext::slices
func (c *CodecContext) SetSlices(v int) {
	c.slices = C.int(v)
}

// SpatialCplxMasking returns the spatial complexity masking. 0 means disabled.
//
// C-Field: AVCodecContext::spatial_cplx_masking
func (c *CodecContext) SpatialCplxMasking() float32 {
	return float32(c.spatial_cplx_masking)
}

// SetSpatialCplxMasking sets the spatial complexity masking. 0 means disabled.
//
// C-Field: AVCodecContext::spatial_cplx_masking
func (c *CodecContext) SetSpatialCplxMasking(v float32) {
	c.spatial_cplx_masking = C.float(v)
}

// StrictStdCompliance returns the strictness to follow the standard (MPEG-4, ...).
//
// C-Field: AVCodecContext::strict_std_compliance
func (c *CodecContext) StrictStdCompliance() int {
	return int(c.strict_std_compliance)
}

// SetStrictStdCompliance sets the strictness to follow the standard (MPEG-4, ...).
//
// C-Field: AVCodecContext::strict_std_compliance
func (c *CodecContext) SetStrictStdCompliance(v int) {
	c.strict_std_compliance = C.int(v)
}

// SubCharencMode returns the subtitles character encoding mode.
//
// Formats or codecs might be adjusting this settings (if they are doing the conversion themselves for instance).
//
// C-Field: AVCodecContext::sub_charenc_mode
func (c *CodecContext) SubCharencMode() int {
	return int(c.sub_charenc_mode)
}

// SubtitleHeaderSize returns something undocumented...
//
// C-Field: AVCodecContext::subtitle_header_size
func (c *CodecContext) SubtitleHeaderSize() int {
	return int(c.subtitle_header_size)
}

// TemporalCplxMasking returns the temporal complexity masking. 0 means disabled.
//
// C-Field: AVCodecContext::temporal_cplx_masking
func (c *CodecContext) TemporalCplxMasking() float32 {
	return float32(c.temporal_cplx_masking)
}

// SetTemporalCplxMasking returns the temporal complexity masking. 0 means disabled.
//
// C-Field: AVCodecContext::temporal_cplx_masking
func (c *CodecContext) SetTemporalCplxMasking(v float32) {
	c.temporal_cplx_masking = C.float(v)
}

// ThreadCount returns the thread count, which is used to decide how many independent tasks should be passed to execute().
//
// C-Field: AVCodecContext::thread_count
func (c *CodecContext) ThreadCount() int {
	return int(c.thread_count)
}

// SetThreadCount sets the thread count, which is used to decide how many independent tasks should be passed to execute().
//
// C-Field: AVCodecContext::thread_count
func (c *CodecContext) SetThreadCount(v int) {
	c.thread_count = C.int(v)
}

// ThreadSafeCallbacks returns whether or not the custom get_buffer() callback can be called synchronously from another threa, which allows faster multithreaded decoding.
//
// C-Field: AVCodecContext::thread_safe_callbacks
func (c *CodecContext) ThreadSafeCallbacks() bool {
	return int(c.thread_safe_callbacks) != 0
}

// SetThreadSafeCallbacks sets whether or not the custom get_buffer() callback can be called synchronously from another threa, which allows faster multithreaded decoding.
//
// C-Field: AVCodecContext::thread_safe_callbacks
func (c *CodecContext) SetThreadSafeCallbacks(v bool) {
	if v {
		c.thread_safe_callbacks = C.int(1)
	} else {
		c.thread_safe_callbacks = C.int(0)
	}
}

// ThreadType returns which multithreading methods to use.
//
// C-Field: AVCodecContext::thread_type
func (c *CodecContext) ThreadType() int {
	return int(c.thread_type)
}

// SetThreadType sets which multithreading methods to use.
//
// C-Field: AVCodecContext::thread_type
func (c *CodecContext) SetThreadType(v int) {
	c.thread_type = C.int(v)
}

// TimeBase returns the unit of time (in seconds).
//
// This is the fundamental unit of time (in seconds) in terms
// of which frame timestamps are represented. For fixed-fps content,
// timebase should be 1/framerate and timestamp increments should be
// identically 1.
// This often, but not always is the inverse of the frame rate or field rate
// for video. 1/time_base is not the average frame rate if the frame rate is not
// constant.
//
// Like containers, elementary streams also can store timestamps, 1/time_base
// is the unit in which these timestamps are specified.
// As example of such codec time base see ISO/IEC 14496-2:2001(E)
// vop_time_increment_resolution and fixed_vop_rate
// (fixed_vop_rate == 0 implies that it is different from the framerate)
//
//  - encoding: MUST be set by user.
// - decoding: the use of this field for decoding is deprecated.
//            Use framerate instead.
//
// C-Field: AVCodecContext::time_base
func (c *CodecContext) TimeBase() Rational {
	return (Rational)(c.time_base)
}

// SetTimeBase sets the unit of time (in seconds).
//
// This is the fundamental unit of time (in seconds) in terms
// of which frame timestamps are represented. For fixed-fps content,
// timebase should be 1/framerate and timestamp increments should be
// identically 1.
// This often, but not always is the inverse of the frame rate or field rate
// for video. 1/time_base is not the average frame rate if the frame rate is not
// constant.
//
// Like containers, elementary streams also can store timestamps, 1/time_base
// is the unit in which these timestamps are specified.
// As example of such codec time base see ISO/IEC 14496-2:2001(E)
// vop_time_increment_resolution and fixed_vop_rate
// (fixed_vop_rate == 0 implies that it is different from the framerate)
//
//  - encoding: MUST be set by user.
// - decoding: the use of this field for decoding is deprecated.
//            Use framerate instead.
//
// C-Field: AVCodecContext::time_base
func (c *CodecContext) SetTimeBase(v Rational) {
	c.time_base = (C.struct_AVRational)(v)
}

// TicksPerFrame returns the ticks per frame.
//
// For some codecs, the time base is closer to the field rate than the frame rate.
//
// Most notably, H.264 and MPEG-2 specify time_base as half of frame duration if no telecine is used ...
//
// Set to time_base ticks per frame. Default 1, e.g., H.264/MPEG-2 set it to 2.
//
// C-Field: AVCodecContext::ticks_per_frame
func (c *CodecContext) TicksPerFrame() int {
	return int(c.ticks_per_frame)
}

// SetTicksPerFrame sets the ticks per frame.
//
// For some codecs, the time base is closer to the field rate than the frame rate.
//
// Most notably, H.264 and MPEG-2 specify time_base as half of frame duration if no telecine is used ...
//
// Set to time_base ticks per frame. Default 1, e.g., H.264/MPEG-2 set it to 2.
//
// C-Field: AVCodecContext::ticks_per_frame
func (c *CodecContext) SetTicksPerFrame(v int) {
	c.ticks_per_frame = C.int(v)
}

// Trellis returns the trellis RD quantization.
//
// C-Field: AVCodecContext::trellis
func (c *CodecContext) Trellis() int {
	return int(c.trellis)
}

// SetTrellis sets the trellis RD quantization.
//
// C-Field: AVCodecContext::trellis
func (c *CodecContext) SetTrellis(v int) {
	c.trellis = C.int(v)
}

// WorkaroundBugs returns which bufs to work around.
//
// C-Field: AVCodecContext::workaround_bugs
func (c *CodecContext) WorkaroundBugs() int {
	return int(c.workaround_bugs)
}

// SetWorkaroundBugs sets which bufs to work around.
//
// C-Field: AVCodecContext::workaround_bugs
func (c *CodecContext) SetWorkaroundBugs(v int) {
	c.workaround_bugs = C.int(v)
}

// AudioServiceType returns the type of service that the audio stream conveys.
//
// C-Field: AVCodecContext::audio_service_type
func (c *CodecContext) AudioServiceType() AudioServiceType {
	return (AudioServiceType)(c.audio_service_type)
}

// SetAudioServiceType sets the type of service that the audio stream conveys.
//
// C-Field: AVCodecContext::audio_service_type
func (c *CodecContext) SetAudioServiceType(v AudioServiceType) {
	c.audio_service_type = C.enum_AVAudioServiceType(v)
}

// ChromaSampleLocation returns the location of chroma samples.
//
// C-Field: AVCodecContext::chroma_sample_location
func (c *CodecContext) ChromaSampleLocation() ChromaLocation {
	return (ChromaLocation)(c.chroma_sample_location)
}

// SetChromaSampleLocation sets the location of chroma samples.
//
// C-Field: AVCodecContext::chroma_sample_location
func (c *CodecContext) SetChromaSampleLocation(v ChromaLocation) {
	c.chroma_sample_location = C.enum_AVChromaLocation(v)
}

// CodecID returns the codec id.
//
// C-Field: AVCodecContext::codec_id
func (c *CodecContext) CodecID() CodecId {
	return (CodecId)(c.codec_id)
}

// CodecType return the codec type.
//
// C-Field: AVCodecContext::codec_type
func (c *CodecContext) CodecType() avutil.MediaType {
	return (avutil.MediaType)(c.codec_type)
}

// ColorPrimaries returns the chromaticity coordinates of the source primaries.
//
// C-Field: AVCodecContext::color_primaries
func (c *CodecContext) ColorPrimaries() ColorPrimaries {
	return (ColorPrimaries)(c.color_primaries)
}

// SetColorPrimaries sets the chromaticity coordinates of the source primaries.
//
// C-Field: AVCodecContext::color_primaries
func (c *CodecContext) SetColorPrimaries(v ColorPrimaries) {
	c.color_primaries = C.enum_AVColorPrimaries(v)
}

// ColorRange returns the MPEG vs JPEG YUV range.
//
// C-Field: AVCodecContext::color_range
func (c *CodecContext) ColorRange() ColorRange {
	return (ColorRange)(c.color_range)
}

// SetColorRange sets the MPEG vs JPEG YUV range.
//
// C-Field: AVCodecContext::color_range
func (c *CodecContext) SetColorRange(v ColorRange) {
	c.color_range = C.enum_AVColorRange(v)
}

// ColorTrc returns the color transfer characteristic.
//
// C-Field: AVCodecContext::color_trc
func (c *CodecContext) ColorTrc() ColorTransferCharacteristic {
	return (ColorTransferCharacteristic)(c.color_trc)
}

// SetColorTrc sets the color transfer characteristic.
//
// C-Field: AVCodecContext::color_trc
func (c *CodecContext) SetColorTrc(v ColorTransferCharacteristic) {
	c.color_trc = C.enum_AVColorTransferCharacteristic(v)
}

// Colorspace returns the YUV colorspace type.
//
// C-Field: AVCodecContext::colorspace
func (c *CodecContext) Colorspace() ColorSpace {
	return (ColorSpace)(c.colorspace)
}

// SetColorspace sets the YUV colorspace type.
//
// C-Field: AVCodecContext::colorspace
func (c *CodecContext) SetColorspace(v ColorSpace) {
	c.colorspace = C.enum_AVColorSpace(v)
}

// FieldOrder returns the field order.
//
// C-Field: AVCodecContext::field_order
func (c *CodecContext) FieldOrder() FieldOrder {
	return (FieldOrder)(c.field_order)
}

// SetFieldOrder sets the field order.
//
// C-Field: AVCodecContext::field_order
func (c *CodecContext) SetFieldOrder(v FieldOrder) {
	c.field_order = C.enum_AVFieldOrder(v)
}

// PixFmt returns the pixel format, see AV_PIX_FMT_xxx.
//
// C-Field: AVCodecContext::pix_fmt
func (c *CodecContext) PixFmt() avutil.PixelFormat {
	return (avutil.PixelFormat)(c.pix_fmt)
}

// SetPixFmt sets the pixel format, see AV_PIX_FMT_xxx.
//
// C-Field: AVCodecContext::pix_fmt
func (c *CodecContext) SetPixFmt(v avutil.PixelFormat) {
	c.pix_fmt = C.enum_AVPixelFormat(v)
}

// RequestSampleFmt returns the desired sample format.
//
// C-Field: AVCodecContext::request_sample_fmt
func (c *CodecContext) RequestSampleFmt() avutil.SampleFormat {
	return (avutil.SampleFormat)(c.request_sample_fmt)
}

// SetRequestSampleFmt sets the desired sample format.
//
// C-Field: AVCodecContext::request_sample_fmt
func (c *CodecContext) SetRequestSampleFmt(fmt avutil.SampleFormat) {
	c.request_sample_fmt = C.enum_AVSampleFormat(fmt)
}

// SampleAspectRatio returns the sample aspect ratio.
//
// sample aspect ratio (0 if unknown) That is the width of a pixel divided by the height of the pixel.
//
// Numerator and denominator must be relatively prime and smaller than 256 for some video standards.
//
// C-Field: AVCodecContext::sample_aspect_ratio
func (c *CodecContext) SampleAspectRatio() Rational {
	return (Rational)(c.sample_aspect_ratio)
}

// SetSampleAspectRatio sets the sample aspect ratio.
//
// sample aspect ratio (0 if unknown) That is the width of a pixel divided by the height of the pixel.
//
// Numerator and denominator must be relatively prime and smaller than 256 for some video standards.
//
// C-Field: AVCodecContext::sample_aspect_ratio
func (c *CodecContext) SetSampleAspectRatio(v Rational) {
	c.sample_aspect_ratio = (C.struct_AVRational)(v)
}

// SampleFmt returns the audio sample format.
//
// C-Field: AVCodecContext::sample_fmt
func (c *CodecContext) SampleFmt() avutil.SampleFormat {
	return (avutil.SampleFormat)(c.sample_fmt)
}

// SetSampleFmt sets the audio sample format.
//
// C-Field: AVCodecContext::sample_fmt
func (c *CodecContext) SetSampleFmt(v avutil.SampleFormat) {
	c.sample_fmt = C.enum_AVSampleFormat(v)
}

// SkipFrame returns the skip decoding for selected frames.
//
// C-Field: AVCodecContext::skip_frame
func (c *CodecContext) SkipFrame() Discard {
	return (Discard)(c.skip_frame)
}

// SetSkipFrame sets the skip decoding for selected frames.
//
// C-Field: AVCodecContext::skip_frame
func (c *CodecContext) SetSkipFrame(v Discard) {
	c.skip_frame = C.enum_AVDiscard(v)
}

// SkipIdct returns the skip IDCT/dequantization for selected frames.
//
// C-Field: AVCodecContext::skip_idct
func (c *CodecContext) SkipIdct() Discard {
	return (Discard)(c.skip_idct)
}

// SetSkipIdct sets the skip IDCT/dequantization for selected frames.
//
// C-Field: AVCodecContext::skip_idct
func (c *CodecContext) SetSkipIdct(v Discard) {
	c.skip_idct = C.enum_AVDiscard(v)
}

// SkipLoopFilter returns the skip loop filtering for selected frames.
//
// C-Field: AVCodecContext::skip_loop_filter
func (c *CodecContext) SkipLoopFilter() Discard {
	return (Discard)(c.skip_loop_filter)
}

// SetSkipLoopFilter sets the skip loop filtering for selected frames.
//
// C-Field: AVCodecContext::skip_loop_filter
func (c *CodecContext) SetSkipLoopFilter(v Discard) {
	c.skip_loop_filter = C.enum_AVDiscard(v)
}
