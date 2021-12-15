package pi_I2C_interface

import (
	"encoding/hex"
	"fmt"
	"os/exec"
	"strings"
)

type DataReaderWriter struct {
	I2cBus   uint
	ChipAddr uint8
}

func (dr *DataReaderWriter) ReadFromAddr(dataAddr uint8) uint8 {
	command := "/usr/sbin/i2cget"
	arg0 := "-y"
	arg1 := fmt.Sprintf("%d", dr.I2cBus)
	arg2 := fmt.Sprintf("0x%X", dr.ChipAddr)
	arg3 := fmt.Sprintf("0x%X", dataAddr)
	cmd := exec.Command(command, arg0, arg1, arg2, arg3)
	dataRaw, err := cmd.Output()
	data := strings.TrimLeft(strings.TrimRight(string(dataRaw), "\n"), "0x")
	if len(data) == 1 {
		data = fmt.Sprintf("0%v", data)
	}
	res, err := hex.DecodeString(data)
	if err != nil || len(res) == 0 {
		return 0
	}
	return res[0]
}

func (dr *DataReaderWriter) ReadFromAddrUInt16(dataAddr uint8) uint16 {
	data := dr.ReadFromAddr(dataAddr)
	return uint16(data)
}

func (dr *DataReaderWriter) ReadFromAddrUInt32(dataAddr uint8) uint32 {
	data := dr.ReadFromAddr(dataAddr)
	return uint32(data)
}

func (dr *DataReaderWriter) WriteToAddr(dataAddr uint8, data uint8) bool {
	command := "/usr/sbin/i2cset"
	arg0 := "-y"
	arg1 := fmt.Sprintf("%d", dr.I2cBus)
	arg2 := fmt.Sprintf("0x%X", dr.ChipAddr)
	arg3 := fmt.Sprintf("0x%X", dataAddr)
	arg4 := fmt.Sprintf("0x%X", data)
	cmd := exec.Command(command, arg0, arg1, arg2, arg3, arg4)
	_, err := cmd.Output()
	if err != nil {
		return false
	}
	return true
}
