# Corsy - CORS Vulnerability Scanner

<div align="center">

[![Go Version](https://img.shields.io/github/go-mod/go-version/themirze/corsy)](https://github.com/themirze/corsy) [![License](https://img.shields.io/github/license/themirze/corsy)](https://github.com/themirze/corsy/blob/main/LICENSE) [![Release](https://img.shields.io/github/v/release/themirze/corsy)](https://github.com/themirze/corsy/releases)

</div>

## 🚀 Features

- ⚡️ **High Performance**: Written in Go for maximum speed and efficiency
- 🔄 **Concurrent Scanning**: Tests multiple URLs simultaneously
- 🎯 **Auto PoC Generation**: Creates HTML proof-of-concept files automatically
- 📝 **Detailed Reports**: Comprehensive output with vulnerability details
- 🎨 **Colored Output**: Easy-to-read terminal interface

## 🔧 Installation

### Quick Install (Linux/MacOS)

\`\`\`bash curl -s https://raw.githubusercontent.com/themirze/corsy/main/install.sh | bash \`\`\`

### Manual Installation

Download the latest binary for your platform from [Releases](https://github.com/themirze/corsy/releases)

#### Linux/MacOS

\`\`\`bash chmod +x corsy sudo mv corsy /usr/local/bin/ \`\`\`

#### Windows

Download and add to your PATH manually

## 📖 Usage

### Basic Usage

\`\`\`bash corsy -l urls.txt -o results.txt \`\`\`

### Command Line Options

\`\`\` -l URL list file (required) -o Output file for results (optional) \`\`\`

### Example URL List Format

\`\`\` https://example.com http://test.com vulnerable.site.com \`\`\`

## 📋 Example Output

\`\`\` [+] CORS Vuln: https://vulnerable-site.com └─> PoC created in output/poc/

[-] https://secure-site.com

[*] Job done. Found 1 vulnerable endpoints. Have a good hack ;) \`\`\`

## 🔍 How It Works

1. Reads URLs from the provided list file
2. Tests each URL for CORS misconfigurations
3. Generates PoC files for vulnerable endpoints
4. Creates a detailed report of findings

## 🛡️ Disclaimer

This tool is for educational purposes and authorized testing only. Always obtain explicit permission before testing any systems. The authors are not responsible for any misuse or damage.

## 🤝 Contributing

1. Fork the repository
2. Create your feature branch (\`git checkout -b feature/amazing-feature\`)
3. Commit your changes (\`git commit -m 'Add amazing feature'\`)
4. Push to the branch (\`git push origin feature/amazing-feature\`)
5. Open a Pull Request

## 📜 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🙏 Credits

Created by Baldwin

## 📞 Contact

- GitHub: [@themirze](https://github.com/themirze)

## 🔄 Updates

Check the [releases page](https://github.com/themirze/corsy/releases) for the latest updates and changes.

---

<div align="center">
Made with ❤️ by Baldwin
</div>
