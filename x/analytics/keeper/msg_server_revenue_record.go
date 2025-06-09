package keeper

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"regexp"
	"strings"
	"time"

	"skillchain/x/analytics/types"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateRevenueRecord(goCtx context.Context, msg *types.MsgCreateRevenueRecord) (*types.MsgCreateRevenueRecordResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Enhanced validation for revenue record
	if err := k.validateRevenueRecordData(msg); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}

	// Generate revenue index if not provided
	revenueIndex := msg.Index
	if revenueIndex == "" {
		revenueIndex = k.generateRevenueIndex(msg.FromAddress, msg.ToAddress, msg.Amount, msg.Timestamp)
	}

	// Check if revenue record already exists
	_, isFound := k.GetRevenueRecord(ctx, revenueIndex)
	if isFound {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "revenue record already exists with this index")
	}

	// Auto-generate timestamp if not provided
	timestamp := msg.Timestamp
	if timestamp == 0 {
		timestamp = uint64(time.Now().Unix())
	}

	// Validate transaction amounts and fees
	if err := k.validateTransactionAmounts(msg.Amount, msg.FeeAmount, msg.PlatformFee); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}

	// Normalize currency code
	currency := k.normalizeCurrency(msg.Currency)

	// Calculate derived financial metrics
	netAmount := k.calculateNetAmount(msg.Amount, msg.FeeAmount, msg.PlatformFee)
	profitMargin := k.calculateProfitMargin(msg.Amount, msg.FeeAmount, msg.PlatformFee)

	// Detect potential fraud patterns
	if k.detectRevenueAnomalies(ctx, msg.FromAddress, msg.ToAddress, msg.Amount, timestamp) {
		// Log suspicious activity but don't block (for now)
		ctx.EventManager().EmitEvent(
			sdk.NewEvent(
				"suspicious_revenue_detected",
				sdk.NewAttribute("from_address", msg.FromAddress),
				sdk.NewAttribute("to_address", msg.ToAddress),
				sdk.NewAttribute("amount", fmt.Sprintf("%d", msg.Amount)),
				sdk.NewAttribute("timestamp", fmt.Sprintf("%d", timestamp)),
			),
		)
	}

	var revenueRecord = types.RevenueRecord{
		Creator:         msg.Creator,
		Index:           revenueIndex,
		TransactionType: k.normalizeTransactionType(msg.TransactionType),
		Amount:          msg.Amount,
		Currency:        currency,
		FromAddress:     msg.FromAddress,
		ToAddress:       msg.ToAddress,
		Timestamp:       timestamp,
		FeeAmount:       msg.FeeAmount,
		ProjectId:       msg.ProjectId,
		PlatformFee:     msg.PlatformFee,
	}

	k.SetRevenueRecord(ctx, revenueRecord)

	// Update financial metrics and aggregations
	k.updateRevenueMetrics(ctx, revenueRecord, netAmount, profitMargin)

	// Update daily/monthly revenue totals
	k.updateRevenueTotals(ctx, msg.Amount, msg.FeeAmount, msg.PlatformFee, currency, timestamp)

	// Emit comprehensive revenue event
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			"revenue_recorded",
			sdk.NewAttribute("creator", msg.Creator),
			sdk.NewAttribute("revenue_index", revenueIndex),
			sdk.NewAttribute("transaction_type", revenueRecord.TransactionType),
			sdk.NewAttribute("amount", fmt.Sprintf("%d", msg.Amount)),
			sdk.NewAttribute("currency", currency),
			sdk.NewAttribute("from_address", msg.FromAddress),
			sdk.NewAttribute("to_address", msg.ToAddress),
			sdk.NewAttribute("fee_amount", fmt.Sprintf("%d", msg.FeeAmount)),
			sdk.NewAttribute("platform_fee", fmt.Sprintf("%d", msg.PlatformFee)),
			sdk.NewAttribute("net_amount", fmt.Sprintf("%d", netAmount)),
			sdk.NewAttribute("profit_margin", fmt.Sprintf("%.2f", profitMargin)),
			sdk.NewAttribute("project_id", msg.ProjectId),
		),
	)

	return &types.MsgCreateRevenueRecordResponse{}, nil
}

