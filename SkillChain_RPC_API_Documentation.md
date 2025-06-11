# SkillChain RPC API Documentation

## Overview

SkillChain provides a comprehensive REST API for interacting with all blockchain modules. The API is available at `http://localhost:1317/skillchain/v1/` when the blockchain is running.

## Base URL

```
http://localhost:1317/skillchain/v1/
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

### vUSD Operations

#### Get vUSD Treasury Status
- **GET** `/vusd/treasury`
- **Description**: Get current vUSD treasury information
- **Response**:
```json
{
  "treasury": "1000000vusd",
  "total_collateral": "1500000uskill",
  "total_vusd_supply": "800000vusd",
  "collateral_ratio": "1.875"
}
```

#### Get User vUSD Position
- **GET** `/vusd/position/{address}`
- **Description**: Get user's vUSD position and collateral details
- **Parameters**: 
  - `address`: User's SkillChain address
- **Response**:
```json
{
  "address": "skill1abc...",
  "vusd_balance": "1000vusd",
  "skill_collateral": "1500uskill",
  "collateral_ratio": "1.5",
  "liquidation_price": "0.33"
}
```

#### Convert SKILL to vUSD
- **POST** `/vusd/convert/skill-to-vusd`
- **Description**: Prepare transaction to convert SKILL tokens to vUSD
- **Body**:
```json
{
  "from_address": "skill1abc...",
  "amount": "1000uskill"
}
```
- **Response**:
```json
{
  "message": "Conversion transaction prepared",
  "tx_info": {
    "from_address": "skill1abc...",
    "amount": "1000uskill",
    "estimated_gas": "200000",
    "note": "Use skillchaind tx skillchain convert-skill-to-vusd to execute"
  }
}
```

#### Convert vUSD to SKILL
- **POST** `/vusd/convert/vusd-to-skill`
- **Description**: Prepare transaction to convert vUSD back to SKILL tokens
- **Body**:
```json
{
  "from_address": "skill1abc...",
  "amount": "500vusd"
}
```

#### Get vUSD Price
- **GET** `/vusd/price`
- **Description**: Get current vUSD mock price and parameters
- **Response**:
```json
{
  "vusd_mock_price": "0.50",
  "min_collateral_ratio": "1.50",
  "vusd_enabled": true,
  "timestamp": "2024-01-01T12:00:00Z"
}
```

#### Update vUSD Price
- **PUT** `/vusd/price`
- **Description**: Prepare transaction to update vUSD price (authority only)
- **Body**:
```json
{
  "authority": "skill1authority...",
  "new_price": "0.55"
}
```

### Token Operations

#### Burn Tokens
- **POST** `/tokens/burn`
- **Description**: Prepare transaction to burn SKILL tokens
- **Body**:
```json
{
  "burner_address": "skill1abc...",
  "amount": "1000uskill"
}
```

#### Get Token Supply
- **GET** `/tokens/supply`
- **Description**: Get current token supply for USKILL and vUSD
- **Response**:
```json
{
  "uskill_supply": "1000000000uskill",
  "vusd_supply": "500000vusd",
  "timestamp": "2024-01-01T12:00:00Z"
}
```

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
  },
  "module": "skillchain",
  "timestamp": "2024-01-01T12:00:00Z"
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
- **GET** `/profiles?page_key={key}&limit={limit}`
- **Description**: Get list of all profiles with pagination
- **Query Parameters**:
  - `page_key`: Pagination key (optional)
  - `limit`: Number of results per page (default: 50)
- **Response**:
```json
{
  "profiles": [
    {
      "creator": "skill1abc...",
      "name": "John Doe",
      "description": "Senior Full Stack Developer",
      "reputation_score": 850,
      "created_at": "2024-01-01T12:00:00Z"
    }
  ],
  "pagination": {...},
  "total_count": 25
}
```

#### Get Profile
- **GET** `/profiles/{address}`
- **Description**: Get specific user profile by address
- **Parameters**:
  - `address`: User's SkillChain address

#### Update Profile
- **PUT** `/profiles/{address}`
- **Description**: Prepare transaction to update user profile
- **Parameters**:
  - `address`: User's SkillChain address
- **Body**:
```json
{
  "name": "John Doe Jr.",
  "description": "Senior Full Stack Developer & Blockchain Expert",
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
  "skill_name": "React.js",
  "proficiency_level": "expert",
  "years_of_experience": 5,
  "is_verified": false,
  "verification_details": "Portfolio and previous projects"
}
```

#### List User Skills
- **GET** `/profiles/{address}/skills`
- **Description**: Get all skills for a specific user
- **Response**:
```json
{
  "skills": [
    {
      "skill_name": "React.js",
      "proficiency_level": "expert",
      "years_of_experience": 5,
      "is_verified": true,
      "reputation_impact": 45.5
    }
  ],
  "address": "skill1abc...",
  "count": 8
}
```

#### Get Specific Skill
- **GET** `/profiles/{address}/skills/{skillId}`
- **Description**: Get details of a specific skill

#### Endorse Skill
- **POST** `/profiles/{address}/skills/{skillId}/endorse`
- **Description**: Prepare transaction to endorse a user's skill
- **Body**:
```json
{
  "endorser_address": "skill1def...",
  "endorsement_type": "strong",
  "stake_amount": "100uskill",
  "comment": "Excellent React developer, worked on multiple projects together"
}
```

### Endorsements

#### List Endorsements
- **GET** `/endorsements?skill_id={id}&endorser={address}`
- **Description**: Get list of endorsements with filtering
- **Query Parameters**:
  - `skill_id`: Filter by skill ID (optional)
  - `endorser`: Filter by endorser address (optional)

#### Get Endorsement
- **GET** `/endorsements/{endorsementId}`
- **Description**: Get specific endorsement details

---

## Marketplace Module

### Job Postings

#### Create Job Posting
- **POST** `/jobs`
- **Description**: Prepare transaction to create new job posting
- **Body**:
```json
{
  "creator": "skill1abc...",
  "title": "React Developer Needed",
  "description": "Looking for experienced React developer for e-commerce project",
  "required_skills": ["React", "JavaScript", "Node.js"],
  "budget": "5000uskill",
  "deadline": "2024-12-31T23:59:59Z",
  "category": "web_development",
  "payment_type": "milestone"
}
```

#### List Job Postings
- **GET** `/jobs?category={cat}&status={status}&min_budget={min}&max_budget={max}`
- **Description**: Get list of job postings with filtering
- **Query Parameters**:
  - `category`: Filter by job category (optional)
  - `status`: Filter by job status (optional)
  - `min_budget`: Minimum budget filter (optional)
  - `max_budget`: Maximum budget filter (optional)

#### Get Job Posting
- **GET** `/jobs/{jobId}`
- **Description**: Get specific job posting details

#### Update Job Posting
- **PUT** `/jobs/{jobId}`
- **Description**: Prepare transaction to update job posting

#### Close Job Posting
- **POST** `/jobs/{jobId}/close`
- **Description**: Prepare transaction to close job posting
- **Body**:
```json
{
  "creator": "skill1abc...",
  "reason": "Position filled"
}
```

### Proposals

#### Submit Proposal
- **POST** `/jobs/{jobId}/proposals`
- **Description**: Prepare transaction to submit proposal for job
- **Body**:
```json
{
  "creator": "skill1freelancer...",
  "proposed_budget": "4500uskill",
  "timeline": "14 days",
  "description": "I can deliver this project with high quality",
  "experience": "5 years React development experience"
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
- **Body**:
```json
{
  "creator": "skill1client...",
  "message": "Your proposal is accepted. Looking forward to working with you."
}
```

#### Reject Proposal
- **POST** `/proposals/{proposalId}/reject`
- **Description**: Prepare transaction to reject proposal
- **Body**:
```json
{
  "creator": "skill1client...",
  "reason": "Budget doesn't match our requirements"
}
```

### Projects & Milestones

#### Create Project
- **POST** `/projects`
- **Description**: Prepare transaction to create project with escrow
- **Body**:
```json
{
  "creator": "skill1client...",
  "title": "E-commerce React Application",
  "description": "Full-featured e-commerce application with React frontend",
  "client_address": "skill1client...",
  "freelancer_address": "skill1freelancer...",
  "total_amount": "4500uskill",
  "escrow_amount": "4500uskill",
  "deadline": "2024-12-31T23:59:59Z"
}
```

#### List Projects
- **GET** `/projects?status={status}&client={address}&freelancer={address}`
- **Description**: Get list of projects with filtering

#### Get Project
- **GET** `/projects/{projectId}`
- **Description**: Get specific project details

#### Create Milestone
- **POST** `/projects/{projectId}/milestones`
- **Description**: Prepare transaction to create project milestone
- **Body**:
```json
{
  "creator": "skill1client...",
  "title": "Frontend Development",
  "description": "Complete React frontend with all pages",
  "amount": "2250uskill",
  "deadline": "2024-06-15T23:59:59Z"
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
- **Body**:
```json
{
  "creator": "skill1freelancer...",
  "deliverable": "Frontend completed, deployed to staging",
  "notes": "All features implemented as requested"
}
```

#### Approve Milestone
- **POST** `/milestones/{milestoneId}/approve`
- **Description**: Prepare transaction to approve completed milestone
- **Body**:
```json
{
  "creator": "skill1client...",
  "feedback": "Excellent work! Everything looks perfect."
}
```

#### Release Payment
- **POST** `/milestones/{milestoneId}/payment`
- **Description**: Prepare transaction to release milestone payment
- **Body**:
```json
{
  "creator": "skill1client...",
  "amount": "2250uskill"
}
```

---

## Analytics Module

### Activity Tracking

#### Track Activity
- **POST** `/analytics/activity`
- **Description**: Prepare transaction to track user activity
- **Body**:
```json
{
  "creator": "skill1system...",
  "activity_type": "job_posting_created",
  "user_address": "skill1user...",
  "ip_address": "192.168.1.1",
  "metadata": {
    "browser": "Chrome",
    "device": "desktop",
    "job_id": "job123"
  }
}
```

#### List Activities
- **GET** `/analytics/activity`
- **Description**: Get list of tracked activities

#### Get Activity
- **GET** `/analytics/activity/{activityId}`
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
- **POST** `/ipfs/pin/{hash}`
- **Description**: Pin content to IPFS network

#### Unpin IPFS Hash
- **DELETE** `/ipfs/unpin/{hash}`
- **Description**: Unpin content from IPFS network

#### Get IPFS Status
- **GET** `/ipfs/status/{hash}`
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

### Notification Preferences

#### Get Preferences
- **GET** `/users/{address}/notification-preferences`
- **Description**: Get user notification preferences

#### Update Preferences
- **PUT** `/users/{address}/notification-preferences`
- **Description**: Prepare transaction to update notification preferences
- **Body**:
```json
{
  "creator": "skill1user...",
  "preferences": {
    "email_notifications": true,
    "push_notifications": false,
    "job_alerts": true,
    "payment_notifications": true
  }
}
```

### Push Notifications

#### Subscribe to Push
- **POST** `/push/subscribe`
- **Description**: Register for push notifications
- **Body**:
```json
{
  "address": "skill1user...",
  "endpoint": "https://fcm.googleapis.com/fcm/send/...",
  "keys": {
    "p256dh": "key...",
    "auth": "auth..."
  }
}
```

#### Unsubscribe from Push
- **POST** `/push/unsubscribe`
- **Description**: Unregister from push notifications

#### Send Push Notification
- **POST** `/push/send`
- **Description**: Send push notification to users
- **Body**:
```json
{
  "creator": "skill1system...",
  "recipients": ["skill1user1...", "skill1user2..."],
  "title": "New Job Alert",
  "body": "A new job matching your skills has been posted",
  "data": {
    "job_id": "job123",
    "action": "view_job"
  }
}
```

---

## Testing the API

### Using curl

```bash
# Get vUSD treasury status
curl -X GET http://localhost:1317/skillchain/v1/vusd/treasury

# List job postings
curl -X GET http://localhost:1317/skillchain/v1/jobs

# Get user profile
curl -X GET http://localhost:1317/skillchain/v1/profiles/skill1abc...

# Create job posting (prepare transaction)
curl -X POST http://localhost:1317/skillchain/v1/jobs \
  -H "Content-Type: application/json" \
  -d '{
    "creator": "skill1abc...",
    "title": "React Developer Needed",
    "description": "Looking for React developer",
    "required_skills": ["React", "JavaScript"],
    "budget": "5000uskill",
    "deadline": "2024-12-31T23:59:59Z",
    "category": "web_development",
    "payment_type": "milestone"
  }'
