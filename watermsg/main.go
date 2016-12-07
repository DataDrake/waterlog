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

package main

import (
	"github.com/DataDrake/waterlog"
	"github.com/DataDrake/waterlog/level"
	"log"
	"os"
)

func main() {
	w := waterlog.New(os.Stdout, "", log.Ltime|log.Ldate|log.Lmicroseconds|log.LUTC)
	w.SetLevel(level.DEBUG)
	w.Debugln("This is a DEBUG")
	w.Errorln("This is an ERROR")
	w.Fatalln("This is a FATAL")
	w.Goodln("This is a GOOD")
	w.Infoln("This is an INFO")
	w.Panicln("This is a PANIC")
	w.Println("This is a PRINT")
	w.Warnln("This is a WARNING")
}
