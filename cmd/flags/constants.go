package flags

// viper and cobra have one way binding i.e all access
// should go through viper which uses unsafe string access
// This is why we define flags as constants so we can have safety in case we change
// flag names in future

// General
const DataDir = "data-dir"
const Color = "color"

// Init Flags
const NoSpawn = "no-spawn"

// Logging Flags
const Verbose = "verbose"
const LogConsole = "log-console"
const LogFormat = "log-format"
const NoLogFile = "no-log-file"
const LogFileName = "log-file-name"

// Telemetry Flags
const NoTelemetry = "no-telemetry"
const TelemetryInspect = "telemetry-inspect"
const TelemtryDebug = "telemetry-debug"

// Sentry flags
const SentryDebug = "sentry-debug"
const SentryDSN = "sentry-dsn"
