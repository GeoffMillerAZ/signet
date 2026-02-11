# configs/schema/groups.cue

#Groups: {
    persistence: [...string] | *["db/**/*.sql", "models/*.go"]
    networking:  [...string] | *["api/**/*.go", "proto/*.proto"]
    logic:       [...string] | *["internal/core/**/*.go"]
}

#Policy: {
    version: string
    groups: #Groups
    enforcement: {
        require_sigterm: bool | *true
        require_migration: bool | *true
    }
}
