package terminalui

import (
	ui "github.com/gizak/termui/v3"
	"image"
	"sort"
)

const (
	MAXLINES = 7
	SCALE    = 4
)

type LineGraph struct {
	*ui.Block
	Labels   []string
	Data     map[string][]float64
	Colors   map[string]ui.Color
	MaxValue float64
}

func NewLineGraph(maxValue float64) *LineGraph {
	lg := &LineGraph{
		Block:    ui.NewBlock(),
		Labels:   make([]string, 0),
		Data:     make(map[string][]float64),
		Colors:   make(map[string]ui.Color),
		MaxValue: maxValue,
	}

	return lg
}

type LineWidget struct {
	*LineGraph
	format       func(key string, value []float64) string
	previousData map[string]float64
	delta        bool
}

func NewLineWidget(maxValue float64, format func(key string, value []float64) string, delta bool) *LineWidget {
	lw := &LineWidget{
		LineGraph:    NewLineGraph(maxValue),
		format:       format,
		previousData: make(map[string]float64),
		delta:        delta,
	}

	return lw
}

func (lw *LineWidget) Update(data map[string]float64) {
	keys := make([]string, 0)
	for k := range data {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	// set color
	for index, key := range keys {
		if index > 7 {
			break
		}

		lw.Colors[key] = ui.Color(index + 1)
	}

	deletedKeys := make([]string, 0)
	for _, key := range lw.Labels {
		if _, ok := data[key]; !ok {
			deletedKeys = append(deletedKeys, key)
		} else {
			if lw.delta {
				lw.Data[key] = append(lw.Data[key], data[key]-lw.previousData[key])
			} else {
				lw.Data[key] = append(lw.Data[key], data[key])
			}
		}
	}

	//add new entry
	for _, key := range keys {
		if _, ok := lw.Data[key]; !ok {
			lw.Data[key] = make([]float64, 0)
			if !lw.delta {
				lw.Data[key] = append(lw.Data[key], data[key])
			}
		}
	}

	//del old entry
	for _, key := range deletedKeys {
		delete(lw.Data, key)
	}

	lw.Labels = keys
	lw.previousData = data
}

func (lw *LineWidget) Draw(buf *ui.Buffer) {
	lw.Block.Draw(buf)

	yStart := lw.Inner.Min.Y
	yEnd := lw.Inner.Max.Y
	xStart := lw.Inner.Min.X
	xEnd := lw.Inner.Max.X

	// draw label
	height := 0
	for _, key := range lw.Labels {
		if yStart+height > yEnd {
			return
		}
		for length, char := range lw.format(key, lw.Data[key]) {
			if length > xEnd {
				break
			}
			buf.SetCell(ui.NewCell(rune(char), ui.NewStyle(lw.Colors[key])), image.Point{X: xStart + length, Y: yStart + height})
		}
		height++
	}

	// draw line
	/*
		newRect := image.Rectangle{
			Min: image.Point{
				X: xStart,
				Y: yStart + height,
			},
			Max: image.Point{
				X: xEnd,
				Y: yStart,
			},
		}
		canvas := ui.NewCanvas()
		canvas.Rectangle = newRect

		normalization := float64(newRect.Dy()) / lw.maxValue

		for _, key := range lw.Labels {
			if len(lw.Data[key])*SCALE > newRect.Dx() {
				lw.Data[key] = lw.Data[key][1:]
			}

			if len(lw.Data[key]) < 2 {
				continue
			}

			data := lw.Data[key]
			//length := len(data)
			previousValue := data[0]
			for index, value := range data[1:] {
				canvas.SetLine(
					image.Point{
						X: (index + 1) * SCALE,
						Y: int(previousValue * normalization),
					},
					image.Point{
						X: index * SCALE,
						Y: int(value * normalization),
					},
					lw.Colors[key],
				)
				previousValue = value
			}
		}

		canvas.Draw(buf)
	*/

}
