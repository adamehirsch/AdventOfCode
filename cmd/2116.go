package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var day2116Cmd = &cobra.Command{
	Use:                   "day2116",
	Short:                 "2021 Advent of Code Day 16",
	DisableFlagsInUseLine: true,
	Run:                   day2116Func,
}

func init() {
	rootCmd.AddCommand(day2116Cmd)
}

var hexMap = map[string]string{
	"0": "0000",
	"1": "0001",
	"2": "0010",
	"3": "0011",
	"4": "0100",
	"5": "0101",
	"6": "0110",
	"7": "0111",
	"8": "1000",
	"9": "1001",
	"A": "1010",
	"B": "1011",
	"C": "1100",
	"D": "1101",
	"E": "1110",
	"F": "1111",
}

type Packet struct {
	version          int
	typeId           int
	packetType       int
	value            int
	remainingPayload string
	subpackets       []Packet
}

func HexToBin(hex string) string {
	bin := ""
	for _, v := range hex {
		bin += hexMap[string(v)]
	}
	return bin
}

func BinToDec(bin string) int {
	i := int64(0)
	i, err := strconv.ParseInt(bin, 2, 64)
	if err != nil {
		fmt.Println(err)
	}
	return int(i)
}

func createNewPacketFromBinaryPayload(binaryPayload string) Packet {
	return Packet{
		version:          BinToDec(binaryPayload[0:3]),
		packetType:       BinToDec(binaryPayload[3:6]),
		remainingPayload: binaryPayload[6:],
	}
}

func parseLiteralPacket(packet Packet) (Packet, int) {
	str := packet.remainingPayload
	binString := ""
	keepGoing := "1"
	count := 0
	for keepGoing != "0" {
		binString += str[1:5]
		keepGoing = str[0:1]
		if keepGoing != "0" {
			str = str[5:]
		}
		count++
	}
	return Packet{
		version:    packet.version,
		typeId:     packet.typeId,
		packetType: packet.packetType,
		value:      BinToDec(binString),
	}, 6 + (count * 5)
}

func parseOperatorPacket(packet Packet) (Packet, int) {
	operatorBitLength := 1
	lengthTypeId := BinToDec(packet.remainingPayload[0:operatorBitLength])
	length, remainingPayload := getPacketLength(lengthTypeId, packet.remainingPayload)

	subPackets, length := readSubPackets(remainingPayload, lengthTypeId, length)
	opPacket := Packet{
		version:    packet.version,
		typeId:     packet.typeId,
		packetType: packet.packetType,
		subpackets: subPackets,
	}
	return opPacket, length + 7
}

func readSubPackets(remainingPayload string, lengthTypeId int, length int) ([]Packet, int) {
	totalBitsUsed := 0
	packets := []Packet{}
	if lengthTypeId == 0 {
		// length type is 0; length is length in bits of all subpackets
		remainingPayload := remainingPayload[0:length]
		for len(remainingPayload) > 0 {
			newPacket := createNewPacketFromBinaryPayload(remainingPayload)
			parsedPacket, bitsUsed := parsePacket(newPacket)
			totalBitsUsed += bitsUsed
			remainingPayload = remainingPayload[bitsUsed:]
			packets = append(packets, parsedPacket)
		}
	} else {
		//length type 11; length is number of sub packets
		for i := 0; i < length; i++ {
			newPacket := createNewPacketFromBinaryPayload(remainingPayload)
			parsedPacket, bitsUsed := parsePacket(newPacket)
			totalBitsUsed += bitsUsed
			remainingPayload = remainingPayload[bitsUsed:]
			packets = append(packets, parsedPacket)
		}
	}
	return packets, totalBitsUsed
}

func getPacketLength(lengthTypeId int, payload string) (int, string) {
	if lengthTypeId == 0 {
		return BinToDec(payload[1:16]), payload[16:]
	} else {
		return BinToDec(payload[1:11]), payload[11:]
	}
}

func parsePacket(packet Packet) (Packet, int) {
	if packet.packetType == 4 {
		return parseLiteralPacket(packet)
	} else {
		return parseOperatorPacket(packet)
	}
}

func decodeOuterPacket(binaryPacket string) (Packet, int) {
	outermostPacket := createNewPacketFromBinaryPayload(binaryPacket)
	return parsePacket(outermostPacket)

}

func day2116Func(cmd *cobra.Command, args []string) {
	// hexPackets, err := utils.Opener("data/2116-sample.txt", true)
	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }
	// binaryPackets := HexToBin(hexPackets)
	binaryPackets := HexToBin("38006F45291200")
	// fmt.Println(binaryPackets)

	decodedPackets, length := decodeOuterPacket(binaryPackets)
	// decodedPackets = decodeOuterPacket("110100101111111000101000")
	fmt.Println(decodedPackets, length)
}
