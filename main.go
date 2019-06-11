package main

import (
	"fmt"
	ui "github.com/gizak/termui/v3"
	"time"
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

		return fmt.Sprintf("%s: %f", key, value[len(value) - 1])
	}

	lw := termui.NewLineWidget(100, format)
	lw.Title = "hahahaha"

	entries := make([]interface{}, 0)
	entries = append(entries, ui.NewRow(1.0/3, ui.NewCol(1.0, lw)))

	grid := ui.NewGrid()
	grid.Set(entries...)

	timer := time.NewTicker(5 * time.Second)

	for {
		lw.Update(map[string]float64{"rx": 14}, false)
		ui.Render(grid)
		select {
		case <- timer.C:
		case <- ui.PollEvents():
			return
		}
	}

}
