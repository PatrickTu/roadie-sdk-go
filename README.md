# roadie-sdk-go

## Overview
This project serves as a wrapper around [Roadie's API](https://docs.roadie.com). The client supports `/v1/*` endpoints.

### Benefits
- Maintain compatibility with clients as the specification evolves over time.
- Simplify usage as end users don't have to worry about the underlying implementation.
## Features
- Supports all authentication methods listed at https://docs.roadie.com/#authentication-methods
- Retrieving an Estimate
- Retrieving Shipment(s)
- Updating an existing Shipment
- Cancelling an existing Shipment
## Future Improvements
### SDK

- Validation for request structs (https://github.com/go-playground/validator)
- Add support for driver interactions (tip, feedback)
- Barcode scanning
- Ability to retrieve artifacts (signature image, pickup image, delivery image, barcode label)
- Custom error handlers to inform end users of expected behavior (rate limiting)
- In-depth testing (invalid body, missing authentication, unexpected responses)
- Release first version

### API

1. Add `Event` as part of the existing data types. Clicking hyperlinks with the following url does not lead to the definition. https://docs.roadie.com/#event
   
2. Discrepancies in the sample [response](https://docs.roadie.com/#retrieve-a-list-of-shipments) when retrieving multiple shipments

    `alternate_id_1` and `alternate_id_2` are marshalled as a string and number

    #### String
    ```json
        "alternate_id_1": "111",
        "alternate_id_2": "222",
    ```

    #### Number
    ```json
        "alternate_id_1": 333,
        "alternate_id_2": 444,
    ```


## Usage

### Installing package

*Assuming the package has been published*

``` go get roadie.com/gopkg/roadie-sdk-go```

### Setting up the client

#### Provide authentication method to access API

API Key

```go
authTransport := NewAuthAPIKeyTransport("key")
```

Bearer Token

```go
authTransport := NewAuthBearerTokenTransport("token")
```

 Basic Auth

```go
authTransport := NewAuthBasicTransport("username", "password")
```

#### Instantiate Client

Pass in the root url to the API and a `*http.Client` with the auth transport.

```go 
client := NewRoadieClient("url", &http.Client{Transport: authTransport})
```

### Available Methods

You can call the following methods directly from the `RoadieClient`

#### Estimate

```go
type EstimateService interface {
	CreateEstimate(CreateEstimateRequest) (Estimate, error)
}
```

#### Shipment

```go
type ShipmentService interface {
	CreateShipment(CreateShipmentRequest) (Shipment, error)
	RetrieveShipment(id int) (Shipment, error)
	RetrieveShipments(ids []int, referenceIds []string) ([]Shipment, error)
	UpdateShipment(id int, request UpdateShipmentRequest) (Shipment, error)
	CancelShipment(id int, request CancelShipmentRequest) error
}
```

### Testing

To run the integration tests, ensure you have [Mockoon](https://mockoon.com/) installed. The responses are based on the examples given on the documentation.

1. Open Mockoon
2. Import `mockoon.json`
3. Start the mock server
4. Run tests