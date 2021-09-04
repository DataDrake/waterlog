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

package main

import (
	log "github.com/DataDrake/waterlog"
	"github.com/DataDrake/waterlog/format"
	"github.com/DataDrake/waterlog/level"
	log2 "log"
)

func main() {
	log.SetFlags(log2.Ltime | log2.Ldate | log2.LUTC)
	log.SetLevel(level.Debug)
	log.SetFormat(format.Partial)
	for _, fmt := range []format.Func{format.Full, format.Partial, format.Min, format.Un} {
		log.SetFormat(fmt)
		log.Debugln("This is a DEBUG")
		log.Errorln("This is an ERROR")
		log.Goodln("This is a GOOD")
		log.Infoln("This is an INFO")
		log.Println("This is a PRINT")
		log.Warnln("This is a WARNING")
	}
	//log.Fatalln("This is a FATAL")
	log.Panicln("This is a PANIC")
}