func (k msgServer) UpdateRevenueRecord(goCtx context.Context, msg *types.MsgUpdateRevenueRecord) (*types.MsgUpdateRevenueRecordResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the revenue record exists
	valFound, isFound := k.GetRevenueRecord(ctx, msg.Index)
	if !isFound {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, "revenue record not found")
	}

	// Enhanced authorization - only creator or financial admin can update
	if msg.Creator != valFound.Creator && !k.isFinancialAdmin(ctx, msg.Creator) {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "only revenue creator or financial admin can update")
	}

	// Validate update data
	if err := k.validateRevenueRecordUpdateData(msg); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}

	// Preserve immutable fields
	originalTimestamp := valFound.Timestamp
	originalCreator := valFound.Creator

	// Validate updated amounts
	if err := k.validateTransactionAmounts(msg.Amount, msg.FeeAmount, msg.PlatformFee); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}

	// Calculate updated financial metrics
	oldNetAmount := k.calculateNetAmount(valFound.Amount, valFound.FeeAmount, valFound.PlatformFee)
	newNetAmount := k.calculateNetAmount(msg.Amount, msg.FeeAmount, msg.PlatformFee)
	newProfitMargin := k.calculateProfitMargin(msg.Amount, msg.FeeAmount, msg.PlatformFee)

	var revenueRecord = types.RevenueRecord{
		Creator:         originalCreator,                      // Preserve original creator
		Index:           msg.Index,
		TransactionType: k.normalizeTransactionType(msg.TransactionType), // Allow type updates
		Amount:          msg.Amount,                           // Allow amount updates
		Currency:        k.normalizeCurrency(msg.Currency),   // Allow currency updates
		FromAddress:     msg.FromAddress,                      // Allow address updates
		ToAddress:       msg.ToAddress,                        // Allow address updates
		Timestamp:       originalTimestamp,                    // Preserve original timestamp
		FeeAmount:       msg.FeeAmount,                        // Allow fee updates
		ProjectId:       msg.ProjectId,                        // Allow project updates
		PlatformFee:     msg.PlatformFee,                      // Allow platform fee updates
	}

	k.SetRevenueRecord(ctx, revenueRecord)

	// Update financial metrics with the difference
	k.adjustRevenueMetrics(ctx, oldNetAmount, newNetAmount, revenueRecord.Currency)

	// Log financial adjustment
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			"revenue_updated",
			sdk.NewAttribute("updater", msg.Creator),
			sdk.NewAttribute("revenue_index", msg.Index),
			sdk.NewAttribute("old_amount", fmt.Sprintf("%d", valFound.Amount)),
			sdk.NewAttribute("new_amount", fmt.Sprintf("%d", msg.Amount)),
			sdk.NewAttribute("old_net_amount", fmt.Sprintf("%d", oldNetAmount)),
			sdk.NewAttribute("new_net_amount", fmt.Sprintf("%d", newNetAmount)),
			sdk.NewAttribute("new_profit_margin", fmt.Sprintf("%.2f", newProfitMargin)),
		),
	)

	return &types.MsgUpdateRevenueRecordResponse{}, nil
}

func (k msgServer) DeleteRevenueRecord(goCtx context.Context, msg *types.MsgDeleteRevenueRecord) (*types.MsgDeleteRevenueRecordResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the revenue record exists
	valFound, isFound := k.GetRevenueRecord(ctx, msg.Index)
	if !isFound {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, "revenue record not found")
	}

	// Enhanced authorization - only creator or financial admin can delete
	if msg.Creator != valFound.Creator && !k.isFinancialAdmin(ctx, msg.Creator) {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "only revenue creator or financial admin can delete")
	}

	// Calculate financial impact before deletion
	netAmount := k.calculateNetAmount(valFound.Amount, valFound.FeeAmount, valFound.PlatformFee)

	// Archive revenue record before deletion (for audit compliance)
	k.archiveRevenueRecord(ctx, valFound)

	// Reverse financial metrics
	k.reverseRevenueMetrics(ctx, valFound, netAmount)

	k.RemoveRevenueRecord(ctx, msg.Index)

	// Emit revenue deletion event
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			"revenue_deleted",
			sdk.NewAttribute("deleter", msg.Creator),
			sdk.NewAttribute("revenue_index", msg.Index),
			sdk.NewAttribute("amount", fmt.Sprintf("%d", valFound.Amount)),
			sdk.NewAttribute("currency", valFound.Currency),
			sdk.NewAttribute("net_amount", fmt.Sprintf("%d", netAmount)),
			sdk.NewAttribute("archived", "true"),
		),
	)

	return &types.MsgDeleteRevenueRecordResponse{}, nil
}

