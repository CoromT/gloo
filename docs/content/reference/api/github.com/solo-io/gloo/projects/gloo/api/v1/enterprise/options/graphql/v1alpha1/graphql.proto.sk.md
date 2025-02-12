
---
title: "graphql.proto"
weight: 5
---

<!-- Code generated by solo-kit. DO NOT EDIT. -->


### Package: `graphql.gloo.solo.io` 
#### Types:


- [PathSegment](#pathsegment)
- [ValueProvider](#valueprovider)
- [GraphQLArgExtraction](#graphqlargextraction)
- [GraphQLParentExtraction](#graphqlparentextraction)
- [TypedValueProvider](#typedvalueprovider)
- [Type](#type)
- [JsonValueList](#jsonvaluelist)
- [JsonValue](#jsonvalue)
- [JsonKeyValue](#jsonkeyvalue)
- [JsonNode](#jsonnode)
- [RequestTemplate](#requesttemplate)
- [RESTResolver](#restresolver)
- [QueryMatcher](#querymatcher)
- [FieldMatcher](#fieldmatcher)
- [Resolution](#resolution)
- [GraphQLSchema](#graphqlschema) **Top-Level Resource**
  



##### Source File: [github.com/solo-io/gloo/projects/gloo/api/v1/enterprise/options/graphql/v1alpha1/graphql.proto](https://github.com/solo-io/gloo/blob/master/projects/gloo/api/v1/enterprise/options/graphql/v1alpha1/graphql.proto)





---
### PathSegment

 
used to reference into json structures by key(s).

```yaml
"key": string
"index": int
"all": bool

```

| Field | Type | Description |
| ----- | ---- | ----------- | 
| `key` | `string` | Key is used to extract named values from a map. Only one of `key`, `index`, or `all` can be set. |
| `index` | `int` | Index is used to extract an element at a certain index from a list. Only one of `index`, `key`, or `all` can be set. |
| `all` | `bool` | all selects all from the current element at in the path. This is useful for extracting list arguments / object arguments. Only one of `all`, `key`, or `index` can be set. |




---
### ValueProvider

 
In the future we may add support for regex and subgroups

```yaml
"graphqlArg": .graphql.gloo.solo.io.ValueProvider.GraphQLArgExtraction
"typedProvider": .graphql.gloo.solo.io.ValueProvider.TypedValueProvider
"graphqlParent": .graphql.gloo.solo.io.ValueProvider.GraphQLParentExtraction
"providerTemplate": string

```

| Field | Type | Description |
| ----- | ---- | ----------- | 
| `graphqlArg` | [.graphql.gloo.solo.io.ValueProvider.GraphQLArgExtraction](../graphql.proto.sk/#graphqlargextraction) | type inferred from schema, no need to provide it. Only one of `graphqlArg`, `typedProvider`, or `graphqlParent` can be set. |
| `typedProvider` | [.graphql.gloo.solo.io.ValueProvider.TypedValueProvider](../graphql.proto.sk/#typedvalueprovider) |  Only one of `typedProvider`, `graphqlArg`, or `graphqlParent` can be set. |
| `graphqlParent` | [.graphql.gloo.solo.io.ValueProvider.GraphQLParentExtraction](../graphql.proto.sk/#graphqlparentextraction) | Fetch value from the graphql_parent of the current field. Only one of `graphqlParent`, `graphqlArg`, or `typedProvider` can be set. |
| `providerTemplate` | `string` | If non-empty, the template to interpolate the extracted value into. For example, if the string is /api/pets/{}/name and the extracted value 123 is the pet ID will then the extracted value is /api/pets/123/name Use {} as the interpolation character (even repeated) regardless of the type of the provided value. |




---
### GraphQLArgExtraction



```yaml
"argName": string
"path": []graphql.gloo.solo.io.PathSegment
"required": bool

```

| Field | Type | Description |
| ----- | ---- | ----------- | 
| `argName` | `string` | The argument name to fetch. The argument value fetched will have a type from the schema that we validate in envoy. If the name is invalid, returns the zero-value primitive or null. |
| `path` | [[]graphql.gloo.solo.io.PathSegment](../graphql.proto.sk/#pathsegment) | Optional: fetches the value in the argument selected at this key. If the key is invalid, returns the zero-value primitive or null. |
| `required` | `bool` | If this is set to true, then a schema error will be returned to user when the graphql arg is not found. Defaults to false, so schema error will not be returned to user when the graphql arg is not found. |




---
### GraphQLParentExtraction

 
Does not do type coercion, but instead if the type does not match the
expected primitive type we throw an error.
In the future we may add support for type coercion.

```yaml
"path": []graphql.gloo.solo.io.PathSegment

```

| Field | Type | Description |
| ----- | ---- | ----------- | 
| `path` | [[]graphql.gloo.solo.io.PathSegment](../graphql.proto.sk/#pathsegment) | Fetches the value in the graphql parent at this key. The value will always be accepted since the parent object is not strongly-typed. If the key is invalid, returns null. |




---
### TypedValueProvider



```yaml
"type": .graphql.gloo.solo.io.ValueProvider.TypedValueProvider.Type
"header": string
"value": string

```

| Field | Type | Description |
| ----- | ---- | ----------- | 
| `type` | [.graphql.gloo.solo.io.ValueProvider.TypedValueProvider.Type](../graphql.proto.sk/#type) | Type that the value will be coerced into. For example if the extracted value is "9", and type is INT, this value will be cast to an int type. |
| `header` | `string` | Fetches the request/response header's value. If not found, uses empty string. Only one of `header` or `value` can be set. |
| `value` | `string` | inline value, use as provided rather than extracting from another source. Only one of `value` or `header` can be set. |




---
### Type

 
Type that the value will be coerced into.
For example if the extracted value is "9", and type is INT,
this value will be cast to an int type.

| Name | Description |
| ----- | ----------- | 
| `STRING` |  |
| `INT` |  |
| `FLOAT` |  |
| `BOOLEAN` |  |




---
### JsonValueList



```yaml
"values": []graphql.gloo.solo.io.JsonValue

```

| Field | Type | Description |
| ----- | ---- | ----------- | 
| `values` | [[]graphql.gloo.solo.io.JsonValue](../graphql.proto.sk/#jsonvalue) |  |




---
### JsonValue



```yaml
"node": .graphql.gloo.solo.io.JsonNode
"valueProvider": .graphql.gloo.solo.io.ValueProvider
"list": .graphql.gloo.solo.io.JsonValueList

```

| Field | Type | Description |
| ----- | ---- | ----------- | 
| `node` | [.graphql.gloo.solo.io.JsonNode](../graphql.proto.sk/#jsonnode) |  Only one of `node`, `valueProvider`, or `list` can be set. |
| `valueProvider` | [.graphql.gloo.solo.io.ValueProvider](../graphql.proto.sk/#valueprovider) |  Only one of `valueProvider`, `node`, or `list` can be set. |
| `list` | [.graphql.gloo.solo.io.JsonValueList](../graphql.proto.sk/#jsonvaluelist) |  Only one of `list`, `node`, or `valueProvider` can be set. |




---
### JsonKeyValue



```yaml
"key": string
"value": .graphql.gloo.solo.io.JsonValue

```

| Field | Type | Description |
| ----- | ---- | ----------- | 
| `key` | `string` |  |
| `value` | [.graphql.gloo.solo.io.JsonValue](../graphql.proto.sk/#jsonvalue) |  |




---
### JsonNode

 
Represents a typed JSON structure

```yaml
"keyValues": []graphql.gloo.solo.io.JsonKeyValue

```

| Field | Type | Description |
| ----- | ---- | ----------- | 
| `keyValues` | [[]graphql.gloo.solo.io.JsonKeyValue](../graphql.proto.sk/#jsonkeyvalue) | if keys repeat, the latest one replaces any earlier values associated with that key. repeated list, rather than a map, to have ordering to allow for merge semantics within the data plane, for example: - gRPC input uses special empty string for input key to set entire body - gRPC wants to replace a certain field in parsed body from GraphQL arg. |




---
### RequestTemplate

 
Defines a configuration for generating outgoing requests for a resolver.

```yaml
"headers": map<string, .graphql.gloo.solo.io.ValueProvider>
"queryParams": map<string, .graphql.gloo.solo.io.ValueProvider>
"outgoingBody": .graphql.gloo.solo.io.JsonValue

```

| Field | Type | Description |
| ----- | ---- | ----------- | 
| `headers` | `map<string, .graphql.gloo.solo.io.ValueProvider>` | Use this attribute to set request headers to your REST service. It consists of a map of strings to value providers. The string key determines the name of the resulting header, the value provided will be the value. at least need ":method" and ":path". |
| `queryParams` | `map<string, .graphql.gloo.solo.io.ValueProvider>` | Use this attribute to set query parameters to your REST service. It consists of a map of strings to value providers. The string key determines the name of the query param, the provided value will be the value. This value is appended to any value set to the :path header in `headers`. Interpolation is done in envoy rather than the control plane to prevent escaped character issues. Additionally, we may be providing values not known until the request is being executed (e.g., graphql parent info). |
| `outgoingBody` | [.graphql.gloo.solo.io.JsonValue](../graphql.proto.sk/#jsonvalue) | Used to construct the outgoing body to the upstream from the graphql value providers. |




---
### RESTResolver

 
control-plane API

```yaml
"upstreamRef": .core.solo.io.ResourceRef
"requestTransform": .graphql.gloo.solo.io.RequestTemplate
"spanName": string

```

| Field | Type | Description |
| ----- | ---- | ----------- | 
| `upstreamRef` | [.core.solo.io.ResourceRef](../../../../../../../../../../solo-kit/api/v1/ref.proto.sk/#resourceref) |  |
| `requestTransform` | [.graphql.gloo.solo.io.RequestTemplate](../graphql.proto.sk/#requesttemplate) | configuration used to compose the outgoing request to a REST API. |
| `spanName` | `string` |  |




---
### QueryMatcher



```yaml
"fieldMatcher": .graphql.gloo.solo.io.QueryMatcher.FieldMatcher

```

| Field | Type | Description |
| ----- | ---- | ----------- | 
| `fieldMatcher` | [.graphql.gloo.solo.io.QueryMatcher.FieldMatcher](../graphql.proto.sk/#fieldmatcher) |  |




---
### FieldMatcher



```yaml
"type": string
"field": string

```

| Field | Type | Description |
| ----- | ---- | ----------- | 
| `type` | `string` | Object type. For example, Query. |
| `field` | `string` | Field with in the object. |




---
### Resolution

 
This is the resolver map for the schema.
For each Type.Field, we can define a resolver.
if a field does not have resolver, the default resolver will be used.
the default resolver takes the field with the same name from the parent, and uses that value
to resolve the field.
if a field with the same name does not exist in the parent, null will be used.

```yaml
"matcher": .graphql.gloo.solo.io.QueryMatcher
"restResolver": .graphql.gloo.solo.io.RESTResolver

```

| Field | Type | Description |
| ----- | ---- | ----------- | 
| `matcher` | [.graphql.gloo.solo.io.QueryMatcher](../graphql.proto.sk/#querymatcher) | Match an object type and field. |
| `restResolver` | [.graphql.gloo.solo.io.RESTResolver](../graphql.proto.sk/#restresolver) |  |




---
### GraphQLSchema

 
Enterprise-Only: THIS FEATURE IS IN TECH PREVIEW. APIs are versioned as alpha and subject to change.
User-facing CR config for resolving client requests to graphql schemas.
Routes that have this config will execute graphql queries, and will not make it to the router filter. i.e. this
filter will terminate the request for these routes.
Note: while users can provide this configuration manually, the eventual UX will
be to generate the Executable Schema CRs from other sources and just have users
configure the routes to point to these schema CRs.

```yaml
"namespacedStatuses": .core.solo.io.NamespacedStatuses
"metadata": .core.solo.io.Metadata
"schema": string
"enableIntrospection": bool
"resolutions": []graphql.gloo.solo.io.Resolution

```

| Field | Type | Description |
| ----- | ---- | ----------- | 
| `namespacedStatuses` | [.core.solo.io.NamespacedStatuses](../../../../../../../../../../solo-kit/api/v1/status.proto.sk/#namespacedstatuses) | NamespacedStatuses indicates the validation status of this resource. NamespacedStatuses is read-only by clients, and set by gloo during validation. |
| `metadata` | [.core.solo.io.Metadata](../../../../../../../../../../solo-kit/api/v1/metadata.proto.sk/#metadata) | Metadata contains the object metadata for this resource. |
| `schema` | `string` | Schema to use in string format. |
| `enableIntrospection` | `bool` | Do we enable introspection for the schema? general recommendation is to disable this for production and hence it defaults to false. |
| `resolutions` | [[]graphql.gloo.solo.io.Resolution](../graphql.proto.sk/#resolution) | The resolver map to use to resolve the schema. Omitted fields will use the default resolver, which looks for a field with that name in the parent's object, and errors if the field cannot be found. |





<!-- Start of HubSpot Embed Code -->
<script type="text/javascript" id="hs-script-loader" async defer src="//js.hs-scripts.com/5130874.js"></script>
<!-- End of HubSpot Embed Code -->
