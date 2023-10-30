connection "akeyless" {
  plugin = "akeyless"

  # Path to the Akeyless CLI executable. Will pull from system path if not specified.
  # This can be set from the AKEYLESS_SHELLER_CLI_PATH environment variable.
  cli_path = ""

  # Name of the Akeyless CLI profile to use. Will use "default" if not specified.
  # This can be set from the AKEYLESS_SHELLER_PROFILE environment variable.
  profile = ""

  # Path to the .akeyless directory. Will use the default path if not specified of ~/.akeyless
  # This can be set from the AKEYLESS_SHELLER_HOME_DIRECTORY_PATH environment variable.
  akeyless_path = ""

  # Buffer time before token expiry to trigger re-authentication "2h" or "10m" (default) if not specified.
  # This can be set from the AKEYLESS_SHELLER_EXPIRY_BUFFER environment variable.
  expiry_buffer = "10m"

  # Debug flag to enable or disable debug logging, defaults to "" which is interpreted as false.
  # This can be set from the AKEYLESS_SHELLER_DEBUG environment variable.  export AKEYLESS_SHELLER_DEBUG=true
  debug = ""
}
