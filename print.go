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
	"os"
)

func (w *WaterLog) eprint(s format.Style, v ...interface{}) {
	if s.Level <= w.level {
		if w.full {
			w.Print(s.Full(w.Time(), v...))
		} else {
			w.Print(s.Min(v...))
		}
	}
}

func (w *WaterLog) eprintf(s format.Style, f string, v ...interface{}) {
	msg := fmt.Sprintf(f, v...)
	if s.Level <= w.level {
		if w.full {
			w.Print(s.Full(w.Time(), msg))
		} else {
			w.Print(s.Min(msg))
		}
	}
}

func (w *WaterLog) eprintln(s format.Style, v ...interface{}) {
	if s.Level <= w.level {
		if w.full {
			w.Println(s.Full(w.Time(), v...))
		} else {
			w.Println(s.Min(v...))
		}
	}
}

// Error is a styled Print
func (w *WaterLog) Error(v ...interface{}) {
	w.eprint(format.Error, v...)
}

// Errorf is a styled Printf
func (w *WaterLog) Errorf(f string, v ...interface{}) {
	w.eprintf(format.Error, f, v...)
}

// Errorln is a styled Println
func (w *WaterLog) Errorln(v ...interface{}) {
	w.eprintln(format.Error, v...)
}

// Debug is a styled Print
func (w *WaterLog) Debug(v ...interface{}) {
	w.eprint(format.Debug, v...)
}

// Debugf is a styled Printf
func (w *WaterLog) Debugf(f string, v ...interface{}) {
	w.eprintf(format.Debug, f, v...)
}

// Debugln is a styled Println
func (w *WaterLog) Debugln(v ...interface{}) {
	w.eprintln(format.Debug, v...)
}

// Fatal is a styled Print followed by os.Exit(1)
func (w *WaterLog) Fatal(v ...interface{}) {
	w.eprint(format.Fatal, v...)
	os.Exit(1)
}

// Fatalf is a styled Printf followed by os.Exit(1)
func (w *WaterLog) Fatalf(f string, v ...interface{}) {
	w.eprintf(format.Fatal, f, v...)
	os.Exit(1)
}

// Fatalln is a styled Println followed by os.Exit(1)
func (w *WaterLog) Fatalln(v ...interface{}) {
	w.eprintln(format.Fatal, v...)
	os.Exit(1)
}

// Good is a styled Print
func (w *WaterLog) Good(v ...interface{}) {
	w.eprint(format.Good, v...)
}

// Goodf is a styled Printf
func (w *WaterLog) Goodf(f string, v ...interface{}) {
	w.eprintf(format.Good, f, v...)
}

// Goodln is a styled Println
func (w *WaterLog) Goodln(v ...interface{}) {
	w.eprintln(format.Good, v...)
}

// Info is a styled Print
func (w *WaterLog) Info(v ...interface{}) {
	w.eprint(format.Info, v...)
}

// Infof is a styled Printf
func (w *WaterLog) Infof(f string, v ...interface{}) {
	w.eprintf(format.Info, f, v...)
}

// Infoln is a styled Println
func (w *WaterLog) Infoln(v ...interface{}) {
	w.eprintln(format.Info, v...)
}

// Output just calls Print
func (w *WaterLog) Output(calldepth int, s string) error {
	w.Print(s)
	return nil
}

// Panic is a styled Print followed by a call to panic("")
func (w *WaterLog) Panic(v ...interface{}) {
	w.eprint(format.Panic, v...)
	panic("")
}

// Panicf is a styled Printf followed by a call to panic("")
func (w *WaterLog) Panicf(f string, v ...interface{}) {
	w.eprintf(format.Panic, f, v...)
	panic("")
}

// Panicln is a styled Println followed by a call to panic("")
func (w *WaterLog) Panicln(v ...interface{}) {
	w.eprintln(format.Panic, v...)
	panic("")
}

// Print is a mutex protect fmt.Fprint
func (w *WaterLog) Print(v ...interface{}) {
	w.mu.Lock()
	defer w.mu.Unlock()
	fmt.Fprint(w.output, v...)
}

// Printf is a mutex protect fmt.Fprintf
func (w *WaterLog) Printf(f string, v ...interface{}) {
	w.mu.Lock()
	defer w.mu.Unlock()
	fmt.Fprintf(w.output, f, v...)
}

// Println is a mutex protect fmt.Fprintln
func (w *WaterLog) Println(v ...interface{}) {
	w.mu.Lock()
	defer w.mu.Unlock()
	fmt.Fprintln(w.output, v...)
}

// Warn is a styled Print
func (w *WaterLog) Warn(v ...interface{}) {
	w.eprint(format.Warn, v...)
}

// Warnf is a styled Printf
func (w *WaterLog) Warnf(f string, v ...interface{}) {
	w.eprintf(format.Warn, f, v...)
}

// Warnln is a styled Println
func (w *WaterLog) Warnln(v ...interface{}) {
	w.eprintln(format.Warn, v...)
}
