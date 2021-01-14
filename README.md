# waterlog
Go Logging that just works

[![Go Report Card](https://goreportcard.com/badge/github.com/DataDrake/waterlog)](https://goreportcard.com/report/github.com/DataDrake/waterlog) [![GoDoc](https://godoc.org/github.com/DataDrake/waterlog?status.svg)](https://godoc.org/github.com/DataDrake/waterlog) [![license](https://img.shields.io/github/license/DataDrake/waterlog.svg)]()

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

### Message Types
The traditional ```log.Logger``` implementation only provides three
different message types:

Type  | Behavior
----- | --------
Fatal | Log message and ```os.Exit(1)```
Panic | Log message and ```panic()```
Print | Log message

For each level it also supports a Print, Printf, and Println variant
of the output.

Waterlog, on the other hand, provides a total of 8 message types.
These levels also support Print, Printf, and Println output modes.
Unlike ```log.Logger```, Waterlog supports logging levels which
correspond with these message types. Higher levels allow more types
of message to be shown. The logging level may be changed at any time
by the ```SetLevel()``` method. The psuedo-type ```Disable``` is also
provided to allow quiet operation. A default level of ```Fatal``` is
assigned to new Waterlog instances. Lastly, Print messages are always
written as unformatted text with no timestamp and ignore the logging
level.

Type    | Behavior                                         | Level
------- | ------------------------------------------------ | -----
Disable | No messages                                      | 0
Panic   | Runtime "exception" message and ```panic()```    | 1
Fatal   | Unrecoverable error message and ```os.Exit(1)``` | 2
Error   | Recoverable error message                        | 3
Warn    | Warning message                                  | 4
Good    | Success message                                  | 5
Info    | General information message                      | 6
Debug   | Developer Debug information.                     | 7
Print   | Always print unstyled text                       | n/a

## Examples

### format.Full Logger

![Go format.Full Example](/images/go.png)

### format.Min Logger

![Go format.Min Example](/images/min.png)

### format.Un Logger

![Go format.Un Example](/images/un.png)

### Makefile

![Makefile Example](/images/makefile.png)

## License

Copyright 2017-2021 Bryan T. Meyers <root@datadrake.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

