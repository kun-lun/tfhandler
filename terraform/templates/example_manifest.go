package templates

var example_artifact = []byte(`
---
schema: v0.1

env_name: kunlun
resource_group_name: tosikunlun3
location: eastus

iaas: azure # this can be azure or some other platform

vm_groups:
  - name: jumpbox
    count: 1
    sku: Standard_B1s
    type: vm
    os_profile:
      admin_name: kunlun
      linux_configuration:
        ssh:
          public_keys:
            - "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDIqA9ospJDZzttzJZk5OBDC9Qo/Y63KDvhJb9t0ix1nC8v2r2bf2uEUQ4FII4cyh6lT9d+qz0sgel09X9VCUhYd8laT/0pO59/VvWr0p6iyyF9M2ai9Z22VAF3YgbZnu2TyLpS/2b8dNDtexQdL7bS/HVbDcpMo/7FQWoAE0oStUVEcPurjCcU+X+gw7N7zI7S3OdWKYEY0Q7WtSJmv9YoU9fmJ2/Krw7ElJYGbvyezlRQRxvyQxPjl61M8XJJizyhw96VP+v12mBe2Kg0tCxPrR0yTGkMVIT/6qfy2zu+YRAO2say8INQF79ZcYhOGcqQvqwDuI62gCU5PHfusTHX"
    storage:
      image:
        offer: UbuntuServer
        publisher: Canonical
        sku: 16.04-LTS
        version: latest
      os_disk:
        managed_disk_type: Standard_LRS
        caching: ReadWrite
        create_option: FromImage
      data_disks:
        - managed_disk_type: Standard_LRS
          caching: ReadWrite
          create_option: Empty
          disk_size_gb: 100
        # - name_convension: uuid # this canbe uuid/group, if uuid, one uuid would be gnerated as name
    networks:
      - subnet_name: snet-1
        network_security_group_name: ssh_nsg
        public_ip: static
        outputs:
          - ip: 192.168.0.3
            public_ip: 202.120.40.101
            host: bananastick.southeast.azure.com	
    roles:
      - name: builtin/jumpbox
        vars:
          jumpbox_public_key: (( jumpbbox_public_key ))

  - name: web-servers
    count: 3
    sku: Standard_B1s
    type: vm # this can be vm or vmss
    os_profile:
      admin_name: kunlun
      linux_configuration:
        ssh:
          public_keys:
            - "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDIqA9ospJDZzttzJZk5OBDC9Qo/Y63KDvhJb9t0ix1nC8v2r2bf2uEUQ4FII4cyh6lT9d+qz0sgel09X9VCUhYd8laT/0pO59/VvWr0p6iyyF9M2ai9Z22VAF3YgbZnu2TyLpS/2b8dNDtexQdL7bS/HVbDcpMo/7FQWoAE0oStUVEcPurjCcU+X+gw7N7zI7S3OdWKYEY0Q7WtSJmv9YoU9fmJ2/Krw7ElJYGbvyezlRQRxvyQxPjl61M8XJJizyhw96VP+v12mBe2Kg0tCxPrR0yTGkMVIT/6qfy2zu+YRAO2say8INQF79ZcYhOGcqQvqwDuI62gCU5PHfusTHX"
    storage:
      image:
        offer: UbuntuServer
        publisher: Canonical
        sku: 16.04-LTS
        version: latest
      os_disk:
        managed_disk_type: Standard_LRS
        caching: ReadWrite
        create_option: FromImage
      data_disks:
        - managed_disk_type: Standard_LRS
          caching: ReadWrite
          create_option: Empty
          disk_size_gb: 100
        # - name_convension: uuid # this canbe uuid/group, if uuid, one uuid would be gnerated as name
    networks:
      - subnet_name: snet-1
        load_balancer_name: load_balancer_1
        load_balancer_backend_address_pool_name: backend_address_pool_1
        network_security_group_name: webserver_nsg
        outputs:
          - ip: 192.168.0.3
            public_ip: 202.120.40.101
            host: bananastick.southeast.azure.com
    roles: # this is used to generate the deployment scripts
      - name: builtin/lamp_web_role
  - name: gluster-fs-servers
    count: 2
    sku: Standard_B1s
    type: vm
    os_profile:
      admin_name: kunlun
      linux_configuration:
        ssh:
          public_keys:
            - "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDIqA9ospJDZzttzJZk5OBDC9Qo/Y63KDvhJb9t0ix1nC8v2r2bf2uEUQ4FII4cyh6lT9d+qz0sgel09X9VCUhYd8laT/0pO59/VvWr0p6iyyF9M2ai9Z22VAF3YgbZnu2TyLpS/2b8dNDtexQdL7bS/HVbDcpMo/7FQWoAE0oStUVEcPurjCcU+X+gw7N7zI7S3OdWKYEY0Q7WtSJmv9YoU9fmJ2/Krw7ElJYGbvyezlRQRxvyQxPjl61M8XJJizyhw96VP+v12mBe2Kg0tCxPrR0yTGkMVIT/6qfy2zu+YRAO2say8INQF79ZcYhOGcqQvqwDuI62gCU5PHfusTHX"
    storage:
      image:
        offer: UbuntuServer
        publisher: Canonical
        sku: 16.04-LTS
        version: latest
      os_disk:
        managed_disk_type: Standard_LRS
        caching: ReadWrite
        create_option: FromImage
      data_disks:
        - managed_disk_type: Standard_LRS
          caching: ReadWrite
          create_option: Empty
          disk_size_gb: 100
    networks:
      - subnet_name: snet-1
        public_ip: static
        network_security_group_name: ssh_nsg

  - name: redis-cache-servers
    count: 1
    sku: Standard_B1ms
    type: vm
    os_profile:
      admin_name: kunlun
      linux_configuration:
        ssh:
          public_keys:
            - "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDIqA9ospJDZzttzJZk5OBDC9Qo/Y63KDvhJb9t0ix1nC8v2r2bf2uEUQ4FII4cyh6lT9d+qz0sgel09X9VCUhYd8laT/0pO59/VvWr0p6iyyF9M2ai9Z22VAF3YgbZnu2TyLpS/2b8dNDtexQdL7bS/HVbDcpMo/7FQWoAE0oStUVEcPurjCcU+X+gw7N7zI7S3OdWKYEY0Q7WtSJmv9YoU9fmJ2/Krw7ElJYGbvyezlRQRxvyQxPjl61M8XJJizyhw96VP+v12mBe2Kg0tCxPrR0yTGkMVIT/6qfy2zu+YRAO2say8INQF79ZcYhOGcqQvqwDuI62gCU5PHfusTHX"
    storage:
      image:
        offer: UbuntuServer
        publisher: Canonical
        sku: 16.04-LTS
        version: latest
      os_disk:
        managed_disk_type: Standard_LRS
        caching: ReadWrite
        create_option: FromImage
      data_disks:
        - managed_disk_type: Standard_LRS
          caching: ReadWrite
          create_option: Empty
          disk_size_gb: 100
    networks:
      - subnet_name: snet-1
        network_security_group_name: ssh_nsg
    roles: # this is used to generate the deployment scripts
      - name: builtin/redis_server

vnets:
  - name: vnet-1
    address_space: "10.0.0.0/16"
    subnets:
      - name: snet-1
        range: 10.0.0.0/24
        gateway: 10.0.0.1

load_balancers:
  - name: load_balancer_1
    sku: Standard
    backend_address_pools:
      - name: backend_address_pool_1
    health_probes:
      - name: http_probe
        protocol: Http # optional values: [Tcp, Http, Https]
        port: 80
        request_path: "/" # required if protocol is set to Http or Https. Otherwise, it is not allowed.
      - name: ssh_probe
        protocol: Tcp
        port: 22
    rules:
      - name: http_rule
        protocol: Tcp # optional values: [Tcp, Udp, All]
        frontend_port: 80
        backend_port: 80
        backend_address_pool_name: backend_address_pool_1
        health_probe_name: http_probe

network_security_groups:
  - name: webserver_nsg
    network_security_rules:
      - name: allow-http
        priority: 200
        direction: Inbound
        access: Allow
        protocol: Tcp
        source_port_range: "*"
        destination_port_range: 80
        source_address_prefix: "*"
        destination_address_prefix: "*"
  - name: ssh_nsg
    network_security_rules:
      - name: allow-ssh
        priority: 100
        direction: Inbound
        access: Allow
        protocol: Tcp
        source_port_range: "*"
        destination_port_range: 22
        source_address_prefix: "*"
        destination_address_prefix: "*"

mysql_databases:
  - migrate_from: # indicate where we should migrate from.
      origin_host: asd
      origin_database: asd
      origin_username: asd
      origin_password: asd
    name: "kunlundb"
    version: "5.7"
    cores: 2
    tier: GeneralPurpose # Possible values: ["Basic", "GeneralPurpose", "MemoryOptimized"]
    family: Gen5 # Possible values: ["Gen4", "Gen5"]
    storage: 5 # storage in GB, optional values from 5 to 4096
    backup_retention_days: 35
    ssl_enforcement: Enabled # Possible values: ["Enabled", "Disabled"]
    username: "kunlunuser"
    password: "abcd1234!"
`)
