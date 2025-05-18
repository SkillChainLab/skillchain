package keeper

import (
	"context"
	"fmt"

	"github.com/SkillChainLab/skillchain/x/job/types"
)

var _ types.QueryServer = Keeper{}

func (k Keeper) ListJobApplications(ctx context.Context, req *types.QueryListJobApplicationsRequest) (*types.QueryListJobApplicationsResponse, error) {
	store := k.getStore(ctx)
	applications := []*types.Application{}

	// Create the prefix for the job ID
	prefix := fmt.Sprintf("Application/value/%d:", req.JobId)
	iterator := store.Iterator([]byte(prefix), nil)
	defer iterator.Close()

	// Get job details first
	job, found := k.GetJob(ctx, req.JobId)
	if !found {
		return nil, fmt.Errorf("job %d not found", req.JobId)
	}

	for ; iterator.Valid(); iterator.Next() {
		var application types.Application
		if err := k.cdc.Unmarshal(iterator.Value(), &application); err != nil {
			return nil, fmt.Errorf("failed to unmarshal application: %w", err)
		}

		// Ensure application has correct status
		if application.Status == "" {
			application.Status = "PENDING"
		}

		// Add job details to the application
		application.JobTitle = job.Title
		application.JobDescription = job.Description
		application.JobBudget = job.Budget

		applications = append(applications, &application)
	}

	return &types.QueryListJobApplicationsResponse{
		Applications: applications,
	}, nil
}
