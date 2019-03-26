package random

import "crypto/rand"

// MaxGeneratorCacheSize is the maximum number of cached entries.
const MaxGeneratorCacheSize = 4096

// SequenceCacher caches and returns previous sequences.
type SequenceCacher interface {
	IsCaching() bool
	CacheSize() int
	GetCachedSequence(n int) []byte
	GetCache() []byte
	Flush()
}

// Generator implements a sequence generator.
type Generator struct {
	cacheSize int
	cache     [][]byte
}

// NewGenerator creates a new Generator
// The generator can optionally cache up to maxCache generated entries.
func NewGenerator(isCaching bool, maxCache int) (result *Generator) {
	result = new(Generator)
	if isCaching {
		if maxCache > MaxGeneratorCacheSize {
			panic("Cache size exceeds max")
		}
		result.cacheSize = maxCache
		result.cache = make([][]byte, maxCache)
	}
	return result
}

func (g *Generator) cacheSequence(seq []byte) {
	g.cache = append(g.cache, seq)
	if len(g.cache) > g.cacheSize {
		g.cache = g.cache[1:] // TODO: test this
	}
}

// NewSequence implements Sequencer. It returns a random sequence.
// This size of the sequence MUST BE RandomSliceByteSize.
// Panics if the sequence of exact size could not be generated.
// This implementation uses 'crypto/rand'.
func (g *Generator) NewSequence() (result []byte) {
	result = make([]byte, RandomSliceByteSize)
	n, err := rand.Read(result)
	if err != nil || n != RandomSliceByteSize {
		panic("Failed to generate random sequence")
	}
	if g.cacheSize != 0 {
		g.cacheSequence(result)
	}
	return result
}

// IsCaching returns if the generator caches entries
func (g *Generator) IsCaching() bool {
	return g.cacheSize != 0
}

// CacheSize implements SequenceCacher. It returns the cache size.
func (g *Generator) CacheSize() int { return g.cacheSize }

// GetCachedSequence implements SequenceCacher. It returns a cached sequence.
// Returns ErrOutOfRange if index is out of cache range.
func (g *Generator) GetCachedSequence(idx int) (result []byte, err error) {
	if idx >= g.cacheSize {
		return nil, ErrOutOfRange
	}
	return g.cache[idx], nil
}

// GetCache implements SequenceCacher. It returns the sequence cache.
func (g *Generator) GetCache() [][]byte { return g.cache }

// Flush implements SequenceCacher. It flushes the sequence cache.
func (g *Generator) Flush() {
	g.cache = make([][]byte, g.cacheSize)
}