package paasio

import "io"

// Define readCounter and writeCounter types here.
type readCounter struct {
	ioReader io.Reader
}

type writeCounter struct {
	ioWriter io.Writer
}

// For the return of the function NewReadWriteCounter, you must also define a type that satisfies the ReadWriteCounter interface.
type readWriteCounter struct {
	ioReadWriter io.ReadWriter
}

func NewWriteCounter(writer io.Writer) WriteCounter {
	return &writeCounter{writer}
}

func NewReadCounter(reader io.Reader) ReadCounter {
	return &readCounter{reader}
}

func NewReadWriteCounter(readwriter io.ReadWriter) ReadWriteCounter {
	return &readWriteCounter{readwriter}
}

func (rc *readCounter) Read(p []byte) (int, error) {
	return rc.ioReader.Read(p)
}

func (rc *readCounter) ReadCount() (int64, int) {
	panic("Please implement the ReadCount function")
}

func (wc *writeCounter) Write(p []byte) (int, error) {
	return wc.ioWriter.Write(p)
}

func (wc *writeCounter) WriteCount() (int64, int) {
	panic("Please implement the WriteCount function")
}

// Read implements ReadWriteCounter.
func (rwc *readWriteCounter) Read(p []byte) (n int, err error) {
	return rwc.ioReadWriter.Read(p)
}

// ReadCount implements ReadWriteCounter.
func (rwc *readWriteCounter) ReadCount() (n int64, nops int) {
	panic("unimplemented")
}

// Write implements ReadWriteCounter.
func (rwc *readWriteCounter) Write(p []byte) (n int, err error) {
	return rwc.ioReadWriter.Write(p)
}

// WriteCount implements ReadWriteCounter.
func (rwc *readWriteCounter) WriteCount() (n int64, nops int) {
	panic("unimplemented")
}
