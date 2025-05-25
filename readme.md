# SkillChain

SkillChain is a decentralized platform built on Cosmos SDK that connects job seekers with employers through a transparent and efficient blockchain-based system.

## Features

### Job Management
- Create and manage job postings
- Submit and track job applications
- Real-time application status updates
- Job search and filtering capabilities

### Notification System
- Application status notifications
- Job update notifications
- Real-time alerts for important events

### Skill Verification System
- Verified Institution Management
  - Create and manage verified institutions
  - Set verification categories and levels
  - Track institution verification status

- Skill Verification Process
  - Submit verification requests
  - Support for multiple skills per request
  - Evidence submission and tracking
  - Verification status management

## Getting Started

### Prerequisites
- Go 1.20 or later
- Ignite CLI
- Cosmos SDK v0.50.13

### Installation

1. Clone the repository:
```bash
git clone https://github.com/SkillChainLab/skillchain.git
cd skillchain
```

2. Install dependencies:
```bash
go mod tidy
```

3. Build the chain:
```bash
ignite chain build
```

4. Start the chain:
```bash
ignite chain serve
```

### Testing the Verification System

1. Create a verified institution:
```bash
skillchaind tx verification create-verified-institution <address> <name> <website> <categories> <level> --from <account> --chain-id skillchain --yes
```

2. Create a verification request:
```bash
skillchaind tx verification create-verification-request <institution_address> <skills> <evidence> --from <account> --chain-id skillchain --yes
```

3. Approve a verification request:
```bash
skillchaind tx verification approve-verification-request <request_id> --from <institution_account> --chain-id skillchain --yes
```

4. Query verification requests:
```bash
skillchaind query verification verification-request <request_id>
```

## Contributing

We welcome contributions to SkillChain! Please see our [Contributing Guidelines](CONTRIBUTING.md) for more information.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Contact

For any questions or support, please open an issue in the GitHub repository.

## 🌟 Overview

SkillChain is a decentralized platform that enables users to create professional profiles, post jobs, apply for positions, and verify their skills through institutions or communities. The platform leverages blockchain technology to ensure transparency and immutability of professional credentials.

## ✅ Status Overview

| Feature                         | Status     | Notes |
|---------------------------------|------------|-------|
| Profile Creation & Update       | ✅ Done     | Tested and working |
| Job Posting & Application       | ✅ Done     | Tested and working |
| Application Tracking            | ✅ Done     | Tested and working |
| Notification System             | ✅ Done     | Notifications for applications and job updates working |
| Job Search & Filtering          | ✅ Done     | Basic search working |
| Skill Verification System       | ✅ Done     | Verified institutions, multi-skill verification, and approval process implemented |
| Institutional Trust System      | ✅ Done     | Institution verification levels and categories implemented |
| Multisig Security              | 🔜 Planned  | Design phase |
| Skill Tokenization             | 🔜 Planned  | Design phase |
| DAO Governance                  | 🔜 Planned  | Design phase |

## 🧩 Modules

- **x/profile** – Manage user profiles, bios, usernames, and professional information
- **x/job** – Job listings, applications, reviews, hiring management, and notifications
- **x/verification** *(coming soon)* – Skill endorsement & validation framework
- **x/governance** *(planned)* – DAO-based community governance
- **x/token** *(planned)* – Skill tokenization and NFT integration

## 🚀 Features

### Profile Module
- Create and manage professional profiles
- Add skills, experiences, and social links
- Update and delete profile information
- List and query profiles

### Job Module
- Create job postings
- List available jobs
- View job details
- Apply for jobs
- Track applications
- Manage job applications
- Time-based job posting expiration
- Search and filter jobs
- Notification system for applications and updates

### Notification System
- Real-time notifications for job applications
- Application status updates
- Job posting updates
- Notification management (mark as read, delete)

### Skill Verification System (Coming Soon)
- Add and request skill verification
- Verification via trusted institutions
- Certificate uploads and references
- Community endorsements (like StackOverflow)
- On-chain verification logs
- Multisig security for verifiers
- Skill tokenization (NFT/Soulbound tokens)

### Governance Features (Planned)
- DAO-based community governance
- Validator selection through voting
- Proposal submission and voting
- Community-driven feature updates

📌 See the [Skill Verification Roadmap](docs/roadmap.md#profil-dogrulama-sistemi)

## 🛠️ Technology Stack

- **Blockchain Framework**: Cosmos SDK
- **Language**: Go
- **Protocol**: Protocol Buffers (proto3)
- **Build Tools**: Make, Buf

## 📦 Installation

1. Clone the repository:
```bash
git clone https://github.com/SkillChainLab/skillchain.git
cd skillchain
```

2. Install dependencies:
```bash
# Using make
make install

# If you don't have 'make':
go install github.com/cosmos/relayer/v2@latest
ignite chain build
```

3. Build the project:
```bash
# Using make
make build

# If you don't have 'make':
ignite chain build
```

## 📦 Usage

### Start Test Network
```bash
# Start a local test network
ignite chain serve
```
This command will:
- Initialize a local blockchain
- Start the chain in development mode
- Enable auto-reloading on code changes
- Provide a local REST API endpoint
- Create test accounts (alice, bob, testuser)

### Create a Profile
```bash
skillchaind tx profile create-profile [username] [bio] [email] [wallet-address] --from [account] --chain-id skillchain --yes
```

### Post a Job
```bash
skillchaind tx job create-job [title] [description] [budget] --from [account] --chain-id skillchain --yes
```

### Apply for a Job
```bash
skillchaind tx job apply-job [job-id] [cover-letter] --from [account] --chain-id skillchain --yes
```

### List Job Applications
```bash
skillchaind query job list-job-applications [job-id]
```

### Check Notifications
```bash
skillchaind query job get-notifications --recipient [address]
```

## 🛠️ Troubleshooting

Common issues and their solutions:

- ❌ `panic: can't find field id` → Missing field in proto definitions. Run `ignite generate proto-go`
- ❌ `error: cannot build app` → Ensure all required modules are installed (`go mod tidy`)
- ❌ `connection refused` → Check if the chain is running (`ignite chain serve`)
- ❌ `invalid account` → Make sure you have created an account (`skillchaind keys add mykey`)
- ❌ `account not found` → Ensure the account is added to genesis (`skillchaind genesis add-genesis-account`)
- ❌ `transaction failed` → Check if you have enough tokens for gas fees

## 🏗️ Project Structure

```