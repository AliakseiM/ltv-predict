# ltv-predict

## Utility for predicting 60s day revenue.

Supports input data in CSV or JSON formats.

Supported models:
- linear regression - lr
- exponential smoothing (Holt's linear trend) - es

Grouping is available by country or campaign

## Example usage
```shell
docker build -t ltv-predict .

docker run ltv-predict -a "country" -m "lr" -s "json"
```
