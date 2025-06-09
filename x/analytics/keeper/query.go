package keeper

import (
	"sort"
	"strings"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"skillchain/x/analytics/types"
)

var _ types.QueryServer = Keeper{}

// GetUserActivityAnalytics returns comprehensive user activity analytics
func (k Keeper) GetUserActivityAnalytics(ctx sdk.Context, userAddress string, timeRange uint64) map[string]interface{} {
	currentTime := uint64(time.Now().Unix())
	startTime := currentTime - timeRange
	
	userActivities := k.GetUserActivitiesByDateRange(ctx, userAddress, startTime, currentTime)
	
	analytics := make(map[string]interface{})
	
	// Basic stats
	analytics["total_activities"] = len(userActivities)
	analytics["time_range_hours"] = timeRange / 3600
	analytics["activities_per_hour"] = float64(len(userActivities)) / float64(timeRange/3600)
	
	// Activity type breakdown
	activityTypes := make(map[string]int)
	actionTypes := make(map[string]int)
	hourlyDistribution := make(map[int]int)
	
	for _, activity := range userActivities {
		activityTypes[activity.ActivityType]++
		actionTypes[activity.Action]++
		
		// Hourly distribution
		hour := int(activity.Timestamp % 86400 / 3600)
		hourlyDistribution[hour]++
	}
	
	analytics["activity_types"] = activityTypes
	analytics["action_types"] = actionTypes
	analytics["hourly_distribution"] = hourlyDistribution
	analytics["most_active_hour"] = k.findMostActiveHour(hourlyDistribution)
	analytics["user_engagement_score"] = k.calculateEngagementScore(userActivities, timeRange)
	
	return analytics
}

// GetRevenueAnalytics returns comprehensive revenue analytics
func (k Keeper) GetRevenueAnalytics(ctx sdk.Context, currency string, timeRange uint64) map[string]interface{} {
	currentTime := uint64(time.Now().Unix())
	startTime := currentTime - timeRange
	
	revenueRecords := k.GetRevenueRecordsByDateRange(ctx, currency, startTime, currentTime)
	
	analytics := make(map[string]interface{})
	
	// Basic financial metrics
	var totalRevenue, totalFees, totalPlatformFees uint64
	transactionTypes := make(map[string]uint64)
	dailyRevenue := make(map[string]uint64)
	
	for _, record := range revenueRecords {
		totalRevenue += record.Amount
		totalFees += record.FeeAmount
		totalPlatformFees += record.PlatformFee
		
		transactionTypes[record.TransactionType] += record.Amount
		
		// Daily breakdown
		day := time.Unix(int64(record.Timestamp), 0).UTC().Format("2006-01-02")
		dailyRevenue[day] += record.Amount
	}
	
	netRevenue := totalRevenue - totalFees - totalPlatformFees
	avgTransactionSize := uint64(0)
	if len(revenueRecords) > 0 {
		avgTransactionSize = totalRevenue / uint64(len(revenueRecords))
	}
	
	analytics["total_revenue"] = totalRevenue
	analytics["total_fees"] = totalFees
	analytics["total_platform_fees"] = totalPlatformFees
	analytics["net_revenue"] = netRevenue
	analytics["transaction_count"] = len(revenueRecords)
	analytics["avg_transaction_size"] = avgTransactionSize
	analytics["profit_margin"] = k.calculateOverallProfitMargin(totalRevenue, totalFees, totalPlatformFees)
	analytics["transaction_types"] = transactionTypes
	analytics["daily_revenue"] = dailyRevenue
	analytics["revenue_growth_rate"] = k.calculateRevenueGrowthRate(ctx, currency, timeRange)
	
	return analytics
}

// GetPlatformMetricsAnalytics returns platform-wide metrics analytics
func (k Keeper) GetPlatformMetricsAnalytics(ctx sdk.Context, timeRange uint64) map[string]interface{} {
	currentTime := uint64(time.Now().Unix())
	startTime := currentTime - timeRange
	
	platformMetrics := k.GetPlatformMetricsByDateRange(ctx, startTime, currentTime)
	
	analytics := make(map[string]interface{})
	
	// Metric type analysis
	metricTypes := make(map[string][]uint64)
	metricTrends := make(map[string]string)
	
	for _, metric := range platformMetrics {
		metricTypes[metric.MetricName] = append(metricTypes[metric.MetricName], metric.MetricValue)
	}
	
	// Calculate trends for each metric
	for metricName, values := range metricTypes {
		trend := k.calculateTrend(values)
		metricTrends[metricName] = trend
		
		// Add statistical analysis
		analytics[metricName+"_avg"] = k.calculateAverage(values)
		analytics[metricName+"_max"] = k.calculateMax(values)
		analytics[metricName+"_min"] = k.calculateMin(values)
		analytics[metricName+"_trend"] = trend
	}
	
	analytics["metric_trends"] = metricTrends
	analytics["total_metrics"] = len(platformMetrics)
	analytics["unique_metrics"] = len(metricTypes)
	
	return analytics
}

