
### What?
Sometimes a user says "container to container networking is broken! AppA can't talk to AppB". After making sure that they have container to container networking (c2c) policies set, the next thing you might do is use tcpdump to inspect the packets.

Tcpdump is a CLI tool that allows you to inspect all of the traffic (not just TCP, despite the name) flowing through your container. This is a great tool for debugging how far packets are making it, determining how long the packets are taking to reach their destination, and figuring out where in the TCP handshake connections are failing.

In this story you are going to look at the packets being sent from boots to swiper. Then you'll watch the packets being sent in response.

### Prerequisites

Go through the **Basic BOSH Knowledge** story to gain familiarity with BOSH, specifically BOSH CLI commands.

### How?
**Setup**
1. You should have two apps pushed named boots and swiper from the previous story ("Allowing network traffic to other apps").
1. There should be no c2c network policies. Remove any policies that are left over from the previous story. (`cf network-policies` and `cf remove-network-policy --help`).

**Curl swiper from boots**
1. Get the overlay IPs of boots and swiper.
 ```
cf ssh boots -c "env | grep CF_INSTANCE_INTERNAL_IP"
cf ssh swiper -c "env | grep CF_INSTANCE_INTERNAL_IP"
 ```
1. Ssh onto boots.
 ```
cf ssh boots
 ```
1. Continually try to curl swiper from boots
 ```
watch -n 15 curl SWIPER_OVERLAY_IP:8080
 ```

**Look at those packets**
1. In another terminal, ssh onto the Diego Cell where boots is running and become root (see the help section if you don't know how to do this).
1. Run `tcpdump`.
    Ahhhhh too much information! ctrl+c! ctrl+c!  On a Diego Cell there are many packets being sent around, and tcpdump gives information about ALL OF THEM. You need to figure out a way to filter this overwhelming stream of information.
1.  Filter by packets where the source IP is BOOTS_OVERLAY_IP and where the destination IP is SWIPER_OVERLAY_IP over any interface. (Interested in learning more about network interfaces? It's a little too much to go into here. Contact Amelia Downs about CF Networking Onboarding week! Or google it.)
 ```
tcpdump -n src BOOTS_OVERLAY_IP and dst SWIPER_OVERLAY_IP -i any
 ```
 Hey! Those are your packets!

 Record the packets you see here from one curl.
 ```
# PUT TCPDUMP OUTPUT HERE
 ```

If swiper was successfully responding, then you should also see packets being sent as a response in the opposite direction.

1. See that no packets are being sent back from swiper to boots
  ```
tcpdump -n src SWIPER_OVERLAY_IP and dst BOOTS_OVERLAY_IP -i any
 ```

**Add c2c policy**
1. Add c2c policy to allow traffic from boots to swiper (`cf add-network-policy --help`)
1. In your original terminal, you should still be curling swiper from soots (if not go back to "curl swiper from boots" section).

**Inspect packets**
1. Look for packets from boots to swiper
 ```
tcpdump -n src BOOTS_OVERLAY_IP and dst SWIPER_OVERLAY_IP -i any
 ```
 Record the packets you see here from one curl.

 ```
 PUT TCPDUMP OUTPUT HERE
 ```
How are these packets different from before?

1. Look for packets from swiper to boots
 ```
tcpdump -n src SWIPER_OVERLAY_IP and dst BOOTS_OVERLAY_IP -i any
 ```

### Expected Result
You should see packets being sent in response from swiper to boots. You should see `200 OK`.

### Help
How to determine what Diego Cell your app is running on:
```
# First determine the IP of the Diego Cell
cf ssh APP_NAME  -c "env | grep CF_INSTANCE_IP"

# Look at bosh output to see which Diego Cell has that IP
bosh instances
```

### Resources
[tcpdump man page](https://www.tcpdump.org/manpages/tcpdump.1.html)
[helpful tcpdump commands](https://www.rationallyparanoid.com/articles/tcpdump.html)
