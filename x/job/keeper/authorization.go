package keeper

import (
	"context"

	"cosmossdk.io/errors"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/SkillChainLab/skillchain/x/job/types"
)

// Authorization errors
var (
	ErrUnauthorizedJobUpdate    = errors.Register(types.ModuleName, 1001, "unauthorized to update job")
	ErrUnauthorizedJobReview    = errors.Register(types.ModuleName, 1002, "unauthorized to review job application")
	ErrUnauthorizedJobDelete    = errors.Register(types.ModuleName, 1003, "unauthorized to delete job")
	ErrUnauthorizedJobView      = errors.Register(types.ModuleName, 1004, "unauthorized to view job details")
	ErrUnauthorizedApplication  = errors.Register(types.ModuleName, 1005, "unauthorized to submit application")
	ErrUnauthorizedNotification = errors.Register(types.ModuleName, 1006, "unauthorized to access notification")
)

// Authorization levels
const (
	AuthLevelNone     = 0
	AuthLevelBasic    = 1
	AuthLevelVerified = 2
	AuthLevelAdmin    = 3
)

// Authorization roles
const (
	RoleNone      = "none"
	RoleUser      = "user"
	RoleVerified  = "verified"
	RoleInstitution = "institution"
	RoleAdmin     = "admin"
)

// Authorization checks for job operations
func (k Keeper) CheckJobUpdateAuthorization(ctx context.Context, jobId uint64, updater string) error {
	job, found := k.GetJob(ctx, jobId)
	if !found {
		return errors.Wrapf(sdkerrors.ErrNotFound, "job %d not found", jobId)
	}

	if updater != job.Creator {
		return errors.Wrapf(ErrUnauthorizedJobUpdate, "only job creator can update job")
	}

	return nil
}

func (k Keeper) CheckJobReviewAuthorization(ctx context.Context, jobId uint64, reviewer string) error {
	job, found := k.GetJob(ctx, jobId)
	if !found {
		return errors.Wrapf(sdkerrors.ErrNotFound, "job %d not found", jobId)
	}

	if reviewer != job.Creator {
		return errors.Wrapf(ErrUnauthorizedJobReview, "only job creator can review applications")
	}

	return nil
}

func (k Keeper) CheckJobDeleteAuthorization(ctx context.Context, jobId uint64, deleter string) error {
	job, found := k.GetJob(ctx, jobId)
	if !found {
		return errors.Wrapf(sdkerrors.ErrNotFound, "job %d not found", jobId)
	}

	if deleter != job.Creator {
		return errors.Wrapf(ErrUnauthorizedJobDelete, "only job creator can delete job")
	}

	return nil
}

func (k Keeper) CheckJobViewAuthorization(ctx context.Context, jobId uint64, viewer string) error {
	_, found := k.GetJob(ctx, jobId)
	if !found {
		return errors.Wrapf(sdkerrors.ErrNotFound, "job %d not found", jobId)
	}

	// Anyone can view jobs (public/private ayrımı yoksa bu şekilde bırakıyoruz)
	return nil
}

// Authorization checks for application operations
func (k Keeper) CheckApplicationAuthorization(ctx context.Context, jobId uint64, applicant string) error {
	job, found := k.GetJob(ctx, jobId)
	if !found {
		return errors.Wrapf(sdkerrors.ErrNotFound, "job %d not found", jobId)
	}

	// Check if applicant is not the job creator
	if applicant == job.Creator {
		return errors.Wrapf(ErrUnauthorizedApplication, "job creator cannot apply to their own job")
	}

	// Check if applicant has already applied
	_, found = k.GetJobApplication(ctx, jobId, applicant)
	if found {
		return errors.Wrapf(ErrUnauthorizedApplication, "already applied to this job")
	}

	return nil
}

// Authorization checks for notification operations
func (k Keeper) CheckNotificationAuthorization(ctx context.Context, notificationId string, user string) error {
	notification, err := k.GetNotificationByID(ctx, notificationId)
	if err != nil {
		return err
	}

	if notification.Recipient != user && notification.Sender != user {
		return errors.Wrapf(ErrUnauthorizedNotification, "unauthorized to access notification")
	}

	return nil
}

// Get user authorization level
func (k Keeper) GetUserAuthLevel(ctx context.Context, user string) int {
	// TODO: Implement user verification status check
	// For now, return basic level for all users
	return AuthLevelBasic
}

// Get user role
func (k Keeper) GetUserRole(ctx context.Context, user string) string {
	// TODO: Implement user role check
	// For now, return basic user role for all users
	return RoleUser
}

// Check if user has required authorization level
func (k Keeper) HasRequiredAuthLevel(ctx context.Context, user string, requiredLevel int) bool {
	userLevel := k.GetUserAuthLevel(ctx, user)
	return userLevel >= requiredLevel
}

// Check if user has required role
func (k Keeper) HasRequiredRole(ctx context.Context, user string, requiredRole string) bool {
	userRole := k.GetUserRole(ctx, user)
	// Role hierarchy
	roleHierarchy := map[string]int{
		RoleNone:       0,
		RoleUser:       1,
		RoleVerified:   2,
		RoleInstitution: 3,
		RoleAdmin:      4,
	}
	return roleHierarchy[userRole] >= roleHierarchy[requiredRole]
} 