// Copyright (c) Liam Stanley <me@liamstanley.io>. All rights reserved. Use
// of this source code is governed by the MIT license that can be found in
// the LICENSE file.
//
// Code generated by cmd/codegen. DO NOT EDIT.
//
// Workarounds Option Group

package ytdlp

import (
	"strconv"
)

// Force the specified encoding (experimental)
//
// Encoding maps to cli flags: --encoding=ENCODING.
func (c *Command) Encoding(encoding string) *Command {
	c.addFlag(&Flag{
		ID:   "encoding",
		Flag: "--encoding",
		Args: []string{encoding},
	})
	return c
}

// Explicitly allow HTTPS connection to servers that do not support RFC 5746 secure
// renegotiation
//
// LegacyServerConnect maps to cli flags: --legacy-server-connect.
func (c *Command) LegacyServerConnect() *Command {
	c.addFlag(&Flag{
		ID:   "legacy_server_connect",
		Flag: "--legacy-server-connect",
		Args: nil,
	})
	return c
}

// Suppress HTTPS certificate validation
//
// NoCheckCertificates maps to cli flags: --no-check-certificates.
func (c *Command) NoCheckCertificates() *Command {
	c.addFlag(&Flag{
		ID:   "no_check_certificate",
		Flag: "--no-check-certificates",
		Args: nil,
	})
	return c
}

// Use an unencrypted connection to retrieve information about the video (Currently
// supported only for YouTube)
//
// PreferInsecure maps to cli flags: --prefer-insecure/--prefer-unsecure.
func (c *Command) PreferInsecure() *Command {
	c.addFlag(&Flag{
		ID:   "prefer_insecure",
		Flag: "--prefer-insecure",
		Args: nil,
	})
	return c
}

// UserAgent maps to cli flags: --user-agent=UA.
func (c *Command) UserAgent(ua string) *Command {
	c.addFlag(&Flag{
		ID:   "user_agent",
		Flag: "--user-agent",
		Args: []string{ua},
	})
	return c
}

// Referer maps to cli flags: --referer=URL.
func (c *Command) Referer(url string) *Command {
	c.addFlag(&Flag{
		ID:   "referer",
		Flag: "--referer",
		Args: []string{url},
	})
	return c
}

// Specify a custom HTTP header and its value, separated by a colon ":". You can
// use this option multiple times
//
// AddHeaders maps to cli flags: --add-headers=FIELD:VALUE.
func (c *Command) AddHeaders(fieldvalue string) *Command {
	c.addFlag(&Flag{
		ID:   "headers",
		Flag: "--add-headers",
		Args: []string{fieldvalue},
	})
	return c
}

// Work around terminals that lack bidirectional text support. Requires bidiv or
// fribidi executable in PATH
//
// BidiWorkaround maps to cli flags: --bidi-workaround.
func (c *Command) BidiWorkaround() *Command {
	c.addFlag(&Flag{
		ID:   "bidi_workaround",
		Flag: "--bidi-workaround",
		Args: nil,
	})
	return c
}

// Number of seconds to sleep between requests during data extraction
//
// SleepRequests maps to cli flags: --sleep-requests=SECONDS.
func (c *Command) SleepRequests(seconds float64) *Command {
	c.addFlag(&Flag{
		ID:   "sleep_interval_requests",
		Flag: "--sleep-requests",
		Args: []string{
			strconv.FormatFloat(seconds, 'g', -1, 64),
		},
	})
	return c
}

// Number of seconds to sleep before each download. This is the minimum time to
// sleep when used along with --max-sleep-interval (Alias: --min-sleep-interval)
//
// SleepInterval maps to cli flags: --sleep-interval/--min-sleep-interval=SECONDS.
func (c *Command) SleepInterval(seconds float64) *Command {
	c.addFlag(&Flag{
		ID:   "sleep_interval",
		Flag: "--sleep-interval",
		Args: []string{
			strconv.FormatFloat(seconds, 'g', -1, 64),
		},
	})
	return c
}

// Maximum number of seconds to sleep. Can only be used along with
// --min-sleep-interval
//
// MaxSleepInterval maps to cli flags: --max-sleep-interval=SECONDS.
func (c *Command) MaxSleepInterval(seconds float64) *Command {
	c.addFlag(&Flag{
		ID:   "max_sleep_interval",
		Flag: "--max-sleep-interval",
		Args: []string{
			strconv.FormatFloat(seconds, 'g', -1, 64),
		},
	})
	return c
}

// Number of seconds to sleep before each subtitle download
//
// SleepSubtitles maps to cli flags: --sleep-subtitles=SECONDS.
func (c *Command) SleepSubtitles(seconds int) *Command {
	c.addFlag(&Flag{
		ID:   "sleep_interval_subtitles",
		Flag: "--sleep-subtitles",
		Args: []string{
			strconv.Itoa(seconds),
		},
	})
	return c
}