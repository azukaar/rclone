//go:build !plan9

// Fork-specific exports for external use.
// This file contains additions to make internal functions accessible.
package sftp

import (
	"context"

	"github.com/rclone/rclone/cmd/serve/proxy"
	"github.com/rclone/rclone/fs"
	"github.com/rclone/rclone/vfs/vfscommon"
)

// DefaultOpt is the default values used for Options
var DefaultOpt = Options{
	ListenAddr:     "localhost:2022",
	AuthorizedKeys: "~/.ssh/authorized_keys",
}

// Server is the external interface to the SFTP server
type Server struct {
	*server
}

// NewServer creates a new SFTP server
func NewServer(ctx context.Context, f fs.Fs, opt *Options, vfsOpt *vfscommon.Options, proxyOpt *proxy.Options) (*Server, error) {
	s, err := newServer(ctx, f, opt, vfsOpt, proxyOpt)
	if err != nil {
		return nil, err
	}
	return &Server{server: s}, nil
}
