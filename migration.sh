#!/bin/bash

base_dir='migrations'
migration_main=''

update_migration_main () {
    # shellcheck disable=SC2116
    migration_main="$(echo $base_dir)/main.go"
}

init_migration () {
    go run "$migration_main" init
}

reset_go_migrations () {
    go run "$migration_main" reset
}

run_go_migrations () {
    init_migration
    go run "$migration_main" up
}

main () {
    source ~/.bash_profile

    update_migration_main

    while [ -n "$1" ]
    do
        case "$1" in
            --reset-migration) reset_go_migrations ;;
            --run-migration) run_go_migrations ;;
            --init-migration) init_migration ;;
        esac
        shift;
    done
}

main "$@"