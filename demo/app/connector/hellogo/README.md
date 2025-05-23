# ndc_go Connector

## Get started

Read the documentation at https://hasura.io/docs/3.0/getting-started/build/add-business-logic?db=Go

## Configuration

See the help of [serve](https://github.com/hasura/ndc-sdk-go#using-this-sdk) command of the connector SDK.

Besides that the connector supports extra environments:

| Name                         | Description                                                                                                                                             | Default Value |
| ---------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------- |
| `QUERY_CONCURRENCY`          | The limit of concurrent query executions by name if there are many request variables in a single request wit h format `<key1>=<value1>;<key2>=<value2>` |               |
| `QUERY_CONCURRENCY_LIMIT`    | The limit of concurrent query executions if there are many request variables in a single request                                                        | `1`           |
| `MUTATION_CONCURRENCY_LIMIT` | The limit of concurrent mutation executions if there are many operations in a single request                                                            | `1`           |
