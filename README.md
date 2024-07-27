# Microsoft API Wrapper

![GitHub Downloads (all assets, all releases)](https://img.shields.io/github/downloads/SecNex/ms-api-wrapper/total)

## Prerequirements

- Entra ID Enterprise Application registered in Azure AD/Entra ID tenant.
- Client ID, client secret, and tenant ID of the app registration.

## Quick Start

1. Download the latest release from GitHub.
2. Create the configuration file `config.json` in the same directory as the executable file.
3. Add the following content to the configuration file:

```json
{
    "client": {
        "client_id": "YOUR_CLIENT_ID",
        "client_secret": "YOUR_CLIENT_SECRET",
        "tenant_id": "YOUR_TENANT_ID"
    }
}
```

4. Start with `go run .` to run the application.

5. Create build with `bash build.sh` to create a binary file.

6. Run the binary file with `./secnex-ms-api-wrapper`.

## Configuration

The configuration file `config.json` must contain the following fields:

- `client_id`: The client ID of the Azure AD application.
- `client_secret`: The client secret of the Azure AD application.
- `tenant_id`: The tenant ID of the Azure AD application.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
