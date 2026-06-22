# Changelog

## v0.1.4

- Adds `PersonImage` to fetch person images through the SDK without storing them.
- Adds `AccessByDate` to fetch access events by date/time range.
- Documents the recommended backend proxy and short-lived frontend cache plan for images.
- Lists `PersonImage` and `AccessByDate` in the public provider endpoints.

## v0.1.1

- Renames public package paths to lowercase `proviscore` and `provisentities`.
- Exports the provider type as `proviscore.ProvisProvider`.

## v0.1.0

Initial public release candidate.

- Adds HMAC-256 request signing for PROVIS API calls.
- Supports courses, quotas, workers, personal data, families, installations, and groups endpoints.
- Adds optional debug output for generated curl commands, response metadata, and response bodies.
- Provides typed request parameters and response models for the supported endpoints.
