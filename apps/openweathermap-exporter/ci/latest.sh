#!/usr/bin/env bash
version=$(curl -sX GET "https://api.github.com/repos/haraldkoch/openweathermap-exporter/releases/latest" | jq --raw-output '. | .tag_name' 2>/dev/null)
version="${version#*v}"
version="${version#*release-}"
printf "%s" "${version}"
