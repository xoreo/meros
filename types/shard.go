package types

import (
	"errors"
	"math"

	"github.com/xoreo/meros/common"
	"github.com/xoreo/meros/core"
	"github.com/xoreo/meros/crypto"
)

var (
	// ErrNilBytes is thrown when a shard is constructed when given nil bytes
	ErrNilBytes = errors.New("bytes to construct new shard must not be nil")

	// ErrCannotCalculateShardSizes is thrown when the []byte to a CalculateShardSizes call is nil
	ErrCannotCalculateShardSizes = errors.New("bytes to calculate shard sizes must not be nil")
)

// Shard is a struct that holds a piece of data that is
// a part of another, bigger piece of data
type Shard struct {
	Size  uint32      `json:"size"`  // The size of the shard
	Bytes []byte      `json:"bytes"` // The actual data of the shard
	Hash  common.Hash `json:"hash"`  // The hash of the shard
}

// NewShard attempts to construct a new shard
func NewShard(bytes []byte) (*Shard, error) {
	if bytes == nil {
		return nil, ErrNilBytes
	}

	newShard := &Shard{
		Size:  uint32(len(bytes)),
		Bytes: bytes,
		Hash:  crypto.Sha3(bytes),
	}

	return newShard, nil
}

// CalculateShardSizes determines the recommended size of each shard
func CalculateShardSizes(raw []byte, n int) ([]uint32, error) {
	rawSize := len(raw)

	// Check that the input is not null
	if rawSize == 0 {
		return nil, ErrCannotCalculateShardSizes
	}

	partition := math.Floor(float64(rawSize / n)) // Calculate the size of each shard
	partitionSize := uint32(partition)            // Convert to a uint32
	modulo := uint32(rawSize % n)                 // Calculate the module mod n

	// Populate a slice of the correct shard sizes
	var sizes []uint32
	for i := 0; i < n; i++ {
		sizes = append(sizes, partitionSize)
	}

	// Adjust for the left over bytes
	if modulo+partitionSize >= partitionSize*uint32(n) {

	}

	sizes[n-1] += modulo // Add the left over bytes to the last element

	return sizes, nil
}

// GenerateShards generates a slice of shards given a string of bytes
func GenerateShards(bytes []byte, n int) ([]*Shard, error) {
	var shards []*Shard // Init the shard slice

	// splitBytes is going to be a 2d array that is returned from the core.split function

	shardSizes, err := CalculateShardSizes(bytes, n) // Calculate the shard sizes
	if err != nil {
		return nil, err
	}

	splitBytes, err := core.SplitBytes(bytes, shardSizes) // Split the bytes into the correct sizes
	if err != nil {
		return nil, err
	}

	// Generate the slices
	for i := 0; i < len(shardSizes); i++ {
		newShard, err := NewShard(
			splitBytes[i],
		)
		if err != nil {
			return nil, err
		}
		shards = append(shards, newShard)
	}

	return nil, nil
}
