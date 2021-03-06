// Copyright 2020 VMware, Inc.
// SPDX-License-Identifier: Apache-2.0

package imageset

import (
	"fmt"
	"sort"

	regname "github.com/google/go-containerregistry/pkg/name"
)

type UnprocessedImageRef struct {
	DigestRef string
	Tag       string
}

type UnprocessedImageRefs struct {
	imgRefs map[UnprocessedImageRef]struct{}
}

func NewUnprocessedImageRefs() *UnprocessedImageRefs {
	return &UnprocessedImageRefs{map[UnprocessedImageRef]struct{}{}}
}

func (i *UnprocessedImageRefs) Add(imgRef UnprocessedImageRef) {
	imgRef.Validate()

	i.imgRefs[imgRef] = struct{}{}
}

func (i *UnprocessedImageRefs) All() []UnprocessedImageRef {
	var result []UnprocessedImageRef
	for imgRef := range i.imgRefs {
		result = append(result, imgRef)
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i].DigestRef < result[j].DigestRef
	})
	return result
}

func (u UnprocessedImageRef) Validate() {
	_, err := regname.NewDigest(u.DigestRef)
	if err != nil {
		panic(fmt.Sprintf("Digest need to be provided: %s", err))
	}
}
