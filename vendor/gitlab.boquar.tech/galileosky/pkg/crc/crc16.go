package crc

type CRC16Func func([]byte) uint16

var DefaultCRC16 CRC16Func = Calc
