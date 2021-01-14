//
// Copyright 2017-2021 Bryan T. Meyers <root@datadrake.com>
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
	"os"
	"sync"
	"time"
)

// WaterLog is a styled log.Logger
type WaterLog struct {
	flag   int
	format format.Func
	level  uint8
	mu     sync.Mutex
	prefix string
	output io.Writer
}

// New creates a WaterLog
func New(out io.Writer, prefix string, flag int) *WaterLog {
	return &WaterLog{flag, format.Full, level.Fatal, sync.Mutex{}, prefix, out}
}

var std = New(os.Stdout, "", log.Ldate|log.Ltime)

// Flags returns the output flags
func (w *WaterLog) Flags() int {
	return w.flag
}

// Flags returns the output flags for std
func Flags() int {
	return std.Flags()
}

// SetFlags replaces the output flags
func (w *WaterLog) SetFlags(flag int) {
	w.flag = flag
}

// SetFlags replaces the output flags
func SetFlags(flag int) {
	std.SetFlags(flag)
}

// Level returns the logging level
func (w *WaterLog) Level() uint8 {
	return w.level
}

// Level returns the logging level
func Level() uint8 {
	return std.Level()
}

// SetLevel changes the logging level
func (w *WaterLog) SetLevel(level uint8) {
	w.level = level
}

// SetLevel changes the logging level
func SetLevel(level uint8) {
	std.SetLevel(level)
}

// SetOutput replaces the internal io.Writer
func (w *WaterLog) SetOutput(output io.Writer) {
	w.output = output
}

// SetOutput replaces the internal io.Writer
func SetOutput(output io.Writer) {
	std.SetOutput(output)
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
func (w *WaterLog) SetFormat(format format.Func) {
	w.format = format
}

// SetFormat changes the printing format
func SetFormat(format format.Func) {
	std.SetFormat(format)
}
