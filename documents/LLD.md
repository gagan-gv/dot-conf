## Entities

1. Company
   - ID (Long)
   - Name (String)
   - Email (String)
   - Admin ID (String)
   - Registered On (Date Time)
   - Modified On (Date Time)
   - Documents Path (String)
2. User
   - ID (String)
   - Email (String)
   - Password (String)
   - Role (ROLE)
   - Company ID (Long)
   - Status (USER STATUS)
   - Registered On (Date Time)
3. Service
   - ID (String)
   - Name (String)
   - Owner IDs (List(String))
   - Company ID (Long)
   - Service Key (String)
   - Created By (String)
   - Created On (Date Time)
   - Modified By (String)
   - Modified On (Date Time)
4. Config
   - ID (String)
   - Name (String)
   - Description (String)
   - Type (TYPE)
   - Value (Any)
   - Service ID (String)
   - Created By (String)
   - Created On (Date Time)
   - Modified By (String)
   - Modified On (Date Time)
5. Config Request
   - ID (Long)
   - Service ID (String)
   - Config ID (String)
   - Approved By (String)
   - Approved On (Date Time)
   - Status (REQUEST STATUS)
6. Comments
   - ID (Long)
   - Request ID (Long)
   - Comment (String)
   - User ID (String)
   - On (Date Time)

## Enums

1. USER STATUS
   - ACTIVATED
   - DEACTIVTED
2. ROLE
   - USER
   - ADMIN
   - SUPER_ADMIN
3. REQUEST STATUS
   - APPROVED
   - OPEN
   - REJECTED

## Methods

1. Company
   - Register
   ```go
   register(CompanyDetails, AdminDetails) {
    Check if company is already registered or not based on email
    Register company
    Register Admin user
    Mail SuperAdmin
   }
   ```
   - Fetch All (Only for SuperAdmin)
   ```go
   fetch() {
    get all companies
   }
   ```
   - Fetch (Only for SuperAdmin)
   ```go
   fetch(CompanyId) {
    get company by id if present
   }
   ```
   - Update (Only for Admin)
   ```go
   update(UpdatedDetails, CompanyId) {
    get company by id if present else return
    update details
   }
   ```
   - Verify (Only for SuperAdmin)
   ```go
   verify(CompanyId) {
    get company by id if present else return
    mail the admin
   }
   ```

2. User
   - Register (Only for Admin)
   ```go
   register(UserDetails) {
    Check if user exist with same email id
    register user
    Mail user
   }
   ```
   - Deactivate (Only for Admin)
   ```go
   deactivate(Email) {
    Check if user exist with email id
    deactivate the user
    Mail Admin
   }
   ```
   - Login
   ```go
   login(LoginDetails) {
    Check if user exists with email id
    if (password is nil) {
        updatePassword
        return
    }
    verify auth
    return
   }
   ```

3. Service
    - Add
    ```go
    add(ServiceDetails) {
     Check if service exists with the same name for the company
     Add service
    }
    ```
    - Delete
    ```go
    delete() {
     Check if service exists with the service id
     Delete only if the user is one of the owners
     Delete all configs
     Mail all owners
    }
    ```
    - Update (Only to update name and owners)
    ```go
    update(ServiceDetails) {
     Check if service exists with the service id
     update details
     mail all owners
    }
    ```
    - Fetch All
    ```go
    fetch() {
     Get all services for the company id
    }
    ```
    - Fetch
    ```go
    fetch(ServiceId) {
     Check if service exists with the id and is for the company
     get all the configs for the service
    }
    ```

4. Configs
   - Add
   ```go
   add(ConfigDetails) {
    Check if config exists with same name
    add config and raise a request
    mail the owners
   }
   ```
   - Delete (Only for owners)
   ```go
   delete(ConfigId) {
    Check if config exists with id only if user is the owner of service
    mail the owners
   }
   ```
   - Update
   ```go
   update(ConfigDetails) {
    Check if config exists with id
    update config and raise a request
    mail the owners
   }
   ```
   - Fetch All
   ```go
   fetch(ServiceId) {
    Fetch all configs for the service Id
   }
   ```
   - Fetch
   ```go
   fetch(ServiceId/Key, ConfigId/Name) {
    Check if config exists by name or id in the service
    return value or complete details
   }
   ```

5. Config Request
   - Get All
   ```go
   get(ServcieId) {
    Get all configs for serviceId
   }
   ```
   - Get
   ```go
   get(RequestId) {
    Check if request exists
    fetch all comments
    return request details
   }
   ```
   - Comment
   ```go
   comment(RequestId) {
    Check if request exists
    add/edit comment
    Mail the requester and onwers
   }
   ```
   - Update (Only for onwers)
   ```go
   update(RequestId) {
    Check if request exists
    Check if request is open
    update status
   }
   ```
