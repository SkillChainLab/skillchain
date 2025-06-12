# SkillChain RPC API Documentation

## Overview

SkillChain provides a comprehensive REST API for interacting with all blockchain modules. The API is available at `http://localhost:1317/skillchain/v1/` when the blockchain is running.

## Base URL

**Versioned (Recommended):**
```
http://localhost:1317/skillchain/v1/
```

**Legacy (Backward Compatibility):**
```
http://localhost:1317/skillchain/
```

## Authentication

Most endpoints require transaction signing with a valid SkillChain address. For query endpoints, no authentication is required.

## Response Format

All API responses follow this standard format:

```json
{
  "data": { ... },
  "message": "Success message",
  "timestamp": "2024-01-01T12:00:00Z"
}
```

## Error Handling

Error responses include HTTP status codes and descriptive messages:

```json
{
  "error": "Error description",
  "code": 400,
  "details": "Additional error details"
}
```

---

## SkillChain Core Module

### Module Parameters

#### Get Module Parameters
- **GET** `/params`
- **Description**: Get SkillChain module parameters
- **Response**:
```json
{
  "params": {
    "vusd_enabled": true,
    "vusd_mock_price": "0.50",
    "min_collateral_ratio": "1.50",
    "price_update_authority": "skill1..."
  }
}
```

### Token Information

#### Get Token Info
- **GET** `/token/info`
- **Description**: Get comprehensive token information including supply, burn status
- **Response**:
```json
{
  "name": "SkillChain",
  "symbol": "SKILL",
  "decimals": 6,
  "description": "SkillChain native token for professional networking",
  "total_supply": "1000000000uskill",
  "circulating_supply": "950000000uskill",
  "burned_amount": "50000000uskill",
  "max_supply": "1000000000uskill",
  "burn_enabled": true,
  "chain_description": "Decentralized professional networking blockchain",
  "website_url": "https://skillchain.io"
}
```

### VUSD Operations

#### Get VUSD Treasury Status
- **GET** `/vusd/treasury`
- **Description**: Get current VUSD treasury information and collateral ratios
- **Response**:
```json
{
  "skill_balance": "1500000uskill",
  "vusd_supply": "800000uvusd",
  "exchange_rate": "0.50"
}
```

#### Get User VUSD Position
- **GET** `/vusd/position/{address}`
- **Description**: Get user's VUSD position and collateral details
- **Parameters**: 
  - `address`: User's SkillChain address
- **Response**:
```json
{
  "position": {
    "address": "skill1abc...",
    "vusd_balance": "1000uvusd",
    "skill_collateral": "1500uskill"
  },
  "vusd_balance": "1000uvusd",
  "skill_collateral": "1500uskill",
  "health_factor": "1.5",
  "exists": true
}
```

#### Convert SKILL to VUSD
- **POST** `/vusd/convert/skill-to-vusd`
- **Description**: Prepare transaction to convert SKILL tokens to VUSD
- **Body**:
```json
{
  "creator": "skill1abc...",
  "amount": "1000000uskill"
}
```
- **Response**:
```json
{
  "message": "SKILL to VUSD conversion transaction prepared",
  "tx_info": {
    "creator": "skill1abc...",
    "amount": "1000000uskill",
    "estimated_gas": "180000",
    "note": "Use skillchaind tx skillchain convert-skill-to-vusd to execute"
  }
}
```

#### Convert VUSD to SKILL
- **POST** `/vusd/convert/vusd-to-skill`
- **Description**: Prepare transaction to convert VUSD back to SKILL tokens
- **Body**:
```json
{
  "creator": "skill1abc...",
  "amount": "500000uvusd"
}
```
- **Response**:
```json
{
  "message": "VUSD to SKILL conversion transaction prepared",
  "tx_info": {
    "creator": "skill1abc...",
    "amount": "500000uvusd",
    "estimated_gas": "180000",
    "note": "Use skillchaind tx skillchain convert-vusd-to-skill to execute"
  }
}
```

### Token Operations

#### Burn Tokens
- **POST** `/tokens/burn`
- **Description**: Prepare transaction to burn SKILL tokens
- **Body**:
```json
{
  "creator": "skill1abc...",
  "amount": "1000000uskill",
  "denom": "uskill"
}
```
- **Response**:
```json
{
  "message": "Token burn transaction prepared",
  "tx_info": {
    "creator": "skill1abc...",
    "amount": "1000000uskill",
    "denom": "uskill",
    "estimated_gas": "150000",
    "note": "Use skillchaind tx skillchain burn-tokens to execute"
  }
}
```

---

## Bank Module (Wallet Integration)

### Wallet Balances

