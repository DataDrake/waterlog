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
	"github.com/DataDrake/waterlog/level"
)

// Style is a style definition used for prints
type Style struct {
	Color  string
	Level  uint8
	Msg    string
	Symbol string
}

// Debug style used for prints
var (
	Debug = Style{"099", level.Debug, "DEBUG", "✚"}
	Error = Style{"208", level.Error, "ERROR", "✗"}
	Fatal = Style{"160", level.Fatal, "FATAL", "🕱"}
	Good  = Style{"040", level.Good, "GOOD", "🗸"}
	Info  = Style{"045", level.Info, "INFO", "⮞"}
	Panic = Style{"200", level.Panic, "PANIC", "☢"}
	Warn  = Style{"220", level.Warn, "WARNING", "🗲"}
)
