# URL Shortener API Testing Documentation

## Introduction

This documentation explains how to test the URL Shortener API using Postman. This API allows you to create, manage, and access short URLs.

## Base URL

- Production: `http://domain.com`
- Development: `http://localhost:8000`

## Endpoints to Test

### 1. Create a New Short URL

**Endpoint:** `POST /api/v1`

**Request Body:**

```json
{
    "url": "https://your-long-url",
    "short": "",
    "expiry": 45
}
```

**Expected Response:**

```json
{
    "url": "https://your-long-url",
    "short": "domain.com/url-id",
    "expiry": 45,
    "rate_limit": 9,
    "rate_limit_reset": 1800
}
```

### 2. Access Original URL

**Endpoint:** `GET /api/v1/:shortID`

**Expected Response:**

```json
{
    "url": "https://your-long-url"
}
```

### 3. Add Tag to Short URL

**Endpoint:** `POST /api/v1/tag`

**Request Body:**

```json
{
    "shortID": "short-id",
    "tag": "name-tag"
}
```

**Expected Response:**

```json
{
    "data": "https://your-long-url",
    "tags": [
        "name-tag"
    ]
}
```

### 4. Update Short URL

**Endpoint:** `PUT /api/v1/:shortID`

**Request Body:**

```json
{
    "url": "https://update-your-long-url",
    "expiry": 45
}
```

**Expected Response:**

```json
{
    "message": "URL updated successfully"
}
```

### 5. Delete Short URL

**Endpoint:** `DELETE /api/v1/:shortID`

**Expected Response:**

```json
{
    "message": "URL deleted successfully"
}
```

## Testing Steps with Postman

1. **Import Collection**
   - Open Postman
   - Click "Import" and select "Raw Text"
   - Paste the following code:

```json
{
  "info": {
    "name": "URL Shortener API Testing",
    "description": "Collection for testing URL Shortener API",
    "schema": "https://schema.getpostman.com/json/collection/v2.1.0/collection.json"
  },
  "item": [
    {
      "name": "Create Short URL",
      "request": {
        "method": "POST",
        "header": [
          {
            "key": "Content-Type",
            "value": "application/json"
          }
        ],
        "body": {
          "mode": "raw",
          "raw": "{\n    \"url\": \"https://your-long-url\",\n    \"short\": \"\",\n    \"expiry\": 45\n}"
        },
        "url": {
          "raw": "{{base_url}}/api/v1",
          "host": ["{{base_url}}"],
          "path": ["api", "v1"]
        }
      }
    },
    {
      "name": "Get Original URL",
      "request": {
        "method": "GET",
        "header": [],
        "url": {
          "raw": "{{base_url}}/api/v1/:shortID",
          "host": ["{{base_url}}"],
          "path": ["api", "v1", ":shortID"],
          "variable": [
            {
              "key": "shortID",
              "value": "your-short-id-here"
            }
          ]
        }
      }
    },
    {
      "name": "Add Tag to Short URL",
      "request": {
        "method": "POST",
        "header": [
          {
            "key": "Content-Type",
            "value": "application/json"
          }
        ],
        "body": {
          "mode": "raw",
          "raw": "{\n    \"shortID\": \"short-id\",\n    \"tag\": \"name-tag\"\n}"
        },
        "url": {
          "raw": "{{base_url}}/api/v1/tag",
          "host": ["{{base_url}}"],
          "path": ["api", "v1", "tag"]
        }
      }
    },
    {
      "name": "Update Short URL",
      "request": {
        "method": "PUT",
        "header": [
          {
            "key": "Content-Type",
            "value": "application/json"
          }
        ],
        "body": {
          "mode": "raw",
          "raw": "{\n    \"url\": \"https://update-your-long-url\",\n    \"expiry\": 45\n}"
        },
        "url": {
          "raw": "{{base_url}}/api/v1/:shortID",
          "host": ["{{base_url}}"],
          "path": ["api", "v1", ":shortID"],
          "variable": [
            {
              "key": "shortID",
              "value": "your-short-id-here"
            }
          ]
        }
      }
    },
    {
      "name": "Delete Short URL",
      "request": {
        "method": "DELETE",
        "header": [],
        "url": {
          "raw": "{{base_url}}/api/v1/:shortID",
          "host": ["{{base_url}}"],
          "path": ["api", "v1", ":shortID"],
          "variable": [
            {
              "key": "shortID",
              "value": "your-short-id-here"
            }
          ]
        }
      }
    }
  ],
  "variable": [
    {
      "key": "base_url",
      "value": "localhost:8000",
      "type": "string"
    }
  ]
}
```

2. **Setup Environment Variable**

   - Create a new environment variable in Postman
   - Add the `base_url` variable with the value:
     - For local testing: `localhost:8000`
     - For production: `domain.com`
3. **Run Tests**

   - Select the appropriate environment
   - Execute requests in the following order:
     1. Create Short URL (save the generated short ID)
     2. Get Original URL (use the short ID from step 1)
     3. Add Tag to Short URL (use the short ID from step 1)
     4. Update Short URL (use the short ID from step 1)
     5. Delete Short URL (use the short ID from step 1)

## Testing Notes

- Ensure the server is running before testing
- For local testing, make sure the PostgreSQL database is connected
- Save the short ID generated from the CREATE endpoint for use in other endpoints
- Pay attention to received response codes:
  - 200: Success
  - 400: Bad Request
  - 404: Not Found
  - 500: Internal Server Error

## Troubleshooting

1. If you get a connection refused error, make sure the server is running on the correct port
2. If you get a database error, verify the database connection is successful
3. If the response doesn't match expectations, check the JSON format being sent

## Version Information

- API Version: v1
- Document Version: 1.0
