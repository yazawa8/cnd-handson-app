# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [user.proto](#user-proto)
    - [CreateUserRequest](#user-CreateUserRequest)
    - [DeleteUserRequest](#user-DeleteUserRequest)
    - [DeleteUserResponse](#user-DeleteUserResponse)
    - [GetUserRequest](#user-GetUserRequest)
    - [ListUsersRequest](#user-ListUsersRequest)
    - [ListUsersResponse](#user-ListUsersResponse)
    - [UpdateUserRequest](#user-UpdateUserRequest)
    - [User](#user-User)
    - [UserResponse](#user-UserResponse)
    - [VerifyPasswordRequest](#user-VerifyPasswordRequest)
    - [VerifyPasswordResponse](#user-VerifyPasswordResponse)
  
    - [UserService](#user-UserService)
  
- [Scalar Value Types](#scalar-value-types)



<a name="user-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## user.proto



<a name="user-CreateUserRequest"></a>

### CreateUserRequest
--- CreateUser ---


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |
| email | [string](#string) |  |  |
| password | [string](#string) |  |  |
| role_id | [string](#string) |  |  |






<a name="user-DeleteUserRequest"></a>

### DeleteUserRequest
--- DeleteUser ---


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="user-DeleteUserResponse"></a>

### DeleteUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| success | [bool](#bool) |  |  |






<a name="user-GetUserRequest"></a>

### GetUserRequest
--- GetUser ---


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="user-ListUsersRequest"></a>

### ListUsersRequest
--- ListUsers ---






<a name="user-ListUsersResponse"></a>

### ListUsersResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| users | [User](#user-User) | repeated |  |






<a name="user-UpdateUserRequest"></a>

### UpdateUserRequest
--- UpdateUser ---


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| name | [string](#string) |  |  |
| email | [string](#string) |  |  |
| password | [string](#string) |  |  |
| role_id | [string](#string) |  |  |






<a name="user-User"></a>

### User
--- User ---


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| name | [string](#string) |  |  |
| email | [string](#string) |  |  |
| role_id | [string](#string) |  |  |
| created_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |
| updated_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |






<a name="user-UserResponse"></a>

### UserResponse
--- UserResponse ---


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user | [User](#user-User) |  |  |






<a name="user-VerifyPasswordRequest"></a>

### VerifyPasswordRequest
--- VerifyPassword ---


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| email | [string](#string) |  |  |
| password | [string](#string) |  |  |






<a name="user-VerifyPasswordResponse"></a>

### VerifyPasswordResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| valid | [bool](#bool) |  |  |





 

 

 


<a name="user-UserService"></a>

### UserService
================================
サービス定義
================================

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| CreateUser | [CreateUserRequest](#user-CreateUserRequest) | [UserResponse](#user-UserResponse) | ユーザーの作成 |
| UpdateUser | [UpdateUserRequest](#user-UpdateUserRequest) | [UserResponse](#user-UserResponse) | ユーザーの更新 |
| GetUser | [GetUserRequest](#user-GetUserRequest) | [UserResponse](#user-UserResponse) | ユーザーの取得 |
| ListUsers | [ListUsersRequest](#user-ListUsersRequest) | [ListUsersResponse](#user-ListUsersResponse) | ユーザー一覧の取得 |
| DeleteUser | [DeleteUserRequest](#user-DeleteUserRequest) | [DeleteUserResponse](#user-DeleteUserResponse) | ユーザーの削除 |
| VerifyPassword | [VerifyPasswordRequest](#user-VerifyPasswordRequest) | [VerifyPasswordResponse](#user-VerifyPasswordResponse) | パスワードの検証 |

 



## Scalar Value Types

| .proto Type | Notes | C++ | Java | Python | Go | C# | PHP | Ruby |
| ----------- | ----- | --- | ---- | ------ | -- | -- | --- | ---- |
| <a name="double" /> double |  | double | double | float | float64 | double | float | Float |
| <a name="float" /> float |  | float | float | float | float32 | float | float | Float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="bool" /> bool |  | bool | boolean | boolean | bool | bool | boolean | TrueClass/FalseClass |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode | string | string | string | String (UTF-8) |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str | []byte | ByteString | string | String (ASCII-8BIT) |

