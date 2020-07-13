#!/bin/bash

base_dir='./migrations'
migration_main=''

update_migration_main () {
    # shellcheck disable=SC2116
    migration_main="$(echo $base_dir)/"
}

init_migration () {
    go run "$migration_main" init
}

reset_go_migrations () {
    go run "$migration_main" reset
}

run_go_migrations () {
    go run "$migration_main" up
}

rerun_go_migrations () {
    reset_go_migrations
    run_go_migrations
}

main () {
    source ~/.bash_profile

    update_migration_main

    while [ -n "$1" ]
    do
        case "$1" in
            --reset) reset_go_migrations ;;
            --rerun) rerun_go_migrations ;;
            --run) run_go_migrations ;;
            --init) init_migration ;;
        esac
        shift;
    done
}

main "$@"