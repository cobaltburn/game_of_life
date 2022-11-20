package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"time"
)

const aliveCell = 'O'
const deadCell = ' '

type Point struct {
	x int
	y int
}

func main() {
	world := startWorld()
	generations, alive := 10, 0

	for i := 0; i < generations; i++ {
		world, alive = nextGeneration(world)
		fmt.Printf("Generation #%d\n", i+1)
		fmt.Printf("Alive: %d\n", alive)
		printWorld(world)
		time.Sleep(500 * time.Millisecond)
		runCmd("clear")
	}
}

func nextGeneration(world [][]rune) (nextGenWorld [][]rune, alive int) {
	nextGenWorld = emptyGrid(len(world))
	alive = 0
	for i := range world {
		for j := range world[i] {
			adjacent := findAdjacent(len(world), i, j)
			count := 0
			for _, point := range adjacent {
				if world[point.y][point.x] == aliveCell {
					count++
				}
			}
			if world[i][j] == aliveCell && count >= 2 && count <= 3 {
				nextGenWorld[i][j] = aliveCell
				alive++
			} else if world[i][j] == deadCell && count == 3 {
				nextGenWorld[i][j] = aliveCell
				alive++
			} else {
				nextGenWorld[i][j] = deadCell
			}
		}
	}
	return
}

func runCmd(name string, arg ...string) {
	cmd := exec.Command(name, arg...)
	cmd.Stdout = os.Stdout
	if err := cmd.Run(); err != nil {
		panic("error running command")
	}
}

func findAdjacent(maxSize, y, x int) (adjacent []Point) {
	adjacent = make([]Point, 8)
	index := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i == 0 && j == 0 {
				continue
			}
			pX, pY := x+j, y+i
			switch {
			case pX < 0:
				pX = maxSize - 1
			case pX >= maxSize:
				pX = 0
			}
			switch {
			case pY < 0:
				pY = maxSize - 1
			case pY >= maxSize:
				pY = 0
			}
			point := Point{y: pY, x: pX}
			adjacent[index] = point
			index++
		}
	}
	return
}

func emptyGrid(size int) (world [][]rune) {
	world = make([][]rune, size)
	for i := range world {
		world[i] = make([]rune, size)
	}
	return
}

func startWorld() (world [][]rune) {
	var size int
	fmt.Printf("input the size of the world: ")
	if _, err := fmt.Scan(&size); err != nil || size <= 0 {
		fmt.Printf("The input size is invalid\n")
		os.Exit(0)
	}
	rand.Seed(time.Now().UnixNano())
	world = emptyGrid(size)
	for i := range world {
		for j := range world[i] {
			if rand.Intn(2) == 1 {
				world[i][j] = aliveCell
			} else {
				world[i][j] = deadCell
			}
		}
	}
	return
}

func printWorld(world [][]rune) {
	for _, row := range world {
		for _, cell := range row {
			fmt.Printf("%c", cell)
		}
		fmt.Println()
	}
}
