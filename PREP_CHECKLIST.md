## Instructions
- Start a [new issue](https://github.com/pivotal/cf-onboarding/issues/new) in the onboarding repo. The title should include the date and location of the Onboarding week (e.g. "SF Onboarding Week, July 16 2018").
- Copy the checklist below into the issue.
- Create the issue
- Use the checkboxes to track your work as you get set up.
- When you finish with the checklist, make sure to close the issue.

## Checklist

### Communication
- [ ] Several weeks before: send out an email to the participants reminding them that they are signed up and to clear their calendars.  You can find a template in [FACILITATING.md](https://github.com/pivotal/cf-onboarding/blob/master/FACILITATING.md).
  - [ ] You will be communicating with these folks a lot. I find it helpful to save a copy/pasteable list of their names <emails> here:
    - `name <email>.name <email>`
- [ ] The week before: send out an email to each of the participants. You can find a template in [FACILITATING.md](https://github.com/pivotal/cf-onboarding/blob/master/FACILITATING.md).
- [ ] Get a [Zoom](https://zoom.us/) room setup for remote participants. {<can save link here>}  (Make sure to include this on your calendar invites.)

### Scheduling
- [ ] Schedule a welcome/orientation session (1hr at 9:30 Monday morning). Discuss the goals of the week and prime the participants to consider what they'd like to get out of Onboarding week.
- [ ] Schedule daily standup (15min at 9:15, starting Tuesday).
- [ ] Schedule afternoon check in (30min at 1:30).
- [ ] Schedule retro at the end of the week (1hr at 4:00), include link to retro board: https://retros.cfapps.io/retros/onboarding-sf (pw: onboarding-sf (pls don't change it))
- [ ] Schedule any sessions that you know the group would be interested in -- for example, BOSH boxes and lines.
- [ ] In the past, non Engineers have appreciated it when we blocked out the entire days of onbaording. This is especially helpful if done a few weeks in advance; so, they have time to move meetings and avoid having new ones scheduled in the first place.

### IaaS Setup
- [ ] Make sure you have [GCP projects](https://console.cloud.google.com) set up.
  - [ ] In SF, we use:
    - [cf-sf-onboarding-env-1](https://console.cloud.google.com/billing/0076DC-766E1F-EBDCB8/reports?project=cf-sf-onboarding-env-1)
    - [cf-sf-onboarding-env-2](https://console.cloud.google.com/billing/0076DC-766E1F-EBDCB8/reports?project=cf-sf-onboarding-env-2)
    - [cf-sf-onboarding-env-3](https://console.cloud.google.com/billing/0076DC-766E1F-EBDCB8/reports?project=cf-sf-onboarding-env-3)
  - [ ] Otherwise, Fill out [this form to create new environments](https://docs.google.com/forms/d/e/1FAIpQLSeJ31997Zma1WtLcCtswiysCFWOG5MXNmlYCpJsiYgdG9kKnA/viewform).
- [ ] Clean up any VMs, networks, and DNS records that may have been leftover from the previous week
- [ ] Add all [participants](https://docs.google.com/spreadsheets/d/1eW23iJRD56CE859o0j6ArSaGtLfP0JlHXSCdXFhbCdI/edit?usp=sharing) as **owners** to each of the GCP projects.
  - [ ] In SF:
    - [IAM for env-1](https://console.cloud.google.com/iam-admin/iam?project=cf-sf-onboarding-env-1)
    - [IAM for env-2](https://console.cloud.google.com/iam-admin/iam?project=cf-sf-onboarding-env-2)
    - [IAM for env-3](https://console.cloud.google.com/iam-admin/iam?project=cf-sf-onboarding-env-3)
- [ ] If your CF on-boarding will include the GCP track (`gcp` tracker label / the `oss` build) then ensure participants have access to the `CF-Onboarding-dns` project in GCP (for DNS propagation). [IAM for DNS](https://console.cloud.google.com/iam-admin/iam?project=cf-onboarding-dns)

### Tracker Setup
- [ ] Make sure you have Tracker projects prepared.
  - [ ] In SF, we use the following trackers projects:
    - [Pair 1](https://www.pivotaltracker.com/n/projects/2125981)
    - [Pair 2](https://www.pivotaltracker.com/n/projects/2125982)
    - [Pair 3](https://www.pivotaltracker.com/n/projects/2089066)
- [ ] Add all participants as members to each of the tracker projects. (A list of emails works here.)
- [ ] Delete everything from the tracker projects.
- [ ] Build the tracker back log csv
  - [ ] Run `./build oss` from the onboarding project, this will generate `onboarding-tracker.csv`
       - Note: you need docker installed `brew cask install docker`
         - To get the command line tool to show up you may have to first start docker via Spotlight Search
         - If docker gives you this error: `failed to dial gRPC: unable to upgrade to h2c, received 502 context canceled`
           - In docker settings: enable experimental, restart docker, disable experiment and restart again.
- [ ] Import the CSV into each Tracker project (Taskbar > More > Import CSV > Choose File)

### LastPass Setup
- [ ] Give all participants access to the shared LastPass folder:
  - LastPass > Sharing Center > Shared-CF SF Onboarding > Manage > Invite Users (A list of emails works here.)

### Workstation Setup
- [ ] Make sure you have a workstation for each pair.
- [ ] Ensure ethernet cable is plugged into iMac, not monitor.
- [ ] Re-image each machine. Here are the steps:
  - [ ] Run `Startup Disk`.
  - [ ] Unlock with the workstation password (you may need to ask the previous facilitator for the machine's password).
  - [ ] Select the network disk and click restart. It will likely be the rightmost start disk, and will definitely include the phrase "Network Disk." If you don't see that, try restarting the machine and try again.
  - [ ] Select deployment option 2, "macOS Sierra" and click the play button. If you look in between the parens of the option name, you'll see the default password for the image. A terminal will open and you’ll see a lot of output, including a bunch of lines that say `installer: PHASE: Running package scripts`. Don’t worry, this is fine. The boot script will search for updates -- you’ll see output about OS X, iTunes, and Safari updates, for example -- and the machine will probably restart a few times as it downloads and installs the updates.
  - [ ] Log into machine with the default password.
  - [ ] Open the security preferences and change the password to something sane. Write this down so you can give it to the participants later.
  - [ ] Open a terminal window and run
     ```
     curl -LOk https://github.com/pivotal-cf-experimental/onboarding-week-construct/archive/master.zip
     unzip master.zip
     cd onboarding-week-construct-master
     ./construct
     ```
     You'll need to babysit this process, as it will prompt you for the password several times.

### Misc

- [ ] Record any issues you have with this process so we can update this checklist.
