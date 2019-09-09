package puppetclient

var rawDDL = "ewogICIkc2NoZW1hIjogImh0dHBzOi8vY2hvcmlhLmlvL3NjaGVtYXMvbWNvcnBjL2RkbC92MS9hZ2VudC5qc29uIiwKICAibWV0YWRhdGEiOiB7CiAgICAibmFtZSI6ICJwdXBwZXQiLAogICAgImRlc2NyaXB0aW9uIjogIk1hbmFnZXMgdGhlIExpZmUgQ3ljbGUgb2YgdGhlIFB1cHBldCBBZ2VudCIsCiAgICAiYXV0aG9yIjogIlIuSS5QaWVuYWFyIDxyaXBAZGV2Y28ubmV0PiIsCiAgICAibGljZW5zZSI6ICJBcGFjaGUtMi4wIiwKICAgICJ2ZXJzaW9uIjogIjIuMy4xIiwKICAgICJ1cmwiOiAiaHR0cHM6Ly9naXRodWIuY29tL2Nob3JpYS1wbHVnaW5zL3B1cHBldC1hZ2VudCIsCiAgICAidGltZW91dCI6IDIwCiAgfSwKICAiYWN0aW9ucyI6IFsKICAgIHsKICAgICAgImFjdGlvbiI6ICJkaXNhYmxlIiwKICAgICAgImlucHV0IjogewogICAgICAgICJtZXNzYWdlIjogewogICAgICAgICAgInByb21wdCI6ICJNZXNzYWdlIiwKICAgICAgICAgICJkZXNjcmlwdGlvbiI6ICJTdXBwbHkgYSByZWFzb24gZm9yIGRpc2FibGluZyB0aGUgUHVwcGV0IGFnZW50IiwKICAgICAgICAgICJ0eXBlIjogInN0cmluZyIsCiAgICAgICAgICAiZGVmYXVsdCI6IG51bGwsCiAgICAgICAgICAib3B0aW9uYWwiOiB0cnVlLAogICAgICAgICAgInZhbGlkYXRpb24iOiAic2hlbGxzYWZlIiwKICAgICAgICAgICJtYXhsZW5ndGgiOiAxMjAKICAgICAgICB9CiAgICAgIH0sCiAgICAgICJvdXRwdXQiOiB7CiAgICAgICAgInN0YXR1cyI6IHsKICAgICAgICAgICJkZXNjcmlwdGlvbiI6ICJTdGF0dXMiLAogICAgICAgICAgImRpc3BsYXlfYXMiOiAiU3RhdHVzIiwKICAgICAgICAgICJkZWZhdWx0IjogIiIKICAgICAgICB9LAogICAgICAgICJlbmFibGVkIjogewogICAgICAgICAgImRlc2NyaXB0aW9uIjogIklzIHRoZSBhZ2VudCBjdXJyZW50bHkgbG9ja2VkIiwKICAgICAgICAgICJkaXNwbGF5X2FzIjogIkVuYWJsZWQiLAogICAgICAgICAgImRlZmF1bHQiOiBudWxsCiAgICAgICAgfQogICAgICB9LAogICAgICAiZGlzcGxheSI6ICJmYWlsZWQiLAogICAgICAiZGVzY3JpcHRpb24iOiAiRGlzYWJsZSB0aGUgUHVwcGV0IGFnZW50IiwKICAgICAgImFnZ3JlZ2F0ZSI6IFsKICAgICAgICB7CiAgICAgICAgICAiZnVuY3Rpb24iOiAiYm9vbGVhbl9zdW1tYXJ5IiwKICAgICAgICAgICJhcmdzIjogWwogICAgICAgICAgICAiZW5hYmxlZCIsCiAgICAgICAgICAgIHsKICAgICAgICAgICAgICAidHJ1ZSI6ICJlbmFibGVkIiwKICAgICAgICAgICAgICAiZmFsc2UiOiAiZGlzYWJsZWQiCiAgICAgICAgICAgIH0KICAgICAgICAgIF0KICAgICAgICB9CiAgICAgIF0KICAgIH0sCiAgICB7CiAgICAgICJhY3Rpb24iOiAiZW5hYmxlIiwKICAgICAgImlucHV0IjogewogICAgICB9LAogICAgICAib3V0cHV0IjogewogICAgICAgICJzdGF0dXMiOiB7CiAgICAgICAgICAiZGVzY3JpcHRpb24iOiAiU3RhdHVzIiwKICAgICAgICAgICJkaXNwbGF5X2FzIjogIlN0YXR1cyIsCiAgICAgICAgICAiZGVmYXVsdCI6ICIiCiAgICAgICAgfSwKICAgICAgICAiZW5hYmxlZCI6IHsKICAgICAgICAgICJkZXNjcmlwdGlvbiI6ICJJcyB0aGUgYWdlbnQgY3VycmVudGx5IGxvY2tlZCIsCiAgICAgICAgICAiZGlzcGxheV9hcyI6ICJFbmFibGVkIiwKICAgICAgICAgICJkZWZhdWx0IjogbnVsbAogICAgICAgIH0KICAgICAgfSwKICAgICAgImRpc3BsYXkiOiAiZmFpbGVkIiwKICAgICAgImRlc2NyaXB0aW9uIjogIkVuYWJsZSB0aGUgUHVwcGV0IGFnZW50IiwKICAgICAgImFnZ3JlZ2F0ZSI6IFsKICAgICAgICB7CiAgICAgICAgICAiZnVuY3Rpb24iOiAiYm9vbGVhbl9zdW1tYXJ5IiwKICAgICAgICAgICJhcmdzIjogWwogICAgICAgICAgICAiZW5hYmxlZCIsCiAgICAgICAgICAgIHsKICAgICAgICAgICAgICAidHJ1ZSI6ICJlbmFibGVkIiwKICAgICAgICAgICAgICAiZmFsc2UiOiAiZGlzYWJsZWQiCiAgICAgICAgICAgIH0KICAgICAgICAgIF0KICAgICAgICB9CiAgICAgIF0KICAgIH0sCiAgICB7CiAgICAgICJhY3Rpb24iOiAibGFzdF9ydW5fc3VtbWFyeSIsCiAgICAgICJpbnB1dCI6IHsKICAgICAgICAibG9ncyI6IHsKICAgICAgICAgICJwcm9tcHQiOiAiUGFyc2UgbG9nIGZyb20gbGFzdF9ydW5fcmVwb3J0LnlhbWwiLAogICAgICAgICAgImRlc2NyaXB0aW9uIjogIldoZXRoZXIgb3Igbm90IHRvIHBhcnNlIHRoZSBsb2dzIGZyb20gbGFzdF9ydW5fcmVwb3J0LnlhbWwiLAogICAgICAgICAgInR5cGUiOiAiYm9vbGVhbiIsCiAgICAgICAgICAiZGVmYXVsdCI6IGZhbHNlLAogICAgICAgICAgIm9wdGlvbmFsIjogdHJ1ZQogICAgICAgIH0KICAgICAgfSwKICAgICAgIm91dHB1dCI6IHsKICAgICAgICAib3V0X29mX3N5bmNfcmVzb3VyY2VzIjogewogICAgICAgICAgImRlc2NyaXB0aW9uIjogIlJlc291cmNlcyB0aGF0IHdlcmUgbm90IGluIGRlc2lyZWQgc3RhdGUiLAogICAgICAgICAgImRpc3BsYXlfYXMiOiAiT3V0IG9mIFN5bmMgUmVzb3VyY2VzIiwKICAgICAgICAgICJkZWZhdWx0IjogLTEKICAgICAgICB9LAogICAgICAgICJmYWlsZWRfcmVzb3VyY2VzIjogewogICAgICAgICAgImRlc2NyaXB0aW9uIjogIlJlc291cmNlcyB0aGF0IGZhaWxlZCB0byBhcHBseSIsCiAgICAgICAgICAiZGlzcGxheV9hcyI6ICJGYWlsZWQgUmVzb3VyY2VzIiwKICAgICAgICAgICJkZWZhdWx0IjogLTEKICAgICAgICB9LAogICAgICAgICJjb3JyZWN0ZWRfcmVzb3VyY2VzIjogewogICAgICAgICAgImRlc2NyaXB0aW9uIjogIlJlc291cmNlcyB0aGF0IHdlcmUgY29ycmVjdGl2ZWx5IGNoYW5nZWQiLAogICAgICAgICAgImRpc3BsYXlfYXMiOiAiQ29ycmVjdGVkIFJlc291cmNlcyIsCiAgICAgICAgICAiZGVmYXVsdCI6IC0xCiAgICAgICAgfSwKICAgICAgICAiY2hhbmdlZF9yZXNvdXJjZXMiOiB7CiAgICAgICAgICAiZGVzY3JpcHRpb24iOiAiUmVzb3VyY2VzIHRoYXQgd2VyZSBjaGFuZ2VkIiwKICAgICAgICAgICJkaXNwbGF5X2FzIjogIkNoYW5nZWQgUmVzb3VyY2VzIiwKICAgICAgICAgICJkZWZhdWx0IjogLTEKICAgICAgICB9LAogICAgICAgICJ0b3RhbF9yZXNvdXJjZXMiOiB7CiAgICAgICAgICAiZGVzY3JpcHRpb24iOiAiVG90YWwgcmVzb3VyY2VzIG1hbmFnZWQgb24gYSBub2RlIiwKICAgICAgICAgICJkaXNwbGF5X2FzIjogIlRvdGFsIFJlc291cmNlcyIsCiAgICAgICAgICAiZGVmYXVsdCI6IDAKICAgICAgICB9LAogICAgICAgICJjb25maWdfcmV0cmlldmFsX3RpbWUiOiB7CiAgICAgICAgICAiZGVzY3JpcHRpb24iOiAiVGltZSB0YWtlbiB0byByZXRyaWV2ZSB0aGUgY2F0YWxvZyBmcm9tIHRoZSBtYXN0ZXIiLAogICAgICAgICAgImRpc3BsYXlfYXMiOiAiQ29uZmlnIFJldHJpZXZhbCBUaW1lIiwKICAgICAgICAgICJkZWZhdWx0IjogLTEKICAgICAgICB9LAogICAgICAgICJ0b3RhbF90aW1lIjogewogICAgICAgICAgImRlc2NyaXB0aW9uIjogIlRvdGFsIHRpbWUgdGFrZW4gdG8gcmV0cmlldmUgYW5kIHByb2Nlc3MgdGhlIGNhdGFsb2ciLAogICAgICAgICAgImRpc3BsYXlfYXMiOiAiVG90YWwgVGltZSIsCiAgICAgICAgICAiZGVmYXVsdCI6IDAKICAgICAgICB9LAogICAgICAgICJsb2dzIjogewogICAgICAgICAgImRlc2NyaXB0aW9uIjogIkxvZyBsaW5lcyBmcm9tIHRoZSBsYXN0IFB1cHBldCBydW4iLAogICAgICAgICAgImRpc3BsYXlfYXMiOiAiTGFzdCBSdW4gTG9ncyIsCiAgICAgICAgICAiZGVmYXVsdCI6IHsKICAgICAgICAgIH0KICAgICAgICB9LAogICAgICAgICJsYXN0cnVuIjogewogICAgICAgICAgImRlc2NyaXB0aW9uIjogIldoZW4gdGhlIEFnZW50IGxhc3QgYXBwbGllZCBhIGNhdGFsb2cgaW4gbG9jYWwgdGltZSIsCiAgICAgICAgICAiZGlzcGxheV9hcyI6ICJMYXN0IFJ1biIsCiAgICAgICAgICAiZGVmYXVsdCI6IDAKICAgICAgICB9LAogICAgICAgICJzaW5jZV9sYXN0cnVuIjogewogICAgICAgICAgImRlc2NyaXB0aW9uIjogIkhvdyBsb25nIGFnbyBkaWQgdGhlIEFnZW50IGxhc3QgYXBwbHkgYSBjYXRhbG9nIGluIGxvY2FsIHRpbWUiLAogICAgICAgICAgImRpc3BsYXlfYXMiOiAiU2luY2UgTGFzdCBSdW4iLAogICAgICAgICAgImRlZmF1bHQiOiAiVW5rbm93biIKICAgICAgICB9LAogICAgICAgICJjb25maWdfdmVyc2lvbiI6IHsKICAgICAgICAgICJkZXNjcmlwdGlvbiI6ICJQdXBwZXQgY29uZmlnIHZlcnNpb24gZm9yIHRoZSBwcmV2aW91c2x5IGFwcGxpZWQgY2F0YWxvZyIsCiAgICAgICAgICAiZGlzcGxheV9hcyI6ICJDb25maWcgVmVyc2lvbiIsCiAgICAgICAgICAiZGVmYXVsdCI6IG51bGwKICAgICAgICB9LAogICAgICAgICJ0eXBlX2Rpc3RyaWJ1dGlvbiI6IHsKICAgICAgICAgICJkZXNjcmlwdGlvbiI6ICJSZXNvdXJjZSBjb3VudHMgcGVyIHR5cGUgbWFuYWdlZCBieSBQdXBwZXQiLAogICAgICAgICAgImRpc3BsYXlfYXMiOiAiVHlwZSBEaXN0cmlidXRpb24iLAogICAgICAgICAgImRlZmF1bHQiOiB7CiAgICAgICAgICB9CiAgICAgICAgfSwKICAgICAgICAic3VtbWFyeSI6IHsKICAgICAgICAgICJkZXNjcmlwdGlvbiI6ICJTdW1tYXJ5IGRhdGEgYXMgcHJvdmlkZWQgYnkgUHVwcGV0IiwKICAgICAgICAgICJkaXNwbGF5X2FzIjogIlN1bW1hcnkiLAogICAgICAgICAgImRlZmF1bHQiOiB7CiAgICAgICAgICB9CiAgICAgICAgfQogICAgICB9LAogICAgICAiZGlzcGxheSI6ICJhbHdheXMiLAogICAgICAiZGVzY3JpcHRpb24iOiAiR2V0IHRoZSBzdW1tYXJ5IG9mIHRoZSBsYXN0IFB1cHBldCBydW4iLAogICAgICAiYWdncmVnYXRlIjogWwogICAgICAgIHsKICAgICAgICAgICJmdW5jdGlvbiI6ICJhdmVyYWdlIiwKICAgICAgICAgICJhcmdzIjogWwogICAgICAgICAgICAiY29uZmlnX3JldHJpZXZhbF90aW1lIiwKICAgICAgICAgICAgewogICAgICAgICAgICAgICJmb3JtYXQiOiAiQXZlcmFnZTogJTAuMmYiCiAgICAgICAgICAgIH0KICAgICAgICAgIF0KICAgICAgICB9LAogICAgICAgIHsKICAgICAgICAgICJmdW5jdGlvbiI6ICJhdmVyYWdlIiwKICAgICAgICAgICJhcmdzIjogWwogICAgICAgICAgICAidG90YWxfdGltZSIsCiAgICAgICAgICAgIHsKICAgICAgICAgICAgICAiZm9ybWF0IjogIkF2ZXJhZ2U6ICUwLjJmIgogICAgICAgICAgICB9CiAgICAgICAgICBdCiAgICAgICAgfSwKICAgICAgICB7CiAgICAgICAgICAiZnVuY3Rpb24iOiAiYXZlcmFnZSIsCiAgICAgICAgICAiYXJncyI6IFsKICAgICAgICAgICAgInRvdGFsX3Jlc291cmNlcyIsCiAgICAgICAgICAgIHsKICAgICAgICAgICAgICAiZm9ybWF0IjogIkF2ZXJhZ2U6ICVkIgogICAgICAgICAgICB9CiAgICAgICAgICBdCiAgICAgICAgfQogICAgICBdCiAgICB9LAogICAgewogICAgICAiYWN0aW9uIjogInJlc291cmNlIiwKICAgICAgImlucHV0IjogewogICAgICAgICJuYW1lIjogewogICAgICAgICAgInByb21wdCI6ICJOYW1lIiwKICAgICAgICAgICJkZXNjcmlwdGlvbiI6ICJSZXNvdXJjZSBOYW1lIiwKICAgICAgICAgICJ0eXBlIjogInN0cmluZyIsCiAgICAgICAgICAiZGVmYXVsdCI6IG51bGwsCiAgICAgICAgICAib3B0aW9uYWwiOiBmYWxzZSwKICAgICAgICAgICJ2YWxpZGF0aW9uIjogIl4uKyQiLAogICAgICAgICAgIm1heGxlbmd0aCI6IDE1MAogICAgICAgIH0sCiAgICAgICAgInR5cGUiOiB7CiAgICAgICAgICAicHJvbXB0IjogIlR5cGUiLAogICAgICAgICAgImRlc2NyaXB0aW9uIjogIlJlc291cmNlIFR5cGUiLAogICAgICAgICAgInR5cGUiOiAic3RyaW5nIiwKICAgICAgICAgICJkZWZhdWx0IjogbnVsbCwKICAgICAgICAgICJvcHRpb25hbCI6IGZhbHNlLAogICAgICAgICAgInZhbGlkYXRpb24iOiAiXi4rJCIsCiAgICAgICAgICAibWF4bGVuZ3RoIjogNTAKICAgICAgICB9LAogICAgICAgICJlbnZpcm9ubWVudCI6IHsKICAgICAgICAgICJwcm9tcHQiOiAiRW52aXJvbm1lbnQiLAogICAgICAgICAgImRlc2NyaXB0aW9uIjogIldoaWNoIFB1cHBldCBlbnZpcm9ubWVudCB0byB1c2UiLAogICAgICAgICAgInR5cGUiOiAic3RyaW5nIiwKICAgICAgICAgICJkZWZhdWx0IjogbnVsbCwKICAgICAgICAgICJvcHRpb25hbCI6IHRydWUsCiAgICAgICAgICAidmFsaWRhdGlvbiI6ICJwdXBwZXRfdmFyaWFibGUiLAogICAgICAgICAgIm1heGxlbmd0aCI6IDUwCiAgICAgICAgfQogICAgICB9LAogICAgICAib3V0cHV0IjogewogICAgICAgICJyZXN1bHQiOiB7CiAgICAgICAgICAiZGVzY3JpcHRpb24iOiAiVGhlIHJlc3VsdCBmcm9tIHRoZSBQdXBwZXQgcmVzb3VyY2UiLAogICAgICAgICAgImRpc3BsYXlfYXMiOiAiUmVzdWx0IiwKICAgICAgICAgICJkZWZhdWx0IjogIiIKICAgICAgICB9LAogICAgICAgICJjaGFuZ2VkIjogewogICAgICAgICAgImRlc2NyaXB0aW9uIjogIldhcyBhIGNoYW5nZSBhcHBsaWVkIGJhc2VkIG9uIHRoZSByZXNvdXJjZSIsCiAgICAgICAgICAiZGlzcGxheV9hcyI6ICJDaGFuZ2VkIiwKICAgICAgICAgICJkZWZhdWx0IjogbnVsbAogICAgICAgIH0KICAgICAgfSwKICAgICAgImRpc3BsYXkiOiAiYWx3YXlzIiwKICAgICAgImRlc2NyaXB0aW9uIjogIkV2YWx1YXRlIFB1cHBldCBSQUwgcmVzb3VyY2VzIiwKICAgICAgImFnZ3JlZ2F0ZSI6IFsKICAgICAgICB7CiAgICAgICAgICAiZnVuY3Rpb24iOiAiYm9vbGVhbl9zdW1tYXJ5IiwKICAgICAgICAgICJhcmdzIjogWwogICAgICAgICAgICAiY2hhbmdlZCIsCiAgICAgICAgICAgIHsKICAgICAgICAgICAgICAidHJ1ZSI6ICJDaGFuZ2VkIiwKICAgICAgICAgICAgICAiZmFsc2UiOiAiTm8gQ2hhbmdlIgogICAgICAgICAgICB9CiAgICAgICAgICBdCiAgICAgICAgfQogICAgICBdCiAgICB9LAogICAgewogICAgICAiYWN0aW9uIjogInJ1bm9uY2UiLAogICAgICAiaW5wdXQiOiB7CiAgICAgICAgImZvcmNlIjogewogICAgICAgICAgInByb21wdCI6ICJGb3JjZSIsCiAgICAgICAgICAiZGVzY3JpcHRpb24iOiAiV2lsbCBmb3JjZSBhIHJ1biBpbW1lZGlhdGVseSBlbHNlIHN1YmplY3QgdG8gZGVmYXVsdCBzcGxheSB0aW1lIiwKICAgICAgICAgICJ0eXBlIjogImJvb2xlYW4iLAogICAgICAgICAgImRlZmF1bHQiOiBudWxsLAogICAgICAgICAgIm9wdGlvbmFsIjogdHJ1ZQogICAgICAgIH0sCiAgICAgICAgInNlcnZlciI6IHsKICAgICAgICAgICJwcm9tcHQiOiAiUHVwcGV0IE1hc3RlciIsCiAgICAgICAgICAiZGVzY3JpcHRpb24iOiAiQWRkcmVzcyBhbmQgcG9ydCBvZiB0aGUgUHVwcGV0IE1hc3RlciBpbiBzZXJ2ZXI6cG9ydCBmb3JtYXQiLAogICAgICAgICAgInR5cGUiOiAic3RyaW5nIiwKICAgICAgICAgICJkZWZhdWx0IjogbnVsbCwKICAgICAgICAgICJvcHRpb25hbCI6IHRydWUsCiAgICAgICAgICAidmFsaWRhdGlvbiI6ICJwdXBwZXRfc2VydmVyX2FkZHJlc3MiLAogICAgICAgICAgIm1heGxlbmd0aCI6IDUwCiAgICAgICAgfSwKICAgICAgICAidGFncyI6IHsKICAgICAgICAgICJwcm9tcHQiOiAiVGFncyIsCiAgICAgICAgICAiZGVzY3JpcHRpb24iOiAiUmVzdHJpY3QgdGhlIFB1cHBldCBydW4gdG8gYSBjb21tYSBsaXN0IG9mIHRhZ3MiLAogICAgICAgICAgInR5cGUiOiAic3RyaW5nIiwKICAgICAgICAgICJkZWZhdWx0IjogbnVsbCwKICAgICAgICAgICJvcHRpb25hbCI6IHRydWUsCiAgICAgICAgICAidmFsaWRhdGlvbiI6ICJwdXBwZXRfdGFncyIsCiAgICAgICAgICAibWF4bGVuZ3RoIjogMTIwCiAgICAgICAgfSwKICAgICAgICAibm9vcCI6IHsKICAgICAgICAgICJwcm9tcHQiOiAiTm8tb3AiLAogICAgICAgICAgImRlc2NyaXB0aW9uIjogIkRvIGEgUHVwcGV0IGRyeSBydW4iLAogICAgICAgICAgInR5cGUiOiAiYm9vbGVhbiIsCiAgICAgICAgICAiZGVmYXVsdCI6IG51bGwsCiAgICAgICAgICAib3B0aW9uYWwiOiB0cnVlCiAgICAgICAgfSwKICAgICAgICAic3BsYXkiOiB7CiAgICAgICAgICAicHJvbXB0IjogIlNwbGF5IiwKICAgICAgICAgICJkZXNjcmlwdGlvbiI6ICJTbGVlcCBmb3IgYSBwZXJpb2QgYmVmb3JlIGluaXRpYXRpbmcgdGhlIHJ1biIsCiAgICAgICAgICAidHlwZSI6ICJib29sZWFuIiwKICAgICAgICAgICJkZWZhdWx0IjogbnVsbCwKICAgICAgICAgICJvcHRpb25hbCI6IHRydWUKICAgICAgICB9LAogICAgICAgICJzcGxheWxpbWl0IjogewogICAgICAgICAgInByb21wdCI6ICJTcGxheSBMaW1pdCIsCiAgICAgICAgICAiZGVzY3JpcHRpb24iOiAiTWF4aW11bSBhbW91bnQgb2YgdGltZSB0byBzbGVlcCBiZWZvcmUgcnVuIiwKICAgICAgICAgICJ0eXBlIjogIm51bWJlciIsCiAgICAgICAgICAiZGVmYXVsdCI6IG51bGwsCiAgICAgICAgICAib3B0aW9uYWwiOiB0cnVlCiAgICAgICAgfSwKICAgICAgICAiZW52aXJvbm1lbnQiOiB7CiAgICAgICAgICAicHJvbXB0IjogIkVudmlyb25tZW50IiwKICAgICAgICAgICJkZXNjcmlwdGlvbiI6ICJXaGljaCBQdXBwZXQgZW52aXJvbm1lbnQgdG8gcnVuIiwKICAgICAgICAgICJ0eXBlIjogInN0cmluZyIsCiAgICAgICAgICAiZGVmYXVsdCI6IG51bGwsCiAgICAgICAgICAib3B0aW9uYWwiOiB0cnVlLAogICAgICAgICAgInZhbGlkYXRpb24iOiAicHVwcGV0X3ZhcmlhYmxlIiwKICAgICAgICAgICJtYXhsZW5ndGgiOiA1MAogICAgICAgIH0sCiAgICAgICAgInVzZV9jYWNoZWRfY2F0YWxvZyI6IHsKICAgICAgICAgICJwcm9tcHQiOiAiVXNlIENhY2hlZCBDYXRhbG9nIiwKICAgICAgICAgICJkZXNjcmlwdGlvbiI6ICJEZXRlcm1pbmUgaWYgdG8gdXNlIHRoZSBjYWNoZWQgY2F0YWxvZyBvciBub3QiLAogICAgICAgICAgInR5cGUiOiAiYm9vbGVhbiIsCiAgICAgICAgICAiZGVmYXVsdCI6IG51bGwsCiAgICAgICAgICAib3B0aW9uYWwiOiB0cnVlCiAgICAgICAgfQogICAgICB9LAogICAgICAib3V0cHV0IjogewogICAgICAgICJzdW1tYXJ5IjogewogICAgICAgICAgImRlc2NyaXB0aW9uIjogIlN1bW1hcnkgb2YgY29tbWFuZCBydW4iLAogICAgICAgICAgImRpc3BsYXlfYXMiOiAiU3VtbWFyeSIsCiAgICAgICAgICAiZGVmYXVsdCI6ICIiCiAgICAgICAgfSwKICAgICAgICAiaW5pdGlhdGVkX2F0IjogewogICAgICAgICAgImRlc2NyaXB0aW9uIjogIlRpbWVzdGFtcCBvZiB3aGVuIHRoZSBydW5vbmNlIGNvbW1hbmQgd2FzIGlzc3VlcyIsCiAgICAgICAgICAiZGlzcGxheV9hcyI6ICJJbml0aWF0ZWQgYXQiLAogICAgICAgICAgImRlZmF1bHQiOiAwCiAgICAgICAgfQogICAgICB9LAogICAgICAiZGlzcGxheSI6ICJmYWlsZWQiLAogICAgICAiZGVzY3JpcHRpb24iOiAiSW52b2tlIGEgc2luZ2xlIFB1cHBldCBydW4iCiAgICB9LAogICAgewogICAgICAiYWN0aW9uIjogInN0YXR1cyIsCiAgICAgICJpbnB1dCI6IHsKICAgICAgfSwKICAgICAgIm91dHB1dCI6IHsKICAgICAgICAiYXBwbHlpbmciOiB7CiAgICAgICAgICAiZGVzY3JpcHRpb24iOiAiSXMgYSBjYXRhbG9nIGJlaW5nIGFwcGxpZWQiLAogICAgICAgICAgImRpc3BsYXlfYXMiOiAiQXBwbHlpbmciLAogICAgICAgICAgImRlZmF1bHQiOiBmYWxzZQogICAgICAgIH0sCiAgICAgICAgImlkbGluZyI6IHsKICAgICAgICAgICJkZXNjcmlwdGlvbiI6ICJJcyB0aGUgUHVwcGV0IGFnZW50IGRhZW1vbiBydW5uaW5nIGJ1dCBub3QgZG9pbmcgYW55IHdvcmsiLAogICAgICAgICAgImRpc3BsYXlfYXMiOiAiSWRsaW5nIiwKICAgICAgICAgICJkZWZhdWx0IjogZmFsc2UKICAgICAgICB9LAogICAgICAgICJlbmFibGVkIjogewogICAgICAgICAgImRlc2NyaXB0aW9uIjogIklzIHRoZSBhZ2VudCBjdXJyZW50bHkgbG9ja2VkIiwKICAgICAgICAgICJkaXNwbGF5X2FzIjogIkVuYWJsZWQiLAogICAgICAgICAgImRlZmF1bHQiOiBudWxsCiAgICAgICAgfSwKICAgICAgICAiZGFlbW9uX3ByZXNlbnQiOiB7CiAgICAgICAgICAiZGVzY3JpcHRpb24iOiAiSXMgdGhlIFB1cHBldCBhZ2VudCBkYWVtb24gcnVubmluZyBvbiB0aGlzIHN5c3RlbSIsCiAgICAgICAgICAiZGlzcGxheV9hcyI6ICJEYWVtb24gUnVubmluZyIsCiAgICAgICAgICAiZGVmYXVsdCI6IGZhbHNlCiAgICAgICAgfSwKICAgICAgICAibGFzdHJ1biI6IHsKICAgICAgICAgICJkZXNjcmlwdGlvbiI6ICJXaGVuIHRoZSBBZ2VudCBsYXN0IGFwcGxpZWQgYSBjYXRhbG9nIGluIGxvY2FsIHRpbWUiLAogICAgICAgICAgImRpc3BsYXlfYXMiOiAiTGFzdCBSdW4iLAogICAgICAgICAgImRlZmF1bHQiOiAwCiAgICAgICAgfSwKICAgICAgICAic2luY2VfbGFzdHJ1biI6IHsKICAgICAgICAgICJkZXNjcmlwdGlvbiI6ICJIb3cgbG9uZyBhZ28gZGlkIHRoZSBBZ2VudCBsYXN0IGFwcGx5IGEgY2F0YWxvZyBpbiBsb2NhbCB0aW1lIiwKICAgICAgICAgICJkaXNwbGF5X2FzIjogIlNpbmNlIExhc3QgUnVuIiwKICAgICAgICAgICJkZWZhdWx0IjogIlVua25vd24iCiAgICAgICAgfSwKICAgICAgICAic3RhdHVzIjogewogICAgICAgICAgImRlc2NyaXB0aW9uIjogIkN1cnJlbnQgc3RhdHVzIG9mIHRoZSBQdXBwZXQgYWdlbnQiLAogICAgICAgICAgImRpc3BsYXlfYXMiOiAiU3RhdHVzIiwKICAgICAgICAgICJkZWZhdWx0IjogInVua25vd24iCiAgICAgICAgfSwKICAgICAgICAiZGlzYWJsZV9tZXNzYWdlIjogewogICAgICAgICAgImRlc2NyaXB0aW9uIjogIk1lc3NhZ2Ugc3VwcGxpZWQgd2hlbiBhZ2VudCB3YXMgZGlzYWJsZWQiLAogICAgICAgICAgImRpc3BsYXlfYXMiOiAiTG9jayBNZXNzYWdlIiwKICAgICAgICAgICJkZWZhdWx0IjogIiIKICAgICAgICB9LAogICAgICAgICJtZXNzYWdlIjogewogICAgICAgICAgImRlc2NyaXB0aW9uIjogIkRlc2NyaXB0aXZlIG1lc3NhZ2UgZGVmaW5pbmcgdGhlIG92ZXJhbGwgYWdlbnQgc3RhdHVzIiwKICAgICAgICAgICJkaXNwbGF5X2FzIjogIk1lc3NhZ2UiLAogICAgICAgICAgImRlZmF1bHQiOiAidW5rbm93biIKICAgICAgICB9CiAgICAgIH0sCiAgICAgICJkaXNwbGF5IjogImFsd2F5cyIsCiAgICAgICJkZXNjcmlwdGlvbiI6ICJHZXQgdGhlIGN1cnJlbnQgc3RhdHVzIG9mIHRoZSBQdXBwZXQgYWdlbnQiLAogICAgICAiYWdncmVnYXRlIjogWwogICAgICAgIHsKICAgICAgICAgICJmdW5jdGlvbiI6ICJib29sZWFuX3N1bW1hcnkiLAogICAgICAgICAgImFyZ3MiOiBbCiAgICAgICAgICAgICJlbmFibGVkIiwKICAgICAgICAgICAgewogICAgICAgICAgICAgICJ0cnVlIjogImVuYWJsZWQiLAogICAgICAgICAgICAgICJmYWxzZSI6ICJkaXNhYmxlZCIKICAgICAgICAgICAgfQogICAgICAgICAgXQogICAgICAgIH0sCiAgICAgICAgewogICAgICAgICAgImZ1bmN0aW9uIjogImJvb2xlYW5fc3VtbWFyeSIsCiAgICAgICAgICAiYXJncyI6IFsKICAgICAgICAgICAgImRhZW1vbl9wcmVzZW50IiwKICAgICAgICAgICAgewogICAgICAgICAgICAgICJ0cnVlIjogInJ1bm5pbmciLAogICAgICAgICAgICAgICJmYWxzZSI6ICJzdG9wcGVkIgogICAgICAgICAgICB9CiAgICAgICAgICBdCiAgICAgICAgfSwKICAgICAgICB7CiAgICAgICAgICAiZnVuY3Rpb24iOiAic3VtbWFyeSIsCiAgICAgICAgICAiYXJncyI6IFsKICAgICAgICAgICAgImFwcGx5aW5nIgogICAgICAgICAgXQogICAgICAgIH0sCiAgICAgICAgewogICAgICAgICAgImZ1bmN0aW9uIjogInN1bW1hcnkiLAogICAgICAgICAgImFyZ3MiOiBbCiAgICAgICAgICAgICJzdGF0dXMiCiAgICAgICAgICBdCiAgICAgICAgfSwKICAgICAgICB7CiAgICAgICAgICAiZnVuY3Rpb24iOiAic3VtbWFyeSIsCiAgICAgICAgICAiYXJncyI6IFsKICAgICAgICAgICAgImlkbGluZyIKICAgICAgICAgIF0KICAgICAgICB9CiAgICAgIF0KICAgIH0KICBdCn0="
