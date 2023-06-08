.PHONY: dev prod

# Objetivo por defecto
all: dev

# Objetivo para el modo de desarrollo
dev:
    @echo "====> Ejecutando en modo de desarrollo <===="
    @if [ -f api/docker/Dockerfile.dev ] && [ "$(shell cat api/docker/Dockerfile | head -n 1)" != "# Entorno: Desarrollo" ]; then \
        mv api/docker/Dockerfile api/docker/Dockerfile.prod; \
        mv api/docker/Dockerfile.dev api/docker/Dockerfile; \
    fi
    @if [ -f subscriber/docker/Dockerfile.dev ] && [ "$(shell cat subscriber/docker/Dockerfile | head -n 1)" != "# Entorno: Desarrollo" ]; then \
        mv subscriber/docker/Dockerfile subscriber/docker/Dockerfile.prod; \
        mv subscriber/docker/Dockerfile.dev subscriber/docker/Dockerfile; \
    fi
    @if [ -f acl/docker/Dockerfile.dev ] && [ "$(shell cat acl/docker/Dockerfile | head -n 1)" != "# Entorno: Desarrollo" ]; then \
        mv acl/docker/Dockerfile acl/docker/Dockerfile.prod; \
        mv acl/docker/Dockerfile.dev acl/docker/Dockerfile; \
    fi
    docker-compose up

# Objetivo para el modo de despliegue
prod:
    @echo "====> Ejecutando en modo de despliegue <===="
    @if [ -f api/docker/Dockerfile.prod ] && [ "$(shell cat api/docker/Dockerfile | head -n 1)" != "# Entorno: Producción" ]; then \
        mv api/docker/Dockerfile api/docker/Dockerfile.dev; \
        mv api/docker/Dockerfile.prod api/docker/Dockerfile; \
    fi
    @if [ -f subscriber/docker/Dockerfile.prod ] && [ "$(shell cat subscriber/docker/Dockerfile | head -n 1)" != "# Entorno: Producción" ]; then \
        mv subscriber/docker/Dockerfile subscriber/docker/Dockerfile.dev; \
        mv subscriber/docker/Dockerfile.prod subscriber/docker/Dockerfile; \
    fi
    @if [ -f acl/docker/Dockerfile.prod ] && [ "$(shell cat acl/docker/Dockerfile | head -n 1)" != "# Entorno: Producción" ]; then \
        mv acl/docker/Dockerfile acl/docker/Dockerfile.dev; \
        mv acl/docker/Dockerfile.prod acl/docker/Dockerfile; \
    fi
    docker-compose up
