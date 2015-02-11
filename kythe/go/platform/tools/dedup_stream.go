/*
 * Copyright 2015 Google Inc. All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

// Binary dedup_stream reads a delimited stream from stdin and writes a delimited stream to stdout.
// Each record in the stream will be hashed, and if that hash value has already been seen, the
// record will not be emitted.
package main

import (
	"crypto/sha512"
	"io"
	"log"
	"os"

	"kythe/go/platform/delimited"
)

func main() {
	written := make(map[[sha512.Size384]byte]struct{})

	var skipped uint64
	rd := delimited.NewReader(os.Stdin)
	wr := delimited.NewWriter(os.Stdout)
	for {
		rec, err := rd.Next()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		hash := sha512.Sum384(rec)
		if _, ok := written[hash]; ok {
			skipped++
			continue
		}
		if err := wr.Put(rec); err != nil {
			log.Fatal(err)
		}
		written[hash] = struct{}{}
	}
	log.Printf("dedup_stream: skipped %d records", skipped)
}