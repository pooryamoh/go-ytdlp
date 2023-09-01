// Copyright (c) Liam Stanley <me@liamstanley.io>. All rights reserved. Use
// of this source code is governed by the MIT license that can be found in
// the LICENSE file.
//
// Code generated by cmd/codegen. DO NOT EDIT.
//
// Verbosity Simulation Option Group

package ytdlp

// Activate quiet mode. If used with --verbose, print the log to stderr
//
// Quiet maps to cli flags: -q/--quiet.
func (c *Command) Quiet() *Command {
	c.addFlag(&Flag{
		ID:   "quiet",
		Flag: "--quiet",
		Args: nil,
	})
	return c
}

// Deactivate quiet mode. (Default)
//
// NoQuiet maps to cli flags: --no-quiet.
func (c *Command) NoQuiet() *Command {
	c.addFlag(&Flag{
		ID:   "quiet",
		Flag: "--no-quiet",
		Args: nil,
	})
	return c
}

// Ignore warnings
//
// NoWarnings maps to cli flags: --no-warnings.
func (c *Command) NoWarnings() *Command {
	c.addFlag(&Flag{
		ID:   "no_warnings",
		Flag: "--no-warnings",
		Args: nil,
	})
	return c
}

// Do not download the video and do not write anything to disk
//
// Simulate maps to cli flags: -s/--simulate.
func (c *Command) Simulate() *Command {
	c.addFlag(&Flag{
		ID:   "simulate",
		Flag: "--simulate",
		Args: nil,
	})
	return c
}

// Download the video even if printing/listing options are used
//
// NoSimulate maps to cli flags: --no-simulate.
func (c *Command) NoSimulate() *Command {
	c.addFlag(&Flag{
		ID:   "simulate",
		Flag: "--no-simulate",
		Args: nil,
	})
	return c
}

// Ignore "No video formats" error. Useful for extracting metadata even if the
// videos are not actually available for download (experimental)
//
// IgnoreNoFormatsError maps to cli flags: --ignore-no-formats-error.
func (c *Command) IgnoreNoFormatsError() *Command {
	c.addFlag(&Flag{
		ID:   "ignore_no_formats_error",
		Flag: "--ignore-no-formats-error",
		Args: nil,
	})
	return c
}

// Throw error when no downloadable video formats are found (default)
//
// NoIgnoreNoFormatsError maps to cli flags: --no-ignore-no-formats-error.
func (c *Command) NoIgnoreNoFormatsError() *Command {
	c.addFlag(&Flag{
		ID:   "ignore_no_formats_error",
		Flag: "--no-ignore-no-formats-error",
		Args: nil,
	})
	return c
}

// Do not download the video but write all related files (Alias: --no-download)
//
// SkipDownload maps to cli flags: --skip-download/--no-download.
func (c *Command) SkipDownload() *Command {
	c.addFlag(&Flag{
		ID:   "skip_download",
		Flag: "--skip-download",
		Args: nil,
	})
	return c
}

// GetUrl sets the "get-url" flag (no description specified).
//
// GetUrl maps to cli flags: -g/--get-url.
func (c *Command) GetUrl() *Command {
	c.addFlag(&Flag{
		ID:   "geturl",
		Flag: "--get-url",
		Args: nil,
	})
	return c
}

// GetTitle sets the "get-title" flag (no description specified).
//
// GetTitle maps to cli flags: -e/--get-title.
func (c *Command) GetTitle() *Command {
	c.addFlag(&Flag{
		ID:   "gettitle",
		Flag: "--get-title",
		Args: nil,
	})
	return c
}

// GetId sets the "get-id" flag (no description specified).
//
// GetId maps to cli flags: --get-id.
func (c *Command) GetId() *Command {
	c.addFlag(&Flag{
		ID:   "getid",
		Flag: "--get-id",
		Args: nil,
	})
	return c
}

// GetThumbnail sets the "get-thumbnail" flag (no description specified).
//
// GetThumbnail maps to cli flags: --get-thumbnail.
func (c *Command) GetThumbnail() *Command {
	c.addFlag(&Flag{
		ID:   "getthumbnail",
		Flag: "--get-thumbnail",
		Args: nil,
	})
	return c
}

// GetDescription sets the "get-description" flag (no description specified).
//
// GetDescription maps to cli flags: --get-description.
func (c *Command) GetDescription() *Command {
	c.addFlag(&Flag{
		ID:   "getdescription",
		Flag: "--get-description",
		Args: nil,
	})
	return c
}

// GetDuration sets the "get-duration" flag (no description specified).
//
// GetDuration maps to cli flags: --get-duration.
func (c *Command) GetDuration() *Command {
	c.addFlag(&Flag{
		ID:   "getduration",
		Flag: "--get-duration",
		Args: nil,
	})
	return c
}

