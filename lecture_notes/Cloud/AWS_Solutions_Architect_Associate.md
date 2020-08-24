### Services

- The AWS Amplify CLI, to rapidly provision and configure our cloud services
- The AWS Amplify JavaScript library, to connect our front end to cloud resources
  - "As we saw above, AWS Amplify is an open source JavaScript library that makes it very easy to integrate a number of cloud services into your web or React Native apps."
  - our src/aws-exports.js file contains all of the configuration we’ll need to pass to the Amplify JS library in order to talk to the AppSync API.
- Amazon Cognito, to handle user sign up authorization
- Amazon Simple Storage Service (S3), to store and serve as many photos as our users care to upload, and to host the static assets for our app
- Amazon DynamoDB, to provide millisecond response times to API queries for album and photo data
- AWS AppSync, to host a GraphQL API for our front end
- AWS Lambda, to create photo thumbnails asynchronously in the cloud
- Amazon Rekognition, to detect relevant labels for uploaded photos
- Amazon Elasticsearch Service, to index and search our photos by their labels

## Udemy Lecture

1. Region 안에는 2개 이상의 AZ가 있음. AZ는 1개 이상의 data center으로 구성.
2. Edge Location : 컨텐츠 caching을 위해 사용되는 endpoints. CloudFront와 CDN이 있음.
   1. 이건 Region에 포함되는 더 작은 개념.

#### IAM이 제공하는 features

- centralized control of AWS account
- Shared Access to AWS account
- Granular Permissions
- Identity Federation (Active Directory, Facebook, Linkdin etc... ) : ID 통합
- Multifactor Authentication
- provide temporary access for users/devices and services
- allows you to set up your own password rotation policy

#### IAM terms

- Users / Groups / Policies (JSON 형태) / Roles
- Roles : AWS resource에 부여하는 권한
- IAM is universal ('글로벌' 개념임. 개별 region에 따로 적용X)
- New Users have No permissions, when first created
- New Users are assigned Access Key ID & Secret Access Keys, when first created
- Group에 부여된 Policy는 Group에 속하는 User들이 상속받음
- Multifactor Authentication : 2차 인증

### S3

- Object-based (0 Bytes to 5TB)
  - Key : name of the object
  - value : data, made up of sequence of bytes
  - Version ID
  - Metadata
  - Subresources : Access control lists, torrents
- unlimited storage
- Files are stored in Buckets (=Folders)
- S3 is a universal namespace (must be unique globally)
  - 성공적으로 업로드되면, HTTP 200 code를 받게 됨!
- Not suitable to install an OS on.
- (중요) How does data consistency work for S3??
  - Read after Write consistency for PUTS of new objects
  - 'Eventual' Consistency for overwrite PUTS and DELETES (can take some time to propagate)
- (중요) Tiered Storage Available
  - S3 Standard : 99.99% availability, 99.99999..% durability, stored redundantly accross multiple devices in multiple facilities, and designed to sustain the loss of 2 facilities concurrently.
  - S3 - IA (Infrequently Accessed) : rapid access는 필요하고 잘 access 하지 않는 데이터, lower fee, but retrieval fee를 내야 함
  - S3 One Zone - IA : do not require the multiple availability Zone data resilience, retrieval fee
  - S3 Intelligent Tiering : Designed to optimize costs by automatically moving data to the most cost-effective access tier, without performance impact
  - S3 Glacier : data archiving 목적, retrieval times는 조정 가능 (minutes to hours), retrieval fee
  - S3 Glacier Deep Archive : lowest-cost storage, retrieval time of 12 hours is acceptable, retrieval fee
- Lifecycle Management
  - 30일 지나면 이 파일을 다른 storage로 올려! 가능
- Encryption
- MFA Delete 가능
  - 지우려면 2 factor authentication 필요하도록
- Secure your data using Access Control Lists (individual objects에 적용 가능) and Bucket Policies (버킷 전체에 적용)
- Cross Region Replication
  - Storage를 region에서 region으로 복제 가능.
  - S3 Transfer Acceleration : CloudFront 의 edge location 이용
- (시험 전) S3 FAQs 읽어 보기 --------------------------------------------- It comes up A LOT!!
- versioning
  - Once enabled, versioning cannot be disabled, only suspended
- S3 Object Lock : to store objects, using WORM (write once, read many)
  - object locks can be on individual objects or applied across the bucket
  - governance mode & compliance mode 가능
    - governance mode : users can overwrite or delete object version or alter its lock settings, if they have special permissions
    - compliance mode : a protected object version can't be overwritten or deleted, 보호 기간 내에는 root user도 불가능함
- S3 Glacier Valut Lock : WORM policy 같은 것 적용 가능
  - 한 번 Lock 되면, change 불가능

#### S3 - Encrpytion