// GetUserBehaviorInsights returns advanced user behavior insights
func (k Keeper) GetUserBehaviorInsights(ctx sdk.Context, userAddress string) map[string]interface{} {
	allActivities := k.GetUserActivitiesByUser(ctx, userAddress)
	
	insights := make(map[string]interface{})
	
	if len(allActivities) == 0 {
		insights["status"] = "no_data"
		return insights
	}
	
	// Sort by timestamp
	sort.Slice(allActivities, func(i, j int) bool {
		return allActivities[i].Timestamp < allActivities[j].Timestamp
	})
	
	// User journey analysis
	journey := k.analyzeUserJourney(allActivities)
	insights["user_journey"] = journey
	
	// Session analysis
	sessions := k.analyzeSessions(allActivities)
	insights["session_analytics"] = sessions
	
	// Behavior patterns
	patterns := k.identifyBehaviorPatterns(allActivities)
	insights["behavior_patterns"] = patterns
	
	// User segmentation
	segment := k.classifyUserSegment(allActivities)
	insights["user_segment"] = segment
	
	// Retention metrics
	retention := k.calculateUserRetention(allActivities)
	insights["retention_metrics"] = retention
	
	return insights
}

// GetFinancialInsights returns advanced financial insights
func (k Keeper) GetFinancialInsights(ctx sdk.Context, currency string) map[string]interface{} {
	allRecords := k.GetRevenueRecordsByCurrency(ctx, currency)
	
	insights := make(map[string]interface{})
	
	if len(allRecords) == 0 {
		insights["status"] = "no_data"
		return insights
	}
	
	// Customer lifetime value analysis
	clv := k.calculateCustomerLifetimeValue(allRecords)
	insights["customer_lifetime_value"] = clv
	
	// Revenue concentration analysis
	concentration := k.analyzeRevenueConcentration(allRecords)
	insights["revenue_concentration"] = concentration
	
	// Fraud risk assessment
	fraudRisk := k.assessFraudRisk(allRecords)
	insights["fraud_risk_assessment"] = fraudRisk
	
	// Seasonal patterns
	seasonal := k.analyzeSeasonalPatterns(allRecords)
	insights["seasonal_patterns"] = seasonal
	
	// Revenue forecasting
	forecast := k.generateRevenueForecast(allRecords)
	insights["revenue_forecast"] = forecast
	
	return insights
}

// GetRealTimeMetrics returns real-time platform metrics
func (k Keeper) GetRealTimeMetrics(ctx sdk.Context) map[string]interface{} {
	currentTime := uint64(time.Now().Unix())
	lastHour := currentTime - 3600
	last24Hours := currentTime - 86400
	
	metrics := make(map[string]interface{})
	
	// Real-time activity metrics
	recentActivities := k.GetActivitiesByDateRange(ctx, lastHour, currentTime)
	dailyActivities := k.GetActivitiesByDateRange(ctx, last24Hours, currentTime)
	
	metrics["activities_last_hour"] = len(recentActivities)
	metrics["activities_last_24h"] = len(dailyActivities)
	metrics["activity_rate_per_hour"] = len(recentActivities)
	
	// Real-time revenue metrics
	recentRevenue := k.GetRevenueByDateRange(ctx, lastHour, currentTime)
	dailyRevenue := k.GetRevenueByDateRange(ctx, last24Hours, currentTime)
	
	metrics["revenue_last_hour"] = k.sumRevenueAmounts(recentRevenue)
	metrics["revenue_last_24h"] = k.sumRevenueAmounts(dailyRevenue)
	metrics["transactions_last_hour"] = len(recentRevenue)
	metrics["transactions_last_24h"] = len(dailyRevenue)
	
	// Active users metrics
	activeUsers := k.getActiveUsers(ctx, last24Hours, currentTime)
	metrics["active_users_24h"] = len(activeUsers)
	metrics["active_users_last_hour"] = len(k.getActiveUsers(ctx, lastHour, currentTime))
	
	// Platform health metrics
	health := k.calculatePlatformHealth(ctx)
	metrics["platform_health_score"] = health
	
	return metrics
}

// Helper functions for analytics calculations

func (k Keeper) GetUserActivitiesByDateRange(ctx sdk.Context, userAddress string, startTime, endTime uint64) []types.UserActivity {
	var activities []types.UserActivity
	allActivities := k.GetAllUserActivity(ctx)
	
	for _, activity := range allActivities {
		if activity.UserAddress == userAddress && activity.Timestamp >= startTime && activity.Timestamp <= endTime {
			activities = append(activities, activity)
		}
	}
	
	return activities
}

