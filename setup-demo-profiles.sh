#!/bin/bash

# SkillChain Demo Profiles Setup Script
# This script creates 10 demo profiles with their skills on the blockchain

echo "🚀 Starting SkillChain Demo Profiles Setup..."
echo "⚠️  WARNING: This uses DEMO wallets - DO NOT use in production!"
echo ""

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Chain ID
CHAIN_ID="skillchain"

# Demo wallet mnemonics and addresses
declare -A WALLETS=(
    ["alice"]="abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon about"
    ["bob"]="abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon" 
    ["carol"]="abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon address"
    ["david"]="abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon age"
    ["emma"]="abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon air"
    ["frank"]="abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon all"
    ["grace"]="abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon always"
    ["henry"]="abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon amateur"
    ["isabella"]="abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon ancient"
    ["jack"]="abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon anger"
)

# Function to create wallet
create_wallet() {
    local name=$1
    local mnemonic=$2
    
    echo -e "${BLUE}Creating wallet for $name...${NC}"
    
    # Delete existing key if it exists
    skillchaind keys delete $name --yes 2>/dev/null || true
    
    # Create new key from mnemonic
    echo "$mnemonic" | skillchaind keys add $name --recover --keyring-backend test
    
    # Get address
    local address=$(skillchaind keys show $name -a --keyring-backend test)
    echo -e "${GREEN}✅ Wallet created: $name -> $address${NC}"
    
    return 0
}

# Function to fund wallet
fund_wallet() {
    local name=$1
    local address=$2
    
    echo -e "${BLUE}Funding wallet $name ($address)...${NC}"
    
    # Fund with faucet
    skillchaind tx bank send alice $address 10000000uskill --chain-id $CHAIN_ID --keyring-backend test --yes --fees 1000uskill
    
    sleep 2
    
    # Check balance
    local balance=$(skillchaind query bank balances $address --chain-id $CHAIN_ID --output json | jq -r '.balances[0].amount // "0"')
    echo -e "${GREEN}✅ Funded: $name has $balance uskill${NC}"
}

# Function to create profile
create_profile() {
    local name=$1
    local display_name=$2
    local bio=$3
    local location=$4
    local website=$5
    local github=$6
    local linkedin=$7
    local twitter=$8
    
    echo -e "${BLUE}Creating profile for $display_name...${NC}"
    
    skillchaind tx profile create-profile \
        "$display_name" \
        "$bio" \
        "$location" \
        "$website" \
        "$github" \
        "$linkedin" \
        "$twitter" \
        "QmZSpwmV3dfwVDVcaJmdga3VVW15SEgXQDY1wiEj8gzpqc" \
        --from $name \
        --chain-id $CHAIN_ID \
        --keyring-backend test \
        --yes \
        --fees 1000uskill
    
    sleep 3
    echo -e "${GREEN}✅ Profile created for $display_name${NC}"
}

# Function to add skill
add_skill() {
    local name=$1
    local skill_name=$2
    local proficiency=$3
    local years=$4
    local address=$5
    
    echo -e "${BLUE}Adding skill $skill_name for $name...${NC}"
    
    # Generate unique skill index
    local skill_index="skill-$(date +%s)-$name-$(echo $skill_name | tr ' ' '-' | tr '[:upper:]' '[:lower:]')"
    
    skillchaind tx profile create-user-skill \
        "$skill_index" \
        "$address" \
        "$skill_name" \
        "$proficiency" \
        $years \
        false \
        "" \
        0 \
        0 \
        --from $name \
        --chain-id $CHAIN_ID \
        --keyring-backend test \
        --yes \
        --fees 1000uskill
    
    sleep 2
    echo -e "${GREEN}✅ Added skill: $skill_name ($proficiency)${NC}"
}

