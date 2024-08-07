# CF Onboarding
## Onboarding Tracker: Local
Onboarding Tracker Local is a hands-on tour of some of the systems you'll be working on / with during your time as a CF contributor. It's good solo-able material, and good preparation for the paired [Onboarding Week](https://github.com/pivotal-cf/onboarding#onboarding-tracker-oss-aka-onboarding-week).

Whether you are Engineering, Product, Design, Docs, or one of dozens of other roles that make our engine run and you're curious about the essentials of the CF or Concourse user experience, then this backlog is for you.

## Onboarding Tracker: OSS (a.k.a. Onboarding Week)
Onboarding Tracker OSS is intended to be a _facilitated_ exploration of Open Source Cloud Foundry, embarked upon with other Cloud Foundry newbies. Unlike Tracker Local, it is more of a puzzle than a tour. It provides:

1. A self-paced learning environment paired with others who are learning too.
1. A coherent, if cursory, overview of a complicated product.
1. Empathy for the customer who uses that product.
1. The opportunity to struggle through these problems and really learn the material.
1. A little knowledge of the breadth of work 70+ teams are doing around the world.

To run an Onboarding Week in your office, **read the [facilitation](FACILITATING.md) docs** and join the **[#cf-onboarding-week](https://pivotal.slack.com/messages/cf-onboarding-week/)** channel on Pivotal Slack.

## Usage
### Import stories to Tracker (from source)
The stories in this repo are divided by epic (e.g. Deploying with GCP, Redis CUPS, etc.) They are provided in .[prolific](https://github.com/onsi/prolific) format. To grab the most recent versions of stories from master or another branch:

1. Clone this repo
1. Run `./build local` or `./build oss` (requires Docker)
1. Import your newly created csv file (`onboarding-tracker.csv`) to a new Tracker project
1. WARNING: concatenating CSVs is a risky/inadvisable business, so stories generated by this script are slightly buggy in inconsequential ways (e.g. the first letter or two of a single story title goes missing).

## Contributing
Depending on personal preferance, either edit stories in the .prolific text files themselves or convert the prolific file to a CSV and upload it to a temporary Tracker project, reverting back to .prolific format once you're ready to make a PR. The second one takes more time, but removes the risk of accidental prolific syntax errors.
