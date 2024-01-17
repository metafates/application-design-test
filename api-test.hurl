GET http://localhost:1234/v1/health
HTTP 200

POST http://localhost:1234/v1/hotels/orders
Content-Type: application/json
{
  "user": {
    "email": "hi@example.com"
  },
  "booking": {
    "time_range": {
      "from": "2024-01-10T12:18:52.198483+03:00",
      "to": "2024-01-11T12:18:52.198483+03:00"
    },
    "room": {
      "type": "luxury"
    }
  }
}

HTTP 201

POST http://localhost:1234/v1/hotels/orders
Content-Type: application/json
{
  "user": {
    "email": "hi@example.com"
  },
  "booking": {
    "time_range": {
      "from": "2024-01-10T12:18:52.198483+03:00",
      "to": "2024-01-11T12:18:52.198483+03:00"
    },
    "room": {
      "type": "luxury"
    }
  }
}

HTTP 409

GET http://localhost:1234/v1/hotels/bookings?email=hi@example.com
Accept: application/json

HTTP 200

[Asserts]
jsonpath "$[0].room.type" == "Luxury"
