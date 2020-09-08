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



