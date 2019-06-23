package ui

import (
	"context"
	"fmt"
	ui "github.com/gizak/termui/v3"
	"github.com/spf13/cobra"
	"math/rand"
	"time"
	"util/pkg/convert"
	singal2 "util/pkg/singal"
	"util/pkg/terminalui"
)

var (
	UiCommand = &cobra.Command{
		Use:  "stats",
		Args: cobra.ExactArgs(1),
		Long: "show xx stats",
		RunE: uiServe,
	}
)

func format(key string, value []float64) string {
	length := len(value)
	if length > 1 {
		return fmt.Sprintf("%s: %s", key, convert.Size(value[length-1]).String())
	}

	return fmt.Sprintf("%s: ", key)
}

func uiServe(cmd *cobra.Command, args []string) error {
	if err := ui.Init(); err != nil {
		return err
	}
	defer ui.Close()

	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	lw := terminalui.NewLineWidget(100, format, false)
	lw.Title = "hahahah"
	lw.Border = true
	p1 := ui.NewRow(2.0/3, ui.NewCol(1.0, lw))

	entries := make([]interface{}, 0)
	entries = append(entries, p1)

	grid := ui.NewGrid()
	grid.Set(entries...)

	ctx := singal2.WitchSingalsContext(context.Background())

	for {
		lw.Update(map[string]float64{
			"rx": float64(rand.Intn(100)),
			"tx": float64(rand.Intn(100)),
		})
		ui.Render(grid)
		select {
		case <-ticker.C:
		case <-ctx.Done():
			return nil
		}
	}
}
