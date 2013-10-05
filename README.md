# Worker for sysbuilder

Running:

    ./bin/buildworker -queues=builder -uri=redis://192.168.1.20:6379

Test:

    redis-cli -r 2 RPUSH resque:queue:builder '{"class":"CreateVM","args":["myhostname.example.com"]}'

