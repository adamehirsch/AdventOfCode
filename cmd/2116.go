package cmd

import (
	"fmt"
	"os"
	"regexp"
	"strconv"

	"github.com/adamehirsch/AdventOfCode/utils"
	"github.com/kr/pretty"
	"github.com/spf13/cobra"
)

var day2116Cmd = &cobra.Command{
	Use:                   "day2116",
	Short:                 "2021 Advent of Code Day 16",
	DisableFlagsInUseLine: true,
	Run:                   day2116Func,
}

var nonZero, _ = regexp.Compile("[^0]")

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
	version    int
	typeId     int
	packetType int
	value      int
	subpackets []Packet
}

func (p Packet) String() string {
	packetType := "literal"
	if p.packetType != 4 {
		packetType = "operator"
	}
	return fmt.Sprintf("{ version: %d, typeId: %d, packetType: %s, value: %d, subpackets: %# v }", p.version, p.typeId, packetType, p.value, pretty.Formatter(p.subpackets))
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

func MakePacket(payload string) (Packet, string) {
	return Packet{
		version:    BinToDec(payload[0:3]),
		packetType: BinToDec(payload[3:6]),
	}, payload[6:]
}

func parseLiteralPacket(packet Packet, payload string) (Packet, string) {
	// a Packet with only numbers in it, thus a concrete value
	binString := ""
	keepGoing := true
	for keepGoing {
		binString += payload[1:5]
		keepGoing = payload[0:1] == "1"
		// trim off the 5 bits we used, leave the rest
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
	// a Packet that indicates an operation to conduct on the numbers in its payload
	lengthTypeId := BinToDec(payload[0:1])
	length, remainingPayload := getPacketLength(lengthTypeId, payload[1:])

	subPackets, rp := readSubPackets(remainingPayload, lengthTypeId, length)
	packet.subpackets = append(packet.subpackets, subPackets...)
	return packet, rp
}

func readSubPackets(payload string, lengthTypeId, length int) ([]Packet, string) {

	newPackets := []Packet{}
	if lengthTypeId == 0 {
		// length type is 0; length is length in bits of all subpackets
		trimmedOff := payload[length:]
		payload = payload[0:length]
		for len(payload) > 0 && nonZero.MatchString(payload) {
			newPacket, remainder := MakePacket(payload)
			parsedPacket, rp := parsePacket(newPacket, remainder)
			payload = rp
			newPackets = append(newPackets, parsedPacket)
		}
		// put the the bits not used in this packet back onto the returned, unused payload
		payload = payload + trimmedOff
	} else {
		//length type 1; length is number of sub packets
		for i := 0; i < length && nonZero.MatchString(payload); i++ {
			newPacket, remainder := MakePacket(payload)
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

	op, payload := parseOperatorPacket(packet, payload)
	op.value = DoPacketMath(op)
	return op, payload
}

func DoPacketMath(op Packet) int {
	if op.subpackets == nil || len(op.subpackets) == 0 {
		return 0
	}
	value := op.subpackets[0].value
	switch op.packetType {
	case 0:
		// "sum"
		for _, p := range op.subpackets[1:] {
			value += p.value
		}
	case 1:
		// "product"
		if len(op.subpackets) > 1 {
			for _, sp := range op.subpackets[1:] {
				value = value * sp.value
			}
		}
	case 2:
		// "minimum"
		for _, sp := range op.subpackets[1:] {
			if sp.value < value {
				value = sp.value
			}
		}
	case 3:
		// maximum
		for _, sp := range op.subpackets[1:] {
			if sp.value > value {
				value = sp.value
			}
		}
	case 5:
		// greater than
		if value > op.subpackets[1].value {
			value = 1
		} else {
			value = 0
		}
	case 6:
		// less than
		if value < op.subpackets[1].value {
			value = 1
		} else {
			value = 0
		}
	case 7:
		// equals
		if value == op.subpackets[1].value {
			value = 1
		} else {
			value = 0
		}
	}
	return value
}

func decodeOuterPacket(binaryPacket string) (Packet, string) {
	outermostPacket, remainingPayload := MakePacket(binaryPacket)
	return parsePacket(outermostPacket, remainingPayload)
}

func SumVersions(p Packet) int {
	s := p.version
	for _, v := range p.subpackets {
		s += SumVersions(v)
	}
	return s
}

func day2116Func(cmd *cobra.Command, args []string) {
	hexPackets, err := utils.Opener("data/2116.txt", true)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	binaryPackets := HexToBin(hexPackets)
	finalPacket, leftover := decodeOuterPacket(binaryPackets)

	fmt.Println(finalPacket, "Leftover:", leftover)
	fmt.Println("Version summary:", SumVersions(finalPacket))
	fmt.Println("Outer Packet Total:", finalPacket.value)

}
