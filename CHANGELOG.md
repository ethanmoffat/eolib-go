# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

Versions before 3.1.0 were reconstructed from the GitHub release notes and the git history. The 2.x versions are
yanked: they were tagged without changing the module path, so Go never accepted them.

## [Unreleased]

### Fixed
- `EoWriter.AddByte`, `AddChar`, `AddShort`, `AddThree` and `AddInt` now return an error for negative values, as
  eolib-dotnet does. Previously a negative value was written as an incorrect encoded number.

## [3.2.0] - 2026-09-27

### Added
- Tests for the packets captured from the official client and server, from the
  [eo-captured-packets](https://github.com/ethanmoffat/eo-captured-packets) repository. Each packet is deserialized,
  its field values are checked against the captured properties via reflection, and it is serialized back to the
  captured bytes.
- golangci-lint checks, run by `make build` and in CI.

### Changed
- The v3 module moved from the `v3/` subdirectory to the repository root. The module path and import paths are
  unchanged (`github.com/ethanmoffat/eolib-go/v3`). Only local `replace` directives pointing at `v3/` need updating.
- Removed the v1 code (`github.com/ethanmoffat/eolib-go`) and the v1 Makefile targets. Version 1 remains available
  from its `v1.x.y` tags, but is no longer maintained.

### Updated
- Pulled in changes for eo-protocol, with impact to generated code:
    - [Change SpellTargetOther caster_direction -> target_type](https://github.com/cirras/eo-protocol/commit/0c0ea73c6a30833541ff4c53a4ebe87432ef359d)
    - [Change LoginMessageCode.Yes 2->250](https://github.com/cirras/eo-protocol/commit/aaba85ffd101fdb848e28b90878f685f2b3c5450)

### Fixed
- Generated code now uses the full import path for types referenced from a package without a known alias. Previously,
  server packets referencing `net/client` types did not compile.

## [3.1.0] - 2026-04-02

### Updated
- Pulled in changes for eo-protocol, with impact to generated code:
    - Update SkillLearn.stat_requirements to correct types
    - Correct handling of chunking in WELCOME_REPLY
    - Rename npc_index to behavior_id in QuestAccept client packet
    - Add Bard emote

## [3.0.1] - 2025-01-28

### Fixed
- CHEST_CLOSE packet special case (optional field followed by dummy) now correctly writes either the optional field OR
  the dummy value, instead of writing both. Deserialization will only read the optional field if there are enough
  bytes available.
- Updated the `go get` command in the README to use v3.

## [3.0.0] - 2024-10-15

This major release fixes package versioning issues intrinsic to the Go module system. All v2 versions of eolib-go
should be considered broken and unusable in the Go module system.

### Updated
- Pulled in latest eo-protocol updates as of 2024-10-15. No impact on generated code function.

## [3.0.0-rc4] - 2024-08-21

### Updated
- Pulled in latest eo-protocol updates as of 2024-08-21. Trade structures and packets now use arrays instead of
  individual player IDs and item lists.

### Fixed
- Calculate element size for array loops over structures. Allows readers to ignore trailing elements with incomplete
  data.

## [3.0.0-rc3] - 2024-07-30

### Updated
- Pulled in eo-protocol fixes for SPELL_REPLY and QUEST_REPORT.

### Fixed
- Use correct type override for enums during serialize/deserialize calls, if present.

## [3.0.0-rc2] - 2024-06-06

### Added
- Separate Go module for v3 (`github.com/ethanmoffat/eolib-go/v3`), and Makefile targets to build it.

## [3.0.0-rc1] - 2024-06-06

### Added
- v3 sub-package with copies of the entire repo directory structure. `go get github.com/ethanmoffat/eolib-go/v3` to
  use the latest version.

### Updated
- Top-level packages reverted to their v1 state.

## [2.1.0] - 2024-06-05 [YANKED]

Not a valid Go module version: the module path was not updated to `/v2`, so this version was never installable.
Use 3.0.0 or later.

### Added
- Break bytes that are not in a chunked section will return an error during code generation.
- Lengths that are not referenced by another instruction will return an error during code generation.
- `ByteSize` field has been added to all generated structs. Returns the size of data that was read during
  deserialization.
- Assert that fixed-length fields are the correct length during serialization.

### Updated
- Length instructions no longer generate a field in generated code. The length used for (de)serialization is implicit
  based on the size of the collection.

### Fixed
- Correct calculation of type sizes for deserialization purposes.
- Corrected call to `Remaining` when deserializing data. No longer used in loop condition.
- Generated code now properly moves between chunks for nested types (switch structs) where `chunked` is specified in
  the outer specification.

## [2.0.1] - 2024-04-30 [YANKED]

Not a valid Go module version: the module path was not updated to `/v2`, so this version was never installable.
Use 3.0.0 or later.

### Fixed
- Use Offset property for `length` fields that specify it. Fixes bug where map signs had incorrect Title length data
  during serialize/deserialize.

## [2.0.0] - 2024-04-11 [YANKED]

Not a valid Go module version: the module path was not updated to `/v2`, so this version was never installable.
Use 3.0.0 or later.

### Added
- Top-level package docs are now available.
- Server pub files are now generated.

### Updated
- Code generation now uses a proper codegen library ([github.com/dave/jennifer](https://github.com/dave/jennifer)).

### Fixed
- `NPC_AGREE` server packet - npc count is now the correct type (char).
- Arena packets have correct chunking delimiters.
- Added comments to various struct members.
- `JUKEBOX_MSG` - remove chunking.
- `GUILD_BUY` - fix typo.
- `GUILD_TAKE` - add missing guild tag.
- `AVATAR_ADMIN` - fix field order.
- `EFFECT_AGREE`/`EFFECT_PLAYER` - update packet data to be arrays.
- `CharacterStatsInfoLookup` - secondary stats field type fixed.

## [1.1.2] - 2023-11-07

### Updated
- Added session ID to bank packets (protocol update).

## [1.1.1] - 2023-08-08

### Fixed
- Add support for `optional` attribute in protocol XML (was previously ignored). Optional values will no longer be
  serialized/deserialized if they are not present in data.
- Fix deserialize bug for slices of primitive types. Data was incorrectly deserialized directly to the current index of
  the slice without first appending a new element to the slice, which would cause a panic.

## [1.1.0] - 2023-08-04

### Added
- EoReader methods: `Slice`, `SliceFromIndex`, and `SliceFromCurrent`.

### Updated
- eo-protocol bumped to version
  [7032f45](https://github.com/Cirras/eo-protocol/commit/7032f4585fadce9f8633934a66b797259a668c5d).
- **Breaking change**: Serialize/Deserialize now take a pointer to an `EoWriter`/`EoReader`.
- **Breaking change**: `GenerateXXX` methods for sequences now take a pointer to a `rand`.

### Fixed
- EoReader `Remaining` method could sometimes return a negative value in chunked reading mode.
- Sequence generation in InitSequence used the incorrect formula.

## [1.0.0] - 2023-07-25

Initial release.

### Added
- Generated code for the EO v0.0.28 protocol.
- Functions for:
    - Encoding/decoding numbers and strings
    - Reading/writing EO data
    - Managing sequence numbers
    - Verifying server hash
- Helper functions for converting a packet ID to a strongly-typed implementation of the `net.Packet` interface.

[Unreleased]: https://github.com/ethanmoffat/eolib-go/compare/v3.2.0...HEAD
[3.2.0]: https://github.com/ethanmoffat/eolib-go/releases/tag/v3.2.0
[3.1.0]: https://github.com/ethanmoffat/eolib-go/releases/tag/v3.1.0
[3.0.1]: https://github.com/ethanmoffat/eolib-go/releases/tag/v3.0.1
[3.0.0]: https://github.com/ethanmoffat/eolib-go/releases/tag/v3.0.0
[3.0.0-rc4]: https://github.com/ethanmoffat/eolib-go/releases/tag/v3.0.0-rc4
[3.0.0-rc3]: https://github.com/ethanmoffat/eolib-go/releases/tag/v3.0.0-rc3
[3.0.0-rc2]: https://github.com/ethanmoffat/eolib-go/releases/tag/v3.0.0-rc2
[3.0.0-rc1]: https://github.com/ethanmoffat/eolib-go/releases/tag/v3.0.0-rc1
[2.1.0]: https://github.com/ethanmoffat/eolib-go/releases/tag/v2.1.0
[2.0.1]: https://github.com/ethanmoffat/eolib-go/releases/tag/v2.0.1
[2.0.0]: https://github.com/ethanmoffat/eolib-go/releases/tag/v2.0.0
[1.1.2]: https://github.com/ethanmoffat/eolib-go/releases/tag/v1.1.2
[1.1.1]: https://github.com/ethanmoffat/eolib-go/releases/tag/v1.1.1
[1.1.0]: https://github.com/ethanmoffat/eolib-go/releases/tag/v1.1.0
[1.0.0]: https://github.com/ethanmoffat/eolib-go/releases/tag/v1.0.0
