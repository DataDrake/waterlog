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

package format

import (
	"github.com/DataDrake/flair/color"
	"github.com/DataDrake/flair/escape"
	"github.com/DataDrake/flair/style"
	"github.com/DataDrake/waterlog/level"
)

// Style is a style definition used for prints
type Style struct {
	Color  color.Color
	Level  uint8
	Msg    string
	Symbol string
	fg     string
	bg     string
	fgf    style.Style
	bgf    style.Style
}

// Debug style used for prints
var (
	Debug = Style{color.Color(99), level.Debug, "DEBUG", "✚", "", "", nil, nil}
	Error = Style{color.Color(208), level.Error, "ERROR", "✗", "", "", nil, nil}
	Fatal = Style{color.Color(160), level.Fatal, "FATAL", "🕱", "", "", nil, nil}
	Good  = Style{color.Color(40), level.Good, "GOOD", "🗸", "", "", nil, nil}
	Info  = Style{color.Color(45), level.Info, "INFO", "⮞", "", "", nil, nil}
	Panic = Style{color.Color(200), level.Panic, "PANIC", "☢", "", "", nil, nil}
	Warn  = Style{color.Color(220), level.Warn, "WARNING", "🗲", "", "", nil, nil}
)

// BG provides the foreground color of this Style as a string
func (s Style) BG() string {
	if len(s.bg) > 0 {
		return s.bg
	}
	s.bg = escape.Combine(s.Color.BG(), color.Black.FG()).String()
	return s.bg
}

// FG provides the foreground color of this Style as a string
func (s Style) FG() string {
	if len(s.fg) > 0 {
		return s.fg
	}
	s.fg = escape.Combine(s.Color.FG(), color.DefaultBG).String()
	return s.fg
}

// BGFunc provides the foreground color of this Style as a function
func (s Style) BGFunc() style.Style {
	if s.bgf != nil {
		return s.bgf
	}
	s.bgf = escape.Combine(s.Color.BG(), color.Black.FG(), escape.Reset.Swap()).Func()
	return s.bgf
}

// FGFunc provides the foreground color of this Style as a function
func (s Style) FGFunc() style.Style {
	if s.fgf != nil {
		return s.fgf
	}
	s.fgf = escape.Combine(s.Color.FG(), color.DefaultBG, escape.Reset.Swap()).Func()
	return s.fgf
}
