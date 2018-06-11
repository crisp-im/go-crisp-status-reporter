// Copyright 2018 Crisp IM SARL All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package crisp_status_reporter


import (
  "github.com/shirou/gopsutil/load"
  "github.com/shirou/gopsutil/cpu"
  "github.com/shirou/gopsutil/mem"
  "fmt"
  "bytes"
  "encoding/json"
  "math"
  "time"
  "io"
  "net/http"
  "net/url"
)


const (
  libraryVersion = "1.1.0"
  reportURL = "https://report.crisp.watch/v1"
  userAgent = "go-crisp-status-reporter/" + libraryVersion
  acceptContentType = "application/json"
  clientTimeout = 10
)


type ReporterBuilder interface {
  ServiceID(string) ReporterBuilder
  NodeID(string) ReporterBuilder
  ReplicaID(string) ReporterBuilder
  Interval(time.Duration) ReporterBuilder
  Build() Reporter
}

type reporterBuilder struct {
  token string
  serviceID *string
  nodeID *string
  replicaID *string
  interval *time.Duration
}

type reporterAuth struct {
  username string
  password string
}

type Reporter struct {
  httpClient *http.Client
  auth reporterAuth
  reportURL string
  token string
  replicaID string
  interval time.Duration
}

type reporterPayload struct {
  Replica   string               `json:"replica"`
  Interval  uint64               `json:"interval"`
  Load      reporterPayloadLoad  `json:"load"`
}

type reporterPayloadLoad struct {
  CPU  float32  `json:"cpu"`
  RAM  float32  `json:"ram"`
}


func (builder *reporterBuilder) ServiceID(serviceID string) ReporterBuilder {
  builder.serviceID = &serviceID

  return builder
}

func (builder *reporterBuilder) NodeID(nodeID string) ReporterBuilder {
  builder.nodeID = &nodeID

  return builder
}

func (builder *reporterBuilder) ReplicaID(replicaID string) ReporterBuilder {
  builder.replicaID = &replicaID

  return builder
}

func (builder *reporterBuilder) Interval(interval time.Duration) ReporterBuilder {
  builder.interval = &interval

  return builder
}

func (builder *reporterBuilder) Build() Reporter {
  if builder.serviceID == nil || *builder.serviceID == "" {
    panic("missing serviceID")
  }
  if builder.nodeID == nil || *builder.nodeID == "" {
    panic("missing nodeID")
  }
  if builder.replicaID == nil || *builder.replicaID == "" {
    panic("missing replicaID")
  }

  reportURL := fmt.Sprintf("%s/%s/%s/", reportURL, url.QueryEscape(*builder.serviceID), url.QueryEscape(*builder.nodeID))

  interval := time.Duration(30 * time.Second)

  if builder.interval != nil {
    interval = *builder.interval
  }

  httpClient := http.DefaultClient
  httpClient.Timeout = time.Duration(clientTimeout * time.Second)

  return Reporter {
    httpClient: httpClient,
    auth: reporterAuth {
      username: "",
      password: builder.token,
    },
    reportURL: reportURL,
    token: builder.token,
    replicaID: *builder.replicaID,
    interval: interval,
  }
}


func (reporter *Reporter) Run() {
  go reporter.manage()
}

func (reporter *Reporter) manage() {
  // Schedule first report after 10 seconds
  time.Sleep(10 * time.Second)

  for {
    if reporter.report() == false {
      // Try reporting again after half the interval (this report failed)
      time.Sleep(reporter.interval / 2)

      reporter.report()
    }

    time.Sleep(reporter.interval)
  }
}

func (reporter *Reporter) report() bool {
  // Generate report payload
  payload := reporterPayload {
    Replica: reporter.replicaID,
    Interval: uint64(reporter.interval.Seconds()),
    Load: reporterPayloadLoad {
      CPU: reporter.getLoadCPU(),
      RAM: reporter.getLoadRAM(),
    },
  }

  // Submit report payload
  req, _ := reporter.newRequest(payload)

  resp, err := reporter.httpClient.Do(req)
  if err == nil && resp != nil && resp.StatusCode == 200 {
    return true
  }

  return false
}

func (reporter *Reporter) getLoadCPU() float32 {
  systemLoad, errLoad := load.Avg()
  cpuCounts, errCPU := cpu.Counts(true)

  if errLoad == nil && errCPU == nil && systemLoad != nil {
    return float32(systemLoad.Load15 / math.Max(float64(cpuCounts), 1.0))
  }

  return 0.0
}

func (reporter *Reporter) getLoadRAM() float32 {
  memoryLoad, err := mem.VirtualMemory()

  if err == nil && memoryLoad != nil {
    return 1.00 - (float32(memoryLoad.Available) / float32(memoryLoad.Total))
  }

  return 0.0
}

func (reporter *Reporter) newRequest(body interface{}) (*http.Request, error) {
  var buf io.ReadWriter
  if body != nil {
    buf = new(bytes.Buffer)
    err := json.NewEncoder(buf).Encode(body)
    if err != nil {
      return nil, err
    }
  }

  req, err := http.NewRequest("POST", reporter.reportURL, buf)
  if err != nil {
    return nil, err
  }

  req.SetBasicAuth(reporter.auth.username, reporter.auth.password)

  req.Header.Add("Accept", acceptContentType)
  req.Header.Add("Content-Type", acceptContentType)
  req.Header.Add("User-Agent", userAgent)

  return req, nil
}


func New(token string) ReporterBuilder {
  return &reporterBuilder {
    token: token,
  }
}
