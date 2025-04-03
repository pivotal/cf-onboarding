
### Bug description

A malicious hacker has cracked into an unsecured bot account with read / write privileges and arbitrarily removed / altered bits of code in the Log Store project in an attempt to have us send broken code through the pipeline and into production! Yikes! Jokes on him or her, though, because we have tests in place that stop code like this from being released :D However, somehow they have locked us out of our commit history so you'll have to use the tests to identify the issues and fix them. Good thing we wrote these tests!

**Note:** the hacker and the damage done is fictitious; we have not actually been hacked.

### Link to branch with code at time of bug

https://github.com/pivotal/log-store-release/tree/broken

### Acceptance criteria / desired outcome

All the tests now pass.
