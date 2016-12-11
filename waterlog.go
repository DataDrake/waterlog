//
// Copyright © 2016 Bryan T. Meyers <bmeyers@datadrake.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

// Package waterlog provides an extended and styled log.Logger functionality.
package waterlog

import (
	"github.com/DataDrake/waterlog/format"
	"github.com/DataDrake/waterlog/level"
	"io"
	"log"
	"sync"
	"time"
)

// WaterLog is a styled log.Logger
type WaterLog struct {
	flag   int
	format uint8
	level  uint8
	mu     sync.Mutex
	prefix string
	output io.Writer
}

// New creates a WaterLog
func New(out io.Writer, prefix string, flag int) *WaterLog {
	return &WaterLog{flag, format.Full, level.Fatal, sync.Mutex{}, prefix, out}
}

// Flags returns the output flags
func (w *WaterLog) Flags() int {
	return w.flag
}

// SetFlags replaces the output flags
func (w *WaterLog) SetFlags(flag int) {
	w.flag = flag
}

// Level returns the logging level
func (w *WaterLog) Level() uint8 {
	return w.level
}

// SetLevel changes the logging level
func (w *WaterLog) SetLevel(level uint8) {
	w.level = level
}

// SetOutput replaces the internal io.Writer
func (w *WaterLog) SetOutput(output io.Writer) {
	w.output = output
}

// Prefix returns the output prefix (UNUSED)
func (w *WaterLog) Prefix() string {
	return w.prefix
}

// SetPrefix replaces the output prefix (UNUSED)
func (w *WaterLog) SetPrefix(prefix string) {
	w.prefix = prefix
}

// Time returns a formatted Timestamp for logging
func (w *WaterLog) Time() string {
	layout := ""
	prev := false
	if w.flag&log.Ldate == log.Ldate {
		layout += "2006-01-02"
		prev = true
	}
	if w.flag&log.Ltime == log.Ltime {
		if prev {
			layout += " "
		}
		layout += "15:04:05"
		if w.flag&log.Lmicroseconds == log.Lmicroseconds {
			layout += ".000000"
		}
	}
	t := time.Now()
	if w.flag&log.LUTC == log.LUTC {
		t = t.UTC()
	}
	return t.Format(layout)
}

// SetFormat changes the printing format
func (w *WaterLog) SetFormat(format uint8) {
	w.format = format
}
