package app

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/gorilla/mux"
	
	analyticstypes "skillchain/x/analytics/types"
	filestoragetypes "skillchain/x/filestorage/types"
	notificationstypes "skillchain/x/notifications/types"
)

// Analytics Module Handlers

func (app *App) handleTrackActivity(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			Creator     string            `json:"creator"`
			ActivityType string           `json:"activity_type"`
			UserAddress  string           `json:"user_address,omitempty"`
			IPAddress    string           `json:"ip_address,omitempty"`
			Metadata     map[string]string `json:"metadata,omitempty"`
		}

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "Activity tracking transaction prepared",
			"tx_info": map[string]interface{}{
				"creator": req.Creator,
				"activity_type": req.ActivityType,
				"user_address": req.UserAddress,
				"ip_address": req.IPAddress,
				"metadata": req.Metadata,
				"estimated_gas": "120000",
				"note": "Use skillchaind tx analytics track-activity to execute",
			},
		})
	}
}

func (app *App) handleListActivities(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		queryClient := analyticstypes.NewQueryClient(clientCtx)
		
		req := &analyticstypes.QueryAllUserActivityRequest{
			Pagination: &query.PageRequest{Limit: 50},
		}

		res, err := queryClient.UserActivityAll(r.Context(), req)
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to query activities: %v", err), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"activities": res.UserActivity,
			"count": len(res.UserActivity),
		})
	}
}

func (app *App) handleGetActivity(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		activityID := vars["activityId"]

		queryClient := analyticstypes.NewQueryClient(clientCtx)
		req := &analyticstypes.QueryGetUserActivityRequest{Index: activityID}

		res, err := queryClient.UserActivity(r.Context(), req)
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to query activity: %v", err), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"activity": res.UserActivity,
			"activity_id": activityID,
		})
	}
}

func (app *App) handleGetUserActivity(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		address := vars["address"]

		if _, err := types.AccAddressFromBech32(address); err != nil {
			http.Error(w, "Invalid address format", http.StatusBadRequest)
			return
		}

		queryClient := analyticstypes.NewQueryClient(clientCtx)
		req := &analyticstypes.QueryAllUserActivityRequest{
			Pagination: &query.PageRequest{Limit: 100},
		}

		res, err := queryClient.UserActivityAll(r.Context(), req)
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to query activities: %v", err), http.StatusInternalServerError)
			return
		}

		// Filter by user address
		var userActivities []interface{}
		for _, activity := range res.UserActivity {
			if activity.UserAddress == address {
				userActivities = append(userActivities, activity)
			}
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"activities": userActivities,
			"address": address,
			"count": len(userActivities),
		})
	}
}

func (app *App) handleRecordMetric(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			Creator    string  `json:"creator"`
			MetricName string  `json:"metric_name"`
			Value      float64 `json:"value"`
			Category   string  `json:"category,omitempty"`
		}

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "Metric recording transaction prepared",
			"tx_info": req,
			"estimated_gas": "100000",
		})
	}
}

func (app *App) handleListMetrics(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		queryClient := analyticstypes.NewQueryClient(clientCtx)
		
		req := &analyticstypes.QueryAllPlatformMetricRequest{
			Pagination: &query.PageRequest{Limit: 50},
		}

		res, err := queryClient.PlatformMetricAll(r.Context(), req)
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to query metrics: %v", err), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"metrics": res.PlatformMetric,
			"count": len(res.PlatformMetric),
		})
	}
}

func (app *App) handleGetMetric(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		metricName := vars["metricName"]

		queryClient := analyticstypes.NewQueryClient(clientCtx)
		req := &analyticstypes.QueryGetPlatformMetricRequest{Index: metricName}

		res, err := queryClient.PlatformMetric(r.Context(), req)
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to query metric: %v", err), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"metric": res.PlatformMetric,
			"metric_name": metricName,
		})
	}
}

func (app *App) handleUserReport(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "User analytics report",
			"total_users": "N/A - implement with actual user count query",
			"active_users": "N/A - implement with active user metrics",
			"timestamp": r.Header.Get("X-Request-Time"),
		})
	}
}

func (app *App) handlePlatformReport(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "Platform analytics report",
			"note": "Implement with aggregated platform metrics",
			"timestamp": r.Header.Get("X-Request-Time"),
		})
	}
}

func (app *App) handleRevenueReport(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "Revenue analytics report",
			"note": "Implement with transaction volume data",
			"timestamp": r.Header.Get("X-Request-Time"),
		})
	}
}

