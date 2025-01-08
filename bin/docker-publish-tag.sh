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
echo "Building image markitos-golang-service-access:$1" && exit
echo "" && \
echo "" && \
echo "Building image markitos-golang-service-access:$1" && \
docker build -t ghcr.io/markitos-devsecops/markitos-golang-service-access:$1 . && \
echo "" && \
echo "Pushing image markitos-golang-service-access:$1" && \
docker push ghcr.io/markitos-devsecops/markitos-golang-service-access:$1 && \
echo "" && \
echo "Image markitos-golang-service-access:$1 pushed" && \
docker image rm --force ghcr.io/markitos-devsecops/markitos-golang-service-access:$1
echo "" && \
echo "Image markitos-golang-service-access:$1 removed"
echo "" && \
echo "Done!"
#:[.'.]:>-------------------------------------