# Main setup function
setup_demo_profiles() {
    echo -e "${YELLOW}📝 Step 1: Creating wallets...${NC}"
    
    # Create all wallets
    for name in "${!WALLETS[@]}"; do
        create_wallet $name "${WALLETS[$name]}"
    done
    
    echo ""
    echo -e "${YELLOW}💰 Step 2: Funding wallets...${NC}"
    
    # Fund all wallets (alice should already have funds from genesis)
    for name in "${!WALLETS[@]}"; do
        if [ "$name" != "alice" ]; then
            local address=$(skillchaind keys show $name -a --keyring-backend test)
            fund_wallet $name $address
        fi
    done
    
    echo ""
    echo -e "${YELLOW}👤 Step 3: Creating profiles...${NC}"
    
    # Alice Johnson - Full-stack Developer
    create_profile "alice" "Alice Johnson" "Full-stack developer with 5 years of experience in React and Node.js. Passionate about creating scalable web applications." "San Francisco, USA" "https://alicejohnson.dev" "alice-johnson-dev" "alice-johnson-sf" "alice_codes"
    
    # Bob Chen - Blockchain Developer  
    create_profile "bob" "Bob Chen" "Blockchain developer and smart contract specialist. Building the future of decentralized applications." "Singapore" "https://bobchen.crypto" "bob-chen-blockchain" "bob-chen-crypto" "bob_blockchain"
    
    # Carol Martinez - UX/UI Designer
    create_profile "carol" "Carol Martinez" "UX/UI Designer with a passion for creating intuitive and beautiful user experiences. 6 years in the industry." "Barcelona, Spain" "https://carolmartinez.design" "carol-design" "carol-martinez-ux" "carol_designs"
    
    # David Kim - DevOps Engineer
    create_profile "david" "David Kim" "DevOps engineer specializing in cloud infrastructure and CI/CD pipelines. AWS and Kubernetes expert." "Seoul, South Korea" "https://davidkim.cloud" "david-kim-devops" "david-kim-cloud" "david_devops"
    
    # Emma Wilson - Data Scientist
    create_profile "emma" "Emma Wilson" "Data scientist and machine learning engineer. Turning data into actionable insights for businesses." "London, UK" "https://emmawilson.ai" "emma-wilson-ml" "emma-wilson-data" "emma_data_sci"
    
    # Frank Rodriguez - Mobile Developer
    create_profile "frank" "Frank Rodriguez" "Mobile app developer with expertise in both iOS and Android platforms. Creating amazing mobile experiences." "Mexico City, Mexico" "https://frankrodriguez.mobile" "frank-mobile-dev" "frank-rodriguez-mobile" "frank_mobile"
    
    # Grace Zhang - Cybersecurity Specialist
    create_profile "grace" "Grace Zhang" "Cybersecurity specialist with focus on penetration testing and security audits. Keeping systems safe." "Toronto, Canada" "https://gracezhang.security" "grace-security" "grace-zhang-security" "grace_security"
    
    # Henry Thompson - Game Developer
    create_profile "henry" "Henry Thompson" "Game developer with passion for creating immersive gaming experiences. Unity and Unreal Engine expert." "Austin, USA" "https://henrythompson.games" "henry-game-dev" "henry-thompson-games" "henry_games"
    
    # Isabella Garcia - Digital Marketing Specialist
    create_profile "isabella" "Isabella Garcia" "Digital marketing specialist with expertise in SEO, social media marketing, and content strategy." "Madrid, Spain" "https://isabellagarcia.marketing" "isabella-marketing" "isabella-garcia-marketing" "isabella_mkt"
    
    # Jack Brown - Product Manager
    create_profile "jack" "Jack Brown" "Product manager with experience in agile methodologies and product strategy. Building products users love." "Sydney, Australia" "https://jackbrown.pm" "jack-product-manager" "jack-brown-pm" "jack_pm"
    
    echo ""
    echo -e "${YELLOW}🛠️ Step 4: Adding skills...${NC}"
    
    # Alice's skills
    local alice_addr=$(skillchaind keys show alice -a --keyring-backend test)
    add_skill "alice" "React" "Expert" 4 $alice_addr
    add_skill "alice" "Node.js" "Advanced" 3 $alice_addr
    add_skill "alice" "TypeScript" "Advanced" 3 $alice_addr
    add_skill "alice" "GraphQL" "Intermediate" 2 $alice_addr
    
    # Bob's skills
    local bob_addr=$(skillchaind keys show bob -a --keyring-backend test)
    add_skill "bob" "Solidity" "Expert" 4 $bob_addr
    add_skill "bob" "Go" "Advanced" 3 $bob_addr
    add_skill "bob" "Cosmos SDK" "Advanced" 2 $bob_addr
    add_skill "bob" "Web3" "Expert" 4 $bob_addr
    
    # Carol's skills
    local carol_addr=$(skillchaind keys show carol -a --keyring-backend test)
    add_skill "carol" "Figma" "Expert" 5 $carol_addr
    add_skill "carol" "Adobe Creative Suite" "Expert" 6 $carol_addr
    add_skill "carol" "User Research" "Advanced" 4 $carol_addr
    add_skill "carol" "Prototyping" "Advanced" 5 $carol_addr
    
    # David's skills
    local david_addr=$(skillchaind keys show david -a --keyring-backend test)
    add_skill "david" "AWS" "Expert" 5 $david_addr
    add_skill "david" "Kubernetes" "Advanced" 4 $david_addr
    add_skill "david" "Docker" "Expert" 5 $david_addr
    add_skill "david" "Terraform" "Advanced" 3 $david_addr
    
    # Emma's skills
    local emma_addr=$(skillchaind keys show emma -a --keyring-backend test)
    add_skill "emma" "Python" "Expert" 6 $emma_addr
    add_skill "emma" "Machine Learning" "Expert" 5 $emma_addr
    add_skill "emma" "TensorFlow" "Advanced" 4 $emma_addr
    add_skill "emma" "SQL" "Expert" 6 $emma_addr
    
    # Frank's skills
    local frank_addr=$(skillchaind keys show frank -a --keyring-backend test)
    add_skill "frank" "Swift" "Expert" 5 $frank_addr
    add_skill "frank" "Kotlin" "Advanced" 4 $frank_addr
    add_skill "frank" "React Native" "Advanced" 3 $frank_addr
    add_skill "frank" "Flutter" "Intermediate" 2 $frank_addr
    
    # Grace's skills
    local grace_addr=$(skillchaind keys show grace -a --keyring-backend test)
    add_skill "grace" "Penetration Testing" "Expert" 4 $grace_addr
    add_skill "grace" "Network Security" "Advanced" 5 $grace_addr
    add_skill "grace" "Python" "Advanced" 4 $grace_addr
    add_skill "grace" "Linux" "Expert" 6 $grace_addr
    
    # Henry's skills
    local henry_addr=$(skillchaind keys show henry -a --keyring-backend test)
    add_skill "henry" "Unity" "Expert" 6 $henry_addr
    add_skill "henry" "C#" "Expert" 7 $henry_addr
    add_skill "henry" "Unreal Engine" "Advanced" 3 $henry_addr
    add_skill "henry" "3D Modeling" "Intermediate" 4 $henry_addr
    
    # Isabella's skills
    local isabella_addr=$(skillchaind keys show isabella -a --keyring-backend test)
    add_skill "isabella" "SEO" "Expert" 5 $isabella_addr
    add_skill "isabella" "Google Analytics" "Advanced" 4 $isabella_addr
    add_skill "isabella" "Social Media Marketing" "Expert" 6 $isabella_addr
    add_skill "isabella" "Content Strategy" "Advanced" 4 $isabella_addr
    
    # Jack's skills
    local jack_addr=$(skillchaind keys show jack -a --keyring-backend test)
    add_skill "jack" "Product Management" "Expert" 5 $jack_addr
    add_skill "jack" "Agile" "Expert" 6 $jack_addr
    add_skill "jack" "User Story Mapping" "Advanced" 4 $jack_addr
    add_skill "jack" "Data Analysis" "Advanced" 3 $jack_addr
}