#### Get Account Balances
- **GET** `/bank/balances/{address}`
- **Description**: Get all token balances for a specific address
- **Parameters**: 
  - `address`: SkillChain wallet address
- **Response**:
```json
{
  "balances": [
    {
      "denom": "uskill",
      "amount": "120000000"
    },
    {
      "denom": "uvusd",
      "amount": "500000"
    }
  ],
  "pagination": {
    "next_key": null,
    "total": "2"
  }
}
```

#### Get Total Supply
- **GET** `/bank/supply`
- **Description**: Get total supply of all tokens
- **Response**:
```json
{
  "supply": [
    {
      "denom": "uskill",
      "amount": "1000000000"
    },
    {
      "denom": "uvusd",
      "amount": "800000"
    }
  ],
  "pagination": {
    "next_key": null,
    "total": "2"
  }
}
```

#### Get Supply by Denomination
- **GET** `/bank/supply/{denom}`
- **Description**: Get supply for specific token denomination
- **Parameters**: 
  - `denom`: Token denomination (e.g., "uskill", "uvusd")
- **Response**:
```json
{
  "amount": {
    "denom": "uskill",
    "amount": "1000000000"
  }
}
```

---

## Profile Module

### Profile Management

#### Create Profile
- **POST** `/profiles`
- **Description**: Prepare transaction to create user profile
- **Body**:
```json
{
  "creator": "skill1abc...",
  "name": "John Doe",
  "description": "Senior Full Stack Developer",
  "avatar": "https://avatar.url",
  "website": "https://johndoe.dev",
  "location": "New York, USA"
}
```

#### List Profiles
- **GET** `/profiles`
- **Description**: Get list of all user profiles
- **Response**:
```json
{
  "profiles": [
    {
      "address": "skill1abc...",
      "name": "John Doe",
      "description": "Senior Full Stack Developer"
    }
  ]
}
```

#### Get Profile
- **GET** `/profiles/{address}`
- **Description**: Get specific user profile
- **Parameters**: 
  - `address`: User's SkillChain address

#### Update Profile
- **PUT** `/profiles/{address}`
- **Description**: Prepare transaction to update user profile
- **Body**:
```json
{
  "name": "John Doe Jr.",
  "description": "Updated description",
  "avatar": "https://new-avatar.url"
}
```

### Skills Management

#### Add Skill
- **POST** `/profiles/{address}/skills`
- **Description**: Prepare transaction to add skill to user profile
- **Body**:
```json
{
  "skill_name": "JavaScript",
  "proficiency_level": "Expert",
  "years_of_experience": 5,
  "is_verified": false
}
```

#### List User Skills
- **GET** `/profiles/{address}/skills`
- **Description**: Get all skills for a specific user

#### Get Skill Details
- **GET** `/profiles/{address}/skills/{skillId}`
- **Description**: Get detailed information about a specific skill

#### Endorse Skill
- **POST** `/profiles/{address}/skills/{skillId}/endorse`
- **Description**: Prepare transaction to endorse a user's skill
- **Body**:
```json
{
  "endorser_address": "skill1xyz...",
  "endorsement_type": "professional",
  "stake_amount": "100000uskill",
  "comment": "Excellent JavaScript skills demonstrated in our project"
}
```

### Endorsements

#### List Endorsements
- **GET** `/endorsements`
- **Description**: Get list of all skill endorsements

#### Get Endorsement
- **GET** `/endorsements/{endorsementId}`
- **Description**: Get specific endorsement details

---

## Marketplace Module

### Job Management

#### Create Job Posting
- **POST** `/jobs`
- **Description**: Prepare transaction to create job posting
- **Body**:
```json
{
  "creator": "skill1client...",
  "title": "Full Stack Developer Needed",
  "description": "Build a React/Node.js application",
  "budget": "5000000uskill",
  "deadline": "2024-12-31T23:59:59Z",
  "required_skills": ["JavaScript", "React", "Node.js"]
}
```

#### List Job Postings
- **GET** `/jobs`
- **Description**: Get list of all job postings

#### Get Job Posting
- **GET** `/jobs/{jobId}`
- **Description**: Get specific job posting details

#### Update Job Posting
- **PUT** `/jobs/{jobId}`
- **Description**: Prepare transaction to update job posting

#### Close Job Posting
- **POST** `/jobs/{jobId}/close`
- **Description**: Prepare transaction to close job posting

### Proposals Management

#### Submit Proposal
- **POST** `/jobs/{jobId}/proposals`
- **Description**: Prepare transaction to submit job proposal
- **Body**:
```json
{
  "creator": "skill1freelancer...",
  "proposal_text": "I can complete this project in 4 weeks",
  "proposed_budget": "4500000uskill",
  "estimated_duration": "4 weeks"
}
```

