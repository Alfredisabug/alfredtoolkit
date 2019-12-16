package features

import (
	"encoding/hex"
	"strconv"
)

/*
PEC1byte 計算2補數 return 1 byte
*/
func PEC1byte(cmd []byte) (PEC []byte) {
	cmdHex := make([]byte, hex.DecodedLen(len(cmd)))
	hex.Decode(cmdHex, cmd)
	sum := 0
	for _, data := range cmdHex {
		sum += int(data)
	}

	//2補數
	turn := (sum ^ 0xFFFF) + 1
	s := strconv.FormatInt(int64(turn), 16)

	cmd = append(cmd, s[len(s)-2], s[len(s)-1])
	// PEC = make([]byte, hex.DecodedLen(len(cmd)))
	// hex.Decode(PEC, cmd)
	return cmd
}

/*
PEC1byte 計算2補數 return 2 bytes
*/
func PEC2byte(cmd []byte) (PEC []byte) {
	cmdHex := make([]byte, hex.DecodedLen(len(cmd)))
	hex.Decode(cmdHex, cmd)
	sum := 0
	for _, data := range cmdHex {
		sum += int(data)
	}

	//2補數
	turn := (sum ^ 0xFFFFFFFF) + 1
	s := strconv.FormatInt(int64(turn), 16)

	cmd = append(cmd, s[len(s)-4], s[len(s)-3], s[len(s)-2], s[len(s)-1])
	// PEC = make([]byte, hex.DecodedLen(len(cmd)))
	// hex.Decode(PEC, cmd)
	return cmd
}
