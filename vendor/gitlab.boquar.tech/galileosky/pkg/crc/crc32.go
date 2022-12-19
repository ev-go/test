package crc

const (
	POLYNOM32_ZIP         uint32 = 0xEDB88320 //(ZIP)
	POLYNOM32_MPEG2       uint32 = 0x04C11DB7 //(MPEG2)
	POLYNOM32_MPEG2_1     uint32 = POLYNOM32_ZIP
	POLYNOM32_MPEG2_2     uint32 = 0xDB710641
	POLYNOM32_IEEE        uint32 = POLYNOM32_MPEG2 //(IEEE 802.3)
	POLYNOM32_IEEE_1      uint32 = POLYNOM32_ZIP
	POLYNOM32_IEEE_2      uint32 = POLYNOM32_MPEG2_2
	POLYNOM32_CASTAGNOLI1 uint32 = 0x1EDC6F41 //(Castagnoli)
	POLYNOM32_CASTAGNOLI2 uint32 = 0x82F63B78
	POLYNOM32_Castagnoli3 uint32 = 0x05EC76F1
	POLYNOM32_KOOPMAN1    uint32 = 0x741B8CD7 //(Koopman)
	POLYNOM32_KOOPMAN2    uint32 = 0xEB31D82E
)

func CalcCRC32(dataArray []uint8, polynom uint32) uint32 {
	if polynom == 0 {
		polynom = POLYNOM32_ZIP
	}

	return CalcCRC32Seaded(0xFFFFFFFF, dataArray, polynom)
}

func CalcCRC32Seaded(seed uint32, dataArray []uint8, polynom uint32) uint32 {
	if polynom == 0 {
		polynom = POLYNOM32_ZIP
	}

	for _, data := range dataArray {
		seed ^= uint32(data)
		for n := 0; n < 8; n++ {
			seed = (-(seed & 1) & polynom) ^ (seed >> 1)
		}
	}

	return seed
}
