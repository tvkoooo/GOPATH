package callback

import (
	"sync"
)

type UriDecodeHandler func ( a interface{}, length int) ()
type MapUriFuncDecode struct {
	M map[uint32]UriDecodeHandler
	L *sync.RWMutex
}






