package main

import (
    "github.com/hajimehoshi/ebiten/v2"
)

type Monitor struct {
    monitors []*ebiten.MonitorType
	currentMonitor *ebiten.MonitorType
}

func NewMonitor() *Monitor {
    var monitors []*ebiten.MonitorType
    monitors = ebiten.AppendMonitors(monitors)
    return &Monitor{
        monitors: monitors,
		currentMonitor: monitors[0],
    }
}

func (m *Monitor) NextMonitor() {
    if len(m.monitors) <= 1 {
        return
    }
    // モニターの切り替え
	for i, monitor := range m.monitors {
        if monitor == m.currentMonitor {
            m.currentMonitor = m.monitors[(i+1)%len(m.monitors)]
            break
        }
    }
}

func (m *Monitor) PreviousMonitor() {
    if len(m.monitors) <= 1 {
        return
    }
    // モニターの切り替え
    for  i, monitor := range m.monitors {
        if monitor == m.currentMonitor {
            m.currentMonitor = m.monitors[(i-1+len(m.monitors))%len(m.monitors)]
            break
        }
    }
}
