
SERVICE=$1

if [[ -z "$SERVICE" ]]; then
  docker compose down -v --remove-orphans
  docker compose rm -f -s
  docker compose up --always-recreate-deps --remove-orphans --renew-anon-volumes --build
else
  docker compose stop "$SERVICE"
  docker compose rm "$SERVICE" -f -s
  docker compose up "$SERVICE" --always-recreate-deps --remove-orphans --renew-anon-volumes --build
fi
