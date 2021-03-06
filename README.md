# go-crisp-status-reporter

[![Test and Build](https://github.com/crisp-im/go-crisp-status-reporter/workflows/Test%20and%20Build/badge.svg?branch=master)](https://github.com/crisp-im/go-crisp-status-reporter/actions?query=workflow%3A%22Test+and+Build%22)

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

## Where can I find my token?

Your private token can be found on your [Crisp dashboard](https://app.crisp.chat/). Go to Settings, then Status Page, and then scroll down to "Configure your Status Reporter". Copy the secret token shown there, and use it while configuring this library in your application.

## How to add monitored node?

You can easily add a push node for the application running this library on your Crisp dashboard. Add the node, and retrieve its `service_id` and `node_id` as follows:

<p align="center">
  <img width="605" src="https://crisp-im.github.io/go-crisp-status-reporter/images/setup.gif" alt="How to add monitored node">
</p>

## Get more help

You can find more help on our helpdesk article: [How to setup the Crisp Status Reporter library?](https://help.crisp.chat/en/article/1koqk09/)

