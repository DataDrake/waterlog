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
	"github.com/DataDrake/flair"
)

const (
	unFmt = " %s  %s  %-7s  %v"
)

// Func defines a format function to use for printing
type Func func(s Style, time string, v ...interface{}) string

// Full prints a complete log message, with a timestep and label for message type
func Full(s Style, time string, v ...interface{}) string {
	full := s.BG() + " %s " + flair.Reverse(" %s ") + "  %-8s " + flair.DefaultBG + s.FGFunc()(" %v")
	args := []interface{}{s.Symbol, time, s.Msg}
	args = append(args, v...)
	return fmt.Sprintf(full, args...)
}

// Partial prints a partial log message, with a timestep and icon for message type
func Partial(s Style, time string, v ...interface{}) string {
	partial := s.BG() + " %s %-8s " + flair.Reset + " %v"
	args := []interface{}{s.Symbol, s.Msg}
	args = append(args, v...)
	return fmt.Sprintf(partial, args...)
}

// Min prints a simplified log message, with only an icon for message type
func Min(s Style, time string, v ...interface{}) string {
	min := s.BG() + " %s " + flair.DefaultBG + s.FGFunc()(" %v")
	args := []interface{}{s.Symbol}
	args = append(args, v...)
	return fmt.Sprintf(min, args...)
}

// Un prints a Full-style log message, without colors
func Un(s Style, time string, v ...interface{}) string {
	args := []interface{}{s.Symbol, time, s.Msg}
	args = append(args, v...)
	return fmt.Sprintf(unFmt, args...)
}
