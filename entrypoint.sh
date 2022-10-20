#!/bin/sh

# Variables
PM_IP="${PM_IP:-"192.168.1.100"}"
PM_PORT="${PM_PORT:-"8006"}"
PM_USER="${PM_USER:-"ubuntu"}"
PM_TOKENID="${PM_TOKENID:-"myUser@pve!myToken"}"
PM_TOKENSECRET="${PM_TOKENSECRET:-"44a2085e-f9ae-41c7-a595-d66ece971203"}"

# Run inventory builder
/pm_inventory_builder -url https://$PM_IP:$PM_PORT -allow-insecure-tls -ansible-user $PM_USER -tokenId $PM_TOKENID -tokenSecret $PM_TOKENSECRET > /data/my_proxmox_inventory.yaml

exit 0
