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

package level

const (
	// Disable turns off logging
	Disable = uint8(0)
	// Panic allows: Panic
	Panic = uint8(1)
	// Fatal allows: Panic, Fatal
	Fatal = uint8(2)
	// Error allows: Panic, Fatal, Error
	Error = uint8(3)
	// Warn allows: Panic, Fatal, Error, Warn
	Warn = uint8(4)
	// Good allows: Panic, Fatal, Error, Warn, Good
	Good = uint8(5)
	// Info allows: Panic, Fatal, Error, Warn, Good, Info
	Info = uint8(6)
	// Debug allows: Panic, Fatal, Error, Warn, Good, Info, Debug
	Debug = uint8(7)
)
