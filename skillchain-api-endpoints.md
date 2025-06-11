# 🚀 SkillChain REST API Endpoints

## 📋 **Overview**
Professional SkillChain-branded REST API endpoints for blockchain operations, marketplace management, and user interactions.

**Base URL:** `http://localhost:1317`

---

## 🔍 **Query Endpoints (GET)**

### 💰 **Bank Operations**
- `GET /skillchain/bank/balances/{address}` - Get account balances
- `GET /skillchain/bank/balance/{address}/{denom}` - Get specific token balance  
- `GET /skillchain/bank/supply` - Get total token supply

### ℹ️ **Chain Information**
- `GET /skillchain/info/status` - Get blockchain status
- `GET /skillchain/info/params` - Get chain parameters

### 👤 **Account Information**
- `GET /skillchain/accounts/{address}/info` - Get comprehensive account info
- `GET /skillchain/accounts/{address}/transactions` - Get transaction history

---

## 📝 **Transaction Endpoints (POST)**

### 💱 **VUSD Operations**

#### Convert SKILL to VUSD
```bash
POST /skillchain/convert/skill-to-vusd
```
**Request Body:**
```json
{
  "from_address": "skill1d2sdxqwcywmt9gd7fqnempzs3zj6zsakgvrsdf",
  "amount": "1000",
  "signature": "optional"
}
```

#### Convert VUSD to SKILL
```bash
POST /skillchain/convert/vusd-to-skill
```
**Request Body:**
```json
{
  "from_address": "skill1d2sdxqwcywmt9gd7fqnempzs3zj6zsakgvrsdf",
  "amount": "500"
}
```

#### Burn Tokens
```bash
POST /skillchain/token/burn
```
**Request Body:**
```json
{
  "from_address": "skill1d2sdxqwcywmt9gd7fqnempzs3zj6zsakgvrsdf",
  "amount": "100",
  "denom": "skill"
}
```

---

### 🏢 **Marketplace Operations**

#### Create Job Posting
```bash
POST /skillchain/marketplace/jobs
```
**Request Body:**
```json
{
  "employer_address": "skill1d2sdxqwcywmt9gd7fqnempzs3zj6zsakgvrsdf",
  "title": "React Developer Required",
  "description": "Need a React developer for e-commerce project",
  "budget": "5000",
  "skills": ["React", "JavaScript", "Node.js"],
  "deadline": "2024-01-15"
}
```

#### Submit Proposal
```bash
POST /skillchain/marketplace/jobs/{jobId}/proposals
```
**Request Body:**
```json
{
  "freelancer_address": "skill18jfkq29mf4nlnjj0sumsxr50g5knc0cwhpx828",
  "proposed_budget": "4500",
  "timeline": "2 weeks",
  "description": "I have 5 years experience in React development"
}
```

#### Accept Proposal
```bash
POST /skillchain/marketplace/proposals/{proposalId}/accept
```
**Request Body:**
```json
{
  "employer_address": "skill1d2sdxqwcywmt9gd7fqnempzs3zj6zsakgvrsdf"
}
```

#### Create Project
```bash
POST /skillchain/marketplace/projects
```
**Request Body:**
```json
{
  "employer_address": "skill1d2sdxqwcywmt9gd7fqnempzs3zj6zsakgvrsdf",
  "freelancer_address": "skill18jfkq29mf4nlnjj0sumsxr50g5knc0cwhpx828",
  "title": "E-commerce Website",
  "budget": "5000",
  "milestones": [
    {
      "description": "Frontend Development",
      "amount": "2500"
    },
    {
      "description": "Backend Integration",
      "amount": "2500"
    }
  ]
}
```

#### Complete Milestone
```bash
POST /skillchain/marketplace/projects/{projectId}/milestones/{milestoneId}/complete
```
**Request Body:**
```json
{
  "freelancer_address": "skill18jfkq29mf4nlnjj0sumsxr50g5knc0cwhpx828",
  "delivery_notes": "Frontend completed with all requirements",
  "files": ["https://github.com/user/project"]
}
```

#### Release Payment
```bash
POST /skillchain/marketplace/projects/{projectId}/payments/release
```
**Request Body:**
```json
{
  "employer_address": "skill1d2sdxqwcywmt9gd7fqnempzs3zj6zsakgvrsdf",
  "milestone_id": "milestone_123",
  "rating": 5,
  "feedback": "Excellent work, delivered on time"
}
```

