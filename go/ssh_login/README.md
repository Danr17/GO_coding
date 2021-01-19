#ssh_login

Try to login to the devices with user and password provided.  

Works for:

* Cisco IOS, NX  
* Fortinet
* Juniper NetScreen, SSG
* F5 BIG-IP ( in interactive mode)

### Script Info:

Ciphers: "aes128-gcm@openssh.com", "chacha20-poly1305@openssh.com", "aes128-ctr", "aes192-ctr", "aes256-ctr", "aes128-cbc", "3des-cbc"
KeyExchanges: "curve25519-sha256@libssh.org", "ecdh-sha2-nistp256", "ecdh-sha2-nistp384", "ecdh-sha2-nistp521", "diffie-hellman-group14-sha1", "diffie-hellman-group1-sha1"  
Timeout: 4s  
Concurent connections: 30

### Usage:

```bash

> .\ssh_login.exe -user=your_username -file=input
Enter Password:
Trying to connect to x.x.x.x

// For interactive shells like BIG-IP
> .\ssh_login.exe -user=your_username -file=input -interactive
Enter Password:
Trying to connect to x.x.x.x
```
