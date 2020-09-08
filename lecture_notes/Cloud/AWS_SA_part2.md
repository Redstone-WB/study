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