# SkillChain

A decentralized, blockchain-based job and skill verification platform built with Cosmos SDK.

## 🌟 Overview

SkillChain is a decentralized platform that enables users to create professional profiles, post jobs, apply for positions, and verify their skills through institutions or communities. The platform leverages blockchain technology to ensure transparency and immutability of professional credentials.

## ✅ Status Overview

| Feature                         | Status     |
|---------------------------------|------------|
| Profile Creation & Update       | ✅ Done     |
| Job Posting & Application       | ✅ Done     |
| Application Tracking            | ✅ Done     |
| Notification System             | ✅ Done     |
| Job Search & Filtering          | ✅ Done     |
| Skill Verification System       | 🔜 In Progress |
| Institutional Trust System      | 🔜 Planned  |
| Multisig Security              | 🔜 Planned  |
| Skill Tokenization             | 🔜 Planned  |
| DAO Governance                  | 🔜 Planned  |

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

### Create a Profile
```bash
skillchaind tx profile create-profile johndoe "Full-stack Dev" --from alice --chain-id skillchain --yes
```

### Post a Job
```bash
skillchaind tx job create-job "Blockchain Engineer" "5+ years with Go/Cosmos SDK" "5000stake" --from bob --chain-id skillchain --yes
```

### Apply for a Job
```bash
skillchaind tx job apply-job 1 "I'm excited to contribute!" --from alice --chain-id skillchain --yes
```

### Review an Application
```bash
skillchaind tx job review-application 1 cosmos1... "ACCEPTED" --from bob --chain-id skillchain --yes
```

### Check Notifications
```bash
skillchaind query job get-notifications --recipient cosmos1...
```

## 🛠️ Troubleshooting

Common issues and their solutions:

- ❌ `panic: can't find field id` → Missing field in proto definitions. Run `ignite generate proto-go`
- ❌ `error: cannot build app` → Ensure all required modules are installed (`go mod tidy`)
- ❌ `connection refused` → Check if the chain is running (`ignite chain serve`)
- ❌ `invalid account` → Make sure you have created an account (`skillchaind keys add mykey`)

## 🏗️ Project Structure

```
skillchain/
├── app/              # Application entry point
├── cmd/              # Command-line interface
├── proto/            # Protocol buffer definitions
├── x/                # Cosmos SDK modules
├── api/              # API definitions
├── docs/             # Documentation
├── testutil/         # Testing utilities
└── tools/            # Development tools
```

## 🧪 Testing

Run the test suite:
```bash
make test
```

## 📚 Documentation

- [Technical Documentation](docs/)
- [API Reference](api/)
- [Roadmap](roadmap.md)

## 🤝 Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📄 License

This project is licensed under the MIT License - see the LICENSE file for details.

## 🔗 Links

- [GitHub Repository](https://github.com/SkillChainLab/skillchain)
- [Documentation](https://docs.skillchain.io)
- [Discord Community](https://discord.gg/skillchain)

## 🙏 Acknowledgments

- Cosmos SDK team for the amazing framework
- All contributors and community members

---

Built with ❤️ by the SkillChain Team
