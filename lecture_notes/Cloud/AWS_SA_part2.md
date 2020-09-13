### Amazon FSx

- Amazon FSx for Windows File server provides a fully managed native MS window file system so you cna easily move your Windows-based apps that require file storage to AWS. 
- Amazon FSx is built on Windows Server
- EFS와의 차이점

![ex](./img/FSX.png)

- SMB based 를 쓰고 있다면? -> Windows FSx를 storage로 쓰면 좋음
- EC2 instance OS가 Windows면 EFS 사용 불가임 (Linux만 가능)
  - NFS나 Storage sharing이 필요하면 EFS 쓰면 됨
- Amazon FSx for Lustre : fully managed file system that is optimized for compute-intensive workloads, such as high-performance computing, machine learning, media data processing workflows, and electronic design automation.
  - 데이터 크고, high-performance 필요한 경우에 선택
  - FSx for Lustre can store data directly on S3.

![ex](./img/FSX_2.png)



### HPC on AWS

- AWS 위에서 HPC 를 achieve 하기 위해 어떤 서비스를 이용할 수 있을까?
  - Data transfer
    - Snowball, Snowmobile (terabytes/petabytes)
    - AWS DataSync to store on S3, EFS, FSx for Windows , ... etc
    - Direct connect : AWS Direct Connect 는 on-premise 와 AWS와의 dedicated network connection을 만들 수 있게 하는 서비스이다. 
  - Compute and networking
    - EC2 instances (GPU or CPU optimized)
    - EC2 Fleets (Spot instances or Spot Fleets)
    - Placement groups (Cluster placement groups)
    - Enhanced networking : uses single root I/O virtualization (SR-IOV), to provide high-performance networking capabilities on supported instance types. / Good network performance를 필요로 할 때 사용
      - SR-IOV : method of device virtualization that provides higher I/O performance and lower CPU utilization
      - provides higher bandwidth, higher packer per second (PPS) performance, and consistently lower inter-instance latencies.
      - Enhanced networking 사용에 대한 additional charge는 없음
      - Elastic Network Adapters : 100 Gbps까지 가능, for supported instance types
        - intel 82599 VF interface (LEGACY) : 10 Gbps까지,, 옛날 instances
      - Elastic Fabric Adapters
        - for HPC and machine-learning apps
        - provides lower, more consistent latency and higher throughput than the TCP transport (traditionally used in cloud-based HPC systems)
        - OS-bypass 사용 가능
          - ML app들이 OS kernel을 bypass해서 EFA와 통신 가능! 훨씬 빨라짐 (lower latency)
          - windows에 지원x, Linux만 가능
  - Storage
    - Instance-attached storage
      - EBS
      - Instance Store : ephemeral
    - Network storage
      - S3 : distributed object-based storage, not a file system
      - EFS : file system, scale IOPS based on total szie, or use provisioned IOPS
      - FSx for Lustre : HPC-optimized distributed file system, backed by S3
  - Orchestration and automation
    - AWS Batch :  supports multi-node parallel jobs, multiple EC2 instances를 가동하도록 single job을 구성할 수 있음
      - Scheduling 가능
    - AWS ParallelCluster
      - Open-source cluster management tool, AWS에 HPC Cluster를 deploy하고, manage하는 tool이다.
      - uses a simple text file to model and provision all the resources needed
      - Automate creation of VPC, subnet, cluster type, and instance types

### AWS WAF

- AWS WAF is a web application firewall, 목적은 monitor HTTP and HTTPS requests that are forwarded to Amazon CloudFront, an Application Load Balancer or API Gateway
  - Application level (Layer 7) aware firewall 이다!
- also lets you control control access to your content
- AWS WAF allows 3 different behaviors
  - Allow all requests except the ones you specify
  - Block all requests except the ones you specify
  - Count the requests that match the properties you specify
- Define conditions by using characteristics of web requests such as :
  - IP addresses that requests originate from.
  - Country that requests originate from
  - Values in request headers
  - Strings that appear in requests, either specific strings or string that match regular expression (regex) patterns
  - Length of requests
  - Presence of SQL code that is likely to be malicious (known as SQL injection)
  - Presense of a script that is likely to be malicious (knwon as cross-site scripting)
- CF > Network ACLs 을 이용할 수도 있음.



## Databases On Aws

- Relational databases on AWS

  - SQL Server, Oracle, MySQL Server, PostGreSQL, Aurora, MariaDB
  - Two Key features
    - Multi-AZ : For Disaster Recovery
      - fail over is automatic, using Multi-AZ
    - Read Replicas : For performance
      - primary database 에 대한 read replica가 존재(백업디비)
      - primary db 손상시에 EC2가 read replica 를 읽도록 수동으로 조치해줘야함

- non relational db

  - collection = table
  - document = row
  - key value pairs = fields = colums
  - 장점 : 특정 doc에 대해서만 field 추가가 자유롭다, field 안에 sub-field 만드는 것이 자유롭다
    - 그만큼 consistency를 유지하는데 신경 써야함
  - Data warehousing : used for business intelligence
    - used to pull in very large and complex datasets
    - usually used by management, to do queries on data
    - Tools like cognos, Jaspersoft, SQL Server Reporting Services, Oracle Hyperion, SAP NetWeaver
    - OLTP : Online Transaction Processing
      - ex> 특정 Order number 에 대해서 Name, Date, 배송지, 배송 현황 등을 가지고 있는 row를 pull up하는데 쓰임
    - OLAP : Online Analytics Processing
      - Pulls in large numbers of records
      - ex> EMEA, Pacific 에서의 Digital Radio Product의 Net Profit을 계산하려고 하는 경우
        - Sum of Radios Sold in EMEA
        - Sum of Radios Sold in Pacific
        - Unit Cost of Radio in each region
        - Sales price of each radio
        - Sales price - unit cost

