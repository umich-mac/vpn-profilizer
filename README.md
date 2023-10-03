# VPN Profilizer

Do you find yourself needing to create dozens of Cisco ~~AnyConnect~~ Secure Client VPN configuration profiles? Do they all look the same except for the group name? Congratulations, you've come to the right place!

## How To Use This

0. Download an appropriate release from the [releases page](https://github.com/umich-mac/vpn-profilizer/releases/latest).
1. A valid `pkcs12` file for signing your profile. Presumably you don't want Jamf to mess with it, which is why it needs to be signed.
2. A CSV file with four columns in this order. The header row is **required**, but the column names are not important.
   * Display Name (what your users will see, and what the resulting file will be named)
   * Remote address of the VPN server, e.g., `vpn.northwinds.contoso.com`
   * The Group Name, e.g., `funnelize-tunnelize-all-traffic`
   * The Shared Secret, e.g., `PacketMangler42!`
3. Modify the `template.configprofile` if desired.
4. Run it: `./vpn-profilizer -certificate path/to/your/dot.p12 -password the-p12-password -csv path/to/the.csv` (Leaving -password off will prompt you for the password)
5. Upload the resulting `.mobileconfig`s to your MDM and deplode away.

Note: the embedded profile ID UUIDs are not preserved between runs, but since you can't replace an uploaded signed profile in Jamf, that was not a concern for us.
