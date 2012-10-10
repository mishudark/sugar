/*
  Copyright (c) 2012 José Carlos Nieto, http://xiam.menteslibres.org/

  Permission is hereby granted, free of charge, to any person obtaining
  a copy of this software and associated documentation files (the
  "Software"), to deal in the Software without restriction, including
  without limitation the rights to use, copy, modify, merge, publish,
  distribute, sublicense, and/or sell copies of the Software, and to
  permit persons to whom the Software is furnished to do so, subject to
  the following conditions:

  The above copyright notice and this permission notice shall be
  included in all copies or substantial portions of the Software.

  THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
  EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
  MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
  NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
  LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
  OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
  WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
*/

package sugar

import (
	"strings"
)

const Separator = "/"

// Trivial tuple
type Tuple map[string]interface{}

// Trivial list
type List []interface{}

func getPath(genmap map[string]interface{}, path string) interface{} {
	chunks := strings.Split(path, Separator)

	switch len(chunks) {
	case 0:
		return nil
	case 1:
		return genmap[chunks[0]]
	default:
		switch genmap[chunks[0]].(type) {
		case map[string]interface{}:
			return getPath(genmap[chunks[0]].(map[string]interface{}), strings.Join(chunks[1:], Separator))
		default:
			return nil
		}
	}

	return nil
}

// Digs deeper in a sugar.Tuple{}, useful when dealing with JSON.
func (self *Tuple) Get(path string) interface{} {
	return getPath(*self, path)
}
