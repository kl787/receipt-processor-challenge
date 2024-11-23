# How to deploy the application using Docker

## Build a docker Image
```
docker build -t receipt-processor-challenge
```

## Run the application in a Docker container
```
docker run -p 10000:10000 receipt-processor-challenge
```

# How to test the application using Docker with example
## Examples
```
{
  "retailer": "Target",
  "purchaseDate": "2022-01-01",
  "purchaseTime": "13:01",
  "items": [
    {
      "shortDescription": "Mountain Dew 12PK",
      "price": "6.49"
    },{
      "shortDescription": "Emils Cheese Pizza",
      "price": "12.25"
    },{
      "shortDescription": "Knorr Creamy Chicken",
      "price": "1.26"
    },{
      "shortDescription": "Doritos Nacho Cheese",
      "price": "3.35"
    },{
      "shortDescription": "   Klarbrunn 12-PK 12 FL OZ  ",
      "price": "12.00"
    }
  ],
  "total": "35.35"
}
```

## How to test the "Process Receipts"
```
curl 'http://localhost:10000/receipts/process -d '{json_data}'
```
### Result
```
{
    "id": "50d48221-4ab6-4cbb-bb92-f551c501260a"
}
```

## How to test the "Get Points"
```
curl 'http://localhost:10000/receipts/{id}/points
```
### Result
```
{
    "points": 109
}
```


