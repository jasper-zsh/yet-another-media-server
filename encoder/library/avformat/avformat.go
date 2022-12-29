package avformat

//#cgo pkg-config: libavutil libavformat libavcodec
//
//#include <libavformat/avformat.h>
//#include <libavcodec/avcodec.h>
//#include <stdlib.h>
import "C"
import (
	"unsafe"
	"yet-another-media-server/encoder/library/avcodec"
	"yet-another-media-server/encoder/library/avutil"
)

type (
	AVFormatContext C.struct_AVFormatContext
	AVInputFormat   C.struct_AVInputFormat
	AVStream        C.struct_AVStream
)

const (
	AVSEEK_FLAG_BACKWARD = 1 ///< seek backward
	AVSEEK_FLAG_BYTE     = 2 ///< seeking based on position in bytes
	AVSEEK_FLAG_ANY      = 4 ///< seek to any frame, even non-keyframes
	AVSEEK_FLAG_FRAME    = 8 ///< seeking based on frame number
)

func AvformatAllocContext() *AVFormatContext {
	return (*AVFormatContext)(C.avformat_alloc_context())
}

func AvformatFreeContext(ctx *AVFormatContext) {
	C.avformat_free_context((*C.struct_AVFormatContext)(unsafe.Pointer(ctx)))
}

func AvformatCloseInput(ctx *AVFormatContext) {
	ptr := (*C.struct_AVFormatContext)(unsafe.Pointer(ctx))
	C.avformat_close_input(&ptr)
}

func AvformatOpenInput(ps **AVFormatContext, url string, fmt *AVInputFormat, options **avutil.AVDictionary) int {
	return int(C.avformat_open_input((**C.struct_AVFormatContext)(unsafe.Pointer(ps)), C.CString(url), (*C.struct_AVInputFormat)(unsafe.Pointer(fmt)), (**C.struct_AVDictionary)(unsafe.Pointer(options))))
}

func AvformatFindStreamInfo(ic *AVFormatContext, options **avutil.AVDictionary) int {
	return int(C.avformat_find_stream_info((*C.struct_AVFormatContext)(unsafe.Pointer(ic)), (**C.struct_AVDictionary)(unsafe.Pointer(options))))
}

func AvReadFrame(s *AVFormatContext, pkt *avcodec.AVPacket) int {
	return int(C.av_read_frame(
		(*C.struct_AVFormatContext)(unsafe.Pointer(s)),
		(*C.struct_AVPacket)(unsafe.Pointer(pkt)),
	))
}

func AvSeekFrame(s *AVFormatContext, streamIndex int, timestamp int64, flags int) int {
	return int(C.av_seek_frame(
		(*C.struct_AVFormatContext)(unsafe.Pointer(s)),
		C.int(streamIndex),
		C.long(timestamp),
		C.int(flags),
	))
}
