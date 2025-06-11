package app

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/gorilla/mux"
	
	marketplacetypes "skillchain/x/marketplace/types"
)

// Job Posting Handlers

func (app *App) handleCreateJobPosting(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			Creator         string   `json:"creator"`
			Title           string   `json:"title"`
			Description     string   `json:"description"`
			RequiredSkills  []string `json:"required_skills"`
			Budget          string   `json:"budget"`
			Deadline        string   `json:"deadline"`
			Category        string   `json:"category"`
			PaymentType     string   `json:"payment_type"`
		}

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		// Validate creator address
		if _, err := types.AccAddressFromBech32(req.Creator); err != nil {
			http.Error(w, "Invalid creator address format", http.StatusBadRequest)
			return
		}

		// Parse budget
		budget, err := types.ParseCoinNormalized(req.Budget)
		if err != nil {
			http.Error(w, "Invalid budget format", http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "Job posting creation transaction prepared",
			"tx_info": map[string]interface{}{
				"creator": req.Creator,
				"title": req.Title,
				"description": req.Description,
				"required_skills": req.RequiredSkills,
				"budget": budget.String(),
				"deadline": req.Deadline,
				"category": req.Category,
				"payment_type": req.PaymentType,
				"estimated_gas": "200000",
				"note": "Use skillchaind tx marketplace create-job-posting to execute",
			},
		})
	}
}

func (app *App) handleListJobPostings(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		queryClient := marketplacetypes.NewQueryClient(clientCtx)
		
		// Parse query parameters
		skillsRequired := r.URL.Query().Get("skills_required")
		isActive := r.URL.Query().Get("is_active")
		minBudget := r.URL.Query().Get("min_budget")
		maxBudget := r.URL.Query().Get("max_budget")
		
		req := &marketplacetypes.QueryAllJobPostingRequest{
			Pagination: &query.PageRequest{
				Limit: 50,
			},
		}

		res, err := queryClient.JobPostingAll(r.Context(), req)
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to query job postings: %v", err), http.StatusInternalServerError)
			return
		}

		// Apply filters
		var filteredJobs []interface{}
		for _, job := range res.JobPosting {
			include := true
			
			if skillsRequired != "" && job.SkillsRequired != skillsRequired {
				include = false
			}
			if isActive != "" {
				if (isActive == "true" && !job.IsActive) || (isActive == "false" && job.IsActive) {
					include = false
				}
			}
			// Add budget filtering logic if needed
			
			if include {
				filteredJobs = append(filteredJobs, job)
			}
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"job_postings": filteredJobs,
			"filters": map[string]string{
				"skills_required": skillsRequired,
				"is_active": isActive,
				"min_budget": minBudget,
				"max_budget": maxBudget,
			},
			"count": len(filteredJobs),
		})
	}
}

func (app *App) handleGetJobPosting(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		jobID := vars["jobId"]

		queryClient := marketplacetypes.NewQueryClient(clientCtx)
		req := &marketplacetypes.QueryGetJobPostingRequest{
			Index: jobID,
		}

		res, err := queryClient.JobPosting(r.Context(), req)
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to query job posting: %v", err), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"job_posting": res.JobPosting,
			"job_id": jobID,
		})
	}
}

func (app *App) handleUpdateJobPosting(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		jobID := vars["jobId"]

		var req struct {
			Creator         string   `json:"creator"`
			Title           string   `json:"title,omitempty"`
			Description     string   `json:"description,omitempty"`
			RequiredSkills  []string `json:"required_skills,omitempty"`
			Budget          string   `json:"budget,omitempty"`
			Deadline        string   `json:"deadline,omitempty"`
		}

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "Job posting update transaction prepared",
			"tx_info": map[string]interface{}{
				"job_id": jobID,
				"creator": req.Creator,
				"updates": req,
				"estimated_gas": "150000",
				"note": "Use skillchaind tx marketplace update-job-posting to execute",
			},
		})
	}
}

func (app *App) handleCloseJobPosting(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		jobID := vars["jobId"]

		var req struct {
			Creator string `json:"creator"`
			Reason  string `json:"reason,omitempty"`
		}

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "Job posting closure transaction prepared",
			"tx_info": map[string]interface{}{
				"job_id": jobID,
				"creator": req.Creator,
				"reason": req.Reason,
				"estimated_gas": "100000",
				"note": "Use skillchaind tx marketplace close-job-posting to execute",
			},
		})
	}
}

// Proposal Handlers

func (app *App) handleSubmitProposal(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		jobID := vars["jobId"]

		var req struct {
			Creator         string `json:"creator"`
			ProposedBudget  string `json:"proposed_budget"`
			Timeline        string `json:"timeline"`
			Description     string `json:"description"`
			Experience      string `json:"experience"`
		}

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		// Validate creator address
		if _, err := types.AccAddressFromBech32(req.Creator); err != nil {
			http.Error(w, "Invalid creator address format", http.StatusBadRequest)
			return
		}

		// Parse proposed budget
		budget, err := types.ParseCoinNormalized(req.ProposedBudget)
		if err != nil {
			http.Error(w, "Invalid proposed budget format", http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "Proposal submission transaction prepared",
			"tx_info": map[string]interface{}{
				"job_id": jobID,
				"creator": req.Creator,
				"proposed_budget": budget.String(),
				"timeline": req.Timeline,
				"description": req.Description,
				"experience": req.Experience,
				"estimated_gas": "180000",
				"note": "Use skillchaind tx marketplace submit-proposal to execute",
			},
		})
	}
}

