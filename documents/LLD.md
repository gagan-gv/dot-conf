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
3. App
   - ID (String)
   - Name (String)
   - Owner IDs (List(String))
   - Company ID (Long)
   - App Key (String)
   - Created By (String)
   - Created On (Date Time)
   - Modified By (String)
   - Modified On (Date Time)
4. Config
   - ID (String)
   - Name (String)
   - Description (String)
   - Type (TYPE)
   - Value (String)
   - App ID (String)
   - Created By (String)
   - Created On (Date Time)
   - Modified By (String)
   - Modified On (Date Time)

## Enums

1. USER STATUS
   - ACTIVATED
   - DEACTIVTED
2. ROLE
   - USER
   - ADMIN
   - SUPER_ADMIN
3. TYPE
   - STRING
   - BOOLEAN
   - OBJECT
   - NUMERIC

## Methods

1. Company
   - Register
   ```go
   register(CompanyDetails, AdminDetails) {
    Check if company is already registered or not based on email
    Register company
    Register Admin user
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

3. App
    - Add
    ```go
    add(ServiceDetails) {
     Check if app exists with the same name for the company
     Add service
    }
    ```
    - Delete
    ```go
    delete() {
     Check if service exists with the appKey
     Delete only if the user is one of the owners
     Delete all configs
     Mail all owners
    }
    ```
    - Update (Only to update name and owners)
    ```go
    update(ServiceDetails) {
     Check if service exists with the appKey
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

4. Configs
   - Add
   ```go
   add(ConfigDetails, AppId) {
    Check if config exists with same name in the app
    add config
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
   update(ConfigId) {
    Check if config exists with id
    update config
    mail the owners
   }
   ```
   - Fetch All
   ```go
   fetch(AppId) {
    Fetch all configs for the app Id
   }
   ```
   - Fetch
   ```go
   fetch(AppKey, ConfigName) {
    Check if config exists by name in the app
    return value or complete details
   }
   ```
