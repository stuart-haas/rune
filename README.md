# RUNE

## Requirements

* Install GoLang: https://go.dev/doc/install
* Give sudo permissions to user:
  * Edit sudoers file with `sudo visudo`.
  * At the end of the file add `<user> ALL=(ALL) NOPASSWD: ALL`, replacing <user> with the username you want to give sudo privileges to.
  * Save the file and run `sudo usermod -aG sudo <user>`, replacing <user> with the username you want to give sudo privileges to.