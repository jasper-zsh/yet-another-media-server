package types

import "yet-another-media-server/media_library/medialibraryclient"

func NewLibraryFromProto(proto *medialibraryclient.Library) *Library {
	return &Library{
		ID:        proto.Id,
		CreatedAt: proto.CreatedAt,
		UpdatedAt: proto.UpdatedAt,
		Name:      proto.Name,
		BasePath:  proto.BasePath,
	}
}

func NewMetadataFromProto(proto *medialibraryclient.Metadata) *Metadata {
	result := &Metadata{
		ID:        proto.Id,
		CreatedAt: proto.CreatedAt,
		UpdatedAt: proto.UpdatedAt,
		Type:      proto.Type,
		Value:     proto.Value,
		Children:  make(map[string][]*Metadata),
	}
	for _, p := range proto.Children {
		list, ok := result.Children[p.Type]
		if !ok {
			list = make([]*Metadata, 0)
		}
		list = append(list, NewMetadataFromProto(p))
		result.Children[p.Type] = list
	}
	return result
}

func NewFileFromProto(proto *medialibraryclient.File) *File {
	result := &File{
		ID:        proto.Id,
		CreatedAt: proto.CreatedAt,
		UpdatedAt: proto.UpdatedAt,
		Type:      proto.Type,
		FilePath:  proto.FilePath,
		Ext:       proto.Ext,
		Children:  make(map[string][]*File),
	}
	for _, p := range proto.Children {
		list, ok := result.Children[p.Type]
		if !ok {
			list = make([]*File, 0)
		}
		list = append(list, NewFileFromProto(p))
		result.Children[p.Type] = list
	}
	return result
}

func NewMediaFromProto(proto *medialibraryclient.Media) *Media {
	result := &Media{
		ID:        proto.Id,
		CreatedAt: proto.CreatedAt,
		UpdatedAt: proto.UpdatedAt,
		Metadatas: make(map[string][]*Metadata),
		Files:     make([]*File, 0, len(proto.Files)),
	}
	for _, p := range proto.Metadatas {
		list, ok := result.Metadatas[p.Type]
		if !ok {
			list = make([]*Metadata, 0)
		}
		list = append(list, NewMetadataFromProto(p))
		result.Metadatas[p.Type] = list
	}
	for _, p := range proto.Files {
		result.Files = append(result.Files, NewFileFromProto(p))
	}
	return result
}
