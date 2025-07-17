#!/bin/bash
go mod init "${MSROOTTMP}"/src/main
go build -o "${MSROOTTMP}"/build/ "${MSROOTTMP}"/src/main #gosetup