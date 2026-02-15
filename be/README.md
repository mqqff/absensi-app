# Go Fiber GORM Boilerplate

This is a boilerplate project built with **Go**, using the **Fiber** web framework and **GORM** as the ORM. It is designed with modern architectural principles to provide a solid, scalable, and maintainable foundation for your web applications.

[![MIT License](https://img.shields.io/badge/license-MIT-blue.svg)](https://opensource.org/licenses/MIT)
[![Go Version](https://img.shields.io/badge/go-1.24%2B-brightgreen.svg)](https://golang.org/dl/)
[![Taskfile](https://img.shields.io/badge/Task-enabled-brightgreen?logo=data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAEAAAABACAYAAACqaXHeAAAEeElEQVR4Xu2bfYjVVRDHP+e+V8xpjplmZmZM2Z9EGaVtaEVUq120s6IoCOJCEgRRUBBEFBERUfADUfCjKAIKooKgIAgK4g9BFAUVu0gLoUNaJmlb2jRly2Rm5sycuTNz7v1x7t29b95995x7s1/u7t3ZmfM7M/OdO8+5M0eq/0Sq/l8p+qgKqgLdYFmwbLAEWAQsA9YCy4Gfge3AG8AnwIeBG4GXgTvAm8A4sA5YJcAPwA/A0cDFwJ3AicAzwA3gG+BnYJp/voDXgF/AycC1wDHAUeAb4CvgI+B/YDpgjYAdwDPAEcAtwDXgCeAR4FvgWeBnYGpAXsA7wFHA7cARwFHgo8B+4Fvgf+BZ4FNgdEAcwJvAscBtwCvAscAjwE7gZ8BvwJPAz4HhAHkCrwE/ADcAtwOXAZ8CvwC/A38BbwA/A8MCA/wK/At8BdwBHAacAjwM/Ae8Dfwe+A4YDghwB/Ba8BtwC3Ac8BrwLPAG8CTwNfA9MDwgwK/An8BvwE3AEcBpwEPAf8ALwGvA68BcwzD/SwHXgY+BJ4HtgYEA/gCeAR4BvgE+A14DngceAw8C3wA3jQzIF3gF2AVcAlwHPAU8AjwPPA08D3wAzDQvwI/AzsBO4CbgYeCbYBlwCvAU8BXwA/Aj8DwwPCPAr8CmwG7gLuA94GngU+BB4AXgQ+BT4g3EBGBAo8B/wEvB68BwwPCBANeA14C3gVeAFYIz/PgK2A4uAsYA1wKLAUmAt8C+wDlgdWH1i+g0sBZYDy4EVwAqgC1gNLAE2AisAi4GFwFbASmARsBC4GVgCrAQWAssBO4GFgOXA5sBBYDZwKbAZOARsBrYCNgJbAfsA64GdgdXA5sBWYBewBdgG7AWeAs8C84F5wLngPWAYsBg4DlwGjAPHgbHAMuAsMBs4BswGzgFzgeHAGGA0cByYCxwHxgHjgUXAmGAxcAwYDlwDlgGjgdHACGA0cCIYDgwFRgOjgbHAUGAksBoYBowChgHDgCHAGGAUGAksBYYDZf41YDwgC/gLeBZYDTwFhgFLgY/Ar8BfwDvA48AwwIBfgF+A24BPgS+BT4DPgA+A34Dfge8Cw/wL/A18DfwYGAfwLfA88CzwGPAUsBv4CjgTeA24GlgQ4FfgC+B7YBNgYEA/gM+BJ4BngUeAXYCfwFvAj8B3wPBAgH8B/wIfAw8De4Bfge+AR4GPgQ+B74Dh/gX+A94BngE2AssBfwJvA38APwKbA3UDfge2A8sBi4C1gBXAduAN4DbgfmA3sBvYDZwD7AZ2AvsB+4DdgE7gC2ARcCaYBFwJbAesA9YDVgM7AI+BvYDdwBbgS3AAuBJYBkwFxgO7ABGAauA7cBNwI3A5sA+YD/gP2BvsB3YDPwFPAfsBm4FngTmA/OBecG8YD4wD1gEjAPmAhOBsUBiMBgYDQxH9gKjgbHAUGAYMBwYDoz/G9yPjADGACMAI6CjwCVgmv8uYDXwM/AZ8CmwDlgLrH3iG+AL4GvgD2ABcBL4CbgLeBl4BHgR+AN4FPgMeA14CfgZ+BGYFojhL4BvA48CHwO/Ao8CPwOvA98CfwM/AtMC8QR+CgwWf0Twf2S6gD/n3zJ+A/k3yV8AAAAASUVORK5CYII=)](https://taskfile.dev)

## 🚀 Key Features

-   **Structured Logging**: Implemented structured logging for easier debugging and application monitoring.
-   **Custom Error & Response Handling**: Consistent custom error and response handlers across the entire application.
-   **JWT Authentication**: Secure JWT-based authentication middleware.
-   **Docker Support**: Fully containerized with `docker-compose` for an isolated and consistent development environment.
-   **Taskfile for Automation**: Uses [Taskfile](https://taskfile.dev/) as a command runner to simplify development tasks.

---

## 🛠️ Tech Stack

This boilerplate is built with the following modern technologies:

-   **[Fiber](https://gofiber.io/)**: An expressive web framework inspired by Express.js, built on top of Fasthttp, the fastest HTTP engine for Go.
-   **[GORM](https://gorm.io/)**: A fantastic Object Relational Mapper (ORM) for Go, simplifying database interactions.
-   **[JWT](https://github.com/golang-jwt/jwt)**: Implementation of JSON Web Tokens for secure, stateless authentication.
-   **[Docker](https://www.docker.com/)**: A containerization platform to build, ship, and run applications in containers.

---

## 🏛️ Architecture

The project is designed following software engineering best practices to ensure scalability, maintainability, and separation of concerns.

-   **Clean Architecture**: Separates code into independent layers (Presentation, Business, Data) to reduce dependencies and improve flexibility.
-   **Modern Design Principles**: Adopts various design principles to build a robust system, including:
	-   **Layered Architecture**
	-   **Domain-Driven Design (DDD)**
	-   **Hexagonal Architecture (Ports and Adapters)**
	-   **SOLID Principles**

---

## 🏁 Getting Started

To get started with this project, follow these steps:

1.  **Clone the repository:**
		```bash
		git clone [https://github.com/mqqff/absensi-app.git](https://github.com/mqqff/absensi-app.git)
		cd go-boilerplate
		```

2.  **Install dependencies:**
		```bash
		go mod tidy
		```

3.  **Run with Docker Compose:**
		```bash
		docker-compose up --build
		```

The application will be running at `http://localhost:3000`.

---

## 🤖 Taskfile Commands

This project uses `Taskfile` to automate common development workflows. First, [install Task](https://taskfile.dev/installation/).

Here are some of the available commands:

-   **Run the development server with hot-reloading:**
		```bash
		task dev
		```

-   **Build the application binary:**
		```bash
		task build
		```

-   **Start all services using Docker Compose:**
		```bash
		task service:up
		```

-   **Stop all services:**
		```bash
		task service:down
		```

Check the `Taskfile.yml` for a full list of available commands.

---

## 📜 License

This project is licensed under the [MIT License](LICENSE).
