# mr-auto-typer
Mrlabs - Auto Typer
<img width="500" height="500" alt="_9b0eaa11-7724-4e77-9fcf-f42e345e9448-removebg-preview" src="https://github.com/user-attachments/assets/1fe34281-c830-40d1-9b1e-e200aaea1be7" />

Example : 
<img width="1058" height="743" alt="image" src="https://github.com/user-attachments/assets/5392ac1c-3679-4b59-a1ce-cfefd6bc127f" />
<img width="1052" height="736" alt="image" src="https://github.com/user-attachments/assets/1ba0fbed-b0bb-491c-9d51-3f2d24333ef7" />

The Auto Typer Software Utility can be used to type Text on Keyboard with a configurable Hot Key or Shortcut Key.
Auto Typer to Type Text easily and quickly without manual text typing.

Key features
- Free text input: duplicates whatever you type
- Delay time: set the interval between keystrokes
- Jitter: small random variations in delay to simulate human typing
- Loop: repeat the input a set number of times or indefinitely
- Delay loop shortcut: a shortcut to add a delay between loops
- Presets: save and load frequently used configurations

Technology stack
- Go (Golang)
- Wails (native GUI integration)
- JavaScript (front-end logic)

Quick usage
1. Open the application.
2. Paste or type your text into the input field.
3. Configure Delay, Jitter, Loop, and Preset options.
4. Save the preset
5. Press Start to run the auto-typing sequence.

## ✔️ Donations
If this project helped you and you'd like to say thanks, please consider sponsoring the project.

## 📦 Download

### Release Stabil
Download stable version from [Releases](https://github.com/MaulanaR/mr-auto-typer/releases) page.

## 📁 Project Structure

```
mr-auto-typer/
├── app.go              # Main application logic
├── main.go            # Wails entry point
├── go.mod             # Go modules
├── wails.json         # Wails configuration
├── frontend/          # Frontend assets
│   ├── index.html     # Main HTML
│   ├── mrlabs.ico     # App icon
│   └── wailsjs/      # Generated Wails bindings
├── .github/
│   └── workflows/    # CI/CD workflows
├── build/            # Build output
└── dist/             # Packaged artifacts
```

## 🎨 UI Features

### Responsive Breakpoints
- **Desktop**: > 768px (3-column layout)
- **Tablet**: ≤ 768px (1-column layout)
- **Mobile**: ≤ 480px (vertical toolbar)

### Keyboard Shortcuts
- `Ctrl+Alt+Enter`: Start typing
- `Ctrl+Alt+S`: Stop typing

## 🔒 Security
- **Gosec**: Static security analysis
- **Dependency Scanning**: Automated vulnerability checks
- **CodeQL**: Advanced security analysis
- **SARIF Reports**: Security findings in GitHub

## 📝 Tokens

Special tokens can be used in text:
- `{backspace}`: Backspace key
- `{delete}`: Delete key  
- `{enter}`: Enter key
- `{tab}`: Tab key

## 🤝 Contributing

1. Fork repository
2. Create feature branch (`git checkout -b feature/amazing-feature`)
3. Commit changes (`git commit -m 'Add amazing feature'`)
4. Push ke branch (`git push origin feature/amazing-feature`)
5. Open Pull Request

## 📄 License

Not an opensource project !

## 🙏 Credits

- [Wails v2](https://wails.io/) - Cross-platform app framework
- [RobotGo](https://github.com/go-vgo/robotgo) - GUI automation
- [GoHook](https://github.com/robotn/gohook) - Global hotkeys

---

**MrLabs** - Building useful tools for developers 🚀
