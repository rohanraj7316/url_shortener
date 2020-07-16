package helpers

import (
	"hash/crc32"
)

// GetHash returns url hash value. need to build all
// logic here regarding hash.
// func GetHash(data []byte) uint32 {
func GetHash() uint32 {
	crc32q := crc32.MakeTable(0xD5828281)
	return crc32.Checksum([]byte("pathname"), crc32q)
}
