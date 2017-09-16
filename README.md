# SSH Configuration Parser

Simple parser for some fields in an ssh-config file.

It extracts these 3 fields:
- Host
- Hostname
- User

The 'parser', if it's even fair to call it that, is extremely dumb. It is not structurally aware, and relies on simple literal matching.

If you want to adjust this for your own needs, then you will potentially want to change these lines (I use 2-spaces for indentation):

```go
if len(line) > 10 && line[0:11] == "  HostName " {
    name = line[11:len(line)]
// ...
if len(line) > 6 && line[0:7] == "  User " {
    user = line[7:len(line)]
// ...
if len(line) > 3 && line[0:4] == "Host" {
    host = line[5:len(line)]
```

The output is a simple list, like this:

```
    root@ap%h.prod.companynames.com       - Aliases [p? !pd !ps                                                ]
    root@store1.prod.companynames.com     - Aliases [ps  s1                                                    ]
    root@app%h.test.companynames.com      - Aliases [1 2 3                                                     ]
    root@%h.companynames.com              - Aliases [app1.test app2.test app3.test                             ]
    root@%h                               - Aliases [app*                                                      ]
    root@log1                             - Aliases [log1 lg1 log lg l1 l                                      ]
    root@app%h                            - Aliases [4 5 6 7 8                                                 ]
    root@lb1.companynames.com             - Aliases [lb1 lb                                                    ]
    root@lb1.test.companynames.com        - Aliases [tlb1 tlb                                                  ]
    root@lb1.prod.companynames.com        - Aliases [plb1 prlb plb                                             ]
  robert@10.0.0.79                        - Aliases [corner                                                    ]
    root@%h.companynames.com              - Aliases [*.test                                                    ]
    root@web1.web.companynames.com        - Aliases [web1 w1                                                   ]
    root@db1.test.companynames.com        - Aliases [tprddb01 tprddb1 tprddb tprdd01 tprdd1 tprdd1 tpd1 tdp tpd]
    root@int01.prod                       - Aliases [intp                                                      ]
    root@int01                            - Aliases [int                                                       ]
    root@db1.prod.companynames.com        - Aliases [prddb01 prddb1 prddb prdd01 prdd1 prdd1 pd1 dp pd         ]
   build@build1.build.companynames.com    - Aliases [build1 bld1 b1 build01 bld01 b01                          ]
   build@build2.build.companynames.com    - Aliases [b build2 bld2 b2 build02 bld02 b02                        ]
    root@registry1.build.companynames.com - Aliases [drg rg registry registry1 drg1 rg1                        ]
    root@smtp1.prod.companynames.com      - Aliases [m1 smtp1 smtp mail mail1                                  ]
ec2-user@%h.companyx.com.au               - Aliases [demo uat                                                  ]
```

## Running

```sh
# Running
go run parseSSHConfig.go
# Building
go build parseSSHConfig.go
```
