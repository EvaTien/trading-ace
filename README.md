# trading-ace

## Description

A campaign for **Uniswap's users**. Target pool address is `0xB4e16d0168e52d35CaCD2c6185b44281Ec28C9Dc`

## Features

1. Get user tasks status by address
2. Get user points history for distributed tasks by address

## Requirements

docker, docker-compose

## Installation

1. Clone the repository:
    ```bash
    git clone https://github.com/EvaTien/trading-ace.git
    cd trading-ace
    ```

2. Run the server:
    ```bash
    docker compose up --build
    ```

## Usage

- **GET /users/{address}** – Get user by address
   - Example Request: `GET http://localhost:8080/users/12345`
   - Example Response:
    ```json
    {
      "address": "12345",
      "onboarding_completed": false,
      "total_amount": 0,
      "total_points": 0,
      "weekly_amount": 0
   }
    ```

- **GET /users/{address}/points-history** – Get user points history by address
   - Example Request: `GET http://localhost:8080/users/12345/points-history`
   - Example Response:
    ```json
   [
      {
         "address": "12345",
         "shared_points": 100,
         "total_points": 100,
         "week_end": "2024-11-27T00:00:00Z",
         "week_start": "2024-11-20T00:00:00Z"
      }
   ]
    ```
