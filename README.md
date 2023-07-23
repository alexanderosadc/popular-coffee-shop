# Popular Cofee Shop
The Popular Coffee Shop Membership Quota Service allows coffee shop members to check their coffee quotas based on their membership type. 
It provides an API endpoint that accepts queries for coffee purchases and returns whether the purchase is within the quota limits.

## How to Run the App

To run the Popular Coffee Shop Membership Quota Service, make sure you have Docker and Docker Compose installed on your machine. Then, follow these steps:

1. Clone the repository to your local machine.
2. Navigate to the root directory of the project containing the docker-compose.yml file.
3. Run the following command: ```docker-compose up```

## How to Use the API
**Endpoint**: http://localhost:8080/buycoffee

**Method**: GET

### Headers:
  
  **user-id**: The unique ID of the user making the request (mandatory).
  
  **membership-type**: The type of membership for the user making the request (mandatory).
  
  membership-type avialiable:
  1. basic
  2. americano_maniac
  3. cofeelover

### Query Parameters:

**cofeeType**: The type of coffee the user wants to buy (mandatory).

cofee-type avialiable:
  1. americano
  2. espresso
  3. capucinno

### Response Format
* If the request is successful and there is no quota violation, the service will respond with a status code of ```200 (OK)``` without any body.
* If the request is not allowed due to a quota violation, the service will respond with a status code of ```429 (Too Many Requests)```
  and the message ```Hours to wait: X.XX``` in the response body, where X.XX represents the remaining time in hours before the quota refreshes.