func (app *App) handleListProposals(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		jobID := vars["jobId"]

		queryClient := marketplacetypes.NewQueryClient(clientCtx)
		req := &marketplacetypes.QueryAllProposalRequest{
			Pagination: &query.PageRequest{
				Limit: 100,
			},
		}

		res, err := queryClient.ProposalAll(r.Context(), req)
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to query proposals: %v", err), http.StatusInternalServerError)
			return
		}

		// Filter proposals by job ID
		var jobProposals []interface{}
		for _, proposal := range res.Proposal {
			if proposal.JobPostingId == jobID {
				jobProposals = append(jobProposals, proposal)
			}
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"proposals": jobProposals,
			"job_id": jobID,
			"count": len(jobProposals),
		})
	}
}

func (app *App) handleGetProposal(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		proposalID := vars["proposalId"]

		queryClient := marketplacetypes.NewQueryClient(clientCtx)
		req := &marketplacetypes.QueryGetProposalRequest{
			Index: proposalID,
		}

		res, err := queryClient.Proposal(r.Context(), req)
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to query proposal: %v", err), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"proposal": res.Proposal,
			"proposal_id": proposalID,
		})
	}
}

func (app *App) handleAcceptProposal(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		proposalID := vars["proposalId"]

		var req struct {
			Creator string `json:"creator"`
			Message string `json:"message,omitempty"`
		}

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "Proposal acceptance transaction prepared",
			"tx_info": map[string]interface{}{
				"proposal_id": proposalID,
				"creator": req.Creator,
				"message": req.Message,
				"estimated_gas": "200000",
				"note": "Use skillchaind tx marketplace accept-proposal to execute",
			},
		})
	}
}

func (app *App) handleRejectProposal(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		proposalID := vars["proposalId"]

		var req struct {
			Creator string `json:"creator"`
			Reason  string `json:"reason,omitempty"`
		}

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "Proposal rejection transaction prepared",
			"tx_info": map[string]interface{}{
				"proposal_id": proposalID,
				"creator": req.Creator,
				"reason": req.Reason,
				"estimated_gas": "120000",
				"note": "Use skillchaind tx marketplace reject-proposal to execute",
			},
		})
	}
}

// Project & Milestone Handlers

func (app *App) handleCreateProject(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			Creator         string `json:"creator"`
			Title           string `json:"title"`
			Description     string `json:"description"`
			ClientAddress   string `json:"client_address"`
			FreelancerAddress string `json:"freelancer_address"`
			TotalAmount     string `json:"total_amount"`
			EscrowAmount    string `json:"escrow_amount"`
			Deadline        string `json:"deadline"`
		}

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		// Validate addresses
		if _, err := types.AccAddressFromBech32(req.Creator); err != nil {
			http.Error(w, "Invalid creator address format", http.StatusBadRequest)
			return
		}
		if _, err := types.AccAddressFromBech32(req.ClientAddress); err != nil {
			http.Error(w, "Invalid client address format", http.StatusBadRequest)
			return
		}
		if _, err := types.AccAddressFromBech32(req.FreelancerAddress); err != nil {
			http.Error(w, "Invalid freelancer address format", http.StatusBadRequest)
			return
		}

		// Parse amounts
		totalAmount, err := types.ParseCoinNormalized(req.TotalAmount)
		if err != nil {
			http.Error(w, "Invalid total amount format", http.StatusBadRequest)
			return
		}

		escrowAmount, err := types.ParseCoinNormalized(req.EscrowAmount)
		if err != nil {
			http.Error(w, "Invalid escrow amount format", http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "Project creation transaction prepared",
			"tx_info": map[string]interface{}{
				"creator": req.Creator,
				"title": req.Title,
				"description": req.Description,
				"client_address": req.ClientAddress,
				"freelancer_address": req.FreelancerAddress,
				"total_amount": totalAmount.String(),
				"escrow_amount": escrowAmount.String(),
				"deadline": req.Deadline,
				"estimated_gas": "250000",
				"note": "Use skillchaind tx marketplace create-project to execute",
			},
		})
	}
}

