// Copyright 2016 DeepFabric, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific language governing permissions and
// limitations under the License.

package pdserver

import (
	"flag"
	"io"

	"github.com/coreos/pkg/capnslog"
	"github.com/deepfabric/elasticell/pkg/log"
)

var (
	logFile  = flag.String("log-file", "", "The external log file. Default log to console.")
	logLevel = flag.String("log-level", "info", "The log level, default is info")
)

// RedirectEmbedEctdLog because of our used embed ectd,
// so we need redirect ectd log to spec.
func RedirectEmbedEctdLog(w io.Writer) {
	capnslog.SetFormatter(capnslog.NewPrettyFormatter(w, false))
}

// InitLog init pd log
func InitLog() {
	log.SetRotateByHour()
	log.SetHighlighting(false)
	log.SetLevelByString(*logLevel)
	if "" != *logFile {
		log.SetOutputByName(*logFile)
	}

	if !log.DebugEnabled() {
		log.SetFlags(log.Ldate | log.Ltime)
	}
}