#!/bin/sh

set -o pipefail

add_account() {
  jq --arg ADDRESS "$1" --arg ASSET "$2" --arg AMOUNT "$3" '.app_state.auth.accounts += [{
        "@type": "/cosmos.auth.v1beta1.BaseAccount",
        "address": $ADDRESS,
        "pub_key": null,
        "account_number": "0",
        "sequence": "0"
    }]' <~/.thornode/config/genesis.json >/tmp/genesis.json
  mv /tmp/genesis.json ~/.thornode/config/genesis.json

  jq --arg ADDRESS "$1" --arg ASSET "$2" --arg AMOUNT "$3" '.app_state.bank.balances += [{
        "address": $ADDRESS,
        "coins": [ { "denom": $ASSET, "amount": $AMOUNT } ],
    }]' <~/.thornode/config/genesis.json >/tmp/genesis.json
  mv /tmp/genesis.json ~/.thornode/config/genesis.json
}

set_mocknet_bond_module() {
  jq '.app_state.state.accounts += [
    {
      "@type": "/cosmos.auth.v1beta1.ModuleAccount",
      "base_account": {
        "account_number": "0",
        "address": "tthor17gw75axcnr8747pkanye45pnrwk7p9c3uhzgff",
        "pub_key": null,
        "sequence": "0"
      },
      "name": "bond",
      "permissions": []
    }
  ]' <~/.thornode/config/genesis.json >/tmp/genesis.json
  mv /tmp/genesis.json ~/.thornode/config/genesis.json

  jq '.app_state.bank.balances += [
    {
      "address": "tthor17gw75axcnr8747pkanye45pnrwk7p9c3uhzgff",
      "coins": [ { "denom": "rune", "amount": "100000000" } ]
    }
  ]' <~/.thornode/config/genesis.json >/tmp/genesis.json
  mv /tmp/genesis.json ~/.thornode/config/genesis.json
}

deploy_evm_contracts() {
  for CHAIN in ETH AVAX BSC; do
    # deploy contract and get address from output
    echo "Deploying $CHAIN contracts"
    if ! python3 scripts/evm/evm-tool.py --chain $CHAIN --rpc "$(eval echo "\$${CHAIN}_HOST")" --action deploy >/tmp/evm-tool.log 2>&1; then
      cat /tmp/evm-tool.log && exit 1
    fi
    cat /tmp/evm-tool.log
    CONTRACT=$(grep </tmp/evm-tool.log "Router Contract Address" | awk '{print $NF}')

    # add contract address to genesis
    echo "$CHAIN Contract Address: $CONTRACT"
    jq --arg CHAIN "$CHAIN" --arg CONTRACT "$CONTRACT" \
      '.app_state.thorchain.chain_contracts += [{"chain": $CHAIN, "router": $CONTRACT}]' \
      ~/.thornode/config/genesis.json >/tmp/genesis.json
    mv /tmp/genesis.json ~/.thornode/config/genesis.json
  done
}

set_eth_contract() {
  jq --arg CONTRACT "$1" '.app_state.thorchain.chain_contracts = [{"chain": "ETH", "address": $CONTRACT}]' ~/.thord/config/genesis.json >/tmp/genesis.json
  mv /tmp/genesis.json ~/.thord/config/genesis.json
}

set_avax_contract() {
  jq --arg CONTRACT "$1" '.app_state.thorchain.chain_contracts = [{"chain": "AVAX", "address": $CONTRACT}]' ~/.thord/config/genesis.json >/tmp/genesis.json
  mv /tmp/genesis.json ~/.thord/config/genesis.json
}

set_bsc_contract() {
  jq --arg CONTRACT "$1" '.app_state.thorchain.chain_contracts = [{"chain": "BSC", "address": $CONTRACT}]' ~/.thord/config/genesis.json >/tmp/genesis.json
  mv /tmp/genesis.json ~/.thord/config/genesis.json
}

gen_bnb_address() {
  if [ ! -f ~/.bond/private_key.txt ]; then
    echo "Generating BNB address"
    mkdir -p ~/.bond
    # because the generate command can get API rate limited, THORNode may need to retry
    n=0
    until [ $n -ge 60 ]; do
      generate >/tmp/bnb && break
      n=$((n + 1))
      sleep 1
    done
    ADDRESS=$(grep </tmp/bnb MASTER= | awk -F= '{print $NF}')
    echo "$ADDRESS" >~/.bond/address.txt
    BINANCE_PRIVATE_KEY=$(grep </tmp/bnb MASTER_KEY= | awk -F= '{print $NF}')
    echo "$BINANCE_PRIVATE_KEY" >/root/.bond/private_key.txt
    PUBKEY=$(grep </tmp/bnb MASTER_PUBKEY= | awk -F= '{print $NF}')
    echo "$PUBKEY" >/root/.bond/pubkey.txt
    MNEMONIC=$(grep </tmp/bnb MASTER_MNEMONIC= | awk -F= '{print $NF}')
    echo "$MNEMONIC" >/root/.bond/mnemonic.txt
  fi
}

# set external ip to localhost in mocknet
if [ "$NET" = "mocknet" ]; then
  EXTERNAL_IP="$(hostname -i)"
  export EXTERNAL_IP
fi
