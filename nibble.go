package nibble

import "io"

// Writer writes half bytes to an io.Writer.
type Writer struct {
	w    io.Writer
	half bool // true if the higher nibble is set
	data byte // will be written when there are two
}

// New creates a new Writer.
func New(w io.Writer) *Writer {
	return &Writer{w: w}
}

// Write writes a nibble to the writer. Only the first half of the bytes is
// written.
func (nw *Writer) Write(bs []byte) (int, error) {
	for _, b := range bs {
		if n, err := nw.write(b); err != nil {
			return n, err
		}
	}
	return len(bs), nil
}

func (nw *Writer) write(b byte) (int, error) {
	if !nw.half {
		nw.data = b >> 4
		nw.half = true
		return 0, nil
	}
	nw.data += b & 0xf0
	nw.half = false
	return nw.w.Write([]byte{nw.data})
}

// Flush writes tle last nibble.
func (nw *Writer) Flush() (int, error) {
	if !nw.half {
		return 0, nil
	}
	return nw.w.Write([]byte{nw.data})
}
