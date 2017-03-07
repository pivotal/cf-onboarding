# CF Onboarding Week

Onboarding Week is a facilitated exploration of the Cloud Foundry product, embarked upon with other Cloud Foundry newbies. It provides:

1. A self-paced learning environment paired with others who are learning too.
1. A coherent, if cursory, overview of a complicated product.
1. Empathy for the customer who uses that product.
1. A little knowledge of the breadth of work 70+ teams are doing around the world.

To run an Onboarding Week in your office, **read the [facilitation](FACILITATING.md) docs** and join the **[#cf-onboarding-week](https://pivotal.slack.com/messages/cf-onboarding-week/)** channel on Pivotal Slack.

## Usage
### Import stories to Tracker (from release)
1. Download latest release from the releases page
1. Use the Import CSV function to add the stories to your project with the CSV file. The account used to import the CSV file will be the requester associated with the resulting stories.

### Import stories to Tracker (from source)
The stories in this repo are divided by epic (e.g. Deploying with GCP, Redis CUPS, etc.) They are provided in .prolific format. To grab the most recent versions of stories from master or another branch:

1. Clone this repo
1. Install the **[prolific tool](https://github.com/dgodd/prolific)** (the version linked here, not the original)
1. Run `prolific file_name.prolific > file_name.csv` on each epic you would like to import to Tracker
1. Import your newly created csv file(s) to your Tracker project

## Contributing
Depending on personal preferance, either edit stories in the .prolific text files themselves or convert the prolific file to a CSV and upload it to a temporary Tracker project, reverting back to .prolific format once you're ready to make a PR. The second one takes more time, but removes the risk of accidental prolific syntax errors.
