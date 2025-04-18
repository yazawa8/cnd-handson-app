# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [proto/session.proto](#proto_session-proto)
    - [AccessTokenResponse](#session-AccessTokenResponse)
    - [GenerateAccessTokenRequest](#session-GenerateAccessTokenRequest)
    - [GenerateRefreshTokenRequest](#session-GenerateRefreshTokenRequest)
    - [RefreshTokenResponse](#session-RefreshTokenResponse)
    - [RevokeRefreshTokenRequest](#session-RevokeRefreshTokenRequest)
    - [RevokeRefreshTokenResponse](#session-RevokeRefreshTokenResponse)
    - [ValidateAccessTokenRequest](#session-ValidateAccessTokenRequest)
    - [ValidateAccessTokenResponse](#session-ValidateAccessTokenResponse)
    - [ValidateRefreshTokenRequest](#session-ValidateRefreshTokenRequest)
    - [ValidateRefreshTokenResponse](#session-ValidateRefreshTokenResponse)
  
    - [AccessTokenService](#session-AccessTokenService)
    - [RefreshTokenService](#session-RefreshTokenService)
  
- [Scalar Value Types](#scalar-value-types)



<a name="proto_session-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## proto/session.proto



<a name="session-AccessTokenResponse"></a>

### AccessTokenResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| access_token | [string](#string) |  |  |
| expires_at | [int64](#int64) |  |  |






<a name="session-GenerateAccessTokenRequest"></a>

### GenerateAccessTokenRequest
--- GenerateAccessToken ---


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user_id | [string](#string) |  |  |






<a name="session-GenerateRefreshTokenRequest"></a>

### GenerateRefreshTokenRequest
--- GenerateRefreshToken ---


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user_id | [string](#string) |  |  |






<a name="session-RefreshTokenResponse"></a>

### RefreshTokenResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| refresh_token | [string](#string) |  |  |
| expires_at | [int64](#int64) |  |  |
| user_id | [string](#string) |  |  |






<a name="session-RevokeRefreshTokenRequest"></a>

### RevokeRefreshTokenRequest
--- RevokeRefreshToken ---


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| refresh_token | [string](#string) |  |  |
| user_id | [string](#string) |  |  |






<a name="session-RevokeRefreshTokenResponse"></a>

### RevokeRefreshTokenResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| success | [bool](#bool) |  |  |






<a name="session-ValidateAccessTokenRequest"></a>

### ValidateAccessTokenRequest
--- ValidateAccessToken ---


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| access_token | [string](#string) |  |  |






<a name="session-ValidateAccessTokenResponse"></a>

### ValidateAccessTokenResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| valid | [bool](#bool) |  |  |
| user_id | [string](#string) |  |  |
| error | [string](#string) |  |  |






<a name="session-ValidateRefreshTokenRequest"></a>

### ValidateRefreshTokenRequest
--- ValidateRefreshToken ---


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| refresh_token | [string](#string) |  |  |
| user_id | [string](#string) |  |  |






<a name="session-ValidateRefreshTokenResponse"></a>

### ValidateRefreshTokenResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| valid | [bool](#bool) |  |  |





 

 

 


<a name="session-AccessTokenService"></a>

### AccessTokenService
================================
サービス定義
================================

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| GenerateAccessToken | [GenerateAccessTokenRequest](#session-GenerateAccessTokenRequest) | [AccessTokenResponse](#session-AccessTokenResponse) | アクセストークンの生成（ログイン後・リフレッシュ時） |
| ValidateAccessToken | [ValidateAccessTokenRequest](#session-ValidateAccessTokenRequest) | [ValidateAccessTokenResponse](#session-ValidateAccessTokenResponse) | アクセストークンの検証 |


<a name="session-RefreshTokenService"></a>

### RefreshTokenService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| GenerateRefreshToken | [GenerateRefreshTokenRequest](#session-GenerateRefreshTokenRequest) | [RefreshTokenResponse](#session-RefreshTokenResponse) | リフレッシュトークンの生成（ログイン時） |
| RevokeRefreshToken | [RevokeRefreshTokenRequest](#session-RevokeRefreshTokenRequest) | [RevokeRefreshTokenResponse](#session-RevokeRefreshTokenResponse) | リフレッシュトークンの無効化（ログアウト時など） |
| ValidateRefreshToken | [ValidateRefreshTokenRequest](#session-ValidateRefreshTokenRequest) | [ValidateRefreshTokenResponse](#session-ValidateRefreshTokenResponse) | リフレッシュトークンの検証 |

 



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

