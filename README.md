# CIDR-Checker
_________ .___________ __________       .__            __
\_   ___ \|   \______ \\______   \ ____ |  |__   ____ |  | __ ___________ 
/    \  \/|   ||    |  \|       _// ___\|  |  \_/ __ \|  |/ // __ \_  __ \
\     \___|   ||    `   \    |   \  \___|   Y  \  ___/|    <\  ___/|  | \/
 \______  /___/_______  /____|_  /\___  >___|  /\___  >__|_ \\___  >__|
        \/            \/       \/     \/     \/     \/     \/    \/

A CIDR checker that can compare 2+n CIDR ranges and tell you if there is any overlap in their network ranges.

Currently written in Golang.

## Use

CIDR checker is simple to use, it accepts 1 or more ips, seperated by spaces, in [CIDR notation format](https://www.rfc-editor.org/rfc/rfc4632) and lets the user know if there are any overlapping CIDRs in that input. For example:

```
$ cidr-checker 10.0.1.0/24 10.0.2.0/24 10.0.3.0/24
2023/09/03 13:59:58 All good no overlapping CIDRs.
```

If you wish to read from stdin using cidr-checker please use xargs. For example:

```
echo "10.0.0.0/24 10.0.1.0/24 10.0.2.0/24" | xargs cidr-checker
2023/09/03 14:00:20 All good no overlapping CIDRs.
```

## Build and Install

You will need [golang's tooling to install](https://go.dev/doc/install), once install.

```
go build . 
mv cidr-checker /usr/local/bin
```

## TODO
- Add help
