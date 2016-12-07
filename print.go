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
	"fmt"
	"github.com/DataDrake/waterlog/format"
)

func (w *WaterLog) eprint(s format.Style, v ...interface{}) {
	if s.Level <= w.level {
		v = append([]interface{}{w.Time(), s.Msg}, v...)
		fmt.Printf(s.Format, v...)
	}
}

func (w *WaterLog) eprintf(s format.Style, f string, v ...interface{}) {
	if s.Level <= w.level {
		v = append([]interface{}{w.Time(), s.Msg}, fmt.Sprintf(f, v...))
		fmt.Printf(s.Format, v...)
	}
}

func (w *WaterLog) eprintln(s format.Style, v ...interface{}) {
	if s.Level <= w.level {
		v = append([]interface{}{w.Time(), s.Msg}, v...)
		fmt.Printf(s.Format+"\n", v...)
	}
}

func (w *WaterLog) Error(v ...interface{}) {
	w.eprint(format.ERROR, v...)
}

func (w *WaterLog) Errorf(f string, v ...interface{}) {
	w.eprintf(format.ERROR, f, v...)
}

func (w *WaterLog) Errorln(v ...interface{}) {
	w.eprintln(format.ERROR, v...)
}

func (w *WaterLog) Debug(v ...interface{}) {
	w.eprint(format.DEBUG, v...)
}

func (w *WaterLog) Debugf(f string, v ...interface{}) {
	w.eprintf(format.DEBUG, f, v...)
}

func (w *WaterLog) Debugln(v ...interface{}) {
	w.eprintln(format.DEBUG, v...)
}

func (w *WaterLog) Fatal(v ...interface{}) {
	w.eprint(format.FATAL, v...)
}

func (w *WaterLog) Fatalf(f string, v ...interface{}) {
	w.eprintf(format.FATAL, f, v...)
}

func (w *WaterLog) Fatalln(v ...interface{}) {
	w.eprintln(format.FATAL, v...)
}

func (w *WaterLog) Good(v ...interface{}) {
	w.eprint(format.GOOD, v...)
}

func (w *WaterLog) Goodf(f string, v ...interface{}) {
	w.eprintf(format.GOOD, f, v...)
}

func (w *WaterLog) Goodln(v ...interface{}) {
	w.eprintln(format.GOOD, v...)
}

func (w *WaterLog) Info(v ...interface{}) {
	w.eprint(format.INFO, v...)
}

func (w *WaterLog) Infof(f string, v ...interface{}) {
	w.eprintf(format.INFO, f, v...)
}

func (w *WaterLog) Infoln(v ...interface{}) {
	w.eprintln(format.INFO, v...)
}

func (w *WaterLog) Output(calldepth int, s string) error {
	return nil
}

func (w *WaterLog) Panic(v ...interface{}) {
	w.eprint(format.ERROR, v...)
}

func (w *WaterLog) Panicf(f string, v ...interface{}) {
	w.eprintf(format.PANIC, f, v...)
}

func (w *WaterLog) Panicln(v ...interface{}) {
	w.eprintln(format.PANIC, v...)
}

func (w *WaterLog) Print(v ...interface{}) {
	fmt.Print(v...)
}

func (w *WaterLog) Printf(f string, v ...interface{}) {
	fmt.Printf(f, v...)
}
func (w *WaterLog) Println(v ...interface{}) {
	fmt.Println(v...)
}

func (w *WaterLog) Warn(v ...interface{}) {
	w.eprint(format.WARN, v...)
}

func (w *WaterLog) Warnf(f string, v ...interface{}) {
	w.eprintf(format.WARN, f, v...)
}

func (w *WaterLog) Warnln(v ...interface{}) {
	w.eprintln(format.WARN, v...)
}
