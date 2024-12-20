#!/bin/bash

source vars.sh

DESCRICAO="Monitor de comandos"
DESTINO=/etc/systemd/system/
SERVICO="${PROGRAMA}.service"
DESTINOSERVICO="${DESTINO}/${SERVICO}"


servico=$(cat << EOF
[Unit]
Description=$DESCRICAO
After=network.target

[Service]
Type=simple
User=$USER
WorkingDirectory=$DIRETORIOBIN
ExecStart=${DIRETORIOBIN}/${PROGRAMA}
Restart=always

[Install]
WantedBy=multi-user.target
EOF
)

# sudo ufw delete allow 8080
case "$1" in
    "-h" | "--help")
        echo "Script para registrar a API de '$DESCRICAO' como um serviço do Systemd"
        echo "$0 [-h | --help | -u | -r ]"
        echo -e "\t\t-h --help Exibe essa ajuda"
        echo -e "\t\t-u -r     Para e desativa o serviço '$SERVICO'"
        ;;
    "-u" | "-r")
        sudo systemctl stop "$SERVICO" &&
            sudo systemctl disable "$SERVICO" &&
            sudo ufw delete allow $PORT
        ;;
    *)
        if [ -f "${DIRETORIOBIN}/${PROGRAMA}" ]; then
            sudo ufw allow $PORT

            echo -e "\nCorpo do serviço '$SERVICO':"
            echo -e "$servico" | sudo tee "$DESTINOSERVICO" &&
                echo -e "\nSalvo serviço '$SERVICO' em '$DESTINO'\n" &&
                sudo systemctl enable "$SERVICO"

            sudo systemctl start "$SERVICO"
            sudo systemctl --no-pager -l status "$SERVICO"
        else
            echo "Execute o script 'build.sh' para compilar a API no diretório 'bin'. Mais detalhes no README.md"
        fi
        ;;
esac