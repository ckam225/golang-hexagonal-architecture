package converter

import (
	"clean-arch-hex/internal/domain/entity"
	"clean-arch-hex/pkg/proto"
)

func PostToProto(p entity.Post) *proto.Post {
	return &proto.Post{
		Id:      int32(p.ID),
		Title:   p.Title,
		Content: p.Content,
		// UserId:  int32(*p.UserId),
	}
}

func PostListToProto(posts []entity.Post) []*proto.Post {
	results := make([]*proto.Post, 0, cap(posts))
	for _, p := range posts {
		results = append(results, PostToProto(p))
	}
	return results
}
