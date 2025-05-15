# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [task.proto](#task-proto)
    - [CreateTaskRequest](#task-CreateTaskRequest)
    - [DeleteTaskRequest](#task-DeleteTaskRequest)
    - [DeleteTaskResponse](#task-DeleteTaskResponse)
    - [GetTaskRequest](#task-GetTaskRequest)
    - [ListTasksRequest](#task-ListTasksRequest)
    - [ListTasksResponse](#task-ListTasksResponse)
    - [Task](#task-Task)
    - [TaskResponse](#task-TaskResponse)
    - [UpdateTaskRequest](#task-UpdateTaskRequest)
  
    - [TaskService](#task-TaskService)
  
- [Scalar Value Types](#scalar-value-types)



<a name="task-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## task.proto



<a name="task-CreateTaskRequest"></a>

### CreateTaskRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| title | [string](#string) |  |  |
| description | [string](#string) |  |  |
| status | [string](#string) |  |  |
| start_time | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |
| end_time | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |
| column_id | [string](#string) |  |  |
| assignee_id | [string](#string) |  |  |






<a name="task-DeleteTaskRequest"></a>

### DeleteTaskRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="task-DeleteTaskResponse"></a>

### DeleteTaskResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| success | [bool](#bool) |  |  |






<a name="task-GetTaskRequest"></a>

### GetTaskRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="task-ListTasksRequest"></a>

### ListTasksRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| column_id | [string](#string) |  | カラムID |
| assignee_id | [string](#string) |  | 担当者ID |
| page | [int32](#int32) |  | ページ番号 |
| page_size | [int32](#int32) |  | ページサイズ |






<a name="task-ListTasksResponse"></a>

### ListTasksResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| tasks | [Task](#task-Task) | repeated |  |
| total_count | [int32](#int32) |  | タスクの総数 |






<a name="task-Task"></a>

### Task
--- Task Entity ---


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| title | [string](#string) |  |  |
| description | [string](#string) |  |  |
| status | [string](#string) |  |  |
| start_time | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |
| end_time | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |
| column_id | [string](#string) |  |  |
| assignee_id | [string](#string) |  |  |






<a name="task-TaskResponse"></a>

### TaskResponse
--- 共通レスポンス ---


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| task | [Task](#task-Task) |  |  |






<a name="task-UpdateTaskRequest"></a>

### UpdateTaskRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| task | [Task](#task-Task) |  | 更新内容 |
| update_mask | [google.protobuf.FieldMask](#google-protobuf-FieldMask) |  |  |





 

 

 


<a name="task-TaskService"></a>

### TaskService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| GetTask | [GetTaskRequest](#task-GetTaskRequest) | [TaskResponse](#task-TaskResponse) | タスクの取得 |
| ListTasks | [ListTasksRequest](#task-ListTasksRequest) | [ListTasksResponse](#task-ListTasksResponse) | タスクの一覧取得 |
| CreateTask | [CreateTaskRequest](#task-CreateTaskRequest) | [TaskResponse](#task-TaskResponse) | タスクの作成 |
| UpdateTask | [UpdateTaskRequest](#task-UpdateTaskRequest) | [TaskResponse](#task-TaskResponse) | タスクの更新 |
| DeleteTask | [DeleteTaskRequest](#task-DeleteTaskRequest) | [DeleteTaskResponse](#task-DeleteTaskResponse) | タスクの削除 |

 



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

