# Changelog

## v0.1.1

- Renames public package paths to lowercase `proviscore` and `provisentities`.
- Exports the provider type as `proviscore.ProvisProvider`.

## v0.1.0

Initial public release candidate.

- Adds HMAC-256 request signing for PROVIS API calls.
- Supports courses, quotas, workers, personal data, families, installations, and groups endpoints.
- Adds optional debug output for generated curl commands, response metadata, and response bodies.
- Provides typed request parameters and response models for the supported endpoints.
