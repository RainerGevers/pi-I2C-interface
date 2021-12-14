package pi_I2C_interface

import (
	"encoding/hex"
	"fmt"
	"os/exec"
	"strings"
)

type DataReader struct {
	I2cBus   uint
	ChipAddr uint8
}

func (dr *DataReader) ReadFromAddr(dataAddr uint8) uint8 {
	cmd := exec.Command(fmt.Sprintf("i2cget -y %d 0x%X 0x%X", dr.I2cBus, dr.ChipAddr, dataAddr))
	dataRaw, err := cmd.Output()
	if err != nil {
		return 0
	}
	data := strings.TrimLeft(strings.TrimRight(string(dataRaw), "\n"), "0x")
	res, err := hex.DecodeString(data)
	if err != nil {
		return 0
	}
	return res[0]
}

func (dr *DataReader) ReadFromAddrUInt16(dataAddr uint8) uint16 {
	data := dr.ReadFromAddr(dataAddr)
	return uint16(data)
}

func (dr *DataReader) ReadFromAddrUInt32(dataAddr uint8) uint32 {
	data := dr.ReadFromAddr(dataAddr)
	return uint32(data)
}
