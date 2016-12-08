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

import "github.com/DataDrake/waterlog/level"

// Style is a style definition used for prints
type Style struct {
	Level  uint8
	Msg    string
	Format string
}

const (
	// DebugFmt is the format string used by Debug prints
	DebugFmt = "\033[30;48;5;099m ✚ \033[7m %s \033[27m %-7s \033[49;38;5;99m %v\033[0m"
	// ErrorFmt is the format string used by Error prints
	ErrorFmt = "\033[30;48;5;208m ✗ \033[7m %s \033[27m %-7s \033[49;38;5;208m %v\033[0m"
	// FatalFmt is the format string used by Fatal prints
	FatalFmt = "\033[30;48;5;160m 🕱 \033[7m %s \033[27m %-7s \033[49;38;5;160m %v\033[0m"
	// GoodFmt is the format string used by Good prints
	GoodFmt = "\033[30;48;5;040m 🗸 \033[7m %s \033[27m %-7s \033[49;38;5;040m %v\033[0m"
	// InfoFmt is the format string used by Info prints
	InfoFmt = "\033[30;48;5;004m ⮞ \033[7m %s \033[27m %-7s \033[49;38;5;004m %v\033[0m"
	// PanicFmt is the format string used by Panic prints
	PanicFmt = "\033[30;48;5;200m 😮 \033[7m %s \033[27m %-7s \033[49;38;5;200m %v\033[0m"
	// WarnFmt is the format string used by Warn prints
	WarnFmt = "\033[30;48;5;220m 🗲 \033[7m %s \033[27m %-7s \033[49;38;5;220m %v\033[0m"
)

// Debug style used for prints
var Debug = Style{level.Debug, "DEBUG", DebugFmt}

// Error style used for prints
var Error = Style{level.Error, "ERROR", ErrorFmt}

// Fatal style used for prints
var Fatal = Style{level.Fatal, "FATAL", FatalFmt}

// Good style used for prints
var Good = Style{level.Good, "GOOD", GoodFmt}

// Info style used for prints
var Info = Style{level.Info, "INFO", InfoFmt}

// Panic style used for prints
var Panic = Style{level.Panic, "PANIC", PanicFmt}

// Warn style used for prints
var Warn = Style{level.Warn, "WARNING", WarnFmt}