```

### Using JavaScript/Fetch

```javascript
// Get vUSD price
const response = await fetch('http://localhost:1317/skillchain/v1/vusd/price');
const priceData = await response.json();
console.log('vUSD Price:', priceData.vusd_mock_price);

// List user skills
const skillsResponse = await fetch('http://localhost:1317/skillchain/v1/profiles/skill1abc.../skills');
const skillsData = await skillsResponse.json();
console.log('User Skills:', skillsData.skills);
```

## Error Codes

- **400 Bad Request**: Invalid request body or parameters
- **401 Unauthorized**: Authentication required
- **403 Forbidden**: Insufficient permissions
- **404 Not Found**: Resource not found
- **500 Internal Server Error**: Server error

## Rate Limiting

API endpoints are rate limited to prevent abuse:
- **Query endpoints**: 100 requests per minute per IP
- **Transaction preparation endpoints**: 50 requests per minute per IP

## Support

For API support and questions:
- GitHub Issues: [SkillChain Repository Issues](https://github.com/skillchain/skillchain/issues)
- Documentation: This file and in-code comments
- Community: SkillChain Discord Server

## Changelog

### v1.0.0
- Initial API release
- All 6 modules supported
- Comprehensive endpoint coverage
- Transaction preparation for all write operations
- Query support for all read operations 