# Function to test search functionality
test_search() {
    echo ""
    echo -e "${YELLOW}🔍 Step 5: Testing search functionality...${NC}"
    
    echo -e "${BLUE}Testing search by name (Alice)...${NC}"
    curl -s "http://localhost:1317/skillchain/v1/profiles/search?name=Alice" | jq '.total_count'
    
    echo -e "${BLUE}Testing search by location (USA)...${NC}"
    curl -s "http://localhost:1317/skillchain/v1/profiles/search?location=USA" | jq '.total_count'
    
    echo -e "${BLUE}Testing search by skill (Python)...${NC}"
    curl -s "http://localhost:1317/skillchain/v1/profiles/search?skill=Python" | jq '.total_count'
    
    echo -e "${BLUE}Testing advanced search (React + USA)...${NC}"
    curl -s "http://localhost:1317/skillchain/v1/profiles/search?name=&location=USA&skill=React&type=advanced" | jq '.total_count'
}

# Function to display final summary
show_summary() {
    echo ""
    echo -e "${GREEN}🎉 Demo profiles setup completed!${NC}"
    echo ""
    echo -e "${YELLOW}📊 Summary:${NC}"
    echo "• 10 demo profiles created"
    echo "• 40 skills added across all profiles"
    echo "• All wallets funded with SKILL tokens"
    echo "• Search functionality ready for testing"
    echo ""
    echo -e "${YELLOW}🔗 Useful endpoints:${NC}"
    echo "• List all profiles: curl http://localhost:1317/skillchain/v1/profiles"
    echo "• Search by name: curl 'http://localhost:1317/skillchain/v1/profiles/search?name=Alice'"
    echo "• Search by skill: curl 'http://localhost:1317/skillchain/v1/profiles/search?skill=React'"
    echo "• Search by location: curl 'http://localhost:1317/skillchain/v1/profiles/search?location=USA'"
    echo ""
    echo -e "${RED}⚠️  IMPORTANT: These are demo wallets - DO NOT use in production!${NC}"
}

# Main execution
main() {
    setup_demo_profiles
    sleep 5
    test_search
    show_summary
}

# Check if blockchain is running
if ! curl -s http://localhost:1317/skillchain/v1/params > /dev/null; then
    echo -e "${RED}❌ Error: Blockchain is not running on localhost:1317${NC}"
    echo "Please start the blockchain first: skillchaind start --api.enable --api.address tcp://0.0.0.0:1317"
    exit 1
fi

# Run main function
main 