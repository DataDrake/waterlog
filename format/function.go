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
	"fmt"
)

const (
	fullFmt    = "\033[30;48;5;%sm %s \033[7m %s \033[27m %-7s \033[49;38;5;%sm %v\033[0m"
	partialFmt = "\033[30;48;5;%sm %s %s \033[49;38;5;%sm %v\033[0m"
	minFmt     = "\033[30;48;5;%sm %s \033[49;38;5;%sm %v\033[0m"
	unFmt      = " %s  %s  %-7s  %v"
)

// Func defines a format function to use for printing
type Func func(s Style, time string, v ...interface{}) string

// Full prints a complete log message, with a timestep and label for message type
func Full(s Style, time string, v ...interface{}) string {
	return fmt.Sprintf(fullFmt, s.Color, s.Symbol, time, s.Msg, s.Color, v)
}

// Partial prints a partial log message, with a timestep and icon for message type
func Partial(s Style, time string, v ...interface{}) string {
	return fmt.Sprintf(partialFmt, s.Color, time, s.Symbol, s.Color, v)
}

// Min prints a simplified log message, with only an icon for message type
func Min(s Style, time string, v ...interface{}) string {
	return fmt.Sprintf(minFmt, s.Color, s.Symbol, s.Color, v)
}

// Un prints a Full-style log message, without colors
func Un(s Style, time string, v ...interface{}) string {
	return fmt.Sprintf(unFmt, s.Symbol, time, s.Msg, v)
}
