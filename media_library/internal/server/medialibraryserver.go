// Code generated by goctl. DO NOT EDIT.
// Source: media_library.proto

package server

import (
	"context"

	"yet-another-media-server/media_library/internal/logic"
	"yet-another-media-server/media_library/internal/svc"
	"yet-another-media-server/media_library/media_library"
)

type MediaLibraryServer struct {
	svcCtx *svc.ServiceContext
	media_library.UnimplementedMediaLibraryServer
}

func NewMediaLibraryServer(svcCtx *svc.ServiceContext) *MediaLibraryServer {
	return &MediaLibraryServer{
		svcCtx: svcCtx,
	}
}

func (s *MediaLibraryServer) CreateLibrary(ctx context.Context, in *media_library.CreateLibraryRequest) (*media_library.Library, error) {
	l := logic.NewCreateLibraryLogic(ctx, s.svcCtx)
	return l.CreateLibrary(in)
}

func (s *MediaLibraryServer) GetLibrary(ctx context.Context, in *media_library.GetLibraryRequest) (*media_library.GetLibraryResponse, error) {
	l := logic.NewGetLibraryLogic(ctx, s.svcCtx)
	return l.GetLibrary(in)
}

func (s *MediaLibraryServer) CreateMedia(ctx context.Context, in *media_library.CreateMediaRequest) (*media_library.Media, error) {
	l := logic.NewCreateMediaLogic(ctx, s.svcCtx)
	return l.CreateMedia(in)
}

func (s *MediaLibraryServer) PaginateMediaByMetadata(ctx context.Context, in *media_library.PaginateMediaByMetadataRequest) (*media_library.PaginateMediaByMetadataResponse, error) {
	l := logic.NewPaginateMediaByMetadataLogic(ctx, s.svcCtx)
	return l.PaginateMediaByMetadata(in)
}

func (s *MediaLibraryServer) GetMedia(ctx context.Context, in *media_library.GetMediaRequest) (*media_library.Media, error) {
	l := logic.NewGetMediaLogic(ctx, s.svcCtx)
	return l.GetMedia(in)
}

func (s *MediaLibraryServer) GetMediaByFilePath(ctx context.Context, in *media_library.GetMediaByFilePathRequest) (*media_library.GetMediaByFilePathResponse, error) {
	l := logic.NewGetMediaByFilePathLogic(ctx, s.svcCtx)
	return l.GetMediaByFilePath(in)
}

func (s *MediaLibraryServer) CreateMetadata(ctx context.Context, in *media_library.CreateMetadataRequest) (*media_library.Metadata, error) {
	l := logic.NewCreateMetadataLogic(ctx, s.svcCtx)
	return l.CreateMetadata(in)
}

func (s *MediaLibraryServer) PaginateMetadata(ctx context.Context, in *media_library.PaginateMetadataRequest) (*media_library.PaginateMetadataResponse, error) {
	l := logic.NewPaginateMetadataLogic(ctx, s.svcCtx)
	return l.PaginateMetadata(in)
}

func (s *MediaLibraryServer) GetMetadataByValue(ctx context.Context, in *media_library.GetMetadataByValueRequest) (*media_library.GetMetadataByValueResponse, error) {
	l := logic.NewGetMetadataByValueLogic(ctx, s.svcCtx)
	return l.GetMetadataByValue(in)
}

func (s *MediaLibraryServer) CreateMediaMetadataRelation(ctx context.Context, in *media_library.CreateMediaMetadataRelationRequest) (*media_library.CreateMediaMetadataRelationResponse, error) {
	l := logic.NewCreateMediaMetadataRelationLogic(ctx, s.svcCtx)
	return l.CreateMediaMetadataRelation(in)
}

func (s *MediaLibraryServer) DeleteMetadata(ctx context.Context, in *media_library.DeleteMetadataRequest) (*media_library.DeleteMetadataResponse, error) {
	l := logic.NewDeleteMetadataLogic(ctx, s.svcCtx)
	return l.DeleteMetadata(in)
}

func (s *MediaLibraryServer) CreateFile(ctx context.Context, in *media_library.CreateFileRequest) (*media_library.File, error) {
	l := logic.NewCreateFileLogic(ctx, s.svcCtx)
	return l.CreateFile(in)
}
