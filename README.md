# waterlog
Go Logging that just works

[![Go Report Card](https://goreportcard.com/badge/github.com/DataDrake/waterlog)](https://goreportcard.com/report/github.com/DataDrake/waterlog) [![license](https://img.shields.io/github/license/DataDrake/waterlog.svg)]()

## Motivation
Logging is terribly boring, but can be an important part of a project. 
This library provides logging capabilities that attempt to "spice up"
what is otherwise a very utilitarian aspect of programming.

### Goals
* Styling that is easy to look at
* More granular logging severities ("levels")
* Compatibility with the existing ```log.Logger``` interface

## Implementation
Waterlog consists of two sets of code. The first is an implementation
in Go which is able to be used as a drop-in replacement for the Go
```log.Logger```. The second is a set of Makefile definitions that
provide similar styling for GNU Make project.

## Examples

### Logger

![Go Example](/images/go.png)

### Makefile

![Makefile Example](/images/makefile.png)

## License

Copyright © 2016 Bryan T. Meyers <bmeyers@datadrake.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

