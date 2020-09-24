# jwt.proto
## Version: version not set

### /health

#### GET
##### Summary:

Gets health status - it should check database connection etc ..

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | A successful response. | [jwtHealthResponse](#jwthealthresponse) |

### /v2/decode/{jwt}

#### GET
##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| jwt | path |  | Yes | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | A successful response. | [jwtJwtResponse](#jwtjwtresponse) |

### /v2/encode

#### GET
##### Summary:

Generates jwt

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| exp | query |  | No | integer |
| short | query |  | No | boolean (boolean) |
| Message | query |  | No | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | A successful response. | [jwtJwtResponse](#jwtjwtresponse) |

### Models


#### jwtHealthResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| message | string |  | No |

#### jwtJwtResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| error | boolean (boolean) |  | No |
| status | integer |  | No |
| data | object |  | No |
| Message | string |  | No |

#### protobufNullValue

`NullValue` is a singleton enumeration to represent the null value for the
`Value` type union.

 The JSON representation for `NullValue` is JSON `null`.

 - NULL_VALUE: Null value.

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| protobufNullValue | string | `NullValue` is a singleton enumeration to represent the null value for the `Value` type union.   The JSON representation for `NullValue` is JSON `null`.   - NULL_VALUE: Null value. |  |