func (k Keeper) GetUserActivitiesByUser(ctx sdk.Context, userAddress string) []types.UserActivity {
	var activities []types.UserActivity
	allActivities := k.GetAllUserActivity(ctx)
	
	for _, activity := range allActivities {
		if activity.UserAddress == userAddress {
			activities = append(activities, activity)
		}
	}
	
	return activities
}

func (k Keeper) GetRevenueRecordsByDateRange(ctx sdk.Context, currency string, startTime, endTime uint64) []types.RevenueRecord {
	var records []types.RevenueRecord
	allRecords := k.GetAllRevenueRecord(ctx)
	
	for _, record := range allRecords {
		if (currency == "" || strings.EqualFold(record.Currency, currency)) && 
		   record.Timestamp >= startTime && record.Timestamp <= endTime {
			records = append(records, record)
		}
	}
	
	return records
}

func (k Keeper) GetRevenueRecordsByCurrency(ctx sdk.Context, currency string) []types.RevenueRecord {
	var records []types.RevenueRecord
	allRecords := k.GetAllRevenueRecord(ctx)
	
	for _, record := range allRecords {
		if strings.EqualFold(record.Currency, currency) {
			records = append(records, record)
		}
	}
	
	return records
}

func (k Keeper) GetPlatformMetricsByDateRange(ctx sdk.Context, startTime, endTime uint64) []types.PlatformMetric {
	var metrics []types.PlatformMetric
	allMetrics := k.GetAllPlatformMetric(ctx)
	
	for _, metric := range allMetrics {
		if metric.Timestamp >= startTime && metric.Timestamp <= endTime {
			metrics = append(metrics, metric)
		}
	}
	
	return metrics
}

func (k Keeper) GetActivitiesByDateRange(ctx sdk.Context, startTime, endTime uint64) []types.UserActivity {
	var activities []types.UserActivity
	allActivities := k.GetAllUserActivity(ctx)
	
	for _, activity := range allActivities {
		if activity.Timestamp >= startTime && activity.Timestamp <= endTime {
			activities = append(activities, activity)
		}
	}
	
	return activities
}

func (k Keeper) GetRevenueByDateRange(ctx sdk.Context, startTime, endTime uint64) []types.RevenueRecord {
	var records []types.RevenueRecord
	allRecords := k.GetAllRevenueRecord(ctx)
	
	for _, record := range allRecords {
		if record.Timestamp >= startTime && record.Timestamp <= endTime {
			records = append(records, record)
		}
	}
	
	return records
}

func (k Keeper) findMostActiveHour(hourlyDistribution map[int]int) int {
	maxActivity := 0
	mostActiveHour := 0
	
	for hour, count := range hourlyDistribution {
		if count > maxActivity {
			maxActivity = count
			mostActiveHour = hour
		}
	}
	
	return mostActiveHour
}

func (k Keeper) calculateEngagementScore(activities []types.UserActivity, timeRange uint64) float64 {
	if len(activities) == 0 || timeRange == 0 {
		return 0.0
	}
	
	// Base score on activity frequency and diversity
	activityTypes := make(map[string]bool)
	for _, activity := range activities {
		activityTypes[activity.ActivityType] = true
	}
	
	frequency := float64(len(activities)) / float64(timeRange/3600) // activities per hour
	diversity := float64(len(activityTypes))
	
	// Engagement score formula (can be refined)
	return (frequency * diversity) * 10.0
}

func (k Keeper) calculateOverallProfitMargin(totalRevenue, totalFees, totalPlatformFees uint64) float64 {
	if totalRevenue == 0 {
		return 0.0
	}
	profit := float64(totalFees + totalPlatformFees)
	return (profit / float64(totalRevenue)) * 100.0
}

func (k Keeper) calculateRevenueGrowthRate(ctx sdk.Context, currency string, timeRange uint64) float64 {
	currentTime := uint64(time.Now().Unix())
	
	// Current period
	currentPeriodStart := currentTime - timeRange
	currentRevenue := k.GetRevenueRecordsByDateRange(ctx, currency, currentPeriodStart, currentTime)
	
	// Previous period
	previousPeriodStart := currentPeriodStart - timeRange
	previousRevenue := k.GetRevenueRecordsByDateRange(ctx, currency, previousPeriodStart, currentPeriodStart)
	
	currentTotal := k.sumRevenueAmounts(currentRevenue)
	previousTotal := k.sumRevenueAmounts(previousRevenue)
	
	if previousTotal == 0 {
		return 0.0
	}
	
	return ((float64(currentTotal) - float64(previousTotal)) / float64(previousTotal)) * 100.0
}

