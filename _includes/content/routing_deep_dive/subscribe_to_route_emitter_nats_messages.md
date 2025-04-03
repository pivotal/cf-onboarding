
### What?
There is one Route Emitter per Diego Cell and its job is to... emit routes.

The Route Emitter tracks the desired route mappings for all of the apps it shares a Cell with. When a change is detected, the Route Emitter emits route registration and unregistration messages to the GoRouter via the NATS message bus. Even when no change is detected, the Route Emitter will periodically emit the current state of the world as a kind of heartbeat.

For this story, let's look at the messages that the Route Emitter is publishing via NATS. Subscribing to these NATs messages can be a helpful debugging technique when it seems like routes are not being registered properly.

### Prerequisites

Go through the **Basic BOSH Knowledge** story to gain familiarity with BOSH, specifically BOSH CLI commands.

### How?
0. Bosh ssh onto a Diego Cell
0. Download ruby and the NATS gem
    ```
    apt-get install -y ruby ruby-dev
    gem install nats
    ```
0. Get the NATS username, password, and server address
    ```
    cat /var/vcap/jobs/route_emitter/config/route_emitter.json | jq . | grep nats
    ```
0. Download nats client
    ```
    curl -sf https://binaries.nats.dev/nats-io/natscli/nats@latest | sh
    mv nats /usr/local/bin/nats
    ```

0. Subscribe using certs
    ```
    nats sub "*.*" -s  nats://NATS_USERNAME:NATS_PASSWORD@NATS_ADDRESS \
      --tlsca=/var/vcap/jobs/route_emitter/config/certs/nats/ca.crt \
      --tlskey=/var/vcap/jobs/route_emitter/config/certs/nats/client.key \
      --tlscert=/var/vcap/jobs/route_emitter/config/certs/nats/client.crt
    ```

 The Route Emitter registers routes every 20 seconds (by default) so that the GoRouter (which subscribes to these messages) has the most up-to-date information about which IPs map to which apps and routes. Depending on how many routes there are, this might be a lot of information.

0. Make a new route for one of your apps and find the route registration message.

### Expected Result
Inspect NATs messages. Look at what route information is sent to the GoRouter.

## Resources
[NATS message bus repo](https://github.com/nats-io/gnatsd)
[NATS ruby gem repo](https://github.com/nats-io/ruby-nats)

