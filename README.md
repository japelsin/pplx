# Perplexity CLI

[![Go](https://github.com/japelsin/pplx/actions/workflows/release.yml/badge.svg)](https://github.com/japelsin/pplx/actions/workflows/release.yml)
[![License](https://img.shields.io/badge/license-MIT-blue)](https://github.com/japelsin/pplx/blob/main/LICENSE)

CLI for searching with [Perplexity](https://www.perplexity.ai/)'s API. Can also be used as a chatbot.

## Prerequisites

- Perplexity account and API key. You'll be prompted for the API key the first time you run `pplx`.

## Installation

### With [Homebrew](https://brew.sh)

```bash
brew install japelsin/tap/pplx
```

### From source

If you have [go](https://go.dev/) installed:

```bash
go install github.com/japelsin/pplx@latest
```

You could also grab the appropriate executable from [releases](https://github.com/japelsin/pplx/releases).

## Usage

### Search

Search command. Most parameters allowed by the [Perplexity API](https://docs.perplexity.ai/api-reference/chat-completions) are available as options. The defaults for some flags can also be set through the config (see below).

```
Usage:
  pplx search [flags]

Flags:
  -a, --api_key string                     API Key
  -f, --frequency_penalty float            Token frequency penalty [0, 1.0]
  -l, --max_tokens int                     Token limit per request
  -m, --model string                       Model to use
  -p, --presence_penalty float             Token presence penalty [-2.0, 2.0]
  -q, --query string                       Your query
  -d, --search_domain_filter stringArray   Domain filter (e.g. '-d https://x.com -d https://y.com')
  -r, --search_recency_filter string       Recency filter (month, week, day or hour)
  -s, --system_prompt string               System prompt
  -t, --temperature float                  Response randomness [0, 2.0]
  -K, --top_k int                          Number of tokens to sample from [0, 2048]
  -P, --top_p float                        Probability cutoff for token selection [0, 1.0]
```

### Config

The following config options are available:

```
Usage:
  pplx config set [command]

Available Commands:
  api_key       Set API key
  max_tokens    Set default max tokens
  model         Set default model
  stream        Set whether to stream response
  system_prompt Set default system prompt
```
