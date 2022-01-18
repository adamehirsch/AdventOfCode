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

func (p Packet) String() string {
	return fmt.Sprintf("{ version: %d, typeId: %d, packetType: %d, value: %d, subpackets: %v }\n", p.version, p.typeId, p.packetType, p.value, p.subpackets)
}

func HexToBin(hex string) string {
	bin := ""
	for _, v := range hex {
		bin += hexMap[string(v)]
	}
	return bin
}

func BinToDec(bin string) int {
	i, err := strconv.ParseInt(bin, 2, 64)
	if err != nil {
		fmt.Println(err)
	}
	return int(i)
}

func createNewPacketFromBinaryPayload(binaryPayload string) (Packet, string) {
	return Packet{
		version:          BinToDec(binaryPayload[0:3]),
		packetType:       BinToDec(binaryPayload[3:6]),
		remainingPayload: binaryPayload[6:],
	}, binaryPayload[6:]
}

// func ParsePacket(input string) (Packet, string) {
// 	newPacket, remainingPayload := createNewPacketFromBinaryPayload(input)

// }

func parseLiteralPacket(packet Packet, payload string) (Packet, string) {
	binString := ""
	keepGoing := true
	for keepGoing {
		binString += payload[1:5]
		keepGoing = payload[0:1] == "1"
		payload = payload[5:]
	}
	p := Packet{
		version:    packet.version,
		typeId:     packet.typeId,
		packetType: packet.packetType,
		value:      BinToDec(binString),
	}
	return p, payload
}

func parseOperatorPacket(packet Packet, payload string) (Packet, string) {
	operatorBitLength := 1
	lengthTypeId := BinToDec(payload[0:operatorBitLength])
	length, remainingPayload := getPacketLength(lengthTypeId, payload[1:])
	subPackets, rp := readSubPackets(remainingPayload, lengthTypeId, length)
	packet.subpackets = append(packet.subpackets, subPackets...)

	return packet, rp
}

func readSubPackets(payload string, lengthTypeId, length int) ([]Packet, string) {

	newPackets := []Packet{}

	if lengthTypeId == 0 {
		// length type is 0; length is length in bits of all subpackets
		packetPayload := payload[0:length]
		for len(packetPayload) > 0 && BinToDec(packetPayload) > 0 {
			newPacket, remainder := createNewPacketFromBinaryPayload(packetPayload)
			parsedPacket, rp := parsePacket(newPacket, remainder)
			packetPayload = rp
			newPackets = append(newPackets, parsedPacket)
		}
	} else {
		//length type 11; length is number of sub packets
		for i := 0; i < length; i++ {
			newPacket, remainder := createNewPacketFromBinaryPayload(payload)
			parsedPacket, rp := parsePacket(newPacket, remainder)
			payload = rp
			newPackets = append(newPackets, parsedPacket)
		}
	}
	return newPackets, payload
}

func getPacketLength(lengthTypeId int, payload string) (int, string) {
	if lengthTypeId == 0 {
		return BinToDec(payload[0:15]), payload[15:]
	} else {
		return BinToDec(payload[0:11]), payload[11:]
	}
}

func parsePacket(packet Packet, payload string) (Packet, string) {
	if packet.packetType == 4 {
		return parseLiteralPacket(packet, payload)
	}
	return parseOperatorPacket(packet, payload)
}

func decodeOuterPacket(binaryPacket string) (Packet, string) {
	outermostPacket, remainingPayload := createNewPacketFromBinaryPayload(binaryPacket)
	return parsePacket(outermostPacket, remainingPayload)
}

func day2116Func(cmd *cobra.Command, args []string) {
	// hexPackets, err := utils.Opener("data/2116-sample.txt", true)
	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }
	// binaryPackets := HexToBin(hexPackets)
	binaryPackets := HexToBin("EE00D40C823060")

	decodedPackets, length := decodeOuterPacket(binaryPackets)
	fmt.Println(decodedPackets, length)
}
