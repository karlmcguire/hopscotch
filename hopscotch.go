package hopscotch

import "fmt"

const (
	H = 4
)

type Bucket struct {
	Bitmap uint64
	Key    uint64
	Value  []byte
}

type Map struct {
	buckets []Bucket
	mask    uint64
}

func New(size uint64) *Map {
	return &Map{
		buckets: make([]Bucket, size),
		mask:    size - 1,
	}
}

func (m *Map) Set(key uint64, value []byte) {
}

func (m *Map) Get(key uint64) []byte {
	return nil
}

func (m *Map) String() string {
	var out string
	for i, bucket := range m.buckets {
		out += fmt.Sprintf(
			"%2d: [%02d] = '%s'\n", i, bucket.Key, string(bucket.Value),
		)
	}
	return out
}
