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
    currentIndex := -1
    for i, monitor := range m.monitors {
        if monitor == m.currentMonitor {
            currentIndex = i
            break
        }
    }
    if currentIndex != -1 {
        m.currentMonitor = m.monitors[(currentIndex+1)%len(m.monitors)]
        ebiten.SetMonitor(m.currentMonitor)
    }
}

func (m *Monitor) PreviousMonitor() {
    if len(m.monitors) <= 1 {
        return
    }
    // モニターの切り替え
    currentIndex := -1
    for i, monitor := range m.monitors {
        if monitor == m.currentMonitor {
            currentIndex = i
            break
        }
    }
    if currentIndex != -1 {
        m.currentMonitor = m.monitors[(currentIndex-1+len(m.monitors))%len(m.monitors)]
        ebiten.SetMonitor(m.currentMonitor)
    }
}
