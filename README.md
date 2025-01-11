# 🚀 Markitos Golang Service Access

## 📖 Description
This project provides a Golang service for accessing various resources. It is designed to be efficient, scalable, and easy to use.

## 🛠️ Installation
To install the project, clone the repository and build the project using the following commands:

```bash
git clone https://github.com/yourusername/markitos-golang-service-access.git
cd markitos-golang-service-access
cp local.app.env app.env
make docker-up
make createdb
```

## 🚀 Usage
To run the service, use the following command:

```bash
make test
make run
```

You can configure the service by editing the `app.env` file from local.app.env or production.app.env.

**Note:** The `app.env` file is ignored in Git, so make sure to create or copy it from the provided templates.

## 🌐 Environment Configuration
The service can be configured using environment variables. Create a `app.env` or copy file (local,production) in the root directory of the project and add your configuration settings. For example:

```env
APP_BBDD_DSN=host=localhost user=admin password=admin dbname=markitos-golang-service-access sslmode=disable TimeZone=Europe/Madrid port=5432 sslmode=disable
APP_ADDRESS=0.0.0.0:3000
```

## 🗄️ Database Requirement
This service requires a PostgreSQL database to function. You can set up the database using Docker. The project includes a `docker-compose.yaml` file to help you get started quickly.

To start the database using Docker, run the following command:

```bash
make docker-up
```

This command will start the PostgreSQL database container as defined in the `docker-compose.yaml` file.

To stop the database container, run:

```bash
make docker-down
```

## 📜 Makefile
The project includes a `Makefile` for managing build and deployment tasks. Here are some useful commands:

- `make docker-up`: Start Docker containers (Database requirement and other services).
- `make docker-down`: Stop Docker containers (Database requirement and other services).
- `make run`: Run the project.
- `make test`: Run tests without verbose output or coverage.
- `make testv`: Run tests with verbose output and without coverage.
- `make testc`: Run tests without verbose output and with coverage.
- `make testcv`: Run tests with verbose output and coverage.
- `make createdb`: Create the database.
- `make dropdb`: Drop the database.
- `make appsec-sast`: Run static application security testing.
- `make appsec-secrets`: Check for secrets in the code.
- `make appsec`: Run all security checks.
- `make docker-login`: Log in to Docker.
- `make docker-publish-tag`: Publish Docker image with a tag.
- `make docker-publish-postgres`: Publish Docker image for PostgreSQL.

## 💡 Examples

### Using Make
To build and run the project using `make`, use the following commands:

```bash
make createdb
make test
make run
```

### Using Bash
To build and run the project using `bash`, use the following commands:

```bash
bash ./bin/createdb.sh
bash ./bin/test.sh
bash ./bin/run.sh
```

### Testing
To run tests using `make`:

```bash
make test       # Run tests without verbose output or coverage
make testv      # Run tests with verbose output and without coverage
make testc      # Run tests without verbose output and with coverage
make testcv     # Run tests with verbose output and coverage
```

To run tests using `bash`:

```bash
bash bin/test.sh       # Run tests without verbose output or coverage
bash bin/testv.sh      # Run tests with verbose output and without coverage
bash bin/testc.sh      # Run tests without verbose output and with coverage
bash bin/testcv.sh     # Run tests with verbose output and coverage
```

### Docker
To manage Docker containers using `make`:

```bash
make docker-up
make docker-down
```

To manage Docker containers using `bash`:

```bash
bash bin/docker-up.sh
bash bin/docker-down.sh
```

### Database
To manage the database using `make`:

```bash
make createdb
make dropdb
```

To manage the database using `bash`:

```bash
bash bin/createdb.sh
bash bin/dropdb.sh
```

### Security
To run security checks using `make`:

```bash
make appsec-sast
make appsec-secrets
make appsec
```

To run security checks using `bash`:

```bash
bash bin/appsec-sast.sh
bash bin/appsec-secrets.sh
bash bin/appsec.sh
```

### Docker Publishing
To publish Docker images using `make`:

```bash
make docker-login
make docker-publish-tag
make docker-publish-postgres
```

To publish Docker images using `bash`:

```bash
bash bin/docker-login.sh $(GITHUB_TOKEN)
bash bin/docker-publish-tag.sh $(or $(TAG),1.0.0)
bash bin/docker-publish-postgres.sh $(or $(TAG),1.0.0)
```

If `TAG` is not provided, the default version `1.0.0` will be used.

### Obtaining GITHUB_TOKEN
To obtain a `GITHUB_TOKEN`, follow these steps:

1. Go to your GitHub account settings.
2. Navigate to "Developer settings" > "Personal access tokens".
3. Click "Generate new token".
4. Select the scopes you need and generate the token.
5. Use this token as `GITHUB_TOKEN` in your commands.

## 🤝 Contributing
Contributions are welcome! Please fork the repository and submit a pull request.

## 📜 License
This project is licensed under the MIT License. See the `LICENSE` file for more details.

---

**Author**: Marco Antonio Rubio Lopez - mArkit0s  
**Contact**: Cultura DevSecOps - markitos.es.info@gmail.com
