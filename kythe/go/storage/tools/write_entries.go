/*
 * Copyright 2014 Google Inc. All rights reserved.
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

// Binary write_entries reads a stream of protobuf entries on os.Stdin and
// writes each to a graphstore server.
//
// Usage:
//   entry_emitter ... | write_entries --graphstore addr
//
// Example:
//   java_indexer_server --port 8181 &
//   graphstore --port 9999 &
//   analysis_driver --analyzer localhost:8181 /tmp/compilation.kindex | \
//     write_entries --workers 10 --graphstore localhost:9999
//
// Example:
//   zcat entries.gz | write_entries --graphstore gs/leveldb
package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime/pprof"
	"sync"
	"sync/atomic"

	"kythe/go/services/graphstore"
	"kythe/go/storage/gsutil"
	"kythe/go/storage/stream"

	spb "kythe/proto/storage_proto"
)

func init() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, `%s - writes the entries from a delimited stream on stdin to a GraphStore
usage: %[1]s [--batch_size entries] [--workers n] --graphstore spec
`, filepath.Base(os.Args[0]))
		flag.PrintDefaults()
		os.Exit(1)
	}
}

var (
	batchSize  = flag.Int("batch_size", 1024, "Maximum entries per write for consecutive entries with the same source")
	numWorkers = flag.Int("workers", 1, "Number of concurrent workers writing to the GraphStore")
	profCPU    = flag.String("cpu_profile", "", "Write CPU profile to the specified file (if nonempty)")

	gs graphstore.Service
)

func init() {
	gsutil.Flag(&gs, "graphstore", "GraphStore to which to write the entry stream")
}

func main() {
	log.SetPrefix("write_entries: ")

	flag.Parse()
	if *numWorkers < 1 {
		gsutil.UsageErrorf("Invalid number of --workers %d (must be ≥ 1)", *numWorkers)
	} else if *batchSize < 1 {
		gsutil.UsageErrorf("Invalid --batch_size %d (must be ≥ 1)", *batchSize)
	} else if gs == nil {
		gsutil.UsageError("Missing --graphstore")
	}

	defer gsutil.LogClose(gs)
	gsutil.EnsureGracefulExit(gs)

	if *profCPU != "" {
		f, err := os.Create(*profCPU)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	writes := graphstore.BatchWrites(stream.ReadEntries(os.Stdin), *batchSize)

	var (
		wg         sync.WaitGroup
		numEntries uint64
	)
	wg.Add(*numWorkers)
	for i := 0; i < *numWorkers; i++ {
		go func() {
			defer wg.Done()
			num, err := writeEntries(gs, writes)
			if err != nil {
				log.Fatal(err)
			}

			atomic.AddUint64(&numEntries, num)
		}()
	}
	wg.Wait()

	log.Printf("Wrote %d entries", numEntries)
}

func writeEntries(s graphstore.Service, reqs <-chan *spb.WriteRequest) (uint64, error) {
	var num uint64

	for req := range reqs {
		num += uint64(len(req.Update))
		if err := s.Write(req); err != nil {
			return 0, err
		}
	}

	return num, nil
}