package hopscotch

import (
	"fmt"
	"math/bits"
)

const (
	H = 4
)

type bucket struct {
	bitmap uint64
	key    uint64
	value  []byte
}

type Map struct {
	buckets []bucket
	mask    uint64
}

func New(size uint64) *Map {
	return &Map{
		buckets: make([]bucket, size),
		mask:    size - 1,
	}
}

func (m *Map) Set(key uint64, value []byte) bool {
	index := key & m.mask

	// if the key already exists
	if m.lookup(key, index) != nil {
		return false
	}

	for i := index; i <= m.mask; i++ {
		bucket := &m.buckets[i]
		if bucket.value == nil {
			for j := uint64(1); i-index >= H; {
				for ; j < H; j++ {
					if m.buckets[i-j].bitmap != 0 {
						offset := uint64(
							bits.TrailingZeros64(m.buckets[i-j].bitmap))
						if offset >= j {
							continue
						}

						m.buckets[i].key = m.buckets[i-j+offset].key
						m.buckets[i].value = m.buckets[i-j+offset].value

						m.buckets[i-j+offset].key = 0
						m.buckets[i-j+offset].value = nil

						m.buckets[i-j].bitmap &= ^(1 << offset)
						m.buckets[i-j].bitmap |= 1 << j

						i = i - j + offset
						break
					}
				}
				if j >= H {
					return false
				}
			}
			// add the item
			bucket.key = key
			bucket.value = value
			setBitmap(&bucket.bitmap, i-index)
			return true
		}
	}
	return false
}

func (m *Map) Get(key uint64) []byte {
	if bucket := m.lookup(key, key&m.mask); bucket == nil {
		return nil
	} else {
		return bucket.value
	}
}

func (m *Map) String() string {
	var out string
	for i, bucket := range m.buckets {
		out += fmt.Sprintf("%2d: [%02d]-[%04b] = %s\n", i,
			bucket.key,
			bucket.bitmap,
			string(bucket.value),
		)
	}
	return out
}

func (m *Map) lookup(key, index uint64) *bucket {
	for i := uint64(0); i < H && (index+i <= m.mask); i++ {
		bucket := &m.buckets[index+i]
		if bucket.bitmap&(1<<i) != 0 {
			if bucket.key == key {
				return bucket
			}
		}
	}
	return nil
}

func setBitmap(bitmap *uint64, offset uint64) {
	*bitmap |= 1 << offset
}
