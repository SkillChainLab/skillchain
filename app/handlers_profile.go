package app

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/gorilla/mux"
	
	profiletypes "skillchain/x/profile/types"
)

// Profile Handlers

func (app *App) handleCreateProfile(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			Creator     string `json:"creator"`
			Name        string `json:"name"`
			Description string `json:"description"`
			Avatar      string `json:"avatar,omitempty"`
			Website     string `json:"website,omitempty"`
			Location    string `json:"location,omitempty"`
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

		// Set default avatar if not provided
		if req.Avatar == "" {
			req.Avatar = "https://ipfs.io/ipfs/QmZSpwmV3dfwVDVcaJmdga3VVW15SEgXQDY1wiEj8gzpqc"
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "Profile creation transaction prepared",
			"tx_info": map[string]interface{}{
				"creator": req.Creator,
				"name": req.Name,
				"description": req.Description,
				"avatar": req.Avatar,
				"website": req.Website,
				"location": req.Location,
				"estimated_gas": "150000",
				"note": "Use skillchaind tx profile create-user-profile to execute",
			},
		})
	}
}

func (app *App) handleListProfiles(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		queryClient := profiletypes.NewQueryClient(clientCtx)
		
		// Parse pagination parameters
		pageKey := r.URL.Query().Get("page_key")
		limit := r.URL.Query().Get("limit")
		if limit == "" {
			limit = "50"
		}

		req := &profiletypes.QueryAllUserProfileRequest{
			Pagination: &query.PageRequest{
				Key:   []byte(pageKey),
				Limit: 50, // Default limit
			},
		}

		res, err := queryClient.UserProfileAll(r.Context(), req)
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to query profiles: %v", err), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"profiles": res.UserProfile,
			"pagination": res.Pagination,
			"total_count": len(res.UserProfile),
		})
	}
}

func (app *App) handleGetProfile(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		identifier := vars["address"] // This could be profile index or owner address
		
		fmt.Printf("DEBUG: Looking for profile with identifier: %s\n", identifier)
		
		queryClient := profiletypes.NewQueryClient(clientCtx)
		
		// First try to get by index
		req := &profiletypes.QueryGetUserProfileRequest{
			Index: identifier,
		}
		
		fmt.Printf("DEBUG: Trying to find by index: %s\n", identifier)
		res, err := queryClient.UserProfile(r.Context(), req)
		if err != nil {
			fmt.Printf("DEBUG: Index lookup failed: %v\n", err)
			// If not found by index, try to search by owner address
			// Get all profiles and search for matching owner
			allReq := &profiletypes.QueryAllUserProfileRequest{
				Pagination: &query.PageRequest{
					Limit: 1000, // Get all profiles to search
				},
			}
			
			fmt.Printf("DEBUG: Fetching all profiles to search by owner\n")
			allRes, allErr := queryClient.UserProfileAll(r.Context(), allReq)
			if allErr != nil {
				fmt.Printf("DEBUG: Failed to get all profiles: %v\n", allErr)
				http.Error(w, fmt.Sprintf("Failed to query profiles: %v", allErr), http.StatusInternalServerError)
				return
			}
			
			fmt.Printf("DEBUG: Found %d profiles to search through\n", len(allRes.UserProfile))
			
			// Search for profile with matching owner
			var foundProfile *profiletypes.UserProfile
			for i, profile := range allRes.UserProfile {
				fmt.Printf("DEBUG: Profile %d - Owner: %s, Creator: %s, Index: %s\n", i, profile.Owner, profile.Creator, profile.Index)
				if profile.Owner == identifier || profile.Creator == identifier {
					fmt.Printf("DEBUG: Found matching profile by owner/creator!\n")
					foundProfile = &profile
					break
				}
			}
			
			if foundProfile == nil {
				fmt.Printf("DEBUG: No profile found for address: %s\n", identifier)
				http.Error(w, fmt.Sprintf("Profile not found for address: %s", identifier), http.StatusNotFound)
				return
			}
			
			fmt.Printf("DEBUG: Returning profile found by owner address\n")
			w.Header().Set("Content-Type", "application/json")
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
			json.NewEncoder(w).Encode(map[string]interface{}{
				"profile": foundProfile,
				"identifier": identifier,
				"found_by": "owner_address",
			})
			return
		}

		fmt.Printf("DEBUG: Found profile by index\n")
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"profile": res.UserProfile,
			"identifier": identifier,
			"found_by": "index",
		})
	}
}

