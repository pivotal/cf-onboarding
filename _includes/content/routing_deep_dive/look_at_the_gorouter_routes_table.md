
### What?
So the Route Emitter emits routes via the NATS message Bus. GoRouter _subscribes_ to those messages and keeps a route table that is uses to route network traffic bound for CF apps and CF components.

Let's take a look at that route table.

### Prerequisites

Go through the **Basic BOSH Knowledge** story to gain familiarity with BOSH, specifically BOSH CLI commands.

### How?
0. Bosh ssh onto the router vm and become root.
0. Install jq (a json manipulation and display tool)
 ```
 apt-get install jq
 ```
0. Get the username and password for the routing api
 ```
 head /var/vcap/jobs/gorouter/config/gorouter.yml
 ```
0. Get the routes table
 Use the following command to confirm the port used
 ```
  netstat -tunlp | grep gorouter
 ```
 then run:
 ```
 curl "http://USERNAME:PASSWORD@localhost:8082/routes" | jq .
 ```
0. Make a new route for one of your apps and find the route in the routes table.

 It should look something like this (I took out some extra bits):
    ```
    "proxy.meow.cloche.c2c.cf-app.com": [   <------ The name of the route!
        {
          "address": "10.0.1.12:61014",     <------ This is the Diego Cell IP and port where GoRouter will send traffic for this route
          "tls": true                       <------ This means the GoRouter will use send traffic to this app over TLS
        }
      ]
    ```

0. Check the environment variables for your app and see that the information matches the routing table
 ```
cf ssh MY-APP -c "env | grep CF_INSTANCE_ADDR"
 ```

### Expected Result
You are able to access the route table on the router VM.

## Resources
[GoRouter routing table docs](https://github.com/cloudfoundry/gorouter#the-routing-table)
