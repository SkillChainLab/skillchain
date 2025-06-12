package app

import (
	"net/http"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// RegisterCustomAPIRoutes registers custom API routes for SkillChain modules
func (app *App) RegisterCustomAPIRoutes(router *mux.Router, clientCtx client.Context) {
	// Enable CORS for all routes
	corsHandler := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"*"}),
	)
	
	// Create skillchain API subrouter (versioned only)
	skillchainRouter := router.PathPrefix("/skillchain/v1").Subrouter()
	
	// Add CORS middleware to router
	skillchainRouter.Use(func(next http.Handler) http.Handler {
		return corsHandler(next)
	})

	// SkillChain Core Module Routes
	skillchainRouter.HandleFunc("/params", app.handleGetParams(clientCtx)).Methods("GET")
	skillchainRouter.HandleFunc("/token/info", app.handleGetTokenInfo(clientCtx)).Methods("GET")
	skillchainRouter.HandleFunc("/vusd/treasury", app.handleGetVUSDTreasury(clientCtx)).Methods("GET")
	skillchainRouter.HandleFunc("/vusd/position/{address}", app.handleGetUserVUSDPosition(clientCtx)).Methods("GET")
	skillchainRouter.HandleFunc("/vusd/convert/skill-to-vusd", app.handleConvertSKILLToVUSD(clientCtx)).Methods("POST")
	skillchainRouter.HandleFunc("/vusd/convert/vusd-to-skill", app.handleConvertVUSDToSKILL(clientCtx)).Methods("POST")
	skillchainRouter.HandleFunc("/tokens/burn", app.handleBurnTokens(clientCtx)).Methods("POST")

	// Bank Module Routes (for wallet integration)
	skillchainRouter.HandleFunc("/bank/balances/{address}", app.handleGetBalances(clientCtx)).Methods("GET")
	skillchainRouter.HandleFunc("/bank/supply", app.handleGetSupply(clientCtx)).Methods("GET")
	skillchainRouter.HandleFunc("/bank/supply/{denom}", app.handleGetSupplyByDenom(clientCtx)).Methods("GET")

	// Profile Module Routes  
	skillchainRouter.HandleFunc("/profiles", app.handleCreateProfile(clientCtx)).Methods("POST")
	skillchainRouter.HandleFunc("/profiles", app.handleListProfiles(clientCtx)).Methods("GET")
	skillchainRouter.HandleFunc("/profiles/{address}", app.handleGetProfile(clientCtx)).Methods("GET")
	skillchainRouter.HandleFunc("/profiles/{address}", app.handleUpdateProfile(clientCtx)).Methods("PUT")
	
	skillchainRouter.HandleFunc("/profiles/{address}/skills", app.handleAddSkill(clientCtx)).Methods("POST")
	skillchainRouter.HandleFunc("/profiles/{address}/skills", app.handleListUserSkills(clientCtx)).Methods("GET")
	skillchainRouter.HandleFunc("/profiles/{address}/skills/{skillId}", app.handleGetSkill(clientCtx)).Methods("GET")
	skillchainRouter.HandleFunc("/profiles/{address}/skills/{skillId}/endorse", app.handleEndorseSkill(clientCtx)).Methods("POST")
	
	skillchainRouter.HandleFunc("/endorsements", app.handleListEndorsements(clientCtx)).Methods("GET")
	skillchainRouter.HandleFunc("/endorsements/{endorsementId}", app.handleGetEndorsement(clientCtx)).Methods("GET")

	// Marketplace Module Routes
	skillchainRouter.HandleFunc("/jobs", app.handleCreateJobPosting(clientCtx)).Methods("POST")
	skillchainRouter.HandleFunc("/jobs", app.handleListJobPostings(clientCtx)).Methods("GET")
	skillchainRouter.HandleFunc("/jobs/{jobId}", app.handleGetJobPosting(clientCtx)).Methods("GET")
	skillchainRouter.HandleFunc("/jobs/{jobId}", app.handleUpdateJobPosting(clientCtx)).Methods("PUT")
	skillchainRouter.HandleFunc("/jobs/{jobId}/close", app.handleCloseJobPosting(clientCtx)).Methods("POST")
	
	skillchainRouter.HandleFunc("/jobs/{jobId}/proposals", app.handleSubmitProposal(clientCtx)).Methods("POST")
	skillchainRouter.HandleFunc("/jobs/{jobId}/proposals", app.handleListProposals(clientCtx)).Methods("GET")
	skillchainRouter.HandleFunc("/proposals/{proposalId}", app.handleGetProposal(clientCtx)).Methods("GET")
	skillchainRouter.HandleFunc("/proposals/{proposalId}/accept", app.handleAcceptProposal(clientCtx)).Methods("POST")
	skillchainRouter.HandleFunc("/proposals/{proposalId}/reject", app.handleRejectProposal(clientCtx)).Methods("POST")
	
	skillchainRouter.HandleFunc("/projects", app.handleCreateProject(clientCtx)).Methods("POST")
	skillchainRouter.HandleFunc("/projects", app.handleListProjects(clientCtx)).Methods("GET")
	skillchainRouter.HandleFunc("/projects/{projectId}", app.handleGetProject(clientCtx)).Methods("GET")
	skillchainRouter.HandleFunc("/projects/{projectId}/milestones", app.handleCreateMilestone(clientCtx)).Methods("POST")
	skillchainRouter.HandleFunc("/projects/{projectId}/milestones", app.handleListMilestones(clientCtx)).Methods("GET")
	skillchainRouter.HandleFunc("/milestones/{milestoneId}", app.handleGetMilestone(clientCtx)).Methods("GET")
	skillchainRouter.HandleFunc("/milestones/{milestoneId}/complete", app.handleCompleteMilestone(clientCtx)).Methods("POST")
	skillchainRouter.HandleFunc("/milestones/{milestoneId}/approve", app.handleApproveMilestone(clientCtx)).Methods("POST")
	skillchainRouter.HandleFunc("/milestones/{milestoneId}/release-payment", app.handleReleasePayment(clientCtx)).Methods("POST")

	// Analytics Module Routes
	skillchainRouter.HandleFunc("/analytics/activity", app.handleTrackActivity(clientCtx)).Methods("POST")
	skillchainRouter.HandleFunc("/analytics/activities", app.handleListActivities(clientCtx)).Methods("GET")
	skillchainRouter.HandleFunc("/analytics/activities/{activityId}", app.handleGetActivity(clientCtx)).Methods("GET")
	skillchainRouter.HandleFunc("/analytics/users/{address}/activity", app.handleGetUserActivity(clientCtx)).Methods("GET")
	skillchainRouter.HandleFunc("/analytics/metrics", app.handleRecordMetric(clientCtx)).Methods("POST")
	skillchainRouter.HandleFunc("/analytics/metrics", app.handleListMetrics(clientCtx)).Methods("GET")
	skillchainRouter.HandleFunc("/analytics/metrics/{metricName}", app.handleGetMetric(clientCtx)).Methods("GET")
	skillchainRouter.HandleFunc("/analytics/reports/users", app.handleUserReport(clientCtx)).Methods("GET")
	skillchainRouter.HandleFunc("/analytics/reports/platform", app.handlePlatformReport(clientCtx)).Methods("GET")
	skillchainRouter.HandleFunc("/analytics/reports/revenue", app.handleRevenueReport(clientCtx)).Methods("GET")

	// File Storage Module Routes
	skillchainRouter.HandleFunc("/files", app.handleUploadFile(clientCtx)).Methods("POST")
	skillchainRouter.HandleFunc("/files", app.handleListFiles(clientCtx)).Methods("GET")
	skillchainRouter.HandleFunc("/files/{fileId}", app.handleGetFile(clientCtx)).Methods("GET")
	skillchainRouter.HandleFunc("/files/{fileId}", app.handleUpdateFile(clientCtx)).Methods("PUT")
	skillchainRouter.HandleFunc("/files/{fileId}", app.handleDeleteFile(clientCtx)).Methods("DELETE")
	skillchainRouter.HandleFunc("/files/{fileId}/download", app.handleDownloadFile(clientCtx)).Methods("GET")
	skillchainRouter.HandleFunc("/files/{fileId}/permissions", app.handleGrantFilePermission(clientCtx)).Methods("POST")
	skillchainRouter.HandleFunc("/files/{fileId}/permissions", app.handleListFilePermissions(clientCtx)).Methods("GET")
	skillchainRouter.HandleFunc("/files/{fileId}/permissions/{granteeAddress}", app.handleGetFilePermission(clientCtx)).Methods("GET")
	skillchainRouter.HandleFunc("/files/{fileId}/permissions/{granteeAddress}", app.handleRevokeFilePermission(clientCtx)).Methods("DELETE")
	skillchainRouter.HandleFunc("/ipfs/{hash}/pin", app.handlePinIPFS(clientCtx)).Methods("POST")
	skillchainRouter.HandleFunc("/ipfs/{hash}/unpin", app.handleUnpinIPFS(clientCtx)).Methods("POST")
	skillchainRouter.HandleFunc("/ipfs/{hash}/status", app.handleIPFSStatus(clientCtx)).Methods("GET")

	// Notifications Module Routes
	skillchainRouter.HandleFunc("/notifications", app.handleCreateNotification(clientCtx)).Methods("POST")
	skillchainRouter.HandleFunc("/notifications", app.handleListNotifications(clientCtx)).Methods("GET")
	skillchainRouter.HandleFunc("/notifications/{notificationId}", app.handleGetNotification(clientCtx)).Methods("GET")
	skillchainRouter.HandleFunc("/notifications/{notificationId}/read", app.handleMarkAsRead(clientCtx)).Methods("POST")
	skillchainRouter.HandleFunc("/notifications/{notificationId}", app.handleDeleteNotification(clientCtx)).Methods("DELETE")
	skillchainRouter.HandleFunc("/users/{address}/notifications", app.handleGetUserNotifications(clientCtx)).Methods("GET")
	skillchainRouter.HandleFunc("/users/{address}/notifications/unread", app.handleGetUnreadNotifications(clientCtx)).Methods("GET")
	skillchainRouter.HandleFunc("/users/{address}/notifications/mark-all-read", app.handleMarkAllAsRead(clientCtx)).Methods("POST")
	skillchainRouter.HandleFunc("/users/{address}/notification-preferences", app.handleGetNotificationPreferences(clientCtx)).Methods("GET")
	skillchainRouter.HandleFunc("/users/{address}/notification-preferences", app.handleUpdateNotificationPreferences(clientCtx)).Methods("PUT")
	skillchainRouter.HandleFunc("/push/subscribe", app.handleSubscribePush(clientCtx)).Methods("POST")
	skillchainRouter.HandleFunc("/push/unsubscribe", app.handleUnsubscribePush(clientCtx)).Methods("POST")
	skillchainRouter.HandleFunc("/push/send", app.handleSendPushNotification(clientCtx)).Methods("POST")
} 