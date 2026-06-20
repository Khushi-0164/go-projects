# 🔒 Secure Password Hasher CLI
An interactive command-line utility designed to hash passwords safely and verify plain-text inputs against existing hashes using secure terminal inputs and the Bcrypt hashing algorithm.

## 🚀 Features
* **Hash Password**: Input a password securely and get a cryptographically safe Bcrypt hash.
* **Verify Password**: Paste a Bcrypt hash and type a password to check if they match.
* **Keystroke Hiding**: Password entries are completely invisible on screen to prevent shoulder surfing.

## 🛠️ Go Concepts Demonstrated
* **Secure Console Inputs**: Utilizing `golang.org/x/term` to temporarily disable console character echoing during prompt inputs.
* **Low-Level System Files**: Translating standard inputs to file descriptor IDs using the `syscall` package.
* **Bcrypt Blowfish Hashing**: Scrambling passwords via `golang.org/x/crypto/bcrypt` which handles unique random salts and CPU cost scaling factors.
* **Constant-Time Verification**: Verifying hash values in constant computing time to defend against database side-channel timing attacks.

## 📖 How to Run
1. **Install dependencies:**
   ```bash
   bash setup.sh
   ```
2. **Hash a Password:**
   ```bash
   go run main.go hash
   ```
3. **Verify a Password:**
   ```bash
   go run main.go verify
   ```