- Redshift : Amazon's Data Warehouse Soultion

  - used for business intelligence or Data Warehousing

- ElasticCache : web service that makes it easy to deploy, operate, and scale an in-memory cache in the cloud

  - speed up performance of existing databases (frequent identical queries)
  - imporve the performance of web apps
    - caching the most common web-queries
    - supports two open-source in-memory caching engines
      - memcached
      - Redis

- AWS Services 정리

  - RDS (OLTP) : SQL, MySQL, PostgreSQL, Oracle, Aurora, MariaDB
  - DynamoDB (No SQL)
  - Red Shift OLAP
  - Elasticache : memcached, Redis

- MySQL Lab Bootstrap code

  - \#!/bin/bash

    yum install httpd php php-mysql -y

    cd /var/www/html

    wget https://wordpress.org/wordpress-5.1.1.tar.gz

    tar -xzf wordpress-5.1.1.tar.gz

    cp -r wordpress/* /var/www/html/

    rm -rf wordpress

    rm -rf wordpress-5.1.1.tar.gz

    chmod -R 755 wp-content

    chown -R apache:apache wp-content

    service httpd start

    chkconfig httpd on
  
- RDS runs on virtual machines

  - However, you cannot log in to these operating systems

- Patching of the RDS OS and DB is amazon's responsibility

- RDS is not serverless

  - cf> Aurora serverless is serverless

- RDS back-up types

  - Automated Backups
    - 'retention period' 내의 어떤 point로도 DB를 recover 할 수 있음
      - second 단위로 가능
    - 'retention period' 는 1-35 days
    - Full daily snapshot & store transaction logs throughout the day
      -  Recovery를 실행하면, AWS는 the most recent daily backup을 고른 뒤, 그 날과 관련된 transaction log를 적용한다.
    - enabled by default, S3에 backupdata가 저장됨, DB사이즈와 동일한 S3 free storage를 받게 됨
    - backups are taken within a defined window
      - During the backup windows, storage I/O may be suspended while your data is backed up -> elevated latency를 경험할 수 있음 (느려짐)
  - Database snapshots
    - user initiated (= done manually)
    - orginal RDS instance를 지워도 남아 있는다. (automated backup 과의 차이점)
  - 두 방법 중 무엇을 쓰던, restored version of the DB would be a new instance, with a new DNS endpoint

- Encryption at rest : MySQL, Oracle, SQL server, PostGreSQL, MariaDB, Aurora 에 대해서 지원됨

  - AWS KMS service를 사용하게 됨
  - RDS instance가 encrypted 되면 관련된 것들도 also encrypted : storage 안의 data, read replicas, snapshots, and automated backups

- Multi-AZ : DB의 exact copy를 다른 AZ에 가지고 있을 수 있도록 함

  - AWS가 이러한 replication을 관리하고, automatically synchronize 시킨다.
  - planned database maintenance, DB instance failure, AZ failure 의 경우 : Amazon RDS will automatically failover to stanby DB
  - Disaster recovery only!
    - not primarily used for improving performance(이걸 위해서는 Read Replica를 써야함)
  - MySQL server, SQL server, Oracle, PostgreSQL, mariaDB
  - You can force a failover from one AZ to another by rebooting the RDS instance

- Read Replica : read-only copy of DBs

  - asynchronous replication from the primary RDS instance to the read replica
  - read-heavy databse workload에 주로 사용 (Used to increase performance)
  - MySQL Server, Oracle, PostgreSQL, mariaDB, Aurora
  - used for scaling! not for disaster recovery!
  - 'automatic backups' 가 켜져 있어야 read replicas를 deploy할 수 있음
  - 각 DB당 read replica를 5개까지 만들 수 있음
  - read replica에 대해서도 read replica를 만들 수 있음
  - Things to know
    - each replica는 각각의 DNS endpoint를 갖는다
    - read replica에 대해서도 Multi-AZ를 사용할 수 있음 
    - Multi-AZ source database에 대해서도 read replica를 만들 수 있음
    - Read replicas can be promoted to their own databases, and this breaks replication
    - You can have a read replica in a second region
    - can be in different regions



### DynamoDB

- NoSQL database solution
  - supports both document and key-value data models
- Stored on SSD storage --> fast!
- Spread across 3 geographically distinct data centers
- Eventual Consistent Reads (Default)
  - 1 second rule : App이 table에 write 하고 나서, 1초정도 지나면 read할때 반영됨
  - Consistency across all copies of data is usually reached within a second. Repeating a read after a short time should return the updated data (Best Read Performance)
- Strongly Consistent Reads 선택 가능
  - A strongly consistent read returns a result that reflects all writes that received a successful response prior to the read
  - 1초 이내의 차이도 중요하다면,, strongly consistent reads를 사용함





