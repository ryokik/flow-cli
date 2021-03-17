/*
 * Flow CLI
 *
 * Copyright 2019-2021 Dapper Labs, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package util

import (
	"fmt"
)

const (
	NoneLog  = "none"
	DebugLog = "debug"
	InfoLog  = "info"
)

// Logger interface
type Logger interface {
	Debug(string)
	Info(string)
	StartProgress(string)
	StopProgress(string)
}

// NewStdoutLogger create new logger
func NewStdoutLogger(level string) Logger {
	return &StdoutLogger{
		level: level,
	}
}

// StdoutLogger stdout logging implementation
type StdoutLogger struct {
	level   string
	spinner *Spinner
}

func (s *StdoutLogger) log(msg string, level string) {
	if s.level == NoneLog || s.level == DebugLog && s.level != level {
		return
	}

	fmt.Println(msg)
}

// Info log
func (s *StdoutLogger) Info(msg string) {
	s.log(msg, InfoLog)
}

// Debug log
func (s *StdoutLogger) Debug(msg string) {
	s.log(msg, DebugLog)
}

func (s *StdoutLogger) StartProgress(msg string) {
	if s.level == NoneLog {
		return
	}

	s.spinner = NewSpinner(msg, "")
	s.spinner.Start()
}

func (s *StdoutLogger) StopProgress(msg string) {
	if s.level == NoneLog {
		return
	}

	if s.spinner != nil {
		s.spinner.Stop(msg)
	}
}
