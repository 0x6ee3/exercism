module TisburyTreasureHunt

let getCoordinate (line: string * string): string =
    snd line

let convertCoordinate (coordinate: string): int * char =
    (int coordinate[0] - int '0', coordinate[1])

let compareRecords (azarasData: string * string) (ruisData: string * (int * char) * string) : bool =
    match azarasData, ruisData with
    | ((_, aC), (_, rC, _)) -> convertCoordinate aC = rC

let createRecord (azarasData: string * string) (ruisData: string * (int * char) * string) : (string * string * string * string) =
    match azarasData, ruisData with
    | ((treasure, coordinate), (location, _, quadrant)) when compareRecords azarasData ruisData -> (coordinate, location, quadrant, treasure)
    | _ -> ("", "", "", "")
