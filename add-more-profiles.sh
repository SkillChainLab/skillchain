#!/bin/bash

# Add a few more demo profiles manually
echo "🚀 Adding more demo profiles..."

# Carol - UX Designer
echo "📝 Creating Carol Martinez (UX Designer)..."
skillchaind tx profile create-profile \
    "Carol Martinez" \
    "UX/UI Designer with a passion for creating intuitive and beautiful user experiences. 6 years in the industry." \
    "Barcelona, Spain" \
    "https://carolmartinez.design" \
    "carol-design" \
    "carol-martinez-ux" \
    "carol_designs" \
    "QmZSpwmV3dfwVDVcaJmdga3VVW15SEgXQDY1wiEj8gzpqc" \
    --from alice \
    --chain-id skillchain \
    --keyring-backend test \
    --yes \
    --fees 1000uskill

sleep 3

# David - DevOps Engineer  
echo "📝 Creating David Kim (DevOps Engineer)..."
skillchaind tx profile create-profile \
    "David Kim" \
    "DevOps engineer specializing in cloud infrastructure and CI/CD pipelines. AWS and Kubernetes expert." \
    "Seoul, South Korea" \
    "https://davidkim.cloud" \
    "david-kim-devops" \
    "david-kim-cloud" \
    "david_devops" \
    "QmZSpwmV3dfwVDVcaJmdga3VVW15SEgXQDY1wiEj8gzpqc" \
    --from alice \
    --chain-id skillchain \
    --keyring-backend test \
    --yes \
    --fees 1000uskill

sleep 3

# Emma - Data Scientist
echo "📝 Creating Emma Wilson (Data Scientist)..."
skillchaind tx profile create-profile \
    "Emma Wilson" \
    "Data scientist and machine learning engineer. Turning data into actionable insights for businesses." \
    "London, UK" \
    "https://emmawilson.ai" \
    "emma-wilson-ml" \
    "emma-wilson-data" \
    "emma_data_sci" \
    "QmZSpwmV3dfwVDVcaJmdga3VVW15SEgXQDY1wiEj8gzpqc" \
    --from alice \
    --chain-id skillchain \
    --keyring-backend test \
    --yes \
    --fees 1000uskill

sleep 3

# Frank - Mobile Developer
echo "📝 Creating Frank Rodriguez (Mobile Developer)..."
skillchaind tx profile create-profile \
    "Frank Rodriguez" \
    "Mobile app developer with expertise in both iOS and Android platforms. Creating amazing mobile experiences." \
    "Mexico City, Mexico" \
    "https://frankrodriguez.mobile" \
    "frank-mobile-dev" \
    "frank-rodriguez-mobile" \
    "frank_mobile" \
    "QmZSpwmV3dfwVDVcaJmdga3VVW15SEgXQDY1wiEj8gzpqc" \
    --from alice \
    --chain-id skillchain \
    --keyring-backend test \
    --yes \
    --fees 1000uskill

sleep 3

# Grace - Cybersecurity Specialist
echo "📝 Creating Grace Zhang (Cybersecurity Specialist)..."
skillchaind tx profile create-profile \
    "Grace Zhang" \
    "Cybersecurity specialist with focus on penetration testing and security audits. Keeping systems safe." \
    "Toronto, Canada" \
    "https://gracezhang.security" \
    "grace-security" \
    "grace-zhang-security" \
    "grace_security" \
    "QmZSpwmV3dfwVDVcaJmdga3VVW15SEgXQDY1wiEj8gzpqc" \
    --from alice \
    --chain-id skillchain \
    --keyring-backend test \
    --yes \
    --fees 1000uskill

sleep 5

echo "🛠️ Adding skills..."

# Get Alice's address for skills
ALICE_ADDR=$(skillchaind keys show alice -a --keyring-backend test)

# Add some skills to the new profiles
echo "Adding Figma skill for Carol..."
skillchaind tx profile create-user-skill \
    "skill-$(date +%s)-carol-figma" \
    $ALICE_ADDR \
    "Figma" \
    "Expert" \
    5 \
    false \
    "" \
    0 \
    0 \
    --from alice \
    --chain-id skillchain \
    --keyring-backend test \
    --yes \
    --fees 1000uskill

sleep 2

echo "Adding AWS skill for David..."
skillchaind tx profile create-user-skill \
    "skill-$(date +%s)-david-aws" \
    $ALICE_ADDR \
    "AWS" \
    "Expert" \
    5 \
    false \
    "" \
    0 \
    0 \
    --from alice \
    --chain-id skillchain \
    --keyring-backend test \
    --yes \
    --fees 1000uskill

sleep 2

echo "Adding Python skill for Emma..."
skillchaind tx profile create-user-skill \
    "skill-$(date +%s)-emma-python" \
    $ALICE_ADDR \
    "Python" \
    "Expert" \
    6 \
    false \
    "" \
    0 \
    0 \
    --from alice \
    --chain-id skillchain \
    --keyring-backend test \
    --yes \
    --fees 1000uskill

sleep 2

echo "Adding Swift skill for Frank..."
skillchaind tx profile create-user-skill \
    "skill-$(date +%s)-frank-swift" \
    $ALICE_ADDR \
    "Swift" \
    "Expert" \
    5 \
    false \
    "" \
    0 \
    0 \
    --from alice \
    --chain-id skillchain \
    --keyring-backend test \
    --yes \
    --fees 1000uskill

sleep 2

echo "Adding Penetration Testing skill for Grace..."
skillchaind tx profile create-user-skill \
    "skill-$(date +%s)-grace-pentest" \
    $ALICE_ADDR \
    "Penetration Testing" \
    "Expert" \
    4 \
    false \
    "" \
    0 \
    0 \
    --from alice \
    --chain-id skillchain \
    --keyring-backend test \
    --yes \
    --fees 1000uskill

sleep 3

echo "✅ Demo profiles added successfully!"
echo "Testing search functionality..."

echo "🔍 Search by location (Spain):"
curl -s "http://localhost:1317/skillchain/v1/profiles/search?location=Spain" | jq '.profiles[] | .displayName'

echo "🔍 Search by skill (Python):"
curl -s "http://localhost:1317/skillchain/v1/profiles/search?skill=Python" | jq '.profiles[] | .displayName'

echo "🔍 Total profiles:"
curl -s "http://localhost:1317/skillchain/v1/profiles" | jq '.profiles | length' 