# Nutanix Resource Provider

The Nutanix Resource Provider lets you manage [Nutanix](https://www.nutanix.com/) resources.

## Installing

This package is available for several languages/platforms:

### Node.js (JavaScript/TypeScript)

To use from JavaScript or TypeScript in Node.js, install using either `npm`:

```bash
npm install @jacobbmay/pulumi-nutanix
```

or `yarn`:

```bash
yarn add @jacobbmay/pulumi-nutanix
```

### Python

To use from Python, install using `pip`:

```bash
pip install pulumi-nutanix
```

### Go

To use from Go, use `go get` to grab the latest version of the library:

```bash
go get github.com/jacobbmay/pulumi-nutanix/sdk/go/...
```

### .NET

To use from .NET, install using `dotnet add package`:

```bash
dotnet add package JacobBMay.Nutanix
```

## Configuration

The following configuration points are available for the `nutanix` provider:

- `nutanix:apiKey` (environment: `NUTANIX_API_KEY`) - the API key for `nutanix`
- `nutanix:region` (environment: `NUTANIX_REGION`) - the region in which to deploy resources

## Reference

For detailed reference documentation, please visit [the Pulumi registry](https://www.pulumi.com/registry/packages/nutanix/api-docs/).