// validateRevenueRecordData performs comprehensive validation
func (k msgServer) validateRevenueRecordData(msg *types.MsgCreateRevenueRecord) error {
	if msg.TransactionType == "" {
		return fmt.Errorf("transaction type cannot be empty")
	}
	if !k.isValidTransactionType(msg.TransactionType) {
		return fmt.Errorf("invalid transaction type: %s", msg.TransactionType)
	}
	if msg.Amount == 0 {
		return fmt.Errorf("amount must be greater than zero")
	}
	if msg.Amount > 1000000000000 { // 1 trillion limit
		return fmt.Errorf("amount exceeds maximum limit")
	}
	if msg.Currency == "" {
		return fmt.Errorf("currency cannot be empty")
	}
	if !k.isValidCurrency(msg.Currency) {
		return fmt.Errorf("invalid currency code: %s", msg.Currency)
	}
	if msg.FromAddress == "" {
		return fmt.Errorf("from address cannot be empty")
	}
	if msg.ToAddress == "" {
		return fmt.Errorf("to address cannot be empty")
	}
	if msg.FromAddress == msg.ToAddress {
		return fmt.Errorf("from and to addresses cannot be the same")
	}
	if msg.FeeAmount > msg.Amount {
		return fmt.Errorf("fee amount cannot exceed transaction amount")
	}
	if msg.PlatformFee > msg.Amount {
		return fmt.Errorf("platform fee cannot exceed transaction amount")
	}
	return nil
}

// validateRevenueRecordUpdateData validates update-specific data
func (k msgServer) validateRevenueRecordUpdateData(msg *types.MsgUpdateRevenueRecord) error {
	if msg.TransactionType == "" {
		return fmt.Errorf("transaction type cannot be empty")
	}
	if !k.isValidTransactionType(msg.TransactionType) {
		return fmt.Errorf("invalid transaction type: %s", msg.TransactionType)
	}
	if msg.Amount == 0 {
		return fmt.Errorf("amount must be greater than zero")
	}
	if msg.Amount > 1000000000000 { // 1 trillion limit
		return fmt.Errorf("amount exceeds maximum limit")
	}
	if msg.Currency == "" {
		return fmt.Errorf("currency cannot be empty")
	}
	if !k.isValidCurrency(msg.Currency) {
		return fmt.Errorf("invalid currency code: %s", msg.Currency)
	}
	if msg.FromAddress == "" {
		return fmt.Errorf("from address cannot be empty")
	}
	if msg.ToAddress == "" {
		return fmt.Errorf("to address cannot be empty")
	}
	if msg.FromAddress == msg.ToAddress {
		return fmt.Errorf("from and to addresses cannot be the same")
	}
	if msg.FeeAmount > msg.Amount {
		return fmt.Errorf("fee amount cannot exceed transaction amount")
	}
	if msg.PlatformFee > msg.Amount {
		return fmt.Errorf("platform fee cannot exceed transaction amount")
	}
	return nil
}

// generateRevenueIndex creates a unique revenue identifier
func (k msgServer) generateRevenueIndex(fromAddress, toAddress string, amount, timestamp uint64) string {
	if timestamp == 0 {
		timestamp = uint64(time.Now().Unix())
	}
	data := fmt.Sprintf("%s:%s:%d:%d", fromAddress, toAddress, amount, timestamp)
	hash := sha256.Sum256([]byte(data))
	return hex.EncodeToString(hash[:])[:16]
}

