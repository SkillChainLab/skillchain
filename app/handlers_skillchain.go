package app

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os/exec"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/types"
	"github.com/gorilla/mux"
	
	skillchaintypes "skillchain/x/skillchain/types"
)

// Query Handlers

func (app *App) handleGetParams(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		queryClient := skillchaintypes.NewQueryClient(clientCtx)
		req := &skillchaintypes.QueryParamsRequest{}

		res, err := queryClient.Params(r.Context(), req)
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to query params: %v", err), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"params": res.Params,
		})
	}
}

func (app *App) handleGetTokenInfo(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		queryClient := skillchaintypes.NewQueryClient(clientCtx)
		req := &skillchaintypes.QueryTokenInfoRequest{}

		res, err := queryClient.TokenInfo(r.Context(), req)
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to query token info: %v", err), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"name": res.Name,
			"symbol": res.Symbol,
			"decimals": res.Decimals,
			"description": res.Description,
			"total_supply": res.TotalSupply,
			"circulating_supply": res.CirculatingSupply,
			"burned_amount": res.BurnedAmount,
			"max_supply": res.MaxSupply,
			"burn_enabled": res.BurnEnabled,
			"chain_description": res.ChainDescription,
			"website_url": res.WebsiteUrl,
		})
	}
}

func (app *App) handleGetVUSDTreasury(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		queryClient := skillchaintypes.NewQueryClient(clientCtx)
		req := &skillchaintypes.QueryVUSDTreasuryRequest{}

		res, err := queryClient.VUSDTreasury(r.Context(), req)
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to query VUSD treasury: %v", err), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"skill_balance": res.SkillBalance,
			"vusd_supply": res.VusdSupply,
			"exchange_rate": res.ExchangeRate,
		})
	}
}

func (app *App) handleGetUserVUSDPosition(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		address := vars["address"]

		if address == "" {
			http.Error(w, "Address is required", http.StatusBadRequest)
			return
		}

		queryClient := skillchaintypes.NewQueryClient(clientCtx)
		req := &skillchaintypes.QueryUserVUSDPositionRequest{
			Address: address,
		}

		res, err := queryClient.UserVUSDPosition(r.Context(), req)
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to query user VUSD position: %v", err), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"position": res.Position,
			"address": address,
			"vusd_balance": res.VusdBalance,
			"skill_collateral": res.SkillCollateral,
			"health_factor": res.HealthFactor,
			"exists": res.Exists,
		})
	}
}

// Transaction Preparation Handlers

func (app *App) handleConvertSKILLToVUSD(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			Creator string `json:"creator"`
			Amount  string `json:"amount"`
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

		// Parse amount
		amount, err := types.ParseCoinNormalized(req.Amount)
		if err != nil {
			http.Error(w, "Invalid amount format", http.StatusBadRequest)
			return
		}

		// Execute actual conversion transaction using CLI
		cmd := exec.Command("skillchaind", "tx", "skillchain", "convert-skill-to-vusd", 
			"--amount", amount.String(), 
			"--from", "alice", 
			"--chain-id", "skillchain", 
			"--yes", 
			"--output", "json")
		
		output, err := cmd.Output()
		if err != nil {
			http.Error(w, fmt.Sprintf("Transaction failed: %v", err), http.StatusInternalServerError)
			return
		}

		// Parse the transaction output
		var txResult map[string]interface{}
		if err := json.Unmarshal(output, &txResult); err != nil {
			http.Error(w, "Failed to parse transaction result", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"success": true,
			"message": "SKILL to VUSD conversion executed successfully",
			"transaction": txResult,
			"conversion_result": map[string]interface{}{
				"original_amount": amount.String(),
				"vusd_received": fmt.Sprintf("%d", amount.Amount.Int64()/2) + "uvusd", // 1:0.5 ratio
				"exchange_rate": "0.5",
			},
		})
	}
}

func (app *App) handleConvertVUSDToSKILL(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			Creator string `json:"creator"`
			Amount  string `json:"amount"`
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

		// Parse amount
		amount, err := types.ParseCoinNormalized(req.Amount)
		if err != nil {
			http.Error(w, "Invalid amount format", http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "VUSD to SKILL conversion transaction prepared",
			"tx_info": map[string]interface{}{
				"creator": req.Creator,
				"amount": amount.String(),
				"estimated_gas": "180000",
				"note": "Use skillchaind tx skillchain convert-vusd-to-skill to execute",
			},
		})
	}
}

func (app *App) handleBurnTokens(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			Creator string `json:"creator"`
			Amount  string `json:"amount"`
			Denom   string `json:"denom"`
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

		// Parse amount
		amount, err := types.ParseCoinNormalized(req.Amount)
		if err != nil {
			http.Error(w, "Invalid amount format", http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "Token burn transaction prepared",
			"tx_info": map[string]interface{}{
				"creator": req.Creator,
				"amount": amount.String(),
				"denom": req.Denom,
				"estimated_gas": "150000",
				"note": "Use skillchaind tx skillchain burn-tokens to execute",
			},
		})
	}
}

// Bank Module Handlers for Wallet Integration

func (app *App) handleGetBalances(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		address := vars["address"]

		// Validate address
		if _, err := types.AccAddressFromBech32(address); err != nil {
			http.Error(w, "Invalid address format", http.StatusBadRequest)
			return
		}

		// Forward to standard Cosmos SDK bank module
		resp, err := http.Get(fmt.Sprintf("http://localhost:1317/cosmos/bank/v1beta1/balances/%s", address))
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to query balances: %v", err), http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()

		// Copy response headers
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "*")

		// Copy status code
		w.WriteHeader(resp.StatusCode)

		// Copy response body
		_, err = w.Write(func() []byte {
			body, _ := io.ReadAll(resp.Body)
			return body
		}())
		if err != nil {
			http.Error(w, "Failed to copy response", http.StatusInternalServerError)
		}
	}
}

func (app *App) handleGetSupply(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Forward to standard Cosmos SDK bank module
		resp, err := http.Get("http://localhost:1317/cosmos/bank/v1beta1/supply")
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to query supply: %v", err), http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()

		// Copy response headers
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")

		// Copy status code
		w.WriteHeader(resp.StatusCode)

		// Copy response body
		_, err = w.Write(func() []byte {
			body, _ := io.ReadAll(resp.Body)
			return body
		}())
		if err != nil {
			http.Error(w, "Failed to copy response", http.StatusInternalServerError)
		}
	}
}

func (app *App) handleGetSupplyByDenom(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		denom := vars["denom"]

		// Forward to standard Cosmos SDK bank module
		resp, err := http.Get(fmt.Sprintf("http://localhost:1317/cosmos/bank/v1beta1/supply/by_denom?denom=%s", denom))
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to query supply: %v", err), http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()

		// Copy response headers
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")

		// Copy status code
		w.WriteHeader(resp.StatusCode)

		// Copy response body
		_, err = w.Write(func() []byte {
			body, _ := io.ReadAll(resp.Body)
			return body
		}())
		if err != nil {
			http.Error(w, "Failed to copy response", http.StatusInternalServerError)
		}
	}
}