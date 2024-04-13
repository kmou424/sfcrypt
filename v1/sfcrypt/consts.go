package sfcrypt

import "runtime"

var MaxThreads = runtime.NumCPU()

// DefaultThreads is the default number of threads to use, default is half of the available CPUs
var DefaultThreads = MaxThreads / 2

// BufferSize is the ratio of the buffer size to the file size, default is 16384 * 64 = 1MB
const BufferSize = 16384 * 64

// BlockRatio is the ratio of the block size to the file size, default is 1 * 16384 * 64 = 1MB
const BlockRatio = 4
