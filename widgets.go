package main

import (
	"fmt"
	"strconv"

	ui "github.com/gizak/termui"
)

type Widgets struct {
	cid    *ui.Par
	cpu    *ui.Gauge
	net    *ui.Gauge
	memory *ui.Gauge
}

func (w *Widgets) SetCPU(val int) {
	w.cpu.BarColor = colorScale(val)
	w.cpu.Label = fmt.Sprintf("%s%%", strconv.Itoa(val))
	if val < 5 && val > 0 {
		val = 5
	}
	w.cpu.Percent = val
}

func (w *Widgets) SetNet(rx int64, tx int64) {
	w.net.Label = fmt.Sprintf("%s / %s", byteFormat(rx), byteFormat(tx))
}

func (w *Widgets) SetMem(val int64, limit int64) {
	if val < 5 {
		val = 5
	}
	w.memory.Percent = round((float64(val) / float64(limit)) * 100)
	w.memory.Label = fmt.Sprintf("%s / %s", byteFormat(val), byteFormat(limit))
}

func NewWidgets(id string) *Widgets {
	cid := ui.NewPar(id)
	cid.Border = false
	cid.Height = 1
	cid.Width = 20
	cid.TextFgColor = ui.ColorWhite
	return &Widgets{cid, mkGauge(), mkGauge(), mkGauge()}
}

func mkGauge() *ui.Gauge {
	g := ui.NewGauge()
	g.Height = 1
	g.Border = false
	g.Percent = 0
	g.PaddingBottom = 0
	g.BarColor = ui.ColorGreen
	g.Label = "-"
	return g
}
