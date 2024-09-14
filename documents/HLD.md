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
2. Config requests
## 3. Capacity Estimations and Constraints
Our application is a read-heavy application. There will be a lot of configuration fetch requests concurrently. Let's assume for every 10,000 fetch requests there will be 1 write request.

**Estimates**: 

Assuming there will be 10 companies onboarding on a weekly basis and for each company details, around ~100KB (details + onboarding documents) of data is required.
```
Company Data = (10 companies * 100KB * 30 days) / (1000 * 7 days) MB = ~4MB per month
```
Assuming there will 100K users onboarded on a monthly basis and for each user ~1KB is required.
```
User Data = (100000 Users * 1KB / 1000) MB = ~100MB per month
```
Assuming there will be 100 services being onboarded every week and each service will require ~5KB of data.
```
Service Data = (100 Services * 5KB * 30 days) / (1000 * 7 days) MB = ~2MB per month
```
Assuming there will be 1000 configurations onboarded every week and each config requires around ~2KB of data.
```
Config Data = (1000 Configs * 2KB * 30 days) / (1000 * 7 days) MB = 8MB per month
```
Now, let's assume that we inform companies about verifying the required configurations every 2 years. Since we have 1000 configurations per week which estimates to  ~4300 configurations per month. Let's also assume that only 10% of the configs will be stale and be deleted, this brings down total configurations to 3900 approx. We know that we'll require 8MB per month to store the configurations.
```
Total Space for Configs (2 yrs) = 3900 configs * 2 yrs * 12 months * 8MB = 750000MB = 750GB
```

So total space required for a year for each aspect is shared below.
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

![DB Schema](https://github.com/user-attachments/assets/3df07b29-d590-4661-b6e6-0205e1df484b)

## 5. Architecture
### 5.1 Company Registration
![Company Registration](https://github.com/gagan-gv/dot-conf/assets/60386381/90216d72-0269-4f3e-b2d6-b4256f392a8c)
### 5.2 User Registration
![User Registration](https://github.com/gagan-gv/dot-conf/assets/60386381/cce546a9-533b-4e2b-89a7-890aa57680a5)
### 5.3 Service Registration
![Service Registration](https://github.com/gagan-gv/dot-conf/assets/60386381/8c1e4831-6a55-4ce5-8e55-e0bf8bd8ff98)
### 5.4 Config Registration/Modification
![Config Registration/Modification](https://github.com/gagan-gv/dot-conf/assets/60386381/4b82e440-6710-4192-aa7a-c53d00078d5c)
### 5.5 Config access via service
![Config access via service](https://github.com/gagan-gv/dot-conf/assets/60386381/103040a4-d296-4068-9ccd-6cc121f5f12e)
## 6. Caching
We can utilize a caching mechanism to store configurations that are actively being used. These configurations can be cached for up to 30 minutes of inactivity or non-usage. An off-the-shelf caching solution such as Guava Cache can be employed for storing this data. Eviction policies can be implemented to manage the cache, where entries are evicted either based on time or when the cache storage reaches its capacity, following the Least Recently Used (LRU) eviction strategy.
## 7. Security and Permissions
The primary concern for companies is ensuring controlled access to configurations using configuration names, especially given the possibility of multiple companies having similar services (e.g., UserService). One approach to mitigate this issue is by assigning separate service keys, thereby reducing access pain points and ensuring appropriate access control. Another crucial concern pertains to the creation and storage of configurations.
