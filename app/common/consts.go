package common

import "runtime"

var MaxRoutines = runtime.NumCPU()

// DefaultRoutines is the default number of goroutines to use, default is half of the available CPUs
var DefaultRoutines = MaxRoutines / 2

// BufferSize is the ratio of the buffer size to the file size, default is 1024 * 1024KB = 1MB
const BufferSize = 1024 * 1024

// BlockRatio is the ratio of the block size to the file size, default is 1 * 1024 * 1024KB = 1MB
const BlockRatio = 1

// SFCryptFileExt is the file extension for sfcrypt files
const SFCryptFileExt = ".sfc"
