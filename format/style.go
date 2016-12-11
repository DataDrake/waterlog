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

package format

import (
	"fmt"
	"github.com/DataDrake/waterlog/level"
)

const (
	// Full indicates a verbose printing style
	Full = uint8(0)
	// Min indicates a minimal printing style
	Min = uint8(1)
	// Un indicated a Full style without colors
	Un = uint8(2)
)

// Style is a style definition used for prints
type Style struct {
	Color  string
	Level  uint8
	Msg    string
	Symbol string
}

const (
	fullFmt = "\033[30;48;5;%sm %s \033[7m %s \033[27m %-7s \033[49;38;5;%sm %v\033[0m"
	minFmt  = "\033[30;48;5;%sm %s \033[49;38;5;%sm %v\033[0m"
	unFmt   = " %s  %s  %-7s  %v"
)

// Debug style used for prints
var Debug = Style{"099", level.Debug, "DEBUG", "✚"}

// Error style used for prints
var Error = Style{"208", level.Error, "ERROR", "✗"}

// Fatal style used for prints
var Fatal = Style{"160", level.Fatal, "FATAL", "🕱"}

// Good style used for prints
var Good = Style{"040", level.Good, "GOOD", "🗸"}

// Info style used for prints
var Info = Style{"004", level.Info, "INFO", "⮞"}

// Panic style used for prints
var Panic = Style{"200", level.Panic, "PANIC", "😮"}

// Warn style used for prints
var Warn = Style{"220", level.Warn, "WARNING", "🗲"}

// Full prints a complete log message, with a timestep and label for message type
func (s Style) Full(time string, v ...interface{}) string {
	if len(v) <= 1 {
		return fmt.Sprintf(fullFmt, s.Color, s.Symbol, time, s.Msg, s.Color, v[0])
	}
	return fmt.Sprintf(fullFmt, s.Color, s.Symbol, time, s.Msg, s.Color, v)
}

// Min prints a simplified log message, with only an icon for message type
func (s Style) Min(v ...interface{}) string {
	if len(v) <= 1 {
		return fmt.Sprintf(minFmt, s.Color, s.Symbol, s.Color, v[0])
	}
	return fmt.Sprintf(minFmt, s.Color, s.Symbol, s.Color, v)
}

// Un prints a Full-style log message, without colors
func (s Style) Un(time string, v ...interface{}) string {
	if len(v) <= 1 {
		return fmt.Sprintf(unFmt, s.Symbol, time, s.Msg, v[0])
	}
	return fmt.Sprintf(unFmt, s.Symbol, time, s.Msg, v)
}