// isValidTransactionType validates transaction types
func (k msgServer) isValidTransactionType(transactionType string) bool {
	validTypes := []string{
		"sale", "purchase", "transfer", "payment", "refund", "withdrawal",
		"deposit", "fee", "commission", "subscription", "donation", "reward",
		"penalty", "tax", "interest", "dividend", "loan", "repayment",
		"marketplace_sale", "skill_purchase", "file_download", "premium_upgrade",
	}
	
	for _, validType := range validTypes {
		if strings.EqualFold(transactionType, validType) {
			return true
		}
	}
	return false
}

// isValidCurrency validates currency codes
func (k msgServer) isValidCurrency(currency string) bool {
	// Common cryptocurrency and fiat currency codes
	validCurrencies := []string{
		"USD", "EUR", "GBP", "JPY", "CNY", "KRW", "BTC", "ETH", "USDT", "USDC",
		"BNB", "ADA", "SOL", "DOT", "AVAX", "MATIC", "ATOM", "SKILL", "CHAIN",
	}
	
	upperCurrency := strings.ToUpper(currency)
	for _, validCurrency := range validCurrencies {
		if upperCurrency == validCurrency {
			return true
		}
	}
	
	// Also accept pattern for token contracts (40 char hex)
	tokenPattern := regexp.MustCompile(`^0x[a-fA-F0-9]{40}$`)
	return tokenPattern.MatchString(currency)
}

// normalizeCurrency converts currency to uppercase
func (k msgServer) normalizeCurrency(currency string) string {
	return strings.ToUpper(strings.TrimSpace(currency))
}

// normalizeTransactionType converts transaction type to lowercase
func (k msgServer) normalizeTransactionType(transactionType string) string {
	return strings.ToLower(strings.TrimSpace(transactionType))
}

// validateTransactionAmounts validates financial amounts
func (k msgServer) validateTransactionAmounts(amount, feeAmount, platformFee uint64) error {
	if amount == 0 {
		return fmt.Errorf("transaction amount must be greater than zero")
	}
	if feeAmount+platformFee > amount {
		return fmt.Errorf("total fees cannot exceed transaction amount")
	}
	return nil
}

// calculateNetAmount calculates net amount after fees
func (k msgServer) calculateNetAmount(amount, feeAmount, platformFee uint64) uint64 {
	if amount <= feeAmount+platformFee {
		return 0
	}
	return amount - feeAmount - platformFee
}

// calculateProfitMargin calculates profit margin percentage
func (k msgServer) calculateProfitMargin(amount, feeAmount, platformFee uint64) float64 {
	if amount == 0 {
		return 0.0
	}
	profit := float64(feeAmount + platformFee)
	return (profit / float64(amount)) * 100.0
}

// detectRevenueAnomalies detects suspicious financial patterns
func (k msgServer) detectRevenueAnomalies(ctx sdk.Context, fromAddress, toAddress string, amount, timestamp uint64) bool {
	// Check for suspicious patterns
	
	// 1. Check for unusually large amounts from new addresses
	if amount > 1000000000 { // > 1 billion units
		recentActivity := k.getRecentRevenueActivity(ctx, fromAddress, timestamp-86400) // Last 24 hours
		if len(recentActivity) == 0 {
			return true // Large amount from new address
		}
	}
	
	// 2. Check for rapid consecutive transactions
	recentTransactions := k.getRecentTransactionsBetween(ctx, fromAddress, toAddress, timestamp-300) // Last 5 minutes
	if len(recentTransactions) > 10 {
		return true // Too many transactions in short time
	}
	
	// 3. Check for round number amounts (potential money laundering)
	if amount%1000000 == 0 && amount >= 10000000 { // Round millions
		return true
	}
	
	return false
}

// updateRevenueMetrics updates financial metrics and KPIs
func (k msgServer) updateRevenueMetrics(ctx sdk.Context, record types.RevenueRecord, netAmount uint64, profitMargin float64) {
	// Emit metrics update events
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			"revenue_metrics_update",
			sdk.NewAttribute("transaction_type", record.TransactionType),
			sdk.NewAttribute("currency", record.Currency),
			sdk.NewAttribute("gross_amount", fmt.Sprintf("%d", record.Amount)),
			sdk.NewAttribute("net_amount", fmt.Sprintf("%d", netAmount)),
			sdk.NewAttribute("profit_margin", fmt.Sprintf("%.2f", profitMargin)),
			sdk.NewAttribute("project_id", record.ProjectId),
		),
	)
}

