# Nest Doorbell events to Openhab

## Overview

Following this guide:https://developers.google.com/nest/device-access/registration?authuser=1
Using tom.deckers@gmail.com

### Google Smart Device Management: Project in Device Access Console

Device Access Console: https://console.nest.google.com/u/1/device-access/project-list
* Doorbell project
* Sets up a pubsub topic that receives all doorbell events


### Google Cloud Function: event to openhab

Google Cloud project: https://console.cloud.google.com/home/dashboard?authuser=1&project=doorbell-1618209550105
This project is used to process the doorbell events on the pubsub topic and to implement
the actual cloud function

### PubSub subscription

* Subscribe to SDM Events through pub/sub subscription 
* Push subscription with Endpoint URL pointing to the cloud function
* View events in gcloud console: `gcloud pubsub subscriptions pull doorbell-events`

### Cloud function

* The cloud function receives the event from the subscription
* Event is forwarded as is to the OPENHAB_SERVER which is configured as environment
  variable

## Development

### Setup

Install gcloud: https://cloud.google.com/sdk/docs/install
Install funcframework:

    go get github.com/GoogleCloudPlatform/functions-framework-go/funcframework
    go install github.com/GoogleCloudPlatform/functions-framework-go/funcframework

### Cloud Function

Main code is `handler.go` which represents the function.

Make commands:

`make build`:

* Builds the gcloud function.  Output is `doorbell-function`

`make deploy`:

* Deploys the function to the google cloud project

Environment variables (configured with the gcloud function)
* OPENHAB_SERVER
* OPENHAB_USER
* OPENHAB_PWD

