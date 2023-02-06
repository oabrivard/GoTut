package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type beacon struct {
	x int
	y int
}

type sensor struct {
	x             int
	y             int
	closestBeacon beacon
	distance      int
}

type linePart struct {
	start    int
	end      int
	isBeacon bool
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	fileName := os.Args[1]

	readFile, err := os.Open(fileName)
	check(err)
	defer readFile.Close()

	LIMIT, err := strconv.Atoi(os.Args[2])
	check(err)

	fileScanner := bufio.NewScanner(readFile)
	sensors := []sensor{}

	for fileScanner.Scan() {
		line := strings.TrimSpace(fileScanner.Text())

		re := regexp.MustCompile(`-?\d+`)
		ints := []int{}
		for _, s := range re.FindAllString(line, -1) {
			value, err := strconv.Atoi(s)
			check(err)
			ints = append(ints, value)
		}

		s := sensor{
			x:             ints[0],
			y:             ints[1],
			closestBeacon: beacon{ints[2], ints[3]},
			distance:      int(math.Abs(float64(ints[0]-ints[2]))) + int(math.Abs(float64(ints[1]-ints[3]))),
		}

		sensors = append(sensors, s)
		fmt.Println(s)
	}

	for TARGET := 0; TARGET <= LIMIT; TARGET++ {

		parts := []linePart{}
		beacons := map[beacon]bool{}
		for _, s := range sensors {
			if s.closestBeacon.y == TARGET {
				if !beacons[s.closestBeacon] {
					part := linePart{
						s.closestBeacon.x,
						s.closestBeacon.x,
						true,
					}
					parts = append(parts, part)
					beacons[s.closestBeacon] = true
				}
			}

			if s.y-s.distance <= TARGET && TARGET <= s.y+s.distance {
				heightDelta := int(math.Abs(float64(s.y - TARGET)))

				segmentWidth := s.distance - heightDelta
				part := linePart{
					s.x - segmentWidth,
					s.x + segmentWidth,
					false,
				}
				parts = append(parts, part)
				// also adds point where sensor stands, hence considering it a non option for the beacon
			}
		}

		sort.Slice(parts, func(i, j int) bool {
			if parts[i].start == parts[j].start {
				return parts[i].end < parts[j].end
			} else {
				return parts[i].start < parts[j].start
			}
		})

		//fmt.Println(parts)
		//fmt.Println(beacons)

		total := 0
		for i := 1; i < len(parts); i++ {
			curr := &parts[i]
			prev := &parts[i-1]

			// by construction, curr.start >= prev.start

			if curr.start > prev.end {
				// disjoint
				if curr.start <= LIMIT && curr.end <= LIMIT {
					diff := int(math.Max(float64(curr.start), 0.0)) - int(math.Max(float64(prev.end), 0.0)) - 1

					if diff == 1 {
						println(curr.start-1, TARGET, (curr.start-1)*4000000+TARGET)
						return
					}

					if diff > 0 {
						total += diff
					}
				}

				//fmt.Println("subtotal 1 :", total, prev, curr)

				if i == len(parts)-1 {

					if curr.start <= LIMIT && curr.end <= LIMIT {
						total += LIMIT - curr.end - 1
					}

					//fmt.Println("subtotal 4 :", total, prev, curr)
				}

			} else if curr.end > prev.end {
				// intersect
				//fmt.Println("subtotal 2 :", total, prev, curr)

				if i == len(parts)-1 {
					//fmt.Println("subtotal 5 :", total, prev, curr)
				}
			} else {
				// take beacon into account
				if curr.isBeacon {
					//				total--
				}

				// curr is included so switch curr with prev
				//fmt.Println("subtotal 3 :", total, prev, curr)
				curr.end = prev.end
				curr.start = prev.start
			}
		}

		//fmt.Println(total)
	}
	fmt.Println("Finished!")
}