1. Encrpytion In Trainsit is achieved by (HTTPS 를 사용하면서 얻어지는 것)
   1. SSL / TLS
2. Encrpytion At Rest (Server Side)
   1. Server Side : Amazon이 도와줌
      1. S3 Managed keys (= SSE - S3; Serversie Encrpytion)
      2. AWS와 User이 함께 : Aws Key management Service, Managed Keys (= SSE - KMS)
      3. Server Side Encrpytion with customer provided keys (= SSE-C)
   2. Client Side Encryption

#### S3 - Performance

1. 여러 개의 prefix (bucketname과 file 사이의 folders) 사용하면 속도 up
   1. per prefix, 3,500 PUT/COPY/POST/DELETE and 5,500 GET/HEAD requests per second
2. SSE -KMS 사용하는 경우, 제약이 있음
   1. Uploading/downloading will count toward the KMS quota
   2. Region에 따라 quota는 다름. 5,500 / 10,000 / 30,000 requests per second
   3. quota increase는 요청할 수 없음. --> 바뀐 듯? https://docs.aws.amazon.com/kms/latest/developerguide/increase-quota.html 요청할 수 있음!
3. Multipart Uploads : 100MB 이상부터 사용가능, 5GB 이상은 의무
4. S3 byte-range fetches : 분할 다운로드 가능, 일부분만 다운로드 가능
5. S3 Select & Glacier select
   1. SQL query를 써서 일부만 다운로드 가능. --> performance 개선 / data transfer 에서 비용 절약 가능
      1.S3, Glacier 둘 다 사용 가능

#### AWS Organizations

- account management service, that enables consolidating multiple AWS accounts into an organization that you create and centrally manage.
- ROOT 아래에 여러 개의 OU (Organization Units)를 두고, 그 아래에 AWS accounts를 둠
  - OU는 일종의 팀, accounts는 users
  - Policy를 OU에 적용할 수도 있고, 개별 accounts에 적용할 수도 있음
  - OU 아래에 OU를 둘 수도 있음
  - OU에 적용되는 Policy는 하부 account나 OU에 inherit 됨
- Paying account (Usually Root account) should be used for billing purpose only. 다른 resources를 deploy하는 것은 권고되지 않음.
- Enable/Disable AWS services, using Service Control Policies (SCP) either on OU or on individual accounts

#### Sharing S3 buckets, between accounts

- 3 가지 방법
  - 1. Using Bucket Policies & IAM (entire bucket), programmatic access only.
  - 2. Using Bucket ACLs & IAM (individual objects), programmatic access only.
  - 3. Cross-account IAM Roles, Programmatic and Console access

#### Cross Region Replication

- Versioning must be enabled, on both the source and destination buckets
- Files in an existing bucket are not replicated automatically
- All subsequent updated files will be replicated automatically
- Delete Markers are not replicated
- Deleting individual versions or delete markers will not be replicated

#### Transfer Acceleration

- Utilize CloudFront Edge Network, to accelerate your uploads to S3
- S3 bucket에 직접 업로드 하는 것 대신에, distinct URL을 사용하여 edge location에 direct upload 한다. Edge Location에서 해당 파일을 타겟 S3로 transfer 한다.

#### Aws DataSync

- Used to move large amounts of data from on-premises to AWS
- Used with NFS- (Network File System) and SMB- (Server Message Block) cmpatible file systems
- Replication can be done hourly, daily, or weekly
- Install the 'Datasync agent' to start the replication
- Can be used to replicate EFS to EFS

#### CloudFront

- A CDN is a system of distributed servers (network) that deliver webpages and other web content to a user, based on the geographical locations of the user, the origion of the webpage, and a content delivery server
- Edge Location : This is the location where content will be cached
  - seperate to an AWS Region / AZ
  - not just READ only, you can write to them too.
  - Objects are chached for the life of the TTL (Time to Live)
  - You can clear cahced objects, but you will be charged
- Origin : S3 bucket or EC2 Instance, Elastic Load Balancer, Route53 등이 될 수 있음.
- Distribution : CDN에 부여되는 이름. CDN which consists of a collection of Edge Locations
  - Web distribution : Used for websites
  - RTMP : Used for media streaming
- CloudFront Signed URLS & Cookies
  - A signed URL is for individual files (1file = 1 URL)
  - A signed cookie is for multiple files (1 cookie = multiple files)
  - Signed URL이나 COOKIE를 만들 때, policy를 붙이게 됨. Policy는 아래 항목들을 가질 수 있음.
    - URL expiration
    - IP ranges
    - Trusted signers (어떤 Aws accounts가 singed URLs를 만들 수 있는지 지정)
  - signed URL/cookies 는 content를 보호하고 싶을 때, 권한을 부여받은 사람에게만 접근 권한을 주고 싶을 때 사용 (ex. 넷플릭스)
  - Origin이 EC2라면? use CloudFront (signed URLs/Cookies)
  - cf> S3 signed URL
    - signed URL을 만든 user과 동일한 IAM 권한을 가지게 됨.
    - 따라서, S3 signed URL을 사용할 경우, S3 bucket에 접근 가능함.

