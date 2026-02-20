# Frontend Masters: Build Go Apps That Scale on AWS

- **Teacher**: Melkey
- **Official Repo**: [https://github.com/Melkeydev/FrontEndMasters](https://github.com/Melkeydev/FrontEndMasters)

---

## Setup

Dependencies installed in the devcontainer:
- [Go](https://go.dev/dl/)
- [AWS CLI](https://docs.aws.amazon.com/cli/latest/userguide/getting-started-install.html)
- [AWS CDK](https://docs.aws.amazon.com/cdk/v2/guide/home.html)

Prepare the environment:
1) Create an [AWS Account](https://portal.aws.amazon.com/billing/signup) if you don't already have one.
2) Configure an AWS user permissions:
    - [Create an IAM or IAM Identity Center administrative account](https://docs.aws.amazon.com/cli/latest/userguide/getting-started-prereqs.html#getting-started-prereqs-iam)
3) Configure the AWS CLI:
    - Configure [short or long term credentials](https://docs.aws.amazon.com/cli/latest/userguide/getting-started-quickstart.html)
    - Confirm your setup with: `aws s3 ls` and `aws sts get-caller-identity`
4) Bootstrap your CDK environment with:
```bash
cdk bootstrap aws://ACCOUNT-NUMBER/REGION
```

Here's [more info](https://docs.aws.amazon.com/cdk/v2/guide/home.html) about boostrapping the CDK.
Your `ACCOUNT-NUMBER` and `REGION` can be found in the AWS console or using the following AWS CLI commands:

Get your `ACCOUNT-NUMBER` using the following AWS CLI command:
```bash
aws sts get-caller-identity --query Account --output text
```

Get your `REGION` using the following AWS CLI command aws configure get region
```bash
aws configure get region
```

## CDK

The `cdk.json` file tells the CDK toolkit how to execute your app.

## Useful commands

 * `cdk deploy`      deploy this stack to your default AWS account/region
 * `cdk diff`        compare deployed stack with current state
 * `cdk synth`       emits the synthesized CloudFormation template
 * `go test`         run unit tests

## Differences between the Original Course and this Version

### Project Structure

In the course there was no emphasis on the project structure,
Melkey soley focused on demonstrating the technologies.  

### Architecture

`amd64` is the default, but `arm64` is cheaper and faster for most Lambda workloads.
