# mobileme
A work-in-progress implementation of MobileMe. At the moment, authentication is assumed to be with the username `someuser` and password `testing1234`.

Some logic is derived from Apple's own code, such as from within [dotMacArchive.cpp](https://github.com/GaloisInc/hacrypto/blob/5c99d7ac73360e9b05452ac9380c1c7dc6784849/src/C/Security-57031.40.6/SecurityTests/clxutils/dotMacArchive/dotMacArchive.cpp).

Other logic is derived from [googlecodeimport-dotmac](https://github.com/audiohacked/googlecodeimport-dotmac/tree/1ddf28eee645148fb183283adc7a115d38a5d84b).

## Domain setup
You will need to have `configuration.apple.com` point to this server in some way or form.
Editing `/etc/hosts` is a fairly straightforward way to approach this.

You will also need to have another domain to be used in place of `mac.com`/`me.com`.
It will need several subdomains pointed to this server:
 - Its root domain
 - `certinfo`