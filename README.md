# Nest Doorbell events to Openhab

## Getting started

Following this guide:https://developers.google.com/nest/device-access/registration?authuser=1
Using tom.deckers@gmail.com

Device Access Console: https://console.nest.google.com/u/1/device-access/project-list
* Doorbell project

Google Cloud project: https://console.cloud.google.com/home/dashboard?authuser=1&project=doorbell-1618209550105
* Subscribe to SDM Events through pub/sub subscription.
* View events in gcloud console: `gcloud pubsub subscriptions pull doorbell-events`

## Google Cloud Function: event to openhab

This project represents the google cloud function that receives events and forwards
them to Openhab.

Environment variables:
* Openhab item link (to post to)