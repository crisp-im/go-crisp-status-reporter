// Copyright 2018 Crisp IM SARL All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main


import (
  Reporter "github.com/crisp-im/go-crisp-status-reporter/crisp_status_reporter"
  "time"
)


func main() {
  // 1. Create Crisp Status Reporter
  builder := Reporter.New("REPLACE_THIS_WITH_A_SECRET_KEY")

  reporter := builder.ProbeID("relay").NodeID("socket-client").ReplicaID("192.168.1.10").Interval(time.Duration(30 * time.Second)).Build()

  // 2. Run Crisp Status Reporter
  reporter.Run()

  // 3. Schedule Crisp Status Reporter end
  time.Sleep(80 * time.Second)
}
