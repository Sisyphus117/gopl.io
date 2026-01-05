package ch6

import "io"

type limitReader struct {
	reader io.Reader
	limit  int64
}

func (l *limitReader) Read(p []byte) (int, error) {
	if int64(len(p)) <= l.limit {
		l.limit -= int64(len(p))
		n, err := l.reader.Read(p)
		return n, err
	} else {
		n, err := l.reader.Read(p[:l.limit])
		l.limit = 0
		return n, err
	}
}

func LimitReader(r io.Reader, n int64) io.Reader {
	return &limitReader{r, n}
}
