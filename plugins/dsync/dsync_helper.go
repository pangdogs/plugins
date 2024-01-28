package dsync

import (
	"git.golaxy.org/core/service"
	"git.golaxy.org/core/util/option"
	"strings"
)

// NewMutex returns a new distributed mutex with given name.
func NewMutex(servCtx service.Context, name string, settings ...option.Setting[DistMutexOptions]) IDistMutex {
	return Using(servCtx).NewMutex(name, settings...)
}

// GetSeparator return name path separator.
func GetSeparator(servCtx service.Context) string {
	return Using(servCtx).GetSeparator()
}

// Path return name path.
func Path(servCtx service.Context, elems ...string) string {
	return strings.Join(elems, GetSeparator(servCtx))
}