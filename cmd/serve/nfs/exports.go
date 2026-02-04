//go:build unix

// Fork-specific exports for external use.
// This file contains additions to make internal functions accessible.
package nfs

// DefaultOpt is the default values used for Options
var DefaultOpt = Options{
	HandleLimit: 1000000,
	HandleCache: cacheMemory,
}
