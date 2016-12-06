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

func (w *WaterLog) Error(v ...interface{}) {
	if ERROR <= w.level {
        v = append([]interface{}{w.Time(),"ERROR"}, v...)
		fmt.Printf(format.BAD, v...)
	}
}

func (w *WaterLog) Errorf(f string, v ...interface{}) {
	if ERROR <= w.level {
        v = []interface{}{w.Time(),"ERROR", fmt.Sprintf(f,v...)}
        fmt.Printf(format.BAD, v...)
	}
}

func (w *WaterLog) Errorln(v ...interface{}) {
	if ERROR <= w.level {
        v = append([]interface{}{w.Time(),"ERROR"}, v...)
		fmt.Printf(format.BAD+"\n", v...)
	}
}

func (w *WaterLog) Debug(v ...interface{}) {
	if DEBUG <= w.level {

	}
}

func (w *WaterLog) Debugf(format string, v ...interface{}) {
	if DEBUG <= w.level {

	}
}

func (w *WaterLog) Debugln(v ...interface{}) {
	if DEBUG <= w.level {

	}
}

func (w *WaterLog) Fatal(v ...interface{}) {
	if FATAL <= w.level {

	}
}

func (w *WaterLog) Fatalf(format string, v ...interface{}) {
	if FATAL <= w.level {

	}
}

func (w *WaterLog) Fatalln(v ...interface{}) {
	if FATAL <= w.level {

	}
}

func (w *WaterLog) Good(v ...interface{}) {
	if GOOD <= w.level {

	}
}

func (w *WaterLog) Goodf(format string, v ...interface{}) {
	if GOOD <= w.level {

	}
}

func (w *WaterLog) Goodln(v ...interface{}) {
	if GOOD <= w.level {

	}
}

func (w *WaterLog) Info(v ...interface{}) {
	if INFO <= w.level {

	}
}

func (w *WaterLog) Infof(format string, v ...interface{}) {
	if INFO <= w.level {

	}
}

func (w *WaterLog) Infoln(v ...interface{}) {
	if INFO <= w.level {

	}
}

func (w *WaterLog) Output(calldepth int, s string) error {
	if PANIC <= w.level {

	}
    return nil
}

func (w *WaterLog) Panic(v ...interface{}) {
	if PANIC <= w.level {

	}
}

func (w *WaterLog) Panicf(format string, v ...interface{}) {
	if PANIC <= w.level {

	}
}

func (w *WaterLog) Panicln(v ...interface{}) {
	if PANIC <= w.level {

	}
}

func (w *WaterLog) Print(v ...interface{}) {

}

func (w *WaterLog) Printf(format string, v ...interface{}) {

}
func (w *WaterLog) Println(v ...interface{}) {

}

func (w *WaterLog) Warn(v ...interface{}) {
	if WARN <= w.level {

	}
}

func (w *WaterLog) Warnf(format string, v ...interface{}) {
	if WARN <= w.level {

	}
}

func (w *WaterLog) Warnln(v ...interface{}) {
	if WARN <= w.level {

	}
}
