package ddcci

import (
	"encoding/binary"
	"fmt"
	"strings"
)

const (
	ReportSize     = 193 // 1 report ID + 192 data bytes
	ddcOffset      = 0x40
	writeCmd       = 0x03 // Set VCP
	readCmd        = 0x01 // Get VCP
	respFuncCode   = 0x02 // Get VCP reply
	respOK         = 0x00
	monitorAddress = 0x6e // sink address in replies
	hostAddress    = 0x51 // source address in host frames
)

func BuildSetFrame(vcp uint16, value uint16, frameLenBase int) []byte {
	buf := make([]byte, ReportSize)
	copy(buf[1:], []byte{0x40, 0xc6})
	copy(buf[1+6:], []byte{0x20, 0, 0x6e, 0, 0x80})

	payload := make([]byte, 0, 4)
	if vcp > 0xff {
		payload = binary.BigEndian.AppendUint16(payload, vcp)
	} else {
		payload = append(payload, byte(vcp))
	}
	payload = binary.BigEndian.AppendUint16(payload, value)

	frame := append([]byte{hostAddress, byte(frameLenBase + len(payload)), writeCmd}, payload...)
	copy(buf[1+ddcOffset:], frame)
	return buf
}

func BuildGetFrame(vcp uint16, frameLenBase int) []byte {
	buf := make([]byte, ReportSize)
	copy(buf[1:], []byte{0x40, 0xc6})
	copy(buf[1+6:], []byte{0x20, 0, 0x6e, 0, 0x80})

	payload := make([]byte, 0, 2)
	if vcp > 0xff {
		payload = binary.BigEndian.AppendUint16(payload, vcp)
	} else {
		payload = append(payload, byte(vcp))
	}
	frame := append([]byte{hostAddress, byte(frameLenBase + len(payload)), readCmd}, payload...)
	copy(buf[1+ddcOffset:], frame)
	return buf
}

// Parse Get VCP reply: 0x6e <len> 0x02 0x00 <opcode> <type> <max_hi><max_lo> <cur_hi><cur_lo> <cksum>
func ParseGetReply(report []byte) (cur, max uint16, vcpType byte, err error) {
	if len(report) < 1+ddcOffset+10 {
		return 0, 0, 0, fmt.Errorf("short HID report (%d bytes)", len(report))
	}
	resp := report[1+ddcOffset:]
	i := 0
	for i < len(resp) && resp[i] != monitorAddress {
		i++
	}
	if i+10 >= len(resp) {
		return 0, 0, 0, fmt.Errorf("DDC/CI response not found")
	}
	r := resp[i:]

	if r[0] != monitorAddress || r[2] != respFuncCode {
		return 0, 0, 0, fmt.Errorf("unexpected reply header: % x", r[:min(8, len(r))])
	}
	if r[3] != respOK {
		return 0, 0, 0, fmt.Errorf("monitor returned error 0x%02x", r[3])
	}
	vcpType = r[5]
	max = binary.BigEndian.Uint16(r[6:8])
	cur = binary.BigEndian.Uint16(r[8:10])
	return cur, max, vcpType, nil
}

func WouldBlock(err error) bool {
	if err == nil {
		return false
	}
	s := strings.ToLower(err.Error())
	return strings.Contains(s, "would block") || strings.Contains(s, "timeout")
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