// File Storage Module Handlers

func (app *App) handleUploadFile(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			Creator     string `json:"creator"`
			FileName    string `json:"file_name"`
			FileSize    int64  `json:"file_size"`
			FileType    string `json:"file_type"`
			IPFSHash    string `json:"ipfs_hash"`
			Description string `json:"description,omitempty"`
		}

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "File upload transaction prepared",
			"tx_info": req,
			"estimated_gas": "180000",
			"note": "Use skillchaind tx filestorage upload-file to execute",
		})
	}
}

func (app *App) handleListFiles(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		queryClient := filestoragetypes.NewQueryClient(clientCtx)
		
		req := &filestoragetypes.QueryAllFileRecordRequest{
			Pagination: &query.PageRequest{Limit: 50},
		}

		res, err := queryClient.FileRecordAll(r.Context(), req)
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to query files: %v", err), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"files": res.FileRecord,
			"count": len(res.FileRecord),
		})
	}
}

func (app *App) handleGetFile(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		fileID := vars["fileId"]

		queryClient := filestoragetypes.NewQueryClient(clientCtx)
		req := &filestoragetypes.QueryGetFileRecordRequest{Index: fileID}

		res, err := queryClient.FileRecord(r.Context(), req)
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to query file: %v", err), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"file": res.FileRecord,
			"file_id": fileID,
		})
	}
}

func (app *App) handleUpdateFile(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		fileID := vars["fileId"]

		var req struct {
			Creator     string `json:"creator"`
			Description string `json:"description,omitempty"`
		}

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "File update transaction prepared",
			"file_id": fileID,
			"updates": req,
			"estimated_gas": "120000",
		})
	}
}

func (app *App) handleDeleteFile(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		fileID := vars["fileId"]

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "File deletion transaction prepared",
			"file_id": fileID,
			"estimated_gas": "100000",
			"note": "Use skillchaind tx filestorage delete-file to execute",
		})
	}
}

func (app *App) handleDownloadFile(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		fileID := vars["fileId"]

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "File download info",
			"file_id": fileID,
			"note": "Implement IPFS gateway integration for actual download",
		})
	}
}

func (app *App) handleGrantFilePermission(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		fileID := vars["fileId"]

		var req struct {
			Creator        string `json:"creator"`
			GranteeAddress string `json:"grantee_address"`
			PermissionType string `json:"permission_type"`
		}

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "File permission grant transaction prepared",
			"file_id": fileID,
			"tx_info": req,
			"estimated_gas": "150000",
		})
	}
}

func (app *App) handleListFilePermissions(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		fileID := vars["fileId"]

		queryClient := filestoragetypes.NewQueryClient(clientCtx)
		req := &filestoragetypes.QueryAllFilePermissionRequest{
			Pagination: &query.PageRequest{Limit: 100},
		}

		res, err := queryClient.FilePermissionAll(r.Context(), req)
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to query file permissions: %v", err), http.StatusInternalServerError)
			return
		}

		// Filter by file ID
		var filePermissions []interface{}
		for _, permission := range res.FilePermission {
			if permission.FileId == fileID {
				filePermissions = append(filePermissions, permission)
			}
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"permissions": filePermissions,
			"file_id": fileID,
			"count": len(filePermissions),
		})
	}
}

func (app *App) handleGetFilePermission(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		fileID := vars["fileId"]
		granteeAddr := vars["granteeAddress"]

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "File permission details",
			"file_id": fileID,
			"grantee_address": granteeAddr,
			"note": "Implement specific permission lookup",
		})
	}
}

func (app *App) handleRevokeFilePermission(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		fileID := vars["fileId"]
		granteeAddr := vars["granteeAddress"]

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "File permission revocation transaction prepared",
			"file_id": fileID,
			"grantee_address": granteeAddr,
			"estimated_gas": "120000",
		})
	}
}

func (app *App) handlePinIPFS(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		hash := vars["hash"]

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "IPFS pin operation",
			"hash": hash,
			"note": "Implement IPFS pinning service integration",
		})
	}
}

func (app *App) handleUnpinIPFS(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		hash := vars["hash"]

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "IPFS unpin operation",
			"hash": hash,
			"note": "Implement IPFS unpinning service integration",
		})
	}
}

func (app *App) handleIPFSStatus(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		hash := vars["hash"]

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "IPFS status check",
			"hash": hash,
			"status": "pinned", // Mock status
			"note": "Implement actual IPFS status checking",
		})
	}
}

// Notifications Module Handlers