---

### 👤 **Profile Operations**

#### Create User Profile
```bash
POST /skillchain/profiles
```
**Request Body:**
```json
{
  "address": "skill18jfkq29mf4nlnjj0sumsxr50g5knc0cwhpx828",
  "name": "Bob Johnson",
  "bio": "Full-stack developer with 5 years experience",
  "skills": ["React", "Node.js", "Python", "Docker"],
  "location": "Istanbul, Turkey",
  "website": "https://bobjohnson.dev",
  "profile_type": "freelancer"
}
```

#### Add User Skill
```bash
POST /skillchain/profiles/{address}/skills
```
**Request Body:**
```json
{
  "skill_name": "React",
  "proficiency": "expert",
  "experience": "5 years",
  "certificates": ["React Professional Certificate"]
}
```

#### Endorse Skill
```bash
POST /skillchain/profiles/{address}/skills/{skillId}/endorse
```
**Request Body:**
```json
{
  "endorser_address": "skill1d2sdxqwcywmt9gd7fqnempzs3zj6zsakgvrsdf",
  "stake_amount": "100",
  "review": "Excellent React developer",
  "rating": 5
}
```

---

### 🔄 **Transfer Operations**

#### Transfer Tokens
```bash
POST /skillchain/transfer
```
**Request Body:**
```json
{
  "from_address": "skill1d2sdxqwcywmt9gd7fqnempzs3zj6zsakgvrsdf",
  "to_address": "skill18jfkq29mf4nlnjj0sumsxr50g5knc0cwhpx828",
  "amount": "1000",
  "denom": "skill",
  "memo": "Payment for React project"
}
```

---

## 🧪 **Example Workflow**

### Complete Freelance Job Process:

1. **Employer creates job:**
```bash
curl -X POST http://localhost:1317/skillchain/marketplace/jobs \
  -H "Content-Type: application/json" \
  -d '{"employer_address": "skill1d2sdxqwcywmt9gd7fqnempzs3zj6zsakgvrsdf", "title": "React Developer", "budget": "5000"}'
```

2. **Freelancer submits proposal:**
```bash
curl -X POST http://localhost:1317/skillchain/marketplace/jobs/job_12345/proposals \
  -H "Content-Type: application/json" \
  -d '{"freelancer_address": "skill18jfkq29mf4nlnjj0sumsxr50g5knc0cwhpx828", "proposed_budget": "4500"}'
```

3. **Employer accepts proposal:**
```bash
curl -X POST http://localhost:1317/skillchain/marketplace/proposals/prop_67890/accept \
  -H "Content-Type: application/json" \
  -d '{"employer_address": "skill1d2sdxqwcywmt9gd7fqnempzs3zj6zsakgvrsdf"}'
```

4. **Payment transfer:**
```bash
curl -X POST http://localhost:1317/skillchain/transfer \
  -H "Content-Type: application/json" \
  -d '{"from_address": "skill1d2sdxqwcywmt9gd7fqnempzs3zj6zsakgvrsdf", "to_address": "skill18jfkq29mf4nlnjj0sumsxr50g5knc0cwhpx828", "amount": "4500", "denom": "skill"}'
```

---

## 📊 **Response Format**

All endpoints return consistent JSON responses:

```json
{
  "status": "success|pending|error",
  "tx_type": "operation_type",
  "message": "Human readable message",
  "data": { /* operation specific data */ }
}
```

---

## 🔧 **Features**

- ✅ **Professional branding** with `/skillchain/` prefix
- ✅ **Enhanced data format** with human-readable fields
- ✅ **Complete workflow support** for freelance marketplace
- ✅ **VUSD conversion** for stable payments
- ✅ **Profile management** with skill endorsements
- ✅ **Token transfers** with memo support
- ✅ **Validation** and error handling
- ✅ **Consistent response format**

---

## 🚀 **Next Steps**

1. **Real transaction integration** with blockchain state
2. **Authentication** and signature verification
3. **Rate limiting** and API security
4. **WebSocket** support for real-time updates
5. **Advanced filtering** and pagination
6. **API versioning** for backward compatibility

---

**🎯 Status:** Production-ready API structure with placeholder data. Ready for frontend integration and real blockchain transaction handling. 