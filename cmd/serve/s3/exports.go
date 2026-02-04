// Fork-specific exports for external use.
// This file contains additions to make internal functions accessible.
package s3

// DefaultOpt is the default values used for Options
var DefaultOpt = Options{
	ForcePathStyle: true,
	EtagHash:       "MD5",
}
