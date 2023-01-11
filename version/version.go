package version

import (
	"fmt"
)

const (
	version = "0.1.9"
	name    = "DoDoOpenGo"
)

// Version Get current version string
func Version() string {
	return fmt.Sprintf("%s/%s", name, version)
}
