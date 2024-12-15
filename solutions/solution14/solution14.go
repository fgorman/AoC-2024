package solution14

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

const TEST = `p=0,4 v=3,-3
p=6,3 v=-1,-3
p=10,3 v=-1,2
p=2,0 v=2,-1
p=0,0 v=1,3
p=3,0 v=-2,-2
p=7,6 v=-1,-3
p=3,0 v=-1,-2
p=9,3 v=2,3
p=7,3 v=-1,2
p=2,4 v=2,-3
p=9,5 v=-3,-3`

type Robot struct {
	PX int
	PY int
	VX int
	VY int
}

func Solution(in string) {
	robots := parseRobots(in)

	safetyFactor := determineSafetyFactor(robots, 100, 101, 103)
	fmt.Println("Safety factor after 100 seconds:", safetyFactor)

	numSecsForEasterEgg := findEasterEgg(robots, 101, 103)
	fmt.Println("Number of seconds to find easter egg:", numSecsForEasterEgg)
}

func parseRobots(in string) []Robot {
	out := []Robot{}

	in = strings.TrimSpace(in)
	lines := strings.Split(in, "\n")
	for _, line := range lines {
		r := Robot{}
		pv := strings.Split(line, " ")
		pos := strings.Split(strings.Split(pv[0], "=")[1], ",")
		vel := strings.Split(strings.Split(pv[1], "=")[1], ",")

		r.PX, _ = strconv.Atoi(pos[0])
		r.PY, _ = strconv.Atoi(pos[1])
		r.VX, _ = strconv.Atoi(vel[0])
		r.VY, _ = strconv.Atoi(vel[1])

		out = append(out, r)
	}

	return out
}

func pmod(x, d int) int {
	x = x % d
	if x >= 0 {
		return x
	}
	if d < 0 {
		return x - d
	}
	return x + d
}

func moveRobot(robot *Robot, secs, width, height int) {
	robot.PX = pmod(robot.PX+robot.VX*secs, width)
	robot.PY = pmod(robot.PY+robot.VY*secs, height)
}

func determineSafetyFactor(robots []Robot, secs, width, height int) (score int) {
	quadrants := map[[4]int]int{
		{0, width/2 - 1, 0, height/2 - 1}:                  0,
		{width/2 + 1, width - 1, 0, height/2 - 1}:          0,
		{0, width/2 - 1, height/2 + 1, height - 1}:         0,
		{width/2 + 1, width - 1, height/2 + 1, height - 1}: 0,
	}

	for _, robot := range robots {
		moveRobot(&robot, secs, width, height)

		for quadrant := range quadrants {
			if robot.PX >= quadrant[0] && robot.PX <= quadrant[1] && robot.PY >= quadrant[2] && robot.PY <= quadrant[3] {
				quadrants[quadrant]++
				break
			}
		}
	}

	score = 1
	for _, num := range quadrants {
		score *= num
	}

	return
}

func easterEggFound(robot []Robot, width, height int) bool {
	grid := make([][]int, height)
	for idx := range grid {
		grid[idx] = make([]int, width)
	}

	for _, robot := range robot {
		grid[robot.PY][robot.PX]++
	}

	count := 0

	for _, line := range grid {
		for _, num := range line {
			if num == 0 {
				count = 0
			} else {
				count++
			}

			if count >= 10 {
				return true
			}
		}
	}

	return false
}

func findEasterEgg(robots []Robot, width, height int) (numSecs int) {
	for numSecs < 100000 {
		testRobots := slices.Clone(robots)
		for idx, robot := range testRobots {
			moveRobot(&robot, numSecs, width, height)
			testRobots[idx] = robot
		}

		if easterEggFound(testRobots, width, height) {
			return
		}

		numSecs++
	}

	return
}
