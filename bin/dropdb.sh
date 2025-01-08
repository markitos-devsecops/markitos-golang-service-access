#!/bin/bash
# =============================================
#  __  __           _    _  __      
# |  \/  |         | |  (_)/ _|     
# | \  / | ___  ___| | ___| |_ ___  
# | |\/| |/ _ \/ __| |/ / |  _/ _ \ 
# | |  | |  __/ (__|   <| | ||  __/ 
# |_|  |_|\___|\___|_|\_\_|_| \___| 
#                                   
#  Creador: Marco Antonio - markitos      
# =============================================
#  🥷 (Cultura DevSecOps) 🗡️
#  Mejor seguro que nunca. 
# =============================================

# Configuración estricta
set -euo pipefail
IFS=$'\n\t'

# Variables globales (modifica según tus necesidades)
SCRIPT_NAME=$(basename "$0")
LOG_FILE="/tmp/${SCRIPT_NAME%.sh}.log"

# Funciones básicas
function log_info() {
    echo "[INFO] $*" | tee -a "$LOG_FILE"
}

function log_error() {
    echo "[ERROR] $*" >&2 | tee -a "$LOG_FILE"
}

cat <<'EOT'

┌────────────────────────────────────────────┐
│                                            │
│ Oh, Vamos allá, autodestruccion en 3.2..1. │
│                                            │
│        Marco Antonio_mArKit0s 2025.        │
│                                            │
└────────────────────────────────────────────┘

EOT

#:[.'.]:>-------------------------------------
#:[.'.]:> Tu lógica aquí
#:[.'.]:>-------------------------------------
log_info "¡Bienvenido al script más burlón de la Cultura DevSecOps!"
docker exec markitos-golang-service-postgres psql -U admin -c "SELECT pg_terminate_backend(pg_stat_activity.pid) FROM pg_stat_activity WHERE pg_stat_activity.datname = 'markitos-golang-service-access' AND pid <> pg_backend_pid();"
docker exec markitos-golang-service-postgres dropdb -U admin markitos-golang-service-access
docker exec markitos-golang-service-postgres psql -U admin -c "DROP USER IF EXISTS \"markitos-golang-service-access\";"
#:[.'.]:>-------------------------------------