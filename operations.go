package m3u8

import (
	"strings"
)

func (mp *MediaPlaylist) FindSegmentsByName(name string) []*MediaSegment {
	segments := make([]*MediaSegment, 0)
	for _, segment := range mp.Segments {
		if strings.ToLower(segment.Title) == strings.ToLower(name) {
			segments = append(segments, segment)
		}
	}
	return segments
}

func (mp *MediaPlaylist) FindSegmentByTag(key, value string) []*MediaSegment {
	segments := make([]*MediaSegment, 0)
	for _, segment := range mp.Segments {
		for _, tag := range segment.CustomTags {
			if strings.ToLower(tag.Key) == strings.ToLower(key) && strings.ToLower(tag.Value) == strings.ToLower(value) {
				segments = append(segments, segment)
			}
		}
	}
	return segments
}

func (ms *MediaSegment) GetTag(key string) *CustomTag {
	for _, tag := range ms.CustomTags {
		if strings.ToLower(tag.Key) == strings.ToLower(key) {
			return tag
		}
	}
	return nil
}

func (ms *MediaSegment) SetTag(key, value string) bool {
	for _, tag := range ms.CustomTags {
		if strings.ToLower(tag.Key) == strings.ToLower(key) {
			tag.Value = value
			return true
		} else {
			ms.CustomTags = append(ms.CustomTags, &CustomTag{key, value})
			return true
		}
	}
	return false
}
