package paasio

import (
	"io"
	"sync"
)

// Define readCounter and writeCounter types here.
type readCounter struct {
	ioReader io.Reader
	n        int64
	nops     int
	mu       sync.Mutex
}

type writeCounter struct {
	ioWriter io.Writer
	n        int64
	nops     int
	mu       sync.Mutex
}

// For the return of the function NewReadWriteCounter, you must also define a type that satisfies the ReadWriteCounter interface.
type readWriteCounter struct {
	ioReadWriter io.ReadWriter
	n            int64
	nops         int
	mu           sync.Mutex
}

func NewWriteCounter(writer io.Writer) WriteCounter {
	return &writeCounter{ioWriter: writer}
}

func NewReadCounter(reader io.Reader) ReadCounter {
	return &readCounter{ioReader: reader}
}

func NewReadWriteCounter(readwriter io.ReadWriter) ReadWriteCounter {
	return &readWriteCounter{ioReadWriter: readwriter}
}

func (rc *readCounter) Read(p []byte) (int, error) {
	rc.mu.Lock()
	defer rc.mu.Unlock()
	n, err := rc.ioReader.Read(p)
	rc.n += int64(n)
	rc.nops++
	return n, err
}

func (rc *readCounter) ReadCount() (int64, int) {
	rc.mu.Lock()
	defer rc.mu.Unlock()
	return rc.n, rc.nops
}

func (wc *writeCounter) Write(p []byte) (int, error) {
	wc.mu.Lock()
	defer wc.mu.Unlock()
	n, err := wc.ioWriter.Write(p)
	wc.n += int64(n)
	wc.nops++
	return n, err
}

func (wc *writeCounter) WriteCount() (int64, int) {
	wc.mu.Lock()
	defer wc.mu.Unlock()
	return wc.n, wc.nops
}

// Read implements ReadWriteCounter.
func (rwc *readWriteCounter) Read(p []byte) (int, error) {
	rwc.mu.Lock()
	defer rwc.mu.Unlock()
	n, err := rwc.ioReadWriter.Read(p)
	rwc.n += int64(n)
	rwc.nops++
	return n, err
}

// ReadCount implements ReadWriteCounter.
func (rwc *readWriteCounter) ReadCount() (int64, int) {
	rwc.mu.Lock()
	defer rwc.mu.Unlock()
	return rwc.n, rwc.nops
}

// Write implements ReadWriteCounter.
func (rwc *readWriteCounter) Write(p []byte) (int, error) {
	rwc.mu.Lock()
	defer rwc.mu.Unlock()
	n, err := rwc.ioReadWriter.Write(p)
	rwc.n += int64(n)
	rwc.nops++
	return n, err
}

// WriteCount implements ReadWriteCounter.
func (rwc *readWriteCounter) WriteCount() (int64, int) {
	return rwc.ReadCount()
}

