#!/bin/bash

# Start the chain
echo "🚀 Starting SkillChain with proper fee configuration..."
ignite chain serve --reset-once &

# Wait for the chain to start
echo "⏰ Waiting for chain to initialize..."
sleep 10

# Update gas prices
echo "💰 Setting minimum gas prices..."
sed -i '' 's/minimum-gas-prices = "0uskill"/minimum-gas-prices = "0.001uskill"/' ~/.skillchain/config/app.toml

# Restart the node to apply changes
echo "🔄 Restarting node with fee configuration..."
pkill skillchaind
sleep 2

# Start with new config (without reset)
ignite chain serve --skip-proto &

echo "✅ SkillChain started with transaction fees enabled!"
echo "💡 Fee rate: 0.001uskill per gas unit"
echo "🌍 Tendermint node: http://0.0.0.0:26657"
echo "🌍 Blockchain API: http://0.0.0.0:1317"
echo "🌍 Token faucet: http://0.0.0.0:4500"

# Keep script running
wait 