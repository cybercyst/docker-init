# docker-new

You want to get started with containerization immediately.
You've heard the benefits, but after poking around in multiple blog posts and documents, you're not sure what to add to get going.
You'd like to quickly start with containerization.
You'd like to try out container based development environments.

**Make good best guesses to start off new Docker users with the most popular languages / frameworks**

## Usage

```bash

# generate a project with whatever tooling / etc you already have

$ cd my-project

$ docker-new
Checking for project type...
We've detected a Go project 🎓

CREATE Dockerfile
CREATE .dockerignore
CREATE docker-compose.yaml

✅ Finished generating files

🚀 Run docker compose up to get started! 🚀

$ docker compose up

```

## Support

Currently supports

- Golang (using go.mod)
- Python (using pyproject.toml)
- Angular
- React

## Demo

[![asciicast](https://asciinema.org/a/KW5lOX439PdMP4qVYm02AgHP3.svg)](https://asciinema.org/a/KW5lOX439PdMP4qVYm02AgHP3)
