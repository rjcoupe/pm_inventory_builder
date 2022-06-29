# pm_inventory_builder
Builds an Ansible inventory from the Proxmox API, grouping VMs by tags

## Usage

```
$ pm_inventory_builder -help
Usage of pm_inventory_builder:
  -allow-insecure-tls
        Allow insecure TLS communication with Proxmox
  -ansible-user string
        SSH user on which Ansible should attempt to connect
  -apiPassword string
        Proxmox Password. Can also be set (and is recommended as such) via the PROXMOX_API_PASSWORD environment variable
  -apiUser string
        Proxmox User. Can also be set via the PROXMOX_API_USERNAME environment variable
  -tokenId string
        Proxmox Token ID - if this is set, username/password parameters are ignored. Can also be set via the PROXMOX_TOKEN_ID environment variable
  -tokenSecret string
        Proxmox Token Secret. Can also be set (and is recommended as such) via the PROXMOX_TOKEN_SECRET environment variable
  -url string
        Proxmox API URL (default "https://localhost:8006")
```

### Example
`$ pm_inventory_builder -url https://192.168.1.100:8006 -allow-insecure-tls -ansible-user ubuntu -tokenId 'myUser@pve!myToken' -tokenSecret '44a2085e-f9ae-41c7-a595-d66ece971203' > my_proxmox_inventory.yaml`
