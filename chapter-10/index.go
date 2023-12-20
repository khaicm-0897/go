// control-statements/boolean-expressions/main.go
package main

import (
	"fmt"
	"math/rand"
	"time"
	"unicode"
	"unicode/utf8"
)

// TrimSpace returns a slice of the string s, with all leading
// and trailing white space removed, as defined by Unicode.
func TrimSpace(s string) string {
	// Fast path for ASCII: look for the first ASCII non-space byte
	start := 0
	for ; start < len(s); start++ {
	   c := s[start]
	   if c >= utf8.RuneSelf {
		  // If we run into a non-ASCII byte, fall back to the
		  // slower Unicode-aware method on the remaining bytes
		  return TrimFunc(s[start:], unicode.IsSpace)
	   }
	   if asciiSpace[c] == 0 {
		  break
	   }
	}
 
	// Now look for the first ASCII non-space byte from the end
	stop := len(s)
	for ; stop > start; stop-- {
	   c := s[stop-1]
	   if c >= utf8.RuneSelf {
		  return TrimFunc(s[start:stop], unicode.IsSpace)
	   }
	   if asciiSpace[c] == 0 {
		  break
	   }
	}
 
	// At this point s[start:stop] starts and ends with an ASCII
	// non-space bytes, so we're done. Non-ASCII cases have already
	// been handled above.
	return s[start:stop]
 }

func main() {
	TrimSpace("abc")
}