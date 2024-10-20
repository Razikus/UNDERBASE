
# UNDERBASE

UNDERBASE is a full-stack web application that demonstrates how to integrate Supabase Auth with a custom Go backend and a Quasar Framework frontend (just for test).

If Supabase is too much for you, and you don't need anything else than authorization - this is for you!

## Features

- User registration and login using Supabase Auth
- Custom Go backend with Echo framework 
- Quasar Framework frontend
- Docker-based deployment
- Single Sign-On (SSO) capability

## Prerequisites

- Docker and Docker Compose
- Go 1.22 or later
- Node.js 14 or later
- Yarn package manager

## Setup and Installation

1. Clone the repository:
   ```
   git clone https://github.com/razikus/underbase.git
   cd underbase
   ```

2. Build and run the Docker containers:
   ```
   docker-compose up --build
   ```

3. The application should now be running at `http://localhost`

## Usage

1. Navigate to `http://localhost` in your web browser
2. Use the interface to register a new user or log in
3. Once logged in, you can test the authenticated endpoint

## Development

This is just a Proof of Concept. 