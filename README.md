# DNS Lookup Tool 🔍

A fast and powerful DNS record lookup tool written in Go. Query different types of DNS records with ease!

![GitHub](https://img.shields.io/github/license/Lokidres/dns-lookup)
![Go Version](https://img.shields.io/badge/Go-1.16%2B-blue)

## Features ✨

- 📝 Multiple DNS record type support:
  - A Records (IPv4 addresses)
  - MX Records (mail servers)
  - NS Records (name servers)
  - TXT Records (text records)
- 🚀 Query multiple record types simultaneously
- 🔄 Automatic URL cleaning (removes http:// and https://)
- 💻 Simple command-line interface
- 📚 Comprehensive help documentation

## Installation 🛠️

```bash
# Clone the repository
git clone https://github.com/Lokidres/dns-lookup.git

# Navigate to the project directory
cd dns-lookup

# Build the project
go build
```

## Usage 📖

Basic usage pattern:
```bash
./dns-lookup [OPTIONS] <domain>
```

### Options

- `--help`: Show help message
- `--a`: Look up A records
- `--mx`: Look up MX records
- `--ns`: Look up NS records
- `--txt`: Look up TXT records
- `--zone`: Look up all available records

### Examples

1. Look up A records:
```bash
./dns-lookup --a google.com
```

2. Look up MX and NS records:
```bash
./dns-lookup --mx --ns example.com
```

3. Look up all available records:
```bash
./dns-lookup --zone github.com
```

## Sample Output 📋

```
=== A Records ===
142.250.187.78

=== MX Records ===
aspmx.l.google.com    preference: 10
alt1.aspmx.l.google.com    preference: 20

=== NS Records ===
ns1.google.com
ns2.google.com
```

## Requirements 📋

- Go 1.16 or higher
- Internet connection for DNS queries

## Contributing 🤝

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License 📄

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Author ✍️

**Lokidres**
- GitHub: [@Lokidres](https://github.com/Lokidres)

## Acknowledgments 🙏

- Go DNS package contributors
- The Go community

## Future Plans 🚀

- [ ] Add AAAA record support (IPv6)
- [ ] Add CNAME record support
- [ ] Add TTL information
- [ ] Add JSON/CSV output formats
- [ ] Implement concurrent queries
- [ ] Add caching mechanism

## Support 💬

If you have any questions or run into issues, please open an issue on the [GitHub repository](https://github.com/Lokidres/dns-lookup/issues). 