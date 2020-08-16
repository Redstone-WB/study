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

### Udemy Lecture

- 1 Region 안에는 2개 이상의 AZ가 있음. AZ는 1개 이상의 data center으로 구성.