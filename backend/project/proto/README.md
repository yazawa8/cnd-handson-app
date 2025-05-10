# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [project.proto](#project-proto)
    - [CreateProjectRequest](#project-CreateProjectRequest)
    - [DeleteProjectRequest](#project-DeleteProjectRequest)
    - [DeleteProjectResponse](#project-DeleteProjectResponse)
    - [GetProjectRequest](#project-GetProjectRequest)
    - [ListProjectsRequest](#project-ListProjectsRequest)
    - [ListProjectsResponse](#project-ListProjectsResponse)
    - [Project](#project-Project)
    - [ProjectResponse](#project-ProjectResponse)
    - [UpdateProjectRequest](#project-UpdateProjectRequest)
  
    - [ProjectService](#project-ProjectService)
  
- [Scalar Value Types](#scalar-value-types)



<a name="project-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## project.proto



<a name="project-CreateProjectRequest"></a>

### CreateProjectRequest
--- CreateProject ---


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |
| description | [string](#string) |  |  |
| owner_id | [string](#string) |  |  |






<a name="project-DeleteProjectRequest"></a>

### DeleteProjectRequest
--- DeleteProject ---


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="project-DeleteProjectResponse"></a>

### DeleteProjectResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| success | [bool](#bool) |  |  |






<a name="project-GetProjectRequest"></a>

### GetProjectRequest
--- GetProject ---


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="project-ListProjectsRequest"></a>

### ListProjectsRequest
--- ListProjects ---


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| owner_id | [string](#string) |  | 特定のユーザーが所有するプロジェクトに絞り込むためのフィールド |






<a name="project-ListProjectsResponse"></a>

### ListProjectsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| projects | [Project](#project-Project) | repeated |  |






<a name="project-Project"></a>

### Project
--- Project ---


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| name | [string](#string) |  |  |
| description | [string](#string) |  |  |
| owner_id | [string](#string) |  |  |
| created_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |
| updated_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |






<a name="project-ProjectResponse"></a>

### ProjectResponse
--- ProjectResponse ---


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| project | [Project](#project-Project) |  |  |






<a name="project-UpdateProjectRequest"></a>

### UpdateProjectRequest
--- UpdateProject ---


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| name | [string](#string) |  |  |
| description | [string](#string) |  |  |





 

 

 


<a name="project-ProjectService"></a>

### ProjectService
================================
サービス定義
================================

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| CreateProject | [CreateProjectRequest](#project-CreateProjectRequest) | [ProjectResponse](#project-ProjectResponse) | プロジェクトの作成 |
| UpdateProject | [UpdateProjectRequest](#project-UpdateProjectRequest) | [ProjectResponse](#project-ProjectResponse) | プロジェクトの更新 |
| GetProject | [GetProjectRequest](#project-GetProjectRequest) | [ProjectResponse](#project-ProjectResponse) | プロジェクトの取得 |
| ListProjects | [ListProjectsRequest](#project-ListProjectsRequest) | [ListProjectsResponse](#project-ListProjectsResponse) | プロジェクト一覧の取得 |
| DeleteProject | [DeleteProjectRequest](#project-DeleteProjectRequest) | [DeleteProjectResponse](#project-DeleteProjectResponse) | プロジェクトの削除 |

 



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

