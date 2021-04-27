# Re-direct to remote environment.
export DOCKER_HOST="ssh://ubuntu@raspberrypi"

# Run your docker-compose commands.
docker-compose pull
docker-compose down
docker-compose up

# All docker-compose commands here will be run on remote-host.

# Switch back to your local environment.
unset DOCKER_HOST