#### List Proposals
- **GET** `/jobs/{jobId}/proposals`
- **Description**: Get all proposals for a specific job

#### Get Proposal
- **GET** `/proposals/{proposalId}`
- **Description**: Get specific proposal details

#### Accept Proposal
- **POST** `/proposals/{proposalId}/accept`
- **Description**: Prepare transaction to accept proposal

#### Reject Proposal
- **POST** `/proposals/{proposalId}/reject`
- **Description**: Prepare transaction to reject proposal

### Project Management

#### Create Project
- **POST** `/projects`
- **Description**: Prepare transaction to create project from accepted proposal

#### List Projects
- **GET** `/projects`
- **Description**: Get list of all active projects

#### Get Project
- **GET** `/projects/{projectId}`
- **Description**: Get specific project details

### Milestone Management

#### Create Milestone
- **POST** `/projects/{projectId}/milestones`
- **Description**: Prepare transaction to create project milestone
- **Body**:
```json
{
  "creator": "skill1client...",
  "title": "Frontend Development",
  "description": "Complete React frontend",
  "payment_amount": "2000000uskill",
  "due_date": "2024-06-30T23:59:59Z"
}
```

#### List Milestones
- **GET** `/projects/{projectId}/milestones`
- **Description**: Get all milestones for a project

#### Get Milestone
- **GET** `/milestones/{milestoneId}`
- **Description**: Get specific milestone details

#### Complete Milestone
- **POST** `/milestones/{milestoneId}/complete`
- **Description**: Prepare transaction to mark milestone as complete

#### Approve Milestone
- **POST** `/milestones/{milestoneId}/approve`
- **Description**: Prepare transaction to approve completed milestone

#### Release Payment
- **POST** `/milestones/{milestoneId}/release-payment`
- **Description**: Prepare transaction to release milestone payment

---

## Analytics Module

### Activity Tracking

#### Track Activity
- **POST** `/analytics/activity`
- **Description**: Prepare transaction to record user activity
- **Body**:
```json
{
  "creator": "skill1user...",
  "activity_type": "profile_view",
  "activity_data": "viewed user profile skill1abc...",
  "metadata": "additional context"
}
```

#### List Activities
- **GET** `/analytics/activities`
- **Description**: Get list of all tracked activities

#### Get Activity
- **GET** `/analytics/activities/{activityId}`
- **Description**: Get specific activity details

#### Get User Activity
- **GET** `/analytics/users/{address}/activity`
- **Description**: Get activity history for specific user

### Platform Metrics

#### Record Metric
- **POST** `/analytics/metrics`
- **Description**: Prepare transaction to record platform metric
- **Body**:
```json
{
  "creator": "skill1system...",
  "metric_name": "daily_active_users",
  "value": 150.0,
  "category": "user_engagement"
}
```

#### List Metrics
- **GET** `/analytics/metrics`
- **Description**: Get list of platform metrics

#### Get Metric
- **GET** `/analytics/metrics/{metricName}`
- **Description**: Get specific metric details

### Reports

#### User Report
- **GET** `/analytics/reports/users`
- **Description**: Get user analytics report

#### Platform Report
- **GET** `/analytics/reports/platform`
- **Description**: Get platform analytics report

#### Revenue Report
- **GET** `/analytics/reports/revenue`
- **Description**: Get revenue analytics report

---

## File Storage Module

### File Management

#### Upload File
- **POST** `/files`
- **Description**: Prepare transaction to register file upload
- **Body**:
```json
{
  "creator": "skill1user...",
  "file_name": "portfolio.pdf",
  "file_size": 2048576,
  "file_type": "application/pdf",
  "ipfs_hash": "QmXxXxXxXxXxXxXxXxXxXxXxXxXxXxXxXxXx",
  "description": "My portfolio document"
}
```

#### List Files
- **GET** `/files`
- **Description**: Get list of uploaded files

#### Get File
- **GET** `/files/{fileId}`
- **Description**: Get specific file metadata

#### Update File
- **PUT** `/files/{fileId}`
- **Description**: Prepare transaction to update file metadata

#### Delete File
- **DELETE** `/files/{fileId}`
- **Description**: Prepare transaction to delete file

#### Download File
- **GET** `/files/{fileId}/download`
- **Description**: Get file download information

### File Permissions

#### Grant File Permission
- **POST** `/files/{fileId}/permissions`
- **Description**: Prepare transaction to grant file access permission
- **Body**:
```json
{
  "creator": "skill1owner...",
  "grantee_address": "skill1user...",
  "permission_type": "read"
}
```

