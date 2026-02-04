// Fork-specific exports for external use.
// This file contains additions to make internal functions accessible.
package mountlib

// GetMountFn returns the mount function for the given mount type.
// Mount types: "mount" (FUSE), "cmount" (cgofuse), "mount2" (FUSE2)
func GetMountFn(mountType string) MountFn {
	mountMu.Lock()
	defer mountMu.Unlock()
	return mountFns[mountType]
}
