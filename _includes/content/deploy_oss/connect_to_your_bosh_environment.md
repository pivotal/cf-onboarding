
### What?
Time to target the Director VM you just `bbl up`ed.

### How?
To set the BOSH environment, run the following commands:

`bbl print-env` to see the environment variables you'll set to use your new BOSH.

`eval "$(bbl print-env)"` to export them. **Note: you will definitely need to include the double-quotes for this command to work properly.** Once you've run it, you should be able to start working with your bosh director. If you ever have a hard time connecting to your BOSH director, your first step should be to re-run `eval "$(bbl print-env)"`.

### Expected Results
When you run `bosh env`, it should print something like:

![image](https://www.pivotaltracker.com/file_attachments/77821681/download)

Annnd that's it!

# Congratulations, you've successfully deployed BOSH! #bigpatontheback