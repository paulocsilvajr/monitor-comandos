#!/bin/bash

echo "CTRL + c para FINALIZAR"
sudo journalctl -u monitor-comandos.service -f
