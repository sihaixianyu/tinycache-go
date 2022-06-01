package hashring

import (
	"hash/crc32"
	"log"
	"sort"
	"strconv"
)

type HashFn func(data []byte) uint32

type HashRing struct {
	hashFn      HashFn
	numReplicas int
	ring        []int
	codeToKey   map[int]string
}

func New(numReplicas int, hashFn HashFn) *HashRing {
	m := &HashRing{
		numReplicas: numReplicas,
		hashFn:      hashFn,
		codeToKey:   make(map[int]string),
	}

	if m.hashFn == nil {
		m.hashFn = crc32.ChecksumIEEE
	}

	return m
}

func (m *HashRing) Put(keys ...string) {
	for _, k := range keys {
		for i := 0; i < m.numReplicas; i++ {
			mixedKey := []byte(strconv.Itoa(i) + k)
			hashCode := int(m.hashFn(mixedKey))
			m.ring = append(m.ring, hashCode)
			m.codeToKey[hashCode] = k
		}
	}
	sort.Ints(m.ring)
}

func (m *HashRing) Get(key string) string {
	if len(m.ring) == 0 {
		return ""
	}

	hashCode := int(m.hashFn([]byte(key)))
	idx := sort.Search(len(m.ring), func(i int) bool {
		return m.ring[i] >= hashCode
	})
	log.Println(m)

	return m.codeToKey[m.ring[idx%len(m.ring)]]
}
