module LogLevels

let message (logLine: string): string =
    match logLine.IndexOf ":" with
    | -1 -> logLine
    | (len) -> logLine[len+1..].Trim()

let logLevel(logLine: string): string =
    match logLine.IndexOf ":" with
    | -1 -> logLine
    | (len) -> logLine[..len-1].Trim([| '['; ']' |]).Trim().ToLower()

let reformat(logLine: string): string =
    $"{message logLine} ({logLevel logLine})"
