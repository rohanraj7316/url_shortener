package helpers

import (
	"hash/crc32"
)

// GetHash returns url hash value
func GetHash(data []byte) uint32 {
	crc32q := crc32.MakeTable(0xD5828281)
	return crc32.Checksum(data, crc32q)
}
