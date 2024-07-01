package evmStructs

import (
	"errors"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"solity/utils/logger"
)

type Chunker struct {
	chunkSize   int
	storage     [][]byte
	isPadded    bool
	paddingSize int
}

/*
NewChunker constructor for the Chunker object, takes in arbitrary sized []byte and chunkSize int
parses the inputted byte array into fixed sized chunks based on the supplied chunks. Applies 0 padding
to the end if needed
*/
func NewChunker(input []byte, chunkSize int) (Chunker, error) {
	// Init return variable
	ret := Chunker{}
	ret.isPadded = false
	ret.paddingSize = 0

	// Chunk Size Check
	if chunkSize <= 0 {
		return ret, errors.New("Invalid chunk size (<= 0)")
	}

	// Set chunk size
	ret.chunkSize = chunkSize

	// Check input length
	if len(input) == 0 {
		return ret, errors.New("Empty data")
	}

	// Determine if a padding is needed
	if len(input)%chunkSize != 0 {

		// Determine the needed padding size
		neededLength := chunkSize - len(input)%chunkSize

		padding := []byte{}

		for i := 0; i < neededLength; i++ {
			padding = append(padding, byte(0))
		}

		// Add the padding to the end
		input = append(input, padding...)

		// Set the padding flag and the padding size
		ret.isPadded = true
		ret.paddingSize = neededLength
	}

	// Create chunks
	for i := 0; i+chunkSize <= len(input); i = i + chunkSize {
		ret.storage = append(ret.storage, input[i:i+chunkSize])
	}

	return ret, nil
}

/*
GetNumberOfChunks returns the total number of chunks
*/
func (c *Chunker) GetNumberOfChunks() int {
	return len(c.storage)
}

/*
GetChunkSize returns the chunk size of the chunks (data length of each chunk)
*/
func (c *Chunker) GetChunkSize() int {
	return c.chunkSize
}

/*
GetChunk returns empty byte array if the supplied index is out of range, else returns that chunk
*/
func (c *Chunker) GetChunk(index int) []byte {

	if index >= len(c.storage) || index < 0 {
		return []byte{}
	}

	return c.storage[index]
}

/*
GetChunkSafe returns empty byte array and error if the supplied index is out of range, else returns that chunk and nil
*/
func (c *Chunker) GetChunkSafe(index int) ([]byte, error) {

	if index >= len(c.storage) || index < 0 {
		return []byte{}, errors.New("Index out of range")
	}

	return c.storage[index], nil
}

/*
ByteIndexToChunk takes index of a byte and returns the chunk index and new index of that byte inside the given chunk,
returns -1, -1 if the index is out of range,
*/
func (c *Chunker) ByteIndexToChunk(index int) (chunkIndex int, newIndex int) {

	if (index / c.chunkSize) >= len(c.storage) {
		return -1, -1
	}

	return index / c.chunkSize, index % c.chunkSize
}

/*
HasPadding returns true if padding has been applied to the byte array prior to the chunking
*/
func (c *Chunker) HasPadding() bool {
	return c.isPadded
}

/*
PaddingSize returns the applied padding size
*/
func (c *Chunker) PaddingSize() int {
	return c.paddingSize
}

/*
GetByteArrayByByteIndex returns byte array with the size length, starting from the starting byte index, if index is out of range
then returns empty byte array. Also return the index of the chunk that the last byte is the member of
*/
func (c *Chunker) GetByteArrayByByteIndex(startingByteIndex int, length int) ([]byte, int) {
	// Init return
	ret := []byte{}

	// Determine the starting chunk corresponding chunk
	chunkStartIndex, newByteStartIndex := c.ByteIndexToChunk(startingByteIndex)

	// If the specified starting byte index is out of range return empty
	if chunkStartIndex == -1 {
		return ret, -1
	}

	// Determine the ending chunk
	chunkEndIndex, newByteEndIndex := c.ByteIndexToChunk(startingByteIndex + length)

	// If the specified starting byte index is out of range return empty
	if chunkEndIndex == -1 {
		return ret, -1
	}

	// If starting index is the same as the ending index
	if chunkStartIndex == chunkEndIndex {
		ret = append(ret, c.GetChunk(chunkStartIndex)[newByteStartIndex:newByteEndIndex]...)
		return ret, chunkStartIndex
	}

	for i := chunkStartIndex; i < chunkEndIndex; i++ {
		// Append the whole chunk
		ret = append(ret, c.GetChunk(i)...)
	}

	ret = append(ret, c.GetChunk(chunkEndIndex)[:newByteEndIndex]...)

	return ret, chunkEndIndex
}

/*
GetIndexOfStartingByte returns the index of the first byte in the given chunk
*/
func (c *Chunker) GetIndexOfStartingByte(startingChunkIndex int) int {
	return startingChunkIndex * c.chunkSize
}

/*
PrintChunkedArrays prints all the chunks one chunk per line
*/
func (c *Chunker) PrintChunkedArrays() {
	for i := 0; i < len(c.storage); i++ {
		logger.LogDf("[%v] %v", i, c.storage[i])
	}
}

/*
PrintChunkedArraysAsHex prints all the chunks one chunk per line
*/
func (c *Chunker) PrintChunkedArraysAsHex() {
	for i := 0; i < len(c.storage); i++ {
		logger.LogDf("[%v] %s", i, hexutil.Encode(c.storage[i])[2:])
	}
}

/*
GetSlice returns the slice of data out of the whole data starting from the start index until the end index
if end index is negative, it is set to last chunk index. Returned data [startIndex, endIndex)
*/
func (c *Chunker) GetSlice(startingChunkIndex int, endChunkIndex int) (theSlice []byte, err error) {
	// Start check
	if startingChunkIndex > len(c.storage) || startingChunkIndex < 0 {
		err = errors.New("starting chunk index is out of range")
		return
	}

	// Check if the end chunk is negative
	if endChunkIndex < 0 {
		endChunkIndex = len(c.storage)
	}

	// Check if the end index is feasible
	if endChunkIndex > len(c.storage) {
		err = errors.New("end chunk index is out of range")
		return
	}

	theSlice = []byte{}

	for i := startingChunkIndex; i < endChunkIndex; i++ {
		theSlice = append(theSlice, c.storage[i]...)
	}

	return
}
