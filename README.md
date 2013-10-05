# Worker for sysbuilder

This is a simple worker to dispatch a job that creates a KVM machine with sysbuilder.

### Running

Assuming you run Redis on 192.168.1.20:

    ./bin/sysbuildworker -queues=builder -uri=redis://192.168.1.20:6379

Create a machine called `myhostname.example.com`:

    redis-cli -r 1 RPUSH resque:queue:builder '{"class":"CreateVM","args":["myhostname.example.com"]}'

Configuration is set in the config.json file. 

    "Command": "/usr/local/bin/sysbuilder/create_vm",
    "Template": "/var/lib/libvirt/images/template.qcow",
    "StoragePath": "/var/lib/libvirt/images"

Set the paths as it reflect your environment.

### Integrating

It should be possible to integrate this pretty simple with Resque or Sidekiq.
