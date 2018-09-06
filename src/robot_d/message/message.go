package message

import (
	"sync"
	"robot_d/common/datagroove"
)

type UriDecodeHandler func (d *datagroove.DataBuff, i interface{}, length int) ()
type MapUriFuncDecode struct {
	M map[uint32]UriDecodeHandler
	L *sync.RWMutex
}

func (p *MapUriFuncDecode)ZhuCe(uri uint32 , i UriDecodeHandler )  {
	p.L.Lock()
	p.M[uri] = i
	p.L.Unlock()
}


func (p *MapUriFuncDecode)UriDecodeHandlerInit(){
	p.M = make(map[uint32]UriDecodeHandler )
	p.L = new(sync.RWMutex)

}