package main

import (
	"fmt"
	ui "github.com/gizak/termui/v3"
	"math/rand"
	"time"
	"util/pkg/convert"
	"util/pkg/termui"
)

func main() {
	if err := ui.Init(); err != nil {
		fmt.Println(err)
		return
	}

	format := func(key string, value []float64) string {
		if len(value) == 0 {
			return key
		}

		v := value[len(value)-1]
		return fmt.Sprintf("%s: %s/s", key, convert.Size(v).String())
	}

	lw := termui.NewLineWidget(100, format)
	lw.Title = "hahahaha"

	entries := make([]interface{}, 0)
	entries = append(entries, ui.NewRow(2.0/3, ui.NewCol(1.0, lw)))

	termWidth, termHeight := ui.TerminalDimensions()
	grid := ui.NewGrid()
	grid.SetRect(0, 0, termWidth, termHeight)
	grid.Set(entries...)

	timer := time.NewTicker(5 * time.Second)

	for {
		lw.Update(map[string]float64{"rx": float64(rand.Intn(100))}, false)
		ui.Render(grid)
		select {
		case <-timer.C:
		case <-ui.PollEvents():
			return
		}
	}

}
