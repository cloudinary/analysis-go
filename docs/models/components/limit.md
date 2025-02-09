# Limit


## Fields

| Field                                                    | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `Type`                                                   | [components.Feature](../../models/components/feature.md) | :heavy_check_mark:                                       | N/A                                                      |
| `UsedByRequest`                                          | **int64*                                                 | :heavy_minus_sign:                                       | The amount of quota used by the current request          |
| `Remaining`                                              | **int64*                                                 | :heavy_minus_sign:                                       | The remaining quota                                      |
| `Limit`                                                  | **int64*                                                 | :heavy_minus_sign:                                       | The total quota limit                                    |
| `ResetTime`                                              | [*time.Time](https://pkg.go.dev/time#Time)               | :heavy_minus_sign:                                       | Time in UTC when the limit will be reset                 |