package streamsplit

import (
	"io"
	"sync"
)

// New returns a splitter, injecting sep every length bytes written to w.
func New(length int, sep []byte, w io.Writer) *splitter {
	if length < 1 {
		panic(`length is less than 1`)
	}
	if len(sep) < 1 {
		panic(`sep is empty`)
	}
	if w == nil {
		panic(`w is nil`)
	}

	return &splitter{
		len: length,
		sep: sep,
		w:   w,
	}
}

// splitter implements io.Writer
var _ io.Writer = (*splitter)(nil)

// splitter injects sep every len bytes written to w.
// For example to chunk the output of a base64.Encoder.
type splitter struct {
	// len indicates where to split data.
	len int
	// sep is the separator used to split data.
	sep []byte
	// w is the Writer the split data is written to.
	w io.Writer
	// mutex serializes writes due to access to splitter.pos
	mutex sync.Mutex
	// pos remembers how long the current line written to w is.
	pos int
}

func (l *splitter) Write(p []byte) (n int, err error) {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	var nn int
	for l.pos+len(p) > l.len {
		// write partial data
		nn, err = l.w.Write(p[:l.len-l.pos])
		n += nn
		if err != nil {
			return
		}

		// keep remainder
		p = p[l.len-l.pos:]
		l.pos = 0

		// write separator
		if _, err = l.w.Write(l.sep); err != nil {
			return
		}
	}

	// write remainder
	nn, err = l.w.Write(p)
	l.pos += nn
	n += nn

	return
}
