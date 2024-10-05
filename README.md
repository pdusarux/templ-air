# TEMPL-AIR 🚀

TEMPL-AIR is a blog-style website built with Go, Gin, and templ, featuring live reloading capabilities using Air. This project is inspired by and based on the Urchin CMS project by [@matheusgomes28](https://github.com/matheusgomes28/urchin).

## Features 🌟

- **Go-Powered Backend**: Utilizes the efficiency and performance of Go for server-side operations.
- **Gin Web Framework**: Employs Gin for routing and handling HTTP requests.
- **templ for HTML**: Uses templ for building HTML templates with Go.
- **Live Reloading**: Implements Air for hot-reloading during development.

## Getting Started 🏁

### Prerequisites

- Go (latest version recommended)
- [Air](https://github.com/cosmtrek/air) for live reloading
- [templ](https://github.com/a-h/templ) for HTML templating

### Installation

1. Clone the repository:

   ```
   git clone https://github.com/pdusarux/templ-air
   cd TEMPL-AIR
   ```

2. Install dependencies:

   ```
   go mod tidy
   ```

3. Install Air (if not already installed):

   ```
   go install github.com/cosmtrek/air@latest
   ```

4. Install templ (if not already installed):
   ```
   go install github.com/a-h/templ/cmd/templ@latest
   ```

### Running the Project

1. Start the server with live reloading:

   ```
   air
   ```

2. Open your browser and navigate to `http://localhost:8080` (or the port you've configured).

## SonarQube Integration 🔍

This project includes SonarQube for continuous code quality inspection. Follow these steps to set up and use SonarQube:

### Prerequisites

- Docker and Docker Compose installed on your machine
- SonarScanner for Go

### Setup

1. Start SonarQube using Docker Compose:

   ```
   docker-compose up -d
   ```

2. Access SonarQube at `http://localhost:9000`

   - Default credentials:
     - Username: admin
     - Password: admin
   - You'll be prompted to change the password on first login

3. Create a new project in SonarQube and generate a token

4. Install SonarScanner for Go:

   ```
   go install github.com/SonarSource/sonar-scanner-cli-binary/sonar-scanner@latest
   ```

5. Set up environment variable for SonarQube token:
   ```
   export SONAR_TOKEN=your_generated_token
   ```

### Configuration

The SonarQube configuration is defined in `sonar-project.properties`:

```properties
sonar.projectKey=TEMPL-AIR
sonar.sources=.
sonar.host.url=http://localhost:9000
sonar.login=${SONAR_TOKEN}
```

### Running Analysis

To run a SonarQube analysis:

1. Ensure SonarQube is running (step 1 in Setup)
2. Run the following command in your project root:

   ```
   sonar-scanner
   ```

3. View the results in the SonarQube dashboard at `http://localhost:9000`

## Contributing 🤝

Contributions are welcome! Please feel free to submit a Pull Request.

## Acknowledgements 🙏

This project is inspired by and based on the [Urchin CMS](https://github.com/matheusgomes28/urchin) by [@matheusgomes28](https://github.com/matheusgomes28).

## Contact 📬

Ployrung Dusarux - [pdusarux@gmail.com]

Project Link: [https://github.com/pdusarux/templ-air](https://github.com/pdusarux/templ-air)
