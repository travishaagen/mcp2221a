package main

import (
	"log"
	"time"

	mcp "github.com/ardnew/mcp2221a"
)

func main() {
	m, err := mcp.New(0, mcp.VID, mcp.PID)
	if nil != err {
		log.Fatalf("Open(): %v", err)
	}
	defer m.Close()

	if err := m.Reset(5 * time.Second); nil != err {
		log.Fatalf("Reset(): %v", err)
	}

	if err := m.I2CSetConfig(mcp.I2CBaudRate); nil != err {
		log.Fatalf("I2CSetConfig(): %v", err)
	}

	if addr, err := m.I2CScan(mcp.I2CMinAddr, mcp.I2CMaxAddr); nil != err {
		log.Fatalf("I2CScan(): %v", err)
	} else {
		for _, a := range addr {
			log.Printf("scan found = 0x%02X", a)
		}
	}

	// read the 16-bit data from device ID register (0xFF) from INA260 with
	// default slave address (0x40)
	if buf, err := m.I2CReadReg(0x40, 0xFF, 2); nil != err {
		log.Fatalf("I2CReadReg(): %v", err)
	} else {

		var ub uint16 = (uint16(buf[0]) << 8) | uint16(buf[1])

		rev := ub & 0x0F // revision is 4 bits (LSB)
		die := ub >> 4   // device ID is remaining 12 bits

		log.Printf("Revision  = %3d {0x%4X} [0b%16b]", rev, rev, rev)
		log.Printf("Device ID = %3d {0x%4X} [0b%16b]", die, die, die)
	}
}
