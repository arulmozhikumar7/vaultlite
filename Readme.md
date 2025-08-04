# VaultLite 🔐

A lightweight, secure, and offline CLI-based secrets manager built in Go.

## ✨ Features

- AES-256-GCM encryption
- Simple CLI commands (`init`, `add`, `get`, `remove`, `list`)
- Cross-platform (Linux, macOS, Windows)
- No external dependencies — everything is local
- Ideal for developers and sysadmins managing secrets in scripts or dev tools

## 📦 Installation

### 🛠 Manual

Download the binary for your platform from the [Releases](https://github.com/arulmozhikumar7/vaultlite/releases) page.

Make it executable and move it to your `PATH`:

```sh
chmod +x vaultlite
sudo mv vaultlite /usr/local/bin/vaultlite
````

## 🚀 Getting Started

### 1. Initialize Vault

```sh
vaultlite init
```

This generates an encrypted config file and sets up your local vault.

### 2. Add a secret

```sh
vaultlite add db_password mysecret123
```

### 3. Retrieve a secret

```sh
vaultlite get db_password
```

### 4. List all keys

```sh
vaultlite list
```

### 5. Remove a secret

```sh
vaultlite remove db_password
```

## 🗃 Storage Details

* Encrypted config: `~/Library/Application Support/vaultlite/config.json` (macOS) or platform-specific app data path
* Encrypted secrets: Stored in `secrets.json` alongside the config
* AES-256-GCM encryption using unique key, IV, and salt

## 🛡 Security Note

Your vault is only as secure as your machine. Use system-level security and backups to protect your local data.

