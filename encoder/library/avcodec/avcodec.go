package avcodec

//#cgo pkg-config: libavcodec
//
//#include <libavcodec/avcodec.h>
//#include "stdlib.h"
import "C"
import (
	"unsafe"
	"yet-another-media-server/encoder/library/avutil"
)

type (
	AVCodecParameters C.struct_AVCodecParameters
	AVCodec           C.struct_AVCodec
	AVCodecContext    C.struct_AVCodecContext
	AVPacket          C.struct_AVPacket
)

func AvcodecFindDecoder(id AVCodecID) *AVCodec {
	return (*AVCodec)(C.avcodec_find_decoder(uint32(id)))
}

func AvcodecAllocContext3(codec *AVCodec) *AVCodecContext {
	return (*AVCodecContext)(C.avcodec_alloc_context3((*C.struct_AVCodec)(unsafe.Pointer(codec))))
}

func AvcodecFreeContext(ctx *AVCodecContext) {
	ptr := (*C.struct_AVCodecContext)(unsafe.Pointer(ctx))
	C.avcodec_free_context(&ptr)
}

func AvcodecParametersToContext(codec *AVCodecContext, par *AVCodecParameters) int {
	return int(C.avcodec_parameters_to_context(
		(*C.struct_AVCodecContext)(unsafe.Pointer(codec)),
		(*C.struct_AVCodecParameters)(unsafe.Pointer(par)),
	))
}

func AvcodecOpen2(avctx *AVCodecContext, codec *AVCodec, options **avutil.AVDictionary) int {
	return int(C.avcodec_open2(
		(*C.struct_AVCodecContext)(unsafe.Pointer(avctx)),
		(*C.struct_AVCodec)(unsafe.Pointer(codec)),
		(**C.struct_AVDictionary)(unsafe.Pointer(options)),
	))
}

func AvPacketAlloc() *AVPacket {
	return (*AVPacket)(C.av_packet_alloc())
}

func AvPacketFree(pkt *AVPacket) {
	//var t *C.struct_AVPacket
	//t = (*C.struct_AVPacket)(unsafe.Pointer(pkt))
	//ptr := C.malloc(C.size_t(unsafe.Sizeof(t)))
	//defer C.free(ptr)
	ptr := (*C.struct_AVPacket)(unsafe.Pointer(pkt))
	C.av_packet_free(&ptr)
}

func AvPacketUnref(pkt *AVPacket) {
	C.av_packet_unref((*C.struct_AVPacket)(unsafe.Pointer(pkt)))
}

func AvcodecSendPacket(avctx *AVCodecContext, avpkt *AVPacket) int {
	return int(C.avcodec_send_packet(
		(*C.struct_AVCodecContext)(unsafe.Pointer(avctx)),
		(*C.struct_AVPacket)(unsafe.Pointer(avpkt)),
	))
}

func AvcodecReceiveFrame(avctx *AVCodecContext, frame *avutil.AVFrame) int {
	return int(C.avcodec_receive_frame(
		(*C.struct_AVCodecContext)(unsafe.Pointer(avctx)),
		(*C.struct_AVFrame)(unsafe.Pointer(frame)),
	))
}

func AvcodecFlushBuffers(avctx *AVCodecContext) {
	C.avcodec_flush_buffers((*C.struct_AVCodecContext)(unsafe.Pointer(avctx)))
}
