package ch6

import "io"

type countWriter struct {
	writer io.Writer
	count  int64
}

func (c *countWriter) Write(p []byte) (int, error) {
	n, err := c.writer.Write(p)
	c.count += int64(n)
	return n, err
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	cw := &countWriter{w, 0}
	return cw, &cw.count
}
