# SpaceTrouble Booking System

## Implementation details

- Destinations, launchpads, and the weekly launch schedule are defined in the `db/init.sql` script, which runs when the PostgreSQL Docker container starts.

- `BookingService.CreateBooking` validates booking details by checking the existence of the destination and launchpad IDs and ensuring a flight is scheduled for the specified launch date. It also verifies that the booking does not conflict with any upcoming SpaceX launches. If all validations pass, a new booking is created with a unique UUID and saved to the repository.

- `LaunchClient` fetches and caches SpaceX upcoming launch data. It includes an HTTP client and an in-memory cache to reduce API calls and improve performance. The GetUpcomingLaunches method checks if cached data is valid; if not, it makes an HTTP request to fetch the data, decodes the JSON response, and updates the cache.

## Run locally

```sh
docker-compose up --build
```

## Stop and Clean Up

```sh
docker-compose down -v
```


## Testing

### Create Booking
#### 201 Created test

```bash
curl --location 'localhost:8080/bookings' \
--header 'Content-Type: application/json' \
--data '{
    "first_name": "John",
    "last_name": "Doe",
    "gender": "Male",
    "birthday": "1990-01-01",
    "launchpad_id": "5e9e4502f509092b78566f87",
    "destination_id": "1443a911-b39c-4404-bdab-dcffe4a6c019",
    "launch_date": "2024-11-02"
}'
```

#### 400 BadRequest samples
##### Destination ID not found
```bash
curl --location 'localhost:8080/bookings' \
--header 'Content-Type: application/json' \
--data '{
    "first_name": "John",
    "last_name": "Doe",
    "gender": "Male",
    "birthday": "1990-01-01",
    "launchpad_id": "5e9e4502f509092b78566f87",
    "destination_id": "1443a911-b39c-4404-bdab-dcffe4a6c01",
    "launch_date": "2024-11-02"
}'
```

Response:
```json
{
    "message": "booking creation failed: not found: destination with ID 1443a911-b39c-4404-bdab-dcffe4a6c01"
}
```
##### Last name not provided
```bash
curl --location 'localhost:8080/bookings' \
--header 'Content-Type: application/json' \
--data '{
    "first_name": "John",
    "last_name": "",
    "birthday": "1990-01-01",
    "launchpad_id": "5e9e4502f509092b78566f87",
    "destination_id": "1443a911-b39c-4404-bdab-dcffe4a6c019",
    "launch_date": "2024-11-02"
}'
```
```json
{
    "message": "last_name is required"
}
```

#### 409 Conflict with SpaceX
```bash
curl --location 'localhost:8080/bookings' \
--header 'Content-Type: application/json' \
--data '{
    "first_name": "John",
    "last_name": "Doe",
    "birthday": "1990-01-01",
    "launchpad_id": "5e9e4502f509094188566f88",
    "destination_id": "69aea949-8fba-4059-a9bc-a5de0d9f9b59",
    "launch_date": "2022-11-01"
}'
```

```json
{
    "message": "booking not possible: conflict"
}
```

### Get All Bookings
```sh
curl --location 'localhost:8080/bookings' \
--data ''
```
### Delete Booking

```bash
curl --location --request DELETE 'localhost:8080/bookings/4fcf7117-c40c-4812-b27e-dbd90841ebbf' \
--data ''
```