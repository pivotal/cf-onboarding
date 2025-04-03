
### What?
[OAuth](http://searchmicroservices.techtarget.com/definition/OAuth) is an open standard for authorization, commonly used as a way for Internet users to authorize websites or applications to access their information on other websites but without giving them the passwords. Designed specifically to work with HTTP, OAuth essentially allows access tokens to be issued to third-party clients by an authorization server, with the approval of the resource owner. The third party then uses the access token to access the protected resources hosted by the resource server.

When you `cf login` via the command line, the cf CLI communicates with the [Cloud Foundry User Account and Authentication (UAA) Server](https://github.com/cloudfoundry/uaa) to exchange your credentials for a token. I was going to think of an explanation for why you'd want to see this token, but honestly it's just a good opportunity to think about OAuth.

### How?
1. Make sure you are logged in and run `cf oauth-token`
1. Visit **[jwt.io](https://jwt.io/)** and paste your token to view it's contents

### Expected Result
The [jwt.io](https://jwt.io/) site will say the token is invalid, but ignore that. You should now have a decoded header and payload!

Are the contents what you expected? What does the header mean? What do you think each of the keys does? (Hint: the links immediately below might help answer that).

### Resources
[Docs: A Note on UAA Tokens](https://github.com/cloudfoundry/uaa/blob/master/docs/UAA-Tokens.md)
[Docs: Validating an ID token](https://developers.google.com/identity/protocols/OpenIDConnect#validatinganidtoken)

### Relevant Repos and Teams
**UAA:** [cloudfoundry/uaa](https://github.com/cloudfoundry/uaa)
**CLI:** [cloudfoundry/cli](https://github.com/cloudfoundry/cli)