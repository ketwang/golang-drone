package termui

import (
	ui "github.com/gizak/termui/v3"
	"image"
	"sort"
)

const MAXLINES  = 7

type LineGraph struct {
	*ui.Block
	Labels []string
	Data   map[string][]float64
	Colors map[string]ui.Color
	MaxValue float64
}

func NewLineGraph(maxValue float64) *LineGraph {
	lg := &LineGraph{
		Block: ui.NewBlock(),
		Labels: make([]string, 0),
		Data: make(map[string][]float64),
		Colors: make(map[string]ui.Color),
	}

	return lg
}


type LineWidget struct {
	*LineGraph
	format func(key string, value []float64) string
	maxValue  float64
	previousData map[string]float64
}

func NewLineWidget(maxValue float64, format func(key string, value []float64) string) *LineWidget {
	lw := &LineWidget{
		LineGraph: NewLineGraph(maxValue),
		format: format,
		maxValue: maxValue,
	}

	return lw
}

func (lw *LineWidget) Update(data map[string]float64, delta bool)  {
	keys := make([]string, len(data))
	for k, _ := range data {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	// set color
	for index, key := range keys {
		if index > 7 {
			break
		}

		lw.Colors[key] = ui.Color(index)
	}

	deletedKeys := make([]string, 0)
	for _, key := range lw.Labels {
		if _, ok := data[key]; !ok {
			deletedKeys = append(deletedKeys, key)
		} else {
			if delta {
				lw.Data[key] = append(lw.Data[key], data[key] - lw.previousData[key])
			} else {
				lw.Data[key] = append(lw.Data[key], data[key])
			}
		}
	}

	//add new entry
	for _, key := range keys {
		if _, ok := lw.Data[key]; ! ok {
			lw.Data[key] = make([]float64, 0)
			lw.Data[key] = append(lw.Data[key], 0)
		}
	}

	//del old entry
	for _, key := range deletedKeys {
		delete(lw.Data, key)
	}

	lw.Labels = keys
	lw.previousData = data
}


func (lw *LineWidget) Draw(buf *ui.Buffer)  {
	lw.Block.Draw(buf)

	yStart := lw.Inner.Min.Y + 1
	yEnd := lw.Inner.Max.Y - 1
	xStart := lw.Inner.Min.X + 1
	xEnd := lw.Inner.Max.X - 1

	index := 0
	for _, key := range lw.Labels {
		if yStart + index > yEnd {
			break
		}
		for length, char := range lw.format(key, lw.Data[key]) {
			if length > xEnd {
				break
			}
			buf.SetCell(ui.NewCell(rune(char), ui.NewStyle(lw.Colors[key])), image.Point{X: xStart + length, Y: yStart + index})
		}
		index++
	}
}
