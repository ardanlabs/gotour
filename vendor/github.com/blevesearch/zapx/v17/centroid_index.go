//  Copyright (c) 2026 Couchbase, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// 		http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//go:build vectors
// +build vectors

package zap

import (
	"encoding/binary"
	"fmt"

	index "github.com/blevesearch/bleve_index_api"
)

func (sb *SegmentBase) GetCoarseQuantizer(field string) (interface{}, error) {
	fieldIDPlus1 := sb.fieldsMap[field]
	if fieldIDPlus1 <= 0 {
		return nil, fmt.Errorf("field %s does not exist in segment", field)
	}

	vectorSection := sb.fieldsSectionsMap[fieldIDPlus1-1][SectionFaissVectorIndex]
	// check if the field has a vector section in the segment.
	if vectorSection <= 0 {
		return nil, fmt.Errorf("field %s does not have a vector section in the segment", field)
	}

	pos := vectorSection
	// doc values
	for i := 0; i < 2; i++ {
		_, n := binary.Uvarint(sb.mem[pos : pos+binary.MaxVarintLen64])
		pos += uint64(n)
	}
	// read the index optimization type
	opt, n := binary.Uvarint(sb.mem[pos : pos+binary.MaxVarintLen64])
	pos += uint64(n)
	// get the optimization type string from the reverse lookup map
	optStr := index.VectorIndexOptimizationsReverseLookup[int(opt)]

	numVecs, n := binary.Uvarint(sb.mem[pos : pos+binary.MaxVarintLen64])
	pos += uint64(n)

	// length of the vector to docID map
	mapLen, n := binary.Uvarint(sb.mem[pos : pos+binary.MaxVarintLen64])
	pos += uint64(n)
	pos += mapLen

	// type of index
	indexType, n := binary.Uvarint(sb.mem[pos : pos+binary.MaxVarintLen64])
	pos += uint64(n)
	indexSize, n := binary.Uvarint(sb.mem[pos : pos+binary.MaxVarintLen64])
	pos += uint64(n)

	params := newFaissIndexParams(optStr, int(numVecs), faissIOFlagsReadOnly)

	// todo: might wanna use the vector cache here, early tests didn't show a big diff
	fIndexBytes, err := sb.fileReader.process(sb.mem[pos : pos+indexSize])
	if err != nil {
		return nil, err
	}
	pos += indexSize

	if faissIndexType(indexType) == faissBIVFIndex {
		binaryIndexSize, n := binary.Uvarint(sb.mem[pos : pos+binary.MaxVarintLen64])
		pos += uint64(n)
		bIndexBytes, err := sb.fileReader.process(sb.mem[pos : pos+binaryIndexSize])
		if err != nil {
			return nil, err
		}
		return newFaissBinaryIndexFromBytes(bIndexBytes, fIndexBytes, params)
	}
	return newFaissFloat32IndexFromBytes(fIndexBytes, params)
}