#### List File Permissions
- **GET** `/files/{fileId}/permissions`
- **Description**: Get list of file permissions

#### Get File Permission
- **GET** `/files/{fileId}/permissions/{granteeAddress}`
- **Description**: Get specific permission details

#### Revoke File Permission
- **DELETE** `/files/{fileId}/permissions/{granteeAddress}`
- **Description**: Prepare transaction to revoke file permission

### IPFS Operations

#### Pin IPFS Hash
- **POST** `/ipfs/{hash}/pin`
- **Description**: Pin content to IPFS network

#### Unpin IPFS Hash
- **POST** `/ipfs/{hash}/unpin`
- **Description**: Unpin content from IPFS network

#### Get IPFS Status
- **GET** `/ipfs/{hash}/status`
- **Description**: Check IPFS hash status

---

## Notifications Module

### Notification Management

#### Create Notification
- **POST** `/notifications`
- **Description**: Prepare transaction to create notification
- **Body**:
```json
{
  "creator": "skill1system...",
  "title": "New Job Proposal",
  "message": "You have received a new proposal for your job posting",
  "priority": "high",
  "recipient": "skill1user..."
}
```

#### List Notifications
- **GET** `/notifications`
- **Description**: Get list of all notifications

#### Get Notification
- **GET** `/notifications/{notificationId}`
- **Description**: Get specific notification details

#### Mark as Read
- **POST** `/notifications/{notificationId}/read`
- **Description**: Prepare transaction to mark notification as read

#### Delete Notification
- **DELETE** `/notifications/{notificationId}`
- **Description**: Prepare transaction to delete notification

### User Notifications

#### Get User Notifications
- **GET** `/users/{address}/notifications`
- **Description**: Get all notifications for specific user

#### Get Unread Notifications
- **GET** `/users/{address}/notifications/unread`
- **Description**: Get unread notifications for specific user

#### Mark All as Read
- **POST** `/users/{address}/notifications/mark-all-read`
- **Description**: Prepare transaction to mark all notifications as read

#### Get Notification Preferences
- **GET** `/users/{address}/notification-preferences`
- **Description**: Get user's notification preferences

#### Update Notification Preferences
- **PUT** `/users/{address}/notification-preferences`
- **Description**: Prepare transaction to update notification preferences

### Push Notifications

#### Subscribe to Push
- **POST** `/push/subscribe`
- **Description**: Subscribe to push notifications
- **Body**:
```json
{
  "address": "skill1user...",
  "endpoint": "https://fcm.googleapis.com/fcm/send/...",
  "keys": {
    "p256dh": "key_data",
    "auth": "auth_data"
  }
}
```

#### Unsubscribe from Push
- **POST** `/push/unsubscribe`
- **Description**: Unsubscribe from push notifications

#### Send Push Notification
- **POST** `/push/send`
- **Description**: Send push notification (system only)

---

## Testing Endpoints

All endpoints can be tested using curl or any HTTP client:

### Example: Get Token Info
```bash
curl -X GET "http://localhost:1317/skillchain/v1/token/info"
```

### Example: Check Wallet Balance
```bash
curl -X GET "http://localhost:1317/skillchain/v1/bank/balances/skill1abc..."
```

### Example: Convert SKILL to VUSD
```bash
curl -X POST "http://localhost:1317/skillchain/v1/vusd/convert/skill-to-vusd" \
  -H "Content-Type: application/json" \
  -d '{
    "creator": "skill1abc...",
    "amount": "1000000uskill"
  }'
```

---

## Error Codes

| Code | Description |
|------|-------------|
| 400  | Bad Request - Invalid parameters |
| 404  | Not Found - Resource doesn't exist |
| 500  | Internal Server Error - Blockchain error |

---

## Rate Limits

Currently no rate limits are implemented, but recommended limits for production:
- Query endpoints: 100 requests/minute
- Transaction endpoints: 10 requests/minute

---

## CORS Support

All endpoints support CORS with the following headers:
- `Access-Control-Allow-Origin: *`
- `Access-Control-Allow-Methods: GET, POST, PUT, DELETE, OPTIONS`
- `Access-Control-Allow-Headers: *`

---

## Legacy Support

All bank endpoints are also available without the `/v1` prefix for backward compatibility:
- `http://localhost:1317/skillchain/bank/balances/{address}`
- `http://localhost:1317/skillchain/bank/supply`
- `http://localhost:1317/skillchain/bank/supply/{denom}`

---

## Last Updated

**Date**: January 2024  
**Version**: 1.0  
**Total Endpoints**: 75+  

For the latest updates and changes, check the GitHub repository or contact the development team. 