func (app *App) handleListProjects(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		queryClient := marketplacetypes.NewQueryClient(clientCtx)
		
		// Parse filters
		isCompleted := r.URL.Query().Get("is_completed")
		clientAddr := r.URL.Query().Get("client")
		freelancerAddr := r.URL.Query().Get("freelancer")
		
		req := &marketplacetypes.QueryAllProjectRequest{
			Pagination: &query.PageRequest{
				Limit: 50,
			},
		}

		res, err := queryClient.ProjectAll(r.Context(), req)
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to query projects: %v", err), http.StatusInternalServerError)
			return
		}

		// Apply filters
		var filteredProjects []interface{}
		for _, project := range res.Project {
			include := true
			
			// Note: Using available fields from Project type
			if clientAddr != "" && project.ClientAddress != clientAddr {
				include = false
			}
			if freelancerAddr != "" && project.FreelancerAddress != freelancerAddr {
				include = false
			}
			
			if include {
				filteredProjects = append(filteredProjects, project)
			}
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"projects": filteredProjects,
			"filters": map[string]string{
				"is_completed": isCompleted,
				"client": clientAddr,
				"freelancer": freelancerAddr,
			},
			"count": len(filteredProjects),
		})
	}
}

func (app *App) handleGetProject(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		projectID := vars["projectId"]

		queryClient := marketplacetypes.NewQueryClient(clientCtx)
		req := &marketplacetypes.QueryGetProjectRequest{
			Index: projectID,
		}

		res, err := queryClient.Project(r.Context(), req)
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to query project: %v", err), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"project": res.Project,
			"project_id": projectID,
		})
	}
}

func (app *App) handleCreateMilestone(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		projectID := vars["projectId"]

		var req struct {
			Creator     string `json:"creator"`
			Title       string `json:"title"`
			Description string `json:"description"`
			Amount      string `json:"amount"`
			Deadline    string `json:"deadline"`
		}

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		// Parse amount
		amount, err := types.ParseCoinNormalized(req.Amount)
		if err != nil {
			http.Error(w, "Invalid amount format", http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "Milestone creation transaction prepared",
			"tx_info": map[string]interface{}{
				"project_id": projectID,
				"creator": req.Creator,
				"title": req.Title,
				"description": req.Description,
				"amount": amount.String(),
				"deadline": req.Deadline,
				"estimated_gas": "180000",
				"note": "Use skillchaind tx marketplace create-milestone to execute",
			},
		})
	}
}

func (app *App) handleListMilestones(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		projectID := vars["projectId"]

		queryClient := marketplacetypes.NewQueryClient(clientCtx)
		req := &marketplacetypes.QueryAllMilestoneRequest{
			Pagination: &query.PageRequest{
				Limit: 100,
			},
		}

		res, err := queryClient.MilestoneAll(r.Context(), req)
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to query milestones: %v", err), http.StatusInternalServerError)
			return
		}

		// Filter milestones by project ID
		var projectMilestones []interface{}
		for _, milestone := range res.Milestone {
			if milestone.ProjectId == projectID {
				projectMilestones = append(projectMilestones, milestone)
			}
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"milestones": projectMilestones,
			"project_id": projectID,
			"count": len(projectMilestones),
		})
	}
}

func (app *App) handleGetMilestone(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		milestoneID := vars["milestoneId"]

		queryClient := marketplacetypes.NewQueryClient(clientCtx)
		req := &marketplacetypes.QueryGetMilestoneRequest{
			Index: milestoneID,
		}

		res, err := queryClient.Milestone(r.Context(), req)
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to query milestone: %v", err), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"milestone": res.Milestone,
			"milestone_id": milestoneID,
		})
	}
}

func (app *App) handleCompleteMilestone(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		milestoneID := vars["milestoneId"]

		var req struct {
			Creator     string `json:"creator"`
			Deliverable string `json:"deliverable"`
			Notes       string `json:"notes,omitempty"`
		}

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "Milestone completion transaction prepared",
			"tx_info": map[string]interface{}{
				"milestone_id": milestoneID,
				"creator": req.Creator,
				"deliverable": req.Deliverable,
				"notes": req.Notes,
				"estimated_gas": "150000",
				"note": "Use skillchaind tx marketplace complete-milestone to execute",
			},
		})
	}
}

func (app *App) handleApproveMilestone(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		milestoneID := vars["milestoneId"]

		var req struct {
			Creator  string `json:"creator"`
			Feedback string `json:"feedback,omitempty"`
		}

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "Milestone approval transaction prepared",
			"tx_info": map[string]interface{}{
				"milestone_id": milestoneID,
				"creator": req.Creator,
				"feedback": req.Feedback,
				"estimated_gas": "120000",
				"note": "Use skillchaind tx marketplace approve-milestone to execute",
			},
		})
	}
}

func (app *App) handleReleasePayment(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		milestoneID := vars["milestoneId"]

		var req struct {
			Creator string `json:"creator"`
			Amount  string `json:"amount"`
		}

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		// Parse amount
		amount, err := types.ParseCoinNormalized(req.Amount)
		if err != nil {
			http.Error(w, "Invalid amount format", http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "Payment release transaction prepared",
			"tx_info": map[string]interface{}{
				"milestone_id": milestoneID,
				"creator": req.Creator,
				"amount": amount.String(),
				"estimated_gas": "180000",
				"note": "Use skillchaind tx marketplace release-payment to execute",
			},
		})
	}
} 