package app

import (
	"encoding/json"
	"fmt"
	"net/http"

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

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "SKILL to VUSD conversion transaction prepared",
			"tx_info": map[string]interface{}{
				"creator": req.Creator,
				"amount": amount.String(),
				"estimated_gas": "180000",
				"note": "Use skillchaind tx skillchain convert-skill-to-vusd to execute",
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

func (app *App) RegisterSkillchainHandlers(router *mux.Router, clientCtx client.Context) {
	// Query endpoints
	router.HandleFunc("/api/skillchain/params", app.handleGetParams(clientCtx)).Methods("GET")
	router.HandleFunc("/api/skillchain/token/info", app.handleGetTokenInfo(clientCtx)).Methods("GET")
	router.HandleFunc("/api/skillchain/vusd/treasury", app.handleGetVUSDTreasury(clientCtx)).Methods("GET")
	router.HandleFunc("/api/skillchain/vusd/position/{address}", app.handleGetUserVUSDPosition(clientCtx)).Methods("GET")
	
	// Transaction preparation endpoints
	router.HandleFunc("/api/skillchain/convert/skill-to-vusd", app.handleConvertSKILLToVUSD(clientCtx)).Methods("POST")
	router.HandleFunc("/api/skillchain/convert/vusd-to-skill", app.handleConvertVUSDToSKILL(clientCtx)).Methods("POST")
	router.HandleFunc("/api/skillchain/burn", app.handleBurnTokens(clientCtx)).Methods("POST")
}