## 1. Overview
The `.conf`, dynamic configuration platform, is designed to provide developers with a flexible solution for managing configurations outside of their codebases. By decoupling configurations from the application code, developers gain the ability to modify settings during runtime, reducing the need for frequent deployments and enhancing overall flexibility and security.
## 2. Requirements
### 2.1 Functional Requirements
1. Users should be able to store configurations in a desired format
1. Services should be able to access configurations only with correct service key
1. All new and modified configuration must be approved by an owner
1. A service can be deleted only by an owner
1. Application should be accessible to the user by REST APIs
1. Configuration should be accessible to the corresponding services via application packages
### 2.2 Non Functional Requirements
1. Application should be highly available
1. There should be minimal latency for service to fetch the configurations
1. Service keys should not be predictable
### 2.3 Extended Requirements
1. Analytics
## 3. Capacity Estimations and Constraints
Our application is a read-heavy application. There will be a lot of configuration fetch requests concurrently. Let's assume for every 10,000 fetch requests there will be 1 write request.
**Estimates**: 
Assuming there will be 10 companies onboarding on a weekly basis and for each company details, around ~100KB (details + onboarding documents) of data is required.
$$null$$Assuming there will 100K users onboarded on a monthly basis and for each user ~1KB is required.
$$null$$Assuming there will be 100 services being onboarded every week and each service will require ~5KB of data.
$$null$$Assuming there will be 1000 configurations onboarded every week and each config requires around ~2KB of data.
$$null$$Now, let's assume that we inform companies about verifying the required configurations every 2 years. Since we have 1000 configurations per week which estimates to  ~4300 configurations per month. Let's also assume that only 10% of the configs will be stale and be deleted, this brings down total configurations to 3900 approx. We know that we'll require 8MB per month to store the configurations.
$$null$$So total space required for a year for each aspect is shared below.
|Aspect|Space|
|-|-|
|Company|48 MB|
|User|1.2 GB|
|Service|24 MB|
|Configuration|375 GB|
|Total|~ 380 GB|
## 4. Database Design
### 4.1 Observations
1. Read-heavy system
1. We need to store millions of records
1. Each object on average is of mid size (~3KB)
1. There are relations between the records
### 4.2 Schema
Since there are multiple relations and we have to store millions of records, a SQL database like MySQL, PostgreSQL, SQLite types of database would be a great fit. A SQL type of database would be a good as there are only specific set of columns required and do not require frequent changes in schema.
![](E:\AppFlowyDataDoNotRename\images\0841946a-ae55-4e2b-8dd4-3a556452c5bb.png)
## 5. Architecture
### 5.1 Company Registration
![](E:\AppFlowyDataDoNotRename\images\1b4f0829-617b-4d77-94a2-13dd2ca49251.png)
### 5.2 User Registration
![](E:\AppFlowyDataDoNotRename\images\580462c6-7702-4a69-882c-cf16500af74c.png)
### 5.3 Service Registration
![](E:\AppFlowyDataDoNotRename\images\97a9c606-f01a-4bde-883d-dc97b51a3cb8.png)
### 5.4 Config Registration/Modification
![](E:\AppFlowyDataDoNotRename\images\273c4358-15de-4c5f-a437-8c57b35c5ebb.png)
### 5.5 Config access via service
![](E:\AppFlowyDataDoNotRename\images\8a7eb165-68c8-4212-81ef-bf97072f0215.png)
## 6. Caching
We can utilize a caching mechanism to store configurations that are actively being used. These configurations can be cached for up to 30 minutes of inactivity or non-usage. An off-the-shelf caching solution such as Guava Cache can be employed for storing this data. Eviction policies can be implemented to manage the cache, where entries are evicted either based on time or when the cache storage reaches its capacity, following the Least Recently Used (LRU) eviction strategy.
## 7. Security and Permissions
The primary concern for companies is ensuring controlled access to configurations using configuration names, especially given the possibility of multiple companies having similar services (e.g., UserService). One approach to mitigate this issue is by assigning separate service keys, thereby reducing access pain points and ensuring appropriate access control.
Another crucial concern pertains to the creation and storage of configurations. While configuration creation can be delegated to either a manager or a user, the approval of configurations could be restricted to service owners. This helps maintain oversight and ensures that only authorized configurations are stored.
