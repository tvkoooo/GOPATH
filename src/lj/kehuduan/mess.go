package kehuduan

import (
	"encoding/binary"
)

type messagehead struct {
	pid   uint32
	mid   uint32
	uid   uint64
	lenth uint32
	linkd uint32
	pxy   uint32
}

func set_pid(messh messagehead, pid uint32) {
	messh.pid = pid
}

func set_mid(messh messagehead, mid uint32) {
	messh.mid = mid
}

func set_uid(messh messagehead, uid uint64) {
	messh.uid = uid
}

func set_linkd(messh messagehead, linkd uint32) {
	messh.linkd = linkd
}

func set_pxy(messh messagehead, pxy uint32) {
	messh.pxy = pxy
}

func Messbuff(messh messagehead, jsonbuff string) []byte {
	var lenth_json uint32
	lenth_json = uint32(len([]rune(jsonbuff)))
	var lenth uint32
	lenth = 28 + lenth_json
	pid := messh.pid
	mid := messh.mid
	uid := messh.uid
	var messbuf []byte
	binary.BigEndian.Uint32(messbuf[0:4]) = pid
	binary.BigEndian.Uint32(messbuf[5:8]) = mid
	binary.BigEndian.Uint64(messbuf[9:16]) = uid
	binary.BigEndian.Uint32(messbuf[17:20]) = lenth
	binary.BigEndian.Uint32(messbuf[21:24]) = 0
	binary.BigEndian.Uint32(messbuf[25:28]) = 0
	copy(messbuf[29:29+lenth_json] ,jsonbuff[0:lenth_json])
	return messbuf
}
