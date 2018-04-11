package messagehead

import (
	"encoding/binary"
)

type Messagehead struct {
	Pid   uint32
	Mid   uint32
	Uid   uint64
	Lenth uint32
	Linkd uint32
	Pxy   uint32
}

func Encodemessh(messh Messagehead, messhbuff []byte, messhlen uint32) ([]byte, uint32) {
	enlenth := uint32(28 + messhlen)
	pid := make([]byte, 4)
	mid := make([]byte, 4)
	uid := make([]byte, 8)
	lenth := make([]byte, 4)
	linkd := make([]byte, 4)
	pxy := make([]byte, 4)
	outbuf := make([]byte, enlenth)
	binary.LittleEndian.PutUint32(pid, messh.Pid)
	binary.LittleEndian.PutUint32(mid, messh.Mid)
	binary.LittleEndian.PutUint64(uid, messh.Uid)
	binary.LittleEndian.PutUint32(lenth, enlenth)
	binary.LittleEndian.PutUint32(linkd, messh.Linkd)
	binary.LittleEndian.PutUint32(pxy, messh.Pxy)
	copy(outbuf[0:3], pid)
	copy(outbuf[4:7], mid)
	copy(outbuf[8:15], uid)
	copy(outbuf[16:19], lenth)
	copy(outbuf[20:24], linkd)
	copy(outbuf[25:27], pxy)
	copy(outbuf[28:28+messhlen], messhbuff[0:messhlen])
	return outbuf, enlenth
}
func Decodemessh(debuff []byte, delen uint32) (Messagehead, []byte, uint32) {
	messhlen := uint32(delen - 28)
	pid := make([]byte, 4)
	mid := make([]byte, 4)
	uid := make([]byte, 8)
	lenth := make([]byte, 4)
	linkd := make([]byte, 4)
	pxy := make([]byte, 4)
	outbuf := make([]byte, delen)
	copy(pid, debuff[0:3])
	copy(mid, debuff[4:7])
	copy(uid, debuff[8:15])
	copy(lenth, debuff[16:19])
	copy(linkd, debuff[20:23])
	copy(pxy, debuff[24:27])
	copy(outbuf[0:messhlen], debuff[28:28+messhlen])
	var messh Messagehead
	messh.Pid = binary.LittleEndian.Uint32(pid)
	messh.Mid = binary.LittleEndian.Uint32(mid)
	messh.Uid = binary.LittleEndian.Uint64(uid)
	messh.Lenth = binary.LittleEndian.Uint32(lenth)
	messh.Linkd = binary.LittleEndian.Uint32(linkd)
	messh.Pxy = binary.LittleEndian.Uint32(pxy)
	return messh, outbuf, messhlen
}

func Get_pid(debuff []byte)uint32{
	pid := make([]byte, 4)
	copy(pid, debuff[0:3])
	outpid := binary.LittleEndian.Uint32(pid)
	return outpid
}
func Get_mid(debuff []byte)uint32{
	mid := make([]byte, 4)
	copy(mid, debuff[4:7])
	outmid := binary.LittleEndian.Uint32(mid)
	return outmid
}
func Get_uid(debuff []byte)uint64{
	uid := make([]byte, 8)
	copy(uid, debuff[8:15])
	outuid := binary.LittleEndian.Uint64(uid)
	return outuid
}
func Get_lenth(debuff []byte)uint32{
	lenth := make([]byte, 4)
	copy(lenth, debuff[16:19])
	outlenth := binary.LittleEndian.Uint32(lenth)
	return outlenth
}
func Get_linkd(debuff []byte)uint32{
	linkd := make([]byte, 4)
	copy(linkd, debuff[20:23])
	outlinkd := binary.LittleEndian.Uint32(linkd)
	return outlinkd
}
func Get_pxy(debuff []byte)uint32{
	pxy := make([]byte, 4)
	copy(pxy, debuff[24:27])
	outpxy := binary.LittleEndian.Uint32(pxy)
	return outpxy
}
func Get_data(debuff []byte, delen uint32)([]byte, uint32){
	messhlen := uint32(delen - 28)
	outbuf := make([]byte, delen)
	copy(outbuf[0:messhlen], debuff[28:28+messhlen])
	return outbuf, messhlen
}
func Change_pid(debuff []byte,pid uint32){
	pid_b := make([]byte, 4)
	binary.LittleEndian.PutUint32(pid_b, pid)
	copy(debuff[0:3],pid_b)
}
func Change_mid(debuff []byte,mid uint32){
	mid_b := make([]byte, 4)
	binary.LittleEndian.PutUint32(mid_b, mid)
	copy(debuff[4:7],mid_b)
}
func Change_uid(debuff []byte,uid uint64){
	uid_b := make([]byte, 8)
	binary.LittleEndian.PutUint64(uid_b, uid)
	copy(debuff[8:15],uid_b)
}
func change_lenth(debuff []byte,lenth uint32){
	lenth_b := make([]byte, 4)
	binary.LittleEndian.PutUint32(lenth_b, lenth)
	copy(debuff[16:19],lenth_b)
}
func Change_data(debuff []byte,databuff []byte, datalen uint32)([]byte,uint32){
	length := uint32 (28 + datalen)
	change_lenth(databuff,datalen)
	change_data := make([]byte, length)
	copy(change_data[:27],debuff[:27])
	copy(debuff[28:28+datalen],databuff[0:datalen])
	return change_data,length
}
func Change_linkd(debuff []byte,linkd uint32){
	linkd_b := make([]byte, 4)
	binary.LittleEndian.PutUint32(linkd_b, linkd)
	copy(debuff[20:23],linkd_b)
}
func Change_pxy(debuff []byte,pxy uint32){
	pxy_b := make([]byte, 4)
	binary.LittleEndian.PutUint32(pxy_b, pxy)
	copy(debuff[24:27],pxy_b)
}
