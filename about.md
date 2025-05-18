SkillChain
SkillChain is a decentralized blockchain application built with the Cosmos SDK and powered by Ignite CLI, designed to facilitate a secure and trustless job marketplace. The platform allows users to create professional profiles, post jobs, and apply to jobs—all verifiable and immutable on-chain.

🧭 Project Scope
The main goal of SkillChain is to:

Enable professionals to build verifiable public profiles.

Allow users to post jobs with budgets and descriptions.

Allow other users to apply to posted jobs with cover letters.

Provide a fully on-chain application history and job management system.

Ensure transparency, security, and data integrity through decentralized architecture.

✅ Features (So far)
🔐 Profile Module
create-profile: Create a professional profile with username and bio.

list-profile: List all created profiles.

show-profile: View profile details by username.

💼 Job Module
create-job: Post a new job with a title, description, and budget.

list-job: List all available jobs.

show-job: View a job post by ID.

📨 Application Module
apply-job: Apply to a job with a coverLetter and reference to a job ID.

list-application: List all applications for a specific job ID.

list-my-applications: Show all applications submitted by a specific user address.

🏛️ Software Architecture
Tech Stack
Layer	Technology
Blockchain Core	Cosmos SDK (v0.50.13)
Development Tool	Ignite CLI
Programming Lang.	Golang (go 1.21+)
Storage	IAVL (via Cosmos SDK)

Modules Structure
x/profile → Handles user profiles.

x/job → Manages job creation, listing, and viewing.

x/job/applications → Submodule to manage job applications per user.

Data Storage
Keys are prefixed (ProfileKeyPrefix, JobKeyPrefix, etc.)

Indexed by usernames or job IDs.

Applications use a compound key (jobID + applicant address).

📌 Future Milestones
Feature	Status
Profile Editing / Deletion	🔜 Planned
Job Editing / Deletion	🔜 Planned
Application Withdrawal	🔜 Planned
On-chain Review System	🔜 Planned
Token Payment Escrow System	🧠 Under Design
REST + gRPC API Gateway	🔜 Planned
Frontend Web App (Next.js)	🔜 Planned
Wallet Integration (Keplr, etc)	🔜 Planned

⚙️ How to Run Locally
bash
Copy
Edit
# Clone the repo
git clone https://github.com/SkillChainLab/skillchain.git
cd skillchain

# Build and start the chain
ignite chain serve
👥 Contributors
@serdarkayaci – Core Developer & Blockchain Architect

SkillChainLab Team – Protocol Design & Implementation

📄 License
MIT License. Free to use, fork, contribute, and build upon.