func (k Keeper) calculateTrend(values []uint64) string {
	if len(values) < 2 {
		return "insufficient_data"
	}
	
	first := values[0]
	last := values[len(values)-1]
	
	if last > first {
		return "increasing"
	} else if last < first {
		return "decreasing"
	} else {
		return "stable"
	}
}

func (k Keeper) calculateAverage(values []uint64) float64 {
	if len(values) == 0 {
		return 0.0
	}
	
	var sum uint64
	for _, value := range values {
		sum += value
	}
	
	return float64(sum) / float64(len(values))
}

func (k Keeper) calculateMax(values []uint64) uint64 {
	if len(values) == 0 {
		return 0
	}
	
	max := values[0]
	for _, value := range values {
		if value > max {
			max = value
		}
	}
	
	return max
}

func (k Keeper) calculateMin(values []uint64) uint64 {
	if len(values) == 0 {
		return 0
	}
	
	min := values[0]
	for _, value := range values {
		if value < min {
			min = value
		}
	}
	
	return min
}

func (k Keeper) sumRevenueAmounts(records []types.RevenueRecord) uint64 {
	var total uint64
	for _, record := range records {
		total += record.Amount
	}
	return total
}

func (k Keeper) getActiveUsers(ctx sdk.Context, startTime, endTime uint64) []string {
	activities := k.GetActivitiesByDateRange(ctx, startTime, endTime)
	userSet := make(map[string]bool)
	
	for _, activity := range activities {
		userSet[activity.UserAddress] = true
	}
	
	users := make([]string, 0, len(userSet))
	for user := range userSet {
		users = append(users, user)
	}
	
	return users
}

func (k Keeper) calculatePlatformHealth(ctx sdk.Context) float64 {
	// Simple health score based on recent activity
	currentTime := uint64(time.Now().Unix())
	last24Hours := currentTime - 86400
	
	activities := k.GetActivitiesByDateRange(ctx, last24Hours, currentTime)
	revenue := k.GetRevenueByDateRange(ctx, last24Hours, currentTime)
	
	activityScore := float64(len(activities)) / 100.0 // Normalize to 0-1 scale
	revenueScore := float64(len(revenue)) / 50.0      // Normalize to 0-1 scale
	
	if activityScore > 1.0 {
		activityScore = 1.0
	}
	if revenueScore > 1.0 {
		revenueScore = 1.0
	}
	
	return (activityScore + revenueScore) * 50.0 // Scale to 0-100
}

// Placeholder functions for advanced analytics (to be implemented)

func (k Keeper) analyzeUserJourney(activities []types.UserActivity) map[string]interface{} {
	return map[string]interface{}{
		"status": "basic_journey_analysis",
		"total_steps": len(activities),
	}
}

func (k Keeper) analyzeSessions(activities []types.UserActivity) map[string]interface{} {
	return map[string]interface{}{
		"status": "basic_session_analysis",
		"total_activities": len(activities),
	}
}

func (k Keeper) identifyBehaviorPatterns(activities []types.UserActivity) map[string]interface{} {
	return map[string]interface{}{
		"status": "basic_pattern_analysis",
		"pattern_count": 1,
	}
}

func (k Keeper) classifyUserSegment(activities []types.UserActivity) string {
	if len(activities) > 100 {
		return "power_user"
	} else if len(activities) > 20 {
		return "regular_user"
	} else {
		return "casual_user"
	}
}

func (k Keeper) calculateUserRetention(activities []types.UserActivity) map[string]interface{} {
	return map[string]interface{}{
		"status": "basic_retention_analysis",
		"total_activities": len(activities),
	}
}

func (k Keeper) calculateCustomerLifetimeValue(records []types.RevenueRecord) map[string]interface{} {
	return map[string]interface{}{
		"status": "basic_clv_analysis",
		"total_records": len(records),
	}
}

func (k Keeper) analyzeRevenueConcentration(records []types.RevenueRecord) map[string]interface{} {
	return map[string]interface{}{
		"status": "basic_concentration_analysis",
		"total_records": len(records),
	}
}

func (k Keeper) assessFraudRisk(records []types.RevenueRecord) map[string]interface{} {
	return map[string]interface{}{
		"status": "basic_fraud_assessment",
		"risk_level": "low",
	}
}

func (k Keeper) analyzeSeasonalPatterns(records []types.RevenueRecord) map[string]interface{} {
	return map[string]interface{}{
		"status": "basic_seasonal_analysis",
		"total_records": len(records),
	}
}

func (k Keeper) generateRevenueForecast(records []types.RevenueRecord) map[string]interface{} {
	return map[string]interface{}{
		"status": "basic_forecast",
		"next_period_estimate": "tbd",
	}
}
