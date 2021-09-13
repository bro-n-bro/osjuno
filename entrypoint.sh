#!/bin/sh
set -e

osjuno_init() {
  flags="--cosmos-prefix osmo"
  if [[ ! -z "${OSJUNO_CLIENT_NAME}" ]]; then
    flags=" ${flags} --client-name ${OSJUNO_CLIENT_NAME}"
  fi
  if [[ ! -z "${OSJUNO_COSMOS_MODULE}" ]]
  then
    flags=" ${flags} --cosmos-modules ${OSJUNO_COSMOS_MODULE}"
  else
    flags=" ${flags} --cosmos-modules auth,osmosis/gamm"
  fi
  if [[ ! -z "${OSJUNO_DATABASE_NAME}" ]]; then
    flags=" ${flags} --database-name ${OSJUNO_DATABASE_NAME}"
  fi
  if [[ ! -z "${OSJUNO_DATABASE_HOST}" ]]; then
    flags=" ${flags} --database-host ${OSJUNO_DATABASE_HOST}"
  fi
  if [[ ! -z "${OSJUNO_DATABASE_PASSWORD}" ]]; then
    flags=" ${flags} --database-password ${OSJUNO_DATABASE_PASSWORD}"
  fi
  if [[ ! -z "${OSJUNO_DATABASE_PORT}" ]]; then
    flags=" ${flags} --database-port ${OSJUNO_DATABASE_PORT}"
  fi
  if [[ ! -z "${OSJUNO_DATABASE_SCHEMA}" ]]; then
    flags=" ${flags} --database-schema ${OSJUNO_DATABASE_SCHEMA}"
  fi
  if [[ ! -z "${OSJUNO_DATABASE_SSL_MODE}" ]]; then
    flags=" ${flags} --database-ssl-mode ${OSJUNO_DATABASE_SSL_MODE}"
  fi
  if [[ ! -z "${OSJUNO_DATABASE_USER}" ]]; then
    flags=" ${flags} --database-user ${OSJUNO_DATABASE_USER}"
  fi
  if [[ ! -z "${OSJUNO_GRPC_ADDRESS}" ]]; then
    flags=" ${flags} --grpc-address ${OSJUNO_GRPC_ADDRESS}"
  fi
  if [[ ! -z "${OSJUNO_GRPC_INSECURE}" ]]; then
    flags=" ${flags} --grpc-insecure ${OSJUNO_GRPC_INSECURE}"
  fi
  if [[ ! -z "${OSJUNO_LOGGING_FORMAT}" ]]; then
    flags=" ${flags} --logging-format ${OSJUNO_LOGGING_FORMAT}"
  fi
  if [[ ! -z "${OSJUNO_LOGGING_LEVEL}" ]]; then
    flags=" ${flags} --logging-level ${OSJUNO_LOGGING_LEVEL}"
  fi
  if [[ ! -z "${OSJUNO_MAX_IDLE_CONNECTIONS}" ]]; then
    flags=" ${flags} --max-idle-connections ${OSJUNO_MAX_IDLE_CONNECTIONS}"
  fi
  if [[ ! -z "${OSJUNO_MAX_OPEN_CONNECTIONS}" ]]; then
    flags=" ${flags} --max-open-connections  ${OSJUNO_MAX_OPEN_CONNECTIONS}"
  fi
  if [[ ! -z "${OSJUNO_PARSING_FAST_SYNC}" ]]; then
    flags=" ${flags} --parsing-fast-sync  ${OSJUNO_PARSING_FAST_SYNC}"
  fi
  if [[ ! -z "${OSJUNO_PARSING_GENESIS_FILE_PATH}" ]]; then
    flags=" ${flags} --parsing-genesis-file-path ${OSJUNO_PARSING_GENESIS_FILE_PATH}"
  fi
  if [[ ! -z "${OSJUNO_PARSING_NEW_BLOCKS}" ]]; then
    flags=" ${flags} --parsing-new-blocks  ${OSJUNO_PARSING_NEW_BLOCKS}"
  fi
  if [[ ! -z "${OSJUNO_PARSING_OLD_BLOCKS}" ]]; then
    flags=" ${flags} --parsing-old-blocks ${OSJUNO_PARSING_OLD_BLOCKS}"
  fi
  if [[ ! -z "${OSJUNO_PARSING__PARSE_GENESIS}" ]]; then
    flags=" ${flags} --parsing-parse-genesis ${OSJUNO_PARSING__PARSE_GENESIS}"
  fi
  if [[ ! -z "${OSJUNO_PARSING_START_HEIGHT}" ]]; then
    flags=" ${flags} --parsing-start-height ${OSJUNO_PARSING_START_HEIGHT}"
  fi
  if [[ ! -z "${OSJUNO_PARSING_WORKERS}" ]]; then
    flags=" ${flags} --parsing-workers ${OSJUNO_PARSING_WORKERS}"
  fi
  if [[ ! -z "${OSJUNO_PRUNING_INTERVAL}" ]]; then
    flags=" ${flags} --pruning-interval ${OSJUNO_PRUNING_INTERVAL}"
  fi
  if [[ ! -z "${OSJUNO_PRUNING_KEEP_EVERY}" ]]; then
    flags=" ${flags} --pruning-keep-every ${OSJUNO_PRUNING_KEEP_EVERY}"
  fi
  if [[ ! -z "${OSJUNO_PRUNING_KEEP_RECENT}" ]]; then
    flags=" ${flags} --pruning-keep-recent ${OSJUNO_PRUNING_KEEP_RECENT}"
  fi
  if [[ ! -z "${OSJUNO_RPC_ADDRESS}" ]]; then
    flags=" ${flags} --rpc-address ${OSJUNO_RPC_ADDRESS}"
  fi
  if [[ ! -z "${OSJUNO_TELEMETRY_ENABLED}" ]]; then
    flags=" ${flags} --telemetry-enabled ${OSJUNO_TELEMETRY_ENABLED}"
  fi
  if [[ ! -z "${OSJUNO_TELEMETRY_PORT}" ]]; then
    flags=" ${flags} --telemetry-port ${OSJUNO_TELEMETRY_PORT}"
  fi

  osjuno init --home ${OSJUNO_HOME} ${flags}
}

# Default CMD
if [ "$1" = 'osjuno' ] && [[ -z "$2" ]]; then
  OSJUNO_HOME=${OSJUNO_HOME:-/osjuno/.osjuno}
  if [ ! -f "${OSJUNO_HOME}/config.toml" ]; then
    osjuno_init
  fi
  osjuno parse --home ${OSJUNO_HOME}
else
  # This allow user to use other commands
  exec "$@"
fi
