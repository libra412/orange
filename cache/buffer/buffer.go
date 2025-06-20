// Byte 复用
package buffer

import "io"

// IoBuffer 接口定义
type IoBuffer interface {
	// Read reads the next len(p) bytes from the buffer or until the buffer
	// is drained. The return value n is the number of bytes read. If the
	// buffer has no data to return, err is io.EOF (unless len(p) is zero);
	// otherwise it is nil.
	Read(p []byte) (n int, err error)

	// ReadOnce make a one-shot read and appends it to the buffer, growing
	// the buffer as needed. The return value n is the number of bytes read. Any
	// error except io.EOF encountered during the read is also returned. If the
	// buffer becomes too large, ReadFrom will panic with ErrTooLarge.
	ReadOnce(r io.Reader) (n int64, err error)

	// ReadFrom reads data from r until EOF and appends it to the buffer, growing
	// the buffer as needed. The return value n is the number of bytes read. Any
	// error except io.EOF encountered during the read is also returned. If the
	// buffer becomes too large, ReadFrom will panic with ErrTooLarge.
	ReadFrom(r io.Reader) (n int64, err error)

	// Grow updates the length of the buffer by n, growing the buffer as
	// needed. The return value n is the length of p; err is always nil. If the
	// buffer becomes too large, Write will panic with ErrTooLarge.
	Grow(n int) error

	// Write appends the contents of p to the buffer, growing the buffer as
	// needed. The return value n is the length of p; err is always nil. If the
	// buffer becomes too large, Write will panic with ErrTooLarge.
	Write(p []byte) (n int, err error)

	// WriteString appends the string to the buffer, growing the buffer as
	// needed. The return value n is the length of s; err is always nil. If the
	// buffer becomes too large, Write will panic with ErrTooLarge.
	WriteString(s string) (n int, err error)

	// WriteByte appends the byte to the buffer, growing the buffer as
	// needed. The return value n is the length of s; err is always nil. If the
	// buffer becomes too large, Write will panic with ErrTooLarge.
	WriteByte(p byte) error

	// WriteUint16 appends the uint16 to the buffer, growing the buffer as
	// needed. The return value n is the length of s; err is always nil. If the
	// buffer becomes too large, Write will panic with ErrTooLarge.
	WriteUint16(p uint16) error

	// WriteUint32 appends the uint32 to the buffer, growing the buffer as
	// needed. The return value n is the length of s; err is always nil. If the
	// buffer becomes too large, Write will panic with ErrTooLarge.
	WriteUint32(p uint32) error

	// WriteUint64 appends the uint64 to the buffer, growing the buffer as
	// needed. The return value n is the length of s; err is always nil. If the
	// buffer becomes too large, Write will panic with ErrTooLarge.
	WriteUint64(p uint64) error

	// WriteTo writes data to w until the buffer is drained or an error occurs.
	// The return value n is the number of bytes written; it always fits into an
	// int, but it is int64 to match the io.WriterTo interface. Any error
	// encountered during the write is also returned.
	WriteTo(w io.Writer) (n int64, err error)

	// Peek returns n bytes from buffer, without draining any buffered data.
	// If n > readable buffer, nil will be returned.
	// It can be used in codec to check first-n-bytes magic bytes
	// Note: do not change content in return bytes, use write instead
	Peek(n int) []byte

	// Bytes returns all bytes from buffer, without draining any buffered data.
	// It can be used to get fixed-length content, such as headers, body.
	// Note: do not change content in return bytes, use write instead
	Bytes() []byte

	// Drain drains a offset length of bytes in buffer.
	// It can be used with Bytes(), after consuming a fixed-length of data
	Drain(offset int)

	// Len returns the number of bytes of the unread portion of the buffer;
	// b.Len() == len(b.Bytes()).
	Len() int

	// Cap returns the capacity of the buffer's underlying byte slice, that is, the
	// total space allocated for the buffer's data.
	Cap() int

	// Reset resets the buffer to be empty,
	// but it retains the underlying storage for use by future writes.
	Reset()

	// Clone makes a copy of IoBuffer struct
	Clone() IoBuffer

	// String returns the contents of the unread portion of the buffer
	// as a string. If the Buffer is a nil pointer, it returns "<nil>".
	String() string

	// Alloc alloc bytes from BytePoolBuffer
	Alloc(int)

	// Free free bytes to BytePoolBuffer
	Free()

	// Count sets and returns reference count
	Count(int32) int32

	// EOF returns whether Io is EOF on the connection
	EOF() bool

	// SetEOF sets the IoBuffer EOF
	SetEOF(eof bool)
	Append(data []byte) error
	CloseWithError(err error)
}
