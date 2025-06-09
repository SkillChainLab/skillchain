# SkillChain Blockchain

A comprehensive blockchain platform for the freelance and gig economy, built with Cosmos SDK. SkillChain enables secure, transparent, and decentralized interactions between freelancers, clients, and service providers.

## 🌟 **Key Features**

### **Custom Modules (6 Modules)**
- **👤 Profile** - User profiles, skills, and endorsement system
- **💼 Marketplace** - Job postings, proposals, projects, and milestones  
- **📊 Analytics** - User activity tracking and platform metrics
- **📁 File Storage** - Decentralized file storage with IPFS integration
- **🔔 Notifications** - Real-time notification system
- **🏗️ SkillChain Core** - Token economics and governance

### **Native Token: USKILL**
- Pure uskill economy with staking support
- Gas fees paid in uskill
- Marketplace payments in uskill
- Deflationary token mechanics with burn functionality
- Powered by Cosmos SDK Bank Module

## 🚀 **Quick Start**

### **Prerequisites**
- Go 1.21+
- [Ignite CLI](https://docs.ignite.com/welcome/install)

### **Installation & Setup**

```bash
# Clone the repository
git clone https://github.com/SkillChainLab/skillchain.git
cd skillchain

# Start the blockchain
ignite chain serve

# Alternative: Start with reset
ignite chain serve --reset-once
```

### **Network Information**
```
🌍 Tendermint RPC: http://localhost:26657
🌍 Blockchain API: http://localhost:1317  
🌍 Token Faucet: http://localhost:4500
```

### **Default Accounts**
```
👤 Alice: skill1szn9xfkglwpjqztscspcyvsn7et7xcjaydp3kv
👤 Bob: skill1p8g303aeh0f3cpj8z55rava0t0sk95pqmhs4vh
```

## 📋 **Project Structure**

```
skillchain/
├── app/                    # Cosmos SDK app configuration
├── x/                      # Custom modules (6 modules)
│   ├── profile/           # User profiles and skills management
│   ├── marketplace/       # Complete marketplace functionality
│   ├── analytics/         # User activity and metrics tracking
│   ├── filestorage/       # Decentralized file management
│   ├── notifications/     # Notification system
│   └── skillchain/        # Core blockchain module
├── proto/                 # Protocol buffer definitions
├── config.yml            # Blockchain configuration with uskill economy
└── cmd/                   # CLI commands
```

## 🔧 **Module Overview**

### **Profile Module**
- Create and manage user profiles with comprehensive information
- Add skills with proficiency levels and experience years
- Skill endorsement system with reputation tracking
- Profile verification and social links integration

### **Marketplace Module**  
- Post job opportunities with detailed requirements
- Submit and manage proposals from freelancers
- Create and track projects with client-freelancer relationships
- Milestone-based payment system with escrow functionality
- Project completion and dispute resolution
- Payment release mechanisms

### **Analytics Module**
- Track comprehensive user activities across platform
- Generate platform usage metrics and statistics
- Revenue tracking and financial analytics
- User behavior analysis with privacy protection
- Session tracking and device information

### **File Storage Module**
- IPFS integration for decentralized file storage
- File permissions and access control management
- Metadata tracking and version control
- Support for multiple file types and formats
- Privacy-compliant file sharing

### **Notifications Module**
- Real-time notification system for platform events
- Priority levels (high, medium, low) for notifications
- Source module tracking for notification origins
- User notification preferences and settings
- Cross-module notification support

### **SkillChain Core Module**
- Token economics and governance functionality
- Burn mechanisms for deflationary tokenomics
- Platform parameter management
- Core blockchain utilities and helper functions

## 💻 **CLI Usage Examples**

### **Profile Management**
```bash
# Create a user profile
skillchaind tx profile create-profile [index] [address] [name] [title] [bio] [linkedin] [github] --from [user]

# Add a skill
skillchaind tx profile create-user-skill [index] [owner] [skill] [level] [years] [verified] [verifier] [date] [endorsements] --from [user]

# Endorse a skill
skillchaind tx profile endorse-skill [skill-id] [endorser] [endorsement-level] --from [user]
```

### **Marketplace Operations**
```bash
# Create a job posting
skillchaind tx marketplace create-job-posting [index] [client] [title] [description] [skills] [budget] [currency] [deadline] [active] [created] --from [user]

# Submit a proposal  
skillchaind tx marketplace create-proposal [index] [job-id] [freelancer] [amount] [currency] [timeline] [cover-letter] [reputation] [accepted] [created] --from [user]

# Accept a proposal
skillchaind tx marketplace accept-proposal [proposal-id] --from [client]

# Create project milestones
skillchaind tx marketplace create-milestone [index] [project-id] [title] [description] [amount] [due-date] [status] [completed] [paid] [submitted] [approved] --from [user]
```

### **File Storage**
```bash
# Store a file
skillchaind tx filestorage create-file-record [index] [owner] [filename] [hash] [size] [content-type] [upload-date] [ipfs-hash] [metadata] [public] --from [user]

# Grant file permissions
skillchaind tx filestorage create-file-permission [index] [file-id] [grantee] [permissions] [granted-by] [granted-at] [expires-at] --from [owner]
```

### **Analytics Tracking**
```bash
# Track user activity
skillchaind tx analytics create-user-activity [index] [user] [type] [action] [resource] [timestamp] [ip] [user-agent] [metadata] --from [user]

# Record platform metrics
skillchaind tx analytics create-platform-metric [index] [metric-name] [value] [category] [timestamp] [metadata] --from [admin]
```

### **Notifications**
```bash
# Create notification
skillchaind tx notifications create-notification [index] [user] [type] [title] [message] [data] [read] [created] [priority] [source-module] [source-action] --from [sender]
```

## 🔍 **Query Examples**

```bash
# Check profile
skillchaind query profile show-user-profile [profile-id]

# View job posting
skillchaind query marketplace show-job-posting [job-id]

# Check token balances (using Cosmos SDK bank module)
skillchaind query bank balances [address]

# View analytics
skillchaind query analytics show-user-activity [activity-id]

# Check file storage
skillchaind query filestorage show-file-record [file-id]

# View notifications
skillchaind query notifications show-notification [notification-id]
```

## 🌐 **API Endpoints**

The blockchain exposes REST API endpoints at `http://localhost:1317`:

### **Core Cosmos SDK**
- `/cosmos/bank/v1beta1/balances/{address}` - Check token balances
- `/cosmos/staking/v1beta1/validators` - Validator information

### **SkillChain Custom Modules**
- `/skillchain/profile/user_profile` - User profiles
- `/skillchain/marketplace/job_posting` - Job postings  
- `/skillchain/marketplace/proposal` - Proposals
- `/skillchain/marketplace/project` - Projects
- `/skillchain/analytics/user_activity` - User activities
- `/skillchain/filestorage/file_record` - File records
- `/skillchain/notifications/notification` - Notifications

## 🔐 **Security Features**

- **Proof of Stake** consensus with uskill token
- **Multi-signature** transaction support
- **Role-based** access control in file permissions
- **Privacy-compliant** IP anonymization in analytics
- **Comprehensive** audit trails for all transactions
- **Encrypted** metadata storage

## 🧪 **Testing**

```bash
# Run module tests
go test ./x/profile/...
go test ./x/marketplace/...
go test ./x/analytics/...
go test ./x/filestorage/...
go test ./x/notifications/...

# Integration tests
make test

# Start test network
ignite chain serve --reset-once
```

## 📊 **Proven Functionality**

### **✅ Fully Tested Marketplace Workflow**
- Profile creation for employers and freelancers
- Job posting with detailed requirements
- Proposal submission and acceptance
- Project creation with milestone tracking
- File sharing and document management
- Payment processing via native bank module
- Comprehensive activity analytics
- Real-time notification system

### **💰 Token Economy**
- **117M uskill** - Alice balance (after transactions)
- **113M uskill** - Bob balance (after receiving payments)
- Gas fees automatically paid in uskill
- Successful token transfers tested

## 📈 **Roadmap**

### **Phase 1: Core Platform** ✅
- [x] 6 custom modules implemented
- [x] Profile management with skills and endorsements
- [x] Complete marketplace functionality
- [x] File storage with IPFS integration
- [x] Analytics and activity tracking
- [x] Notification system
- [x] Full marketplace workflow tested

### **Phase 2: Advanced Features** 🚧
- [ ] Multi-chain interoperability
- [ ] Advanced smart contracts
- [ ] Mobile applications
- [ ] Enhanced governance features
- [ ] Advanced dispute resolution

### **Phase 3: Ecosystem** 🔮
- [ ] DeFi integrations
- [ ] NFT marketplace for digital assets
- [ ] Social features and networking
- [ ] Global expansion and partnerships

## 🤝 **Contributing**

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📄 **License**

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🔗 **Links**

- **Documentation**: [docs/](./docs/)
- **API Reference**: http://localhost:1317/swagger/
- **Community**: [Discussions](https://github.com/SkillChainLab/skillchain/discussions)

## 💬 **Support**

For support and questions:
- Open an [issue](https://github.com/SkillChainLab/skillchain/issues)
- Join our community discussions
- Check the documentation

---

**Built with ❤️ using Cosmos SDK and Ignite CLI**