// updateRevenueTotals updates daily/monthly aggregations
func (k msgServer) updateRevenueTotals(ctx sdk.Context, amount, feeAmount, platformFee uint64, currency string, timestamp uint64) {
	// Calculate daily and monthly keys
	t := time.Unix(int64(timestamp), 0).UTC()
	dailyKey := t.Format("2006-01-02")
	monthlyKey := t.Format("2006-01")
	
	// Emit aggregation events
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			"revenue_totals_update",
			sdk.NewAttribute("daily_key", dailyKey),
			sdk.NewAttribute("monthly_key", monthlyKey),
			sdk.NewAttribute("currency", currency),
			sdk.NewAttribute("amount", fmt.Sprintf("%d", amount)),
			sdk.NewAttribute("fee_amount", fmt.Sprintf("%d", feeAmount)),
			sdk.NewAttribute("platform_fee", fmt.Sprintf("%d", platformFee)),
		),
	)
}

// adjustRevenueMetrics adjusts metrics on updates
func (k msgServer) adjustRevenueMetrics(ctx sdk.Context, oldNetAmount, newNetAmount uint64, currency string) {
	adjustment := int64(newNetAmount) - int64(oldNetAmount)
	
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			"revenue_metrics_adjustment",
			sdk.NewAttribute("currency", currency),
			sdk.NewAttribute("adjustment", fmt.Sprintf("%d", adjustment)),
			sdk.NewAttribute("old_net_amount", fmt.Sprintf("%d", oldNetAmount)),
			sdk.NewAttribute("new_net_amount", fmt.Sprintf("%d", newNetAmount)),
		),
	)
}

// reverseRevenueMetrics reverses metrics on deletion
func (k msgServer) reverseRevenueMetrics(ctx sdk.Context, record types.RevenueRecord, netAmount uint64) {
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			"revenue_metrics_reversal",
			sdk.NewAttribute("currency", record.Currency),
			sdk.NewAttribute("reversed_amount", fmt.Sprintf("%d", record.Amount)),
			sdk.NewAttribute("reversed_net_amount", fmt.Sprintf("%d", netAmount)),
			sdk.NewAttribute("transaction_type", record.TransactionType),
		),
	)
}

// getRecentRevenueActivity gets recent activity for an address
func (k msgServer) getRecentRevenueActivity(ctx sdk.Context, address string, sinceTimestamp uint64) []types.RevenueRecord {
	var recentActivity []types.RevenueRecord
	allRecords := k.GetAllRevenueRecord(ctx)
	
	for _, record := range allRecords {
		if (record.FromAddress == address || record.ToAddress == address) && record.Timestamp >= sinceTimestamp {
			recentActivity = append(recentActivity, record)
		}
	}
	
	return recentActivity
}

// getRecentTransactionsBetween gets recent transactions between addresses
func (k msgServer) getRecentTransactionsBetween(ctx sdk.Context, fromAddress, toAddress string, sinceTimestamp uint64) []types.RevenueRecord {
	var recentTransactions []types.RevenueRecord
	allRecords := k.GetAllRevenueRecord(ctx)
	
	for _, record := range allRecords {
		if record.FromAddress == fromAddress && record.ToAddress == toAddress && record.Timestamp >= sinceTimestamp {
			recentTransactions = append(recentTransactions, record)
		}
	}
	
	return recentTransactions
}

// isFinancialAdmin checks if user has financial admin privileges
func (k msgServer) isFinancialAdmin(ctx sdk.Context, userAddress string) bool {
	// In production, this would check against financial admin list
	return false
}

// archiveRevenueRecord archives record for audit compliance
func (k msgServer) archiveRevenueRecord(ctx sdk.Context, record types.RevenueRecord) {
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			"revenue_record_archived",
			sdk.NewAttribute("revenue_index", record.Index),
			sdk.NewAttribute("transaction_type", record.TransactionType),
			sdk.NewAttribute("amount", fmt.Sprintf("%d", record.Amount)),
			sdk.NewAttribute("currency", record.Currency),
			sdk.NewAttribute("archived_at", fmt.Sprintf("%d", time.Now().Unix())),
		),
	)
}
