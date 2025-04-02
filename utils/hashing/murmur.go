package hashing

import (
	"encoding/binary"
	"math"
)

const (
	c1 uint32 = 0xcc9e2d51
	c2 uint32 = 0x1b873593
	c3 uint32 = 0x85ebca6b
	c4 uint32 = 0xc2b2ae35
)

func MurmurHash3(key []byte, seed uint32) uint32 {
	length := len(key)
	hash := seed

	for i := 0; i < length-3; i += 4 {
		k := binary.LittleEndian.Uint32(key[i:])
		k *= c1
		k = rotateLeft(k, 15)
		k *= c2
		hash ^= k
		hash = rotateLeft(hash, 13)
		hash = hash*5 + 0xe6546b64
	}

	var k uint32
	tailIndex := length - length%4
	switch length & 3 {
	case 3:
		k ^= uint32(key[tailIndex+2]) << 16
		fallthrough
	case 2:
		k ^= uint32(key[tailIndex+1]) << 8
		fallthrough
	case 1:
		k ^= uint32(key[tailIndex])
		k *= c1
		k = rotateLeft(k, 15)
		k *= c2
		hash ^= k
	}

	hash ^= uint32(length)
	hash = fmix32(hash)

	return hash
}

func GetLocationIndex(key []byte, numNodes int, seed uint32) int {
	if numNodes <= 0 {
		return 0
	}

	hash := MurmurHash3(key, seed)
	return int(hash % uint32(numNodes))
}

func GetConsistentLocation(key []byte, numNodes int, seed uint32) int {
	if numNodes <= 0 {
		return 0
	}

	hash := MurmurHash3(key, seed)
	return int(hash%uint32(numNodes)) + 1
}

func rotateLeft(x uint32, r uint8) uint32 {
	return (x << r) | (x >> (32 - r))
}

func fmix32(h uint32) uint32 {
	h ^= h >> 16
	h *= 0x85ebca6b
	h ^= h >> 13
	h *= 0xc2b2ae35
	h ^= h >> 16
	return h
}

func GetVirtualNodeIndex(key []byte, numVirtualNodes, numPhysicalNodes int, seed uint32) int {
	if numVirtualNodes <= 0 || numPhysicalNodes <= 0 {
		return 0
	}

	hash := MurmurHash3(key, seed)
	virtualIndex := int(hash % uint32(numVirtualNodes*numPhysicalNodes))
	return (virtualIndex / numVirtualNodes) % numPhysicalNodes
}

func GetWeightedLocation(key []byte, weights []float64, seed uint32) int {
	if len(weights) == 0 {
		return 0
	}

	hash := MurmurHash3(key, seed)
	hashFloat := float64(hash) / float64(math.MaxUint32)

	cumulativeWeight := 0.0
	for i, weight := range weights {
		cumulativeWeight += weight
		if hashFloat <= cumulativeWeight {
			return i
		}
	}

	return len(weights) - 1
}
