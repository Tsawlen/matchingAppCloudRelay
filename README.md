# CloudRelay Service
This service works as a pub/sub server in golang, where services can subscribe and publish to.

## Setup
Just download the repository and run `docker compose up` (WIP)

## Commands
- To subscribe to a topic the subscriber has to open a websocket connection to: `<ip>:8082/subscribe` with a header containing the topic
- To publish the REST endpoint `<ip>:8082/publish` has to be called. In the headers the fields `topic` and `service` are required. The body can contain any json, which should be delivered to all subscribers