# csv2keychain
![](https://travis-ci.com/fibbery/csv2keychain.svg?branch=master)
## About
command-line for load credentials(account/password) csv file from chrome to macos keychain

## Test Environment
1. Chrome 78.0.3904.87
2. MacOS Mojave 10.14.3
3. Go 1.12

## installing
You can install the tools via *homebrew*
```shell
$ homebrew tap fibbery/tap && homebrew install csv2keychain
```

## Notice
1. You must use the Chrome export the credentials csv file
2. The csv file must be this format :  domain,  url, account, password.  Tools will verify this.
3. Tools will select suit the normal domain item load , exclude ip 
4. Only for Mac

## How to use
```shell
$ csv2keychain -f password.csv
```