func (app *App) handleUpdateProfile(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		address := vars["address"]

		var req struct {
			Name        string `json:"name,omitempty"`
			Description string `json:"description,omitempty"`
			Avatar      string `json:"avatar,omitempty"`
			Website     string `json:"website,omitempty"`
			Location    string `json:"location,omitempty"`
		}

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		// Validate address
		if _, err := types.AccAddressFromBech32(address); err != nil {
			http.Error(w, "Invalid address format", http.StatusBadRequest)
			return
		}

		// If avatar is being set to empty, provide default
		if req.Avatar == "" {
			req.Avatar = "https://ipfs.io/ipfs/QmZSpwmV3dfwVDVcaJmdga3VVW15SEgXQDY1wiEj8gzpqc"
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "Profile update transaction prepared",
			"tx_info": map[string]interface{}{
				"creator": address,
				"updates": req,
				"estimated_gas": "120000",
				"note": "Use skillchaind tx profile update-user-profile to execute",
			},
		})
	}
}

// Skills Handlers

func (app *App) handleAddSkill(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		address := vars["address"]

		var req struct {
			SkillName           string `json:"skill_name"`
			ProficiencyLevel    string `json:"proficiency_level"`
			YearsOfExperience   int32  `json:"years_of_experience"`
			IsVerified          bool   `json:"is_verified"`
			VerifiedBy          string `json:"verified_by,omitempty"`
			VerificationDetails string `json:"verification_details,omitempty"`
		}

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		// Validate address
		if _, err := types.AccAddressFromBech32(address); err != nil {
			http.Error(w, "Invalid address format", http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "Add skill transaction prepared",
			"tx_info": map[string]interface{}{
				"creator": address,
				"skill": req,
				"estimated_gas": "100000",
				"note": "Use skillchaind tx profile create-user-skill to execute",
			},
		})
	}
}

func (app *App) handleListUserSkills(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		address := vars["address"]

		queryClient := profiletypes.NewQueryClient(clientCtx)
		req := &profiletypes.QueryAllUserSkillRequest{
			Pagination: &query.PageRequest{
				Limit: 100,
			},
		}

		res, err := queryClient.UserSkillAll(r.Context(), req)
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to query skills: %v", err), http.StatusInternalServerError)
			return
		}

		// Filter skills by user address (use Owner instead of Creator)
		var userSkills []interface{}
		for _, skill := range res.UserSkill {
			if skill.Owner == address {
				userSkills = append(userSkills, skill)
			}
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"skills": userSkills,
			"address": address,
			"count": len(userSkills),
		})
	}
}

func (app *App) handleGetSkill(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		address := vars["address"]
		skillID := vars["skillId"]

		queryClient := profiletypes.NewQueryClient(clientCtx)
		req := &profiletypes.QueryGetUserSkillRequest{
			Index: skillID,
		}

		res, err := queryClient.UserSkill(r.Context(), req)
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to query skill: %v", err), http.StatusInternalServerError)
			return
		}

		// Verify skill belongs to user (use Owner instead of Creator)
		if res.UserSkill.Owner != address {
			http.Error(w, "Skill does not belong to the specified user", http.StatusForbidden)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"skill": res.UserSkill,
			"address": address,
			"skill_id": skillID,
		})
	}
}

func (app *App) handleEndorseSkill(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		address := vars["address"]
		skillID := vars["skillId"]

		var req struct {
			EndorserAddress string `json:"endorser_address"`
			EndorsementType string `json:"endorsement_type"`
			StakeAmount     string `json:"stake_amount"`
			Comment         string `json:"comment,omitempty"`
		}

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		// Validate endorser address
		if _, err := types.AccAddressFromBech32(req.EndorserAddress); err != nil {
			http.Error(w, "Invalid endorser address format", http.StatusBadRequest)
			return
		}

		// Parse stake amount
		stakeAmount, err := types.ParseCoinNormalized(req.StakeAmount)
		if err != nil {
			http.Error(w, "Invalid stake amount format", http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "Skill endorsement transaction prepared",
			"tx_info": map[string]interface{}{
				"skill_owner": address,
				"skill_id": skillID,
				"endorser": req.EndorserAddress,
				"endorsement_type": req.EndorsementType,
				"stake_amount": stakeAmount.String(),
				"comment": req.Comment,
				"estimated_gas": "180000",
				"note": "Use skillchaind tx profile endorse-skill to execute",
			},
		})
	}
}

