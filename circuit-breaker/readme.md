# Go Circuit Breaker with Redis Pub/Sub

This project demonstrates a simple **circuit breaker** implementation using **Go** and **Redis Pub/Sub**. The `post service` calls the `profile service` (mock!), and the circuit breaker ensures that if the `profile service` is down, the `post service` will not make further requests until the profile service is available again.

## Architecture

- **Profile Service**:
  - Publishes its status (`up` or `down`) to a Redis Pub/Sub channel (`profile_service_status`).
  - Simulates being "up" or "down" at regular intervals.
  
- **Post Service**:
  - Subscribes to the Redis channel to receive status updates from the `profile service`.
  - If the `profile service` is "down", the circuit is tripped and no calls are made to the profile service.
  - Once the `profile service` is "up" again, the circuit is closed, and calls resume.

## How It Works

1. **Profile Service** periodically publishes its current status (`up` or `down`) to a Redis channel.
2. **Post Service** listens to this status via Redis Pub/Sub.
3. If the profile service is "down", the circuit breaker in the post service is tripped, preventing calls to the profile service.
4. When the profile service becomes available again, the post service resumes making calls.

## Dependencies

- Go 1.18+
- Docker

## Setup

1. **Install Go Dependencies**:
   ```bash
   go mod tidy
   ```

2. **Start Redis** (if not running already):
   ```bash
    docker run -d -p 6379:6379 redis
   ```

3. **Run Profile Service**:
   ```bash
   go run profile_service.go
   ```

4. **Run Post Service**:
   ```bash
   go run post_service.go
   ```
