# syntax=docker/dockerfile:1

ARG IMAGE=python:3.11-alpine3.17

FROM $IMAGE AS development

WORKDIR /app

# Copy project files
# NOTE: For more efficient builds, copy only
# required files here.
# For example, for a Python project, you may only
# need pyproject.toml and the actual *.py files
# that make up your project
COPY . .

RUN pip install .

EXPOSE 8000

CMD ["uvicorn", "{{ app }}", "--host", "0.0.0.0"]

FROM $IMAGE AS production

WORKDIR /app

# Copy project files
# NOTE: For more efficient builds, copy only
# required files here.
# For example, for a Python project, you may only
# need pyproject.toml and the actual *.py files
# that make up your project
COPY . .

RUN pip install .

EXPOSE 8000

ENTRYPOINT ["uvicorn", "{{ app }}", "--host", "0.0.0.0"]
