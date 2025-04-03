
### What?
Application Security Groups are a collection of egress (outbound) rules that specify the protocols, ports, and IP ranges where application containers can send traffic. Security groups define rules that *allow* traffic instead of omitting it, which means that the order of evaluation for security groups that apply to the same space, org, or deployment is unimportant. Application containers use these rules to filter and log outbound network traffic.

When applications begin staging, they need traffic rules permissive enough to allow them to pull resources from the network. After an application is running, the traffic rules can be more restrictive and secure. To distinguish between these two security requirements, administrators can define different security groups for *staging* containers versus *runtime* containers.

To provide granular control when securing a deployment, an administrator can also assign security groups to apply across a CF deployment, or to specific spaces or orgs within a deployment.

### How?
1. As admin view the list of security groups 
`cf security-groups`
1. View the security groups assigned to staging containers 
`cf staging-security-groups`
1. View the security groups assigned to running containers 
`cf running-security-groups`
1. View the specific rules of each group using 
`cf security-group $group_name`

#### Expected Result
If you've deployed a full Cloud Foundry on GCP you should have two security groups applied to both staging and running apps: `public_networks` and `dns`.

If you run `cf security-group public_networks` you'll see that it allows traffic on all ports with all protocols on all IPs, save four gaps that correspond to [private IPv4 address spaces](https://en.wikipedia.org/wiki/Private_network#Private_IPv4_address_spaces) and the [APIPA reserved range](https://www.pctechbytes.com/networking/fix-169-254-address-problem/). The `dns` security group allows access to any IP, but only on [port 53](https://en.wikipedia.org/wiki/Domain_Name_System#Protocol_transport). The combination of the two groups is that private IPs can be accessed only on port 53, and all other IPs have all ports open.

(If you're working with PCF Dev, you should see three security groups, one of which is named `all_pcfdev` and opens all egress traffic. Because of the `all_pcfdev` security group any other group would be redundant.)

### Resources
[Application Security Groups Documentation](https://docs.cloudfoundry.org/adminguide/app-sec-groups.html)
[Typical Application Security Groups](https://docs.cloudfoundry.org/adminguide/app-sec-groups.html#typical-groups)
["Taking Security to the Next Level—Application Security Groups" by Abby Kearns](https://tanzu.vmware.com/content/blog/taking-security-to-the-next-level-application-security-groups)
["Making sense of Cloud Foundry security group declarations" by Sadique Ali](https://sdqali.in/blog/2015/05/21/making-sense-of-cloud-foundry-security-group-declarations/)