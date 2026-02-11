# 🧩 Cuelang Composability: The "Slot" Architecture

**Status:** Draft | **Version:** 1.0

Signet replaces rigid YAML templates with **Composable Cuelang Schemas**. This allows for "Strong Governance" (Mandatory Slots) mixed with "Great DX" (Flexible Patches).

## 1. The Slot Concept
A Pipeline is defined as a set of **Slots**. Some are mandatory (Core), some are optional (User).

```cue
// schema/pipeline.cue (The Law)
#Pipeline: {
    meta: {
        name: string
        team: string
    }

    // Mandatory Slots (Enforced by Signet)
    steps: {
        init:    #Step & { name: "Signet Integrity Check" }
        build:   #Step // User defines this, but must output an artifact
        scan:    #Step & { name: "Signet Overseer Scan" }
        seal:    #Step & { name: "Signet Ledger Write" }
    }
}
```

2. Unification Strategy
We use Cuelang's unification (&) to merge the Global Policy with the User Patch.

The Global Policy (Locked)
Located in .signet/core/policy.cue. Managed by Platform Team.

```
#GlobalPolicy: {
    steps: {
        // Enforce that scanning always happens before sealing
        scan: {
            required: true
            timeout: "10m"
        }
    }
}
```
The User Patch (Flexible)
Located in .signet/app/patch.cue. Managed by Developers.

```
#UserPatch: {
    meta: {
        name: "Payment Service"
        team: "Checkout"
    }

    steps: {
        build: {
            command: "go build -o main ."
            env: { CGO_ENABLED: "0" }
        }
    }
}
```

The Runtime Merger
The Signet Engine executes:

```go
finalConfig := Unify(GlobalPolicy, UserPatch)
if err != nil {
    // Fails if UserPatch tries to override a locked field
    return Error("Policy Violation: Cannot override scan step")
}
```
3. The Catalog (Composable Bricks)
We provide a library of "Verified Partials" that developers can import to fill their slots.

catalog.golang.build

catalog.terraform.plan

catalog.npm.test

Usage in User Patch:

```
import "github.com/company/signet/catalog"

#UserPatch: {
    steps: {
        // Developer injects a pre-made, compliant brick
        build: catalog.golang.build & {
            version: "1.23"
        }
    }
}
```

4. Integrity Protection
To prevent developers from editing the Core Policy:

CODEOWNERS: Locks .signet/core/ and .github/workflows/.

Runtime Hash: The Engine calculates the hash of the loaded Cuelang files. If the Core files do not match the known "Gold Standard" hash from the Registry, the Engine aborts.
