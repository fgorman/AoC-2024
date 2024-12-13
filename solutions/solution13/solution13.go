package solution13

import (
	"fmt"
	"strconv"
	"strings"
)

const TEST = `Button A: X+94, Y+34
Button B: X+22, Y+67
Prize: X=8400, Y=5400

Button A: X+26, Y+66
Button B: X+67, Y+21
Prize: X=12748, Y=12176

Button A: X+17, Y+86
Button B: X+84, Y+37
Prize: X=7870, Y=6450

Button A: X+69, Y+23
Button B: X+27, Y+71
Prize: X=18641, Y=10279`

type MachineConfig struct {
	ButtonAX int
	ButtonAY int
	ButtonBX int
	ButtonBY int
	PrizeX   int
	PrizeY   int
}

func Solution(in string) {
	//in = TEST

	machinesConfigs := getMachinesConfigs(in)

	totalCost := getTotalCost(machinesConfigs)
	fmt.Println("Total cost for all machines:", totalCost)

	totalCostWith10000000000000 := getTotalCostWith10000000000000(machinesConfigs)
	fmt.Println("Total cost for all machines with 10000000000000:", totalCostWith10000000000000)
}

func getMachinesConfigs(in string) []MachineConfig {
	machinesConfigs := []MachineConfig{}

	in = strings.TrimSpace(in)

	machineConfigsStrs := strings.Split(in, "\n\n")

	for _, machineConfigStr := range machineConfigsStrs {
		mcStrSplit := strings.Split(machineConfigStr, "\n")

		aXYStr := strings.Split(strings.TrimSpace(strings.Split(mcStrSplit[0], ":")[1]), ",")
		bXYStr := strings.Split(strings.TrimSpace(strings.Split(mcStrSplit[1], ":")[1]), ",")
		prizeXYStr := strings.Split(strings.TrimSpace(strings.Split(mcStrSplit[2], ":")[1]), ",")

		aX, _ := strconv.Atoi(strings.Split(aXYStr[0], "+")[1])
		aY, _ := strconv.Atoi(strings.Split(aXYStr[1], "+")[1])
		bX, _ := strconv.Atoi(strings.Split(bXYStr[0], "+")[1])
		bY, _ := strconv.Atoi(strings.Split(bXYStr[1], "+")[1])
		prizeX, _ := strconv.Atoi(strings.Split(prizeXYStr[0], "=")[1])
		prizeY, _ := strconv.Atoi(strings.Split(prizeXYStr[1], "=")[1])

		machinesConfigs = append(machinesConfigs, MachineConfig{aX, aY, bX, bY, prizeX, prizeY})
	}

	return machinesConfigs
}

func getMachinesSolution(config MachineConfig) (numA, numB int) {
	x1, x2 := config.ButtonAX, config.ButtonBX
	y1, y2 := config.ButtonAY, config.ButtonBY
	pX, pY := config.PrizeX, config.PrizeY

	numA = (y2*pX + -x2*pY) / ((x1 * y2) - (x2 * y1))
	numB = (-y1*pX + x1*pY) / ((x1 * y2) - (x2 * y1))

	if numA*x1+numB*x2 == pX && numA*y1+numB*y2 == pY {
		return
	}

	numA = 0
	numB = 0

	return
}

func getTotalCost(configs []MachineConfig) (totalCost int) {
	for _, config := range configs {
		x, y := getMachinesSolution(config)
		if x <= 100 && y <= 100 {
			totalCost += x*3 + y
		}
	}
	return
}

func getTotalCostWith10000000000000(configs []MachineConfig) (totalCost int) {
	for _, config := range configs {
		config.PrizeX += 10000000000000
		config.PrizeY += 10000000000000

		x, y := getMachinesSolution(config)

		totalCost += x*3 + y
	}

	return
}
