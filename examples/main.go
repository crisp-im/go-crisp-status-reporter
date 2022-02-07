// Copyright 2018 Crisp IM SAS All rights reserved.
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

  reporter := builder.ServiceID("d657b4c1-dd07-4f94-ac7a-d4c3b4b219c1").NodeID("5eca824b-4134-4126-982d-2c2338ecf3ab").ReplicaID("192.168.1.10").Interval(time.Duration(30 * time.Second)).Build()

  // 2. Run Crisp Status Reporter
  reporter.Run()

  // 3. Schedule Crisp Status Reporter end
  time.Sleep(80 * time.Second)
}
