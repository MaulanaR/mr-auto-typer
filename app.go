// -------- File: app.go --------
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync/atomic"
	"time"

	"github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
	wruntime "github.com/wailsapp/wails/v2/pkg/runtime"
)

// NOTE: Use for lawful local testing/automation only.

type Token struct{ Kind, Value string }

type Preset struct {
	Text        string `json:"text"`
	MinMS       int    `json:"min_ms"`
	MaxMS       int    `json:"max_ms"`
	Loops       int    `json:"loops"`
	LoopDelayMS int    `json:"loop_delay_ms"`
	LoopJitter  int    `json:"loop_jitter_ms"`
}

type App struct {
	ctx     context.Context
	running atomic.Bool
	cancel  atomic.Bool
}

func NewApp() *App { rand.Seed(time.Now().UnixNano()); return &App{} }

func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	go a.listenHotkeyStop()
}

// ---- Public methods (bound to frontend) ----
func (a *App) Start(p Preset) {
	if a.running.Load() {
		return
	}
	if strings.TrimSpace(p.Text) == "" {
		wruntime.EventsEmit(a.ctx, "status", "Input kosong")
		return
	}
	if p.MaxMS < p.MinMS {
		p.MinMS, p.MaxMS = p.MaxMS, p.MinMS
	}
	a.cancel.Store(false)
	a.running.Store(true)
	go a.executeTyping(p)
}

func (a *App) Stop() { a.cancel.Store(true) }

func (a *App) SavePreset(p Preset) (string, error) {
	path := defaultPresetPath()
	return path, savePreset(path, p)
}
func (a *App) LoadPreset() (Preset, error) { return loadPreset(defaultPresetPath()) }
func (a *App) OS() string                  { return runtime.GOOS }

// ---- Core typing ----
func (a *App) executeTyping(p Preset) {
	defer func() { a.running.Store(false); wruntime.EventsEmit(a.ctx, "status", "Idle") }()

	// Countdown 3, 2, 1
	for i := 3; i > 0; i-- {
		if a.cancel.Load() {
			return
		}
		wruntime.EventsEmit(a.ctx, "status", fmt.Sprintf("Mulai dalam %d... (Segera pindah ke aplikasi target)", i))
		time.Sleep(2 * time.Second)
	}

	if a.cancel.Load() {
		return
	}

	wruntime.EventsEmit(a.ctx, "status", "Running… Fokuskan ke window target. Ctrl+Alt+S untuk STOP")

	minD := time.Duration(p.MinMS) * time.Millisecond
	maxD := time.Duration(p.MaxMS) * time.Millisecond
	loopDelay := time.Duration(p.LoopDelayMS) * time.Millisecond
	jitter := time.Duration(p.LoopJitter) * time.Millisecond

	toks := tokenize(p.Text)
	if len(toks) == 0 {
		wruntime.EventsEmit(a.ctx, "status", "Tidak ada token")
		return
	}

	iter := 0
	for {
		if a.cancel.Load() {
			return
		}
		for _, tk := range toks {
			if a.cancel.Load() {
				return
			}
			if tk.Kind == "text" {
				typeText(tk.Value, minD, maxD, &a.cancel)
			} else {
				pressCmd(tk.Value, minD)
			}
		}
		iter++
		if p.Loops > 0 && iter >= p.Loops {
			return
		}
		d := loopDelay
		if jitter > 0 {
			off := time.Duration(rand.Int63n(int64(jitter)*2+1)) - jitter
			d += off
			if d < 0 {
				d = 0
			}
		}
		if d > 0 {
			time.Sleep(d)
		}
	}
}

func tokenize(s string) []Token {
	var out []Token
	lower := strings.ToLower(s)
	i := 0
	for i < len(s) {
		if strings.HasPrefix(lower[i:], "{backspace}") {
			out = append(out, Token{"cmd", "backspace"})
			i += len("{backspace}")
		} else if strings.HasPrefix(lower[i:], "{delete}") {
			out = append(out, Token{"cmd", "delete"})
			i += len("{delete}")
		} else if strings.HasPrefix(lower[i:], "{enter}") {
			out = append(out, Token{"cmd", "enter"})
			i += len("{enter}")
		} else if strings.HasPrefix(lower[i:], "{tab}") {
			out = append(out, Token{"cmd", "tab"})
			i += len("{tab}")
		} else {
			// ambil karakter satu per satu sampai ketemu token
			start := i
			for i < len(s) && !strings.HasPrefix(lower[i:], "{backspace}") &&
				!strings.HasPrefix(lower[i:], "{delete}") &&
				!strings.HasPrefix(lower[i:], "{enter}") &&
				!strings.HasPrefix(lower[i:], "{tab}") {
				i++
			}
			out = append(out, Token{"text", s[start:i]})
		}
	}
	return out
}

func typeText(text string, minDelay, maxDelay time.Duration, cancel *atomic.Bool) {
	for _, r := range text {
		if cancel.Load() {
			return
		}
		robotgo.TypeStr(string(r))
		d := minDelay
		if maxDelay > minDelay {
			span := maxDelay - minDelay
			d = minDelay + time.Duration(rand.Int63n(int64(span)+1))
		}
		time.Sleep(d)
	}
}

func pressCmd(cmd string, minDelay time.Duration) {
	switch cmd {
	case "backspace":
		robotgo.KeyTap("backspace")
	case "delete":
		robotgo.KeyTap("delete")
	case "enter":
		robotgo.KeyTap("enter")
	case "tab":
		robotgo.KeyTap("tab")
	}
	time.Sleep(minDelay)
}

// Global hotkey Ctrl+Alt+S to stop
func (a *App) listenHotkeyStop() {
	for {
		if a.running.Load() {
			// true saat user menekan Ctrl+Alt+S
			if hook.AddEvents("s", "alt", "ctrl") {
				a.cancel.Store(true)
				wruntime.EventsEmit(a.ctx, "status", "Stopped by hotkey")
				time.Sleep(300 * time.Millisecond) // debounce
			}
		} else {
			time.Sleep(200 * time.Millisecond)
		}
	}
}

// Preset I/O
func defaultPresetPath() string {
	exe, _ := os.Executable()
	return filepath.Join(filepath.Dir(exe), "preset.json")
}
func savePreset(path string, p Preset) error {
	b, _ := json.MarshalIndent(p, "", "  ")
	return os.WriteFile(path, b, 0o644)
}
func loadPreset(path string) (Preset, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return Preset{}, err
	}
	var p Preset
	err = json.Unmarshal(b, &p)
	return p, err
}
