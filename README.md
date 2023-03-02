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

 - Add `Event` as part of the existing data types. Clicking hyperlinks with the following url does not lead to the definition. https://docs.roadie.com/#event


## Usage

### Installing package

*Assuming the package has been published internally*

``` go get roadie.com/gopkg/roadie-sdk-go```

### Setting up the client

Authentication

### Available Methods

### Testing

To run test suite, ensure you have [Mockoon](https://mockoon.com/) installed.