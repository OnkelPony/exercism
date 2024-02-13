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
	mu       sync.RWMutex
}

type writeCounter struct {
	ioWriter io.Writer
	n        int64
	nops     int
	mu       sync.RWMutex
}

// For the return of the function NewReadWriteCounter, you must also define a type that satisfies the ReadWriteCounter interface.
type readWriteCounter struct {
	readCounter
	writeCounter
}

func NewWriteCounter(writer io.Writer) WriteCounter {
	return &writeCounter{ioWriter: writer}
}

func NewReadCounter(reader io.Reader) ReadCounter {
	return &readCounter{ioReader: reader}
}

func NewReadWriteCounter(readwriter io.ReadWriter) ReadWriteCounter {
	return &readWriteCounter{readCounter: readCounter{ioReader: readwriter}, writeCounter: writeCounter{ioWriter: readwriter}}
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
	rc.mu.RLock()
	defer rc.mu.RUnlock()
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
	wc.mu.RLock()
	defer wc.mu.RUnlock()
	return wc.n, wc.nops
}

// Read implements ReadWriteCounter.
func (rwc *readWriteCounter) Read(p []byte) (int, error) {
	return rwc.readCounter.Read(p)
	}

// ReadCount implements ReadWriteCounter.
func (rwc *readWriteCounter) ReadCount() (int64, int) {
	return rwc.readCounter.ReadCount()
}

// Write implements ReadWriteCounter.
func (rwc *readWriteCounter) Write(p []byte) (int, error) {
	return rwc.writeCounter.Write(p)
}

// WriteCount implements ReadWriteCounter.
func (rwc *readWriteCounter) WriteCount() (int64, int) {
	return rwc.writeCounter.WriteCount()
}

