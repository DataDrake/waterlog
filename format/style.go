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

type Style struct {
    Level  uint8
    Msg    string
    Format string
}

const (
	DEBUG_FMT = "\033[30;48;5;099m ✚ \033[7m %s \033[27m %-7s \033[49;38;5;99m %v\033[0m"
	ERROR_FMT = "\033[30;48;5;208m ✗ \033[7m %s \033[27m %-7s \033[49;38;5;208m %v\033[0m"
	FATAL_FMT = "\033[30;48;5;160m 🕱 \033[7m %s \033[27m %-7s \033[49;38;5;160m %v\033[0m"
	GOOD_FMT  = "\033[30;48;5;040m 🗸 \033[7m %s \033[27m %-7s \033[49;38;5;040m %v\033[0m"
	INFO_FMT  = "\033[30;48;5;004m ⮞ \033[7m %s \033[27m %-7s \033[49;38;5;004m %v\033[0m"
	PANIC_FMT = "\033[30;48;5;200m 😮 \033[7m %s \033[27m %-7s \033[49;38;5;200m %v\033[0m"
	WARN_FMT  = "\033[30;48;5;220m 🗲 \033[7m %s \033[27m %-7s \033[49;38;5;220m %v\033[0m"
)


var DEBUG = Style{level.DEBUG,"DEBUG",DEBUG_FMT}
var ERROR = Style{level.ERROR,"ERROR",ERROR_FMT}
var FATAL = Style{level.FATAL,"FATAL",FATAL_FMT}
var GOOD  = Style{level.GOOD, "GOOD" ,GOOD_FMT}
var INFO  = Style{level.INFO, "INFO", INFO_FMT}
var PANIC = Style{level.ERROR,"PANIC",PANIC_FMT}
var WARN  = Style{level.ERROR,"WARNING",WARN_FMT}