// GetFilename sets the "get-filename" flag (no description specified).
//
// GetFilename maps to cli flags: --get-filename.
func (c *Command) GetFilename() *Command {
	c.addFlag(&Flag{
		ID:   "getfilename",
		Flag: "--get-filename",
		Args: nil,
	})
	return c
}

// GetFormat sets the "get-format" flag (no description specified).
//
// GetFormat maps to cli flags: --get-format.
func (c *Command) GetFormat() *Command {
	c.addFlag(&Flag{
		ID:   "getformat",
		Flag: "--get-format",
		Args: nil,
	})
	return c
}

// Quiet, but print JSON information for each video. Simulate unless --no-simulate
// is used. See "OUTPUT TEMPLATE" for a description of available keys
//
// DumpJson maps to cli flags: -j/--dump-json.
func (c *Command) DumpJson() *Command {
	c.addFlag(&Flag{
		ID:   "dumpjson",
		Flag: "--dump-json",
		Args: nil,
	})
	return c
}

// Force download archive entries to be written as far as no errors occur, even if
// -s or another simulation option is used (Alias: --force-download-archive)
//
// ForceWriteArchive maps to cli flags: --force-write-archive/--force-write-download-archive/--force-download-archive.
func (c *Command) ForceWriteArchive() *Command {
	c.addFlag(&Flag{
		ID:   "force_write_download_archive",
		Flag: "--force-write-archive",
		Args: nil,
	})
	return c
}

// Output progress bar as new lines
//
// Newline maps to cli flags: --newline.
func (c *Command) Newline() *Command {
	c.addFlag(&Flag{
		ID:   "progress_with_newline",
		Flag: "--newline",
		Args: nil,
	})
	return c
}

// Do not print progress bar
//
// NoProgress maps to cli flags: --no-progress.
func (c *Command) NoProgress() *Command {
	c.addFlag(&Flag{
		ID:   "noprogress",
		Flag: "--no-progress",
		Args: nil,
	})
	return c
}

// Show progress bar, even if in quiet mode
//
// Progress maps to cli flags: --progress.
func (c *Command) Progress() *Command {
	c.addFlag(&Flag{
		ID:   "noprogress",
		Flag: "--progress",
		Args: nil,
	})
	return c
}

// Display progress in console titlebar
//
// ConsoleTitle maps to cli flags: --console-title.
func (c *Command) ConsoleTitle() *Command {
	c.addFlag(&Flag{
		ID:   "consoletitle",
		Flag: "--console-title",
		Args: nil,
	})
	return c
}

// Print various debugging information
//
// Verbose maps to cli flags: -v/--verbose.
func (c *Command) Verbose() *Command {
	c.addFlag(&Flag{
		ID:   "verbose",
		Flag: "--verbose",
		Args: nil,
	})
	return c
}

// Write downloaded intermediary pages to files in the current directory to debug
// problems
//
// WritePages maps to cli flags: --write-pages.
func (c *Command) WritePages() *Command {
	c.addFlag(&Flag{
		ID:   "write_pages",
		Flag: "--write-pages",
		Args: nil,
	})
	return c
}

// LoadPages sets the "load-pages" flag (no description specified).
//
// LoadPages maps to cli flags: --load-pages.
func (c *Command) LoadPages() *Command {
	c.addFlag(&Flag{
		ID:   "load_pages",
		Flag: "--load-pages",
		Args: nil,
	})
	return c
}

// YoutubePrintSigCode sets the "youtube-print-sig-code" flag (no description specified).
//
// YoutubePrintSigCode maps to cli flags: --youtube-print-sig-code.
func (c *Command) YoutubePrintSigCode() *Command {
	c.addFlag(&Flag{
		ID:   "youtube_print_sig_code",
		Flag: "--youtube-print-sig-code",
		Args: nil,
	})
	return c
}

// Display sent and read HTTP traffic
//
// PrintTraffic maps to cli flags: --print-traffic/--dump-headers.
func (c *Command) PrintTraffic() *Command {
	c.addFlag(&Flag{
		ID:   "debug_printtraffic",
		Flag: "--print-traffic",
		Args: nil,
	})
	return c
}

// CallHome sets the "call-home" flag (no description specified).
//
// CallHome maps to cli flags: -C/--call-home.
func (c *Command) CallHome() *Command {
	c.addFlag(&Flag{
		ID:   "call_home",
		Flag: "--call-home",
		Args: nil,
	})
	return c
}

// NoCallHome sets the "no-call-home" flag (no description specified).
//
// NoCallHome maps to cli flags: --no-call-home.
func (c *Command) NoCallHome() *Command {
	c.addFlag(&Flag{
		ID:   "call_home",
		Flag: "--no-call-home",
		Args: nil,
	})
	return c
}