func (app *App) handleCreateNotification(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			Creator     string `json:"creator"`
			Title       string `json:"title"`
			Message     string `json:"message"`
			Priority    string `json:"priority"`
			UserAddress string `json:"user_address,omitempty"`
		}

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "Notification creation transaction prepared",
			"tx_info": req,
			"estimated_gas": "120000",
			"note": "Use skillchaind tx notifications create-notification to execute",
		})
	}
}

func (app *App) handleListNotifications(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		queryClient := notificationstypes.NewQueryClient(clientCtx)
		
		req := &notificationstypes.QueryAllNotificationRequest{
			Pagination: &query.PageRequest{Limit: 50},
		}

		res, err := queryClient.NotificationAll(r.Context(), req)
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to query notifications: %v", err), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"notifications": res.Notification,
			"count": len(res.Notification),
		})
	}
}

func (app *App) handleGetNotification(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		notificationID := vars["notificationId"]

		queryClient := notificationstypes.NewQueryClient(clientCtx)
		req := &notificationstypes.QueryGetNotificationRequest{Index: notificationID}

		res, err := queryClient.Notification(r.Context(), req)
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to query notification: %v", err), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"notification": res.Notification,
			"notification_id": notificationID,
		})
	}
}

func (app *App) handleMarkAsRead(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		notificationID := vars["notificationId"]

		var req struct {
			Creator string `json:"creator"`
		}

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "Mark notification as read transaction prepared",
			"notification_id": notificationID,
			"creator": req.Creator,
			"estimated_gas": "80000",
		})
	}
}

func (app *App) handleDeleteNotification(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		notificationID := vars["notificationId"]

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "Notification deletion transaction prepared",
			"notification_id": notificationID,
			"estimated_gas": "100000",
		})
	}
}

func (app *App) handleGetUserNotifications(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		address := vars["address"]

		if _, err := types.AccAddressFromBech32(address); err != nil {
			http.Error(w, "Invalid address format", http.StatusBadRequest)
			return
		}

		queryClient := notificationstypes.NewQueryClient(clientCtx)
		req := &notificationstypes.QueryAllNotificationRequest{
			Pagination: &query.PageRequest{Limit: 100},
		}

		res, err := queryClient.NotificationAll(r.Context(), req)
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to query notifications: %v", err), http.StatusInternalServerError)
			return
		}

		// Filter notifications by user address
		var userNotifications []interface{}
		for _, notification := range res.Notification {
			if notification.UserAddress == address {
				userNotifications = append(userNotifications, notification)
			}
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"notifications": userNotifications,
			"address": address,
			"count": len(userNotifications),
		})
	}
}

func (app *App) handleGetUnreadNotifications(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		address := vars["address"]

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "Unread notifications for user",
			"address": address,
			"note": "Implement filtering for unread notifications",
		})
	}
}

func (app *App) handleMarkAllAsRead(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		address := vars["address"]

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "Mark all notifications as read transaction prepared",
			"address": address,
			"estimated_gas": "150000",
		})
	}
}

func (app *App) handleGetNotificationPreferences(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		address := vars["address"]

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "User notification preferences",
			"address": address,
			"note": "Implement notification preferences storage and retrieval",
		})
	}
}

func (app *App) handleUpdateNotificationPreferences(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		address := vars["address"]

		var req struct {
			Creator     string            `json:"creator"`
			Preferences map[string]bool   `json:"preferences"`
		}

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "Notification preferences update transaction prepared",
			"address": address,
			"preferences": req.Preferences,
			"estimated_gas": "120000",
		})
	}
}

func (app *App) handleSubscribePush(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			Address  string `json:"address"`
			Endpoint string `json:"endpoint"`
			Keys     map[string]string `json:"keys"`
		}

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "Push notification subscription registered",
			"address": req.Address,
			"endpoint": req.Endpoint,
		})
	}
}

func (app *App) handleUnsubscribePush(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			Address  string `json:"address"`
			Endpoint string `json:"endpoint"`
		}

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "Push notification subscription removed",
			"address": req.Address,
			"endpoint": req.Endpoint,
		})
	}
}

func (app *App) handleSendPushNotification(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			Creator   string   `json:"creator"`
			Recipients []string `json:"recipients"`
			Title      string   `json:"title"`
			Body       string   `json:"body"`
			Data       map[string]string `json:"data,omitempty"`
		}

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "Push notification sent",
			"creator": req.Creator,
			"recipients": req.Recipients,
			"title": req.Title,
			"body": req.Body,
		})
	}
} 