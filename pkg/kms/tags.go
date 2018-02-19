package kms

import "errors"

func (k *kms) ListResourceTags(req *ListResourceTagsRequest) (*ListResourceTagsResult, error) {
	k.lock.Lock()
	defer k.lock.Unlock()

	if req.Marker != "" {
		return nil, errors.New("InvalidMarkerException")
	}
	if req.Limit != 0 {
		return nil, errors.New("LimitNotSupported")
	}

	key := k.get(req.KeyID)
	if key == nil {
		return nil, errors.New("NotFoundException")
	}

	tags := []Tag{}
	for key, value := range key.tags {
		tags = append(tags, Tag{TagKey: key, TagValue: value})
	}

	return &ListResourceTagsResult{
		Tags:      tags,
		Truncated: false,
	}, nil
}

func (k *kms) TagResource(req *TagResourceRequest) error {
	k.lock.Lock()
	defer k.lock.Unlock()

	key := k.get(req.KeyID)
	if key == nil {
		return errors.New("NotFoundException")
	}

	for _, tag := range req.Tags {
		key.tags[tag.TagKey] = tag.TagValue
	}

	return nil
}

func (k *kms) UntagResource(req *UntagResourceRequest) error {
	k.lock.Lock()
	defer k.lock.Unlock()

	key := k.get(req.KeyID)
	if key == nil {
		return errors.New("NotFoundException")
	}

	for _, tag := range req.TagKeys {
		delete(key.tags, tag)
	}

	return nil
}
