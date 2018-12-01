// Based on ssh/terminal:
// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build js

package term

import "io"

// IsTerminal always returns false on JS.
func IsTerminal(w io.Writer) bool {
	return false
}
