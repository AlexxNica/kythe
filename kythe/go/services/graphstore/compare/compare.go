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

// Package compare implements comparisons between Kythe values as used in the
// implementation of a graphstore.Service.
package compare

import (
	"bytes"

	spb "kythe/proto/storage_proto"
)

// An Order represents an ordering relationship between values.
type Order int

// LT, EQ, and GT are the standard values for an Order.
const (
	LT Order = -1 // lhs < rhs
	EQ Order = 0  // lhs == rhs
	GT Order = 1  // lhs > rhs
)

// Strings returns LT if s < t, EQ if s == t, or GT if s > t.
func Strings(s, t string) Order {
	switch {
	case s < t:
		return LT
	case s > t:
		return GT
	default:
		return EQ
	}
}

// VNames returns LT if v1 precedes v2, EQ if v1 and v2 are equal, or GT if v1
// follows v2, in standard order.  The ordering for VNames is defined by
// lexicographic comparison of [signature, corpus, root, path, language].
func VNames(v1, v2 *spb.VName) Order {
	if v1 == v2 {
		return EQ
	} else if c := Strings(v1.GetSignature(), v2.GetSignature()); c != EQ {
		return c
	} else if c := Strings(v1.GetCorpus(), v2.GetCorpus()); c != EQ {
		return c
	} else if c := Strings(v1.GetRoot(), v2.GetRoot()); c != EQ {
		return c
	} else if c := Strings(v1.GetPath(), v2.GetPath()); c != EQ {
		return c
	}
	return Strings(v1.GetLanguage(), v2.GetLanguage())
}

// VNamesEqual reports whether v1 and v2 are equal.
func VNamesEqual(v1, v2 *spb.VName) bool { return VNames(v1, v2) == EQ }

// Entries reports whether e1 is LT, GT, or EQ to e2 in entry order, ignoring
// fact values (if any).
//
// The ordering for entries is defined by lexicographic comparison of
// [source, edge kind, fact name, target].
func Entries(e1, e2 *spb.Entry) Order {
	if e1 == e2 {
		return EQ
	}
	if c := VNames(e1.GetSource(), e2.GetSource()); c != EQ {
		return c
	} else if c := Strings(e1.GetEdgeKind(), e2.GetEdgeKind()); c != EQ {
		return c
	} else if c := Strings(e1.GetFactName(), e2.GetFactName()); c != EQ {
		return c
	}
	return VNames(e1.GetTarget(), e2.GetTarget())
}

// ValueEntries reports whether e1 is LT, GT, or EQ to e2 in entry order,
// including fact values (if any).
func ValueEntries(e1, e2 *spb.Entry) Order {
	if c := Entries(e1, e2); c != EQ {
		return c
	}
	return Order(bytes.Compare(e1.GetFactValue(), e2.GetFactValue()))
}

// EntriesEqual reports whether e1 and e2 are equivalent, including their fact
// values (if any).
func EntriesEqual(e1, e2 *spb.Entry) bool { return ValueEntries(e1, e2) == EQ }