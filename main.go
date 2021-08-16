package main

import (
	"fmt"
	tm "github.com/buger/goterm"
	"time"
)

func main() {
	testAvailableCourse()
}

func testPath(grid string)  {
	g := ParseGrid(grid)

	_, _, history := Find(g.From(), g.To(), true)

	for _, p := range history {
		s := p.noder.(*Spot)
		g[s.X][s.Y].Feature = SpotPath
		tm.Clear()
		tm.MoveCursor(1, 1)

		time.Sleep(200 * time.Millisecond)
		tm.Color("RED STRING", tm.RED)
		tm.Println("[info] ", int32(p.cost), int32(p.rank), p.closed)
		tm.Println(g.RenderPath())
		tm.Flush()
	}
}

func clearConsole() {
	fmt.Println("\033[2J")
}

func testAvailableCourse() {
	testPath(`
f......x...t
.......x....
.......x...x
..xxxx.xxx.x
x..xxx.xxx.x
.x..........
`)
}