// Endorsements Handlers

func (app *App) handleListEndorsements(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		queryClient := profiletypes.NewQueryClient(clientCtx)
		
		req := &profiletypes.QueryAllSkillEndorsementRequest{
			Pagination: &query.PageRequest{
				Limit: 100,
			},
		}

		res, err := queryClient.SkillEndorsementAll(r.Context(), req)
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to query endorsements: %v", err), http.StatusInternalServerError)
			return
		}

		// Filter by query parameters if provided
		skillID := r.URL.Query().Get("skill_id")
		endorserAddr := r.URL.Query().Get("endorser")
		
		var filteredEndorsements []interface{}
		for _, endorsement := range res.SkillEndorsement {
			include := true
			
			if skillID != "" && endorsement.SkillName != skillID {
				include = false
			}
			if endorserAddr != "" && endorsement.Endorser != endorserAddr {
				include = false
			}
			
			if include {
				filteredEndorsements = append(filteredEndorsements, endorsement)
			}
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"endorsements": filteredEndorsements,
			"filters": map[string]string{
				"skill_id": skillID,
				"endorser": endorserAddr,
			},
			"count": len(filteredEndorsements),
		})
	}
}

func (app *App) handleGetEndorsement(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		endorsementID := vars["endorsementId"]

		queryClient := profiletypes.NewQueryClient(clientCtx)
		req := &profiletypes.QueryGetSkillEndorsementRequest{
			Index: endorsementID,
		}

		res, err := queryClient.SkillEndorsement(r.Context(), req)
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to query endorsement: %v", err), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"endorsement": res.SkillEndorsement,
			"endorsement_id": endorsementID,
		})
	}
}

// handleSearchProfiles handles advanced profile search
func (app *App) handleSearchProfiles(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Parse query parameters
		name := r.URL.Query().Get("name")
		location := r.URL.Query().Get("location") 
		skill := r.URL.Query().Get("skill")
		bio := r.URL.Query().Get("bio")
		github := r.URL.Query().Get("github")
		searchType := r.URL.Query().Get("type") // "name", "location", "skill", "bio", "github", "advanced"
		
		fmt.Printf("DEBUG: Search request - name: %s, location: %s, skill: %s, bio: %s, github: %s, type: %s\n", 
			name, location, skill, bio, github, searchType)

		// Get all profiles for client-side search (since we don't have gRPC methods for search yet)
		queryClient := profiletypes.NewQueryClient(clientCtx)
		allReq := &profiletypes.QueryAllUserProfileRequest{
			Pagination: &query.PageRequest{
				Limit: 1000, // Get all profiles
			},
		}

		allRes, err := queryClient.UserProfileAll(r.Context(), allReq)
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to query profiles: %v", err), http.StatusInternalServerError)
			return
		}

		var matchingProfiles []profiletypes.UserProfile

		// Perform search based on type
		switch searchType {
		case "name":
			if name != "" {
				matchingProfiles = searchProfilesByName(allRes.UserProfile, name)
			}
		case "location":
			if location != "" {
				matchingProfiles = searchProfilesByLocation(allRes.UserProfile, location)
			}
		case "skill":
			if skill != "" {
				// For skill search, we need to get user skills too
				matchingProfiles = searchProfilesBySkill(allRes.UserProfile, skill, clientCtx, r)
			}
		case "bio":
			if bio != "" {
				matchingProfiles = searchProfilesByBio(allRes.UserProfile, bio)
			}
		case "github":
			if github != "" {
				matchingProfiles = searchProfilesByGithub(allRes.UserProfile, github)
			}
		case "advanced":
			matchingProfiles = searchProfilesAdvanced(allRes.UserProfile, name, location, skill, bio, clientCtx, r)
		default:
			// Default: search all fields if any search term is provided
			if name != "" || location != "" || skill != "" || bio != "" || github != "" {
				matchingProfiles = searchProfilesAdvanced(allRes.UserProfile, name, location, skill, bio, clientCtx, r)
			} else {
				matchingProfiles = allRes.UserProfile // Return all if no search criteria
			}
		}

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"profiles": matchingProfiles,
			"total_count": len(matchingProfiles),
			"search_criteria": map[string]string{
				"name": name,
				"location": location,
				"skill": skill,
				"bio": bio,
				"github": github,
				"type": searchType,
			},
		})
	}
}

