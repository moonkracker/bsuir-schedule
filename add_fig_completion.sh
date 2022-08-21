#!/bin/bash

if ! command -v bsuir-schedule &> /dev/null
then
    echo "bsuir-schdeule could not be found"
    exit
fi

if command -v fig &> /dev/null
then
    bsuir-schedule generate-fig-spec > ~/.fig/autocomplete/src/bsuir-schedule.ts
    npx @fig/publish-spec --spec-path ~/.fig/autocomplete/src/bsuir-schedule.ts
    exit
else 
    echo "fig not found"
    exit
fi