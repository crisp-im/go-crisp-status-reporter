# go-crisp-status-reporter

[![Build Status](https://img.shields.io/travis/crisp-im/go-crisp-status-reporter/master.svg)](https://travis-ci.org/crisp-im/go-crisp-status-reporter)

**Crisp Status Reporter for Golang.**

Crisp Status Reporter is used to actively submit health information to Crisp Status from your apps. Apps are best monitored via application probes, which are able to report detailed system information such as CPU and RAM load. This lets Crisp Status show if an application host system is under high load.

## How to use?

### Create reporter

`crisp-status-reporter` can be instantiated as such:

```go
import (
  Reporter "github.com/crisp-im/go-crisp-status-reporter/crisp_status_reporter"
  "time"
)

// Build reporter
builder := Reporter.New("YOUR_TOKEN_SECRET")

// Service ID containing the parent Node for Replica (given by Crisp)
// Node ID containing Replica (given by Crisp)
// Unique Replica ID for instance (ie. your IP on the LAN)
// Reporting interval (in seconds; defaults to 30 seconds if not set)
reporter := builder.ServiceID("YOUR_SERVICE_ID").NodeID("YOUR_NODE_ID").ReplicaID("192.168.1.10").Interval(time.Duration(30 * time.Second)).Build()

// Run reporter (starts reporting)
reporter.Run()
```