// Helper functions for profile search
func searchProfilesByName(profiles []profiletypes.UserProfile, searchTerm string) []profiletypes.UserProfile {
	var matches []profiletypes.UserProfile
	searchLower := strings.ToLower(searchTerm)
	
	for _, profile := range profiles {
		if strings.Contains(strings.ToLower(profile.DisplayName), searchLower) {
			matches = append(matches, profile)
		}
	}
	return matches
}

func searchProfilesByLocation(profiles []profiletypes.UserProfile, location string) []profiletypes.UserProfile {
	var matches []profiletypes.UserProfile
	locationLower := strings.ToLower(location)
	
	for _, profile := range profiles {
		if strings.Contains(strings.ToLower(profile.Location), locationLower) {
			matches = append(matches, profile)
		}
	}
	return matches
}

func searchProfilesByBio(profiles []profiletypes.UserProfile, searchTerm string) []profiletypes.UserProfile {
	var matches []profiletypes.UserProfile
	searchLower := strings.ToLower(searchTerm)
	
	for _, profile := range profiles {
		if strings.Contains(strings.ToLower(profile.Bio), searchLower) {
			matches = append(matches, profile)
		}
	}
	return matches
}

func searchProfilesByGithub(profiles []profiletypes.UserProfile, github string) []profiletypes.UserProfile {
	var matches []profiletypes.UserProfile
	githubLower := strings.ToLower(github)
	
	for _, profile := range profiles {
		if strings.Contains(strings.ToLower(profile.Github), githubLower) {
			matches = append(matches, profile)
		}
	}
	return matches
}

func searchProfilesBySkill(profiles []profiletypes.UserProfile, skillName string, clientCtx client.Context, r *http.Request) []profiletypes.UserProfile {
	// Get all user skills
	queryClient := profiletypes.NewQueryClient(clientCtx)
	skillsReq := &profiletypes.QueryAllUserSkillRequest{
		Pagination: &query.PageRequest{
			Limit: 1000,
		},
	}
	
	skillsRes, err := queryClient.UserSkillAll(r.Context(), skillsReq)
	if err != nil {
		return []profiletypes.UserProfile{} // Return empty if can't get skills
	}
	
	// Find users with matching skills
	userAddresses := make(map[string]bool)
	skillLower := strings.ToLower(skillName)
	
	for _, skill := range skillsRes.UserSkill {
		if strings.Contains(strings.ToLower(skill.SkillName), skillLower) {
			userAddresses[skill.Owner] = true
		}
	}
	
	// Filter profiles by users who have the skill
	var matches []profiletypes.UserProfile
	for _, profile := range profiles {
		if userAddresses[profile.Owner] {
			matches = append(matches, profile)
		}
	}
	
	return matches
}

func searchProfilesAdvanced(profiles []profiletypes.UserProfile, name, location, skill, bio string, clientCtx client.Context, r *http.Request) []profiletypes.UserProfile {
	var matches []profiletypes.UserProfile
	
	// Get user skills if skill search is needed
	var userSkillMap map[string]bool
	if skill != "" {
		queryClient := profiletypes.NewQueryClient(clientCtx)
		skillsReq := &profiletypes.QueryAllUserSkillRequest{
			Pagination: &query.PageRequest{
				Limit: 1000,
			},
		}
		
		skillsRes, err := queryClient.UserSkillAll(r.Context(), skillsReq)
		if err == nil {
			userSkillMap = make(map[string]bool)
			skillLower := strings.ToLower(skill)
			
			for _, userSkill := range skillsRes.UserSkill {
				if strings.Contains(strings.ToLower(userSkill.SkillName), skillLower) {
					userSkillMap[userSkill.Owner] = true
				}
			}
		}
	}
	
	for _, profile := range profiles {
		match := true
		
		// Check name match
		if name != "" {
			if !strings.Contains(strings.ToLower(profile.DisplayName), strings.ToLower(name)) {
				match = false
			}
		}
		
		// Check location match
		if location != "" && match {
			if !strings.Contains(strings.ToLower(profile.Location), strings.ToLower(location)) {
				match = false
			}
		}
		
		// Check bio match
		if bio != "" && match {
			if !strings.Contains(strings.ToLower(profile.Bio), strings.ToLower(bio)) {
				match = false
			}
		}
		
		// Check skill match
		if skill != "" && match && userSkillMap != nil {
			if !userSkillMap[profile.Owner] {
				match = false
			}
		}
		
		if match {
			matches = append(matches, profile)
		}
	}
	
	return matches
} 