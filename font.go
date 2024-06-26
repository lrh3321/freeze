package main

import (
	"encoding/base64"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/alecthomas/chroma/v2/formatters/svg"
	"github.com/charmbracelet/freeze/font"
)

func fontOptions(config *Config) ([]svg.Option, error) {
	if config.Font.File != "" {
		bts, err := os.ReadFile(config.Font.File)
		if err != nil {
			return nil, fmt.Errorf("invalid font file: %w", err)
		}

		var format svg.FontFormat
		switch ext := filepath.Ext(config.Font.File); strings.ToLower(ext) {
		case ".ttf":
			format = svg.TRUETYPE
		case ".woff2":
			format = svg.WOFF2
		case ".woff":
			format = svg.WOFF
		default:
			return nil, fmt.Errorf("%s is not a supported font extension", ext)
		}

		return []svg.Option{
			svg.EmbedFont(
				config.Font.Family,
				base64.StdEncoding.EncodeToString(bts),
				format,
			),
			svg.FontFamily(config.Font.Family),
		}, nil
	}
	if config.Font.Family != "JetBrains Mono" {
		return []svg.Option{
			svg.FontFamily(config.Font.Family),
		}, nil
	}
	config.Font.Family = "JetBrains Mono"
	fontBase64 := font.JetBrainsMono
	if !config.Font.Ligatures {
		fontBase64 = font.JetBrainsMonoNL
	}
	return []svg.Option{
		svg.EmbedFont(config.Font.Family, fontBase64, svg.WOFF2),
		svg.FontFamily(config.Font.Family),
	}, nil
}