#### AWS Snowball (Snow Family)

- Snowball : 대용량 데이터 Migration에 쓰이는 storage (Hardware)
  - Very Big Disk, petabyte-scale data transport solution
  - Storage 뿐만 아니라 Compute capability도 보유
  - 가능 : Import to S3, Export from S3
- Snowmobile
  - Exabyte-scale data transfer service
  - data center migration 같은데도 쓰임

#### Storage Gateway (Virtual | Physical device 둘 다 가능)

- service that connects an on-premises software appliance with cloud-based storage to provide seamless and secure integration between an organization's on-premises IT environment and AWS's storage infra.
- AWS Storage Gateway's software appliance는 VM image 형태로 다운로드가 가능하다. 이 VM image를 datacenter의 host에 설치할 수 있음.
- Gateway 설치 후, AWS account와 연동 (activation process)하면, 콘솔을 사용하여 storage gateway option을 만들어 낼 수 있음.
- Gateway 3종
  - File Gateway (NFS & SMB) : file들을 S3에 저장하기 위한 목적
  - Volume Gateway (isCSI) : Virtual 하드디스크 드라이브의 복사본을 S3에 저장하기 위한 목적
    - Stored Volumes : provide on-premise applications with low-latency access to their 'entire' datasets
      - data는 EBS snapshot 형태로 asynchronously backup 됨
      - Entire Dataset is stored 'on site', and is asynchronously backed up to S3
    - Cached Volumes
      - Entire Dataset is store 'on S3' and the most frequently accessed data is cached on site
  - Tape Gateway (VTL, Virtual Tape Library)
    - data 저장 방식으로 Tape 쓰고 있는 경우에 쓰는 것

#### Athena vs Macie

- Athena : interactive query service, that enables you to analyze query data located in S3, using standard SQL
  - pay per query / per TB scanned
  - No need to set up complex ETL processes
  - Works directly with data stored in S3
  - 어디에 쓰이나?
    - query log files, stored in S3 (e.g. ELB logs, S3 access logs, ... etc)
    - generate biz reports, on data stored in S3
    - Analyze AWS cost and usage reports
    - Run queries on click-stream data
- PII (Personally Identifiable Information) : 개인 식별 정보
- Macie : Security service (ML and NLP를 사용하는) to discover, classify, and protect sensitive data stored in S3
  - S3에 저장된 object가 PII인지 판별하기 위해 AI 사용
  - Dashboards, reporting and alerts
  - Works directly with data stored in S3
  - Can also analyze CloudTrail logs
  - Great for PCI-DSS (신용카드 결제) and preventing ID theft

### EC2

- RECIZABLE COMPUTE CAPACITY IN THE CLOUD
- On demand : Allows you to pay a fixed rate by the hour (or by the second) with no commitment
  - app with short term, spiky, or 예측 불가능한 workloads that cannot be interrupted
  - Apps being developed, or tested on Amazon EC2 for the first time
- Reserved : Provides you with a capacity reservation, and offer a significant discount on the hourly charge for an instance. 1년 or 3년 계약
  - Standard Reserved Instances : On demand에 비해 최대 75% off, 미리 많이 돈 낼수록 / 계약이 길수록 싸짐
  - Convertible Reserved Instaces : On demand에 비해 최대 54% off, CPU나 memory 등을 중간에 늘릴 수 있음.
  - Scheduled Reserved Instacnes : 특정 time window에 예약 가능.
  - Apps with steady state or predictable usage
  - Apps that require reserved capacity
- Spot
  - apps that have flexible start and end times
  - apps that are only feasible at very low compute prices
  - users with urgent computing needs for large amounts of additional capacity
  - 특징 : Amazon EC2에 의해 종료될 경우, partial hour of usage에 대해서는 청구되지 않음. 그러나, 사용자가 직접 종료시킨 경우 사용 시간만큼 charge 됨.
- Dedicated Hosts : Physical EC2 server didicated for your use. License 비용을 아낄 수 있음.
  - Regulatory requirements (ex. no multi-tenant virtualization이 요구되는 경우)
  - Great for licensing, which does not support multi-tenancy or cloud deployments
  - Can be purchased On-Demand (hourly)
  - Can be purchased as a Reservation for up to 70% off the On-demand price

#### 아래는 참고 사항 (시험 X)

![EC2_1](./img/EC2_Instance_Types_2.png)

![EC2_2](./img/EC2_Instance_Types.png)

- 'Fight DR. MC PXZ (in) AU (strailia) 로 외울 수 있음.

####
