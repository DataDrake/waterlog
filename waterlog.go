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

package waterlog

import (
    "io"
    "time"
)

type WaterLog struct {
	flag   int
	level  uint8
	prefix string
	output io.Writer
}

func New(out io.Writer, prefix string, flag int) *WaterLog {
	return &WaterLog{flag, FATAL, prefix, out}
}

func (w *WaterLog) Flags() int {
	return w.flag
}

func (w *WaterLog) SetFlags(flag int) {
	w.flag = flag
}

func (w *WaterLog) Level() uint8 {
	return w.level
}

func (w *WaterLog) SetLevel(level uint8) {
	w.level = level
}

func (w *WaterLog) SetOutput(output io.Writer) {
	w.output = output
}

func (w *WaterLog) Prefix() string {
	return w.prefix
}

func (w *WaterLog) SetPrefix(prefix string) {
	w.prefix = prefix
}

func (w *WaterLog) Time() string {
    return time.Now().Format("15:04:05")
}