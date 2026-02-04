// Fork-specific exports for external use.
// This file contains additions to make internal functions accessible.
package webdav

import (
	"context"

	"github.com/rclone/rclone/cmd/serve/proxy"
	"github.com/rclone/rclone/fs"
	"github.com/rclone/rclone/vfs/vfscommon"
)

// DefaultOpt is the default values used for Options
var DefaultOpt = Options{}

// NewServer creates a new WebDAV server
func NewServer(ctx context.Context, f fs.Fs, opt *Options, vfsOpt *vfscommon.Options, proxyOpt *proxy.Options) (*WebDAV, error) {
	return newWebDAV(ctx, f, opt, vfsOpt, proxyOpt)
}
